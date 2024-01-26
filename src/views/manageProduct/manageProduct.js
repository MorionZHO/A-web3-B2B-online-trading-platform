import React, { useState, useForm } from 'react';
import Navbar from '../../components/navbar/navbar'
import { PlusOutlined } from '@ant-design/icons';
import { Card } from 'antd';
import './manageProduct.css'
import { Link } from 'react-router-dom';
import image1 from '../../assets/img/feature_prod_01.jpg';
import image2 from '../../assets/img/feature_prod_02.jpg';
import image3 from '../../assets/img/feature_prod_03.jpg';
import { Button, Modal, Checkbox, Form, Input, Space, Select, InputNumber, Upload } from 'antd';

const onFinish = (values) => {
    console.log('Success:', values);
};
const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
};

const { Meta } = Card;
function ManageProduct() {
    const [form] = Form.useForm();
    const [imageUrl, setImageUrl] = useState();
    const [isModalOpen, setIsModalOpen] = useState(false);
    const showModal = () => {
        setIsModalOpen(true);
    };
    const handleOk = () => {
        setIsModalOpen(false);
    };
    const handleCancel = () => {
        setIsModalOpen(false);
    };
    const products = [
        {
            id: 1,
            productName: 'test',
            inventary: 200,
            unitPrice: 10,
            imageData: image1,
            merchantName: 'ZHH',
        },
        {
            id: 2,
            productName: 'test',
            inventary: 100,
            unitPrice: 10,
            imageData: image2,
            merchantName: 'ZHH'
        },
        {
            id: 3,
            productName: 'test',
            inventary: 200,
            unitPrice: 10,
            imageData: image3,
            merchantName: 'ZHH'
        },
        {
            id: 4,
            productName: 'test',
            inventary: 200,
            unitPrice: 10,
            imageData: image1,
            merchantName: 'ZHH'
        },
        {
            id: 5,
            productName: 'test',
            inventary: 200,
            unitPrice: 10,
            imageData: image2,
            merchantName: 'ZHH'
        },
        {
            id: 6,
            productName: 'test',
            inventary: 200,
            unitPrice: 10,
            imageData: image3,
            merchantName: 'ZHH'
        }

    ]

    const uploadButton = (
        <button
            style={{
                border: 0,
                background: 'none',
            }}
            type="button"
        >
            {/* {loading ? <LoadingOutlined /> : <PlusOutlined />} */}
            <PlusOutlined></PlusOutlined>
            <div
                style={{
                    marginTop: 8,
                }}
            >
                Upload
            </div>
        </button>
    );
    return (
        <>
            <Navbar></Navbar>

            <div className='product-card-container'>
                <Card hoverable={true} className='product-card' onClick={showModal}><PlusOutlined className='plus-icon' /></Card>
                {
                    products.map((product) => (
                        // 注意这里使用的是圆括号而不是花括号
                        <Card
                            key={product.id}
                            title={product.productName}
                            className='product-card'
                            hoverable
                            cover={<img alt="example" className='card-image' src={`${product.imageData}`} />}
                        >
                            <Meta description={`Inventory remaining: ${product.inventary}`} />
                        </Card>
                    ))
                }
            </div>
            <Modal title="Shelf New Product" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>

                <Form
                    form={form}
                    name="basic"
                    labelCol={{
                        span: 8,
                    }}
                    wrapperCol={{
                        span: 16,
                    }}
                    style={{
                        maxWidth: 600,
                    }}
                    initialValues={{
                        remember: true,
                    }}
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                    autoComplete="off"
                >
                    <Form.Item
                        label="Product Name"
                        name="productName"
                        rules={[
                            {
                                required: true,
                                message: 'Please input the product name!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="Unit Price"
                        name="unitPrice"
                        rules={[
                            {
                                required: true,
                                message: 'Please set the unit price!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label='Quantity'
                        name='number'
                        rules={[
                            {
                                required: true,
                                message: 'Please select the quantity you want to purchase!',
                            },
                        ]}
                    >
                        <InputNumber min={1} max={100000} defaultValue={1} style={{ display: 'inline-block' }} />
                    </Form.Item>
                    <Form.Item
                        label="Product Type"
                        name="productType"
                        rules={[
                            {
                                required: true,
                                message: 'Please select the type of your product!',
                            },
                        ]}

                    >
                        <Select
                            defaultValue="shoes"
                            style={{
                                width: 120,
                            }}
                            // onChange={handleChange}
                            options={[
                                {
                                    value: 'shoes',
                                    label: 'Shoes',
                                },
                                {
                                    value: 'clothes',
                                    label: 'Clothes',
                                },
                                {
                                    value: 'electronics',
                                    label: 'Electronics',
                                },
                                {
                                    value: 'homegoods',
                                    label: 'Home Goods',
                                },
                            ]}
                        />
                    </Form.Item>
                    <Form.Item
                        label="Picture Upload"
                        name="pictureUpload"
                        rules={[
                            {
                                required: true,
                                message: 'Please select the type of your product!',
                            },
                        ]}

                    >

                        <Upload
                            name="avatar"
                            listType="picture-card"
                            className="avatar-uploader"
                            showUploadList={false}
                            action="https://run.mocky.io/v3/435e224c-44fb-4773-9faf-380c5e6a2188"
                        // beforeUpload={beforeUpload}
                        // onChange={handleChange}
                        >
                            {imageUrl ? (
                                <img
                                    src={imageUrl}
                                    alt="avatar"
                                    style={{
                                        width: '100%',
                                    }}
                                />
                            ) : (
                                uploadButton
                            )}
                        </Upload>
                    </Form.Item>


                    {/* <Form.Item
                        wrapperCol={{
                            offset: 8,
                            span: 16,
                        }}
                    >
                        <Button type="primary" htmlType="submit">
                            Submit
                        </Button>
                    </Form.Item> */}
                </Form>
            </Modal>
        </>
    )
}

export default ManageProduct;