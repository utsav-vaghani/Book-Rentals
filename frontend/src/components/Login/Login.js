import React from 'react';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import styles from './Login.module.css';
import booksImg from '../../images/books_02.svg';

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const submit_handler = (e) => {
        e.preventDefault();
        console.log(username + password);
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
                        style={{ backgroundColor: 'white' }}
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <TextField
                        id="pass"
                        label="Password"
                        variant="outlined"
                        type="password"
                        autoComplete="current-password"
                        fullWidth
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <div className={styles.buttonArea}>
                        <Button
                            variant="contained"
                            className={styles.submitButton}
                            style={{ backgroundColor: '#F65944' }}
                            color="primary"
                            type="submit"
                        >
                            {' '}
                            Sign In
                        </Button>
                        <div className={styles.loginText}>
                            or,{' '}
                            <Link to="/account/register">
                                Create an Account
                            </Link>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    );
}

export default Login;
