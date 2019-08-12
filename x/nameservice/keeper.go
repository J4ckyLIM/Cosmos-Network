package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/sdk-application-tutorial/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey  sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

/**
  * Function constructor that create a new instances of the nameservice keeper
  * Takes 3 parameters
  * coinKeeper, type: bank.Keeper
  * storekey, type:  sdk.Storekey
  * cdc, type: *codec.Codec
  * Return Keeper, type: Keeper
*/
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}


/**
  * Function to set a name for the Whois, encoding with amino to marshal
  * Takes 3 parameters
  * ctx, type: sdk.Context
  * name, type:  string
  * whois, type: Whois
*/
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois Whois){
	if whois.Owner.Empty(){
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}


/**
  * Function to get the Whois metadata from the name
  * Takes 2 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * Return whois or new whois
*/
func (k Keeper) GetWhois(ctx sdk.Context, name string) Whois {
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, name){
		return NewWhois()
	}
	bz := store.Get([]byte(name))
	var whois Whoisk
	k.cdc.MustUnmarshalBinaryBare(bz, &whois)
	return whois
}


/**
  * Function to delete a Whois metadata from the name
  * Takes 2 parameters
  * ctx, type: sdk.Context
  * name, type: string
*/
func (k Keeper) DeleteWhois(ctx sdk.Context, name string) {
	store := ctx.KVStorer(k.storeKey)
	store.Delete([]byte(name))
}


/**
  * Function to get the string that resolves to the name
  * Takes 2 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * Return value, type: string
*/
func (k Keeper) ResolveName(ctx sdk.Context, name string) string{
	return k.GetWhois(ctx, name).Value
}

/**
  * Function to set the value string that resolves to the name
  * Takes 3 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * value, type: string
*/
func (k keeper) SetName(ctx sdk.Context, name string, value, string){
	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

/**
  * Function to check if the name has already an Owner or not
  * Takes 2 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * Return Bool (true or false)
*/
func (k keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetWhois(ctx, name).Owner.Empty()
}

/**
  * Function to get the current Owner of the name
  * Takes 2 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * Return Owner, type: sdk.AccAddress
*/
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhois(ctx, name).Owner
}

/**
  * Function to set the current Owner of the name
  * Takes 3 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * owner, type: sdk.AccAddress
*/
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whois := k.GetWhois(ctx, name)
	whois.Owner = owner
	k.SetWhois(ctx, name, whois)
}

/**
  * Function to get the current price of the name
  * Takes 2 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * Return price, type: sdk.Coins
*/
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx, name).Price
}

/**
  * Function to set the current price of the name
  * Takes 3 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * price, type: sdk.Coins
*/
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}

/**
  * Function to check if the name is present in the store or not
  * Takes 2 parameters
  * ctx, type: sdk.Context
  * name, type: string
  * Return Bool (true or false)
*/
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

/**
  * Function to get an iterator over all names in which the keys are the name and the values are the whois
  * Takes 1 parameter
  * ctx, type: sdk.Context
  * Return sdk.Iterator
*/
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

