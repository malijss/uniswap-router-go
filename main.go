package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	uniswap "uniswap/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//Token struct
type Token struct {
	symbol  string
	address string
	decimal int
}

//enter your private key of infura
const privateKey string = ""

func main() {

	//inputAmount in from token in app.uniswap.com
	const inputAmount = 1

	tokenMap := make(map[string]Token)
	tokenMap["WETH"] = Token{"WETH", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", 18}
	tokenMap["DEA"] = Token{"DEA", "0x80aB141F324C3d6F2b18b030f1C4E95d4d658778", 18}
	tokenMap["DEUS"] = Token{"DEUS", "0x3b62F3820e0B035cc4aD602dECe6d796BC325325", 18}
	tokenMap["USDC"] = Token{"USDC", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", 6}

	//routes you want to check it from contract
	routes := []Token{
		tokenMap["DEA"],
		tokenMap["DEUS"],
		tokenMap["WETH"],
		tokenMap["USDC"],
	}

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//uniswap contract address
	address := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	instance, err := uniswap.NewUniswap(address, client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("uniswap contract is loaded")

	//convert amountIn toWei
	amountIn := new(big.Int).Mul(big.NewInt(inputAmount), big.NewInt(int64(math.Pow10(routes[0].decimal))))

	b, err := instance.GetAmountsOut(&bind.CallOpts{}, amountIn, getPaths(routes))
	if err != nil {
		log.Fatal(err)
	}
	outputAmount := new(big.Float)
	outputAmount, _ = outputAmount.SetString(b[len(b)-1].String())

	//convert amountOut fromWei
	outputAmountDecimal := new(big.Float).Quo(outputAmount, big.NewFloat(math.Pow10(routes[len(routes)-1].decimal)))

	fmt.Println(outputAmountDecimal)
}

func getPaths(routes []Token) []common.Address {
	cAddres := []common.Address{}
	for _, token := range routes {
		cAddres = append(cAddres, common.HexToAddress(token.address))
	}
	return cAddres
}
