import React from 'react'
import './Navbar.scss'
import { IoIosArrowDown } from "react-icons/io";
import { FaSearch,FaMoon } from "react-icons/fa";
import { MdOutlineQrCode2,MdLanguage } from "react-icons/md";
import { Link } from "react-router-dom";
import { useSelector } from 'react-redux'
import { RootState } from '../../redux/store';
import { useDispatch } from 'react-redux'
import { logout } from '../../redux/user/userSlice';

const Navbar: React.FC = () =>{
  const dispatch = useDispatch();
  const handleLogOut = () => {
    localStorage.removeItem('token');
    dispatch(logout()); 
  };
  const userSelector = useSelector((state: RootState) => state.user);
  return (
    <div className="navbar">
      <div className="nav-left">
        <Link to="/">
          <span className="homelogo">/LOGO</span>
        </Link>
        <span className="menu">Buy Crypto</span>
        <span className="menu">Markets</span>
        <span className="menu">Trade<IoIosArrowDown /></span>
        <span className="menu">Futures<IoIosArrowDown /></span>
        <span className="menu">Earn</span>
        <span className="menu">Square<IoIosArrowDown /></span>
        <span className="menu">More<IoIosArrowDown /></span>
      </div>
      <div className="nav-right">
        <FaSearch className="menu-icon"/>
        {userSelector.username ? (
          <Link to="/login">
            <button className="login-btn" onClick={handleLogOut}>LogOut</button>
          </Link>
        ) : (
          <>
            <Link to="/login">
              <button className="login-btn">Log In</button>
            </Link>
            <Link to="/register">
              <button className="signup-btn">Sign Up</button>
            </Link>
          </>
        )}
        {
          userSelector.username && (
            <div className="nav-username">
              {userSelector.username}
            </div>
          )
        }
        <MdOutlineQrCode2 className="menu-icon"/>
        <MdLanguage className="menu-icon"/>
        <FaMoon className="menu-icon"/>
      </div>
    </div>
  )
}

export default Navbar
