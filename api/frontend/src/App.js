import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Home from "./Home";

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
