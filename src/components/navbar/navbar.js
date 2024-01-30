import React from 'react';
import { useEffect, useState } from 'react';
import { UserOutlined, SearchOutlined } from '@ant-design/icons';
import { Button, Dropdown, Space, message, Input } from 'antd';
import './navbar.css';
import axios from 'axios';
import { Link } from 'react-router-dom';
import storage from '../../utils/storage';
import imagePic from './logo.png'

const localCache = storage.getItem('localCache') ? storage.getItem('localCache') : undefined;

const { Search } = Input;

const isLoggedIn = () => {
  if (localCache !== undefined) {
    const token = localCache['authToken']
    return !!token; // 如果 token 存在则返回 true，否则返回 false
  }
  else return false
};

const isSeller = () => {
  if (localCache !== undefined) {
    const role = localCache['userInfo']['role']
    if (role === 'seller') return true;
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
  if (key === '3') {
    storage.clearAll()
    window.location.href = '/login'
  }
  else if (key === '2') {
    window.location.href = '/manageProduct'
  }
  else if (key === '1') {
    window.location.href = '/userProfile'
  }
};


function Navbar({ onSearch }) {
  const [items, setItems] = useState([])
  const [showSearch, setShowSearch] = useState(false);
  function toggleSearch() {
    setShowSearch(!showSearch);
  }


  useEffect(() => {
    getDropdownItems();
  }, []); // 空依赖数组表示这个 effect 只在组件挂载时运行一次
  const getDropdownItems = () => {
    if (isLoggedIn()) {
      if (isSeller()) {
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
            label: 'Logout',
            key: 3,
            danger: true,
          }
        ])
      } else {
        setItems([
          {
            label: 'User Profile',
            key: 1,
          },
          {
            label: 'Logout',
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
          {
            isSeller()
              ? <Link className="nav-link" onClick={() => handleNavigation('/manageProduct')}>Products</Link>
              : <Link className="nav-link" onClick={() => handleNavigation('/shop')}>Shop</Link>
          }

          <Link className="nav-link" onClick={() => handleNavigation('/orders')}>Orders</Link>
          <Link to="/about" className="nav-link">About Us</Link>

        </nav>
        <div className="navbar-icons">
          {
            showSearch
              ? <MySearch switchIcon={toggleSearch} searchHandle={onSearch}></MySearch>
              : <button aria-label="Search" title='Search' onClick={toggleSearch}>
                <SearchOutlined style={{ color: '#3293C6', fontSize: '24px', width: '50px' }} />
              </button>
          }

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

function MySearch({ switchIcon, searchHandle }) {
  const handleSearch = (value, _e, info) => {
    switchIcon()
    searchHandle(['search', value])
    console.log(value);
  }
  return (
    <Search
      placeholder="input search keywords"
      allowClear
      onSearch={handleSearch}
      style={{
        width: 250,
      }}
    />
  )

}

export default Navbar;
