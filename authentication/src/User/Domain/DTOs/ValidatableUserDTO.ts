import { IsString, IsEmail, IsDate, IsNumber } from 'class-validator';
import UserDTO from './UserDTO';

class ValidatableUserDto extends UserDTO
{
  @IsString()
  declare id: string;

  @IsString()
  declare sub: string;

  @IsString()
  declare name: string;

  @IsEmail()
  declare email: string;

  @IsString()
  declare picture: string;

  @IsNumber()
  declare role: number;

  @IsDate()
  declare createdAt: Date;

  @IsDate()
  declare updatedAt: Date;
}

export default ValidatableUserDto;
