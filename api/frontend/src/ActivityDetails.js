import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Modal from "./Modal";

function ActivityDetails() {
  const { id } = useParams();
  const [activity, setActivity] = useState(null);
  const [modal, setModal] = useState({ show: false, message: "", success: true });

  useEffect(() => {
    fetch(`http://localhost:8080/activities/${id}`)
      .then((res) => res.json())
      .then((data) => setActivity(data))
      .catch((err) => console.error("Error al cargar detalles", err));
  }, [id]);

  const handleEnroll = (scheduleId) => {
    const userId = 1; // 🔒 Usar valor real cuando implementes login completo
    fetch(`http://localhost:8080/users/${userId}/enrollments`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ schedule_id: scheduleId }),
    })
      .then(async (res) => {
        const data = await res.json();
        if (!res.ok) {
          if (data.error?.includes("already enrolled")) {
            throw new Error("Ya estás inscripto en este horario");
          } else if (data.error?.includes("no available capacity")) {
            throw new Error("No hay cupo disponible");
          } else {
            throw new Error(data.error || "Inscripción fallida");
          }
        }
        return data;
      })
      .then((data) => {
        setModal({ show: true, message: data.message, success: true });
        return fetch(`http://localhost:8080/activities/${id}`);
      })
      .then((res) => res.json())
      .then((updated) => setActivity(updated))
      .catch((err) =>
        setModal({ show: true, message: err.message, success: false })
      );
  };

  if (!activity) return <p style={{ color: "white" }}>Cargando actividad...</p>;

  return (
    <div style={{ padding: "30px", color: "white" }}>
      <div style={{
        border: "1px solid #00c6ff",
        borderRadius: "10px",
        padding: "20px",
        marginBottom: "30px",
        backgroundColor: "rgba(255,255,255,0.05)"
      }}>
        <h1>{activity.title}</h1>
        <p><strong>Categoría:</strong> {activity.category}</p>
        <p><strong>Instructor:</strong> {activity.instructor}</p>
        <p><strong>Descripción:</strong> {activity.description}</p>
      </div>

      <h2 style={{ marginBottom: "20px" }}>Horarios Disponibles</h2>
      <div style={{ display: "flex", flexWrap: "wrap", gap: "20px" }}>
        {activity.schedules.map((s) => (
          <div key={s.id} style={{
            border: "1px solid #00c6ff",
            borderRadius: "10px",
            padding: "15px",
            backgroundColor: "rgba(255,255,255,0.05)",
            minWidth: "250px",
            flex: "1"
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
              cursor: "pointer",
              marginTop: "10px",
              width: "100%"
            }}>
              Inscribirme
            </button>
          </div>
        ))}
      </div>

      <Modal
        show={modal.show}
        message={modal.message}
        success={modal.success}
        onClose={() => setModal({ ...modal, show: false })}
      />
    </div>
  );
}

export default ActivityDetails;
