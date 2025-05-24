import React, { useEffect, useState } from "react";

function Home() {
  const [activities, setActivities] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/activities")
      .then((res) => res.json())
      .then((data) => setActivities(data))
      .catch((err) => console.error("Error al cargar actividades", err));
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h1>Actividades Disponibles</h1>
      {activities.length === 0 ? (
        <p>No hay actividades disponibles.</p>
      ) : (
        <ul>
          {activities.map((act) => (
            <li key={act.id}>
              <strong>{act.title}</strong> - {act.category}
              <br />
              Instructor: {act.instructor}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default Home;