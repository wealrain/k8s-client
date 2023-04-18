package event

import (
	"k8s-client/pkg/resource/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Event struct {
	Type          string `json:"type"` // Event type, warning or normal
	Message       string `json:"message"`
	Source        string `json:"source"`
	Count         int32  `json:"count"`
	Age           int64  `json:"age"`
	LastSeenTime  string `json:"lastSeen"`
	FirstSeenTime string `json:"firstSeen"`
}

func GetEventList(client k8sClient.Interface, namespace string) ([]Event, error) {

	channel := common.GetEventListChannelWithOptions(client, namespace, metav1.ListOptions{})

	events := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var eventList []Event

	for _, event := range events.Items {
		eventInfo := toEvent(&event)
		eventList = append(eventList, eventInfo)
	}

	return eventList, nil
}

func toEvent(event *corev1.Event) Event {

	firstSeenTime, lastSeenTime := event.FirstTimestamp, event.LastTimestamp
	eventTime := metav1.NewTime(event.EventTime.Time)
	if firstSeenTime.IsZero() {
		firstSeenTime = eventTime
	}
	if lastSeenTime.IsZero() {
		lastSeenTime = eventTime
	}
	return Event{
		Type:          FillEventType(event),
		Message:       event.Message,
		Source:        event.Source.Component,
		Count:         event.Count,
		Age:           eventTime.Time.Unix(),
		LastSeenTime:  lastSeenTime.Format("2006-01-02 15:04:05"),
		FirstSeenTime: firstSeenTime.Format("2006-01-02 15:04:05"),
	}
}

func ToInterface(eventList []Event) []interface{} {
	var result []interface{}
	for _, event := range eventList {
		result = append(result, event)
	}
	return result
}
