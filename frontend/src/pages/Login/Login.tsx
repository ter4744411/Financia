import React,{useState,ChangeEvent} from 'react'
import "./Login.scss"
import image1 from "./login_img.png"
import image2 from "./login_img2.png"
import { MdEmail } from "react-icons/md";
import { FaLock } from "react-icons/fa";
import { IoIosArrowDroprightCircle } from "react-icons/io";
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';
import { login } from '../../redux/user/userSlice'
import { useDispatch } from 'react-redux'

type User = {
  email : string;
  password : string;
}

const Login: React.FC = () => {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [errorstate,setErrorstate] = useState<string>('');
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const handleSubmit = async (e:React.FormEvent<HTMLFormElement>)=>{
      e.preventDefault();
      const data: User = {
        email: email,
        password: password,
      };
      console.log("form info : ",data)
      try {
        const res = await axios.post('http://localhost:8080/loginuser', data);
        console.log("res from frontend: ",res)
        setErrorstate("")
        const token = res.data.token;
        localStorage.setItem('token', token);
        const userinfo = await axios.get('http://localhost:8080/userinfo', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log("current user info: ",userinfo)
        dispatch(login({
          //res มี username เป็น propoty ดูจาก usercontroller
          idnumber : userinfo.data.idnumber,
          username : userinfo.data.username,
          email : userinfo.data.email,
          role: userinfo.data.role,
        }))
        console.log(`idnumber : ${userinfo.data.idnumber},
          username : ${userinfo.data.username},
          email : ${userinfo.data.email},
          role: ${userinfo.data.role},`)
        {userinfo.data.role === "Admin"? navigate(`/admindashboard`) : navigate(`/userdashboard?idnumber=${userinfo.data.idnumber}`)}
      } catch (error: any) {
        console.error('Error occur:', error.response.data);
        setErrorstate( error.response.data.message)
      }
    }
    
      const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
        setEmail(event.target.value);
      };
    
      const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
      };
    
  return (
    <div className="login">
      <div className="container">
        <div className="left">
            <img src={image2} alt=""/>
            <img src={image1} alt=""/>
        </div>
        <div className={errorstate === "" ? "right" : "right-error"}>
            <span className="login-header">เข้าสู่ระบบ  <span style={{color:"#f0b90b"}}>/ LOGO</span></span>
            <form onSubmit={handleSubmit}>
                <div className="input-box">
                    <MdEmail className="form-icon"/>
                    <input type="email" name="email" placeholder="ที่อยู่อีเมล" onChange={handleEmailChange}/>
                </div>
                <div className="input-box">
                    <FaLock className="form-icon"/>
                    <input type="password" name="password" placeholder="รหัสผ่าน" onChange={handlePasswordChange}/>
                </div>
                <button style={{fontFamily:"Kanit"}} type="submit">ต่อไป <IoIosArrowDroprightCircle style={{fontSize:"20px"}}/></button>
                {errorstate && <div className="error">{errorstate}</div>}
            </form>
            <Link to="/register">
                <span className="already-member">ยังไม่มีบัญชี สมัครที่นี่ <IoIosArrowDroprightCircle style={{color:"#f0b90b",fontSize:"20px"}}/></span>
            </Link>
            <span className="login-footer">ได้ใบอนุญาตประกอบธุรกิจจากกระทรวงการคลัง โดยอยู่ภายใต้การกำกับดูแลของสำนักงาน ก.ล.ต.</span>
        </div>
      </div>
    </div>
  )
}

export default Login
