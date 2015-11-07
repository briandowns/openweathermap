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
	"fmt"
	"os"
	"testing"
)

// TestRetrieveIcon will test the retrieval of icons from the API.
func TestRetrieveIcon(t *testing.T) {
	tmpDir := "/tmp"
	iconFiles := []string{"01d.png", "n7m.png"}

	for _, iconFile := range iconFiles {
		size, err := RetrieveIcon(tmpDir, iconFile)
		if err != nil {
			t.Error(err)
		}

		f, err := os.Stat(fmt.Sprintf("%s/%s", tmpDir, iconFile))
		if err != nil {
			t.Error(err)
		}

		if f.Size() != size {
			t.Error("Size of downloaded file does not match actual size of file")
		}

		err = os.Remove(fmt.Sprintf("%s/%s", tmpDir, iconFile))
		if err != nil {
			t.Error(err)
		}
	}
}
