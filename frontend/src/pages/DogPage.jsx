import React from 'react';
import "../styles/App.css"
import Header from "../components/UI/header/header";
import Footer from "../components/UI/footer/Footer";
import DogFoodLogo from "../img/categoryimg/dog-food.svg"
import DogThreatsLogo from "../img/categoryimg/dog-threats.svg"
import Vitamins from "../img/categoryimg/vitamins.svg"
import DogCollarLogo from "../img/categoryimg/dog-collar.svg"
import CarrierLogo from "../img/categoryimg/carrier.svg"
import {Link} from "react-router-dom";
import TextBanner from "../components/UI/text banner/TextBanner";
import ItemsForCategoryPage from "../components/items/ItemsForCategoryPage";
import ItemsForm from "../components/items/ItemForm/ItemsForm";
const DogPage = () => {
    return (
        <div className="App">
            <Header></Header>
            <TextBanner text="Категории товаров для собак"/>
            <div className={"App__category"}>
                <Link to="/for-dog/feeds" style={{marginLeft: 120}}>
                    <button className="App__btn">
                        <img src={DogFoodLogo} style={{width: 120}}/>
                        <p>Корм</p>
                    </button>
                </Link>
                <Link to="/for-dog/threats">
                    <button className="App__btn">
                        <img src={DogThreatsLogo} style={{width: 120}}/>
                        <p>Лакомства</p>
                    </button>
                </Link>
                <Link to="/for-dog/vitamins">
                    <button className="App__btn">
                        <img src={Vitamins} style={{width: 85}}/>
                        <p>Витамины</p>
                    </button>
                </Link>

                <Link to="/for-dog/accessories">
                    <button className="App__btn">
                        <img src={DogCollarLogo} style={{width: 120}}/>
                        <p>Аксессуары</p>
                    </button>
                </Link>

                <Link to="/for-dog/carries">
                    <button className="App__btn">
                        <img src={CarrierLogo} style={{width: 120}}/>
                        <p>Переноски</p>
                    </button>
                </Link>

            </div>
            <ItemsForCategoryPage type={"category"} category={"Собаки"}></ItemsForCategoryPage>
            <Footer></Footer>
        </div>
    );
};

export default DogPage;