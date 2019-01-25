import { User } from '../types/user';
import jsonUsers from '../datasets/users.json';

export class UserProvider {
    public getUsers(): User[] { 
        return jsonUsers; 
    }

    public getUser(id: number): User {
        const jsonUser: any = jsonUsers.find((jsonUser: any) => jsonUser.id === id);
        return jsonUser;
    }

    public createUser(email: string): User {
        var ids: number[] = jsonUsers.map((jsonUser: any) => jsonUser.id);
        const nextId: number = Math.max(...ids) + 1;
        jsonUsers.push({
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
        const jsonUser: any = jsonUsers.find((jsonUser: any) => jsonUser.id === id);
        const index: number = jsonUsers.indexOf(jsonUser, 0);
        if (index === -1) {
            return false;
        }
        jsonUsers.splice(index, 1);
        const deletedUser: any = jsonUsers.find((jsonUser: any) => jsonUser.id === id);
        return deletedUser === undefined; 
    }
}
