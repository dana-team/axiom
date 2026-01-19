import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Link } from "react-router-dom";
import type { Cluster, ClusterResources } from "@/consts";
import { capitalize } from "@/utils";


interface ClusterCardProps {
  cluster: Cluster;
}


const ClusterCard = ({ cluster }: ClusterCardProps) => {
  return (
    <Link to={`/clusters/${cluster.clusterID}`} className="block">
      <Card className="transition hover:ring-2 flex flex-col h-full">
        <CardHeader>
          <div className="flex items-center justify-between h-12">
            <CardTitle className="text-2xl">{cluster.name}</CardTitle>
            <div className="flex gap-2 text-muted-foreground">
              <Badge variant="outline">{cluster.kubernetesVersion}</Badge>
              {cluster.clusterDnsConfig.searchDomains ? (
                <Badge variant="outline">
                  {cluster.clusterDnsConfig.searchDomains[0]}
                </Badge>
              ) : null}
            </div>
          </div>
        </CardHeader>

        <CardContent className="flex flex-col gap-y-2">
          <p className="text-muted-foreground">Segments:</p>
          <div className="flex gap-3 flex-wrap">
            {cluster.segments && cluster.segments.length > 0 ? (
              cluster.segments.map((segment: string) => (
                <Badge key={segment} variant="outline">{segment}</Badge>
              ))
            ) : (
              <p className="text-sm text-muted-foreground italic">Not found</p>
            )}
          </div>
        </CardContent>
        
        <CardContent className="flex-grow">
          <Badge
            variant="outline"
            className="flex gap-2 flex-row justify-around w-full"
          > 
            {Object.entries(cluster.clusterResources as ClusterResources)
            .filter(([, value]) => value !== "0")
            .map(([key, value]) => (
              <div key={key} className="flex flex-col h-24 justify-center gap-1">
                <p className="text-lg text-muted-foreground">{capitalize(key)}:</p>
                <p className="text-sm">{value}</p>
              </div>
            ))}
          </Badge>
        </CardContent>
      </Card>
    </Link>
  );
};

export default ClusterCard;
