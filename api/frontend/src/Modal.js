"use client"
import "./Modal.css"

function Modal({ show, message, success, onClose }) {
  if (!show) return null

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className={`modal-content ${success ? "success" : "error"}`} onClick={(e) => e.stopPropagation()}>
        <div className={`modal-message ${success ? "success" : "error"}`}>
          {success ? "✅" : "❌"} {message}
        </div>
        <button className="modal-close-btn" onClick={onClose}>
          Cerrar
        </button>
      </div>
    </div>
  )
}

export default Modal
