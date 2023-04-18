package endpoint

import (
	corev1 "k8s.io/api/core/v1"
)

func GetEndpointAddresses(endpoint *corev1.Endpoints) []string {
	addresses := []string{}

	for _, subset := range endpoint.Subsets {
		for _, address := range subset.Addresses {
			addresses = append(addresses, address.IP)
		}
	}

	return addresses
}
