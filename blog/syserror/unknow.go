package syserror

type UnknowError struct {
	msg string
	reason error
}

func (c *UnknowError) Code() int{
	return 1000
}
func (c *UnknowError) Error() string{
	if len(c.msg) == 0{
		return "unkown error"
	}
	return c.msg
}
func (c *UnknowError) ErrorReson() error{
	return c.reason
}
