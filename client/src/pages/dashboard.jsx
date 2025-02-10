import { Link, Route, Routes } from "react-router-dom"
import ProtectedRoute from "../components/protectedroute"
import UserPage from "./user"
import ArtistPage from "./artist"
import { Content, Header } from "antd/es/layout/layout"
import MusicPage from "./music"

export default function Dashboard() {
    const role = localStorage.getItem("role")

    return(
        <>
            <Header
                style={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    color: "white",
                }}
            >
                <h1>Artist Management System</h1>
                <div
                    style={{
                        display: "flex",
                        gap: "40px"
                    }}
                >
                    <div
                        style={{
                            display: "flex",
                            gap: "10px"
                        }}
                    >
                        {role === 'super_admin' && <Link to={"/dashboard/users"}>Users</Link>}
                        <Link to={"/dashboard/artists"}>Artists</Link>
                    </div>
                    <div>
                        <span>Logout</span>
                    </div>
                </div>
            </Header>
            <Content>
                <Routes>
                    <Route path="users" element={
                        <ProtectedRoute role={role} route={"/users"}>
                            <UserPage />
                        </ProtectedRoute>
                    } />
                    <Route path="artists" element={
                        <ProtectedRoute role={role} route={"/artists"}>
                            <ArtistPage isManager={role === 'super_admin' || role === 'artist_manager'} />
                        </ProtectedRoute>
                    } />
                    <Route path="artist/:artistID" element={
                        <ProtectedRoute role={role} route={"/artist"}>
                            <MusicPage />
                        </ProtectedRoute>
                    } />
                </Routes>
            </Content>
        </>
    )
}