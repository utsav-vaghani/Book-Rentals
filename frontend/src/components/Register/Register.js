import React from "react";
import { Link } from "react-router-dom";
import "../../App.css";
import avatar from "../../images/avatar.png";

const Register = () => {
  return (
    <div className="register">
      <img src={avatar} className="avatar" alt="" />
      <h1>Register Here</h1>
      <form>
        <p>Name</p>
        <input type="text" name="name" placeholder="Enter Name..." />
        <p>Email</p>
        <input type="email" name="email" placeholder="Enter Email..." />
        <p>Address</p>
        <textarea name="address" placeholder="Enter Address.." />
        <p>Password</p>
        <input type="password" name="password" placeholder="Enter Password" />
        <p>Confirm Password</p>
        <input
          type="password"
          name="cfpassword"
          placeholder="Enter Confirm Password"
        />
        <input type="submit" name="" value="SignUp" />
        <Link href="/accoutn/signin">Already Have An Account</Link>
      </form>
    </div>
  );
};

export default Register;
