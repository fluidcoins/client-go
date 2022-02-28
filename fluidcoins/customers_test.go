package fluidcoins

import (
	"fmt"

	"github.com/fluidcoins/client-go/util"
)

func randomCustomer() Customer {
	c := Customer{
		ID:            util.RandomID(),
		MerchantID:    util.RandomID(),
		FullName:      util.RandomName(),
		Reference:     fmt.Sprintf("CUS_%s", util.RandomID()),
		Email:         util.RandomEmail(),
		PhoneNumber:   util.RandomPhoneNumber(12),
		IsBlacklisted: false,
		Domain:        RandomDomain(),
	}
	return c
}
