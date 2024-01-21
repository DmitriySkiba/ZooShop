import React, {useContext, useState} from 'react';
import MyInput from "../../UI/input/MyInput";
import Mybutton from "../../UI/button/Mybutton";
import axios from "axios";
import Modal from "../../UI/modal/modal";
import "./UserForm.css"
import { AuthContext } from '../../../context';
import UserService from "../../../services/user.service";

const UserForm = () => {
    const [active, setActive] = useState(false)
    const [message, setMessage] = useState('')
    const [loginValue, setLoginValue] = useState("")
    const [passwordValue, setPasswordValue] = useState("")

    const {reg, setReg} = useContext(AuthContext)
    const {userLogin, setUserLogin} = useContext(AuthContext)
    const {userStatus, setUserStatus} = useContext(AuthContext)
    const {userID, setUserID} = useContext(AuthContext)
    const login = async (e) => {
        e.preventDefault()
        if (loginValue !== "" && passwordValue !== ""){
            try {
                const responseLogin = await UserService.Login(loginValue, passwordValue)
                if (responseLogin  !== "Invalid login or password"){
                    setMessage("Вы успешно вошли!")
                    setUserID(responseLogin)
                    setUserLogin(loginValue)
                    setReg(true)
                    const responseStatus = await UserService.getStatus(loginValue)
                    if (responseStatus === 1){
                        setUserStatus(responseStatus)
                    }
                }else{
                    setMessage("Неверный логин или пароль!")
                }
            } catch (error) {
                setMessage("Error")
            } finally {
                setActive(true)
            }
        }else {
            setMessage("Поля не должны быть пустыми!")
            setActive(true)
        }
    }
    return (
        <form className={"Form"}>
            <MyInput
                name="login"
                id="login"
                type="text"
                onChange={e => {setLoginValue(e.target.value)}}
                placeholder="Логин"
            />
            <MyInput
                name='password'
                id="password"
                type="password"
                onChange={e => {setPasswordValue(e.target.value)}}
                placeholder="Пароль"
            />
            <br/>
            <button style={{marginTop:20}} type={"reset"}>Очистить</button>
            <button onClick={login} type={"submit"}>Войти</button>
            <Modal active={active} setActive={setActive} style={{width:150}}>{message}</Modal>
        </form>

    );
};

export default UserForm;