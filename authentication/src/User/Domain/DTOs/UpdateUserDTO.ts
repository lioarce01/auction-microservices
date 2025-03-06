import { IsString, IsEmail, IsUrl, IsNumber } from 'class-validator';

class UpdateUserDTO
{
  @IsString()
  name?: string;

  @IsEmail()
  email?: string;

  @IsUrl()
  picture?: string;
}

export default UpdateUserDTO;
