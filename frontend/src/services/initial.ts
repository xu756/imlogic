export const fetchInitialData = async () => {

    return {
        config: {
            title: '云工桥安',
        },
        user: {
            name: 'admin',
            JwtToken: '1234-5678-9101-1121',
            connId: '1234-5678-9101-1121',
            role: 'SYSTEM_ADMIN',
            email: '756334744@qq.com',
        }
    }
}