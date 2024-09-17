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
	RequestTypeInt     RequestTypes = 0
	RequestTypeInt64   RequestTypes = 1
	RequestTypeUint64  RequestTypes = 2
	RequestTypeFloat64 RequestTypes = 3
)

type RandomRequest struct {
	RequestType RequestTypes
	Min, Max    int
	Return      chan *RandomResponse
}

type RandomResponse struct {
	Err   error `json:"-"`
	Value any   `json:"value"`
}
