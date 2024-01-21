import React from 'react';
import Header from "../components/UI/header/header";
import {Link} from "react-router-dom";
import CatFoodLogo from "../img/categoryimg/cat-food.svg";
import CatThreatsLogo from "../img/categoryimg/cat-threats.svg";
import Vitamins from "../img/categoryimg/vitamins.svg";
import DogCollarLogo from "../img/categoryimg/dog-collar.svg";
import CarrierLogo from "../img/categoryimg/carrier.svg";
import Footer from "../components/UI/footer/Footer";
import TextBanner from "../components/UI/text banner/TextBanner";
import "../styles/App.css"
import ItemsForCategoryPage from "../components/items/ItemsForCategoryPage";
const CatPage = () => {
    return (
        <div className="App">
            <Header></Header>

            <TextBanner text="Категории товаров для кошек"/>

            <div className={"App__category"}>
                <Link to="/for-cat/feeds" style={{marginLeft: 120}}>
                    <button className="App__btn">
                        <img src={CatFoodLogo} style={{width: 120}}/>
                        <p>Корм</p>
                    </button>
                </Link>
                <Link to="/for-cat/threats">
                    <button className="App__btn">
                        <img src={CatThreatsLogo} style={{width: 120}}/>
                        <p>Лакомства</p>
                    </button>
                </Link>
                <Link to="/for-cat/vitamins">
                    <button className="App__btn">
                        <img src={Vitamins} style={{width: 85}}/>
                        <p>Витамины</p>
                    </button>
                </Link>

                <Link to="/for-cat/accessories">
                    <button className="App__btn">
                        <img src={DogCollarLogo} style={{width: 120}}/>
                        <p>Аксессуары</p>
                    </button>
                </Link>

                <Link to="/for-cat/carries">
                    <button className="App__btn">
                        <img src={CarrierLogo} style={{width: 120}}/>
                        <p>Переноски</p>
                    </button>
                </Link>

            </div>
            <ItemsForCategoryPage type={"category"} category={"Кошки"}></ItemsForCategoryPage>
            <Footer></Footer>
        </div>
    );
};

export default CatPage;