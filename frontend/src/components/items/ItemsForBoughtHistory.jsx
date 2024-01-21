import React, {useContext, useEffect, useState} from 'react';
import ItemsService from "../../services/item.service";
import TextBanner from "../UI/text banner/TextBanner";
import MySelect from "../UI/mySelect/mySelect";
import Item from "./Item";
import {AuthContext} from "../../context";

const ItemsForBoughtHistory = () => {
    const [items, setItems] = useState([])
    const [brands, setBrands] = useState([])
    const [error, setError] = useState(' ')
    const [isLoading, setIsLoading] = useState(true)
    const [selectedSort, setSelectedSort] = useState(" ")
    const [selectedBrand, setSelectedBrand] = useState("Выбрать бренд")

    const {userID, setUserID} = useContext(AuthContext)

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await ItemsService.getStoryBought(userID);
                setItems(response);
            } catch (error) {
                setError(error.message);
            } finally {
                setIsLoading(false);
            }
        };
        if (userID !== 0){
            let ignore = fetchData();
        }else {
            setError("Error: undefined user")
            setIsLoading(false);
        }
    }, []);

    const sortItems = (sort) =>{
        setSelectedSort(sort)
        if (sort === "nameASC"){
            setItems([...items].sort((a, b) => a["name"].localeCompare(b["name"])))
        }else if(sort === "nameDSC"){
            setItems([...items].sort((a, b) => b["name"].localeCompare(a["name"])))
        }else if(sort === "priceASC"){
            setItems([...items].sort((a, b) => b["price"] - a["price"]))
        }else if(sort === "priceDSC"){
            setItems([...items].sort((a, b) => a["price"] - b["price"]))
        }

    }
    return (
        <div>
            <TextBanner text={"История покупок:"}/>
            <div className={"Items"}>
                <MySelect
                    value={selectedSort}
                    onChange={sortItems}
                    defaultValue={"Сортировка по"}
                    options={[
                        {value: 'nameASC', name: 'По названию (A-Z)'},
                        {value: 'nameDSC', name: 'По названию (Z-A)'},
                        {value: 'priceDSC', name: 'По цене (возр)'},
                        {value: 'priceASC', name: 'По цене (убыв)'},
                    ]}
                />
                <div className={"ItemsList"}>
                    {error && <div className={"Error"}>{error}</div>}
                    {isLoading ? (
                            <div className={"Message"}>Загрузка...</div>
                        )
                        : items ? (
                            items.map((item, index) => {
                                    return Item(item)
                            })
                        ) : (
                            <div className={"Message"}>Товары не найдены</div>
                        )}
                </div>
            </div>
        </div>
    );
};

export default ItemsForBoughtHistory;