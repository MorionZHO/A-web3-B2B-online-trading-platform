import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Card, Image, Descriptions, Button, Modal, InputNumber, Select, Form, Input } from 'antd';
import './itemDetails.css';
import Navbar from '../../components/navbar/navbar';
import { fetchProductDetail } from '../../api/product';
import FormItem from 'antd/es/form/FormItem';

const { Option } = Select;

function ItemDetails() {
    const { id } = useParams();
    const [product, setProduct] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);

    // 打开 Modal 的方法
    const openModal = () => {
        setIsModalOpen(true);
    };

    // 关闭 Modal 的方法
    const closeModal = () => {
        setIsModalOpen(false);
    };

    let productTitle = 'Camera'

    let resProduct = [
        {
            key: '1',
            label: 'Merchant',
            children: 'Leica Co., Ltd.',
        },
        {
            key: '2',
            label: 'Stock quantity',
            children: '18',
        },
        {
            key: '3',
            label: 'Warehouse Address',
            children: 'Hangzhou, Zhejiang',
        },
        {
            key: '4',
            label: 'Description',
            span: 2,
            children: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod temp incididunt ut labore et dolore magna aliqua. Quis ipsum suspendisse. Donec condimentum elementum convallis. Nunc sed orci a diam ultrices aliquet interdum quis nulla.',
        },
        {
            key: '5',
            label: 'Price',
            children: '$120',
        },
    ];

    useEffect(() => {
        setProduct(resProduct)
        fetchProductDetail(id).then((res) => {
            // setProduct(res)
            console.log(res)
        });
    }, []);

    return (
        <>
            <Navbar></Navbar>
            <Dialogue isModalOpen={isModalOpen} closeModal={closeModal} product={product}></Dialogue>
            <div className='detail-wrapper'>
                <Card>
                    <Image.PreviewGroup
                        items={[
                            'https://gw.alipayobjects.com/zos/antfincdn/cV16ZqzMjW/photo-1473091540282-9b846e7965e3.webp',
                            'https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp',
                        ]}
                    >
                        <Image
                            style={{ width: 300 }}
                            src={`data:image/jpeg;base64,${product.imageData}`}
                        />
                    </Image.PreviewGroup>
                </Card>
                <Card>
                    <Descriptions title={productTitle} items={product} />
                    <Button type="primary" size='large' onClick={() => { openModal() }} className='purchase-button'>Buy
                    </Button>
                </Card>

            </div>
        </>)
}

export default ItemDetails;

function Dialogue({ isModalOpen, closeModal, product }) {
    const [confirmLoading, setConfirmLoading] = useState(false);
    const [modalText, setModalText] = useState('Please select the quantity you want to purchase');
    const [quantity, setQuantity] = useState(1);

    const [form] = Form.useForm();
    
    const handleOk = () => {
        setModalText('The modal will be closed after two seconds');
        setConfirmLoading(true);
        setTimeout(() => {
            setConfirmLoading(false);
            closeModal();
        }, 2000);
    };

    const onFinish = (values) => {
        console.log('Success:', values);
    };
    return (
        <>
            {/* <Button type="primary" onClick={showModal}>
                Open Modal with async logic
            </Button> */}
            <Modal
                title="Purchase Confirm"
                open={isModalOpen}
                onOk={handleOk}
                confirmLoading={confirmLoading}
                onCancel={closeModal}
            >
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
                    autoComplete="off"
                >
                    <FormItem
                        label='Price'
                    >
                        120 USD
                    </FormItem>

                    <FormItem
                        label='Quantity'
                        name='number'
                        rules={[
                            {
                                required: true,
                                message: 'Please select the quantity you want to purchase!',
                            },
                        ]}
                    >
                        <InputNumber min={1} max={18} defaultValue={1} onChange={(value) => { setQuantity(value) }} style={{ display: 'inline-block' }} />
                    </FormItem>
                    <Form.Item
                        label="Shipping address"
                        name="address"
                        rules={[
                            {
                                required: true,
                                message: 'Please input your address!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Password"
                        name="password"
                        rules={[
                            {
                                required: true,
                                message: 'Please input your password!',
                            },
                        ]}
                    >
                        <Input.Password />
                    </Form.Item>

                    <Form.Item
                        wrapperCol={{
                            offset: 8,
                            span: 16,
                        }}
                    >
                        <Button type="primary" htmlType="submit">
                            Submit
                        </Button>
                    </Form.Item>
                </Form>






            </Modal>
        </>
    );
}
