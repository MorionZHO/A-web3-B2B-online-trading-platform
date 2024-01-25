import React, { useState, useEffect } from 'react';
import Navbar from '../../components/navbar/navbar';
import { Menu } from 'antd';
import Itemcard from './itemCard';
import './shop.css'
import { getAllProducts, searchProduct } from '../../api/product';

export default function Shop() {
    const [keyWord,setKeyWord] = useState("All")
    
    const setKeyWords = (k) => {
        setKeyWord(k);
        console.log(`keyword:${keyWord}`)
    }

    return (
        <>
            <Navbar onSearch={setKeyWord}></Navbar>
            <div className='shop-container'>
                <div className='sidebar'>
                    <Sidebar onData={setKeyWords}></Sidebar>
                </div>
                <div className='item-cards'>
                    <Itemcard keyWord={keyWord}></Itemcard>
                </div>
            </div>
        </>
    )
}

function getItem(label, key, icon, children, type) {
    return {
        key,
        icon,
        children,
        label,
        type,
    };
}

const items = [
    {
        type: 'divider',
    },
    getItem('All', 'All', null, null),
    getItem('Shoes', 'Shoes', null, null),
    getItem('Clothes', 'Clothes', null, null),
    getItem('Electronics', 'Electronics', null, null),
    getItem('Home goods', 'Home goods', null, null)

];

const rootSubmenuKeys = ['All', 'Shoes', 'Clothes', 'Electronics', 'Home goods'];

function Sidebar({ onData }) {
    const [openKeys, setOpenKeys] = useState(['All']);
    const onOpenChange = (keys) => {
        const latestOpenKey = keys.find((key) => openKeys.indexOf(key) === -1);
        if (latestOpenKey && rootSubmenuKeys.indexOf(latestOpenKey) === -1) {
            setOpenKeys(keys);
        } else {
            setOpenKeys(latestOpenKey ? [latestOpenKey] : []);
        }
    };
    const onClick = (e) => {
        console.log(e.key)
        onData(e.key)
    };
    return (
        <>
            <section style={{ marginLeft: '5px' }}>
                <h2 style={{ marginLeft: '15px', marginBottom: '10px', fontWeight: 'normal' }}>Catagories</h2>
                <Menu
                    onClick={onClick}
                    openKeys={openKeys}
                    onOpenChange={onOpenChange}
                    style={{ background: '#f0f2f5', border: '0px' }}
                    mode="inline"
                    items={items}
                />
            </section>
        </>
    )
}