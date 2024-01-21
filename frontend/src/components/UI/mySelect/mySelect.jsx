import React from 'react';
import "./mySelect.css"

const mySelect = ({options, defaultValue, value, onChange}) => {
    return (
        <select
            className={"Myselect"}
            value={value}
            onChange={event => onChange(event.target.value)}
        >
            <option disabled value="">{defaultValue}</option>
            {options.map(option =>
                <option key={option.value} value={option.value}>
                    {option.name}
                </option>
            )}
        </select>
    );
};

export default mySelect;