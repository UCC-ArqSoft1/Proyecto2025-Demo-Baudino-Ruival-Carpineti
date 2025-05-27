import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import "./Home.css";

function Home() {
  const [activities, setActivities] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/activities")
      .then((res) => res.json())
      .then((data) => setActivities(data))
      .catch((err) => console.error("Error al cargar actividades", err));
  }, []);

  return (
    <div className="home-container">
      <h1 className="home-title">Actividades Disponibles</h1>
      <div className="activities-grid">
        {activities.length === 0 ? (
          <p className="no-activities">No hay actividades disponibles.</p>
        ) : (
          activities.map((act) => (
            <Link to={`/activities/${act.id}`} key={act.id} className="activity-card">
              <h2>{act.title}</h2>
              <p><strong>Categoría:</strong> {act.category}</p>
              <p><strong>Instructor:</strong> {act.instructor}</p>
            </Link>
          ))
        )}
      </div>
    </div>
  );
}

export default Home;