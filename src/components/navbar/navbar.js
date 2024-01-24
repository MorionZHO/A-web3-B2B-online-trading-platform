import React from 'react';
import { useEffect, useState } from 'react';
import { UserOutlined, SearchOutlined } from '@ant-design/icons';
import { Button, Dropdown, Space, message } from 'antd';
import './navbar.css';
import axios from 'axios';
import { Link } from 'react-router-dom';
import storage from '../../utils/storage';

const localCache = storage.getItem('localCache')?storage.getItem('localCache'):undefined;

const isLoggedIn = () => {
  if(localCache!==undefined){
    const token=localCache['authToken']
    return !!token; // 如果 token 存在则返回 true，否则返回 false
  }
  else return false
};

const isSeller = () =>{
  if(localCache!==undefined){
    const role=localCache['userInfo']['role']
    console.log(role)
    if(role==='seller')return true;
  }
  return false
}


function handleNavigation(targetPath) {
  if (!isLoggedIn()) {
    window.location.href = '/login'
  } else {
    window.location.href = `${targetPath}`
  }
}


const onClick = ({ key }) => {
  if(key==='3'){
    storage.clearAll()
    window.location.href='/login'
  }
  else if(key === '2'){
    window.location.href='/manageProduct'
  }
};


function Navbar() {
  const [items, setItems] = useState([])
  
  useEffect(() => {
    getDropdownItems();
  }, []); // 空依赖数组表示这个 effect 只在组件挂载时运行一次
  const getDropdownItems = () => {
    if (isLoggedIn()) {
      if(isSeller()){
        setItems([
          {
            label: 'User Profile',
            key: 1,
          },
          {
            label: 'Manage Products',
            key: 2,
          },
          {
            label:'Logout',
            key: 3,
            danger: true,
          }
        ])
      }else{
        setItems([
          {
            label: 'User Profile',
            key: 1,
          },
          {
            label:'Logout',
            key: 3,
            danger: true,
          }
        ])
      }
      
    } else {
      setItems([
        {
          label: <a href="/login">Login</a>,
          key: 1,
        },
      ])
    }
  };


  return (
    <header className="navbar-container">
      <div className="navbar">
        <div className="navbar-logo"><Link to="/" className="navbar-logo">
          Web3Trade
        </Link></div>
        <nav className="navbar-links">
          <Link to="/" className="nav-link">Home</Link>
          <Link  className="nav-link" onClick={() => handleNavigation('/shop')}>Shop</Link>
          <Link  className="nav-link" onClick={() => handleNavigation('/orders')}>Orders</Link>
          <Link to="/about" className="nav-link">About Us</Link>

        </nav>
        <div className="navbar-icons">
          <button aria-label="Search" title='Search'><SearchOutlined style={{ color: '#3293C6', fontSize: '24px' }} /></button>
          <Space direction='vertical'>
            <Space wrap>
              <Dropdown menu={{ items, onClick }}
                trigger={['hover']}
                placement='bottom'
              >
                <button aria-label="User Profile" title='User Profile' onClick={(e) => e.preventDefault()}><UserOutlined style={{ color: '#3293C6', fontSize: '24px' }} /></button>
              </Dropdown>
            </Space>
          </Space>
        </div>
      </div>
    </header>
  );
}



export default Navbar;
