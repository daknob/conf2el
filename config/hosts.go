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

// Host contains a single host that conf2el will SSH into, the credentials to use, and the commands to run
type Host struct {
	Name        string   `yaml:"name"`
	Hostname    string   `yaml:"hostname"`
	Credentials string   `yaml:"credentials"` // Must contain the Name of a Credentials
	Commands    []string `yaml:"commands"`    // Can be a list of Name s of Command s
}

// Validate validates the SSH host and returns the
// first error that it found.
func (host *Host) Validate() error {
	if !regexp.MustCompile(`^[a-z0-9A-Z\-]{1,64}$`).MatchString(host.Name) {
		return fmt.Errorf("name must be a-z, A-Z, 0-9, and - and between 1 and 64 characters")
	}
	if !regexp.MustCompile(`^[a-z0-9A-Z\-]{1,64}$`).MatchString(host.Credentials) {
		return fmt.Errorf("credentials name must be a-z, A-Z, 0-9, and - and between 1 and 64 characters")
	}
	if len(host.Commands) < 1 {
		return fmt.Errorf("no commands specified for host")
	}
	for n, c := range host.Commands {
		if !regexp.MustCompile(`^[a-z0-9A-Z\-]{1,64}$`).MatchString(c) {
			return fmt.Errorf("command #%d name must be a-z, A-Z, 0-9, and - and between 1 and 64 characters", n+1)
		}
	}

	return nil
}
