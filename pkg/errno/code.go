package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal Server Error"}
	NotFound            = &Errno{Code: 10002, Message: "404 Not Found"}
)
