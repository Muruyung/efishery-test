import { Body, ClassSerializerInterceptor, Controller, Post, UseInterceptors, ValidationPipe } from "@nestjs/common";
import { UserInputError } from "apollo-server-express";
import { CreateUserInput } from "src/users/dto/create-user.input";
import { UserObject } from "src/users/dto/user.object";
import { UserDto } from "../users/dto/user.dto";
import { AuthService } from "./auth.service";
import { LoginUserBody } from "./dto/login-user.body";
import { LoginUserResponse } from "./dto/login-user.response";
import { AccessTokenBody } from "./dto/access-token.body";
import { RefreshTokenBody } from "./dto/refresh-token.body";
import { RefreshTokenResponse } from "./dto/refresh-token.response";
import { RegisterUserBody } from "./dto/register-user.body";
import { RegisterUserResponse } from "./dto/register-user.response";

@Controller("auth")
export class AuthController {
  constructor(private authService: AuthService) {}

  /**
   * login is used to log in user
   * @param loginUserBody body request is phone_number and password
   * @returns user and generated token data
   */
  @Post("login")
  @UseInterceptors(ClassSerializerInterceptor)
  async login(@Body(new ValidationPipe({ transform: true })) loginUserBody: LoginUserBody) {
    const user = await this.authService.validateUser(loginUserBody.phoneNumber, loginUserBody.password)
    if (!user) {
      return new UserInputError("phone number or password incorrect.");
    }

    const accessToken = await this.authService.generateAccessToken(user);
    const refreshToken = await this.authService.generateRefreshToken(user);

    const payload = new LoginUserResponse();
    payload.user = new UserObject(user);
    payload.accessToken = accessToken;
    payload.refreshToken = refreshToken;

    return payload;
  }

  /**
   * revoke is used to revoke refresh token so that they cannot be reused by users
   * @param refreshTokenBody refresh token to be revoked
   * @returns 
   */
  @Post("revoke")
  @UseInterceptors(ClassSerializerInterceptor)
  async revoke(@Body(new ValidationPipe({ transform: true })) refreshTokenBody: RefreshTokenBody) {
    const revoke = await this.authService.revokeRefreshToken(refreshTokenBody.refreshToken)
    return revoke;
  }

  /**
   * resolveAccessToken is used to check if the access token is valid
   * @param accessTokenBody access token to be checked
   * @returns payload token
   */
  @Post("access-token-check")
  @UseInterceptors(ClassSerializerInterceptor)
  async resolveAccessToken(@Body(new ValidationPipe({ transform: true })) accessTokenBody: AccessTokenBody) {
    const payload = await this.authService.resolveAccessToken(accessTokenBody.accessToken)
    return payload;
  }

  /**
   * refresh is used to generate new access token by refresh token
   * @param refreshInput refresh token
   * @returns user and generated access token data
   */
  @Post("refresh")
  @UseInterceptors(ClassSerializerInterceptor)
  async refresh(@Body(new ValidationPipe({ transform: true })) refreshInput: RefreshTokenBody) {
    const { user, token } = await this.authService.createAccessTokenFromRefreshToken(refreshInput.refreshToken);

    const payload = new RefreshTokenResponse();
    payload.user = new UserDto(user);
    payload.accessToken = token;

    return payload;
  }

  /**
   * register is used to create new user
   * @param registerInput the field that will be used to user input when register
   * @returns user and generated token data
   */
  @Post("register")
  @UseInterceptors(ClassSerializerInterceptor)
  async register(@Body(new ValidationPipe({ transform: true })) registerInput: RegisterUserBody) {
    
    const user = await this.authService.register(registerInput);

    const accessToken = await this.authService.generateAccessToken(user);
    const refreshToken = await this.authService.generateRefreshToken(user);

    const payload = new RegisterUserResponse();
    payload.user = new UserDto(user);
    payload.accessToken = accessToken;
    payload.refreshToken = refreshToken;

    return payload;
  }
}
