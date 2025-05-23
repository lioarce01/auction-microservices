import { User as PrismaUser } from '@prisma/client';
import User from '@User/Domain/Entities/User';

class UserTransformer
{
  static toDomain(
    user: PrismaUser,
  ): User
  {
    return new User(
      user.sub,
      user.email,
      user.name,
      user.picture,
      user.roleId,
      user.createdAt,
      user.updatedAt,
      user.id,
    );
  }
}

export default UserTransformer;
