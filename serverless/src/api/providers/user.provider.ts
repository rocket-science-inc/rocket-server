import { User } from '../types/user';
import users from './users.json';

export class UserProvider {
    public getUsers(): User[] { 
        return users; 
    }
    public getUser(id: number): User {
        const user: User = users.find((user: User) => user.id === id);
        return user; 
    }
}
