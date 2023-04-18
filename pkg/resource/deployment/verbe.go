package deployment

import (
	"context"
	"errors"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func PauseDeployment(client kubernetes.Interface, namespace, name string) error {
	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if deployment.Spec.Paused {
		return errors.New("deployment already paused")
	}

	deployment.Spec.Paused = true
	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	return err
}

func ResumeDeployment(client kubernetes.Interface, namespace, name string) error {
	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if !deployment.Spec.Paused {
		return errors.New("deployment already resumed")
	}

	deployment.Spec.Paused = false
	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	return err
}

const RestartAnnotation = "kubectl.kubernetes.io/restartedAt"

func RestartDeployment(client kubernetes.Interface, namespace, name string) error {
	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if deployment.Spec.Template.ObjectMeta.Annotations == nil {
		deployment.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}

	deployment.Spec.Template.ObjectMeta.Annotations[RestartAnnotation] = metav1.Now().Format(time.RFC3339)
	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	return err
}
