"use client"

import { useEffect, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import "./Home.css"
import { getUserRole, getCookie } from "./utils/cookies"

function Home() {
  const [activities, setActivities] = useState([])
  const [search, setSearch] = useState("")
  const navigate = useNavigate()

  const isAdmin = getUserRole() === "admin"
  const token = getCookie("token")

  useEffect(() => {
    if (isAdmin) {
      fetch("http://localhost:8080/admin/activities", {
        headers: { Authorization: token }
      })
        .then((res) => res.json())
        .then((data) => setActivities(data))
        .catch((err) => console.error("Error al cargar actividades", err))
    } else {
      fetch("http://localhost:8080/activities")
        .then((res) => res.json())
        .then((data) => setActivities(data))
        .catch((err) => console.error("Error al cargar actividades", err))
    }
  }, [isAdmin, token])

  const handleSearch = () => {
    if (isAdmin) {
      // (Opcional: podrías implementar búsqueda para admins en el backend)
      fetch("http://localhost:8080/admin/activities", {
        headers: { Authorization: token }
      })
        .then((res) => res.json())
        .then((data) => setActivities(data))
        .catch((err) => console.error("Error al buscar", err))
    } else {
      fetch(`http://localhost:8080/activities/search?keyword=${search}`)
        .then((res) => res.json())
        .then((data) => setActivities(data))
        .catch((err) => console.error("Error al buscar", err))
    }
  }

  const handleClear = () => {
    setSearch("")
    if (isAdmin) {
      fetch("http://localhost:8080/admin/activities", {
        headers: { Authorization: token }
      })
        .then((res) => res.json())
        .then((data) => setActivities(data))
        .catch((err) => console.error("Error al cargar actividades", err))
    } else {
      fetch("http://localhost:8080/activities")
        .then((res) => res.json())
        .then((data) => setActivities(data))
        .catch((err) => console.error("Error al cargar actividades", err))
    }
  }

  const handleLogout = () => {
    document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;"
    navigate("/")
  }

  return (
    <div className="home-container">
      <div className="home-content">
        <div className="home-header">
          <h1 className="home-title">🏃‍♂️ Actividades Deportivas</h1>
          <p className="home-subtitle">Encuentra la actividad perfecta para ti</p>

          <div className="navigation-buttons">
            <button className="my-activities-btn" onClick={() => navigate("/my-activities")}>
              📋 Mis Actividades
            </button>
            {isAdmin && (
              <button className="admin-btn" onClick={() => navigate("/admin/create-activity")}>
                ➕ Crear Actividad
              </button>
            )}
          </div>
        </div>

        <div className="search-section">
          <div className="search-bar">
            <input
              type="text"
              placeholder="🔍 Buscar actividades por nombre..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              onKeyPress={(e) => e.key === "Enter" && handleSearch()}
            />
            <button className="search-btn" onClick={handleSearch}>
              Buscar
            </button>
            <button className="clear-btn" onClick={handleClear}>
              Limpiar
            </button>
          </div>
        </div>

        <div className="activities-grid">
          {activities.length === 0 ? (
            <div className="no-activities">
              <p>🚫 No hay actividades disponibles</p>
            </div>
          ) : (
            activities.map((act) => (
              <Link to={`/activities/${act.id}`} key={act.id} className="activity-card">
                {act.image && (
                  <div className="activity-card-image">
                    <img
                      src={act.image || "/placeholder.svg"}
                      alt={act.title}
                      onError={(e) => {
                        e.target.style.display = "none"
                      }}
                    />
                  </div>
                )}
                <div className="activity-card-content">
                  <h2>🏆 {act.title}</h2>
                  <p>
                    <strong>📂 Categoría:</strong> {act.category}
                  </p>
                  <p>
                    <strong>👨‍🏫 Instructor:</strong> {act.instructor}
                  </p>
                  <p>
                    <strong>📝 Descripción:</strong> {act.description.substring(0, 100)}...
                  </p>
                </div>
              </Link>
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

export default Home
