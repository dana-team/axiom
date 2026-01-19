import { Outlet } from "react-router-dom";
import { Nav } from "@/Nav";

export function Layout() {
  return (
    <div className="flex flex-col h-screen">
      <header>
        <Nav />
      </header>
      <main className="flex-1 overflow-auto bg-gradient-to-br dark:from-primary/30 via-primary-dark/10 dark:to-primary-light/20">
        <Outlet />
      </main>
    </div>
  );
}
