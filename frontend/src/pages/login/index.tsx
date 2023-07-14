import React, {useEffect} from 'react';
import {Button, Checkbox, Form, Input, message} from 'antd';
import './login.scss';
import Public from '@/services/public';

export default () => {
    const [messageApi, contextHolder] = message.useMessage();
    const [loading, setLoading] = React.useState(false);
    const onFinish = (values: any) => {
        console.log('Success:', values);
        setLoading(true)
        Public.login(values.username, values.password).then((e) => {
            console.log(e)
            setLoading(false)
        })
    };
    useEffect(() => {
        Public.getCaptcha().then((e) => {
            console.log(e)
        }, (e) => {
            console.log(e)
        })
    }, [])
    const onFinishFailed = (errorInfo: any) => {
        for (let i in errorInfo.errorFields) {
            messageApi.error(errorInfo.errorFields[i].errors[0]).then(r =>
                console.log(r)
            );
        }
    };


    return (
        <>
            {contextHolder}
            <div className="login">
                <Form
                    className={"login-box"}
                    labelCol={{span: 4, pull: 1}}
                    initialValues={{remember: true, username: 'admin', password: '123456'}}
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                    autoComplete="off"
                >
                    <Form.Item
                        label="用户名"
                        name="username"
                        rules={[{required: true, message: '请输入用户名!'}]}
                    >
                        <Input/>
                    </Form.Item>

                    <Form.Item
                        label="密码"
                        name="password"
                        rules={[{required: true, message: '请输入密码!'}]}
                    >
                        <Input.Password/>
                    </Form.Item>

                    <Form.Item name="remember" valuePropName="checked" wrapperCol={{offset: 16, span: 8}}>
                        <Checkbox>记住我</Checkbox>
                    </Form.Item>

                    <Form.Item>
                        <Button loading={loading} block type="primary" htmlType="submit">
                            Submit
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </>
    )
}