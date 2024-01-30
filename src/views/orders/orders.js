import React, { useEffect, useState } from 'react';
import Navbar from '../../components/navbar/navbar'
import { Menu } from 'antd';
import Itemcard from '../shop/itemCard';
import { LikeOutlined, MessageOutlined, StarOutlined } from '@ant-design/icons';
import { Avatar, List, Space, Card } from 'antd';
import './orders.css'
import { getUserOrder } from '../../api/order';
import storage from '../../utils/storage';
const localCache = storage.getItem('localCache') ? storage.getItem('localCache') : undefined;

const data = Array.from({
    length: 25,
}).map((_, i) => ({
    title: `cup`,
    descrip:
        'Ant Design, a design language for background applications, is refined by Ant UED Team.',
    content:
        'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
}));

function Orders() {
    const [orderList, setOrderList] = useState([])
    let role = localCache['userInfo']['role']
    let payload = {}
    if (role === 'seller') {
        payload = {
            role: role,
            Token: localCache['authToken'],
            orderState: 0,
            companyName: localCache['userInfo']['companyName']
        }
    }
    else {
        payload = {
            role: role,
            Token: localCache['authToken'],
            orderState: 0,
        }
    }

    useEffect(() => {
        getUserOrder(payload).then(res => {
            console.log(res)
            setOrderList(res)
            
        })
    }, [])

    return (
        <>
            <Navbar></Navbar>
            <List
                itemLayout="vertical"
                size="large"
                className='list-item'
                dataSource={orderList}

                renderItem={(item) => (
                    <div className='list-wrapper'>
                        <Card className='list-card'>
                            <List.Item

                                key={item.Order.id}
                                extra={
                                    <img
                                        className='order-image'
                                        alt="logo"
                                        src={`data:image/jpeg;base64,${item.Imagedata}`}
                                    />
                                }
                            >
                                <List.Item.Meta
                                    style={{ marginTop: -10 }}
                                    title={<h2>{item.Order.productName.toUpperCase()}</h2>}
                                    description={<div><p>Order time: {item.Order.orderTime}</p><p><p>Product Provider: {item.Order.sellerCompanyName}</p></p></div>}
                                />

                                {`Purchase quantity: ${item.Order.number}`}

                                <br></br>
                                {`Unit price: ${item.Order.unitPrice} SGD`}<br></br>
                                {`Total price: ${item.Order.totalPrice} SGD`}<br></br>

                            </List.Item>
                        </Card>
                    </div>
                )}
            />
        </>
    )
}

export default Orders;