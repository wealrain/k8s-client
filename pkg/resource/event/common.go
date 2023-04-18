package event

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
)

var FailedReasonPartials = []string{"failed", "err", "exceeded", "invalid", "unhealthy",
	"mismatch", "insufficient", "conflict", "outof", "nil", "backoff"}

func FillEventType(event *corev1.Event) string {
	if len(event.Type) == 0 {

		for _, partial := range FailedReasonPartials {
			if event.Reason != "" && strings.Contains(strings.ToLower(event.Reason), partial) {
				event.Type = corev1.EventTypeWarning
				break
			}
		}

		if len(event.Type) == 0 {
			event.Type = corev1.EventTypeNormal

		}

	}

	return event.Type

}
