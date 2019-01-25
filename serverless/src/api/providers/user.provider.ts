import { User } from '../types/user';
import users from './users.json';

export class UserProvider {
    public getUsers(): User[] { 
        return users; 
    }

    public getUser(id: number): User {
        console.log(users);
        const jsonUser: any = users.find((user: User) => user.id == id);
        console.log(jsonUser);
        return jsonUser; 
    }

    public createUser(email: string): User {
        var ids: number[] = users.map(user => user.id);
        const nextId: number = Math.max(...ids) + 1;
        users.push({
            id: nextId,
            email: email,
            createdAt: new Date().toDateString(),
            authToken: "",
            firstName: "",
            lastName: "",
            middleName: "",
            isEmailVerified: false,
            mobile: "",
            dateOfBirth: "",
            companyName: "",
        });
        return this.getUser(nextId);
    }

    public deleteUser(id: number): boolean {
        const jsonUser: any = users.find(user => user.id === id);
        const index: number = users.indexOf(jsonUser, 0);
        if (index === -1) {
            return false;
        }
        users.splice(index, 1);
        const deletedUser: any = users.find(user => user.id === id);
        return deletedUser === null; 
    }
}
