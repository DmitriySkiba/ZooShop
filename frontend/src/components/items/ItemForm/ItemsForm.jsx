import React, {useEffect, useState} from 'react';
import MyInput from "../../UI/input/MyInput";
import Mybutton from "../../UI/button/Mybutton";
import axios from "axios";
import Modal from "../../UI/modal/modal";
import "./ItemsForm.css"
import itemService from "../../../services/item.service";
import ItemsService from "../../../services/item.service";

const ItemsForm = ({nameProp ="", priceProp="", descriptionProp="", brandProp="", photoProp="", categoryProp="", subcategoryProp="", vcode}) => {
    const [name, setName] = useState("");
    const [price, setPrice] = useState(0);
    const [description, setDescription] = useState("");
    const [brand, setBrand] = useState("");
    const [photo, setPhoto] = useState("");
    useEffect(() => {
        setName(nameProp)
        setPrice(Number(priceProp))
        setDescription(descriptionProp)
        setBrand(brandProp)
        setPhoto(photoProp)
    }, [nameProp]);
    let categoryList = [
        {
            "Собаки": ["Корма", "Лакомства", "Витамины", "Аксессуары", "Переноски"]
        },
        {
            "Кошки": ["Корма", "Лакомства", "Витамины", "Аксессуары", "Переноски"]
        },
        {
            "Рыбы": ["Корма", "Витамины", "Аквариумы", "Декорации", "Аксессуары"]
        },
        {
            "Птицы": ["Корма", "Лакомства", "Витамины", "Аксессуары", "Клетки"]
        },
        {
            "Грызуны": ["Корма", "Лакомства", "Витамины", "Аксессуары", "Клетки"]
        },
        {
            "Рептилии": ["Корма", "Витамины", "Декорации", "Террариумы"]
        }
    ]
    const [active, setActive] = useState(false)
    const [message, setMessage] = useState()
    const [activeCategory, setCategory] = useState([])
    const [categoryValue, setCategoryValue] = useState("");
    const [subcategoryValue, setSubcategoryValue] = useState("");
    const handleCategoryChange = (e) => {
        const selectedCategory = e.target.value;

        const foundCategory = categoryList.find((category) =>
            Object.keys(category)[0] === selectedCategory
        );

        if (foundCategory) {
            setCategory(foundCategory[selectedCategory]);
            setCategoryValue(selectedCategory);
            setSubcategoryValue(foundCategory[selectedCategory][0]);
        } else {
            setCategory([]);
            setCategoryValue("");
            setSubcategoryValue("");
        }
    };
     async function UploadItem(e) {
         e.preventDefault()
         let selcategory = document.getElementById("category")
         let selsubcategory = document.getElementById("subcategory")
         let Newitem = {
             status: "admin",
             name: name,
             price: Number(price),
             description: description,
             photo: photo,
             category: categoryValue,
             subcategory: subcategoryValue,
             brand: brand,
         }
         try {
             console.log(Newitem)
             const response = await itemService.uploadItem(Newitem)
             setMessage(response)
         } catch (error) {
             setMessage(error)
         } finally {
             setActive(true)
         }
     }
    async function UpdateItem(e){
        e.preventDefault()
        let Updateitem = {
            status: "admin",
            name: name,
            price: Number(price),
            description: description,
            photo: photo,
            brand: brand,
            vcode: Number(vcode),
            category: categoryValue,
            subcategory: subcategoryValue
        }
        try {
            const response = await itemService.updateItem(Updateitem)
            setMessage(response)
        } catch (error) {
            setMessage(error)
        } finally {
            setActive(true)
        }
    }

    return (
        <form className={"Form"}>
            <MyInput
                name="name"
                id="name"
                value={name}
                onChange={e => setName(e.target.value)}
                type="text"
                placeholder="Название товара"
            />
            <MyInput
                name='price'
                id="price"
                value={price}
                onChange={e => setPrice(e.target.value)}
                type="number"
                placeholder="Цена"
            />
            <MyInput
                name='description'
                id="description"
                value={description}
                onChange={e => setDescription(e.target.value)}
                type="text"
                placeholder="Описание"
            />
            <MyInput
                name='brand'
                id="brand"
                value={brand}
                onChange={e => setBrand(e.target.value)}
                type="text"
                placeholder="Бренд"
            />
            <MyInput
                name='photo'
                id="photo"
                value={photo}
                onChange={e => setPhoto(e.target.value)}
                type="url"
                placeholder="Фото(url)"
            />
            <select name='category' id="category" onChange={handleCategoryChange}>
                <option value="" className={"Form__option"}>Выберите категорию</option>
                {categoryList.map((category, index) => {
                    const categoryName = Object.keys(category)[0];
                    return <option key={index} value={categoryName} className={"Form__option"}>{categoryName}</option>;

                })}
            </select>

            <select name='subcategory' id="subcategory" onChange={e => setSubcategoryValue(e.target.value)}>
                <option value="" className={"Form__option"}>Выберите подкатегорию</option>
                {activeCategory.map((subcategory, index) => (
                    <option key={index} value={subcategory} className={"Form__option"} >{subcategory}</option>
                ))}
            </select>
            {vcode === undefined? (
                <button onClick={UploadItem}>Создать товар</button>
            ) : (
                <button onClick={UpdateItem}>Обновить товар</button>
            )}
            <button type={"reset"}>Очистить</button>
            <Modal active={active} setActive={setActive} style={{width:150}}>{message}</Modal>
        </form>

    );
};

export default ItemsForm;