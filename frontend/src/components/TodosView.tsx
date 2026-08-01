import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Checkbox } from "@/components/ui/checkbox"
import { Trash2, Plus, LogOut } from "lucide-react"
import { useEffect, useState } from "react"
import { deleteTodo, getTodos, toggleTodoCompletion, toggleTodoIncompletion } from "@/api/todos"
import { createTodo } from "@/api/todos"

// This is the DESIGN/LAYOUT only. The logic is left as TODOs for you.
// Before using: `npx shadcn@latest add checkbox`

export type Todo = {
    id: string;
    title: string;
    description: string;
    completed: boolean;
    created_at: string;
    updated_at: string;
}

type TodosViewProps = {
    onLogout?: () => void
}

export function TodosView({ onLogout }: TodosViewProps) {
    const [todos, setTodos] = useState<Todo[]>([])
    const [newTitle, setNewTitle] = useState("")
    const [newDescription, setNewDescription] = useState("")


    useEffect(() => {
        loadTodos()
    }, [])

    const remaining = todos.filter((t) => !t.completed).length

    const loadTodos = () => getTodos().then((fetchedTodos) => {
        setTodos(fetchedTodos)
    })

    const handleLogout = () => {
        // Clear the authentication token
        localStorage.removeItem("token");
        onLogout?.();
    }

    const handleAddTodo = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        try {
            console.log("Adding todo:", newTitle, newDescription)
            await createTodo(newTitle, newDescription)
            loadTodos()
            setNewTitle("")
            setNewDescription("")
        } catch (error) {
            console.error("Error adding todo:", error)
        }
    }

    const handleToggle = async (todo: Todo) => {
        if (todo.completed) {
            console.log("Toggling todo to incomplete:", todo.id)
            await toggleTodoIncompletion(todo.id)
        }
        if (!todo.completed) {
            console.log("Toggling todo to complete:", todo.id)
            await toggleTodoCompletion(todo.id)
        }
        loadTodos()
    }

    const handleDelete = async (todo: Todo) => {
        console.log("Deleting todo:", todo.id)
        await deleteTodo(todo.id)
        loadTodos()
    }

    return (
        <Card className="w-full max-w-md">
            <CardHeader>
                <div className="flex items-baseline justify-between">
                    <div>
                        <h1 className="text-lg font-medium">Your todos</h1>
                        <p className="text-sm text-muted-foreground">
                            {remaining} of {todos.length} remaining
                        </p>
                    </div>
                    <Button variant="ghost" size="sm" onClick={handleLogout}>
                        <LogOut className="mr-1 size-4" />
                        Logout
                    </Button>
                </div>
            </CardHeader>

            <CardContent className="flex flex-col gap-4">
                <form className="flex flex-col gap-2" onSubmit={handleAddTodo}>
                    <Input placeholder="Task title" onChange={(e) => setNewTitle(e.target.value)} value={newTitle} />
                    <div className="flex gap-2">
                        <Input placeholder="Description" className="flex-1" onChange={(e) => setNewDescription(e.target.value)} value={newDescription} />
                        <Button type="submit">
                            <Plus className="mr-1 size-4" />
                            Add
                        </Button>
                    </div>
                </form>

                {todos.length === 0 && (
                    <p className="py-6 text-center text-sm text-muted-foreground">
                        No todos yet — add one above.
                    </p>
                )}

                {/* List */}
                <div className="flex flex-col">
                    {todos.map((todo) => (
                        <div
                            key={todo.id}
                            className="flex items-start gap-3 border-t py-3 first:border-t-0"
                        >
                            <Checkbox checked={todo.completed} onCheckedChange={() => handleToggle(todo)} className="mt-1" />

                            <div className="min-w-0 flex-1">
                                <p
                                    className={
                                        todo.completed
                                            ? "text-muted-foreground line-through"
                                            : ""
                                    }
                                >
                                    {todo.title}
                                </p>
                                <p
                                    className={
                                        "text-sm " +
                                        (todo.completed
                                            ? "text-muted-foreground line-through"
                                            : "text-muted-foreground")
                                    }
                                >
                                    {todo.description}
                                </p>
                            </div>

                            <Button
                                variant="ghost"
                                size="icon"
                                aria-label="Delete task"
                                className="text-destructive"
                                onClick={() => handleDelete(todo)}
                            >
                                <Trash2 className="size-4" />
                            </Button>
                        </div>
                    ))}
                </div>
            </CardContent>
        </Card>
    )
}
