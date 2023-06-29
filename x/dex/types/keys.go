package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/sei-protocol/goutils"
)

const (
	// ModuleName defines the module name
	ModuleName = "dex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dex"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func AddressKeyPrefix(contractAddr string) []byte {
	addr, _ := sdk.AccAddressFromBech32(contractAddr)
	return address.MustLengthPrefix(addr)
}

func ContractKeyPrefix(p string, contractAddr string) []byte {
	return goutils.ImmutableAppend([]byte(p), AddressKeyPrefix(contractAddr)...)
}

func DenomPrefix(denom string) []byte {
	length := uint16(len(denom))
	bz := make([]byte, 2)
	binary.BigEndian.PutUint16(bz, length)
	return goutils.ImmutableAppend(bz, []byte(denom)...)
}

func PairPrefix(priceDenom string, assetDenom string) []byte {
	return goutils.ImmutableAppend(DenomPrefix(priceDenom), DenomPrefix(assetDenom)...)
}

func OrderBookPrefix(long bool, contractAddr string, priceDenom string, assetDenom string) []byte {
	return goutils.ImmutableAppend(
		OrderBookContractPrefix(long, contractAddr),
		PairPrefix(priceDenom, assetDenom)...,
	)
}

func OrderBookContractPrefix(long bool, contractAddr string) []byte {
	var prefix []byte
	if long {
		prefix = KeyPrefix(LongBookKey)
	} else {
		prefix = KeyPrefix(ShortBookKey)
	}
	return goutils.ImmutableAppend(prefix, AddressKeyPrefix(contractAddr)...)
}

// `Price` constant + contract + price denom + asset denom
func PricePrefix(contractAddr string, priceDenom string, assetDenom string) []byte {
	return goutils.ImmutableAppend(
		PriceContractPrefix(contractAddr),
		PairPrefix(priceDenom, assetDenom)...,
	)
}

func PriceContractPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(PriceKey), AddressKeyPrefix(contractAddr)...)
}

func RegisteredPairPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(RegisteredPairKey), AddressKeyPrefix(contractAddr)...)
}

func OrderPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(OrderKey), AddressKeyPrefix(contractAddr)...)
}

func AssetListPrefix(assetDenom string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(AssetListKey), DenomPrefix(assetDenom)...)
}

func NextOrderIDPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(NextOrderIDKey), AddressKeyPrefix(contractAddr)...)
}

func MatchResultPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(MatchResultKey), AddressKeyPrefix(contractAddr)...)
}

func GetSettlementOrderIDPrefix(orderID uint64, account string) []byte {
	accountBytes := goutils.ImmutableAppend([]byte(account), []byte("|")...)
	orderIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(orderIDBytes, orderID)
	return goutils.ImmutableAppend(accountBytes, orderIDBytes...)
}

func GetSettlementKey(orderID uint64, account string, settlementID uint64) []byte {
	settlementIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(settlementIDBytes, settlementID)
	return goutils.ImmutableAppend(GetSettlementOrderIDPrefix(orderID, account), settlementIDBytes...)
}

func MemOrderPrefixForPair(contractAddr string, priceDenom string, assetDenom string) []byte {
	return goutils.ImmutableAppend(
		goutils.ImmutableAppend(KeyPrefix(MemOrderKey), AddressKeyPrefix(contractAddr)...),
		PairPrefix(priceDenom, assetDenom)...,
	)
}

func MemCancelPrefixForPair(contractAddr string, priceDenom string, assetDenom string) []byte {
	return goutils.ImmutableAppend(
		goutils.ImmutableAppend(KeyPrefix(MemCancelKey), AddressKeyPrefix(contractAddr)...),
		PairPrefix(priceDenom, assetDenom)...,
	)
}

func MemOrderPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(MemOrderKey), AddressKeyPrefix(contractAddr)...)
}

func MemCancelPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(MemCancelKey), AddressKeyPrefix(contractAddr)...)
}

func MemDepositPrefix(contractAddr string) []byte {
	return goutils.ImmutableAppend(KeyPrefix(MemDepositKey), AddressKeyPrefix(contractAddr)...)
}

func MemDepositSubprefix(creator, denom string) []byte {
	return goutils.ImmutableAppend([]byte(creator), DenomPrefix(denom)...)
}

func ContractKey(contractAddr string) []byte {
	return AddressKeyPrefix(contractAddr)
}

func OrderCountPrefix(contractAddr string, priceDenom string, assetDenom string, long bool) []byte {
	var prefix []byte
	if long {
		prefix = KeyPrefix(LongOrderCountKey)
	} else {
		prefix = KeyPrefix(ShortOrderCountKey)
	}
	return goutils.ImmutableAppend(prefix, goutils.ImmutableAppend(AddressKeyPrefix(contractAddr), PairPrefix(priceDenom, assetDenom)...)...)
}

const (
	LongBookKey = "LongBook-value-"

	ShortBookKey = "ShortBook-value-"

	OrderKey               = "order"
	AccountActiveOrdersKey = "account-active-orders"
	CancelKey              = "cancel"

	TwapKey             = "TWAP-"
	PriceKey            = "Price-"
	SettlementEntryKey  = "SettlementEntry-"
	NextSettlementIDKey = "NextSettlementID-"
	NextOrderIDKey      = "noid"
	RegisteredPairKey   = "rp"
	AssetListKey        = "AssetList-"
	MatchResultKey      = "MatchResult-"
	LongOrderCountKey   = "loc-"
	ShortOrderCountKey  = "soc-"

	MemOrderKey   = "MemOrder-"
	MemDepositKey = "MemDeposit-"
	MemCancelKey  = "MemCancel-"
)
