import {defineConfig} from '@umijs/max';

export default defineConfig({
    antd: {},
    access: {},
    model: {},
    initialState: {},
    request: {},
    layout: {
        title: 'imlogic',

    },
    routes: [
        {
            path: '/',
            exact: false,
            component: '@/layouts/index',
            routes: [
                {
                    exact: true,
                    name: '首页',
                    path: '/home',
                    component: '@/pages/home'
                },
            ]
        },

        {
            exact: true,
            name: '登录',
            path: '/login',
            headerRender: false,
            // 不展示页脚
            footerRender: false,
            // 不展示菜单
            menuRender: false,
            // 不展示菜单顶栏
            menuHeaderRender: false,
            component: '@/pages/login'

        }

    ],
    npmClient: 'pnpm',
});
