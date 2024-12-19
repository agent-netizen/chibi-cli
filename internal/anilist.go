package internal

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"

    "github.com/CosmicPredator/chibi/types"
)

type AniListClient struct {
    HttpClient *http.Client
    ApiUrl     string
}

func (a *AniListClient) ExecuteGraqhQL(query string, variables map[string]interface{}, target interface{}) error {
    accessToken := types.NewTokenConfig()
    if err := accessToken.ReadFromJsonFile(); err != nil {
        return &types.TokenNotFoundError{}
    }

    jsonBuffer, err := json.Marshal(map[string]interface{}{"query": query, "variables": variables})
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", a.ApiUrl, bytes.NewBuffer(jsonBuffer))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken.AccessToken))
    response, err := a.HttpClient.Do(req)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    responseBytes, err := io.ReadAll(response.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal(responseBytes, &target)
    if err != nil {
        return err
    }

    return nil
}

func NewAnilistClient() *AniListClient {
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    return &AniListClient{
        HttpClient: client,
        ApiUrl:     "https://graphql.anilist.co",
    }
}
