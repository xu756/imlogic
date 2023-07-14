import {fetchInitialData} from '@/services/initial';
import {RunTimeLayoutConfig} from "@@/plugin-layout/types";

export async function getInitialState() {
    return await fetchInitialData();
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