import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function ActivityDetails() {
  const { id } = useParams(); // Obtiene el ID desde la URL
  const [activity, setActivity] = useState(null);

  useEffect(() => {
    fetch(`http://localhost:8080/activities/${id}`)
      .then((res) => res.json())
      .then((data) => setActivity(data))
      .catch((err) => console.error("Error al cargar detalles", err));
  }, [id]);

  if (!activity) return <p>Cargando actividad...</p>;

  return (
    <div style={{ padding: "20px" }}>
      <h1>{activity.title}</h1>
      <p><strong>Categoría:</strong> {activity.category}</p>
      <p><strong>Instructor:</strong> {activity.instructor}</p>
      {/* Agregá más datos si tu backend devuelve más */}
    </div>
  );
}

export default ActivityDetails;
