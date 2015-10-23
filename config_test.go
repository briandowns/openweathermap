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
	"testing"
)

func TestApiKeyNotFoundNegativeTC(t *testing.T) {
	t.Parallel()

	c, err := NewCurrent("f", "ru")
	if err != nil {
		t.Error(err)
	}

	bFetchErr := c.CurrentByName("Stockholm")
	if bFetchErr != ErrApiKeyNotFound {
		t.Error(errors.New("CurrentByName did not return ApiKeyNotFound!"))
	}

	bFetchErr = c.CurrentByCoordinates(
		&Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
	)
	if bFetchErr != ErrApiKeyNotFound {
		t.Error(errors.New("CurrentByCoordinates did not return ApiKeyNotFound!"))
	}

	bFetchErr = c.CurrentByID(5344157)
	if bFetchErr != ErrApiKeyNotFound {
		t.Error(errors.New("CurrentByID did not return ApiKeyNotFound!"))
	}
}

func setupConfigTC() {
	Init()
}

func teardownConfigTC() {
	SetApiKey("")
	SetAuthenticationCredentials("", "")
}

// TestInitConfigRequest
func TestInitRequest(t *testing.T) {
	t.Parallel()

	setupConfigTC()
	c, err := NewCurrent("f", "en")
	if err != nil {
		t.Error(err)
	}

	if bFetchErr := c.CurrentByName("Berlin"); bFetchErr != nil {
		t.Error(bFetchErr)
	}
	teardownConfigTC()
}

// TestInitConfig
func TestInit(t *testing.T) {
	t.Parallel()

	lOrigAPIKey := os.Getenv(envVarNameAPIKey)
	lOrigUsername := os.Getenv(envVarNameUsername)
	lOrigPassword := os.Getenv(envVarNamePassword)

	os.Setenv(envVarNameAPIKey, "12345678901234567890123456789012")
	os.Setenv(envVarNameUsername, "Username")
	os.Setenv(envVarNamePassword, "Password")

	Init()

	if GetApiKey() != "12345678901234567890123456789012" {
		t.Error(errors.New("InitConfig failed parsing API key!"))
	}

	if GetUsername() != "Username" {
		t.Error(errors.New("InitConfig failed parsing Username!"))
	}

	if GetPassword() != "Password" {
		t.Error(errors.New("InitConfig failed parsing password!"))
	}

	os.Setenv(envVarNameAPIKey, lOrigAPIKey)
	os.Setenv(envVarNameUsername, lOrigUsername)
	os.Setenv(envVarNamePassword, lOrigPassword)
}

// TestInitConfigNegativeNoApiKey
func TestInitNegativeNoApiKey(t *testing.T) {
	t.Parallel()

	lOrigAPIKey := os.Getenv(envVarNameAPIKey)

	os.Setenv(envVarNameAPIKey, "")

	if Init() != ErrNoApiKeyInEnvVar {
		t.Error(errors.New("InitConfig failed parsing API key!"))
	}

	os.Setenv(envVarNameAPIKey, lOrigAPIKey)
}
