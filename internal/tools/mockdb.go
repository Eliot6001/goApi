package tools 

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"james": {
		username: "james",
		AuthToken: "aa_AA_bb_BB"
	},
	"jason": {
		username: "jason",
		AuthToken: "bb_BB_cc_CC"
	},
	"jack": {
		username: "jack",
		AuthToken: "bb_BB_cc_DD"
	},
}

var mockAuthorDetails = map[string]ArticleDetails{
	"nietzsche": {
		Author: "nietzsche",
		NumberOfArticles: 10 
	},
	"freud": {
		Author: "freud",
		NumberOfArticles: 11 
	},
	"fyodor": {
		Author: "fyodor",
		NumberOfArticles: 30 
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

func (d *mockDB) GetAuthorArticles(username string) *AuthorDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = ArticleDetails{}
	clientData, ok := mockAuthorDetails[Author]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
