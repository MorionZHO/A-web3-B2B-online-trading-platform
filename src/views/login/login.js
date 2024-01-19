import React from 'react';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Checkbox, Form, Input, Card, Select } from 'antd';
import './login.css';
import Navbar from '../../components/navbar/navbar'


function Login() {
    const onFinish = (values) => {
        console.log('Received values of form: ', values);
    };

    return (
        <>
            <Navbar></Navbar>
            <div className='login-wrapper'>
                <Card className='login-card'>
                    <div className='form-logo'>
                        Web3Trade
                    </div>
                    <Form
                        name="normal_login"
                        className="login-form"
                        initialValues={{
                            remember: true,
                        }}
                        onFinish={onFinish}
                    >
                        <Form.Item
                            name="Mobile Number"
                            rules={[
                                {
                                    required: true,
                                    message: 'Please input your Mobile Number!',
                                },
                            ]}
                        >
                            <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Mobile Number" />
                        </Form.Item>
                        <Form.Item
                            name="Password"
                            rules={[
                                {
                                    required: true,
                                    message: 'Please input your Password!',
                                },
                            ]}
                        >
                            <Input
                                prefix={<LockOutlined className="site-form-item-icon" />}
                                type="password"
                                placeholder="Password"
                            />
                        </Form.Item>
                        <Form.Item>
                            <Select placeholder="Role">
                                <Select.Option value="buyer">Buyer</Select.Option>
                                <Select.Option value="seller">Seller</Select.Option>
                            </Select>

                        </Form.Item>
                   

                        <Form.Item>
                            <Button type="primary" htmlType="submit" className="login-form-button">
                                Log in
                            </Button>
                            Or <a href="/register">register now!</a>
                        </Form.Item>
                    </Form>
                </Card>
            </div>
        </>
    );
};
export default Login;