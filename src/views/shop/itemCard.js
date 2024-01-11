import React from 'react';
import { Card } from 'antd';
import './itemCard.css'; // 引入自定义样式文件
import image1 from '../../assets/img/category_img_01.jpg'
import image2 from '../../assets/img/category_img_02.jpg'
import image3 from '../../assets/img/category_img_03.jpg'
import image4 from '../../assets/img/feature_prod_01.jpg'
import image5 from '../../assets/img/feature_prod_02.jpg'
import image6 from '../../assets/img/feature_prod_03.jpg'



const { Meta } = Card;
const data = [
    {
        title: 'Card title 1',
        src:image1
    },
    {
        title: 'Card title 2',
        src:image2
    },
    {
        title: 'Card title 3',
        src:image3
    },
    {
        title: 'Card title 4',
        src:image4
    },
    {
        title: 'Card title 5',
        src:image5
    },
    {
        title: 'Card title 6',
        src:image6
    },
   
    // ...更多的数据项
];

export default function Itemcard() {
    return (
        <div className="card-container">
            {data.map((item, index) => (
                <Card key={index} title={item.title} className='custom-card' hoverable cover={<img alt="example" className='card-image' src={item.src} />}>
                    <p>{item.description}</p>
                    <Meta title="Europe Street beat" description="www.instagram.com" />
                </Card>
            ))}
        </div>
    );
}
