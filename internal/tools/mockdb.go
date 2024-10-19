package tools 

import (
	"time"
	"fmt"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"james": {
		Username: "james",
		AuthToken: "aa_AA_bb_BB",
	},
	"jason": {
		Username: "jason",
		AuthToken: "bb_BB_cc_CC",
	},
	"jack": {
		Username: "jack",
		AuthToken: "bb_BB_cc_DD",
	},
}

var mockAuthorDetails = map[string]AuthorDetails{
	"nietzsche": {
		Author: "nietzsche",
		NumberOfArticles: 10, 
	},
	"freud": {
		Author: "freud",
		NumberOfArticles: 11, 
	},
	"fyodor": {
		Author: "fyodor",
		NumberOfArticles: 30, 
	},
}

func (d *mockDB) GetUserLoginDetails(Username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[Username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetAuthorArticles(Author string) *AuthorDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)
  fmt.Println(Author)
	var clientData = AuthorDetails{}
	clientData, ok := mockAuthorDetails[Author]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
