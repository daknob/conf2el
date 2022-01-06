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
	"io/ioutil"
	"regexp"
)

// Credentials contains a single credential that can be used to SSH into a host
type Credentials struct {
	Name       string `yaml:"name"`
	Username   string `yaml:"username"`
	SSHKeyPath string `yaml:"keypath"`
	Password   string `yaml:"password"`
}

// Validate validates the SSH credentials and returns the
// first error that it found.
func (cred *Credentials) Validate() error {
	if cred.Name == "" {
		return fmt.Errorf("no name given")
	}
	if cred.Username == "" {
		return fmt.Errorf("no SSH username given")
	}
	if cred.SSHKeyPath == "" && cred.Password == "" {
		return fmt.Errorf("no SSH key path or password given")
	}
	if cred.SSHKeyPath != "" && cred.Password != "" {
		return fmt.Errorf("both SSH key and password given")
	}
	if cred.SSHKeyPath != "" {
		skb, err := ioutil.ReadFile(cred.SSHKeyPath)
		if err != nil {
			return fmt.Errorf("failed to read SSH key: %v", err)
		}
		if len(skb) == 0 {
			return fmt.Errorf("the SSH key is empty")
		}
	}
	if !regexp.MustCompile(`^[a-z0-9A-Z\-]{1,64}$`).MatchString(cred.Name) {
		return fmt.Errorf("name must be a-z, A-Z, 0-9, or - and between 1 and 64 characters")
	}

	return nil
}
