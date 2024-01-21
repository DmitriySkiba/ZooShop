import axios from "axios";
import header from "../components/UI/header/header";

export const userService ={
    async Login(login, password){
        const response = await axios.post(
            "http://localhost:8081/authorisation", {login: login, password: password},
            {headers: {"Content-Type": "application/x-www-form-urlencoded"}})
        return (response.data)
    },

    async Registration(login, password, name){
        const response = await axios.post(
            "http://localhost:8081/registration",
            {login: login, password: password, name: name},
            {headers: {"Content-Type": "application/x-www-form-urlencoded"}})
        return (response.data)
    },
    async getStatus(login){
        const response = await axios.post(
            "http://localhost:8081/check_status",
            {login: login},
            {headers: {"Content-Type": "application/x-www-form-urlencoded"}})
        return (response.data)
    },
    async getSale(login){
        const response = await axios.post(
            "http://localhost:8081/sale",
            {login: login},
            {headers: {"Content-Type": "application/x-www-form-urlencoded"}})
        console.log("getSale:"+String(response.data))
        return (response.data)
    },
    async buyItem(login, sum){
        const response = await axios.post(
            "http://localhost:8081/bought",
            {login: login, sum: sum},
            {headers: {"Content-Type": "application/x-www-form-urlencoded"}})
        console.log(response)
        return (response.data)
    },

}
export default userService