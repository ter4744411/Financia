import React,{useEffect,useState} from 'react'
import './AdminDashboard.scss'
import { IoIosArrowDown } from "react-icons/io";
import { GoHomeFill } from "react-icons/go";
import { TbCertificate } from "react-icons/tb";
import { IoReceipt } from "react-icons/io5";
import { FaTrophy } from "react-icons/fa";
import { IoMdPersonAdd } from "react-icons/io";
import { BiSolidUserAccount } from "react-icons/bi";
import { MdOutlineSupervisorAccount } from "react-icons/md";
import { IoMdSettings } from "react-icons/io";
import userpic from "../UserDashboard/userpic.png"
import { IoShieldCheckmarkSharp } from "react-icons/io5";
import { FaXTwitter } from "react-icons/fa6";
import userinfo from "../UserDashboard/userinfo-right.png"
import { useSelector } from 'react-redux'
import { RootState } from '../../redux/store';
import Papa from 'papaparse';
import { Link } from 'react-router-dom';
import { useLocation,useNavigate } from 'react-router-dom';
import { logout } from '../../redux/user/userSlice';
import { useDispatch } from 'react-redux'

interface User {
    idnumber: string;
    email: string;
    username: string;
    role: string;
  }

const AdminDashboard: React.FC = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [user, setUser] = useState<{ idnumber: string; username: string; role: string } | null>(null);

  useEffect(() => {
      fetchUserById();
  }, []);

  const fetchUserById = async () => {
    try {
      const token = localStorage.getItem("token");
      const response = await fetch(`http://localhost:8080/admindashboard`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`, // Replace token with your actual JWT token
          'Content-Type': 'application/json'
        }
      });
      console.log(response)
      const userData = await response.json();
      setUser(userData);
      if (userData.message){
        alert(userData.message)
        dispatch(logout()); 
        navigate("/login")
      }
      console.log("userdata: ",userData)
    } catch (error) {
      console.error('Error fetching user data:', error);
    }
  };
    const [userList, setUserList] = useState<any[]>([]);
    const userSelector = useSelector((state: RootState) => state.user);
    console.log("from admindashboard.jsx the crrent state username :",userSelector.username);

    useEffect(() => {
        fetch('../../../users.csv')
          .then(response => response.text())
          .then(data => {
            const parsed = Papa.parse<User>(data, {
                header: true,
                dynamicTyping: true,
                skipEmptyLines: true,
                complete: (results) => {
                    setUserList(results.data);
                    if (results.errors.length) {
                      console.error("Errors while parsing:", results.errors);
                    }
                  }
              });
            console.log("csv user: ",parsed)
          })
          .catch(error => {
            console.error('Error fetching the CSV file:', error);
          });
      }, []);


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
                    <span><span style={{color:"#fcd535",fontWeight:"500"}}>Username :</span> {userSelector.username}</span>
                    <span><span style={{color:"#fcd535",fontWeight:"500"}}>Email :</span> {userSelector.email}</span>
                    <span><span style={{color:"#fcd535",fontWeight:"500"}}>Role :</span> {userSelector.role}</span>
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
        <div className="user-list">
        {userList.map((user, index) => (
            <Link key={index} to={`/userdashboard?idnumber=${user.idnumber}`}>
              <div className="user-box">
                <span><span className="idnumber">Idnumber:</span> {user.idnumber}</span>
                <span><span className="username">Username:</span> {user.username}</span>
                <span><span className="role">Role:</span> {user.role}</span>
              </div>
            </Link>
          ))}
        </div>
      </div>
    </div>
  )
}

export default AdminDashboard
