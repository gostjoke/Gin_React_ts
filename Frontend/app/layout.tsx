import type { Metadata } from "next"
import { SidebarProvider, SidebarTrigger, SidebarInset } from "../components/ui/sidebar"
import { AppSidebar } from "../components/app-sidebar"
import "./globals.css"

export const metadata: Metadata = {
  title: "Gin + Next.js App",
  description: "Full-stack application with Gin backend and Next.js frontend",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className="antialiased bg-gray-50">
        <SidebarProvider>
          <div className="flex min-h-screen">
            <AppSidebar />
            <SidebarInset className="flex-1">
              <div className="flex flex-col h-screen">
                <header className="bg-white border-b border-gray-200 px-6 py-4">
                  <div className="flex items-center gap-3">
                    <SidebarTrigger />
                    <h1 className="text-xl font-semibold text-gray-900">Dashboard</h1>
                  </div>
                </header>
                <main className="flex-1 overflow-auto bg-white">
                  <div className="p-6">
                    {children}
                  </div>
                </main>
              </div>
            </SidebarInset>
          </div>
        </SidebarProvider>
      </body>
    </html>
  )
}