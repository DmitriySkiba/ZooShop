import React, {useEffect, useState} from 'react';
import Item from "./Item";
import ItemsService from "../../services/item.service"
import "./Items.css"
import TextBanner from "../UI/text banner/TextBanner";
import MySelect from "../UI/mySelect/mySelect";
const ItemsForBrandsPage = (props) => {
    const [items, setItems] = useState([])
    const [category, setCategory] = useState([])
    const [error, setError] = useState(' ')
    const [itemsExist, setItemsExist] = useState(false)
    const [isLoading, setIsLoading] = useState(true)
    const [selectedSort, setSelectedSort] = useState(" ")
    const [selectedCategory, setSelectedCategory] = useState("Выбрать категорию")

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await ItemsService.getItems("date", "", "");

                setItems(response);
            } catch (error) {
                setError(error.message);
            } finally {
                setIsLoading(false);
            }
        };

        let ignore = fetchData();
    }, []);

    useEffect(() => {
        const updateCategories = () => {
            const categorySet = new Set();
            items.forEach((item) => {
                if (props.brand.toLowerCase() === item["brand"].toLowerCase()){
                    categorySet.add(item["category"]);
                }
            });
            const categoryList = Array.from(categorySet).map((category) => ({
                value: category,
                name: category,
            }));
            categoryList.unshift({value: "Выбрать категорию", name: "Выбрать категорию (по умолч.)"})
            setCategory(categoryList);
        };
        updateCategories();
    }, [items]);
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
            <TextBanner text={"Список товаров бренда "+props.brand+":"}/>
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
                {isLoading ? (
                        <></>
                    )
                    : (<MySelect
                            value={selectedCategory}
                            onChange={setSelectedCategory}
                            defaultValue={"Выбрать категорию"}
                            options={category}
                        />
                    )}
                <div className={"ItemsList"}>
                    {error && <div className={"Error"}>{error}</div>}
                    {isLoading ? (
                            <div className={"Message"}>Загрузка...</div>
                        )
                        : items.length ? (
                            items.map((item, index) => {
                                if (selectedCategory === "Выбрать категорию" && props.brand.toLowerCase() === item["brand"].toLowerCase()) {
                                    if (itemsExist === false){
                                        setItemsExist(true)
                                    }
                                    return Item(item)
                                } else if (selectedCategory === item["category"] && props.brand.toLowerCase() === item["brand"].toLowerCase()) {
                                    if (itemsExist === false){
                                        setItemsExist(true)
                                    }
                                    return Item(item)
                                } else if (itemsExist === false && index === items.length-1) {
                                    return <div className={"Message"}>Товары не найдены</div>
                                }
                            })
                        ) : (
                            <div className={"Message"}>Товары не найдены</div>
                        )}
                </div>
            </div>
        </div>
    );
};

export default ItemsForBrandsPage ;