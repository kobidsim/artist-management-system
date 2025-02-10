import { Navigate } from "react-router-dom"

const allowedRoutes = {
    "super_admin": ["/users", "/artists", "/artist"],
    "artist_manager": ["/artists", "/artist"],
    "artist": ["/artists", "/artist"]

}

export default function ProtectedRoute({role, route, children}) {

    return (
        <>
            {allowedRoutes[role].includes(route) ? children : <Navigate to="/unauthorized" />}
        </>
    )
}