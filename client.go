package gtapi

import (
	"log"
	"net/http"
)
//NewClient is init struct conn
func NewClient(clientID string) *Session {
	if clientID == "" {
		log.Fatal("No ClientID provided")
		return nil
	}

	return &Session{
		Client:     &http.Client{},
		URL:        APITwitchURL,
		ClientID:   clientID,
	}
}