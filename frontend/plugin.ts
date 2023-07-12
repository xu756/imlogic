import type {IApi} from 'umi';

export default (api: IApi) => {
    api.onDevCompileDone((opts) => {
        opts;
        console.log('> onDevCompileDone', opts.isFirstCompile);
    });
};
