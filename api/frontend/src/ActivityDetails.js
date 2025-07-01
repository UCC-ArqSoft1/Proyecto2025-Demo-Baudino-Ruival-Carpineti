"use client"

import { useEffect, useState } from "react"
import { useParams, useNavigate } from "react-router-dom"
import Modal from "./Modal"
import { getCookie, decodeJWT, getUserRole } from "./utils/cookies"
import "./ActivityDetails.css"

function ActivityDetails() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [activity, setActivity] = useState(null)
  const [modal, setModal] = useState({ show: false, message: "", success: true })
  const [userId, setUserId] = useState(null)

  useEffect(() => {
    // Obtener userID del token
    const token = getCookie("token")
    if (token) {
      const payload = decodeJWT(token)
      if (payload && payload.jti) {
        setUserId(Number.parseInt(payload.jti))
      }
    }

    fetch(`http://localhost:8080/activities/${id}`)
      .then((res) => res.json())
      .then((data) => setActivity(data))
      .catch((err) => console.error("Error al cargar detalles", err))
  }, [id])

  const handleEnroll = (scheduleId) => {
    if (!userId) {
      setModal({ show: true, message: "Debes iniciar sesión para inscribirte", success: false })
      return
    }

    const token = getCookie("token")

    fetch(`http://localhost:8080/users/${userId}/enrollments`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify({ schedule_id: scheduleId }),
    })
      .then((res) => {
        if (!res.ok)
          return res.json().then((err) => {
            throw new Error(err.error || "Inscripción fallida")
          })
        return res.json()
      })
      .then((data) => {
        setModal({ show: true, message: data.message, success: true })
        return fetch(`http://localhost:8080/activities/${id}`)
      })
      .then((res) => res.json())
      .then((updated) => setActivity(updated))
      .catch((err) => setModal({ show: true, message: err.message, success: false }))
  }

  const isAdmin = getUserRole() === "admin"

  const handleDelete = async () => {
    const token = getCookie("token")
    if (!window.confirm("¿Seguro que deseas eliminar esta actividad?")) return
    const res = await fetch(`http://localhost:8080/admin/activities/${activity.id}`, {
      method: "DELETE",
      headers: { Authorization: token },
    })
    if (res.ok) {
      window.location.href = "/activities"
    } else {
      alert("Error al eliminar actividad")
    }
  }

  if (!activity) {
    return (
      <div className="activity-details-container">
        <div className="activity-details-content">
          <div className="loading-message">
            <p>⏳ Cargando actividad...</p>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="activity-details-container">
      <div className="activity-details-content">
        <div className="activity-hero">
          <h1>🏆 {activity.title}</h1>
          {activity.image && (
            <div className="activity-image-container">
              <img
                src={activity.image || "/placeholder.svg"}
                alt={activity.title}
                className="activity-image"
                onError={(e) => {
                  e.target.style.display = "none"
                }}
              />
            </div>
          )}
          <p>
            <strong>📂 Categoría:</strong> {activity.category}
          </p>
          <p>
            <strong>👨‍🏫 Instructor:</strong> {activity.instructor}
          </p>
          <p>
            <strong>📝 Descripción:</strong> {activity.description}
          </p>
          <p>
            <strong>⏱️ Duración:</strong> {activity.duration} minutos
          </p>
          <p>
            <strong>📊 Estado:</strong> {activity.status}
          </p>
          <div className="back-button-container">
            <button className="back-btn" onClick={() => navigate("/activities")}>
              ← Volver a Actividades
            </button>
          </div>
        </div>

        <div className="schedules-section">
          <h2 className="schedules-title">📅 Horarios Disponibles</h2>
          <div className="schedules-grid">
            {activity.schedules.map((s) => (
              <div key={s.id} className="schedule-card">
                <p>
                  <strong>📆 Día:</strong> {s.week_day}
                </p>
                <p>
                  <strong>🕐 Inicio:</strong> {s.start_time}
                </p>
                <p>
                  <strong>🕕 Fin:</strong> {s.end_time}
                </p>
                <p>
                  <strong>👥 Cupo:</strong> {s.capacity} personas
                </p>
                <button className="enroll-btn" onClick={() => handleEnroll(s.id)}>
                  ✅ Inscribirme
                </button>
              </div>
            ))}
          </div>
        </div>

        {isAdmin && (
          <div className="admin-actions">
            <button className="edit-btn" onClick={() => (window.location.href = `/admin/edit-activity/${activity.id}`)}>
              ✏️ Editar
            </button>
            <button className="delete-btn" onClick={handleDelete}>
              🗑️ Eliminar
            </button>
          </div>
        )}

        <Modal
          show={modal.show}
          message={modal.message}
          success={modal.success}
          onClose={() => setModal({ ...modal, show: false })}
        />
      </div>
    </div>
  )
}

export default ActivityDetails
