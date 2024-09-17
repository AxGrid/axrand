package internal

/*
 _____  __  __  _____  _____  ___  _____
/  _  \/  \/  \/   __\/  _  \/___\|  _  \
|  _  |>-    -<|  |_ ||  _  <|   ||  |  |
\__|__/\__/\__/\_____/\__|\_/\___/|_____/
zed (17.09.2024)
*/

type RequestTypes int8

var (
	RequestTypeInt64   RequestTypes = 0
	RequestTypeUint64  RequestTypes = 1
	RequestTypeFloat64 RequestTypes = 2
)

type RandomRequest struct {
	RequestType RequestTypes
	Return      chan *RandomResponse
}

type RandomResponse struct {
	Err error
	Out any
}
