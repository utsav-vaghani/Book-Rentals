import React from "react";
import { Link } from "react-router-dom";
import { FaUserAlt } from "react-icons/fa";
import logo from "../images/logo.png";
import cart from "../images/cart.png";
import "../App.css";

const Header = () => {
  return (
    <div className="navbar">
      <div className="logo">
        <Link to="/">
          <img alt="Book Rental" src={logo} />
        </Link>
      </div>
      <nav>
        <ul id="MenuItems">
          <li>
            <Link to="/about">About</Link>
          </li>
          <li>
            <Link to="/contact">Contact Us</Link>
          </li>
          <li>
            <Link to="/account/signin">SignIn</Link>
          </li>
          <li>
            <Link to="/account/register">SignUp</Link>
          </li>
          <li>
            <Link to="/">
              <FaUserAlt />
            </Link>
          </li>
        </ul>
      </nav>
      <li className="cart">
        <Link to="/cart">
          <img alt="" src={cart} className="cart" width="30px" height="30px" />
          <span>0</span>
        </Link>
      </li>
      <img alt="" src="" className="menu-icon" />
    </div>
  );
};

export default Header;
