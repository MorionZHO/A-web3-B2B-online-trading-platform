import React from 'react';
import { Descriptions,Card } from 'antd';
import Navbar from '../../components/navbar/navbar';
const items = [
    {
        key: '1',
        label: 'UserName',
        children: 'Zhou Maomao',
    },
    {
        key: '2',
        label: 'Telephone',
        children: '1810000000',
    },
    {
        key: '3',
        label: 'Live',
        children: 'Hangzhou, Zhejiang',
    },
    {
        key: '4',
        label: 'Remark',
        children: 'empty',
    },
    {
        key: '5',
        label: 'Address',
        children: 'No. 18, Wantang Road, Xihu District, Hangzhou, Zhejiang, China',
    },
];

function UserInfo() {
    return (
        <>
        <Navbar></Navbar>
        <Card><Descriptions title="User Info" items={items} /></Card>
        </>
    );

}

export default UserInfo;