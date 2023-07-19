export interface InitialData {
    config: {
        title: string
    },
    user: {
        Id: number,
        name: string,
        JwtToken: string,
        Avatar: string,
        connId: string,
        role: Array<Number>,
        group: Array<Number>,
    }
}

export const fetchInitialData = async (): Promise<InitialData> => {
    return {
        config: {
            title: '云工桥安',
        },
        user: {
            Id: 1,
            name: 'admin',
            JwtToken: '1234-5678-9101-1121',
            Avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
            connId: '1234-5678-9101-1121',
            role: [1],
            group: [1],
        }
    }
}