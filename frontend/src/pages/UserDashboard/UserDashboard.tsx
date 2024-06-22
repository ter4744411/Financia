import React,{ useEffect,useState } from 'react'
import './UserDashboard.scss'
import { IoIosArrowDown } from "react-icons/io";
import { GoHomeFill } from "react-icons/go";
import { TbCertificate } from "react-icons/tb";
import { IoReceipt } from "react-icons/io5";
import { FaTrophy } from "react-icons/fa";
import { IoMdPersonAdd } from "react-icons/io";
import { BiSolidUserAccount } from "react-icons/bi";
import { MdOutlineSupervisorAccount } from "react-icons/md";
import { IoMdSettings } from "react-icons/io";
import userpic from "./userpic.png"
import { IoShieldCheckmarkSharp } from "react-icons/io5";
import { FaXTwitter } from "react-icons/fa6";
import userinfo from "./userinfo-right.png"
import userbalance from "./userbalance.png"
import userholding from "./userholding.png"
import { useLocation,useNavigate } from 'react-router-dom';
import { logout } from '../../redux/user/userSlice';
import { useDispatch } from 'react-redux'
import { useSelector } from 'react-redux';
import { RootState } from '../../redux/store';

const UserDashboard: React.FC = () => {
  const dispatch = useDispatch();
let location = useLocation();
const navigate = useNavigate();
const userSelector = useSelector((state: RootState) => state.user);
const [user, setUser] = useState<{ idnumber: string; username: string; role: string } | null>(null);

useEffect(() => {
  const searchParams = new URLSearchParams(location.search);
  const idnumber: string | null = searchParams.get('idnumber');
  if (idnumber) {
    fetchUserById(idnumber);
    }
  }, [location]);

  const fetchUserById = async (idnumber: string) => {
  try {
  const token = localStorage.getItem("token");
  const response = await fetch(`http://localhost:8080/userdashboard/${idnumber}`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`, // Replace token with your actual JWT token
      'Content-Type': 'application/json'
    }
  });
  console.log(response)
  const userData = await response.json();
  setUser(userData);
  if (userData.message && userSelector.role !== "Admin"){
    alert(userData.message)
    dispatch(logout());
    navigate("/login")
  }
  console.log("userdata: ",userData)
  } catch (error) {
  console.error('Error fetching user data:', error);
    }
  };
  
  
  return (
    <div className="userdashboard">
      <div className="left">
        <div><GoHomeFill /> &nbsp;&nbsp;Dashboard</div>
        <div style={{justifyContent:"space-between"}}><span><TbCertificate /> &nbsp;&nbsp;Assets</span> <IoIosArrowDown /></div>
        <div style={{justifyContent:"space-between"}}><span><IoReceipt /> &nbsp;&nbsp;Orders</span> <IoIosArrowDown /></div>
        <div><FaTrophy /> &nbsp;&nbsp;Rewards Hub</div>
        <div><IoMdPersonAdd /> &nbsp;&nbsp;Referral</div>
        <div style={{justifyContent:"space-between"}}><span><BiSolidUserAccount /> &nbsp;&nbsp;Account</span> <IoIosArrowDown /></div>
        <div><MdOutlineSupervisorAccount /> &nbsp;&nbsp;Sub Accounts</div>
        <div><IoMdSettings /> &nbsp;&nbsp;Settings</div>
      </div>
      <div className="right">
        <div className="userinfo">
            <div className="userinfo-left">
                <img src={userpic} alt=""/>
                <div className="userinfo-name">
                    <span><span style={{color:"#fcd535",fontWeight:"500"}}>Username :</span> {user?.username}</span>
                    <span><span style={{color:"#fcd535",fontWeight:"500"}}>UserID :</span> {user?.idnumber}</span>
                    <span><span style={{color:"#fcd535",fontWeight:"500"}}>Role :</span> {user?.role}</span>
                    <div className="userinfo-verify">
                        <IoShieldCheckmarkSharp className="verify-icon"/>
                        <FaXTwitter className="verify-icon"/>
                    </div>
                </div>
            </div>
            <div className="userinfo-right">
                <img src={userinfo} alt=""/>
            </div>
        </div>
        <img src={userbalance} alt=""/>
        <img src={userholding} alt=""/>
      </div>
    </div>
  )
}

export default UserDashboard
