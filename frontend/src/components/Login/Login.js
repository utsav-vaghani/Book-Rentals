import React from "react";
import avatar from "../../images/avatar.png";
import "../../App.css";
import { Link } from "react-router-dom";
const Login = () => {
  return (
    <div className="register">
      <img src={avatar} className="avatar" alt="" />
      <h1>LogIn</h1>
      <form>
        <p>Email</p>
        <input type="email" name="email" placeholder="Enter Email..." />
        <p>Password</p>
        <input
          type="password"
          name="password"
          placeholder="Enter Password..."
        />
        <input type="submit" name="" value="Login" />
        <Link to="/account/register">Don't Have An Account?</Link>
        {/* TODO */}
        <Link to="/">Forgot Your Password</Link>
      </form>
    </div>
  );
};

export default Login;
