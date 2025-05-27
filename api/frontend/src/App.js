<<<<<<< Updated upstream
import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Home from "./Home";
import ActivityDetails from "./ActivityDetails"; // 👈 Este archivo debe crearse

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/activities" element={<Home />} />
        <Route path="/activities/:id" element={<ActivityDetails />} />
      </Routes>
    </BrowserRouter>
=======
import React from 'react';
import './App.css';
import Home from './Home'; // ✅ Importación correcta

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>HELLO WORLD OF REACT</h1>
        <Home /> {/* ✅ Mostramos el componente Home */}
      </header>
    </div>
>>>>>>> Stashed changes
  );
}

export default App;