# Official Go API client for Fluidcoins crypto payment gateway

A suite of APIs to create, manage and accept cryptocurrency payments in Africa

## Installation
Put the package under your project folder and add the following in import:
```sh
$ go get -u -v github.com/fluidcoins/client-go
```

## Documentation For Authorization

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "sk_key from the dashboard here",
	Prefix: "Bearer",
})
r, err := client.Service.Operation(auth, args)
```

