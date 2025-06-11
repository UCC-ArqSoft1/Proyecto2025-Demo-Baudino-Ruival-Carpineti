import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import "./Home.css";

function Home() {
  const [activities, setActivities] = useState([]);
  const [search, setSearch] = useState("");
  const [filteredActivities, setFilteredActivities] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/activities")
      .then((res) => res.json())
      .then((data) => {
        setActivities(data);
        setFilteredActivities(data);
      })
      .catch((err) => console.error("Error al cargar actividades", err));
  }, []);

  const handleSearch = () => {
    const results = activities.filter((a) =>
      a.title.toLowerCase().includes(search.toLowerCase())
    );
    setFilteredActivities(results);
  };

  const handleReset = () => {
    setSearch("");
    setFilteredActivities(activities);
  };

  return (
    <div className="home-container">
      <h1 className="home-title">Actividades Disponibles</h1>

      <div className="search-bar">
        <input
          type="text"
          placeholder="Buscar por nombre:"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
        <button onClick={handleSearch}>Buscar</button>
        <button onClick={handleReset}>Limpiar</button>
      </div>

      <div className="activities-grid">
        {filteredActivities.length === 0 ? (
          <p className="no-activities">No se encontraron actividades.</p>
        ) : (
          filteredActivities.map((act) => (
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
