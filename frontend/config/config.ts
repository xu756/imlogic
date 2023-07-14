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
            flatMenu: true,
            layout: 'mix',
            component: '@/layouts/index',
            routes: [
                {path: '/', redirect: '/home'},
                {path: '/home', name: '首页', component: '@/pages/home', icon: 'HomeOutlined',},
                {path: '/chat', name: '聊天', component: '@/pages/chat', icon: 'TeamOutlined'},
            ]
        },
        {path: '/login', name: '登录', component: '@/pages/login', layout: false}

    ],
    npmClient: 'pnpm',
});
