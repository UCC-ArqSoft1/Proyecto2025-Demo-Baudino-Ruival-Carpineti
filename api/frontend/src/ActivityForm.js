import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getCookie } from "./utils/cookies";

const emptySchedule = { week_day: "", start_time: "", end_time: "", capacity: 0 };

function ActivityForm({ mode }) {
  const [form, setForm] = useState({
    title: "",
    description: "",
    category: "",
    instructor: "",
    duration: 60,
    image: "",
    status: "activo",
    schedules: [ { ...emptySchedule } ]
  });
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { id } = useParams();

  useEffect(() => {
    if (mode === "edit" && id) {
      fetch(`http://localhost:8080/activities/${id}`)
        .then(res => res.json())
        .then(data => {
          setForm({
            title: data.title,
            description: data.description,
            category: data.category,
            instructor: data.instructor,
            duration: data.duration,
            image: data.image,
            status: data.status,
            schedules: data.schedules.map(s => ({
              week_day: s.week_day,
              start_time: s.start_time,
              end_time: s.end_time,
              capacity: s.capacity
            }))
          });
        });
    }
  }, [mode, id]);

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleScheduleChange = (idx, e) => {
    const newSchedules = form.schedules.map((s, i) =>
      i === idx ? { ...s, [e.target.name]: e.target.value } : s
    );
    setForm({ ...form, schedules: newSchedules });
  };

  const addSchedule = () => {
    setForm({ ...form, schedules: [...form.schedules, { ...emptySchedule }] });
  };

  const removeSchedule = idx => {
    setForm({ ...form, schedules: form.schedules.filter((_, i) => i !== idx) });
  };

  const handleSubmit = async e => {
    e.preventDefault();
    setError("");
    const token = getCookie("token");
    const url = mode === "edit" ? `http://localhost:8080/admin/activities/${id}` : "http://localhost:8080/admin/activities";
    const method = mode === "edit" ? "PUT" : "POST";
    const mappedForm = {
      ...form,
      duration: Number(form.duration),
      schedules: form.schedules.map(s => ({
        week_day: s.week_day,
        start_time: s.start_time,
        end_time: s.end_time,
        capacity: Number(s.capacity)
      }))
    };
    const res = await fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
        "Authorization": token
      },
      body: JSON.stringify(mappedForm)
    });
    if (res.ok) {
      navigate("/activities");
    } else {
      let errorMsg = "Error al guardar la actividad";
      try {
        const data = await res.json();
        if (data && data.error) errorMsg = data.error;
      } catch {}
      setError(errorMsg);
    }
  };

  return (
    <div style={{ padding: 30 }}>
      <h2>{mode === "edit" ? "Editar" : "Crear"} Actividad</h2>
      {error && <div style={{ color: "red" }}>{error}</div>}
      <form onSubmit={handleSubmit}>
        <input name="title" value={form.title} onChange={handleChange} placeholder="Título" required />
        <input name="description" value={form.description} onChange={handleChange} placeholder="Descripción" required />
        <input name="category" value={form.category} onChange={handleChange} placeholder="Categoría" required />
        <input name="instructor" value={form.instructor} onChange={handleChange} placeholder="Instructor" required />
        <input name="duration" type="number" value={form.duration} onChange={handleChange} placeholder="Duración (min)" required />
        <input name="image" value={form.image} onChange={handleChange} placeholder="Imagen (URL)" />
        <select name="status" value={form.status} onChange={handleChange} required>
          <option value="activo">Activo</option>
          <option value="inactivo">Inactivo</option>
        </select>
        <h3>Horarios</h3>
        {form.schedules.map((s, idx) => (
          <div key={idx} style={{ marginBottom: 10, border: "1px solid #ccc", padding: 10 }}>
            <input name="week_day" value={s.week_day} onChange={e => handleScheduleChange(idx, e)} placeholder="Día" required />
            <input name="start_time" value={s.start_time} onChange={e => handleScheduleChange(idx, e)} placeholder="Inicio (HH:MM)" required />
            <input name="end_time" value={s.end_time} onChange={e => handleScheduleChange(idx, e)} placeholder="Fin (HH:MM)" required />
            <input name="capacity" type="number" value={s.capacity} onChange={e => handleScheduleChange(idx, e)} placeholder="Cupo" required />
            <button type="button" onClick={() => removeSchedule(idx)} disabled={form.schedules.length === 1}>Eliminar</button>
          </div>
        ))}
        <button type="button" onClick={addSchedule}>Agregar horario</button>
        <br /><br />
        <button type="submit">{mode === "edit" ? "Guardar cambios" : "Crear actividad"}</button>
        <button type="button" onClick={() => navigate("/activities")}>Cancelar</button>
      </form>
    </div>
  );
}

export default ActivityForm; 