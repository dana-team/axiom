import { Badge } from "@/components/ui/badge";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Link, useParams } from "react-router-dom";
import { Clock } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { SiGrafana, SiRedhatopenshift } from "react-icons/si";
import { ClusterLink } from "@/components/ClusterLink";
import { useStore } from "@/store/useStore";
import { buildConsoleUrl, capitalize } from "@/utils";
import type { ClusterResourcesInfoProps, NetworkInfoProps, NodeInfoProps, StorageClassesInfoProps, WebhookInfoProps } from "@/consts";

const grafanaUrl = import.meta.env.VITE_GRAFANA_URL;

function WebhookInfo({
  mutatingWebhooks,
  validatingWebhooks,
}: WebhookInfoProps) {
  const sections = [
    { title: "Mutating Webhooks", items: mutatingWebhooks },
    { title: "Validating Webhooks", items: validatingWebhooks },
  ];

  return (
    <div className="flex flex-row gap-4 justify-around">
      {sections.map((section) => (
        <Card
          key={section.title}
          className="p-4 transition hover:ring-2 w-1/2"
        >
          <CardTitle className="text-xl p-2">{section.title}</CardTitle>
          <CardContent>
          {section.items ? 
            <ScrollArea className="max-h-48 overflow-y-auto pr-2">
              {section.items.map((webhook, index) => (
                <div key={index} className="flex flex-col">
                  <div className="flex flex-row gap-2">
                    <p className="text-md text-muted-foreground">Name: </p>
                    <p>{webhook}</p>
                  </div>
                  <Separator className="my-2" />
                </div>
              ))}
            </ScrollArea> 
          : <>No {section.title} found on cluster</> }
          </CardContent>
        </Card>
      ))}
    </div>
  );
}

function NetworkInfo({ apiServerAdresses, routerLBAddress, dnsServers, segments }: NetworkInfoProps) {
  const sections = [
    { title: "API Servers:", items: apiServerAdresses },
    { title: "Router LBs:", items: routerLBAddress },
    { title: "DNS Servers:", items: dnsServers },
    { title: "Segments:", items: segments }
  ];

  return (
    <div className="grid grid-cols-2 gap-4">
      {sections.map((section) => (
        section.items && section.items.length > 0 ? (
          <Card
            key={section.title}
            className="flex flex-col gap-3 p-4 transition hover:ring-2"
          >
            <p className="text-lg font-medium text-muted-foreground">
              {section.title}
            </p>
            <ScrollArea className="h-40 flex items-center justify-center">
              <div className="space-y-3 pr-4">
                {section.items.map((address) => (
                  <p key={address} className="text-lg font-medium">{address}</p>
                ))}
              </div>
            </ScrollArea>
          </Card>
        ) : null
      ))}
    </div>
  );
}

function StorageClassesInfo({
  storageClasses,
}: StorageClassesInfoProps) {
  return (
    <>
      {storageClasses.map((sc) => (
        <Card className="transition hover:ring-2 flex flex-col gap-4">
          <CardContent>
            <div className="flex flex-row gap-2">
              <p className="text-md text-muted-foreground">Name: </p>
              <p>{sc.name}</p>
            </div>
            <div className="flex flex-row gap-2">
              <p className="text-md text-muted-foreground">Type: </p>
              <p>{sc.provisioner}</p>
            </div>
          </CardContent>
        </Card>
      ))}
    </>
  );
}

function NodeInfo({ node }: NodeInfoProps) {
  return (
    <Card className="transition hover:ring-2 flex flex-col gap-4">
      <CardTitle className="pl-4">{node.hostname}</CardTitle>
      <CardContent>
        <p className="text-sm">IP: {node.internalIP}</p>
        <p className="text-sm">Kubelet Version: {node.kubeletVersion}</p>
        <p className="text-sm">OS: {node.osImage}</p>
      </CardContent>
    </Card>
  );
}

function ClusterResourcesInfo({
  clusterResources,
}: ClusterResourcesInfoProps) {
  const resourceEntries = Object.entries(clusterResources).filter(
    ([, value]) => value !== "0"
  );

  return (
    <div className="grid grid-cols-2 gap-3">
      {resourceEntries.map(([key, value]) => (
        <Card
          key={key}
          className="flex flex-col gap-3 p-4 transition hover:ring-2 justify-center items-center min-h-24"
        >
          <p className="text-lg font-medium text-muted-foreground text-center">
            {capitalize(key)}
          </p>
          <p className="text-2xl font-bold text-center">{value}</p>
        </Card>
      ))}
    </div>
  );
}

