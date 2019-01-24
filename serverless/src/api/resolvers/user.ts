import { User } from '../types/user';

export default {
    User: {
        id: (user: User) => user.id,
        authToken: (user: User) => user.authToken,
        createdAt: (user: User) => user.createdAt,
        firstName: (user: User) => user.firstName,
        lastName: (user: User) => user.lastName,
        middleName: (user: User) => user.middleName,
        email: (user: User) => user.email,
        isEmailVerified: (user: User) => user.isEmailVerified,
        mobile: (user: User) => user.mobile,
        dateOfBirth: (user: User) => user.dateOfBirth,
        companyName: (user: User) => user.companyName
    },
};
