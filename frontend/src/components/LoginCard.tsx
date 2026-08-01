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
import { Input } from "./ui/input"
import { useState } from "react"
import { login } from "@/api/auth"

type LoginCardProps = {
    onLoginSuccess?: (token: string) => void
    onSwitchToRegister?: () => void
}

export function LoginCard({ onLoginSuccess, onSwitchToRegister }: LoginCardProps) {
    const [email, setEmail] = useState<string>("")
    const [password, setPassword] = useState<string>("")

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        console.log("Email:", email)
        console.log("Password:", password)

        try {
            const response = await login(email, password)
            if (response && response.token) {
                console.log("Login successful. Token:", response.token)
                localStorage.setItem("token", response.token)
                if (onLoginSuccess) {
                    onLoginSuccess(response.token)
                }
            } else {
                console.error("Login failed. Response:", response)
                throw new Error("Login failed. No token received.")
            }
        } catch (error) {
            console.error("Login failed:", error)
        }
    }

    return (
        <Card className="w-full max-w-sm">
            <CardHeader>
                <CardTitle>Login to your account</CardTitle>
                <CardDescription>
                    Enter your email below to login to your account
                </CardDescription>
                <CardAction>
                    <Button variant="link" onClick={onSwitchToRegister}>
                        Sign Up
                    </Button>
                </CardAction>
            </CardHeader>
            <CardContent>
                <form id="login-form" className="grid gap-4" onSubmit={handleSubmit}>
                    <div className="flex flex-col gap-6">
                        <div className="grid gap-2">
                            <Label htmlFor="email">Email</Label>
                            <Input
                                id="email"
                                type="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                placeholder="m@example.com"
                                required
                            />
                        </div>
                        <div className="grid gap-2">
                            <div className="flex items-center">
                                <Label htmlFor="password">Password</Label>
                                <a
                                    href="#"
                                    className="ml-auto inline-block text-sm underline-offset-4 hover:underline"
                                >
                                    Forgot your password?
                                </a>
                            </div>
                            <Input
                                id="password"
                                type="password"
                                onChange={(e) => setPassword(e.target.value)}
                                value={password}
                                required
                            />
                        </div>
                    </div>
                </form>
            </CardContent>
            <CardFooter className="flex-col gap-2">
                <Button type="submit" className="w-full" form="login-form">
                    Login
                </Button>
            </CardFooter>
        </Card>
    )
}
