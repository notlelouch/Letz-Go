package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"elon": {
		AuthToken: "ELON@123",
		Username: "elon"
	}
	"aryan": {
		AuthToken: "ARYAN@123",
		Username: "aryan"
	}
	"dogy": {
		AuthToken: "DOGY@123",
		Username: "dogy"
	}
}

var mockCoinDetails = map[string]CoinDetails{
	"elon": {
		Coins: 100,
		Username: "elon"
	}
	"aryan": {
		Coins: 10000,
		Username: "aryan"
	}
	"dogy": {
		Coins: 1000,
		Username: "dogy"
	}
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 3)

	var clientData = LoginDetails{}
	clientData, ok = mockLoginDetails[username]

	if !ok {
		return null
	}

	return &clientData

}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 3)

	var clientData = CoinDetails{}
	clientData, ok = mockCoinDetails[username]

	if !ok {
		return null
	}

	return &clientData

}

func (d *mockDB) SetupDatabase() error {
	return nill
} 