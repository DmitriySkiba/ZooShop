import React, {useEffect} from 'react';
import "./cartItem.css"
const CartItem = (props) => {

    return (
        <div className={"cartItem"}>
            <h2>{props.props.name}</h2>
            <h1>Цена: {props.props.price} ₽</h1>
        </div>
    );
};

export default CartItem;