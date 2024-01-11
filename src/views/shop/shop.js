import React ,{useState} from 'react';
import Navbar from '../../components/navbar/navbar';
import { Menu} from 'antd';
import Itemcard from './itemCard';
import './shop.css'

export default function Shop() {
    return (
        <>
            <Navbar></Navbar>
            <div className='shop-container'>
                <div className='sidebar'>
                <Sidebar></Sidebar>
                </div>
                <div className='item-cards'>
                <Itemcard></Itemcard>
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
    getItem('Navigation One', 'sub1', null, [
        getItem('Option 1', '1'),
        getItem('Option 2', '2'),
        getItem('Option 3', '3'),
        getItem('Option 4', '4'),
    ]),
    getItem('Navigation Two', 'sub2', null, [
        getItem('Option 5', '5'),
        getItem('Option 6', '6'),
        getItem('Submenu', 'sub3', null, [getItem('Option 7', '7'), getItem('Option 8', '8')]),
    ]),
    
    getItem('Navigation Three', 'sub4', null, [
        getItem('Option 9', '9'),
        getItem('Option 10', '10'),
        getItem('Option 11', '11'),
        getItem('Option 12', '12'),
    ]),
    
];

const rootSubmenuKeys = ['sub1', 'sub2', 'sub4'];
function Sidebar() {
    const [openKeys, setOpenKeys] = useState(['sub1']);
    const onOpenChange = (keys) => {
        const latestOpenKey = keys.find((key) => openKeys.indexOf(key) === -1);
        if (latestOpenKey && rootSubmenuKeys.indexOf(latestOpenKey) === -1) {
          setOpenKeys(keys);
        } else {
          setOpenKeys(latestOpenKey ? [latestOpenKey] : []);
        }
      };
    const onClick = (e) => {
        console.log('click ', e);
    };
    return (
    <>
    <section style={{marginLeft:'5px'}}>
    <h2 style={{marginLeft:'15px',marginBottom:'10px',fontWeight:'normal'}}>Catagories</h2>
        <Menu
            onClick={onClick}
            openKeys={openKeys}
            onOpenChange={onOpenChange}
            style={{background:'#f0f2f5',border:'0px'}}
            mode="inline"
            items={items}
        />
        </section>
        </>
    )
}