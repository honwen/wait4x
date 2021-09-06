// Copyright 2021 Mohammad Abdolirad
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCommandError(t *testing.T) {
	var commandError *CommandError
	err := NewCommandError("error message")

	assert.ErrorAs(t, err, &commandError)
	assert.Equal(t, ExitError, commandError.ExitCode)
	assert.Equal(t, "error message", commandError.Message)
}

func TestTimedOutError(t *testing.T) {
	var commandError *CommandError
	err := NewTimedOutError()

	assert.ErrorAs(t, err, &commandError)
	assert.Equal(t, ExitTimedOut, commandError.ExitCode)
	assert.Equal(t, TimedOutErrorMessage, commandError.Message)
}
