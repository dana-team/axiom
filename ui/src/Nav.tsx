import { Link,useNavigate } from "react-router-dom";
import ThemeToggle from "@/components/ThemeToggle";
// import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Button, buttonVariants } from "@/components/ui/button";
import { BookMarked, LogOut, Search } from "lucide-react";
import { AnimatePresence, motion } from "framer-motion";
import { Input } from "@/components/ui/input";
import { useEffect, useState } from "react";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
} from "@/components/ui/accordion";
import NavFiltersDrawer from "@/components/NavFiltersDrawer";
import { useDebounceValue } from "@/hooks/use-debounce-value";
import { fetchFilteredClusters } from "./ClusterService";

const MotionLink = motion(Link);
const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;

export function Nav() {
  const navigate = useNavigate();
  const [showFilters] = useState(false);
  const [debouncedValue, setValue] = useDebounceValue("", 500);

  useEffect(() => {
    fetchFilteredClusters(debouncedValue)
  },[debouncedValue])

  const handleLogout = () => {
    navigate("/");
  };

  return (
    <Accordion
      type="single"
      collapsible
      value={showFilters ? "filters" : ""}
      className="w-full bg-sidebar text-sidebar-foreground z-10"
    >
      <AccordionItem
        value="filters"
        className="border-b border-sidebar-border relative"
      >
        <nav className="flex items-center justify-between h-16 px-4 py-2">
          <div className="flex items-center gap-2">
            {/* <Button
              variant="ghost"
              size="sm"
              onClick={() => setShowFilters((prev) => !prev)}
              className="gap-2"
            >
              <SlidersHorizontal className="w-4 h-4" />
              Advanced Filters
            </Button> */}
          </div>

          <div className="absolute left-1/2 transform -translate-x-1/2 w-88">
            <div className="relative">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                type="search"
                placeholder="Search"
                className="pl-10 text-sm"
                defaultValue=""
                onChange={(event) => setValue(event.target.value)}
              />
            </div>
          </div>

          <div className="flex flex-row-reverse items-center gap-4">
            <Button variant="secondary" size="icon" onClick={handleLogout}>
              <LogOut className="w-5 h-5" />
              <span className="sr-only">Log out</span>
            </Button>
            {/* <Avatar>
              <AvatarFallback>TD</AvatarFallback>
            </Avatar> */}
            <ThemeToggle />
            <MotionLink
              to={`https://${BACKEND_URL}/swagger/index.html`}
              target="_blank"
              rel="noreferrer"
              className={
                buttonVariants({ variant: "outline", size: "sm" }) +
                " border-2 font-semibold px-8 py-4 rounded-full"
              }
              whileHover={{ scale: 1.03 }}
              whileTap={{ scale: 0.95 }}
            >
              <BookMarked className="mr-2 h-5 w-5" />
              <span>Documentation</span>
            </MotionLink>
          </div>
        </nav>

        <AnimatePresence>
          {showFilters && (
            <AccordionContent
              forceMount
              className="absolute left-0 right-0 bg-sidebar px-4 py-4 border-t border-sidebar-border z-50 shadow-lg"
            >
              <motion.div
                initial={{ height: 0, opacity: 0 }}
                animate={{ height: "auto", opacity: 1 }}
                exit={{ height: 0, opacity: 0 }}
                transition={{ duration: 0.2, ease: "easeInOut" }}
              >
                <NavFiltersDrawer />
              </motion.div>
            </AccordionContent>
          )}
        </AnimatePresence>
      </AccordionItem>
    </Accordion>
  );
}
