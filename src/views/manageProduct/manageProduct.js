import React, { useState, useForm, useEffect } from 'react';
import Navbar from '../../components/navbar/navbar'
import { PlusOutlined, LoadingOutlined } from '@ant-design/icons';
import { Card } from 'antd';
import './manageProduct.css'
import { Button, Modal, Form, Input, Space, Select, InputNumber, Upload, message } from 'antd';
import storage from '../../utils/storage';
import { shelfProduct, getMyProducts, editProduct } from '../../api/product';

const localCache = storage.getItem('localCache') ? storage.getItem('localCache') : undefined;

const { Meta } = Card;
function ManageProduct() {
    const [form] = Form.useForm();
    const [editForm] = Form.useForm();
    const [imageUrl, setImageUrl] = useState();
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [loading, setLoading] = useState(false);
    const [imgBase64, setImageBase64] = useState("")
    const [quant, setQuant] = useState(1)
    const [productType, setProductType] = useState('')
    const [imageName, setImageName] = useState('')
    const [shelfedProducts, setShelfedProducts] = useState([])
    const [editable, setEditable] = useState(false)
    const [editedProduct, setEditedProduct] = useState({})
    const [editLoading, setEditLoading]=useState(false)

    useEffect(() => {
        let companyName = localCache['userInfo']['companyName']
        getMyProducts(companyName).then(res => {
            setShelfedProducts(res)
        })
    }, [])


    const onFinishFailed = (errorInfo) => {
        message.error('Failed:', errorInfo)
    };

    const showEditModal = (prod) => {
        setEditable(true)
        setEditedProduct(prod)
        console.log(prod)
    }

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
        setEditable(false)
    };


    const uploadButton = (
        <button
            style={{
                border: 0,
                background: 'none',
            }}
            type="button"
        >
            {loading ? <LoadingOutlined /> : <PlusOutlined />}

            <div
                style={{
                    marginTop: 8,
                }}
            >
                Upload
            </div>
        </button>
    );

    function handelUpload() {
        setLoading(true)
        const formValues = form.getFieldsValue();
        let payload = {
            MerchantName: localCache['userInfo']['companyName'],
            ProductName: formValues['productName'],
            UnitPrice: parseFloat(formValues['unitPrice']),
            ProductType: productType,
            Inventory: parseInt(quant),
            ImageName: imageName,
            ImageData: imgBase64,
            ProductState: 0
        }
        // console.log(payload)
        shelfProduct(payload).then(res => {
            if (res.code === 200) {
                message.success('The product was successfully put on the shelves!')
                getMyProducts(localCache['userInfo']['companyName']).then(res => {
                    setShelfedProducts(res)
                })
                setLoading(false)
                setIsModalOpen(false)
            }
        })

    }

    function handleUpdate() {
        setEditLoading(true)
        const updateValues = editForm.getFieldValue();
        let payload = {
            ProductName: editedProduct.productName,
            MerchantName: localCache['userInfo']['companyName'],
            UnitPrice: parseFloat(updateValues.newUnitPrice),
            Inventory: parseInt(updateValues.newInventory)
        }
        // console.log(payload)
        editProduct(payload).then(res => {
            if(res.code===200){
                setEditLoading(false)
                setEditable(false)
                getMyProducts(localCache['userInfo']['companyName']).then(res => {
                    setShelfedProducts(res)
                })
                console.log(res)
            }
            
        })

    }

    const fileToBase64 = (file) => {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();
            reader.onload = () => resolve(reader.result);
            reader.onerror = error => reject(error);
            reader.readAsDataURL(file);
        });
    };

    const getBase64 = (img, callback) => {
        const reader = new FileReader();
        reader.addEventListener('load', () => callback(reader.result));
        reader.readAsDataURL(img);
    };

    const handleUploadChange = (info) => {
        if (info.file.status === 'done' || info.file.status === 'uploading') {

            setImageName(info.file.name)
            fileToBase64(info.file.originFileObj)
                .then(base64 => {
                    let base64String = base64.split(',')[1]
                    setImageBase64(base64String)
                })
                .catch(error => {
                    console.log(error);
                });
            getBase64(info.file.originFileObj, (url) => {
                setLoading(false);
                setImageUrl(url);
            })
        }
    };

    const beforeUpload = (file) => {
        const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
        if (!isJpgOrPng) {
            message.error('You can only upload JPG/PNG file!');
        }

        return isJpgOrPng;
    };

    return (
        <>
            <Navbar></Navbar>

            <div className='product-card-container'>
                <Card
                    hoverable={true}
                    className='product-card'
                    onClick={showModal}>
                    <PlusOutlined className='plus-icon' />
                </Card>
                {
                    shelfedProducts.map((product) => (
                        // 注意这里使用的是圆括号而不是花括号
                        <Card
                            key={product.id}
                            title={product.productName}
                            className='product-card'
                            hoverable
                            onClick={() => { showEditModal(product) }}
                            cover={<img alt="example" className='card-image' src={`data:image/jpeg;base64,${product.imageData}`} />}
                        >
                            <Meta description={`Unit Price: ${product.unitPrice} SGD`} style={{ textAlign: 'center', marginBottom: '5px' }} />
                            <Meta description={`Inventory remaining: ${product.inventory}`} style={{ textAlign: 'center', marginBottom: '5px' }} />
                        </Card>
                    ))
                }
            </div>
            <Modal title="Edit Product" open={editable} onOk={handleUpdate} onCancel={handleCancel} destroyOnClose={true} confirmLoading={editLoading}>
                <Form
                    form={editForm}
                    name="edit"
                    labelCol={{
                        span: 8,
                    }}
                    wrapperCol={{
                        span: 16,
                    }}
                    style={{
                        maxWidth: 800,
                        marginTop: '30px'
                    }}
                    initialValues={{
                        newInventory: editedProduct.inventory,
                        newUnitPrice: editedProduct.unitPrice
                    }}

                    autoComplete="off"
                >
                    <Form.Item
                        label="Update the Inventory"
                        name="newInventory"
                        rules={[
                            {
                                required: true,
                                message: 'Please input new inventory of your product!',
                            },
                        ]}
                    >
                        <Input defaultValue={`${editedProduct.inventory}`} />
                    </Form.Item>

                    <Form.Item
                        label="New Unit Price"
                        name="newUnitPrice"
                        rules={[
                            {
                                required: true,
                                message: 'Please input your password!',
                            },
                        ]}
                    >
                        <Input defaultValue={`${editedProduct.unitPrice}`} />
                    </Form.Item>
                </Form>
            </Modal>
            <Modal title="Shelf New Product" open={isModalOpen} onOk={handelUpload} onCancel={handleCancel} destroyOnClose={true} confirmLoading={loading}>
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
                        maxWidth: 800,
                        marginTop: '30px'
                    }}
                    initialValues={{
                        remember: true,
                    }}
                    onFinish={handelUpload}
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
                        <InputNumber min={1} max={100000} defaultValue={1} style={{ display: 'inline-block' }} onChange={(value) => { setQuant(value) }} />
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
                            showSearch
                            style={{
                                width: 120,
                            }}
                            onChange={(value) => { setProductType(value) }}
                            options={[
                                {
                                    value: 'Shoes',
                                    label: 'Shoes',
                                },
                                {
                                    value: 'Clothes',
                                    label: 'Clothes',
                                },
                                {
                                    value: 'Electronics',
                                    label: 'Electronics',
                                },
                                {
                                    value: 'Home Goods',
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
                                message: 'Please upload a picture of your product!',
                            },
                        ]}

                    >
                        <Upload
                            name="avatar"
                            listType="picture-card"
                            className="avatar-uploader"
                            beforeUpload={beforeUpload}
                            onChange={handleUploadChange}
                            showUploadList={false}

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
                </Form>
            </Modal>
        </>
    )
}

export default ManageProduct;