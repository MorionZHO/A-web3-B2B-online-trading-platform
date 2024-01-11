import React from 'react';
import { UserOutlined, SearchOutlined } from '@ant-design/icons';
import { Button, Dropdown, Space } from 'antd';
import './navbar.css';

function Navbar() {
    return (
        <header className="navbar-container">
            <div className="navbar">
                <div className="navbar-logo">Web3Trade</div>
                <nav className="navbar-links">
                    <a href="/" className="nav-link">Home</a>
                    <a href="/about" className="nav-link">About</a>
                    <a href="/shop" className="nav-link">Shop</a>
                    <a href="/contact" className="nav-link">Contact</a>
                </nav>
                <div className="navbar-icons">
                    <button aria-label="Search" title='Search'><SearchOutlined style={{ color: '#3293C6', fontSize: '24px' }} /></button>
                    <Space direction='vertical'>
                        <Space wrap>
                            <Dropdown menu={{ items, }}
                                placement='bottom'
                            >
                                <button aria-label="User Profile" title='User Profile'><UserOutlined style={{ color: '#3293C6', fontSize: '24px' }} /></button>
                            </Dropdown>
                        </Space>
                    </Space>


                </div>

            </div>
        </header>
    );

  
}
var items =  [
    {
      label: 'User Profile',
      key: '1',
    },
    {
      label: 'My Orders',
      key: '2',
    },
    {
      label: 'Login',
      key: '3',
    },
  ];

export default Navbar;
