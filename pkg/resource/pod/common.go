package pod

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// 采用kubectl的方式获取pod的状态信息
// 主要添加了容器状态的判断，包括初始化容器和容器
// 容器状态包括：Waiting、Running、Terminated
// 在Terminated状态下, 会判断是否是正常退出，如果是正常退出，则不会显示为失败
func getPodStatus(pod *corev1.Pod) string {
	restarts := 0
	readyContainers := 0

	// pod当前状态
	reason := string(pod.Status.Phase)
	if pod.Status.Reason != "" {
		reason = pod.Status.Reason
	}

	initializing := false

	// 初始化容器阶段状态
	for i := range pod.Status.InitContainerStatuses {
		container := pod.Status.InitContainerStatuses[i]
		restarts += int(container.RestartCount)

		switch {
		case container.State.Terminated != nil && container.State.Terminated.ExitCode == 0:
			continue
		case container.State.Terminated != nil:
			// 初始化容器退出状态，是否是正常退出，可以通过Reason来判断，如果Reason为空，则是正常退出
			// 如果Reason不为空，则是异常退出。如果是异常退出，则需要判断是通过Signal退出还是通过ExitCode退出
			if len(container.State.Terminated.Reason) > 0 {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Init: Signal %d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("Init: ExitCode %d", container.State.Terminated.ExitCode)
				}
			} else {
				reason = "Init:" + container.State.Terminated.Reason
			}
			initializing = true
		case container.State.Waiting != nil && len(container.State.Waiting.Reason) > 0 && container.State.Waiting.Reason != "PodInitializing":
			// 初始化容器等待状态
			reason = fmt.Sprintf("Init: %s", container.State.Waiting.Reason)
			initializing = true
		default:
			// 初始化容器运行状态，显示已经初始化的容器序号
			reason = fmt.Sprintf("Init: %d/%d", i, len(pod.Spec.InitContainers))
			initializing = true
		}
		break
	}

	// 容器阶段状态
	if !initializing {
		restarts = 0
		hasRunning := false
		for i := len(pod.Status.ContainerStatuses) - 1; i >= 0; i-- {
			container := pod.Status.ContainerStatuses[i]
			restarts += int(container.RestartCount)
			if container.State.Waiting != nil && container.State.Waiting.Reason != "" {
				reason = container.State.Waiting.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason != "" {
				reason = container.State.Terminated.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason == "" {
				// 容器非正常状态，判断是通过Signal退出还是通过ExitCode退出
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Signal: %d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("ExitCode: %d", container.State.Terminated.ExitCode)
				}
			} else if container.Ready && container.State.Running != nil {
				hasRunning = true
				readyContainers++
			}
		}

		if reason == "Completed" && hasRunning {
			if hasPodReadyCondition(pod.Status.Conditions) {
				reason = string(corev1.PodRunning)
			} else {
				reason = "NotReady"
			}
		}
	}

	if pod.DeletionTimestamp != nil && pod.Status.Reason == "NodeLost" {
		reason = string(corev1.PodUnknown)
	} else if pod.DeletionTimestamp != nil {
		reason = "Terminating"
	}

	if len(reason) == 0 {
		reason = string(corev1.PodUnknown)
	}

	return reason
}

func hasPodReadyCondition(conditions []corev1.PodCondition) bool {
	for _, condition := range conditions {
		if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

// 所有容器的重启次数
func getRestartCount(pod *corev1.Pod) int32 {
	var restartCount int32
	for _, container := range pod.Status.ContainerStatuses {
		restartCount += container.RestartCount
	}
	return restartCount
}

// 所有容器的镜像
func getImages(pod *corev1.Pod) []string {
	var images []string
	for _, container := range pod.Spec.Containers {
		images = append(images, container.Image)
	}
	return images
}
