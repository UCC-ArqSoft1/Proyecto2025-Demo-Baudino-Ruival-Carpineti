import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./Home.css";

function Home() {
  const [activities, setActivities] = useState([]);
  const [search, setSearch] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    fetch("http://localhost:8080/activities")
      .then((res) => res.json())
      .then((data) => setActivities(data))
      .catch((err) => console.error("Error al cargar actividades", err));
  }, []);

  const handleSearch = () => {
    fetch(`http://localhost:8080/activities/search?keyword=${search}`)
      .then((res) => res.json())
      .then((data) => setActivities(data))
      .catch((err) => console.error("Error al buscar", err));
  };

  const handleClear = () => {
    setSearch("");
    fetch("http://localhost:8080/activities")
      .then((res) => res.json())
      .then((data) => setActivities(data))
      .catch((err) => console.error("Error al cargar actividades", err));
  };

  const handleLogout = () => {
    document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    navigate("/");
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
        <button onClick={handleClear}>Limpiar</button>
      </div>

      <div className="activities-grid">
        {activities.length === 0 ? (
          <p className="no-activities">No hay actividades disponibles.</p>
        ) : (
          activities.map((act) => (
            <Link
              to={`/activities/${act.id}`}
              key={act.id}
              className={`activity-card ${act.category.toLowerCase()}`}
            >
              <h2>{act.title}</h2>
              <p>
                <strong>Categoría:</strong> {act.category}
              </p>
              <p>
                <strong>Instructor:</strong> {act.instructor}
              </p>
            </Link>
          ))
        )}
      </div>
      <button className="logout-btn" onClick={handleLogout}>Cerrar sesión</button>
    </div>
  );
}

export default Home;
