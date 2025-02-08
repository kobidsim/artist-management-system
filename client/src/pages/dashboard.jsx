import { Route, Routes } from "react-router-dom"
import ProtectedRoute from "../components/protectedroute"
import UserPage from "./user"
import ArtistPage from "./artist"

export default function Dashboard() {
    const role = localStorage.getItem("role")

    return(
        <>
            <Routes>
                <Route path="users" element={
                    <ProtectedRoute role={role} route={"/users"}>
                        <UserPage />
                    </ProtectedRoute>
                } />
                <Route path="artists" element={
                    <ProtectedRoute role={role} route={"/artists"}>
                        <ArtistPage />
                    </ProtectedRoute>
                } />
            </Routes>
        </>
    )
}