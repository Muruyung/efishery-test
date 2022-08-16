import { BadRequestException, Injectable, UnprocessableEntityException } from "@nestjs/common";
import { JwtService } from "@nestjs/jwt";
import { InjectRepository } from "@nestjs/typeorm";
import { TokenExpiredError } from "jsonwebtoken";
import { Repository } from "typeorm";
import { Users } from "../users/entities/user.entity";
import { UsersService } from "../users/users.service";
import { RefreshToken } from "./entities/refresh-token.entity";
import { RegisterUserBody } from "./dto/register-user.body";
import { CreateUserInput } from "src/users/dto/create-user.input";

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
    @InjectRepository(RefreshToken)
    private refreshTokenRepository: Repository<RefreshToken>,
  ) {}
  
  /**
   * validateUser is used to validate user login input
   * @param phoneNumber user phone number
   * @param pass user password
   * @returns user data
   */
  async validateUser(phoneNumber: string, pass: string) {
    return await this.usersService.findByLogin(phoneNumber, pass);
  }
  
  /**
   * generateAccessToken is used to generate new access token
   * @param user user data (id)
   * @returns generated access token
   */
  async generateAccessToken(user: Pick<Users, "id">) {
    const payload = { sub: String(user.id) };
    return await this.jwtService.signAsync({
      ...payload,
      exp: Math.floor(Date.now()/1000) + Number(process.env.TOKEN_EXPIRED),
    });
  }

  /**
   * createRefreshToken is used to save refresh token into database
   * @param user user data (id)
   * @param expire token expired time
   * @returns refresh token data
   */
  async createRefreshToken(user: Pick<Users, "id">, expire: number) {
    const expiration = new Date();
    expiration.setTime(expiration.getTime() + expire);

    const token = this.refreshTokenRepository.create({
      user,
      expires: expiration,
    });

    return await this.refreshTokenRepository.save(token);
  }

  /**
   * generateRefreshToken is used to generate new refresh token
   * @param user user data (id)
   * @returns generated refresh token
   */
  async generateRefreshToken(user: Pick<Users, "id">) {
    const payload = { sub: String(user.id) };
    const expiresIn = Number(process.env.REFRESH_TOKEN_EXPIRED)
    const token = await this.createRefreshToken(user, expiresIn*1000);
    return await this.jwtService.signAsync({
      ...payload,
      jwtId: String(token.id),
      exp: Math.floor(Date.now()/1000) + expiresIn,
    });
  }

  /**
   * revokeRefreshToken is used to revoke refresh token so that they cannot be reused by users
   * @param refresh refresh token to be revoked
   * @returns 
   */
  async revokeRefreshToken(refresh: string) {
    const { user } = await this.resolveRefreshToken(refresh);
    const payload = await this.jwtService.verify(refresh);

    const refreshToken = await this.refreshTokenRepository.preload({
      id: payload.jwtId,
      revoked: true,
    })
    this.refreshTokenRepository.save(refreshToken)
    return "Refresh token has been revoked."
  }

  /**
   * resolveRefreshToken is used to check if the refresh token is valid
   * @param refresh refresh token to be checked
   * @returns user and token data if valid
   */
  async resolveRefreshToken(refresh: string) {
    try {
      const payload = await this.jwtService.verify(refresh);

      if (!payload.sub || !payload.jwtId) {
        throw new UnprocessableEntityException("Refresh token malformed");
      }

      const token = await this.refreshTokenRepository.findOneBy({
        id: payload.jwtId,
      });

      if (!token) {
        throw new UnprocessableEntityException("Refresh token not found");
      }

      if (token.revoked) {
        throw new UnprocessableEntityException("Refresh token revoked");
      }

      const user = await this.usersService.findOne(payload.sub);

      if (!user) {
        throw new UnprocessableEntityException("Refresh token malformed");
      }

      return { user, token };
    } catch (e) {
      if (e instanceof TokenExpiredError) {
        throw new UnprocessableEntityException("Refresh token expired");
      } else {
        throw new UnprocessableEntityException("Refresh token malformed");
      }
    }
  }

  /**
   * resolveAccessToken is used to check if the access token is valid
   * @param access access token to be checked
   * @returns payload token
   */
  async resolveAccessToken(access: string) {
    try {
      const payload = await this.jwtService.verify(access);
      console.log(payload)
      
      if (!payload.sub || payload.jwtId) {
        throw new UnprocessableEntityException("Access token malformed");
      }

      const user = await this.usersService.findOne(payload.sub);

      if (!user) {
        throw new UnprocessableEntityException("Access token malformed");
      }

      return { payload };
    } catch (e) {
      if (e instanceof TokenExpiredError) {
        throw new UnprocessableEntityException("Access token expired");
      } else {
        throw new UnprocessableEntityException("Access token invalid");
      }
    }
  }

  /**
   * createAccessTokenFromRefreshToken is used to create new access token from current refresh token
   * @param refresh refresh token
   * @returns user and token data
   */
  async createAccessTokenFromRefreshToken(refresh: string) {
    const { user } = await this.resolveRefreshToken(refresh);

    const token = await this.generateAccessToken(user);

    return { user, token };
  }

  /**
   * register is used to create new user
   * @param registerInput the field that will be used to user input when register
   * @returns user data
   */
  async register(registerInput: RegisterUserBody) {
    var createUser = new CreateUserInput;
    var password = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";

    for (var i = 0; i < 4; i++)
      password += possible.charAt(Math.floor(Math.random() * possible.length));

    createUser.password = password;
    createUser.name = registerInput.name;
    createUser.phoneNumber = registerInput.phoneNumber;
    createUser.role = registerInput.role;

    return await this.usersService.create(createUser);
  }
}
