package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"elon": {
		AuthToken: "123ABC",
		Username:  "elon",
	},
	"aryan": {
		AuthToken: "456DEF",
		Username:  "aryan",
	},
	"dogy": {
		AuthToken: "789GHI",
		Username:  "dogy",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"elon": {
		Coins:    100,
		Username: "elon",
	},
	"aryan": {
		Coins:    200,
		Username: "aryan",
	},
	"dogy": {
		Coins:    300,
		Username: "dogy",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
