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
)

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

var ApiKeyNotFound = errors.New("Api key not found!")
