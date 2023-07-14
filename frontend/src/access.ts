export default function (initialState: any) {
    const {config, user} = initialState;
    console.log(user.JwtToken)
    return {
        user: true
    };
}