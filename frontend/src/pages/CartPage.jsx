import React, {useContext, useEffect, useState} from 'react';
import Header from "../components/UI/header/header";
import Footer from "../components/UI/footer/Footer";
import {AuthContext} from "../context";
import ItemsService from "../services/item.service";
import TextBanner from "../components/UI/text banner/TextBanner";
import "../styles/Cart.css"
import CartItem from "../components/items/cartItem";
import PurchaseForm from "../components/items/PurchaseForm/purchaseForm";
import Modal from "../components/UI/modal/modal";
import {Link} from "react-router-dom";

const CartPage = () => {
    const {ignoreThat, setIgnoreThat} = useContext(AuthContext)
    const {userStatus, setUserStatus} = useContext(AuthContext)
    const {reg, setReg} = useContext(AuthContext)
    const {userLogin, setUserLogin} = useContext(AuthContext)
    const {userID, setUserID} = useContext(AuthContext)
    const [items, setItems] = useState([])
    const [error, setError] = useState("")
    const [isLoading, setIsLoading] = useState(true)
    const [message, setMessage] = useState("")
    const [modalMessageActive, setModalMessageActive] = useState(false)
    const [selectedItemsData, setSelectedItemsData] = useState([])
    useEffect(() => {
        const fetchCart = async () =>{
            try {
                const response = await ItemsService.getCartItems(userID);
                setItems(response);
            } catch (error) {
                setError(error.message);
            } finally {
                setIsLoading(false);
            }
        }
        let ignore = fetchCart()
    }, []);

    const handleCheckboxChange = async (event) => {
        const { value, checked } = event.target;

        if (checked) {
            try {
                const response = await ItemsService.selectItem(Number(value));
                setSelectedItemsData((prevSelectedItemsData) => [
                    ...prevSelectedItemsData,
                    {...response}
                ]);
            } catch (error) {
                console.error(error);
            }
        } else {
            setSelectedItemsData((prevSelectedItemsData) =>
                prevSelectedItemsData.filter((itemData) => itemData.vcode !== Number(value))
            );
        }
    };

    const deleteItemFromCart = async (event)=>{
        const value = event.target.value
        setItems((prevItems) =>
        prevItems.filter((item) => item.vcode !== Number(value)))
        setSelectedItemsData((prevSelectedItemsData) =>
            prevSelectedItemsData.filter((itemData) => itemData.vcode !== Number(value))
        );

        try {
            const response = await ItemsService.deleteCartItem(userID, Number(value));
            setMessage(response)
        } catch (error) {
            setMessage(error)
        }finally {
            setModalMessageActive(true)
        }
    }

    useEffect(() => {
        selectedItemsData.map(async (item) => {
            try {
                const response = await ItemsService.deleteCartItem(userID, item.vcode);
                setItems((prevItems) =>
                    prevItems.filter((item1) => item1.vcode !== item.vcode))
                setSelectedItemsData((prevSelectedItemsData) =>
                    prevSelectedItemsData.filter((itemData) => itemData.vcode !== item.vcode)
                );
            } catch (error) {
                setMessage(error)
                setModalMessageActive(true)
            }
        })
    }, [ignoreThat]);


    useEffect(() => {
        console.log(selectedItemsData)
    }, [selectedItemsData]);

    return (
        <div className="App">
            <Header></Header>
            <Link to={"/bought-history"}>
                <button className={"btn"}>История покупок</button>
            </Link>
            <TextBanner text={"Корзина:"}/>
            {error && <div className={"Error"}>{error}</div>}
            {isLoading ? (
                <div className={"Message"}>Загрузка...</div>
            ) : (
                <div className={"Cart__body"}>
                    <div className={"Cart__items"}>
                        {items.length ? (
                            items.map((item) => (
                                <div className={"Cart__item"}>
                                    <input type={"checkbox"} value={item.vcode} onChange={handleCheckboxChange}/>
                                    <CartItem props={item}/>
                                    <button value={item.vcode} onClick={deleteItemFromCart}>Удалить</button>
                                </div>
                            ))
                        ) : (
                            <div className={"Message"}>Товары не найдены</div>
                        )}
                    </div>
                    <div className={"Cart__BoughtForm"}>
                        <PurchaseForm itemsArr={selectedItemsData}/>
                    </div>
                </div>)}
            <Footer></Footer>
            <Modal active={modalMessageActive} setActive={setModalMessageActive}>
                {message}
            </Modal>
        </div>
    );
};

export default CartPage;
