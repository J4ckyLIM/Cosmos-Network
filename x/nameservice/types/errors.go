package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeNameDoesNotExist sdk.CodeType = 101
)



/**
  * Function used if the name to delete isn't existing
  * Takes 1 parameter
  * codespace, type: sdk.CodespaceType
*/
func ErrNameDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNameDoesNotExist)
}