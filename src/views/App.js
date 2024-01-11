import '../App.css';
import About from './about/about';
import Home from './home/home';
import Shop from './shop/shop'
import Contact from './contact/contact'
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
            <Route path="/shop" element={<Shop/>}/>
            <Route path="/contact" element ={<Contact/>}/>
          </Routes>
        
      </Router>

    </ConfigProvider>

  );
}

export default App;
