import React, {useEffect, useState} from 'react';
import Item from "./Item";
import ItemsService from "../../services/item.service"
import "./Items.css"
import TextBanner from "../UI/text banner/TextBanner";
import MySelect from "../UI/mySelect/mySelect";
const ItemsForCategoryPage = (props) => {
    const [items, setItems] = useState([])
    const [brands, setBrands] = useState([])
    const [error, setError] = useState(' ')
    const [isLoading, setIsLoading] = useState(true)
    const [selectedSort, setSelectedSort] = useState(" ")
    const [selectedBrand, setSelectedBrand] = useState("Выбрать бренд")

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await ItemsService.getItems("date", props.type, props.category);
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
        const updateBrands = () => {
            const brandsSet = new Set();
            items.forEach((item) => {
                brandsSet.add(item["brand"].toLowerCase());
            });
            const brandsList = Array.from(brandsSet).map((brand) => ({
                value: brand,
                name: brand[0].toUpperCase()+brand.slice(1),
            }));
            brandsList.unshift({value: "Выбрать бренд", name: "Выбрать бренд (по умолч.)"})
            setBrands(brandsList);
        };
        if (items) {
            updateBrands();
        }
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
            <TextBanner text={"Список товаров:"}/>
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
                            value={selectedBrand}
                            onChange={setSelectedBrand}
                            defaultValue={"Выбрать бренд"}
                            options={brands}
                        />
                    )}
                <div className={"ItemsList"}>
                    {error && <div className={"Error"}>{error}</div>}
                    {isLoading ? (
                            <div className={"Message"}>Загрузка...</div>
                        )
                        : items ? (
                            items.map((item, index) => {
                                if (selectedBrand === "Выбрать бренд"){
                                    return Item(item)
                                }else if(selectedBrand === item["brand"].toLowerCase()){
                                    return Item(item)
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

export default ItemsForCategoryPage;