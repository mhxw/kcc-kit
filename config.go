package kcckit

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	Token         Token
	Mojito        Mojito
	Torches       Torches
	Chain         Chain
	Witnet        Witnet
	TorchesOralce TorchesOralce
	MojitoOracle  MojitoOracle
}

type Token struct {
	BTCAddress   string
	ETHAddress   string
	SKCSAddress  string
	WKCSAddress  string
	WOKTAddress  string
	OKBAddress   string
	USDTAddress  string
	USDCAddress  string
	MJTAddress   string
	CHEAddress   string
	APEAddress   string
	MULTIAddress string
	SAXAddress   string
	SANDAddress  string
	MANAAddress  string
	MLSAddress   string
	AAVEAddress  string
	CRVAddress   string
	UNIAddress   string
	CFXAddress   string
	LINKAddress  string
}

type Torches struct {
	TorchesTroller string
	CompoundLens   string
	TToken         TToken
}

type TToken struct {
	TKCS  string
	TUSDT string
	TUSDC string
	TBTC  string
	TETH  string
	TSKCS string
}

type Mojito struct {
	MojitoFactory    string
	Masterchef       string
	MasterchefV2     string
	MojitoRouter     string
	SwapMiningAddr   string
	MojitoOracle     string
	LP               Lp
	LpPid            LpPid
	Path             Path
	EIP712DomainHash string
	PermitTypeHash   string
}

type Lp struct {
	LpMJTUSDT  string
	LpMJTUSDC  string
	LpUSDTUSDC string
	LpWKCSMJT  string
	LpWKCSUSDC string
	LpWKCSUSDT string
}

type LpPid struct {
	LpMJTUSDTPid  int64
	LpMJTUSDCPid  int64
	LpUSDTUSDCPid int64
}

type Path struct {
	PathUsdcUsdt []common.Address
	PathUsdtUsdc []common.Address
	PathUsdtMjt  []common.Address
	PathMjtUsdt  []common.Address
	PathUsdtChe  []common.Address
	PathCheUsdt  []common.Address
}

type Chain struct {
	Network    string
	ChainId    int64
	TXUrl      string
	RPCUrl     string
	WSUrl      string
	Multicall  string
	Multicall2 string
}

type Witnet struct {
	WitnetPriceRouter string
	BtcUsdPairId      string
	EthUsdPairId      string
	KcsUsdtPairId     string
	MjtKcsPairId      string
	UsdtUsdPairId     string
	UsdcUsdPairId     string
}

type MojitoOracle struct {
	MojitoOracleAddr string
	MOProxyAddr      string
	BtcUsdtPairId    string
	EthUsdtPairId    string
	KcsUsdtPairId    string
	MjtUsdtPairId    string
	UsdcUsdtPairId   string
}

type TorchesOralce struct {
	BtcUsd  EACAndOCR
	EthUsd  EACAndOCR
	KcsUsd  EACAndOCR
	MjtUsd  EACAndOCR
	UsdtUsd EACAndOCR
	UsdcUsd EACAndOCR
}

type EACAndOCR struct {
	EAC string
	OCR string
}
