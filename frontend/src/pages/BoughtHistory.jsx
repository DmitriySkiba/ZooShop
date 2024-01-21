import React from 'react';
import Header from "../components/UI/header/header";
import Footer from "../components/UI/footer/Footer";
import ItemsForBoughtHistory from "../components/items/ItemsForBoughtHistory";

const BoughtHistory = () => {
    return (
        <div className="App">
            <Header></Header>
            <ItemsForBoughtHistory/>
            <Footer></Footer>
        </div>
    );
};

export default BoughtHistory;