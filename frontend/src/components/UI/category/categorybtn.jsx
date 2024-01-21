import React, { useState, useEffect, useRef } from 'react';
import Mybutton from "../button/Mybutton";
import {CategoryItemsData} from "./CategoryItemsData";
import CategoryItems from "./CategoryItems";
import "./Category.css"


const depthLevel = 0;

const Categorybtn = () => {
    const [dropdown, setDropdown] = useState(false);

    let ref = useRef();

    useEffect(() => {
        const handler = (event) => {
            if (dropdown && ref.current && !ref.current.contains(event.target)) {
                setDropdown(false);
            }
        };
        document.addEventListener("mousedown", handler);
        document.addEventListener("touchstart", handler);
        return () => {
            document.removeEventListener("mousedown", handler);
            document.removeEventListener("touchstart", handler);
        };
    }, [dropdown]);

    return (
        <div>
            <Mybutton style={{height: 50, width: 125}}
                      aria-haspopup="category"
                      aria-expanded={dropdown ? "true" : "false"}
                      onClick={() => setDropdown((prev) => !prev)}>
                Категории
            </Mybutton>
            <ul className={`menus ${dropdown ? "show" : ""}`} ref={ref} style={{zIndex:2}}>
                {CategoryItemsData.map((menu, index) => {
                    return <CategoryItems items={menu} key={index} depthLevel={depthLevel} style={{zIndex:3}}/>;
                })}
            </ul>
        </div>
    );
};

export default Categorybtn;