import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Checkbox } from "@/components/ui/checkbox"
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogDescription,
    DialogFooter,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import { Trash2, Plus, Pencil, AlertCircle } from "lucide-react"
import { useEffect, useState } from "react"
import { deleteTodo, getTodos, toggleTodoCompletion, toggleTodoIncompletion, updateTodo } from "@/api/todos"
import { createTodo } from "@/api/todos"
import type { CreateTodoRequest, Todo, UpdateTodoRequest } from "@/types/todo"
import { Alert, AlertDescription } from "./ui/alert"


export function TodosView() {
    const [todos, setTodos] = useState<Todo[]>([])
    const [newTitle, setNewTitle] = useState<string>("")
    const [newDescription, setNewDescription] = useState<string>("")
    const [filter, setFilter] = useState<"all" | "completed" | "incomplete">("all")
    const [editingTodo, setEditingTodo] = useState<Todo | null>(null)
    const [editTitle, setEditTitle] = useState<string>("")
    const [editDescription, setEditDescription] = useState<string>("")

    const [error, setError] = useState<string | null>(null)
    const [editError, setEditError] = useState<string | null>(null)





    useEffect(() => {
        loadTodos()
    }, [])

    const remaining = todos.filter((t) => !t.completed).length

    const loadTodos = () => getTodos().then((fetchedTodos) => {
        setTodos(fetchedTodos)
    })

    const handleAddTodo = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        setError(null)
        try {
            const newTodo: CreateTodoRequest = {
                title: newTitle,
                description: newDescription
            };
            await createTodo(newTodo)
            loadTodos()
            setNewTitle("")
            setNewDescription("")
        } catch (error) {
            console.error("Error adding todo:", error)
            setError(error instanceof Error ? error.message : "An unknown error occurred.")
        }
    }

    const handleToggle = async (todo: Todo) => {
        try {
            if (todo.completed) {
                await toggleTodoIncompletion(todo.id)
            }
            if (!todo.completed) {
                await toggleTodoCompletion(todo.id)
            }
            loadTodos()
        } catch (error) {
            console.error("Error toggling todo completion:", error)
            setError(error instanceof Error ? error.message : "An unknown error occurred.")
        }
    }

    const handleDelete = async (todo: Todo) => {
        setError(null)
        try {
            await deleteTodo(todo.id)
            loadTodos()

        } catch (error) {
            console.error("Error deleting todo:", error)
            setError(error instanceof Error ? error.message : "An unknown error occurred.")
        }
    }

    const handleEditSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        setEditError(null)
        try {
            if (editingTodo) {
                const updateRequest: UpdateTodoRequest = {
                    title: editTitle,
                    description: editDescription
                };
                await updateTodo(editingTodo.id, updateRequest)
                loadTodos()
                setEditingTodo(null)
            }
        } catch (error) {
            console.error("Error updating todo:", error)
            setEditError(error instanceof Error ? error.message : "An unknown error occurred.")
        }
    }

    let visibleTodos = todos
    if (filter === "completed") {
        visibleTodos = todos.filter((t) => t.completed)
    }
    if (filter === "incomplete") {
        visibleTodos = todos.filter((t) => !t.completed)
    }


    return (
        <>
            <Card className="w-full max-w-md">
                <CardHeader>
                    <div className="flex items-baseline justify-between">
                        <div>
                            <h1 className="text-lg font-medium">Your todos</h1>
                            <p className="text-sm text-muted-foreground">
                                {remaining} of {todos.length} remaining
                            </p>
                        </div>
                    </div>
                </CardHeader>

                <CardContent className="flex flex-col gap-4">
                    {error && (
                        <Alert variant="destructive" className="mt-2">
                            <AlertCircle className='h-4 w-4' />
                            <AlertDescription className='ml-2'>
                                {error}
                            </AlertDescription>

                        </Alert>
                    )}
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

                    <div className="flex gap-1">
                        <Button type="button" size="sm" variant={filter === "all" ? "secondary" : "ghost"} onClick={() => setFilter("all")}>
                            All
                        </Button>
                        <Button type="button" size="sm" variant={filter === "completed" ? "secondary" : "ghost"} onClick={() => setFilter("completed")}>
                            Completed
                        </Button>
                        <Button type="button" size="sm" variant={filter === "incomplete" ? "secondary" : "ghost"} onClick={() => setFilter("incomplete")}>
                            Incomplete
                        </Button>
                    </div>

                    {visibleTodos.length === 0 && filter === "all" && (
                        <p className="py-6 text-center text-sm text-muted-foreground">
                            No todos yet — add one above.
                        </p>
                    )}

                    {visibleTodos.length === 0 && filter === "completed" && (
                        <p className="py-6 text-center text-sm text-muted-foreground">
                            No completed todos.
                        </p>
                    )}

                    {visibleTodos.length === 0 && filter === "incomplete" && (
                        <p className="py-6 text-center text-sm text-muted-foreground">
                            No incomplete todos.
                        </p>
                    )}

                    {/* List */}
                    <div className="flex flex-col">
                        {visibleTodos.map((todo) => (
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
                                    aria-label="Edit task"
                                    onClick={() => {
                                        setEditingTodo(todo)
                                        setEditTitle(todo.title)
                                        setEditDescription(todo.description)
                                    }}
                                >
                                    <Pencil className="size-4" />
                                </Button>

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
            <Dialog open={editingTodo !== null} onOpenChange={(open) => { if (!open) setEditingTodo(null) }}>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Edit todo</DialogTitle>
                        <DialogDescription>Update the title or description.</DialogDescription>
                    </DialogHeader>
                    {editError && (
                        <Alert variant="destructive" className="mt-2">
                            <AlertCircle className='h-4 w-4' />
                            <AlertDescription className='ml-2'>
                                {editError}
                            </AlertDescription>
                        </Alert>
                    )}

                    <form id="edit-form" className="flex flex-col gap-4" onSubmit={handleEditSubmit}>
                        <div className="grid gap-2">
                            <Label htmlFor="edit-title">Title</Label>
                            <Input id="edit-title" value={editTitle} onChange={(e) => setEditTitle(e.target.value)} />
                        </div>
                        <div className="grid gap-2">
                            <Label htmlFor="edit-description">Description</Label>
                            <Input id="edit-description" value={editDescription} onChange={(e) => setEditDescription(e.target.value)} />
                        </div>
                    </form>

                    <DialogFooter>
                        <Button type="button" variant="outline" onClick={() => setEditingTodo(null)}>
                            Cancel
                        </Button>
                        <Button type="submit" form="edit-form">
                            Save
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>
        </>
    )
}
