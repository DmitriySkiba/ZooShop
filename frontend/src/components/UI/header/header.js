import React, {useContext, useEffect, useState} from 'react';
import Mybutton from "../button/Mybutton";
import Search from "../Search/Search";
import "./headerStyle.css"
import Logo from "../../../img/logo.svg"
import CartLogo from "../../../img/cart.svg"
import UserLogo from "../../../img/user.svg"
import BoxLogo from "../../../img/box.svg"
import Categorybtn from "../category/categorybtn";
import Modal from "../modal/modal"
import ItemsForm from "../../items/ItemForm/ItemsForm";
import UserForm from "../../items/UserForm/UserForm";
import { AuthContext } from '../../../context';
import { Link } from 'react-router-dom';
import RegForm from '../../items/UserForm/RegForm';



const Header = () => {
    const [modalActive, setModalActive] = useState(false)
    const [userModalActive, setUserModalActive] = useState(false)
    const [regModalActive, setRegModalActive] = useState(false)

    const {ignoreThat, setIgnoreThat} = useContext(AuthContext)
    const {userStatus, setUserStatus} = useContext(AuthContext)
    const {reg, setReg} = useContext(AuthContext)
    const {userLogin, setUserLogin} = useContext(AuthContext)
    const {userID, setUserID} = useContext(AuthContext)
    useEffect(() => {
        if (userModalActive){
            setUserModalActive(false)
        }
        if (regModalActive){
            setRegModalActive(false)
        }
    }, [reg]);
    const quit = () =>{
        setIgnoreThat(0)
        setUserID(0)
        setReg(false)
        setUserLogin("")
        if (userStatus === 1){
            setUserStatus(0)
        }
    }
    return (
        <div className="Header">
            <Link to="/" className="Header__logo">
                <img src={Logo} style={{height:96}}/>
                <h1 style={{padding: 25}}>ZooMarket</h1>
            </Link>

            <div className="Header__category">
                <Categorybtn></Categorybtn>
            </div>

                <div className="Header__search">
                    <Search/>
                </div>
             {!reg ? (
                <div className="Header__btn">
                <Mybutton style={{marginRight: 10}} onClick={() => setUserModalActive(true)}>
                    <img src={UserLogo} style={{width:50, height: 50}}/>
                    <p>Войти</p>
                </Mybutton>
                <Mybutton style={{marginRight: 10}} onClick={() => setRegModalActive(true)}>
                    <img src={UserLogo} style={{width:50, height: 50}}/>
                    <h6>Регистрация</h6>
                </Mybutton>
            </div>
             ) : (<div className="Header__btn">
                 {userStatus === 1 ?(
                     <Mybutton style={{marginRight: 10}} onClick={() => setModalActive(true)}>
                         <p style={{fontSize: 14}}>Добавить товар</p>
                     </Mybutton>
                 ) : console.log()}
             <Mybutton style={{marginRight: 10}} onClick={quit}>
                 <p>Выход</p>
             </Mybutton>
              <Link to={"/cart"}>
                  <Mybutton style={{marginRight: 10}}>
                      <img src={CartLogo} style={{width:50, height: 50}}/>
                      <p>Корзина</p>
                  </Mybutton>
             </Link>

         </div>)}
            
            <Modal active={modalActive} setActive={setModalActive}>
                <ItemsForm/>
            </Modal>
            <Modal active={userModalActive} setActive={setUserModalActive}>
                <UserForm/>
            </Modal>

            <Modal active={regModalActive} setActive={setRegModalActive}>
                <RegForm/>
            </Modal>
        </div>
    );
};

export default Header;