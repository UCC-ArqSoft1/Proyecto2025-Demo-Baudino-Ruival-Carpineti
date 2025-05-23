import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Home from "./Home"; // Asegurate de que Home.js existe y está bien escrito

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/activities" element={<Home />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
