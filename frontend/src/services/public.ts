import {request} from "@@/plugin-request";


export default class Public {

    public static async getCaptcha() {

        return request('/public/get_captcha/get_code', {
            method: 'POST',
            data: {
                session_id: '123',
                sign: 'web',
                timestamp: new Date().getTime()
            }
        })
    }


    public static async login(username: string, password: string) {
        return request('/public/login/by_password', {
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