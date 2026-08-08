import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
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
import { Trash2, Plus, Pencil } from "lucide-react"
import type { CreateTransactionRequest, SummaryResponse, Transaction, TransactionType, UpdateTransactionRequest } from "@/types/transaction"
import { useEffect, useState } from "react"
import { createTransaction, deleteTransaction, getSummary, getTransactions, updateTransaction } from "@/api/transactions"

export function TransactionsView() {
    const [transactions, setTransactions] = useState<Transaction[]>([])
    const [summary, setSummary] = useState<SummaryResponse>({
        total_income: "0.00",
        total_expense: "0.00",
        balance: "0.00",
    })
    const [newType, setNewType] = useState<TransactionType>("income")
    const [newAmount, setNewAmount] = useState<string>("")
    const [newDescription, setNewDescription] = useState<string>("")
    const [newCategory, setNewCategory] = useState<string>("")
    const [newTransactionDate, setNewTransactionDate] = useState<string>("")

    const [editingTransaction, setEditingTransaction] = useState<Transaction | null>(null)
    const [editType, setEditType] = useState<TransactionType>("income")
    const [editAmount, setEditAmount] = useState<string>("")
    const [editDescription, setEditDescription] = useState<string>("")
    const [editCategory, setEditCategory] = useState<string>("")
    const [editTransactionDate, setEditTransactionDate] = useState<string>("")


    useEffect(() => {
        loadTransactions()
        loadSummary()
    }, [])

    const loadTransactions = async () => getTransactions().then((data) => setTransactions(data))
    const loadSummary = async () => getSummary().then((data) => setSummary(data))

    const handleAdd = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const transactionRequest: CreateTransactionRequest = {
                type: newType,
                amount: newAmount,
                description: newDescription,
                category: newCategory,
                transaction_date: newTransactionDate,
            }

            await createTransaction(transactionRequest)
            loadTransactions()
            loadSummary()
            setNewAmount("")
            setNewDescription("")
            setNewCategory("")
            setNewTransactionDate("")
        } catch (error) {
            console.error("Error adding transaction:", error)
        }
    }

    const handleEditSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            if (editingTransaction) {
                const updateRequest: UpdateTransactionRequest = {
                    type: editType,
                    amount: editAmount,
                    description: editDescription,
                    category: editCategory,
                    transaction_date: editTransactionDate,
                }
                await updateTransaction(editingTransaction.id, updateRequest)
                loadTransactions()
                loadSummary()
                setEditingTransaction(null)
                setEditAmount("")
                setEditDescription("")
                setEditCategory("")
                setEditTransactionDate("")
            }
        } catch (error) {
            console.error("Error editing transaction:", error)
        }
    }

    const handleDelete = async (transaction: Transaction) => {
        try {
            await deleteTransaction(transaction.id)
            loadTransactions()
            loadSummary()
        } catch (error) {
            console.error("Error deleting transaction:", error)
        }
    }

    return (
        <>
            {/* Summary cards */}
            <div className="grid grid-cols-3 gap-2">
                <div className="rounded-lg border p-3">
                    <p className="text-xs text-muted-foreground">Income</p>
                    <p className="text-lg font-medium text-green-600">{summary.total_income}</p>
                </div>
                <div className="rounded-lg border p-3">
                    <p className="text-xs text-muted-foreground">Expense</p>
                    <p className="text-lg font-medium text-red-600">{summary.total_expense}</p>
                </div>
                <div className="rounded-lg border p-3">
                    <p className="text-xs text-muted-foreground">Balance</p>
                    <p className="text-lg font-medium">{summary.balance}</p>
                </div>
            </div>

            <Card className="w-full max-w-md">
                <CardContent className="flex flex-col gap-4 pt-6">
                    <form className="flex flex-col gap-2" onSubmit={handleAdd}>
                        <div className="flex gap-1">
                            <Button type="button" size="sm" variant={newType === 'income' ? "secondary" : "ghost"} className="flex-1" onClick={() => setNewType('income')}>Income</Button>
                            <Button type="button" size="sm" variant={newType === 'expense' ? "secondary" : "ghost"} className="flex-1" onClick={() => setNewType("expense")}>Expense</Button>
                        </div>
                        <Input placeholder="Amount (e.g. 1500.00)" inputMode="decimal" value={newAmount} onChange={(e) => setNewAmount(e.target.value)} />
                        <Input placeholder="Description" value={newDescription} onChange={(e) => setNewDescription(e.target.value)} />
                        <div className="flex gap-2">
                            <Input placeholder="Category" className="flex-1" value={newCategory} onChange={(e) => setNewCategory(e.target.value)} />
                            <Input type="date" className="flex-1" value={newTransactionDate} onChange={(e) => setNewTransactionDate(e.target.value)} />
                            <Button type="submit">
                                <Plus className="mr-1 size-4" />
                                Add
                            </Button>
                        </div>
                    </form>

                    {transactions.length === 0 && (
                        <p className="py-6 text-center text-sm text-muted-foreground">
                            No transactions yet — add one above.
                        </p>
                    )}

                    {/* List */}
                    <div className="flex flex-col">
                        {transactions.map((transaction) => (
                            <div
                                key={transaction.id}
                                className="flex items-center gap-3 border-t py-3 first:border-t-0"
                            >
                                <div className="min-w-0 flex-1">
                                    <p className="font-medium">{transaction.description}</p>
                                    <p className="text-xs text-muted-foreground">
                                        {transaction.category} · {transaction.transaction_date}
                                    </p>
                                </div>

                                <p className={
                                    "font-medium tabular-nums " +
                                    (transaction.type === "income" ? "text-green-600" : "text-red-600")
                                }>
                                    {transaction.type === "income" ? "+" : "-"}{transaction.amount}
                                </p>

                                <Button variant="ghost" size="icon" aria-label="Edit transaction" onClick={() => {
                                    setEditingTransaction(transaction)
                                    setEditType(transaction.type)
                                    setEditAmount(transaction.amount)
                                    setEditDescription(transaction.description)
                                    setEditCategory(transaction.category)
                                    setEditTransactionDate(transaction.transaction_date)
                                }}>
                                    <Pencil className="size-4" />
                                </Button>

                                <Button variant="ghost" size="icon" aria-label="Delete transaction" className="text-destructive" onClick={() => handleDelete(transaction)}>
                                    <Trash2 className="size-4" />
                                </Button>
                            </div>
                        ))}
                    </div>
                </CardContent>
            </Card>

            <Dialog open={editingTransaction !== null} onOpenChange={(open) => { if (!open) setEditingTransaction(null) }}>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Edit transaction</DialogTitle>
                        <DialogDescription>Update the details below.</DialogDescription>
                    </DialogHeader>

                    <form id="edit-transaction-form" className="flex flex-col gap-4" onSubmit={handleEditSubmit}>
                        <div className="flex gap-1">
                            <Button type="button" size="sm" variant="secondary" className="flex-1" onClick={() => setEditType("income")}>Income</Button>
                            <Button type="button" size="sm" variant="ghost" className="flex-1" onClick={() => setEditType("expense")}>Expense</Button>
                        </div>
                        <div className="grid gap-2">
                            <Label htmlFor="edit-amount">Amount</Label>
                            <Input id="edit-amount" inputMode="decimal" value={editAmount} onChange={(e) => setEditAmount(e.target.value)} />
                        </div>
                        <div className="grid gap-2">
                            <Label htmlFor="edit-description">Description</Label>
                            <Input id="edit-description" value={editDescription} onChange={(e) => setEditDescription(e.target.value)} />
                        </div>
                        <div className="flex gap-2">
                            <div className="grid flex-1 gap-2">
                                <Label htmlFor="edit-category">Category</Label>
                                <Input id="edit-category" value={editCategory} onChange={(e) => setEditCategory(e.target.value)} />
                            </div>
                            <div className="grid flex-1 gap-2">
                                <Label htmlFor="edit-date">Date</Label>
                                <Input id="edit-date" type="date" value={editTransactionDate} onChange={(e) => setEditTransactionDate(e.target.value)} />
                            </div>
                        </div>
                    </form>

                    <DialogFooter>
                        <Button type="button" variant="outline" onClick={() => setEditingTransaction(null)}>
                            Cancel
                        </Button>
                        <Button type="submit" form="edit-transaction-form">
                            Save
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>
        </>
    )
}
