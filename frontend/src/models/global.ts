// 全局共享数据示例
import {useModel} from "@umijs/max";

export default () => {
    return useModel('@@initialState').initialState
}
