package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/CosmicPredator/chibi/types"
)

type AuthRequest struct {
	clientId string
	redirectUri string
}

func (a AuthRequest) GetAuthURL() string {
	return fmt.Sprintf(
		"https://anilist.co/api/v2/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code",
		a.clientId,
		a.redirectUri,
	)
}

func (a AuthRequest) Login(authCode string) error {
	if err := saveAccessToken(&a, authCode); err != nil {
		return err
	}

	return nil
}

func saveAccessToken(a *AuthRequest, authCode string) error {
	client := NewAnilistClient()

	body := map[string]string {
		"grant_type": "authorization_code",
		"client_id": a.clientId,
		"client_secret": "pYCQmRe7KMEFaWTIzWPMPsMqSJOWhyGsjj06BNrO",
		"redirect_uri": a.redirectUri,
		"code": authCode,
	}

	jsonString, err := json.Marshal(body)

	if err != nil {
		return err
	}

	response, err := client.HttpClient.Post(
		"https://anilist.co/api/v2/oauth/token",
		"application/json",
		bytes.NewBuffer(jsonString),
	)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	
	res, _ := io.ReadAll(response.Body)

	var tokenConfig types.TokenConfig
	err = json.Unmarshal(res, &tokenConfig)

	if err != nil {
		return err
	}

	err = tokenConfig.FlushToJsonFile()

	if err != nil {
		return err
	}
	return nil
}


func NewAuthRequest() *AuthRequest {
	return &AuthRequest{
		clientId: "4593",
		redirectUri: "https://anilist.co/api/v2/oauth/pin",
	}
}