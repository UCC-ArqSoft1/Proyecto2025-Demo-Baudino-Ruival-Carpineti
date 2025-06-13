import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Modal from "./Modal";
import { getCookie, decodeJWT } from "./utils/cookies"; 

function ActivityDetails() {
  const { id } = useParams();
  const [activity, setActivity] = useState(null);
  const [modal, setModal] = useState({ show: false, message: "", success: true });
  const [userId, setUserId] = useState(null); 

  useEffect(() => {
    // Obtener userID del token
    const token = getCookie('token');
    if (token) {
      const payload = decodeJWT(token);
      if (payload && payload.jti) {
        setUserId(parseInt(payload.jti));
      }
    }

    
    fetch(`http://localhost:8080/activities/${id}`)
      .then((res) => res.json())
      .then((data) => setActivity(data))
      .catch((err) => console.error("Error al cargar detalles", err));
  }, [id]);

  const handleEnroll = (scheduleId) => {
    if (!userId) { // Validar sesión
      setModal({ show: true, message: "Debes iniciar sesión para inscribirte", success: false });
      return;
    }

    // Usar userId dinámico
    fetch(`http://localhost:8080/users/${userId}/enrollments`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ schedule_id: scheduleId }),
    })
      .then((res) => {
        if (!res.ok) return res.json().then(err => { throw new Error(err.error || "Inscripción fallida") });
        return res.json();
      })
      .then((data) => {
        setModal({ show: true, message: data.message, success: true });
        return fetch(`http://localhost:8080/activities/${id}`);
      })
      .then((res) => res.json())
      .then((updated) => setActivity(updated))
      .catch((err) => setModal({ show: true, message: err.message, success: false }));
  };

  const getBackground = () => {
    if (!activity) return "none";
    if (activity.title.toLowerCase().includes("yoga")) return "url('/yoga-bg(2).jpg.png')";
    if (activity.title.toLowerCase().includes("spinning")) return "url('/spinning-bg(1).jpg.png')";
    return "none";
  };

  if (!activity) return <p style={{ color: "white", padding: "20px" }}>Cargando actividad...</p>;

  return (
    <div style={{ padding: "30px", color: "white" }}>
      <div style={{
        padding: "30px",
        borderRadius: "10px",
        border: "1px solid #00c6ff",
        backgroundImage: getBackground(),
        backgroundSize: "cover",
        backgroundPosition: "center",
        backgroundRepeat: "no-repeat",
        backgroundBlendMode: "overlay",
        backgroundColor: "rgba(0,0,0,0.6)",
        marginBottom: "30px"
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
            flex: "1 1 200px",
            border: "1px solid #00c6ff",
            borderRadius: "10px",
            padding: "15px",
            backgroundColor: "rgba(255,255,255,0.05)"
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
