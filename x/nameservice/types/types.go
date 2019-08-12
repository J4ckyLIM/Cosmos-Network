package types

import (
	"fmt"
	"strings"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Whois is a struct that contains all the metadata of a name
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

func (w Whois) String() string {
	return strings.Trimspace(fmt.Sprintf(`
	Owner: %s
	Value: %s
	Price: %s`, w.Owner, w.Value, w.Price))
}