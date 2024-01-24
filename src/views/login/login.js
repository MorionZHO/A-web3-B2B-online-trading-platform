import React from 'react';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Form, Input, Card, Select, message ,Spin} from 'antd';
import './login.css';
import Navbar from '../../components/navbar/navbar'
import { handleLogin } from '../../api/auth';
import storage from '../../utils/storage';

const { Option } = Select;

function Login() {
    const [spinning, setSpinning] = React.useState(false);

    const showLoader = () => {
        setSpinning(true);
      };

    const onFinishLogin = (values) => {
        showLoader();
        const payload = {
            phoneNumber: `+${values.prefix}${values.mobileNumber}`,
            password: values.password,
            role: values.role
        }
        handleLogin(payload).then((res) => {
            if(res.code===200){
                message.success("login success!")
                console.log(res)
                let localCache={}
                localCache['authToken']=res.token
                localCache['userInfo']=res.userInfo
                storage.setItem('localCache',localCache)
                window.location.href='/'
            }
            
        });

    };

    const prefixSelector = (
        <Form.Item name="prefix" noStyle>
            <Select
                style={{
                    width: 70,
                }}
            >
                <Option value="86">+86</Option>
                <Option value="65">+65</Option>
            </Select>
        </Form.Item>
    );

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
                            prefix: '65'
                        }}
                        onFinish={onFinishLogin}
                    >
                        <Form.Item
                            name="mobileNumber"
                            rules={[
                                {
                                    required: true,
                                    message: 'Please input your Mobile Number!',
                                },
                            ]}
                        >
                            <Input addonBefore={prefixSelector} prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Mobile Number" />
                        </Form.Item>
                        <Form.Item
                            name="password"
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
                        <Form.Item name='role'
                            rules={[
                                {
                                    required: true,
                                    message: 'Please select your role!',
                                },
                            ]}
                        >
                            <Select placeholder="Role">
                                <Select.Option value="buyer">Buyer</Select.Option>
                                <Select.Option value="seller">Seller</Select.Option>
                            </Select>

                        </Form.Item>


                        <Form.Item>
                            <Button type="primary" htmlType="submit" className="login-form-button" >
                                <Spin spinning={spinning}/>
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