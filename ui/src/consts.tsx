export type Theme = "light" | "dark";

export type NodeInfoProps = {
  node: Node;
};

export type ClusterResourcesInfoProps = {
  clusterResources: ClusterResources;
};

export type NetworkInfoProps = {
  apiServerAdresses: string[];
  routerLBAddress: string[];
  dnsServers: string[] | undefined;
  segments: string[];
};

export type StorageClassesInfoProps = {
  storageClasses: StorageProvisioner[];
};

export type WebhookInfoProps = {
  mutatingWebhooks: string[];
  validatingWebhooks: string[];
};

export type ClusterResponse = {
  data: Cluster[];
  success: boolean;
  count: number;
}

export type Cluster = {
  clusterID: string;
  name?: string;
  kubernetesVersion: string;
  clusterDnsConfig: ClusterDnsConfig;
  clusterResources: ClusterResources;
  nodeInfo: Node[];
  routerLBAddress: string[];
  apiServerAddresses: string[];
  identityProviders: string[];
  storageProvisioners: StorageProvisioner[];
  mutatingWebhooks: string[];
  validatingWebhooks: string[];
  segments: string[];
};

export type ClusterDnsConfig = {
  searchDomains?: string[];
  servers?: string[];
};

type ClusterResourceKey = 'cpu' | 'memory' | 'storage' | 'gpu' | 'pods';

export type ClusterResources = {
  [key in ClusterResourceKey]: string;
}

export type Node = {
  name: string;
  internalIP: string;
  hostname: string;
  osImage: string;
  kubeletVersion: string;
};

export type StorageProvisioner = {
  name: string;
  provisioner: string;
};
