import Criteria from '@Main/Infrastructure/Criteria/Criteria';
import { RepoToken } from '@Shared/DI/Tokens/DITokens';
import User from '@User/Domain/Entities/User';
import IUserRepository from '@User/Domain/Repositories/UserRepository';
import { inject, injectable } from 'tsyringe';


@injectable()
class ListUsersUseCase
{
  constructor(
    @inject(RepoToken.UserRepository) private userRepository: IUserRepository,
  )
  { }

  async execute(criteria: Criteria): Promise<User[] | []>
  {

    const result = await this.userRepository.list(criteria);

    return result;
  }
}

export default ListUsersUseCase;
