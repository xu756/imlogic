import {RunTimeLayoutConfig} from "@@/plugin-layout/types";
import {RuntimeAntdConfig} from "@@/plugin-antd/types";
import {RequestConfig} from "@@/plugin-request/request";
import {message, notification, theme} from "antd";
import {useLocation} from "@@/exports";

export function getInitialState() {
    return {
        config: {
            title: '云工桥安',
        },
        user: {
            name: 'admin',
            JwtToken: '1234-5678-9101-1121',
            connId: '1234-5678-9101-1121',
            session_id: '1234-5678-9101-1121',
            role: 'SYSTEM_ADMIN',
            email: '756334744@qq.com',
        }
    }
}

export const layout: RunTimeLayoutConfig = () => {
    return {
        logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
        menu: {
            locale: false,
        },
        rightRender: (initialState: any) => {
            return (
                <>
                    <div>头像，退出按钮</div>
                </>
            )
        },

    };
};

export const antd: RuntimeAntdConfig = (memo: any) => {
    memo.theme ??= {};
    memo.theme.algorithm = theme.darkAlgorithm; // 配置 antd5 的预设 dark 算法

    memo.appConfig = {
        message: {
            // 配置 message 最大显示数，超过限制时，最早的消息会被自动关闭
            maxCount: 3,
        }
    }

    return memo;
};

//请求

// 错误处理方案： 错误类型
enum ErrorShowType {
    SILENT = 0,    // 不提示
    WARNMESSAGE = 1,   // 警告提示
    ERRORMESSAGE = 2,  // 报错提示
    NOTIFICATION = 3,   // 通知
    REDIRECT = 9,   // 页面跳转
}

// 与后端约定的响应数据格式
interface ResponseStructure {
    success: boolean;
    data: any;
    errorCode?: ErrorShowType;
    errorMessage?: string;
}


export const request: RequestConfig = {
    timeout: 1000,
    headers: {'X-Requested-With': 'XMLHttpRequest'},
    // 请求拦截器
    requestInterceptors: [
        (config: any) => {
            // 拦截请求配置，进行个性化处理。
            config.url = "/imlogic" + config.url
            // config.headers('Authorization', 'Bearer ' + localStorage.getItem('token'))
            return config
        }
    ],
    errorConfig: {
        // 错误抛出
        errorThrower: (res: ResponseStructure) => {
            const {success, data, errorCode, errorMessage} = res;
            if (!success || undefined) {
                const error: any = new Error(errorMessage);
                error.name = 'BizError';
                error.info = {success, data, errorCode, errorMessage};
                throw error; // 抛出自制的错误
            }
        },
        // 错误接收及处理

        errorHandler: (error: any, opts: any) => {
            console.error(error);
            if (opts?.skipErrorHandler) throw error;
            // 我们的 errorThrower 抛出的错误。
            if (error.response) {
                // Axios 的错误
                console.log(error.response.data);
                // 请求成功发出且服务器也响应了状态码，但状态代码超出了 2xx 的范围
                const {errorCode, errorMessage} = error.response.data;
                switch (errorCode) {
                    case ErrorShowType.SILENT:
                        // do nothing
                        break;
                    case ErrorShowType.WARNMESSAGE:
                        message.error(errorMessage);
                        break;
                    case ErrorShowType.ERRORMESSAGE:
                        message.error(errorMessage);
                        break;
                    case ErrorShowType.NOTIFICATION:
                        notification.open({
                            description: errorMessage,
                            message: errorCode,
                        });
                        break;
                    case ErrorShowType.REDIRECT:
                        // TODO: redirect
                        useLocation().pathname = '/login';
                        break;
                    default:
                        message.error(errorMessage);
                }
            } else if (error.request) {
                // 请求已经成功发起，但没有收到响应
                // \`error.request\` 在浏览器中是 XMLHttpRequest 的实例，
                // 而在node.js中是 http.ClientRequest 的实例
                message.error('None response! Please retry.');
            } else {
                // 发送请求时出了点问题
                message.error('Request error, please retry.');
            }
        },

    },

    // 响应拦截器
    responseInterceptors: [
        (response: any) => {
            // 拦截响应数据，进行个性化处理
            const data: ResponseStructure = response.data;
            if (!data.success) {
                message.error('请求失败！');
            }
            return response.data;
        }
    ]


}

