import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import styles from './Register.module.css';
import booksImg from '../../images/books_02.svg';

function Register() {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [address, setAddress] = useState('');
    const [password, setPassword] = useState('');

    const submit_handler = (e) => {
        e.preventDefault();
        console.log(name + email + address + password);

        //After logic of submit
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
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                    <TextField
                        id="email"
                        label="Email"
                        variant="outlined"
                        fullWidth
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <TextField
                        id="address"
                        label="Address"
                        variant="outlined"
                        fullWidth
                        multiline
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={address}
                        onChange={(e) => setAddress(e.target.value)}
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
                            style={{
                                backgroundColor: '#F65944',
                                whiteSpace: 'nowrap',
                            }}
                            color="primary"
                            type="submit"
                        >
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
        </div>
    );
}

export default Register;
