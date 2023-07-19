import React, {useEffect} from 'react';
import {Button, Col, Form, Input, message, Row} from 'antd';
import './login.scss';
import Public from '@/services/public';

export default () => {
    const [form] = Form.useForm();
    const [messageApi, contextHolder] = message.useMessage();
    const [loading, setLoading] = React.useState(false);
    const [captcha, setCaptcha] = React.useState('');
    const [session_id, setSession_id] = React.useState('');
    const onFinish = (values: any) => {
        setLoading(true)
        Public.login(values.username, values.password, values.code, session_id).then((e) => {
            messageApi.success('登录成功')
        }).catch((e) => {
            console.log(e)
        }).finally(() => {
            init()
            setLoading(false)
        })

    };
    useEffect(() => {
        init()
    }, [])

    const init = () => {
        Public.getCaptcha().then((e) => {
            setCaptcha(e.img)
            setSession_id(e.session_id)
            form.resetFields(['code'])
        }, (e) => {
            console.log(e)
        })
    }
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
                    form={form}
                    className={"login-box"}
                    labelCol={{span: 5}}
                    initialValues={{username: 'admin', password: '123456'}}
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                    autoComplete="off"
                    labelAlign={"left"}

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
                    <Form.Item
                        label="验证码"
                        name="code"
                        rules={[{required: true, message: '请输入验证码!'}]}
                    >
                        <Row>
                            <Col onClick={init} span={12}><img src={captcha} width={"100%"}/></Col>
                            <Col span={10} push={2}> <Input/></Col>
                        </Row>
                    </Form.Item>

                    <Form.Item>
                        <Button loading={loading} block type="primary" htmlType="submit">
                            登录
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </>
    )
}