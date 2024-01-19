import '../App.css';
import About from './about/about';
import Home from './home/home';
import Shop from './shop/shop'
import Login from './login/login'
import Register from './register/register';
import Orders from './orders/orders'
import React from 'react';
import { Button, ConfigProvider } from 'antd';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link
} from 'react-router-dom';

function App() {
  return (
    <ConfigProvider
      theme={{
        token: {
          colorPrimary: '#3293C6',
          borderRadius: 5,
        },
      }}>
      <Router>
        {/* 路由切换 */}
        <Routes>
          <Route path="/about" element={<About />} />
          <Route path="/" element={<Home />} />
          <Route path="/shop" element={<Shop />} />
          <Route path="/orders" element={<Orders />} />
          <Route path='/login' element={<Login />}></Route>
          <Route path='/register' element={<Register></Register>}></Route>
        </Routes>

      </Router>

    </ConfigProvider>

  );
}

export default App;
