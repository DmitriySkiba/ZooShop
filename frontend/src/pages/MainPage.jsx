import React from 'react';
import Header from "../components/UI/header/header";
import banner from "../img/banner.svg";
import DogLogo from "../img/dog.svg";
import CatLogo from "../img/cat.svg";
import FishLogo from "../img/fish.svg";
import ParrotLogo from "../img/parrot.svg";
import MiceLogo from "../img/mice.svg";
import LizardLogo from "../img/lizard.svg";
import RCLogo from "../img/brands/royal-canin.svg";
import PedigLogo from "../img/brands/pedigree.svg";
import PurinaLogo from "../img/brands/purina.svg";
import WhiskasLogo from "../img/brands/whiskas.svg";
import KiteLogo from "../img/brands/kite.svg";
import ShebaLogo from "../img/brands/sheba.svg";
import DreamLogo from "../img/brands/dreamies.svg";
import MnyamsLogo from "../img/brands/mnyams.svg";
import BritLogo from "../img/brands/brit.svg";
import CesarLogo from "../img/brands/cesar.svg";
import ProbalLogo from "../img/brands/probalance.svg";
import EightinOneLogo from "../img/brands/eight.svg";
import PerfectfitLogo from "../img/brands/perfectfit.svg";
import EukaLogo from "../img/brands/euka.svg";
import "../styles/App.css"
import {Link} from "react-router-dom";
import Footer from "../components/UI/footer/Footer";
import TextBanner from "../components/UI/text banner/TextBanner";

const MainPage = () => {
    return (
        <div className="App">
            <Header></Header>
            <div className="App__body">
                <div>
                    <img src={banner} className="App__img"/>
                </div>
                <div className="App__category">
                    <Link to="/for-dog">
                        <button className="App__btn">
                            <img src={DogLogo} style={{width: 135}}/>
                            <p>Собаки</p>
                        </button>
                    </Link>
                    <Link to="/for-cat">
                        <button className="App__btn">
                            <img src={CatLogo} style={{width: 135}}/>
                            <p>Кошки</p>
                        </button>
                    </Link>

                    <Link to="for-fish">
                        <button className="App__btn">
                            <img src={FishLogo} style={{width: 130}}/>
                            <p>Рыбы</p>
                        </button>
                    </Link>
                    <Link to="for-bird">
                        <button className="App__btn">
                            <img src={ParrotLogo} style={{width: 135}}/>
                            <p>Птицы</p>
                        </button>
                    </Link>
                    <Link to="for-rodent">
                        <button className="App__btn">
                            <img src={MiceLogo} style={{width: 120}}/>
                            <p>Грызуны</p>
                        </button>
                    </Link>
                    <Link to="for-reptile">
                        <button className="App__btn">
                            <img src={LizardLogo} style={{width: 115}}/>
                            <p>Рептилии</p>
                        </button>
                    </Link>
                </div>

                <TextBanner text="Популярные бренды"/>

                <div className="App__brands">
                    <div style={{display: "flex"}}>
                        <Link to='/royalcanin'>
                            <button className="App__brands_btn" style={{marginLeft: 35}}>
                                <img src={RCLogo}/>
                            </button>
                        </Link>
                        <Link to="/pedigree">
                            <button className="App__brands_btn">
                                <img src={PedigLogo}/>
                            </button>
                        </Link>
                        <Link to="/purina">
                            <button className="App__brands_btn">
                                <img src={PurinaLogo}/>
                            </button>
                        </Link>
                        <Link to="/whiskas">
                            <button className="App__brands_btn">
                                <img src={WhiskasLogo}/>
                            </button>
                        </Link>

                        <Link to="/kitekat">
                            <button className="App__brands_btn">
                                <img src={KiteLogo}/>
                            </button>
                        </Link>

                        <Link to="/sheba">
                            <button className="App__brands_btn">
                                <img src={ShebaLogo}/>
                            </button>
                        </Link>

                        <Link to="/dreamies">
                            <button className="App__brands_btn">
                                <img src={DreamLogo}/>
                            </button>
                        </Link>

                    </div>

                    <div style={{display: "flex"}}>
                        <Link to="/mnyams">
                            <button className="App__brands_btn" style={{marginLeft: 35}}>
                                <img src={MnyamsLogo}/>
                            </button>
                        </Link>

                        <Link to="/brit">
                            <button className="App__brands_btn">
                                <img src={BritLogo}/>
                            </button>
                        </Link>

                        <Link to="/cesar">
                            <button className="App__brands_btn">
                                <img src={CesarLogo}/>
                            </button>
                        </Link>

                        <Link to="/probalance">
                            <button className="App__brands_btn">
                                <img src={ProbalLogo}/>
                            </button>
                        </Link>

                        <Link to="/eightinone">
                            <button className="App__brands_btn">
                                <img src={EightinOneLogo}/>
                            </button>
                        </Link>

                        <Link to="/perfectfit">
                            <button className="App__brands_btn">
                                <img src={PerfectfitLogo}/>
                            </button>
                        </Link>

                        <Link to="/eukanuba">
                            <button className="App__brands_btn">
                                <img src={EukaLogo}/>
                            </button>
                        </Link>
                    </div>
                </div>

                <Footer/>
            </div>

        </div>
    );
};

export default MainPage;