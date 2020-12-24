import jwt from 'jsonwebtoken';

export const setUser = (token) => {
    localStorage.setItem('token', token);
    const user = jwt.decode(token);
    localStorage.setItem('user', JSON.stringify(user));
};

export const getUser = () => {
    const token = localStorage.getItem('token');
    const user = JSON.parse(localStorage.getItem('user'));
    user['token'] = token;
    return user;
};
