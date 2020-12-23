package errno

import "fmt"

type Errno struct {
	Code int
	Message string
}

//实现接口 Error
func (err *Errno) Error() string{
	return err.Message
}

type Err struct {
	Code int
	Message string
	Err error
}

// 返回一个err
func New(errno *Errno,err error)*Err{
	return &Err{Code: errno.Code,Message: errno.Message,Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

//实现接口 Error
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}


func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	// 用法
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}