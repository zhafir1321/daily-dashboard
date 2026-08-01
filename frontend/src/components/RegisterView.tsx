import { Button } from "@/components/ui/button"
import {
    Card,
    CardAction,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import { register } from "@/api/auth"

// LAYOUT ONLY. Wire up the logic yourself:
// - state for each field (name, email, phone_number, password)
// - value + onChange on each Input
// - handleSubmit: preventDefault -> register(req) -> success flow
// - "Log in" action -> switch back to login (onSwitchToLogin prop from App)

export type RegisterReq = {
    email: string
    name: string
    password: string
    phone_number: string
}

type RegisterViewProps = {
    onSwitchToLogin?: () => void
}

export function RegisterView({ onSwitchToLogin }: RegisterViewProps) {
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const [phoneNumber, setPhoneNumber] = useState("")
    const [password, setPassword] = useState("")

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        try {
            const registerReq: RegisterReq = {
                name,
                email,
                phone_number: phoneNumber,
                password,
            }
            await register(registerReq)
            onSwitchToLogin?.()

        } catch (error) {
            console.error(error)
        }
    }

    return (
        <Card className="w-full max-w-sm">
            <CardHeader>
                <CardTitle>Create your account</CardTitle>
                <CardDescription>
                    Enter your details below to sign up
                </CardDescription>
                <CardAction>
                    {/* TODO: onClick={onSwitchToLogin} */}
                    <Button variant="link" onClick={onSwitchToLogin}>Log in</Button>
                </CardAction>
            </CardHeader>

            <CardContent>
                {/* TODO: onSubmit={handleSubmit} */}
                <form id="register-form" className="flex flex-col gap-4" onSubmit={handleSubmit}>
                    <div className="grid gap-2">
                        <Label htmlFor="name">Name</Label>
                        <Input id="name" type="text" placeholder="Name" required value={name} onChange={(e) => setName(e.target.value)} />
                    </div>

                    <div className="grid gap-2">
                        <Label htmlFor="email">Email</Label>
                        <Input id="email" type="email" placeholder="m@example.com" required value={email} onChange={(e) => setEmail(e.target.value)} />
                    </div>

                    <div className="grid gap-2">
                        <Label htmlFor="phone">Phone number</Label>
                        <Input id="phone" type="tel" placeholder="081234567890" required value={phoneNumber} onChange={(e) => setPhoneNumber(e.target.value)} />
                    </div>

                    <div className="grid gap-2">
                        <Label htmlFor="password">Password</Label>
                        <Input id="password" type="password" required value={password} onChange={(e) => setPassword(e.target.value)} />
                    </div>
                </form>
            </CardContent>

            <CardFooter>
                <Button type="submit" className="w-full" form="register-form">
                    Create account
                </Button>
            </CardFooter>
        </Card>
    )
}
