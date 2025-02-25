import { Route, Routes } from 'react-router-dom'
import './App.css'
import Login from './pages/login'
import Register from './pages/register'
import Dashboard from './pages/dashboard'
import Unauthorized from './pages/unauthorized'
import NotFound from './pages/notfound'

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

        <Route path='/dashboard/*' element={
          <Dashboard/>
        } />

        <Route path='/unauthorized' element={<Unauthorized/>} />
        
        <Route path='*' element={<NotFound/>}/>
      </Routes>
    </>
  )
}

export default App
