import React from "react";
import { Link } from "react-router-dom";
import avatar from "../images/avatar.png";

const Login = () => {
  return (
    <div className="bg">
      <div className="register">
        <img src={avatar} className="avatar" alt="" />
        <h1>SignIn</h1>
        <form>
          <p>Username</p>
          <input type="text" name="" placeholder="Enter Username" />
          <p>Password</p>
          <input type="password" name="" placeholder="Enter Password" />
          <input type="submit" name="" value="SignIn" />
          <Link to="/account/register">Don't Have An Account?</Link>
          <br />
          <Link to="/">Forgot Your Password</Link>
        </form>
      </div>
    </div>
  );
};

export default Login;
