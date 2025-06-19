import React from "react";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from "./Login";
import Home from "./inicio";
import ActivityDetails from "./ActivityDetails";
import { getCookie } from "./utils/cookies";

function PrivateRoute({ children }) {
  const token = getCookie("token");
  return token ? children : <Navigate to="/" />;
}

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/activities" element={<PrivateRoute><Home /></PrivateRoute>} />
        <Route path="/activities/:id" element={<PrivateRoute><ActivityDetails /></PrivateRoute>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;