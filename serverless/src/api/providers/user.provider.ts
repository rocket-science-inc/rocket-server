import { Injectable } from '@graphql-modules/di';
import users from './users.json';

@Injectable()
export class UserProvider {
    getUsers() { 
        console.log(users);
        return users; 
    }

    getUser(id: number) {
        const user = users.find(({ id }) => id === id);
        console.log(user);
        return user; 
    }
}