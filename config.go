// Copyright 2015 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openweathermap

import (
	"errors"
	"os"
)

// Config will hold default settings to be passed into the
// "NewCurrent, NewForecast, etc}" functions.
type Config struct {
	Mode     string // user choice of JSON or XML
	Unit     string // measurement for results to be displayed.  F, C, or K
	Lang     string // should reference a key in the LangCodes map
	APIKey   string // API Key for connecting to the OWM
	Username string // Username for posting data
	Password string // Pasword for posting data
}

// CheckAPIKeyExists will see if an API key has been set.
func (c *Config) CheckAPIKeyExists() bool {
	return len(c.APIKey) > 1
}

// config will hold default settings to be passed into the OWM requests
var config = Config{}

// SetApiKey sets the API key
func SetApiKey(aNewApiKey string) {
	config.APIKey = aNewApiKey
}

// GetApiKey returns the API key. This value can be set using SetApiKey function
func GetApiKey() string {
	return config.APIKey
}

// SetAuthenticationCredentials sets username and password
func SetAuthenticationCredentials(aUsername, aPassword string) {
	config.Username = aUsername
	config.Password = aPassword
}

func GetUsername() string {
	return config.Username
}

func GetPassword() string {
	return config.Password
}

var ErrApiKeyNotFound = errors.New("OWM Api key not found!")

var ErrNoApiKeyInEnvVar = errors.New("OWM API key not found in environment variables!")

const envVarNameAPIKey = "OWM_API_KEY"
const envVarNameUsername = "OWM_USERNAME"
const envVarNamePassword = "OWM_PASSWORD"

// InitConfig parses environment variables to initialize configuration
func Init() error {
	var lApiKey string

	if lApiKey = os.Getenv(envVarNameAPIKey); !ValidateApiKey(lApiKey) {
		return ErrNoApiKeyInEnvVar
	}

	lUsername := os.Getenv(envVarNameUsername)
	lPassword := os.Getenv(envVarNamePassword)

	config.APIKey = lApiKey
	config.Username = lUsername
	config.Password = lPassword

	return nil
}

func ValidateApiKey(aApiKey string) bool {
	return len(aApiKey) == 32
}
