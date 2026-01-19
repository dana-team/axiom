package common

import (
	"context"
	"fmt"
	"net"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"

	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RetryOnConflictUpdate attempts to update a Kubernetes object with retry mechanism in case of conflicts.
// It retrieves the object first and then applies the update operation with built-in retry logic
func RetryOnConflictUpdate[T client.Object](ctx context.Context, obj T, k8sClient client.Client, name, namespace string, updateOperation func(T) error) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		if err := k8sClient.Get(ctx, client.ObjectKey{Name: name, Namespace: namespace}, obj); err != nil {
			return err
		}
		return updateOperation(obj)
	})
}

// FormatMiB converts a Kubernetes resource.Quantity value to a human-readable string in MiB units
func FormatMiB(q *resource.Quantity) string {
	bytes := q.Value()
	mib := float64(bytes) / (1024 * 1024)
	return fmt.Sprintf("%.0fMi", mib)
}

// FormatGiB converts a Kubernetes resource.Quantity value to a human-readable string in GiB units
func FormatGiB(q *resource.Quantity) string {
	bytes := q.Value()
	gib := float64(bytes) / (1024 * 1024 * 1024)
	return fmt.Sprintf("%.1fGi", gib)
}

func StripDomain(host string) string {
	parts := strings.Split(host, ".")
	if len(parts) > 2 {
		return strings.Join(parts[:len(parts)-2], ".")
	}
	return parts[0]
}

func CreateSegmentFromIPAndPrefix(ip string, prefix int) (string, error) {
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return "", fmt.Errorf("failed to parse IP address %s", ip)
	}

	parsedIp = parsedIp.To4()
	if parsedIp == nil {
		return "", fmt.Errorf("failed to convert IP address %s to IPv4", ip)
	}

	mask := net.CIDRMask(prefix, 32)
	network := parsedIp.Mask(mask)
	ipNet := net.IPNet{IP: network, Mask: mask}

	return ipNet.String(), nil
}

func FilterUniqueStrings(slice []string) []string {
	seen := make(map[string]struct{}, len(slice))
	out := make([]string, 0, len(slice))
	for _, s := range slice {
		// Skip internal/loopback ranges
		if strings.HasPrefix(s, "10.") || strings.HasPrefix(s, "127.") {
			continue
		}

		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			out = append(out, s)
		}
	}
	return out
}
