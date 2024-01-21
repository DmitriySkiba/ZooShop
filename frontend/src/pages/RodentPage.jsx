import React from 'react';
import Header from "../components/UI/header/header";
import {Link} from "react-router-dom";
import RodentFoodLogo from "../img/categoryimg/cat-food.svg";
import RodentThreatsLogo from "../img/categoryimg/bird-threats.svg";
import Vitamins from "../img/categoryimg/vitamins.svg";
import DogCollarLogo from "../img/categoryimg/dog-collar.svg";
import RodentCageLogo from "../img/categoryimg/rodent-cage.svg";
import Footer from "../components/UI/footer/Footer";
import TextBanner from "../components/UI/text banner/TextBanner";
import "../styles/App.css"
import ItemsForCategoryPage from "../components/items/ItemsForCategoryPage";
const RodentPage = () => {
    return (
        <div className="App">
            <Header></Header>
            <TextBanner text="Категории товаров для грызунов"/>
            <div className={"App__category"}>
                <Link to="/for-rodent/feeds" style={{marginLeft: 120}}>
                    <button className="App__btn">
                        <img src={RodentFoodLogo} style={{width: 120}}/>
                        <p>Корм</p>
                    </button>
                </Link>
                <Link to="/for-rodent/threats">
                    <button className="App__btn">
                        <img src={RodentThreatsLogo} style={{width: 28}}/>
                        <p>Лакомства</p>
                    </button>
                </Link>
                <Link to="/for-rodent/vitamins">
                    <button className="App__btn">
                        <img src={Vitamins} style={{width: 85}}/>
                        <p>Витамины</p>
                    </button>
                </Link>

                <Link to="/for-rodent/accessories">
                    <button className="App__btn">
                        <img src={DogCollarLogo} style={{width: 120}}/>
                        <p>Аксессуары</p>
                    </button>
                </Link>

                <Link to="/for-rodent/cages">
                    <button className="App__btn">
                        <img src={RodentCageLogo} style={{width: 130}}/>
                        <p>Переноски</p>
                    </button>
                </Link>

            </div>
            <ItemsForCategoryPage type={"category"} category={"Грызуны"}></ItemsForCategoryPage>
            <Footer></Footer>
        </div>
    );
};

export default RodentPage;