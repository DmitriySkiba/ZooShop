import React from 'react';
import Header from "../components/UI/header/header";
import {Link} from "react-router-dom";
import BirdFoodLogo from "../img/categoryimg/bird-feed.svg";
import BirdThreatsLogo from "../img/categoryimg/bird-threats.svg";
import Vitamins from "../img/categoryimg/vitamins.svg";
import DogCollarLogo from "../img/categoryimg/dog-collar.svg";
import BirdCageLogo from "../img/categoryimg/bird-cage.svg";
import Footer from "../components/UI/footer/Footer";
import TextBanner from "../components/UI/text banner/TextBanner";
import "../styles/App.css"
import ItemsForCategoryPage from "../components/items/ItemsForCategoryPage";
const BirdPage = () => {
    return (
        <div className="App">
            <Header></Header>

            <TextBanner text="Категории товаров для птиц"/>

            <div className={"App__category"}>
                <Link to="/for-bird/feeds" style={{marginLeft: 120}}>
                    <button className="App__btn">
                        <img src={BirdFoodLogo} style={{width: 100}}/>
                        <p>Корм</p>
                    </button>
                </Link>
                <Link to="/for-bird/threats">
                    <button className="App__btn">
                        <img src={BirdThreatsLogo} style={{width: 28}}/>
                        <p>Лакомства</p>
                    </button>
                </Link>
                <Link to="/for-bird/vitamins">
                    <button className="App__btn">
                        <img src={Vitamins} style={{width: 85}}/>
                        <p>Витамины</p>
                    </button>
                </Link>

                <Link to="/for-bird/accessories">
                    <button className="App__btn">
                        <img src={DogCollarLogo} style={{width: 120}}/>
                        <p>Аксессуары</p>
                    </button>
                </Link>

                <Link to="/for-bird/cages">
                    <button className="App__btn">
                        <img src={BirdCageLogo} style={{width: 120}}/>
                        <p>Клетки</p>
                    </button>
                </Link>

            </div>
            <ItemsForCategoryPage type={"category"} category={"Птицы"}></ItemsForCategoryPage>
            <Footer></Footer>
        </div>
    );
};

export default BirdPage;