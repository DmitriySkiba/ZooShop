import App from "../App";
import DogPage from "../pages/DogPage";
import MainPage from "../pages/MainPage";
import ErrorPage from "../pages/ErrorPage";
import SubcategoryPage from "../pages/SubcategoryPage";
import CatPage from "../pages/CatPage";
import ReptilePage from "../pages/ReptilePage";
import RodentPage from "../pages/RodentPage";
import FishPage from "../pages/FishPage";
import BirdPage from "../pages/BirdPage";
import BrandsPage from "../pages/BrandsPage";
import SearchPage from "../pages/SearchPage";
import ItemPage from "../pages/ItemPage";
import CartPage from "../pages/CartPage";
import BoughtHistory from "../pages/BoughtHistory";

export const routes = [
    {path: "/", element: <MainPage/>, exact: true},
    {path: "/error", element: <ErrorPage/>, exact: true},
    {path: "/search/:id", element: <SearchPage/>, exact: true},
    {path: "/item/:id", element: <ItemPage/>, exact: true},
    {path: "/cart", element: <CartPage/>, exact: true},
    {path: "/bought-history", element: <BoughtHistory/>, exact: true},

    {path: "/for-dog", element: <DogPage/>, exact: true},
    {path: "/for-dog/feeds", element: <SubcategoryPage subcategory={"Корма"} category={"Собаки"}/>, exact: true},
    {path: "/for-dog/threats", element: <SubcategoryPage subcategory={"Лакомства"} category={"Собаки"}/>, exact: true},
    {path: "/for-dog/vitamins", element: <SubcategoryPage subcategory={"Витамины"} category={"Собаки"}/>, exact: true},
    {path: "/for-dog/accessories", element: <SubcategoryPage subcategory={"Аксессуары"} category={"Собаки"}/>, exact: true},
    {path: "/for-dog/carries", element: <SubcategoryPage subcategory={"Переноски"} category={"Собаки"}/>, exact: true},

    {path: "/for-cat", element: <CatPage/>, exact: true},
    {path: "/for-cat/feeds", element:  <SubcategoryPage subcategory={"Корма"} category={"Кошки"}/>, exact: true},
    {path: "/for-cat/threats", element: <SubcategoryPage subcategory={"Лакомства"} category={"Кошки"}/>, exact: true},
    {path: "/for-cat/vitamins", element: <SubcategoryPage subcategory={"Витамины"} category={"Кошки"}/>, exact: true},
    {path: "/for-cat/accessories", element: <SubcategoryPage subcategory={"Аксессуары"} category={"Кошки"}/>, exact: true},
    {path: "/for-cat/carries", element: <SubcategoryPage subcategory={"Переноски"} category={"Кошки"}/>, exact: true},

    {path: "/for-fish", element: <FishPage/>, exact: true},
    {path: "/for-fish/feeds", element: <SubcategoryPage subcategory={"Корма"} category={"Рыбы"}/>, exact: true},
    {path: "/for-fish/vitamins", element: <SubcategoryPage subcategory={"Витамины"} category={"Рыбы"}/>, exact: true},
    {path: "/for-fish/aquariums", element: <SubcategoryPage subcategory={"Аквариумы"} category={"Рыбы"}/>, exact: true},
    {path: "/for-fish/decorations", element: <SubcategoryPage subcategory={"Декорации"} category={"Рыбы"}/>, exact: true},
    {path: "/for-fish/accessories", element: <SubcategoryPage subcategory={"Аксессуары"} category={"Рыбы"}/>, exact: true},

    {path: "/for-bird", element: <BirdPage/>, exact: true},
    {path: "/for-bird/feeds", element: <SubcategoryPage subcategory={"Корма"} category={"Птицы"}/>, exact: true},
    {path: "/for-bird/threats", element: <SubcategoryPage subcategory={"Лакомства"} category={"Птицы"}/>, exact: true},
    {path: "/for-bird/vitamins", element: <SubcategoryPage subcategory={"Витамины"} category={"Птицы"}/>, exact: true},
    {path: "/for-bird/accessories", element: <SubcategoryPage subcategory={"Аксессуары"} category={"Птицы"}/>, exact: true},
    {path: "/for-bird/cages", element: <SubcategoryPage subcategory={"Клетки"} category={"Птицы"}/>, exact: true},

    {path: "/for-rodent", element: <RodentPage/>, exact: true},
    {path: "/for-rodent/feeds", element: <SubcategoryPage subcategory={"Корма"} category={"Грызуны"}/>, exact: true},
    {path: "/for-rodent/threats", element: <SubcategoryPage subcategory={"Лакомства"} category={"Грызуны"}/>, exact: true},
    {path: "/for-rodent/vitamins", element: <SubcategoryPage subcategory={"Витамины"} category={"Грызуны"}/>, exact: true},
    {path: "/for-rodent/accessories", element: <SubcategoryPage subcategory={"Аксессуары"} category={"Грызуны"}/>, exact: true},
    {path: "/for-rodent/cages", element: <SubcategoryPage subcategory={"Клетки"} category={"Грызуны"}/>, exact: true},

    {path: "/for-reptile", element: <ReptilePage/>, exact: true},
    {path: "/for-reptile/feeds", element: <SubcategoryPage subcategory={"Корма"} category={"Рептилии"}/>, exact: true},
    {path: "/for-reptile/vitamins", element: <SubcategoryPage subcategory={"Витамины"} category={"Рептилии"}/>, exact: true},
    {path: "/for-reptile/decoration", element: <SubcategoryPage subcategory={"Декорации"} category={"Рептилии"}/>, exact: true},
    {path: "/for-reptile/terrariums", element: <SubcategoryPage subcategory={"Террариумы"} category={"Рептилии"}/>, exact: true},


    {path: "/brit", element: <BrandsPage brand={"Brit"}/>, exact: true},
    {path: "/cesar", element: <BrandsPage brand={"Cesar"}/>, exact: true},
    {path: "/dreamies", element: <BrandsPage brand={"Dreamies"}/>, exact: true},
    {path: "/eightinone", element: <BrandsPage brand={"EightInOne"}/>, exact: true},
    {path: "/eukanuba", element: <BrandsPage brand={"Eukanuba"}/>, exact: true},
    {path: "/kitekat", element: <BrandsPage brand={"Kitekat"}/>, exact: true},
    {path: "/mnyams", element: <BrandsPage brand={"Mnyams"}/>, exact: true},
    {path: "/pedigree", element: <BrandsPage brand={"Pedigree"}/>, exact: true},
    {path: "/perfectfit", element: <BrandsPage brand={"PerfectFit"}/>, exact: true},
    {path: "/probalance", element: <BrandsPage brand={"Probalance"}/>, exact: true},
    {path: "/purina", element: <BrandsPage brand={"Purina"}/>, exact: true},
    {path: "/royalcanin", element: <BrandsPage brand={"Royal Canin"}/>, exact: true},
    {path: "/sheba", element: <BrandsPage brand={"Sheba"}/>, exact: true},
    {path: "/whiskas", element: <BrandsPage brand={"Whiskas"}/>, exact: true}

]