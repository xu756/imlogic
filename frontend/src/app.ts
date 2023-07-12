import {useModel} from '@umijs/max';

export function getInitialState() {
    return useModel('global');
}

export const layout = () => {
    return {
        logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
        menu: {
            locale: false,
        },
    };
};