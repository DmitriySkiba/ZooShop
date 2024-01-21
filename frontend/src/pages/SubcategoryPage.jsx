import React from 'react';
import Header from "../components/UI/header/header";
import Footer from "../components/UI/footer/Footer";
import ItemsForSubcategoryPage from "../components/items/ItemsForSubcategoryPage";

const SubcategoryPage = (props) => {
    return (
        <div className="App">
                <Header></Header>
                <ItemsForSubcategoryPage subcategory={props.subcategory} type={"category"} category={props.category}></ItemsForSubcategoryPage>
                <Footer></Footer>
        </div>
    );
};

export default SubcategoryPage;