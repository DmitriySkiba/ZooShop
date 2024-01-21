import axios from "axios";

export const itemService ={
    async getItems(sort, column, meaning){
        const response = await axios.post(
            "http://localhost:8080/allAndSort",
            {Sort: sort , Column: column, Meaning:meaning},
            {headers:{"Content-Type": "application/json"}
            })
        console.log(response)
        return (JSON.parse(response.data))
    },

    async selectItem(id){
        console.log(id)
        const response = await axios.post(
            "http://localhost:8080/select",
            {vcode: Number(id)},
            {headers:{"Content-Type": "application/json"}
            })
        return (JSON.parse(response.data))
    },

    async deleteItem(id){
        const response = await axios.post(
            "http://localhost:8080/delete",
            {Status:"admin", Vcode: Number(id)},
            {headers:{"Content-Type": "application/json"}
            })
        console.log(response.data)
        return (response.data)
    },

    async uploadItem(newItem){
        const response = await axios.post(
            "http://localhost:8080/register",
            {...newItem},
            {headers: {"Content-Type": "application/json"}})
        console.log(response.data)
        return (response.data)
    },

    async updateItem(item){
        const response = await axios.post(
            "http://localhost:8080/update",
            {...item},
            {headers: {"Content-Type": "application/json"}})
        console.log(response)
        return (response.data)
    },
    async buyItem(id, vcode){
        const response = await axios.post(
            "http://localhost:8080/buy",
            {id: id, vcode: vcode},
            {headers: {"Content-Type": "application/json"}})
        console.log(response)
        return (response.data)
    },
    async addToCartItem(id, vcode){
        const response = await axios.post(
            "http://localhost:8080/addTobasket",
            {id: id, vcode: vcode},
            {headers: {"Content-Type": "application/json"}})
        console.log(response)
        return (response.data)
    },
    async deleteCartItem(id, vcode){
        const response = await axios.post(
            "http://localhost:8080/deleteInBasket",
            {id: id, vcode: vcode},
            {headers: {"Content-Type": "application/json"}})
        console.log(response)
        return (response.data)
    },
    async getCartItems(id){
        const response = await axios.post(
            "http://localhost:8080/viewBasket",
            {id: id},
            {headers: {"Content-Type": "application/json"}})
        if (response.data !== "Basket is empty"){
            return (JSON.parse(response.data))
        }else return []
    },
    async getStoryBought(id){
        const response = await axios.post(
            "http://localhost:8080/storyBought",
            {id: id},
            {headers: {"Content-Type": "application/json"}})
        console.log(JSON.parse(response.data))
        return (JSON.parse(response.data))
    }
}
export default itemService