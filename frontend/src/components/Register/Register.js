import React , {useState} from "react";
import { Link } from "react-router-dom";
import "../../App.css";
import avatar from "../../images/avatar.png";

function Register() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("")
  const [address, setAddress] = useState("")
  const [password, setPassword] = useState("")
  const [re_password, setRe_password] = useState("")

 const submit_handler = e => {
    e.preventDefault();
    console.log(name + email + address + password);

    //After logic of submit
  }
  return (
    <div className="register">
      <img src={avatar} className="avatar" alt="" />
      <h1>Register Here</h1>
      <form>
        <p>Name</p>
        <input type="text" name="name" placeholder="Enter Name..." value={name} onChange={e => setName(e.target.value)}/>
        <p>Email</p>
        <input type="email" name="email" placeholder="Enter Email..." value={email} onChange={e => setEmail(e.target.value)}/>
        <p>Address</p>
        <textarea name="address" placeholder="Enter Address.." value={address} onChange={e => setAddress(e.target.value)}/>
        <p>Password</p>
        <input type="password" name="password" placeholder="Enter Password" value={password} onChange={e => setPassword(e.target.value)}/>
        <p>Confirm Password</p>
        <input
          type="password"
          name="cfpassword"
          placeholder="Enter Confirm Password"
          value={re_password} onChange={e => setRe_password(e.target.value)}
        />
        <input type="submit" name="" value="SignUp" onClick={submit_handler} />
        <Link href="/accoutn/signin">Already Have An Account</Link>
      </form>
    </div>
  );
};

export default Register;
