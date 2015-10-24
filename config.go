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
type ConfigData struct {
	Mode     string // user choice of JSON or XML
	Unit     string // measurement for results to be displayed.  F, C, or K
	Lang     string // should reference a key in the LangCodes map
	APIKey   string // API Key for connecting to the OWM
	Username string // Username for posting data
	Password string // Pasword for posting data
}

// CheckAPIKeyExists will see if an API key has been set.
func (c *ConfigData) CheckAPIKeyExists() bool {
	return len(c.APIKey) > 1
}

// SetApiKey sets the API key
func (me *ConfigData) SetApiKey(aNewApiKey string) {
	me.APIKey = aNewApiKey
}

// GetApiKey returns the API key. This value can be set using SetApiKey function
func (me *ConfigData) GetApiKey() string {
	return Config.APIKey
}

// SetAuthenticationCredentials sets username and password
func (me *ConfigData) SetAuthenticationCredentials(aUsername, aPassword string) {
	me.Username = aUsername
	me.Password = aPassword
}

func (me *ConfigData) GetUsername() string {
	return me.Username
}

func (me *ConfigData) GetPassword() string {
	return me.Password
}

// config will hold default settings to be passed into the OWM requests
var Config = ConfigData{}

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

	Config.APIKey = lApiKey
	Config.Username = lUsername
	Config.Password = lPassword

	return nil
}

func ValidateApiKey(aApiKey string) bool {
	return len(aApiKey) == 32
}
