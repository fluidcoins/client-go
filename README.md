# Official Go API client for Fluidcoins crypto payment gateway

A suite of APIs to create, manage and accept cryptocurrency payments in Africa

## Installation
Put the package under your project folder and add the following in import:
```sh
$ go get -u -v github.com/fluidcoins/client-go
```

## Documentation For Authorization

```golang

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/fluidcoins/client-go/fluidcoins"
)

func main() {

	c, err := fluidcoins.New(fluidcoins.SecretKey("sk_test_e93f2247654a4f2fb9021b1789b00a8c"))
	if err != nil {
		log.Fatal(err)
	}

	t, _, err := c.Transaction.GetByID(context.Background(), "TRANS_li5p6OWTfEn1Pc1hovawR")
	if err != nil {
	    log.Fatal(err)
	}


	json.NewEncoder(os.Stdout).Encode(t)

	trans, _, err := c.Transaction.List(context.Background(), &fluidcoins.TransactionListOptions{
		Perpage: 15,
		Status:  fluidcoins.Overpaid,
	})
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(os.Stdout).Encode(trans)
}
```

