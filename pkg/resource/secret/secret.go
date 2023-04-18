package secret

import (
	"context"
	"k8s-client/pkg/errors"
	"k8s-client/pkg/resource/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Secret struct {
	Name string   `json:"name"`
	Age  int64    `json:"age"` // unix timestamp
	Keys []string `json:"keys"`
	Type string   `json:"type"`
}

func GetSecretList(client k8sClient.Interface, namespace string) ([]Secret, error) {
	channel := common.GetSecretListChannelWithOptions(client, namespace, metav1.ListOptions{})
	secrets := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var secretList []Secret

	for _, secret := range secrets.Items {
		secretInfo := toSecret(&secret)
		secretList = append(secretList, secretInfo)
	}

	return secretList, nil
}

func toSecret(secret *corev1.Secret) Secret {
	return Secret{
		Name: secret.Name,
		Age:  secret.CreationTimestamp.Unix(),
		Keys: getKeys(secret.Data),
		Type: string(secret.Type),
	}
}

func getKeys(data map[string][]byte) []string {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	return keys
}

func GetSecretData(client k8sClient.Interface, namespace, name string) map[string][]byte {
	secret, err := client.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	errors.HandleError(err)
	return secret.Data
}

func ToInterface(secretList []Secret) []interface{} {
	var result []interface{}
	for _, secret := range secretList {
		result = append(result, secret)
	}
	return result
}
