"use client"

import { useEffect, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { getCookie, decodeJWT } from "./utils/cookies"
import "./MyActivities.css"

function MyActivities() {
  const [myActivities, setMyActivities] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState("")
  const navigate = useNavigate()

  useEffect(() => {
    const fetchMyActivities = async () => {
      try {
        // Obtener userID del token
        const token = getCookie("token")
        if (!token) {
          navigate("/")
          return
        }

        const payload = decodeJWT(token)
        if (!payload || !payload.jti) {
          navigate("/")
          return
        }

        const userId = payload.jti

        // Hacer petición al backend
        const response = await fetch(`http://localhost:8080/users/${userId}/activities`, {
          method: "GET",
          headers: {
            Authorization: token,
            "Content-Type": "application/json",
          },
        })

        if (!response.ok) {
          throw new Error("Error al cargar mis actividades")
        }

        const data = await response.json()
        setMyActivities(data || [])
      } catch (err) {
        console.error("Error fetching my activities:", err)
        setError("Error al cargar tus actividades")
      } finally {
        setLoading(false)
      }
    }

    fetchMyActivities()
  }, [navigate])

  const handleLogout = () => {
    document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;"
    navigate("/")
  }

  if (loading) {
    return (
      <div className="my-activities-container">
        <div className="my-activities-content">
          <div className="loading-message">
            <p>⏳ Cargando tus actividades...</p>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="my-activities-container">
      <div className="my-activities-content">
        <div className="my-activities-header">
          <h1 className="my-activities-title">📋 Mis Actividades</h1>
          <p className="my-activities-subtitle">Actividades en las que estás inscrito</p>

          <div className="navigation-controls">
            <button className="back-to-activities-btn" onClick={() => navigate("/activities")}>
              ← Volver a Todas las Actividades
            </button>
          </div>
        </div>

        {error && (
          <div className="error-section">
            <p>❌ {error}</p>
          </div>
        )}

        <div className="my-activities-grid">
          {myActivities.length === 0 ? (
            <div className="no-activities">
              <p>📝 No estás inscrito en ninguna actividad aún</p>
              <button className="browse-activities-btn" onClick={() => navigate("/activities")}>
                🔍 Explorar Actividades
              </button>
            </div>
          ) : (
            myActivities.map((activity) => (
              <div key={activity.id} className="my-activity-card">
                {activity.image && (
                  <div className="my-activity-card-image">
                    <img
                      src={activity.image || "/placeholder.svg"}
                      alt={activity.title}
                      onError={(e) => {
                        e.target.style.display = "none"
                      }}
                    />
                  </div>
                )}
                <div className="my-activity-card-content">
                  <div className="enrolled-badge">✅ Inscrito</div>
                  <h2>🏆 {activity.title}</h2>
                  <p>
                    <strong>📂 Categoría:</strong> {activity.category}
                  </p>
                  <p>
                    <strong>👨‍🏫 Instructor:</strong> {activity.instructor}
                  </p>
                  <p>
                    <strong>📝 Descripción:</strong> {activity.description.substring(0, 100)}...
                  </p>

                  <div className="my-schedules-section">
                    <h4>📅 Mis Horarios:</h4>
                    <div className="my-schedules-list">
                      {activity.schedules.map((schedule) => (
                        <div key={schedule.id} className="my-schedule-item">
                          <span className="schedule-day">📆 {schedule.week_day}</span>
                          <span className="schedule-time">
                            🕐 {schedule.start_time} - {schedule.end_time}
                          </span>
                        </div>
                      ))}
                    </div>
                  </div>

                  <Link to={`/activities/${activity.id}`} className="view-details-btn">
                    👁️ Ver Detalles
                  </Link>
                </div>
              </div>
            ))
          )}
        </div>

        <div className="logout-section">
          <button className="logout-btn" onClick={handleLogout}>
            🚪 Cerrar Sesión
          </button>
        </div>
      </div>
    </div>
  )
}

export default MyActivities
