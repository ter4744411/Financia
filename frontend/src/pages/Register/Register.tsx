import React,{useState,ChangeEvent} from 'react'
import './Register.scss'
import image1 from "./digital_asset.png"
import { MdEmail } from "react-icons/md";
import { FaLock } from "react-icons/fa";
import { IoPerson } from "react-icons/io5";
import { IoIosArrowDroprightCircle } from "react-icons/io";
import { Link,useNavigate } from 'react-router-dom';
import axios from 'axios';

type User = {
  username: string;
  email : string;
  password : string;
}

const Register = () => {
    const navigate = useNavigate();
    const [username,setUsername] = useState<string>('');
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [errorstate,setErrorstate] = useState<string>('');

    const handleSubmit = async (e:React.FormEvent<HTMLFormElement>)=>{
      e.preventDefault();
      const user: User = {
        username: username,
        email: email,
        password: password,
      };
      console.log("form info : ",user)
      try {
        const res = await axios.post('http://localhost:8080/registerclient', user);
        console.log("res from frontend: ",res)
        setErrorstate("")
        navigate("/login")
      } catch (error: any) {
        console.log(error)
        console.error(error.response.data.message);
        setErrorstate(error.response.data.message);
      }
    }
    const handleUsernameChange = (event: ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value);
      };
    
      const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
        setEmail(event.target.value);
      };
    
      const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
      };
    
      // const handleImageChange = (e: ChangeEvent<HTMLInputElement>) => {
      //   const file = e.target.files && e.target.files[0];
      //   setImage(file); 
      // };
    
  return (
    <div className="register">
      <div className="container">
        <div className="left">
            <img src={image1} alt=""/>
        </div>
        <div className={errorstate === "" ? "right" : "right-error"}>
            <span className="register-header">สมัครสมาชิก</span>
            <form onSubmit={handleSubmit}>
                <div className="input-box">
                    <MdEmail className="form-icon"/>
                    <input type="email" name="email" placeholder="ที่อยู่อีเมล" onChange={handleEmailChange}/>
                </div>
                <div className="input-box">
                    <FaLock className="form-icon"/>
                    <input type="password" name="password" placeholder="รหัสผ่าน" onChange={handlePasswordChange}/>
                </div>
                <div className="input-box">
                    <IoPerson className="form-icon"/>
                    <input type="text" name="username" placeholder="ชื่อผู้ใช้" onChange={handleUsernameChange}/>
                </div>
                {/* <div className="input-img">
                    <label htmlFor="image">รูปภาพ</label>
                    <input type="file" id="image" onChange={handleImageChange} />
                </div> */}
                <button style={{fontFamily:"Kanit"}} type="submit">สมัครสมาชิก <IoIosArrowDroprightCircle style={{fontSize:"20px"}}/></button>
                {errorstate && <div className="error">{errorstate}</div>}
            </form>
            <Link to="/login">
              <span className="already-member">มีบัญชีอยู่แล้ว ล็อกอินที่นี่ <IoIosArrowDroprightCircle style={{color:"#f0b90b",fontSize:"20px"}}/></span>
            </Link>
            <span className="login-footer">ได้ใบอนุญาตประกอบธุรกิจจากกระทรวงการคลัง โดยอยู่ภายใต้การกำกับดูแลของสำนักงาน ก.ล.ต.</span>
        </div>
      </div>
    </div>
  )
}

export default Register
