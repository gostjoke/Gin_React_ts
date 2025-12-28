import { columns, Payment } from "./columns"
import { DataTable } from "./data-table"
import { AppBreadcrumb } from "@/components/app-breadcrumb"

async function getData(): Promise<Payment[]> {
  // Fetch data from your API here.
  return [
    {
      id: "728ed52f",
      amount: 100,
      status: "pending",
      email: "m@example.com",
    },
    {
        id: "a3b1c9e4",
        amount: 250,
        status: "success",
        email: "aa@gmail.com",
    }
    // ...
  ]
}

export default async function DemoPage() {
  const data = await getData()

  return (
    <div className="container mx-auto py-2">
      <AppBreadcrumb 
        items={[
          { label: "首頁", href: "/" },
          { label: "付款管理" }
        ]} 
      />
      
      <div className="mt-6">
        <DataTable columns={columns} data={data} />
      </div>
    </div>
  )
}