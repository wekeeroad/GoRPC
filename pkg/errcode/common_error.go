package errcode

var (
	Success          = NewError(0, "Success")
	Fail             = NewError(10000000, "Internal error")
	InvalidParams    = NewError(10000001, "Invalid params")
	Unauthorized     = NewError(10000002, "Authorized error")
	NotFound         = NewError(10000003, "Not found")
	Unknown          = NewError(10000004, "Unknown")
	DeadlineExceeded = NewError(10000005, "Deadline exceeded")
	AccessDenied     = NewError(10000006, "Access denied")
	LimitAccess      = NewError(10000007, "Limit access")
	MethodNotAllowed = NewError(10000008, "Not support the method")
)
