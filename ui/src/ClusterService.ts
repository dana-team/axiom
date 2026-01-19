import axios from "axios";
import { useStore } from "@/store/useStore"
import type { ClusterResponse } from "@/consts";

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;

export const fetchFilteredClusters = async (name: string = ""): Promise<void> => {
    try {
      useStore.setState({loading: true, error: null})
      const res = await axios.get<ClusterResponse>(`${BACKEND_URL}/clusters?name=${name}`);
      const data = res.data;
      if (data.success) {
        useStore.setState({clusters: data.data, totalItems: data.count || 0})
      } else {
        useStore.setState({ error: "Failed to fetch clusters", totalItems: 0 });
      }
    } catch (err: any) {
      if (axios.isAxiosError(err) && err.code === "ECONNABORTED") {
          useStore.setState({ error: "Request timed out", totalItems: 0});
      }
      useStore.setState({ error: err.message || "Failed to fetch clusters" , totalItems: 0});
      useStore.setState({loading: false})
    } finally {
      useStore.setState({loading: false})
    }
}