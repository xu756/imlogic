export default function (initialState: any) {
    const {userId, role} = initialState;
    console.log('global', initialState);

    return {
        canReadFoo: true,
        canUpdateFoo: role === 'admin',
    };
}