import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import type { Cluster } from "@/consts";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";

interface ClustersListProps {
  clusters: Cluster[];
  onAction?: (clusterOrId: Cluster) => void;
  actionLabel?: string;
}

interface ClustersSelectProps {
  clusters: Cluster[];
  selectedClusters: Cluster[];
  onAdd: (cluster: Cluster) => void;
  onRemove: (cluster: Cluster) => void;
  search: string;
  setSearch: (val: string) => void;
}

const ClustersList = ({
  clusters,
  onAction,
  actionLabel,
}: ClustersListProps) => {
  return (
    <ScrollArea className="border rounded-md flex-1 overflow-auto w-1/2 flex flex-col">
      <ul className="p-2 space-y-1">
        {clusters.map((cluster) => (
          <li
            key={cluster.clusterID}
            className="flex items-center justify-between hover:bg-muted px-2 rounded border"
          >
            <span>{cluster.name}</span>
            {onAction && (
              <Button
                variant="ghost"
                size="sm"
                onClick={() => onAction(cluster)}
              >
                {actionLabel}
              </Button>
            )}
          </li>
        ))}
      </ul>
    </ScrollArea>
  );
};

const ClustersSelect = ({
  clusters,
  selectedClusters,
  onAdd,
  onRemove,
  search,
  setSearch,
}: ClustersSelectProps) => {
  const filteredClusters = clusters.filter(
    (cluster) =>
      cluster.name?.toLowerCase().includes(search.toLowerCase()) &&
      !selectedClusters.some(
        (selectedCluster) => selectedCluster.clusterID === cluster.clusterID
      )
  );

  return (
    <>
      <Input
        placeholder="Search clusters..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />

      <div className="flex gap-4 mt-2 h-116">
        <ClustersList
          clusters={filteredClusters}
          onAction={onAdd}
          actionLabel="+"
        />

        <Separator orientation="vertical" className="mx-1" />

        <ClustersList
          clusters={selectedClusters}
          onAction={onRemove}
          actionLabel="✕"
        />
      </div>
    </>
  );
};

export default ClustersSelect;
