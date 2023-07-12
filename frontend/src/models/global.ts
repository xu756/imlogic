// 全局共享数据示例
import {useEffect, useState} from 'react';

export default () => {
    const [config, setConfig] = useState({});
    const [user, setUser] = useState({})
    return {
        config,
        user,
    };
}
