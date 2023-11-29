package error_utils

type CustomErr struct {
	Code    int
	Message string
	Detail  string
}

func (slf *CustomErr) Error() string {
	return slf.Message
}
