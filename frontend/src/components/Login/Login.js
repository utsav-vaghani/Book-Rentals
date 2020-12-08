import React from "react";
import { useState } from 'react'
import avatar from "../../images/avatar.png";
import "../../App.css";
import { Link } from "react-router-dom";

function Login() {

  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const submit_handler = (e) => {
    e.preventDefault();
    console.log(username + password)
  }
  return (
    <div className="register">
      <img src={avatar} className="avatar" alt="" />
      <h1>LogIn</h1>
      <form>
        <p>Email</p>
        <input type="email" name="email" placeholder="Enter Email..." value={username} onChange={(e) => setUsername(e.target.value)} />
        <p>Password</p>
        <input
          type="password"
          name="password"
          placeholder="Enter Password..."
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <input type="submit" onClick={submit_handler} />
        <Link to="/account/register">Don't Have An Account?</Link>
        {/* TODO */}
        <Link to="/">Forgot Your Password</Link>
      </form>
    </div>
  );
};

export default Login;

