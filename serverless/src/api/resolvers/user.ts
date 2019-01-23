export default {
    User: {
        id: user => user.id,
        authToken: user => user.authToken,
        createdAt: user => user.createdAt,
        firstName: user => user.firstName,
        lastName: user => user.lastName,
        middleName: user => user.middleName,
        email: user => user.email,
        isEmailVerified: user => user.isEmailVerified,
        mobile: user => user.mobile,
        dateOfBirth: user => user.dateOfBirth,
        companyName: user => user.companyName
    },
};
