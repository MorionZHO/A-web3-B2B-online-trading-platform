import React from 'react';
import './home.css';
import Navbar from '../../components/navbar/navbar'
import './carousel.css'
import { Carousel } from 'antd';
import image1 from '../../assets/img/banner_img_01.jpg';
import image2 from '../../assets/img/banner_img_02.jpg';
import image3 from '../../assets/img/banner_img_03.jpg';


export default function Home() {
    return (
        <>
        <Navbar></Navbar>
        <Mycarousel></Mycarousel>
        </>
    )
}

const imgStyle = {
    height: '700px',
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

            </Carousel>
        </section>
    )
}