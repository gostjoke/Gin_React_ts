// app/page.tsx
import { AppBreadcrumb } from "@/components/app-breadcrumb"
import { ChatDemo } from "./chatbot"

export default function HomePage() {
  return (
    <>
      <div className="container mx-auto py-2">
        <AppBreadcrumb 
          items={[
            { label: "首頁"},
          ]} 
        />    
      </div>
      <div>
        <h1 className="text-xl font-bold">
          Dashboard
        </h1>
        <p className="text-muted-foreground">
          Layout is working.
        </p>
        <ChatDemo />
      </div>
    </>
  )
}
