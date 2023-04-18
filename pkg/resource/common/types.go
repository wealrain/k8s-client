package common

import (
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

var ListEverything = metav1.ListOptions{
	LabelSelector: labels.Everything().String(),
	FieldSelector: fields.Everything().String(),
}

func ToYaml(obj interface{}) string {
	data, err := yaml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(data)
}
