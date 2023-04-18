package service

import (
	corev1 "k8s.io/api/core/v1"
)

func GetExternalEndpoints(service *corev1.Service) []string {
	extEndpoints := []string{}

	if service.Spec.Type == corev1.ServiceTypeLoadBalancer {
		for _, ingress := range service.Status.LoadBalancer.Ingress {
			if ingress.IP != "" {
				extEndpoints = append(extEndpoints, ingress.IP)
			}
			if ingress.Hostname != "" {
				extEndpoints = append(extEndpoints, ingress.Hostname)
			}
		}
	}

	for _, ip := range service.Spec.ExternalIPs {
		extEndpoints = append(extEndpoints, ip)
	}

	return extEndpoints
}

func GetPorts(service *corev1.Service) []string {
	ports := []string{}

	for _, port := range service.Spec.Ports {
		// 端口:协议
		ports = append(ports, port.Name+":"+string(port.Protocol))
	}

	return ports
}

 

