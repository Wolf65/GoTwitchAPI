package gtapi

import (
	//"fmt"
	"errors"
	"encoding/json"
	"net/http"
)

func (s *Session) GetStreamByUser(userID string) (*GetStreamByUserOutput,error) { //fixmi
	var getStreamByUserOutput GetStreamByUserOutput
	req, err := http.NewRequest("GET", s.URL+"/streams/"+userID, nil)
	if err != nil {
		return &getStreamByUserOutput, err
	}
	req.Header.Set("Client-ID", s.ClientID)
	res, err := s.Client.Do(req)
	if err != nil {
		return &getStreamByUserOutput, err
	}
	if err := json.NewDecoder(res.Body).Decode(&getStreamByUserOutput); err != nil {
		return &getStreamByUserOutput, err
	}
	defer res.Body.Close()
	if getStreamByUserOutput.Stream == nil {
		return &getStreamByUserOutput, errors.New("stream offline")
	}
	return &getStreamByUserOutput, nil
}

func (s *Session) GetTopStream() (*GetTopStreamOutput, error) {
	var getTopStreamOutput GetTopStreamOutput 
	req, err := http.NewRequest("GET", s.URL+"/streams"+"?limit=5", nil)
	if err != nil {
		return &getTopStreamOutput, err
	}
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Set("Client-ID", s.ClientID)
	res, err := s.Client.Do(req)
	if err != nil {
		return &getTopStreamOutput, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&getTopStreamOutput)
	if err != nil {
		return &getTopStreamOutput, err	
	}
	return &getTopStreamOutput, nil
}