"use client"

import { useState, useEffect } from "react"
import { useNavigate, useParams } from "react-router-dom"
import { getCookie } from "./utils/cookies"
import "./ActivityForm.css"

const emptySchedule = { week_day: "", start_time: "", end_time: "", capacity: 0 }

function ActivityForm({ mode }) {
  const [form, setForm] = useState({
    title: "",
    description: "",
    category: "",
    instructor: "",
    duration: 60,
    image: "",
    status: "activo",
    schedules: [{ ...emptySchedule }],
  })
  const [error, setError] = useState("")
  const navigate = useNavigate()
  const { id } = useParams()

  useEffect(() => {
    if (mode === "edit" && id) {
      fetch(`http://localhost:8080/activities/${id}`)
        .then((res) => res.json())
        .then((data) => {
          setForm({
            title: data.title,
            description: data.description,
            category: data.category,
            instructor: data.instructor,
            duration: data.duration,
            image: data.image,
            status: data.status,
            schedules: data.schedules.map((s) => ({
              week_day: s.week_day,
              start_time: s.start_time,
              end_time: s.end_time,
              capacity: s.capacity,
            })),
          })
        })
    }
  }, [mode, id])

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleScheduleChange = (idx, e) => {
    const newSchedules = form.schedules.map((s, i) => (i === idx ? { ...s, [e.target.name]: e.target.value } : s))
    setForm({ ...form, schedules: newSchedules })
  }

  const addSchedule = () => {
    setForm({ ...form, schedules: [...form.schedules, { ...emptySchedule }] })
  }

  const removeSchedule = (idx) => {
    setForm({ ...form, schedules: form.schedules.filter((_, i) => i !== idx) })
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError("")
    const token = getCookie("token")
    const url =
      mode === "edit" ? `http://localhost:8080/admin/activities/${id}` : "http://localhost:8080/admin/activities"
    const method = mode === "edit" ? "PUT" : "POST"
    const mappedForm = {
      ...form,
      duration: Number(form.duration),
      schedules: form.schedules.map((s) => ({
        week_day: s.week_day,
        start_time: s.start_time,
        end_time: s.end_time,
        capacity: Number(s.capacity),
      })),
    }
    const res = await fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify(mappedForm),
    })
    if (res.ok) {
      navigate("/activities")
    } else {
      let errorMsg = "Error al guardar la actividad"
      try {
        const data = await res.json()
        if (data && data.error) errorMsg = data.error
      } catch {}
      setError(errorMsg)
    }
  }

  return (
    <div className="form-container">
      <div className="form-content">
        <div className="form-header">
          <h2>{mode === "edit" ? "✏️ Editar Actividad" : "➕ Crear Nueva Actividad"}</h2>
        </div>

        <div className="form-main">
          {error && <div className="form-error">❌ {error}</div>}

          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <label>🏆 Título de la Actividad</label>
              <input name="title" value={form.title} onChange={handleChange} placeholder="Ej: Yoga Matutino" required />
            </div>

            <div className="form-group">
              <label>📝 Descripción</label>
              <textarea
                name="description"
                value={form.description}
                onChange={handleChange}
                placeholder="Describe la actividad deportiva..."
                required
              />
            </div>

            <div className="form-group">
              <label>📂 Categoría</label>
              <input
                name="category"
                value={form.category}
                onChange={handleChange}
                placeholder="Ej: Yoga, Spinning, Pilates"
                required
              />
            </div>

            <div className="form-group">
              <label>👨‍🏫 Instructor</label>
              <input
                name="instructor"
                value={form.instructor}
                onChange={handleChange}
                placeholder="Nombre del instructor"
                required
              />
            </div>

            <div className="form-group">
              <label>⏱️ Duración (minutos)</label>
              <input
                name="duration"
                type="number"
                value={form.duration}
                onChange={handleChange}
                placeholder="60"
                required
                min="1"
              />
            </div>

            <div className="form-group">
              <label>🖼️ Imagen (URL)</label>
              <input
                name="image"
                value={form.image}
                onChange={handleChange}
                placeholder="https://ejemplo.com/imagen.jpg"
              />
            </div>

            <div className="form-group">
              <label>📊 Estado</label>
              <select name="status" value={form.status} onChange={handleChange} required>
                <option value="activo">✅ Activo</option>
                <option value="inactivo">❌ Inactivo</option>
              </select>
            </div>

            <div className="schedules-section">
              <h3 className="schedules-title">📅 Horarios de la Actividad</h3>

              {form.schedules.map((s, idx) => (
                <div key={idx} className="schedule-item">
                  <h4>Horario #{idx + 1}</h4>

                  {form.schedules.length > 1 && (
                    <button type="button" className="remove-schedule-btn" onClick={() => removeSchedule(idx)}>
                      🗑️ Eliminar
                    </button>
                  )}

                  <div className="schedule-grid">
                    <div className="form-group">
                      <label>📆 Día de la semana</label>
                      <input
                        name="week_day"
                        value={s.week_day}
                        onChange={(e) => handleScheduleChange(idx, e)}
                        placeholder="Lunes"
                        required
                      />
                    </div>

                    <div className="form-group">
                      <label>🕐 Hora de inicio</label>
                      <input
                        name="start_time"
                        value={s.start_time}
                        onChange={(e) => handleScheduleChange(idx, e)}
                        placeholder="09:00"
                        required
                      />
                    </div>

                    <div className="form-group">
                      <label>🕕 Hora de fin</label>
                      <input
                        name="end_time"
                        value={s.end_time}
                        onChange={(e) => handleScheduleChange(idx, e)}
                        placeholder="10:00"
                        required
                      />
                    </div>

                    <div className="form-group">
                      <label>👥 Capacidad</label>
                      <input
                        name="capacity"
                        type="number"
                        value={s.capacity}
                        onChange={(e) => handleScheduleChange(idx, e)}
                        placeholder="20"
                        required
                        min="1"
                      />
                    </div>
                  </div>
                </div>
              ))}

              <button type="button" className="add-schedule-btn" onClick={addSchedule}>
                ➕ Agregar Otro Horario
              </button>
            </div>

            <div className="form-actions">
              <button type="submit" className="submit-btn">
                {mode === "edit" ? "💾 Guardar Cambios" : "✅ Crear Actividad"}
              </button>
              <button type="button" className="cancel-btn" onClick={() => navigate("/activities")}>
                ❌ Cancelar
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}

export default ActivityForm
