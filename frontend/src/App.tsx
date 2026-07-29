import { useState } from 'react'
import { LoginCard } from './components/LoginCard'

function App() {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem("token"))



  return (
    <>
      <div className="min-h-screen flex items-center justify-center">
        {token ? (
          <div>
            <h1 className="text-2xl font-bold mb-4">Welcome!</h1>
            <p className="mb-4">You are logged in.</p>
            <button
              className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              onClick={() => {
                localStorage.removeItem("token")
                setToken(null)
              }}
            >
              Logout
            </button>
          </div>
        ) : (
          <LoginCard onLoginSuccess={setToken} />
        )}
      </div>
    </>
  )
}
export default App

