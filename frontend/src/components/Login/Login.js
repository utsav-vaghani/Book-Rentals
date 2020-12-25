import React from "react";
import { useState, useEffect } from "react";
import { Link, Redirect } from "react-router-dom";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import { useQuery } from "../../utils/useQuery";
import { setUser } from "../../utils/auth";
import styles from "./Login.module.css";
import booksImg from "../../images/books_02.svg";

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL;

function Login() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });
  const [buttonDisabled, setButtonDisabled] = useState(false);
  const [loginSuccess, setLoginSuccess] = useState(false);

  const query = useQuery();
  const registrationStatus = query.get("registrationStatus");

  useEffect(() => {
    if (query.get("registrationStatus") === "success")
      toast("Registration Successful! Please Login");
  }, [registrationStatus]);

  const submit_handler = async (e) => {
    e.preventDefault();
    console.log(BACKEND_URL);
    // Checking if any filed is empty
    if (
      !Object.values(formData).every((field) => field !== "" && field !== null)
    )
      toast(
        `please enter ${Object.keys(formData).find(
          (key) => formData[key] === ""
        )}`
      );
    else {
      try {
        setButtonDisabled(true);
        const res = await axios.post(`${BACKEND_URL}/api/auth/login`, formData);

        if (res.status === 200) {
          setUser(res.data.token);
          setLoginSuccess(true);
        }

        toast(res.data.message);
        setButtonDisabled(false);
      } catch (error) {
        toast(error.response.data.message);
        setButtonDisabled(false);
      }
    }
  };
  return (
    <div className={styles.login}>
      <div className={styles.leftCol}>
        <img src={booksImg} className={styles.coverImg} alt="Books" />
      </div>
      <div className={styles.rightCol}>
        <h1 className={styles.heading}>Login</h1>

        <form onSubmit={(e) => submit_handler(e)}>
          <TextField
            id="email"
            label="Email"
            variant="outlined"
            fullWidth
            margin="normal"
            style={{ backgroundColor: "white" }}
            value={formData.email}
            onChange={(e) =>
              setFormData({
                ...formData,
                [e.target.id]: e.target.value,
              })
            }
          />
          <TextField
            id="password"
            label="Password"
            variant="outlined"
            type="password"
            autoComplete="current-password"
            fullWidth
            margin="normal"
            style={{ backgroundColor: "white" }}
            value={formData.password}
            onChange={(e) =>
              setFormData({
                ...formData,
                [e.target.id]: e.target.value,
              })
            }
          />
          <div className={styles.buttonArea}>
            <Button
              variant="contained"
              className={styles.submitButton}
              style={{ backgroundColor: "#F65944" }}
              color="primary"
              type="submit"
              disabled={buttonDisabled}
            >
              {" "}
              Sign In
            </Button>
            <div className={styles.loginText}>
              or, <Link to="/account/register">Create an Account</Link>
            </div>
          </div>
        </form>
      </div>
      <ToastContainer />
      {loginSuccess && <Redirect push to="/" />}
    </div>
  );
}

export default Login;
