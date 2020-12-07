import React from "react";
import { Link } from "react-router-dom";
import "../App.css";
import avatar from "../images/avatar.png";

const Register = () => {
  return (
    <div className="bg">
      <div className="register">
        <img src={avatar} className="avatar" alt="" />
        <h1>Register Here</h1>
        <form>
          <p>Username</p>
          <input type="text" name="" placeholder="Enter Username" />
          <p>Password</p>
          <input type="password" name="" placeholder="Enter Password" />
          <p>Confirm Password</p>
          <input type="password" name="" placeholder="Enter Confirm Password" />
          <input type="submit" name="" value="SignUp" />
          <Link to="/account/signin">Already Have An Account</Link>
        </form>
      </div>
    </div>
  );
};

export default Register;
