import React, { useState } from 'react';
import { Link, Redirect } from 'react-router-dom';
import axios from 'axios';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import styles from './Register.module.css';
import booksImg from '../../images/books_02.svg';
import { AiOutlineLoading } from 'react-icons/ai';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL;

function Register() {
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        address: '',
        password: '',
    });

    const [buttonDisabled, setButtonDisabled] = useState(false);
    const [registerSuccess, setRegisterSuccess] = useState(false);

    const submit_handler = async (e) => {
        e.preventDefault();
        console.log(formData);

        // Checking if any filed is empty
        if (
            !Object.values(formData).every(
                (field) => field !== '' && field !== null,
            )
        )
            toast(
                `please enter ${Object.keys(formData).find(
                    (key) => formData[key] === '',
                )}`,
            );
        else {
            try {
                setButtonDisabled(true);
                const res = await axios.post(
                    `${BACKEND_URL}/api/auth/register`,
                    formData,
                );
                if (res.status === 201) setRegisterSuccess(true);
                toast(res.data.message);
                setButtonDisabled(false);
            } catch (error) {
                toast(error.response.data.message);
                setButtonDisabled(false);
            }
        }
    };
    return (
        <div className={styles.register}>
            <div className={styles.leftCol}>
                <h1 className={styles.heading}>Create Account</h1>
                <form onSubmit={(e) => submit_handler(e)}>
                    <TextField
                        id="name"
                        label="Name"
                        variant="outlined"
                        fullWidth
                        style={{ backgroundColor: 'white' }}
                        value={formData.name}
                        onChange={(e) =>
                            setFormData({
                                ...formData,
                                [e.target.id]: e.target.value,
                            })
                        }
                    />
                    <TextField
                        id="email"
                        label="Email"
                        variant="outlined"
                        fullWidth
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={formData.email}
                        onChange={(e) =>
                            setFormData({
                                ...formData,
                                [e.target.id]: e.target.value,
                            })
                        }
                    />
                    <TextField
                        id="address"
                        label="Address"
                        variant="outlined"
                        fullWidth
                        multiline
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={formData.address}
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
                        style={{ backgroundColor: 'white' }}
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
                            style={{
                                backgroundColor: '#F65944',
                                whiteSpace: 'nowrap',
                            }}
                            color="primary"
                            type="submit"
                            disabled={buttonDisabled}
                        >
                            {buttonDisabled && (
                                <AiOutlineLoading
                                    className={styles.rotate}
                                    style={{ marginRight: '0.5rem' }}
                                />
                            )}
                            Sign up
                        </Button>
                        <div className={styles.loginText}>
                            Already have an account?{' '}
                            <Link to="/account/signin">Login</Link>
                        </div>
                    </div>
                </form>
            </div>
            <div className={styles.rightCol}>
                <img src={booksImg} className={styles.coverImg} alt="Books" />
            </div>
            <ToastContainer />
            {registerSuccess && (
                <Redirect
                    push
                    to="/account/signin?registrationStatus=success"
                />
            )}
        </div>
    );
}

export default Register;
