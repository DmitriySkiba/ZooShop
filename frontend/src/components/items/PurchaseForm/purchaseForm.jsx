import React, {useContext, useEffect, useState} from 'react';
import {AuthContext} from "../../../context";
import UserService from "../../../services/user.service";
import "./purchaseForm.css"
import ItemService from "../../../services/item.service";
import Modal from "../../UI/modal/modal";
const PurchaseForm = ({itemsArr}) => {
    const {reg, setReg} = useContext(AuthContext)
    const {userLogin, setUserLogin} = useContext(AuthContext)
    const {userID, setUserID} = useContext(AuthContext)
    const {ignoreThat, setIgnoreThat} = useContext(AuthContext)
    const [items, setItems]= useState([])
    const [priceWithoutSale, setPriceWithoutSale] = useState(0)
    const [priceWithSale, setPriceWithSale] = useState(0)
    const [sale, setSale]=useState(0)
    const [active, setActive] = useState(false)
    const [message, setMessage] = useState("")


    useEffect(() => {
        console.log("Вход1")
        console.log(itemsArr)
        setItems(itemsArr)
    }, [itemsArr]);
    useEffect( () => {
        console.log("Вход2")
        console.log(items)
        let sumWithoutSale = 0
        let sumWithSale = 0
        items.map((item) =>{
            sumWithoutSale += item["price"]
            sumWithSale += (item["price"]-item["price"]*sale)
        })
        setPriceWithoutSale(sumWithoutSale)
        setPriceWithSale(sumWithSale)
    }, [items]);

    useEffect( () => {
        const getSale = async ()=>{
            try {
                const responseSale = await UserService.getSale(userLogin)
                if (responseSale !== "Error executing query") {
                    setSale(responseSale)
                } else {
                    setMessage("Sale error: " + responseSale)
                }
            } catch (error) {
                setMessage(error)
                setActive(true)
            }
        }
        let ignore = getSale()
    }, [reg, ignoreThat]);

    const buyItem = async () =>{
        try {
            const responseBuy = await UserService.buyItem(userLogin, priceWithSale.toFixed(0))
            if (responseBuy === "Success") {
                for (var i = 0; i < items.length; i++){
                    try{
                        const responseBuyItem = await ItemService.buyItem(userID, items[i]["vcode"])
                    }catch (error){
                        setMessage(error)
                        setActive(true)
                        return;
                    }
                }
                setMessage("Товар был успешно оплачен")
                setIgnoreThat(ignoreThat+1)
            } else {
                setMessage("Buy error: " + responseBuy)
            }
        } catch (error) {
            setMessage(error)
        } finally {
            setActive(true)
        }
    }

    return (
        <div>
            <h3>Оплата товаров:</h3>
            {items.map((item) => {
                return <h4>
                    {item["name"]}: {item["price"]} ₽
                </h4>
            })}
            <br/>
            <h3>Общая цена: {priceWithoutSale} ₽</h3>
            <h3>Ваша скидка: {(priceWithoutSale-priceWithSale).toFixed(2)} ₽</h3>
            <h2>Окончательная цена: {priceWithSale.toFixed(2)} ₽</h2>
            <button className={"purchaseForm__button"} onClick={buyItem}>Купить</button>
            <Modal active={active} setActive={setActive}>{message}</Modal>
        </div>
    );
};

export default PurchaseForm;