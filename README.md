# Pretty Fail - pf

Generic library agnostic HTTP errors made to be displayed as JSON.


## Usage

```go
// just an error
pf.BadRequest()
// Error: HTTP 400: Bad Request
// Output:
// {
//   "code": "bad_request",
//   "message": "Bad Request"
// }

// with msg
pf.BadRequest("validation fail")
// Error: HTTP 400: validation fail
// Output:
// {
//   "code": "bad_request",
//   "message": "Bad Request"
// }


// with format
pf.BadRequestf("validation fail: %v", err)
// Error: HTTP 400: validation fail: err string
// Output:
// {
//   "code": "bad_request",
//   "message": "Bad Request"
// }

// with format for the public
pf.BadRequestf("bad input: %s", input).Safe()
// Error: HTTP 400: bad input: x
// Output:
// {
//   "code": "bad_request",
//   "message": "bad input: x"
// }


// with detailed error for the public
validations := []string{"fail one", "fail two"}
pf.BadRequestDetail(validations).Safe()
// Error: HTTP 400: Bad Request:[fail one fail two]
// Output:
// {
//   "code": "bad_request",
//   "message": "Bad Request",
//   "detail": [
//     "fail one",
//     "fail two"
//   ]
// }
```
