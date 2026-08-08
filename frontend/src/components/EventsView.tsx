import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogDescription,
    DialogFooter,
} from "@/components/ui/dialog"
import { Trash2, Plus, Pencil, Calendar, Clock } from "lucide-react"
import type { CreateEventRequest, Event, UpdateEventRequest } from "@/types/event"
import { useEffect, useState } from "react"
import { createEvent, deleteEvent, getEvents, updateEvent } from "@/api/events"

export function EventsView() {
    const [newTitle, setNewTitle] = useState<string>("")
    const [newDescription, setNewDescription] = useState<string>("")
    const [newDateEvent, setNewDateEvent] = useState<string>("")
    const [newDateTime, setNewDateTime] = useState<string>("")
    const [events, setEvents] = useState<Event[]>([])
    const [editingEvent, setEditingEvent] = useState<Event | null>(null)
    const [editTitle, setEditTitle] = useState<string>("")
    const [editDescription, setEditDescription] = useState<string>("")
    const [editDateEvent, setEditDateEvent] = useState<string>("")
    const [editDateTime, setEditDateTime] = useState<string>("")


    useEffect(() => {
        loadEvents()
    }, [])

    const loadEvents = async () => getEvents().then((events) => setEvents(events))

    const handleAdd = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        try {
            const eventRequest: CreateEventRequest = {
                title: newTitle,
                description: newDescription,
                event_date: newDateEvent,
                event_time: newDateTime || null,
            };
            await createEvent(eventRequest);
            await loadEvents();
            setNewTitle("");
            setNewDescription("");
            setNewDateEvent("");
            setNewDateTime("");
        } catch (error) {
            console.error("Error adding event:", error);
        }
    }

    const handleEdit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        try {
            if (editingEvent) {
                const updatedEvent: UpdateEventRequest = {
                    title: editTitle,
                    description: editDescription,
                    event_date: editDateEvent,
                    event_time: editDateTime || null,
                }
                await updateEvent(editingEvent.id, updatedEvent)
                await loadEvents()
                setEditingEvent(null)
                setEditTitle("")
                setEditDescription("")
                setEditDateEvent("")
                setEditDateTime("")
            }
        } catch (error) {
            console.error("Error editing event:", error);
        }
    }

    const handleDelete = async (event: Event) => {
        try {
            await deleteEvent(event.id);
            await loadEvents();
        } catch (error) {
            console.error("Error deleting event:", error);
        }
    }



    return (
        <>
            <Card className="w-full max-w-md">
                <CardHeader>
                    <h1 className="text-lg font-medium">Your calendar</h1>
                    <p className="text-sm text-muted-foreground">{events.length} events</p>
                </CardHeader>

                <CardContent className="flex flex-col gap-4">
                    <form className="flex flex-col gap-2" onSubmit={handleAdd}>
                        <Input placeholder="Event title" value={newTitle} onChange={(e) => setNewTitle(e.target.value)} />
                        <Input placeholder="Description" value={newDescription} onChange={(e) => setNewDescription(e.target.value)} />
                        <div className="flex gap-2">
                            <Input type="date" className="flex-1" value={newDateEvent} onChange={(e) => setNewDateEvent(e.target.value)} />
                            <Input type="time" className="flex-1" value={newDateTime} onChange={(e) => setNewDateTime(e.target.value)} />
                            <Button type="submit">
                                <Plus className="mr-1 size-4" />
                                Add
                            </Button>
                        </div>
                    </form>

                    {events.length === 0 && (
                        <p className="py-6 text-center text-sm text-muted-foreground">
                            No events yet — add one above.
                        </p>
                    )}

                    {/* List */}
                    <div className="flex flex-col">
                        {events.map((event) => (
                            <div
                                key={event.id}
                                className="flex items-start gap-3 border-t py-3 first:border-t-0"
                            >
                                <div className="min-w-0 flex-1">
                                    <p className="font-medium">{event.title}</p>
                                    <p className="text-sm text-muted-foreground">{event.description}</p>
                                    <div className="mt-1 flex items-center gap-3 text-xs text-muted-foreground">
                                        <span className="flex items-center gap-1">
                                            <Calendar className="size-3" />
                                            {event.event_date}
                                        </span>
                                        <span className="flex items-center gap-1">
                                            <Clock className="size-3" />
                                            {event.event_time ?? "All day"}
                                        </span>
                                    </div>
                                </div>

                                <Button variant="ghost" size="icon" aria-label="Edit event" onClick={() => {
                                    setEditingEvent(event);
                                    setEditTitle(event.title);
                                    setEditDescription(event.description);
                                    setEditDateEvent(event.event_date);
                                    setEditDateTime(event.event_time ?? "");
                                }}>
                                    <Pencil className="size-4" />
                                </Button>

                                <Button variant="ghost" size="icon" aria-label="Delete event" className="text-destructive" onClick={() => handleDelete(event)}>
                                    <Trash2 className="size-4" />
                                </Button>
                            </div>
                        ))}
                    </div>
                </CardContent>
            </Card>
            <Dialog open={editingEvent !== null} onOpenChange={(open) => { if (!open) setEditingEvent(null) }}>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Edit event</DialogTitle>
                        <DialogDescription>Update the details below.</DialogDescription>
                    </DialogHeader>

                    <form id="edit-event-form" className="flex flex-col gap-4" onSubmit={handleEdit}>
                        <div className="grid gap-2">
                            <Label htmlFor="edit-event-title">Title</Label>
                            <Input id="edit-event-title" value={editTitle} onChange={(e) => setEditTitle(e.target.value)} />
                        </div>
                        <div className="grid gap-2">
                            <Label htmlFor="edit-event-description">Description</Label>
                            <Input id="edit-event-description" value={editDescription} onChange={(e) => setEditDescription(e.target.value)} />
                        </div>
                        <div className="flex gap-2">
                            <div className="grid flex-1 gap-2">
                                <Label htmlFor="edit-event-date">Date</Label>
                                <Input id="edit-event-date" type="date" value={editDateEvent} onChange={(e) => setEditDateEvent(e.target.value)} />
                            </div>
                            <div className="grid flex-1 gap-2">
                                <Label htmlFor="edit-event-time">Time</Label>
                                <Input id="edit-event-time" type="time" value={editDateTime} onChange={(e) => setEditDateTime(e.target.value)} />
                            </div>
                        </div>
                    </form>

                    <DialogFooter>
                        <Button type="button" variant="outline" onClick={() => setEditingEvent(null)}>
                            Cancel
                        </Button>
                        <Button type="submit" form="edit-event-form">
                            Save
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>
        </>
    )
}
