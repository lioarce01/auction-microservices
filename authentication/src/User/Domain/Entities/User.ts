import Base from "@Main/Domain/Entities/Base";
import UserDTO from "@User/Domain/DTOs/UserDTO";


export default class User extends Base
{
  constructor(
    public readonly sub: string,
    public readonly email: string,
    public readonly name: string,
    public readonly picture: string,
    public readonly roleId: number,
    createdAt: Date,
    updatedAt: Date,
    id?: string,
  )
  {
    super(createdAt, updatedAt, id);
  }

  getSub(): string
  {
    return this.sub;
  }

  getEmail(): string
  {
    return this.email;
  }

  getName(): string
  {
    return this.name;
  }

  getPicture(): string
  {
    return this.picture;
  }

  getRole(): number
  {
    return this.roleId;
  }

  public toDto(): UserDTO
  {
    return new UserDTO(
      this.id ?? '',
      this.sub,
      this.name,
      this.email,
      this.picture,
      this.roleId,
      this.createdAt,
      this.updatedAt,
    );
  }
}
