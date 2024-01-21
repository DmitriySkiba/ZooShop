import React from 'react';
import {useParams} from "react-router-dom";
import Header from "../components/UI/header/header";
import Footer from "../components/UI/footer/Footer";
import ItemsForSearchPage from "../components/items/ItemsForSearchPage";
const SearchPage = () => {
    const {id} = useParams();
    return (
        <div className="App">
            <Header></Header>
            <ItemsForSearchPage searchQuery={id}></ItemsForSearchPage>
            <Footer></Footer>
        </div>
    );
};

export default SearchPage;