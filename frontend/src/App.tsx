import { useState } from 'react'
import { LoginCard } from './components/LoginCard'
import { TodosView } from './components/TodosView'

function App() {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem("token"))



  return (
    <>
      <div className="min-h-screen flex items-center justify-center">
        {token ? (
          <div>
            <TodosView onLogout={() => {
              setToken(null)
            }} />
          </div>
        ) : (
          <LoginCard onLoginSuccess={setToken} />
        )}
      </div>
    </>
  )
}
export default App

