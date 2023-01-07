# go-requests an HTTP client for Go
_________

An HTTP client that is ready for production in Go for a lot of useful features while only using the standard library of the language.
## Installation
To use the library for HTTP calls, you must first import the corresponding HTTP package:
```bash
go get github.com/cploutarchou/go-requests
``` 
    
## Usage

### Client
1. Create a new client with the default configuration
```go
	builder := requests.NewBuilder()
	client := builder.Build()
```
2. Create a  new client with a custom configuration
```go
	builder := requests.NewBuilder()
	
	builder.Headers().
		SetContentType("application/json").
        SetAccept("application/json").
        SetUserAgent("go-requests")
        builder.SetRequestTimeout(10 * time.Second).
        SetResponseTimeout(10 * time.Second).
        SetMaxIdleConnections(10)
	client := builder.Build()
```
3. Create a new client with our own http.Client
```go
	customClient := &http.Client{
	Timeout: 10 * time.Second, 
	Transport: &http.Transport{
		MaxIdleConnsPerHost:   10, 
		ResponseHeaderTimeout: 10 * time.Second, 
		DialContext: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).DialContext,
	},
	}
	builder := requests.NewBuilder()
	builder.Headers().
		SetContentType("application/json").
		SetAccept("application/json").SetUserAgent("go-requests")
	builder.SetHTTPClient(customClient)
	client := builder.Build()
```
____________________
### Requests
1. GET request
```go
    response, err := client.Get("https://httpbin.org/get")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(body))
```