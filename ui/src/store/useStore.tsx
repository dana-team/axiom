import type { Cluster, Theme } from "@/consts";
import { create } from "zustand";
import { persist } from "zustand/middleware";
import { fetchFilteredClusters } from "@/ClusterService";

type State = {
  clusters: Cluster[];
  theme: Theme;
  loading: boolean;
  error: string | null;
  page: number;
  totalItems: number;
  setTheme: (theme: Theme) => void;
  setPage: (page: number) => void;
  fetchClusters: () => void;
};

const setDocumentTheme = (theme: Theme): void => {
  document.documentElement.classList.remove("light", "dark");
  document.documentElement.classList.add(theme);
};

export const useStore = create<State>()(
  persist(
    (set) => ({
      clusters: [],
      totalItems: 0,
      loading: false,
      error: null,
      page: 1,
      theme: "light",
      setTheme: (theme) => {
        set({ theme });
        setDocumentTheme(theme);
      },
      setPage: (page: number) => set({ page }),
      fetchClusters: async () => {
        await fetchFilteredClusters();
      },
    }),
    {
      name: "app-storage",
      partialize: (state) => ({ theme: state.theme }),
    }
  )
);
