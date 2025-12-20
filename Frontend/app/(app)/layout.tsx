import type { Metadata } from "next"
import {
  SidebarProvider,
  SidebarTrigger,
  SidebarInset,
} from "../../components/ui/sidebar"
import { AppSidebar } from "../../components/app-sidebar"
import { NavigationMenuCustom } from "../../components/app-navbar"
import "@/app/globals.css"

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
          {/* ===== Whole Screen ===== */}
          <div className="flex h-screen w-screen overflow-hidden">

            {/* ===== Left Sidebar ===== */}
            <AppSidebar />

            {/* ===== Right Content (FULL WIDTH) ===== */}
            <SidebarInset
              className="
                flex-1 flex flex-col
                w-full max-w-none
                overflow-hidden
                bg-white
              "
            >
              {/* ===== Header (FULL WIDTH) ===== */}
              <header className="shrink-0 sticky top-0 z-30 bg-white border-b w-full">
                <div className="h-12 flex items-center justify-between px-6 w-full">
                  <SidebarTrigger />
                  <NavigationMenuCustom />
                </div>
              </header>

              {/* ===== Main Content (Scrollable) ===== */}
              <main className="flex-1 overflow-auto bg-gray-50 w-full">
                <div className="p-6 w-full max-w-none">
                  {children}
                </div>
              </main>
            </SidebarInset>
          </div>
        </SidebarProvider>
      </body>
    </html>
  )
}
