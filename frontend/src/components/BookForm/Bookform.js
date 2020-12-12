import React, { useState } from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import styles from './Bookform.module.css';
import booksImg from '../../images/books_02.svg';
import CurrencyTextField from '@unicef/material-ui-currency-textfield';

import IconButton from '@material-ui/core/IconButton';
import PhotoCamera from '@material-ui/icons/PhotoCamera';
function Bookform() {
    const [Title, setTitle] = useState('');
    const [Author, setAuthor] = useState('');
    const [Description, setDescription] = useState('');
    const [OwnerID, setOwnerID] = useState('');
    const [Price, setPrice] = useState('');
    const [Stock, setStock] = useState('');
    const [NumberofPages, setNumberofPages] = useState('');
    const [Picture,setPicture]=useState('')
        

    const submit_handler = (e) => {
        e.preventDefault();
        console.log(Title+Author+Description+OwnerID+Price+Stock+NumberofPages+Picture);

        //After logic of submit
    };


    return (
        <div className={styles.bookform}>
            <div className={styles.leftCol}>
                <h1 className={styles.heading}>Add Book</h1>
                <form onSubmit={(e) => submit_handler(e)}>
                    <TextField
                        id="title"
                        label="Title"
                        variant="outlined"
                        fullWidth
                        style={{ backgroundColor: 'white' }}
                        value={Title}
                        onChange={(e) => setTitle(e.target.value)}
                    />
                    <TextField
                        id="author"
                        label="Author"
                        variant="outlined"
                        fullWidth
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={Author}
                        onChange={(e) => setAuthor(e.target.value)}
                    />
                    <TextField
                        id="description"
                        label="Description"
                        variant="outlined"
                        fullWidth
                        multiline
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={Description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                    <TextField
                        id="ownerID"
                        label="OwnerID"
                        variant="outlined"
                        fullWidth
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={OwnerID}
                        onChange={(e) => setOwnerID(e.target.value)}
                    />
                    <CurrencyTextField
                        id="price"
                        label="Price"
                        variant="outlined"
                        fullWidth
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        currencySymbol="₹"
                        value={Price}
                        textAlign="left"
                        onChange={(e) => setPrice(e.target.value)}
                    />
                    <TextField
                        id="stock"
                        label="Stock"
                        variant="outlined"
                        fullWidth
                        inputProps={{ min: "0", step: "1" }}
                        margin="normal"
                        type="number"
                        style={{ backgroundColor: 'white' }}
                        value={Stock}
                        onChange={(e) => setStock(e.target.value)}
                    />
                    <TextField
                        id="numberofPages"
                        label="NumberofPages"
                        variant="outlined"
                        fullWidth
                        inputProps={{ min: "0", step: "1" }}
                        type="number"
                        margin="normal"
                        style={{ backgroundColor: 'white' }}
                        value={NumberofPages}
                        onChange={(e) => setNumberofPages(e.target.value)}
                    />
                    <div className={styles.buttonArea}>
                        <label htmlFor="icon-button-photo">
                            <IconButton backgroundColor="white">
                                Image 
                                <PhotoCamera />
                            </IconButton>
                        </label>
                        <input
                            accept="image/*"
                            id="icon-button-photo"
                            value={Picture}
                            onChange={(e)=>setPicture(e.target.value)}
                            type="file"
                        />
                    </div>
                    
                
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
                            Add Book
                        </Button>
                    </div>
                </form>
            </div>
            <div className={styles.rightCol}>
                <img src={booksImg} className={styles.coverImg} alt="Books" />
            </div>
        </div>
    );
}

export default Bookform;
