import {request} from "@@/plugin-request";

export default class UserServices {
    public static async login(username: string, password: string) {
        return request('/api/login', {
            method: 'POST',
            data: {
                username: username,
                password: password
            }
        })
    }

    static async register(username: string, password: string) {
        return request('/api/register', {
            method: 'POST',
            data: {
                username: username,
                password: password
            }
        })
    }

    static async logout() {
        return request('/api/logout', {
            method: 'POST'
        })
    }

    static async getUserInfo() {
        return request('/api/user/info', {
            method: 'GET'
        })
    }


}