import { useEffect, useState } from "react";

function Hotel(){

    const [hotel, setHotel] = useState({});

    useEffect(() => {
        fetch("http://localhost:8080/hotels/1")
        .then((res)=> {
            return res.json();
        })
        .then((data) => {
            setHotel(data);
        })
    });


    return(
        <div>
            <h1>Holiday Inn</h1>
            <div>
                <p>location: Av Colon 500</p>
                <p>rating: 5.00</p>
                <p>Descripcion del Hotel</p>
                <img src="https://es.holidayinnrosario.com/web/uploads/sliders/1/sliders/0-desktop.jpg?1584710871"/>
            </div>
            
        </div>
    )
}

export default Hotel;