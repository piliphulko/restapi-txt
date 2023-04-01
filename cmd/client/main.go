package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type httpClient struct {
	http.Client
}

func main() {
	client := httpClient{http.Client{}}
	ctx, cf := context.WithTimeout(context.Background(), time.Second*5)
	defer cf()
	m := map[string]interface{}{"login": "arestovich123", "passwort": "aperon235"}
	if err := client.CreateAccount(ctx, m); err != nil {
		log.Fatal(err)
	}
	jwt, err := client.TakeJWT(ctx, m)
	if err != nil {
		log.Fatal(err)
	}
	date, err := client.GetData(ctx, jwt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}

func (client httpClient) CreateAccount(ctx context.Context, m map[string]interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8080/create/account", bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	request, err := client.Do(req)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(request.Body)
	var reqErr struct {
		Err string
	}

	if err := decoder.Decode(&reqErr); err != nil {
		return nil
	} else {
		return errors.New(reqErr.Err)
	}
}

func (client httpClient) TakeJWT(ctx context.Context, m map[string]interface{}) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8080/create/jwt", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	request, err := client.Do(req)
	if err != nil {
		return "", err
	}
	decoder := json.NewDecoder(request.Body)
	var jwt struct {
		JWT string
	}
	if err := decoder.Decode(&jwt); err != nil {
		return "", err
	}
	return jwt.JWT, nil
}

func (client httpClient) GetData(ctx context.Context, jwt string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/abc/date", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "BEARER "+jwt)
	request, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer request.Body.Close()
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
