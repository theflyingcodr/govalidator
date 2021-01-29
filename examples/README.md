# Examples

Three examples showing some useage examples of the library.

* cli - validating flag input in a command line application
* http-inline - simple stdlib http server, showing request validation within a handler
* http-validate - sample of using an error handler combined with the validation.Validator interface

To test them, just run `go run main.go` targeting the relevant directory.

For the http servers, an example valid POST request is:

POST http://localhost:1234/
```json
{
    "name":"My Name",
    "dob":"2000-10-12T07:20:50.52Z",
    "count":1,
    "isEnabled":false
}
```

An invalid request:

POST http://localhost:1234/
```json
{
    "name":"My Name",
    "dob":"2000-10-12T07:20:50.52Z",
    "count":0,
    "isEnabled":true
}
```

The invalid request will then return this response with a 400 status code:

````json
{
    "errors": {
        "count": [
            "value 0 should be greater than 0"
        ],
        "isEnabled": [
            "value true does not evaluate to false"
        ]
    }
}
````
