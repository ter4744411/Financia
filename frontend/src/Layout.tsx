import React from 'react';
import Navbar from './components/Navbar/Navbar';
import { Outlet } from 'react-router-dom';

const Layout: React.FC = () => {
  return (
    <>
     <Navbar/>
     <Outlet/> 
    </>
  )
}

export default Layout
