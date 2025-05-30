import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function ActivityDetails() {
  const { id } = useParams();
  const [activity, setActivity] = useState(null);
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch(`http://localhost:8080/activities/${id}`)
      .then((res) => res.json())
      .then((data) => setActivity(data))
      .catch((err) => console.error("Error al cargar detalles", err));
  }, [id]);

  const handleEnroll = (scheduleId) => {
    const userId = 1; // ✳️ Cambia esto según cómo manejes el login real
    fetch(`http://localhost:8080/users/${userId}/enrollments`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ schedule_id: scheduleId }),
    })
      .then((res) => {
        if (!res.ok) throw new Error("Inscripción fallida");
        return res.json();
      })
      .then((data) => setMessage(data.message))
      .catch((err) => setMessage(err.message));
  };

  if (!activity) return <p>Cargando actividad...</p>;

  return (
    <div style={{ padding: "30px", color: "white" }}>
      <h1>{activity.title}</h1>
      <p><strong>Categoría:</strong> {activity.category}</p>
      <p><strong>Instructor:</strong> {activity.instructor}</p>
      <p><strong>Descripción:</strong> {activity.description}</p>

      <h2 style={{ marginTop: "30px" }}>Horarios Disponibles</h2>
      {activity.schedules.map((s) => (
        <div key={s.id} style={{
          border: "1px solid #00c6ff",
          borderRadius: "10px",
          padding: "15px",
          marginBottom: "15px",
          backgroundColor: "rgba(255,255,255,0.05)",
        }}>
          <p><strong>Día:</strong> {s.week_day}</p>
          <p><strong>Inicio:</strong> {s.start_time}</p>
          <p><strong>Fin:</strong> {s.end_time}</p>
          <p><strong>Cupo:</strong> {s.capacity}</p>
          <button onClick={() => handleEnroll(s.id)} style={{
            padding: "10px 20px",
            backgroundColor: "#00c6ff",
            border: "none",
            borderRadius: "5px",
            color: "#fff",
            cursor: "pointer"
          }}>
            Inscribirme
          </button>
        </div>
      ))}
      {message && <p style={{ color: "lime", marginTop: "20px" }}>{message}</p>}
    </div>
  );
}

export default ActivityDetails;
