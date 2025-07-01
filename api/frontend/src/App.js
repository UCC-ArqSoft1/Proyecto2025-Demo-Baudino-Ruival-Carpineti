import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom"
import Login from "./Login"
import Home from "./inicio"
import ActivityDetails from "./ActivityDetails"
import MyActivities from "./MyActivities"
import { getCookie, getUserRole } from "./utils/cookies"
import ActivityForm from "./ActivityForm"

function PrivateRoute({ children }) {
  const token = getCookie("token")
  return token ? children : <Navigate to="/" />
}

function AdminRoute({ children }) {
  const token = getCookie("token")
  const isAdmin = getUserRole() === "admin"
  return token && isAdmin ? children : <Navigate to="/activities" />
}

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route
          path="/activities"
          element={
            <PrivateRoute>
              <Home />
            </PrivateRoute>
          }
        />
        <Route
          path="/activities/:id"
          element={
            <PrivateRoute>
              <ActivityDetails />
            </PrivateRoute>
          }
        />
        <Route
          path="/my-activities"
          element={
            <PrivateRoute>
              <MyActivities />
            </PrivateRoute>
          }
        />
        <Route
          path="/admin/create-activity"
          element={
            <AdminRoute>
              <ActivityForm mode="create" />
            </AdminRoute>
          }
        />
        <Route
          path="/admin/edit-activity/:id"
          element={
            <AdminRoute>
              <ActivityForm mode="edit" />
            </AdminRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  )
}

export default App
