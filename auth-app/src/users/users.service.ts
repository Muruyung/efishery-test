import { Logger, HttpException, HttpStatus, Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository } from 'typeorm';
import * as bcrypt from 'bcrypt';
import { CreateUserInput } from './dto/create-user.input';
import { Users } from './entities/user.entity';

@Injectable()
export class UsersService {
  private readonly logger = new Logger(UsersService.name);
  
  constructor(
    @InjectRepository(Users) private usersRepository: Repository<Users>,
    private connection: DataSource,
  ) {}

  /**
   * create is used to create new user
   * @param createUserInput the field that will be used to user input when register
   * @returns user data
   */
  async create(createUserInput: CreateUserInput) {
    return await this.connection.transaction(async manager => {
      const isExist = await manager.findOneBy(Users, { phoneNumber: createUserInput.phoneNumber });
      if (isExist) {
        throw new HttpException('User already exist', HttpStatus.BAD_REQUEST);
      }
      const salt = await bcrypt.genSalt(11);
      createUserInput.password = await bcrypt.hash(createUserInput.password, salt);
      const user = this.usersRepository.create({
        password: createUserInput.password,
        phoneNumber: createUserInput.phoneNumber,
        name: createUserInput.name,
        role: createUserInput.role,
      });
      await manager.save(Users, user)

      return user;
    });
  }

  /**
   * findOne is used to get user data by id
   * @param id user id
   * @returns user data
   */
  async findOne(id: string) {
    var user = await this.usersRepository.findOneBy({ id });
    if (!user) {
      throw new NotFoundException(`User #${id} not found`);
    }
    return user;
  }

  /**
   * findByLogin is used to get user data by login
   * @param phoneNumber user phone number
   * @param password user password
   * @returns user data
   */
  async findByLogin(phoneNumber: string, password: string) {
    const user = await this.usersRepository.findOneBy({ phoneNumber });
    if (!user) {
      throw new NotFoundException(`User with phone number ${phoneNumber} not found`);
    }

    const isIdentic = await bcrypt.compare(password, user.password);
    if (!isIdentic) {
      throw new HttpException(`Invalid credentials`, HttpStatus.UNAUTHORIZED);
    }
    return user;
  }
}
