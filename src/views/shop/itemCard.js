import React,{useState,useEffect} from 'react';
import { Card } from 'antd';
import './itemCard.css'; // 引入自定义样式文件
import { Link } from 'react-router-dom';
import image1 from '../../assets/img/category_img_01.jpg'
import image2 from '../../assets/img/category_img_02.jpg'
import image3 from '../../assets/img/category_img_03.jpg'
import image4 from '../../assets/img/feature_prod_01.jpg'
import image5 from '../../assets/img/feature_prod_02.jpg'
import image6 from '../../assets/img/feature_prod_03.jpg'




const { Meta } = Card;
const data = [
    {
        id: 1,
        imageData:image1,
        merchantName:"ZHH",
        productName:"plastic bag",
        productType:"",
        unitPrice:6
    },
    {
        id: 2,
        imageData:image2,
        merchantName:"ZHH",
        productName:"hook up",
        productType:"",
        unitPrice:6
    },
    {
        id: 3,
        imageData:image3,
        merchantName:"ZHH",
        productName:"dish tower",
        productType:"",
        unitPrice:6
    },
    {
        id: 4,
        imageData:image4,
        merchantName:"ZHH",
        productName:"apron",
        productType:"",
        unitPrice:6
    },
    {
        id: 5,
        imageData:image5,
        merchantName:"ZHH",
        productName:"Custom alarm clock",
        productType:"",
        unitPrice:6
    },
    {
        id: 6,
        imageData:image6,
        merchantName:"ZHH",
        productName:"seasoning box",
        productType:"",
        unitPrice:6
    },

    // ...更多的数据项
];

export default function Itemcard(productList) {
    // const [products,setProducts]=useState([])
    // useEffect(()=>{
    //     setProducts(productList);
    //     console.log(products)
    //     // setProducts(data)
    // })
    useEffect(() => {
        console.log(productList); // 这将显示传入的 productList
    }, []); // 添加 productList 作为依赖


    return (
        <div className="card-container">
            {productList.map((productList) => (
                <Link key={productList.id} className='link' to={`/shop/${productList.id}`}>
                    <Card title={productList.productName} className='custom-card' hoverable cover={<img alt="example" className='card-image' src={`data:image/jpeg;base64,${productList.imageData}`} />}>
                        <p>{productList.description}</p>
                        <Meta title={productList.title} description={`provided by ${productList.merchantName}`} />
                    </Card>
                </Link>
            ))}
        </div>
        // <div className="card-container">
        //     {products.map((product) => (
        //         <Link className='link' to={`/shop/${product.id}`}>
        //         <Card key={products.id} title={product.productName} className='custom-card' hoverable cover={<img alt="example" className='card-image' src={`data:image/jpeg;base64,${product.imageData}`}/>}>
        //             <p>{product.description}</p>
        //             <Meta title={product.title} description={`provided by ${product.merchantName}`} />
        //         </Card>
        //         </Link>
        //     ))}
        // </div>
    );
}

