import React, { useState } from 'react'
import "./Home.scss"
import { MdOutlineQrCode2 } from "react-icons/md";
import { FcGoogle } from "react-icons/fc";
import { FaApple } from "react-icons/fa";
import popular from "./popular.png"
import { useNavigate,Link } from 'react-router-dom';
import { useSelector } from 'react-redux'
import { RootState } from '../../redux/store';
import { logout } from '../../redux/user/userSlice';

const Home = () => {
    const userSelector = useSelector((state: RootState) => state.user);
  return (
    <div className="home">
      <div className="container">
        <div className="home-left">
            <div className="home-left-upper">
                <span className="user-count">201,382,051</span>
                <span className="user-trust">USERS</span>
                <span className="user-trust">TRUST US</span>
                <div className="left-signup">
                {userSelector.username ? (
                userSelector.role === "Admin" ? (
                  <Link to="/admindashboard">
                    <button className="home-btn">Go Admin Dashboard</button>
                  </Link>
                ) : (
                  <Link to="/userdashboard">
                    <button className="home-btn">Go User Dashboard</button>
                  </Link>
                )
              ) : (
                <>
                  <input placeholder='Email/Phone number' type="text" className="signup-input" />
                  <Link to="/register">
                    <button className="signup-btn">Sign Up</button>
                  </Link>
                </>
              )}
                </div>
            </div>
            <div className="home-left-lower">
                <div className="connection">
                    <span>Or Connect With</span>
                    <div className="connection-icon">
                        <FcGoogle className="icon"/>
                        <FaApple  className="icon"/>
                    </div>
                </div>
                <div className="download">
                    <span>App Download</span>
                    <MdOutlineQrCode2 className="icon"/>
                </div>
            </div>
        </div>
        <div className="home-right">
            <div className="home-right-upper">
                <div className="popular-newlisting">
                    <span className="popular">Popular</span>
                    <span className="newlisting">New Listing</span>
                </div>
                <img src={popular} alt=""/>
                <span className="view-all">View All 350+ Coins &gt;</span>
            </div>
            <div className="home-right-lower">
                <span>US Spot Bitcoin ETFs Show Strong Start in June
Historical Analysis Reveals Bitcoin's Average Q3 Returns
Bitcoin's Circulation Speed Mirrors That of 13 Years Ago, Says CryptoQuant CEO
Sandstorm Pledges $4 Million Financing to Crypto Trading Platform Upbots</span>
                <span className="view-all">View All News &gt;</span>
            </div>
        </div>
      </div>
    </div>
  )
}

export default Home
