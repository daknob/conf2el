// Copyright 2022 The conf2el Authors.
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

package config

import (
	"fmt"
	"regexp"
)

// Command contains a command to execute in the remote host
type Command struct {
	Name     string `yaml:"name"`
	Shell    string `yaml:"command"`  // String to run in SSH connection and capture output
	Compress bool   `yaml:"compress"` // Whether to gzip the output
}

// Validate validates the commands and returns the
// first error that it found.
func (cmd *Command) Validate() error {
	if !regexp.MustCompile(`^[a-z0-9A-Z\-]{1,64}$`).MatchString(cmd.Name) {
		return fmt.Errorf("name must be a-z, A-Z, 0-9, and - and between 1 and 64 characters")
	}
	if cmd.Shell == "" {
		return fmt.Errorf("no command given")
	}
	return nil
}