export const ClusterPanel = () => {
  const { clusters } = useStore();
  const { id: clusterId } = useParams<{ id: string }>();
  const cluster = clusters.find((c) => c.clusterID === clusterId);
  const timestamp = new Date().toLocaleString();

  const clusterDetails = [
    { label: "Version", value: cluster?.kubernetesVersion },
    {
      label: "Network",
      value: cluster?.clusterDnsConfig?.searchDomains
        ? cluster.clusterDnsConfig.searchDomains[0]
        : "",
    },
  ];

  const linkItems = [
    {
      url:
        cluster?.clusterDnsConfig.searchDomains && cluster?.name
          ? buildConsoleUrl(cluster?.name)
          : "",
      label: "Link to OpenShift",
      colorClass: "text-red-600",
      Icon: SiRedhatopenshift,
    },
    {
      url: grafanaUrl,
      label: "Link to Grafana",
      colorClass: "text-orange-500",
      Icon: SiGrafana,
    },
  ];

  return (
    <div className="w-1/4 p-6 space-y-6 border-r border-border bg-muted/30">
      <div className="flex items-center justify-between">
        <Link
          to="/clusters"
          className="gap-1 text-sm text-muted-foreground hover:text-foreground"
        >
          ← Back
        </Link>
      </div>

      <div>
        <h1 className="text-2xl font-bold text-foreground">{cluster?.name}</h1>
        <p className="text-sm text-muted-foreground mt-1 flex items-center gap-1">
          <Clock className="w-4 h-4 text-muted-foreground" />
          Last updated: {timestamp}
        </p>
      </div>

      <Separator />

      <div className="space-y-2 pt-2">
        {clusterDetails.map(({ label, value }) => (
          <div key={label} className="flex items-center justify-between">
            <span className="text-muted-foreground">{label}:</span>
            <Badge variant="outline">{value}</Badge>
          </div>
        ))}
      </div>

      <Separator />

      <div className="flex gap-6 pt-1">
        {linkItems.map((item, index) => (
          <ClusterLink
            key={index}
            url={item.url}
            label={item.label}
            colorClass={item.colorClass}
            Icon={item.Icon}
          />
        ))}
      </div>
    </div>
  );
};

type CustomWidgetProps = {
  id: number;
  title?: string;
  children: React.ReactNode;
};

const CustomWidget = ({ id, title, children }: CustomWidgetProps) => (
  <Card key={id} className="mb-4">
    <CardHeader>
      <CardTitle className="text-lg">{title}</CardTitle>
    </CardHeader>
    <CardContent>{children}</CardContent>
  </Card>
);

export default function ClustersPage() {
  const { clusters } = useStore();
  const { id: clusterId } = useParams<{ id: string }>();
  const cluster = clusters.find((c) => c.clusterID === clusterId);

  return (
    <div className="flex border border-border bg-background shadow-sm h-full">
      <ClusterPanel />
      {cluster ? (
        <div className="w-3/4 p-6 bg-gradient-to-br dark:from-primary/30 via-primary-dark/10 dark:to-primary-light/20 overflow-auto">
          <ScrollArea>
            <CustomWidget title="Nodes" id={1} key={1}>
              <ScrollArea className="max-h-80 overflow-y-auto">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4 m-4">
                  {cluster.nodeInfo.map((n) => (
                    <NodeInfo node={n} />
                  ))}
                </div>
              </ScrollArea>
            </CustomWidget>
            
            <div className="basis-2/3">
              <CustomWidget title="Network Info" id={3} key={3}>
                <NetworkInfo
                  dnsServers={cluster.clusterDnsConfig.servers}
                  apiServerAdresses={cluster.apiServerAddresses}
                  routerLBAddress={cluster.routerLBAddress}
                  segments={cluster.segments}
                />
              </CustomWidget>
            </div>
            <div className="basis-1/3 shrink-0">
              <CustomWidget title="Cluster Resources" id={2} key={2}>
                <ClusterResourcesInfo
                  clusterResources={cluster.clusterResources}
                />
              </CustomWidget>
            </div>
            <CustomWidget title="Storage Classes" id={4} key={4}>
              {cluster.storageProvisioners ? 
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <StorageClassesInfo 
                  storageClasses={cluster.storageProvisioners}
                /> 
              </div>:
              <p className="text-lg">No storage provisioners configured for cluster {cluster.name}</p>
            }
            </CustomWidget>
            <CustomWidget title="Webhooks" id={5} key={5}>
              <div>
                <WebhookInfo
                  mutatingWebhooks={cluster.mutatingWebhooks}
                  validatingWebhooks={cluster.validatingWebhooks}
                />
              </div>
            </CustomWidget>
          </ScrollArea>
        </div>
      ) : (
        <></>
      )}
    </div>
  );
}
