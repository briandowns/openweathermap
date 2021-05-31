// Copyright 2021 Marc-André Levasseur
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
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneCall(t *testing.T) {
	APIKey := os.Getenv("OWM_API_KEY")
	fmt.Printf("APIKey :%s\n", APIKey)
	oc, err := NewOneCall("C", "EN", APIKey, []ExcludeOption{ExcludeHourly, ExcludeAlerts, ExcludeMinutely})
	assert.NoError(t, err)
	assert.NotNil(t, oc)

	err = oc.PerformOneCall(45.508, -73.5878)
	assert.NoError(t, err)

}

func TestErrorExcludes(t *testing.T) {
	APIKey := os.Getenv("OWM_API_KEY")
	oc, err := NewOneCall("C", "EN", APIKey, []ExcludeOption{ExcludeHourly, ExcludeMinutely, 18})
	assert.Error(t, err)
	assert.True(t, errors.Is(err, errExcludesInvalid))
	assert.Nil(t, oc)
}
