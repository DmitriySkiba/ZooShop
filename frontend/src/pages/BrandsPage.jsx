import React from 'react';
import Header from "../components/UI/header/header";
import Footer from "../components/UI/footer/Footer";
import ItemsForBrandsPage from "../components/items/ItemsForBrandsPage";

const BrandsPage = (props) => {
    return (
        <div className="App">
            <Header></Header>
            <ItemsForBrandsPage type={"brand"} category={props.brand} brand={props.brand}></ItemsForBrandsPage>
            <Footer></Footer>
        </div>

    );
};

export default BrandsPage;