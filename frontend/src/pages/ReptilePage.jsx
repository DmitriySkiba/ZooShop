import React from 'react';
import Header from "../components/UI/header/header";
import {Link} from "react-router-dom";
import FoodLogo from "../img/categoryimg/cat-food.svg";
import Vitamins from "../img/categoryimg/vitamins.svg";
import DecorationsLogo from "../img/categoryimg/decoration.svg";
import TerrariumLogo from "../img/categoryimg/terrarium.svg";
import Footer from "../components/UI/footer/Footer";
import TextBanner from "../components/UI/text banner/TextBanner";
import "../styles/App.css"
import ItemsForCategoryPage from "../components/items/ItemsForCategoryPage";
const ReptilePage = () => {
    return (
        <div className="App">
            <Header></Header>
            <TextBanner text="Категории товаров для рептилий"/>
            <div className={"App__category"}>
                <Link to="/for-reptile/feeds" style={{marginLeft: 240}}>
                    <button className="App__btn">
                        <img src={FoodLogo} style={{width: 120}}/>
                        <p>Корм</p>
                    </button>
                </Link>
                <Link to="/for-reptile/vitamins">
                    <button className="App__btn">
                        <img src={Vitamins} style={{width: 85}}/>
                        <p>Витамины</p>
                    </button>
                </Link>
                <Link to="/for-reptile/decorations">
                    <button className="App__btn">
                        <img src={DecorationsLogo} style={{width: 120}}/>
                        <p>Декорации</p>
                    </button>
                </Link>
                <Link to="/for-reptile/decorations">
                    <button className="App__btn">
                        <img src={TerrariumLogo} style={{width: 130}}/>
                        <p>Террариумы</p>
                    </button>
                </Link>
            </div>
            <ItemsForCategoryPage type={"category"} category={"Рептилии"}></ItemsForCategoryPage>
            <Footer></Footer>
        </div>
    );
};

export default ReptilePage;