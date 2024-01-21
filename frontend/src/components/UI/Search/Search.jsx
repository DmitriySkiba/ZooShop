import React, {useState} from 'react';
import classes from './Search.module.css'
import Mybutton from "../button/Mybutton";
import MyInput from "../input/MyInput";
import {Link} from "react-router-dom";
const Search = () => {
    const [searchQuery, setSearchQuery]= useState("")
    return (
        <form className={classes.Form}>
            <MyInput
                value = {searchQuery}
                onChange ={e => setSearchQuery(e.target.value)}
                placeholder ={"Поиск..."}
            />
            <Link to={"/search/" + searchQuery}>
                <Mybutton style={{width:100, height:50}}>Найти</Mybutton>
            </Link>
        </form>
    );
};

export default Search;