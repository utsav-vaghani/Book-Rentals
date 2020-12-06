import React from "react";
import avatar from "../../images/avatar.png";
import "../../App.css";
import { Link } from "react-router-dom";
const Login = () => {
  return (
    <div className="register">
      <img src={avatar} className="avatar" alt="" />
      <h1>Register Here</h1>
      <form>
        <p>Username</p>
        <input type="text" name="" placeholder="Enter Username" />
        <p>Password</p>
        <input type="password" name="" placeholder="Enter Password" />
        <input type="submit" name="" value="SignUp" />
        <Link to="/">Don't Have An Account?</Link>
        <Link to="/">Forgot Your Password</Link>
      </form>
    </div>
  );
};

export default Login;
