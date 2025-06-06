import React from "react";

function Modal({ show, message, success, onClose }) {
  if (!show) return null;

  const modalStyle = {
    position: "fixed",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    backgroundColor: success ? "#d4edda" : "#f8d7da",
    color: success ? "#155724" : "#721c24",
    border: `1px solid ${success ? "#c3e6cb" : "#f5c6cb"}`,
    padding: "20px",
    borderRadius: "8px",
    zIndex: 9999,
    minWidth: "300px",
    textAlign: "center",
  };

  return (
    <div style={modalStyle}>
      <p>{message}</p>
      <button onClick={onClose} style={{
        marginTop: "10px",
        padding: "8px 16px",
        borderRadius: "5px",
        border: "none",
        backgroundColor: "#007bff",
        color: "white",
        cursor: "pointer"
      }}>
        Cerrar
      </button>
    </div>
  );
}

export default Modal;
