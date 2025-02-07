import { Route, Routes } from 'react-router-dom'
import './App.css'
import Login from './pages/login'
import Register from './pages/register'

function App() {

  return (
    <>
      <Routes>
        <Route path='/' element={
          <Login/>
        } />

        <Route path='/register' element={
          <Register/>
        } />
      </Routes>
    </>
  )
}

export default App
