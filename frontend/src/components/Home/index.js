import React from "react";
import { Link } from "react-router-dom";
import intro from "../images/intro.jpg";
import "./Home.css";

const Home = () => {
  return (
    <div className="header">
      <div className="container">
        <div className="row">
          <div className="col2">
            <h1>
              Welcome To Free Book <br />
              Rental Store!
            </h1>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipisicing elit,
              <br />
              sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.{" "}
            </p>
            <Link to="" class="btn">
              Explore Now &#8594;{" "}
            </Link>
          </div>
          <div class="col-2">
            <img src={intro} alt="" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
