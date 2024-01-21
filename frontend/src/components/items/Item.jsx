import React from 'react';
import "./item.css"
import {Link} from "react-router-dom";
const Item = (props) => {
    return (
        <Link to={"/item/"+props["vcode"]} className={"item"}>
                <img src={props.photo} className={"item__img"}/>
                <h2>{props.name}</h2>
                <h1>Цена: {props.price} ₽</h1>
        </Link>
    );
};

export default Item;