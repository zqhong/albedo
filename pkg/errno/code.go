package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal Server Error"}
	RouteNotFound       = &Errno{Code: 10002, Message: "Route Not Found"}
	MethodNotFound      = &Errno{Code: 10002, Message: "Method Not Found"}
)
