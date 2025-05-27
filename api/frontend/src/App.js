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
  );
}

export default App;