"use client"

import { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import "./Login.css"
import { getCookie } from "./utils/cookies"

function Login() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")
  const navigate = useNavigate()

  useEffect(() => {
    const token = getCookie("token")
    if (token) {
      navigate("/activities")
    }
  }, [navigate])

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError("")

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      })

      if (!response.ok) {
        const data = await response.json()
        setError(data.error || "Credenciales incorrectas")
        return
      }

      const data = await response.json()
      document.cookie = `token=${data.token}; path=/; SameSite=Strict`
      navigate("/activities")
    } catch (err) {
      setError("Error al iniciar sesión")
    }
  }

  return (
    <div className="login-container">
      <div className="login-form">
        <h2>Iniciar Sesión</h2>
        {error && <div className="error">{error}</div>}
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Usuario"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <input
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button type="submit">Entrar</button>
        </form>
      </div>
    </div>
  )
}

export default Login
