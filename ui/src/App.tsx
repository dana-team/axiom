import { Routes, Route } from "react-router-dom";
import { Layout } from "@/Layout";
import Home from "@/routes/Home";
import Clusters from "@/routes/Clusters";
import NotFound from "@/routes/NotFound";
import { useEffect } from "react";
import ClusterInfo from "@/routes/ClusterInfo";
import type { Theme } from "./consts";
import { useStore } from "@/store/useStore";

function App() {
  const setTheme = useStore((state) => state.setTheme);

  useEffect(() => {

    const appStorage = localStorage.getItem("app-storage");
    if (appStorage) {
      const parsedAppStorage = JSON.parse(appStorage);
      const savedTheme = parsedAppStorage?.state?.theme as Theme;
      setTheme(savedTheme);
    } else {
      setTheme("light");
    }
    
  }, []);

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="*" element={<NotFound />} />

      <Route element={<Layout />}>
        <Route path="/clusters" element={<Clusters />} />
        <Route path="/clusters/:id" element={<ClusterInfo />} />
      </Route>
    </Routes>
  );
}

export default App;
