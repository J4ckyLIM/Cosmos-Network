package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

/**
 * Function constructor that create a new instances of the MsgSetName
 * Takes 3 parameters
 * name, type: string
 * value, type:  string
 * owner, type: sdk.AccAddress
 * Return MsgSetName, type: MsgSetName
 */
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

/**
 * NewMsgBuyName is the constructor function for MsgBuyName
 * Takes 3 parameters
 * name, type: string
 * bid, type:  sdk.Coins
 * buyer, type: sdk.AccAddress
 * Return MsgBuyName, type: MsgBuyName
 */
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Msg Interface

/**
 * Function that return the name of the module based on the route
 */
func (msg MsgSetName) Route() string { return RouterKey }

/**
 * Function that return the action based on the type
 */
func (msg MsgSetName) Type() string { return "set_name" }

/**
 * Function to validate the message
 * Return sdk.Error or nil
 */
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

/**
 * Function to encode the message for signing
 */
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

/**
 * Function to defines whose signature is required
 */
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
