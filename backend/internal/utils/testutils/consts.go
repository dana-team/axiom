package testutils

import v1alpha "github.com/dana-team/axiom/operator/api/v1alpha1"

const (
	TestDbName          = "test_db"
	TestDbCollection    = "clusterInfo"
	TestCluster1ID      = "cluster-1"
	TestCluster2ID      = "cluster-2"
	TestCluster1Version = "v1.24.0"
	TestCluster2Version = "v1.25.0"
)

var TestClusters = []interface{}{
	v1alpha.ClusterInfoStatus{
		ClusterID:          TestCluster1ID,
		ApiServerAddresses: []string{"172.201.39.229"},
		KubernetesVersion:  TestCluster1Version,
		IdentityProviders:  []string{"my_htpasswd_provider"},
		ClusterResources: v1alpha.ClusterResources{
			CPU:     "36",
			Memory:  "144280Mi",
			Pods:    "1500",
			Storage: "389754Mi",
			GPU:     "0",
		},
		NodeInfo: []v1alpha.NodeInfo{
			{
				Name:           "ocp-kubero-t9tcr-worker-westeurope1-bzcmj",
				InternalIP:     "10.0.128.6",
				Hostname:       "ocp-kubero-t9tcr-worker-westeurope1-bzcmj",
				OSImage:        "linux",
				KubeletVersion: "v1.28.15+1879980",
			},
			{
				Name:           "ocp-kubero-t9tcr-master-0",
				InternalIP:     "10.0.0.8",
				Hostname:       "ocp-kubero-t9tcr-master-0",
				OSImage:        "linux",
				KubeletVersion: "v1.28.15+1879980",
			},
		},
		StorageProvisioners: []v1alpha.StorageProvisioner{
			{
				Name:        "azurefile-csi",
				Provisioner: "file.csi.azure.com",
			},
			{
				Name:        "managed-csi",
				Provisioner: "disk.csi.azure.com",
			},
		},
		MutatingWebhooks: []string{
			"hns-mutating-webhook-configuration",
			"machine-api",
			"nmstate",
			"cert-manager-webhook",
		},
		ValidatingWebhooks: []string{
			"controlplanemachineset.machine.openshift.io",
			"hns-validating-webhook-configuration",
			"machine-api",
		},
	},
	v1alpha.ClusterInfoStatus{
		ClusterID:          TestCluster2ID,
		ApiServerAddresses: []string{"172.201.39.229"},
		KubernetesVersion:  TestCluster2Version,
		IdentityProviders:  []string{"my_htpasswd_provider"},
		ClusterResources: v1alpha.ClusterResources{
			CPU:     "36",
			Memory:  "144280Mi",
			Pods:    "1500",
			Storage: "389754Mi",
			GPU:     "0",
		},
		NodeInfo: []v1alpha.NodeInfo{
			{
				Name:           "ocp-kubero-t9tcr-worker-westeurope1-bzcmj",
				InternalIP:     "10.0.128.6",
				Hostname:       "ocp-kubero-t9tcr-worker-westeurope1-bzcmj",
				OSImage:        "linux",
				KubeletVersion: "v1.28.15+1879980",
			},
			{
				Name:           "ocp-kubero-t9tcr-master-0",
				InternalIP:     "10.0.0.8",
				Hostname:       "ocp-kubero-t9tcr-master-0",
				OSImage:        "linux",
				KubeletVersion: "v1.28.15+1879980",
			},
		},
		StorageProvisioners: []v1alpha.StorageProvisioner{
			{
				Name:        "azurefile-csi",
				Provisioner: "file.csi.azure.com",
			},
			{
				Name:        "managed-csi",
				Provisioner: "disk.csi.azure.com",
			},
		},
		MutatingWebhooks: []string{
			"hns-mutating-webhook-configuration",
			"machine-api",
			"nmstate",
			"cert-manager-webhook",
		},
		ValidatingWebhooks: []string{
			"controlplanemachineset.machine.openshift.io",
			"hns-validating-webhook-configuration",
			"machine-api",
		},
	},
}
