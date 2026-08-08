import { useState } from 'react'
import { Button } from './components/ui/button'
import { LoginCard } from './components/LoginCard'
import { TodosView } from './components/TodosView'
import { EventsView } from './components/EventsView'
import { TransactionsView } from './components/TransactionsView'
import { RegisterView } from './components/RegisterView'

function App() {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem("token"))
  const [authScreen, setAuthScreen] = useState<'login' | 'register'>('login')
  const [view, setView] = useState<"todos" | "events" | "transactions">("todos")


  const handleLogout = () => {
    localStorage.removeItem("token")
    setToken(null)
  }

  return (
    <>
      <div className="min-h-screen flex items-center justify-center">
        {token ? (
          <div className="w-full max-w-md flex flex-col gap-4">
            <nav className="flex items-center gap-1 rounded-lg border p-1">
              <Button variant={view === "todos" ? "secondary" : "ghost"} size="sm" className="flex-1" onClick={() => setView("todos")}>Todos</Button>
              <Button variant={view === "events" ? "secondary" : "ghost"} size="sm" className="flex-1" onClick={() => setView("events")}>Calendar</Button>
              <Button variant={view === "transactions" ? "secondary" : "ghost"} size="sm" className="flex-1" onClick={() => setView("transactions")}>Money</Button>
              <Button variant="ghost" size="sm" onClick={handleLogout}>
                Logout
              </Button>
            </nav>

            {view === "todos" ? (
              <TodosView />
            ) : view === "events" ? (
              <EventsView />
            ) : (
              <TransactionsView />
            )}
          </div>
        ) : authScreen === 'login' ? (
          <LoginCard onLoginSuccess={(token) => {
            setToken(token)
          }} onSwitchToRegister={() => setAuthScreen('register')} />
        ) : (
          <RegisterView onSwitchToLogin={() => setAuthScreen('login')} />
        )}
      </div>
    </>
  )
}
export default App

