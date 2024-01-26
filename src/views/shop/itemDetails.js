import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Card, Image, Descriptions, Button, Modal, InputNumber, Select, Form, Input, Popconfirm } from 'antd';
import './itemDetails.css';
import Navbar from '../../components/navbar/navbar';
import { fetchProductDetail } from '../../api/product';
import FormItem from 'antd/es/form/FormItem';
import storage from '../../utils/storage';
import { submitOrder } from '../../api/order';
import { Carousel } from 'antd';
import image1 from '../../assets/img/banner_img_01.jpg';
import image2 from '../../assets/img/banner_img_02.jpg';
import image3 from '../../assets/img/banner_img_03.jpg';
import image4 from '../../assets/img/category_img_01.jpg';
import image5 from '../../assets/img/category_img_02.jpg';
import image6 from '../../assets/img/category_img_03.jpg';
import image7 from '../../assets/img/feature_prod_01.jpg';
import image8 from '../../assets/img/feature_prod_02.jpg';
import image9 from '../../assets/img/feature_prod_03.jpg';


// const { Option } = Select;

function ItemDetails() {
    const { id } = useParams();
    const [product, setProduct] = useState({});
    const [resProduct, setResProduct] = useState([])
    const [isModalOpen, setIsModalOpen] = useState(false);

    function updateInventory(newInventory) {
        setResProduct(resProduct.map((item) => {
            if (item.label === 'Inventory') {
                return {...item,children:newInventory}
            }
            return item;
        }))
    }

    // 打开 Modal 的方法
    const openModal = () => {
        setIsModalOpen(true);
    };

    // 关闭 Modal 的方法
    const closeModal = () => {
        setIsModalOpen(false);
    };


    useEffect(() => {
        setProduct(resProduct)
        fetchProductDetail(id).then((res) => {
            let ID = 1;
            let productDescribe = []
            for (let des in res) {
                if (des === 'merchantName') {
                    productDescribe.push({ key: `${ID}`, label: 'Merchant', children: res[des] });
                    ID++;
                }
                else if (des === 'inventory') {
                    productDescribe.push({ key: `${ID}`, label: 'Inventory', children: res[des] })
                    ID++;
                }
                else if (des === 'productName') {
                    productDescribe.push({ key: `${ID}`, label: 'Product Name', children: res[des].toUpperCase() })
                    ID++;
                }
                else if (des === 'productType') {
                    productDescribe.push({ key: `${ID}`, label: 'Product Type', children: res[des] })
                    ID++
                }
                else if (des === 'unitPrice') {
                    productDescribe.push({ key: `${ID}`, label: 'Unit Price', children: `${res[des]} SGD` })
                    ID++
                }
            }
            setProduct(res)
            setResProduct(productDescribe)
            console.log(res)
        });
    }, []);

    return (
        <>
            <Navbar></Navbar>
            <Dialogue isModalOpen={isModalOpen} closeModal={closeModal} product={product} update={updateInventory}></Dialogue>
            <div className='detail-wrapper'>
                <Card>

                    <Image
                        style={{ width: 300 }}
                        src={`data:image/jpeg;base64,${product.imageData}`}
                    />

                </Card>
                <Card>
                    <Descriptions title={product.productName} items={resProduct} />
                    <Button type="primary" size='large' onClick={() => { openModal() }} className='purchase-button'>Buy
                    </Button>
                </Card>

            </div>
            <Mycarousel></Mycarousel>
        </>)
}

export default ItemDetails;

function Dialogue({ isModalOpen, closeModal, product, update }) {
    const [modalText, setModalText] = useState('Please select the quantity you want to purchase');
    const [quantity, setQuantity] = useState(1);
    const [confirmLoading, setConfirmLoading] = useState(false);

    const [form] = Form.useForm();

    const localCache = storage.getItem('localCache') ? storage.getItem('localCache') : undefined;

    const onFinish = (values) => {
        console.log('Success:', values);
    };

    function getCurrentDateTime() {
        const now = new Date();

        const year = now.getFullYear();
        const month = String(now.getMonth() + 1).padStart(2, '0'); // 月份是从0开始的
        const day = String(now.getDate()).padStart(2, '0');

        const hours = String(now.getHours()).padStart(2, '0');
        const minutes = String(now.getMinutes()).padStart(2, '0');
        const seconds = String(now.getSeconds()).padStart(2, '0');

        return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    }

    function getRandomInt(min, max) {
        min = Math.ceil(min);
        max = Math.floor(max);
        return Math.floor(Math.random() * (max - min + 1)) + min;
    }

    function handleOrderSubmit() {
        setConfirmLoading(true);
        let address = form.getFieldValue('address')
        let date = getCurrentDateTime()
        let buyerToken = localCache['authToken']
        let companyName = product.merchantName
        let id = getRandomInt(1, 1000000)
        let unitPrice = product.unitPrice
        let totalPrice = unitPrice * quantity
        let number = quantity
        let orderState = 0

        let payload = {
            BuyerToken: buyerToken,
            SellerCompanyName: companyName,
            ProductName: product.productName,
            Address: address,
            OrderTime: date,
            Id: id.toString(),
            UnitPrice: unitPrice,
            TotalPrice: totalPrice,
            Number: number,
            OrderState: orderState
        }

        submitOrder(payload).then(res => {
            //loading界面
            if (res.code === 200) {
                let newInventory = res.number;
                update(newInventory)
                setConfirmLoading(false);
            }

        })
        closeModal()
    }
    return (
        <>
            <Modal
                confirmLoading={confirmLoading}
                title="Place an order"
                open={isModalOpen}
                onOk={handleOrderSubmit}
                onCancel={closeModal}
                destroyOnClose={true}
                okText={'Confirm'}
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
                    onFinish={onFinish}
                    autoComplete="off"

                >
                    <FormItem
                        label='Price'
                    >
                        {`${product.unitPrice} SGD`}
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
                        <InputNumber min={1} max={product.inventory} defaultValue={1} onChange={(value) => { setQuantity(value) }} style={{ display: 'inline-block' }} />
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

                </Form>
            </Modal>
        </>
    );
}


const imgStyle = {
    height: '500px',
    textAlign: 'center',
    marginLeft: 'auto',
    marginRight: 'auto',
    flex: '1'

};
function Mycarousel() {
    return (
        <section style={{ background: '#efefef' }}>
            <Carousel autoplay dotPosition='buttom'>

                <div >
                    <div><img src={image1} style={imgStyle}></img></div>
                </div>

                <div >
                    <img src={image2} style={imgStyle}></img>
                </div>
                <div >
                    <img src={image3} style={imgStyle}></img>
                </div>

                <div >
                    <img src={image4} style={imgStyle}></img>
                </div>

                <div >
                    <img src={image5} style={imgStyle}></img>
                </div>

                <div >
                    <img src={image6} style={imgStyle}></img>
                </div>

                <div >
                    <img src={image7} style={imgStyle}></img>
                </div>

                <div >
                    <img src={image8} style={imgStyle}></img>
                </div>

                <div >
                    <img src={image9} style={imgStyle}></img>
                </div>

            </Carousel>
        </section>
    )
}
