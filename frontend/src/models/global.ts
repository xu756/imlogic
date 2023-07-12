// 全局共享数据示例
import {useEffect, useState} from 'react';

export default () => {
    const [config, setConfig] = useState({
        title: '云工桥安',
    });
    const [user, setUser] = useState({
        name: '张三',
        age: 18
    })

    useEffect(() => {
        console.log('global', global);
    }, []);

    return {
        config,
        user,
    };
}
