package pod

import (
	"context"
	"io"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

var lineReadLimit int64 = 5000
var streamReadLimit int64 = 200

func GetLogsFile(client kubernetes.Interface, namespace, podName, containerName string) (string, error) {
	pod, err := client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})

	if err != nil {
		return "", err
	}
	// 未指定容器名时，默认为第一个容器
	if len(containerName) == 0 {
		containerName = pod.Spec.Containers[0].Name
	}

	// 获取日志流
	stream, err := client.CoreV1().RESTClient().Get().
		Namespace(namespace).
		Name(podName).
		Resource("pods").
		SubResource("log").
		VersionedParams(&corev1.PodLogOptions{
			Container:  containerName,
			Follow:     false,
			Previous:   false,
			Timestamps: true,
			TailLines:  &lineReadLimit,
		}, scheme.ParameterCodec).
		Stream(context.TODO())

	if err != nil {
		return "", err
	}

	defer stream.Close()
	// 读取日志流
	result, err := io.ReadAll(stream)
	if err != nil {
		return "", err
	}

	logs := string(result)

	return logs, nil

}

func GetLog(client kubernetes.Interface, namespace, podName, containerName string) (io.ReadCloser, error) {
	pod, err := client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}
	// 未指定容器名时，默认为第一个容器
	if len(containerName) == 0 {
		containerName = pod.Spec.Containers[0].Name
	}

	// 获取日志流
	stream, err := client.CoreV1().RESTClient().Get().
		Namespace(namespace).
		Name(podName).
		Resource("pods").
		SubResource("log").
		VersionedParams(&corev1.PodLogOptions{
			Container: containerName,
			Follow:    true,
			Previous:  false,
			TailLines: &streamReadLimit,
		}, scheme.ParameterCodec).
		Stream(context.TODO())

	if err != nil {
		return nil, err
	}

	return stream, nil

}
