import IUserRepository from '@User/Domain/Repositories/UserRepository';
import { container } from 'tsyringe';
import { RepoToken, UsecaseToken } from './Tokens/DITokens';
import PrismaUserRepository from '@User/Infrastructure/Repositories/PrismaUserRepository';
import ListUsersUseCase from '@User/Application/UseCases/List';
import GetOneUserUseCase from '@User/Application/UseCases/GetOne';
import GetByIdentifierUseCase from '@User/Application/UseCases/GetByIdentifier';
import UpdateUserUseCase from '@User/Application/UseCases/Update';
import DeleteUserUseCase from '@User/Application/UseCases/Delete';
import SaveUserUseCase from '@User/Application/UseCases/Save';
import GetMeUseCase from '@User/Application/UseCases/GetMe';


function setupContainer()
{
  console.log('📌 Registering dependencies...');

  container.registerSingleton<IUserRepository>(
    RepoToken.UserRepository,
    PrismaUserRepository,
  )

  // REGISTER USER USE CASES
  container.registerSingleton(UsecaseToken.User.ListUsers, ListUsersUseCase);
  container.registerSingleton(UsecaseToken.User.GetOneUser, GetOneUserUseCase);
  container.registerSingleton(UsecaseToken.User.GetByIdentifier, GetByIdentifierUseCase);
  container.registerSingleton(UsecaseToken.User.UpdateUser, UpdateUserUseCase);
  container.registerSingleton(UsecaseToken.User.DeleteUser, DeleteUserUseCase);
  container.registerSingleton(UsecaseToken.User.SaveUser, SaveUserUseCase);
  container.registerSingleton(UsecaseToken.User.GetMe, GetMeUseCase);

  console.log('✅ Dependencies registered successfully');
}

export default setupContainer;
