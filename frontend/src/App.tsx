import { useState } from 'react'
import { Button } from './components/ui/button'
import { LoginCard } from './components/LoginCard'
import { TodosView } from './components/TodosView'
import { EventsView } from './components/EventsView'
import { RegisterView } from './components/RegisterView'

function App() {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem("token"))
  const [authScreen, setAuthScreen] = useState<'login' | 'register'>('login')
  const [view, setView] = useState<"todos" | "events">("todos")


  const handleLogout = () => {
    localStorage.removeItem("token")
    setToken(null)
  }

  return (
    <>
      <div className="min-h-screen flex items-center justify-center">
        {token ? (
          // LOGGED IN — LAYOUT ONLY, you wire the logic:
          // 1. add: const [view, setView] = useState<"todos" | "events">("todos")
          // 2. nav buttons: variant={view === "todos" ? "secondary" : "ghost"} + onClick={() => setView("todos")}  (same for events)
          // 3. logout button onClick: localStorage.removeItem("token"); setToken(null)
          // 4. swap the <TodosView /> below for: {view === "todos" ? <TodosView /> : <EventsView />}
          <div className="w-full max-w-md flex flex-col gap-4">
            <nav className="flex items-center gap-1 rounded-lg border p-1">
              <Button variant={view === "todos" ? "secondary" : "ghost"} size="sm" className="flex-1" onClick={() => setView("todos")} >Todos</Button>
              <Button variant={view === "events" ? "secondary" : "ghost"} size="sm" className="flex-1" onClick={() => setView("events")}>Calendar</Button>
              <Button variant="ghost" size="sm" onClick={handleLogout}>
                Logout
              </Button>
            </nav>

            {view === "todos" ? <TodosView /> : <EventsView />}
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

