# uniswap-router-go
load Uniswap V2: Router 2 contract in go

  #example
  
	tokenMap := make(map[string]Token)
	tokenMap["WETH"] = Token{"WETH", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", 18}
	tokenMap["DEA"] = Token{"DEA", "0x80aB141F324C3d6F2b18b030f1C4E95d4d658778", 18}
	tokenMap["DEUS"] = Token{"DEUS", "0x3b62F3820e0B035cc4aD602dECe6d796BC325325", 18}
	tokenMap["USDC"] = Token{"USDC", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", 6}

	routes := []Token{
		tokenMap["DEA"],
		tokenMap["DEUS"],
		tokenMap["WETH"],
		tokenMap["USDC"],
	}

	const inputAmount = 1
  
  #install dependencies and run

    go get -d ./..
  
    go run . 
  
#output 

    uniswap contract is loaded
    201.504389
  
  
