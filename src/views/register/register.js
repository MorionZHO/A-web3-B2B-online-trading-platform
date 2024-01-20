import React, { useState } from 'react';
import {
    Button,
    Form,
    Input,
    Select,
    Card
} from 'antd';
import './register.css'
import { handleRegister } from '../../api/auth';
import Navbar from '../../components/navbar/navbar';


const { Option } = Select;
const formItemLayout = {
    labelCol: {
        xs: {
            span: 24,
        },
        sm: {
            span: 10,
        },
    },
    wrapperCol: {
        xs: {
            span: 24,
        },
        sm: {
            span: 16,
        },
    },
};
const tailFormItemLayout = {
    wrapperCol: {
        xs: {
            span: 24,
            offset: 0,
        },
        sm: {
            span: 16,
            offset: 8,
        },
    },
};
function Register() {
    const [form] = Form.useForm();

    const onFinishRegister = (values) => {
        console.log('Received values of form: ', values);
        let payload= new FormData()
        let phoneNumber=`+${values.prefix}${values.phoneNumber}`
        payload.append("phoneNumber",phoneNumber)
        payload.append("password",values.password)
        payload.append("role",values.role)
        if(values.role==='seller'){
            payload.append("companyName",values.companyName)
        }
        handleRegister(payload).then(res=>{
            if(res.code===200){
                console.log(res)
            }
        })
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
                    <div className="register-logo">
                        Web3Trade
                    </div>
                    <Form
                        {...formItemLayout}
                        form={form}
                        name="register"
                        onFinish={onFinishRegister}
                        initialValues={{
                            prefix: '65',
                        }}
                        style={{
                            width: '100%',
                        }}
                        scrollToFirstError
                    >
                        <Form.Item
                            name="phone"
                            label="Phone Number"
                            rules={[
                                {
                                    required: true,
                                    message: 'Please input your phone number!',
                                },
                            ]}
                        >
                            <Input
                                addonBefore={prefixSelector}
                                style={{
                                    width: '100%',
                                }}
                            />
                        </Form.Item>

                        <Form.Item
                            name="password"
                            label="Password"
                            rules={[
                                {
                                    required: true,
                                    message: 'Please input your password!',
                                },
                            ]}
                            hasFeedback
                        >
                            <Input.Password />
                        </Form.Item>

                        <Form.Item
                            name="confirm"
                            label="Confirm Password"
                            dependencies={['password']}
                            hasFeedback
                            rules={[
                                {
                                    required: true,
                                    message: 'Please confirm your password!',
                                },
                                ({ getFieldValue }) => ({
                                    validator(_, value) {
                                        if (!value || getFieldValue('password') === value) {
                                            return Promise.resolve();
                                        }
                                        return Promise.reject(new Error('The new password that you entered do not match!'));
                                    },
                                }),
                            ]}
                        >
                            <Input.Password />
                        </Form.Item>

                        <Form.Item name="role"
                            label="Role"
                            hasFeedback
                            rules={[
                                {
                                    required: true,
                                    message: 'Please select the role that you want to register!',
                                },
                            ]}
                        >
                            <Select placeholder="Role">
                                <Select.Option value="buyer">Buyer</Select.Option>
                                <Select.Option value="seller">Seller</Select.Option>
                            </Select>
                        </Form.Item>

                        <Form.Item
                            shouldUpdate={(prevValues, currentValues) => prevValues.select !== currentValues.select}//条件渲染
                            noStyle
                        >
                            {({ getFieldValue }) =>
                                getFieldValue('select') === 'seller' ? (
                                    <Form.Item name='companyName'
                                        label="Company Name"
                                        tooltip="Please enter the full name of your registered company."
                                        rules={[
                                            {
                                                required: true,
                                                message: 'Please enter the full name of your registered company.',
                                            },
                                        ]}
                                    >
                                        <Input />
                                    </Form.Item>
                                ) : null
                            }
                        </Form.Item>

                        <Form.Item {...tailFormItemLayout}>
                            <Button type="primary" htmlType="submit">
                                Register
                            </Button>
                        </Form.Item>
                    </Form>
                </Card>
            </div>
        </>
    );
};
export default Register;