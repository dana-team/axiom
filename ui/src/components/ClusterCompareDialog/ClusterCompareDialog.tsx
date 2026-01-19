import { useEffect, useState, type JSX } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import type { Cluster } from "@/consts";
import ClustersSelect from "@/components/ClusterCompareDialog/ClustersSelect";
import ComparisonDisplay from "@/components/ClusterCompareDialog/ComparisonDisplay";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Columns } from "lucide-react";

interface ClusterCompareDialogProps {
  clusters: Cluster[];
}

const MOCK_JSON = `{
  "differences": [
    { "key": "version", "clusterA": "1.22", "clusterB": "1.24" },
    { "key": "region", "clusterA": "us-west", "clusterB": "eu-central" },
    { "key": "nodeCount", "clusterA": 3, "clusterB": 5 },
    { "key": "networkPolicy", "clusterA": "enabled", "clusterB": "disabled" },
    { "key": "logging", "clusterA": "Stackdriver", "clusterB": "CloudWatch" },
    { "key": "autoscaling", "clusterA": true, "clusterB": false },
    { "key": "storageClass", "clusterA": "standard", "clusterB": "fast-ssd" },
    { "key": "kubeProxyMode", "clusterA": "iptables", "clusterB": "ipvs" },
    { "key": "ingressController", "clusterA": "nginx", "clusterB": "traefik" },
    { "key": "apiEndpoint", "clusterA": "https://api.cluster-a.io", "clusterB": "https://api.cluster-b.io" }
  ]
}`;

const ClusterCompareDialog = ({ clusters }: ClusterCompareDialogProps) => {
  const [searchInput, setSearchInput] = useState("");
  const [selectedClusters, setSelectedClusters] = useState<Cluster[]>([]);
  const [open, setOpen] = useState(false);
  const [view, setView] = useState<"select" | "result">("select");
  const [jsonResult, setJsonResult] = useState("");

  const handleAdd = (cluster: Cluster): void => {
    if (!selectedClusters.some((s) => s.clusterID === cluster.clusterID)) {
      setSelectedClusters((prev) => [...prev, cluster]);
    }
  };

  const handleRemove = (clusterToRemove: Cluster): void => {
    setSelectedClusters((prev) =>
      prev.filter((cluster) => cluster.clusterID !== clusterToRemove.clusterID)
    );
  };

  const handleAction = (type: "diff" | "common"): void => {
    console.log(`${type.toUpperCase()} selected clusters:`, selectedClusters);
    setJsonResult(MOCK_JSON);
    setView("result");
  };

  const handleBack = (): void => {
    setView("select");
  };

  const handleCloseDialog = (): void => {
    setView("select");
    setSearchInput("");
    setSelectedClusters([]);
    setJsonResult("");
  };

  const viewButtons: Record<typeof view, () => JSX.Element | null> = {
    select: () => (
      <div className="flex justify-end gap-2">
        <Button
          disabled={selectedClusters.length < 2}
          onClick={() => handleAction("diff")}
          className="bg-primary-dark/90 hover:bg-primary-dark/80 dark:bg-primary-lighter/70 dark:hover:bg-primary/60 
        text-white px-4 transition"
        >
          Find Differences
        </Button>
        <Button
          variant="secondary"
          disabled={selectedClusters.length < 2}
          onClick={() => handleAction("common")}
          className="bg-primary-dark/90 hover:bg-primary-dark/80 dark:bg-primary-lighter/70 dark:hover:bg-primary/60 
        text-white px-4 transition"
        >
          Find Commonalities
        </Button>
      </div>
    ),

    result: () => (
      <Button
        onClick={handleBack}
        className="bg-primary-dark/90 hover:bg-primary-dark/80 dark:bg-primary-lighter/70 dark:hover:bg-primary/60 
      text-white px-4 transition"
      >
        Back
      </Button>
    ),
  };

  useEffect(() => {
    handleCloseDialog();
  }, [open]);

  const isSelectView = view === "select";

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Tooltip>
          <TooltipTrigger asChild>
            <Button
              size="icon"
              className="bg-primary-dark/90 hover:bg-primary-dark/80 dark:bg-primary-lighter/70 dark:hover:bg-primary/60 
        text-white transition round"
              onClick={() => setOpen(true)}
            >
              <Columns className="h-5 w-5" />
              <span className="sr-only">Compare Clusters</span>
            </Button>
          </TooltipTrigger>
          <TooltipContent side="left" className="text-sm">
            Compare Clusters
          </TooltipContent>
        </Tooltip>
      </DialogTrigger>

      <DialogContent className="min-w-2xl bg-card">
        <DialogHeader className="mb-2">
          <DialogTitle>
            {isSelectView
              ? "Select 2 or more Clusters to Compare"
              : "Comparison Result"}
          </DialogTitle>
        </DialogHeader>

        {isSelectView ? (
          <ClustersSelect
            clusters={clusters}
            selectedClusters={selectedClusters}
            onAdd={handleAdd}
            onRemove={handleRemove}
            search={searchInput}
            setSearch={setSearchInput}
          />
        ) : (
          <ComparisonDisplay value={jsonResult} />
        )}

        <div className="mt-4 gap-2 flex justify-between">
          <div className="flex justify-start gap-2">
            <Button variant="outline" onClick={() => setOpen(false)}>
              Cancel
            </Button>
            {!isSelectView && viewButtons.result()}
          </div>

          {isSelectView && viewButtons.select()}
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default ClusterCompareDialog;
