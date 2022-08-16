import { TypeOrmModuleOptions } from '@nestjs/typeorm';
import * as dotenv from 'dotenv';

dotenv.config();

class ConfigService {

  constructor(private env: { [k: string]: string | undefined }) { }

  /**
   * getValue is used to get value from .env file
   * @param key key for the value to retrieve
   * @param throwOnMissing condition if key not found
   * @returns value of the key
   */
  private getValue(key: string, throwOnMissing = true): string {
    const value = this.env[key];
    if (!value && throwOnMissing) {
      throw new Error(`config error - missing env.${key}`);
    }

    return value;
  }

  /**
   * ensureValues is used to get many value by array of key
   * @param keys array of key
   * @returns value of the keys
   */
  public ensureValues(keys: string[]) {
    keys.forEach(k => this.getValue(k, true));
    return this;
  }

  /**
   * getPort is used to get PORT value
   * @returns PORT value
   */
  public getPort() {
    return this.getValue('PORT', true);
  }

  /**
   * isProduction is used to get production mode value
   * @returns boolean (true=Production, false=Dev)
   */
  public isProduction() {
    const mode = this.getValue('MODE', false);
    return mode != 'DEV';
  }

  /**
   * getTypeOrmConfig is used to get typeorm config from .env file
   * @returns TypeOrmModuleOptions
   */
  public getTypeOrmConfig(): TypeOrmModuleOptions {
    return {
      type: 'postgres',
      host: this.getValue('POSTGRES_HOST'),
      port: parseInt(this.getValue('POSTGRES_PORT')),
      username: this.getValue('POSTGRES_USER'),
      password: this.getValue('POSTGRES_PASSWORD'),
      database: this.getValue('POSTGRES_DATABASE'),
      entities: ['dist/**/*.entity{.ts,.js}'],
      synchronize: Boolean(this.getValue("SYNC")),
      ssl: this.isProduction(),
      logging: true,
    };
  }

  /**
   * getMigrationConfig is used to get migration config from .env file
   * @returns TypeOrmModuleOptions
   */
  public getMigrationConfig(): TypeOrmModuleOptions {
    return {
      type: 'postgres',
      host: this.getValue('POSTGRES_HOST'),
      port: parseInt(this.getValue('POSTGRES_PORT')),
      username: this.getValue('POSTGRES_USER'),
      password: this.getValue('POSTGRES_PASSWORD'),
      database: this.getValue('POSTGRES_DATABASE'),
      entities: ['dist/**/*.entity{.ts,.js}'],
      synchronize: Boolean(this.getValue("SYNC")),
      migrations: ['src/migration/*{.ts,.js}'],
      migrationsTableName: 'migrations',
      ssl: this.isProduction(),
    };
  }
}

const configService = new ConfigService(process.env)
  .ensureValues([
    'POSTGRES_HOST',
    'POSTGRES_PORT',
    'POSTGRES_USER',
    'POSTGRES_PASSWORD',
    'POSTGRES_DATABASE',
    'SYNC',
  ]);

export { configService };