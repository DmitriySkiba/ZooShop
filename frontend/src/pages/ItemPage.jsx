import React, {useContext, useEffect, useState} from 'react';
import Header from "../components/UI/header/header";
import Footer from "../components/UI/footer/Footer";
import {Link, useParams} from "react-router-dom";
import ItemsService from "../services/item.service";
import TextBanner from "../components/UI/text banner/TextBanner";
import "../styles/ItemPage.css"
import Modal from "../components/UI/modal/modal";
import ItemsForm from "../components/items/ItemForm/ItemsForm";
import Item from "../components/items/Item";
import { AuthContext } from '../context';
import PurchaseForm from "../components/items/PurchaseForm/purchaseForm";

const ItemPage = () => {
    const {userStatus, setUserStatus} = useContext(AuthContext)
    const {reg, setReg} = useContext(AuthContext)
    const {userLogin, setUserLogin} = useContext(AuthContext)
    const {userID, setUserID} = useContext(AuthContext)
    const {ignoreThat, setIgnoreThat} = useContext(AuthContext)

    const {id} = useParams();
    const [modalDeleteActive, setModalDeleteActive]=useState(false)
    const [modalUpdateActive, setModalUpdateActive] = useState(false)
    const [modalMessageActive, setModalMessageActive]=useState(false)
    const [modalPurchaseActive, setModalPurchaseActive] = useState(false)
    const [regMessage, setRegMessage] = useState("Сначала войдите в аккаунт!")
    const [regModal, setRegModal] = useState(false)
    const [message, setMessage] = useState("")
    const [error, setError] = useState(' ')
    const [isLoading, setIsLoading] = useState(true)
    const [item, setItem]=useState([])


    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await ItemsService.selectItem(id);
                setItem(response)
            } catch (error) {
                setError(error)
            } finally {
                setIsLoading(false)
            }
        };

        let ignore = fetchData();
    }, []);

    const deleteItem = async () =>{
        try {
            const response = await ItemsService.deleteItem(id);
            setMessage(response)
        } catch (error) {
            setMessage(error)
        } finally {
            setModalDeleteActive(false)
            setModalMessageActive(true)
        }  
    }
    const addToCartItem = async () =>{
        if (reg){
            try {
                const response = await ItemsService.addToCartItem(userID, item["vcode"]);
                if (response === "New goods added successfully into basket"){
                    setMessage("Товар успешно добавлен в корзину")
                }
            } catch (error) {
                setMessage(error)
            } finally {
                setModalMessageActive(true)
            }
        }else{
            setRegModal(true)
        }
    }

    const itemBuy = async () => {
        !reg ? (
            setRegModal(true)
        ) : (
            setModalPurchaseActive(true)
        )
    }  
    
    return (
        <div className="App">
            <Header></Header>
            {error && <div className={"Error"}>{error}</div>}
            {isLoading ? (
                    <div className={"Message"}>Загрузка...</div>
                )
                : item ? (
                    <div>
                        <TextBanner text={item["name"]}/>
                        {userStatus === 1 ? (
                        <div className={"admin"}>
                            <button className={"itemPage__button"} onClick={() => setModalUpdateActive(true)}>Изменить</button>
                            <button className={"itemPage__button"} onClick={() => setModalDeleteActive(true)}>Удалить</button>
                        </div>) : (console.log())}
                        <div className={"itemPage"}>
                            <img src={item["photo"]} className={"itemPage__img"}/>
                            <div className={"itemPage__body"}>
                                <div className={"itemPage__price"}>Цена: {item["price"]} ₽</div>
                                <button className={"itemPage__button"} onClick={() => itemBuy()}>Купить</button>
                                <button className={"itemPage__button"} onClick={addToCartItem}>Добавить в корзину</button>
                                <div className={"itemPage__desciption"}>
                                    <h3>Описание</h3>
                                    {item["description"]}
                                </div>
                            </div>
                        </div>
                    </div>
                ) : (
                    <div className={"Message"}>Товар не найден</div>
                )}
            <Footer></Footer>
            {userStatus === 1 ? (
                <div>
                    <Modal active={modalDeleteActive} setActive={setModalDeleteActive}>
                        <h2>Вы точно хотите удалить?</h2>
                        <button className={"itemPage__button"} onClick={() => deleteItem()}>Абсолютно</button>
                        <button className={"itemPage__button"} onClick={() => setModalDeleteActive(false)}>Нет</button>
                    </Modal>
                    <Modal active={modalUpdateActive} setActive={setModalUpdateActive}>
                        <ItemsForm
                            nameProp={item["name"]}
                            priceProp={item["price"]}
                            descriptionProp={item["description"]}
                            brandProp={item["brand"]}
                            photoProp={item["photo"]}
                            categoryProp={item["category"]}
                            subcategoryProp={item["subcategory"]}
                            vcode={item["vcode"]}
                        />
                    </Modal>
                </div>
            ) : (console.log())}

            <Modal active={modalMessageActive} setActive={setModalMessageActive}>
                <h2>{message}</h2>
                <Link to={"/"}>
                    <button className={"itemPage__button"}>Вернуться в главное меню.</button>
                </Link>
            </Modal>
            <Modal active={regModal} setActive={setRegModal}>{regMessage}</Modal>
            {reg ? (
                <Modal active={modalPurchaseActive} setActive={setModalPurchaseActive}>
                    <PurchaseForm itemsArr={[item]}/>
                </Modal>
            ) : (console.log())}
        </div>
    );
};

export default ItemPage;