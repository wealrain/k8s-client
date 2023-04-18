package ingress

import netv1 "k8s.io/api/networking/v1"

func GetLoadBalancer(ingress *netv1.Ingress) []string {
	var loadBalancer []string
	for _, ingress := range ingress.Status.LoadBalancer.Ingress {
		if ingress.IP != "" {
			loadBalancer = append(loadBalancer, ingress.IP)
		}
		if ingress.Hostname != "" {
			loadBalancer = append(loadBalancer, ingress.Hostname)
		}
	}
	return loadBalancer
}
