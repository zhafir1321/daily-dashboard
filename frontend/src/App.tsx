import { useState } from 'react'
import { LoginCard } from './components/LoginCard'
import { TodosView } from './components/TodosView'
import { RegisterView } from './components/RegisterView'

function App() {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem("token"))
  const [authScreen, setAuthScreen] = useState<'login' | 'register'>('login')



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
          <div>
            {authScreen === 'login' ? (
              <LoginCard onLoginSuccess={(token) => {
                setToken(token)
              }} onSwitchToRegister={() => setAuthScreen('register')} />
            ) : (
              <RegisterView onSwitchToLogin={() => setAuthScreen('login')} />
            )}
          </div>
        )}
      </div>
    </>
  )
}
export default App

