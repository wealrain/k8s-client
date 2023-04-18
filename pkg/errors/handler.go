package errors

// 全局异常处理器
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type ClusterError struct {
	message string
}

func NewClusterError(message string) *ClusterError {
	return &ClusterError{message: message}
}

func (e *ClusterError) Error() string {
	return e.message
}
