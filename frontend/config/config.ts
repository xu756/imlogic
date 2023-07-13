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
    mfsu: false,
    routes: [
        // {
        //     path: '/',
        //     exact: false,
        //     redirect: '/home',
        //     component: '@/layouts',
        //     routes: [
        //
        //     ]
        // },
        {
            name: '首页',
            path: '/home',
            component: '@/pages/home',
            headerRender: true,
            access: 'user',
        },
        {
            exact: true,
            name: '登录',
            path: '/login',
            layout: false,
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
