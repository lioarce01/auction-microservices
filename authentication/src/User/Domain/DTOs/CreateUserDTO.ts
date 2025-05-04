import { IsString, IsEmail, IsUrl, IsNumber } from 'class-validator';

class CreateUserDTO
{
  @IsString()
  sub!: string;

  @IsString()
  name!: string;

  @IsEmail()
  email!: string;

  @IsUrl()
  picture!: string;

  @IsNumber()
  roleId!: number;
}

export default CreateUserDTO;
