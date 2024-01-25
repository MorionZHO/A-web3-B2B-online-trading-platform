import React, { useState, useEffect } from 'react';
import { Card, Pagination } from 'antd';
import './itemCard.css'; // 引入自定义样式文件
import { Link } from 'react-router-dom';
import { getAllProducts, searchProduct, getProductByType } from '../../api/product';


const { Meta } = Card;

export default function Itemcard({ keyWord }) {
    const [products, setProducts] = useState([])
    useEffect(() => {
        if (typeof keyWord == "string") {
            console.log("catalogue")
            if (keyWord === 'All') {
                getAllProducts().then(res => {
                    setProducts(res);
                })
            }
            else {
                getProductByType(keyWord).then(res => {
                    setProducts(res)
                    
                })
            }
        }
        else {
            searchProduct(keyWord[1]).then(res => {
                setProducts(res)
                console.log(res)
            })
            console.log("searching")
        }

    }, [keyWord])


    return (
        <div className="card-container">
            {products.map((products) => (
                <Link className='link' to={`/shop/${products.id}`} key={products.id}>
                    <Card key={products.id} title={products.productName} className='custom-card' hoverable cover={<img alt="example" className='card-image' src={`data:image/jpeg;base64,${products.imageData}`} />}>
                        <p>{products.description}</p>
                        <Meta title={products.title} description={`provided by ${products.merchantName}`} />
                    </Card>
                </Link>
            ))}
        </div>
    );
}

