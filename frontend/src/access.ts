export default function (initialState) {
    const {userId, role} = initialState;

    return {
        canReadFoo: true,
        canUpdateFoo: role === 'admin',
        canDeleteFoo: (foo) => {
            return foo.ownerId === userId;
        },
    };
}