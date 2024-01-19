import React from 'react';
import { useEffect } from 'react';
import { UserOutlined, SearchOutlined } from '@ant-design/icons';
import { Button, Dropdown, Space, message } from 'antd';
import './navbar.css';
import axios from 'axios';
import { Link } from 'react-router-dom';

const isLoggedIn = () => {
  const token = localStorage.getItem('authToken');
  return !!token; // 如果 token 存在则返回 true，否则返回 false
};


function handleNavigation(targetPath){
  if(!isLoggedIn()){
    window.location.href='/login'
  }else{
    window.location.href=`${targetPath}`
  }
}

function handleUserClick(key) {
  if(key===1){
    window.location.href='/login'
  }
  // message.info(`Click on itme ${key}`)
  switch (key) {
    case '1':
      window.location.href = '/';
      break;
    case '2': 
      window.location.href = '/';
      break;
    case '3': 
      window.location.href = '/login';
      break;

  }
}

const onClick = ({ key }) => {
  handleUserClick(key)
};


function Navbar() {
  return (
    <header className="navbar-container">
      <div className="navbar">
        <div className="navbar-logo"><Link to="/" className="navbar-logo">
          Web3Trade
        </Link></div>
        <nav className="navbar-links">
          <Link to="/" className="nav-link">Home</Link>
          <Link to="/shop" className="nav-link">Shop</Link>
          <Link to="/Orders" className="nav-link" onClick={() => handleNavigation('/orders')}>Orders</Link>
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
var items = [
  {
    label: 'User Profile',
    key: 1,
  },
  {
    label: 'My Orders',
    key: 2,
  },
  {
    label: 'Log in',
    key: 3,
  },
  {
    label: 'Log out',
    key: 4,
    danger: true,
  }
];


export default Navbar;
