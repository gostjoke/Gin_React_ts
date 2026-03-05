import { Calendar, Home, Inbox, Search, Settings, Users } from "lucide-react"

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "./ui/sidebar"

// Menu items.
const items = [
  {
    title: "Dashboard",
    url: "/",
    icon: Home,
  },
  {
    title: "Payments",
    url: "/payments",
    icon: Home,
  },
  {
    title: "PM",
    url: "/pm",
    icon: Home,
  },
  {
    title: "test",
    url: "/test",
    icon: Home,
  },
  {
    title: "Users",
    url: "/users",
    icon: Users,
  },
  {
    title: "Inbox",
    url: "/inbox",
    icon: Inbox,
  },
  {
    title: "Calendar",
    url: "/calendar",
    icon: Calendar,
  },
  {
    title: "Search",
    url: "/search",
    icon: Search,
  },
  {
    title: "Settings",
    url: "/settings",
    icon: Settings,
  },
]

const itemstest = Array(4).fill(items).flat()

export function AppSidebar() {
  return (
    <Sidebar className="w-64 border-r border-gray-200 bg-white">
      <SidebarContent className="p-4">
        <div className="mb-6 flex flex-col items-center">
          {/* <h2 className="text-lg font-bold text-gray-900 mb-1">Program Temp</h2> */}
          <img src="/images/FoxconnLogo.png" alt="dead image" style={{ width: '75%', height: 'auto' }} />
          {/* add an line to separate */}
          <div className="w-full border-t border-gray-200 mt-4"></div>
        </div>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu className="space-y-1">
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton 
                    asChild 
                    className="w-full flex items-center gap-3 px-3 py-2 text-gray-700 rounded-md hover:bg-gray-100 hover:text-gray-900 transition-colors"
                  >
                    <a href={item.url}>
                      <item.icon className="w-5 h-5" />
                      <span className="font-medium">{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}