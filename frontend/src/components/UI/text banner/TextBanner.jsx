import React from 'react';
import classes from "./TextBanner.module.css"
const TextBanner = (items) => {
    return (
        <div className={classes.TextBanner}>
            <h1>{items.text}</h1>
        </div>
    );
};

export default TextBanner;