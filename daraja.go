// Package daraja is a wrapper for Safaricom's Daraja API. It provides methods for making requests against the API and response wrappers to get response fields in a quick and easy way.
package daraja

import (
	"encoding/base64"
	"fmt"
	"time"
)

type Environment byte

const (
	SANBOX Environment = iota
	PRODUCTION
)

type mpesa struct {
	_latestToken       string
	_lastTokenTime     int64
	CertPath           string
	ConsumerKey        string
	ConsumerSecret     string
	GlobalShortcode    int32
	SecurityCredential string
	Environment        Environment
}

func Mpesa(key, secret string, env Environment) mpesa {
	return mpesa{
		ConsumerKey:    key,
		ConsumerSecret: secret,
		Environment:    env,
	}
}

func (m *mpesa) GetAuthToken() string {
	if (time.Now().Unix() - m._lastTokenTime) < 3599 {
		return m._latestToken
	}

	b := []byte(m.ConsumerKey + ":" + m.ConsumerSecret)
	encoded := base64.StdEncoding.EncodeToString(b)

	headers := map[string]string{
		"Authorization": "Basic " + encoded,
	}
	data, err := getRequest[authResponse](m, routes["oauth"], headers)
	if err != nil {
		panic(fmt.Sprintf("an error occurred while trying to get an auth token: %v", err))
	}

	// cache values
	m._latestToken = data.AccessToken
	m._lastTokenTime = time.Now().Unix()

	println("reached")
	fmt.Printf("data: %v\n", data)
	return data.AccessToken
}

func (m *mpesa) STKPush() *stkPush {
	return &stkPush{}
}
