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

import "fmt"

// Config contains a conf2el configuration file
type Config struct {
	// SSH Credentials available
	Creds []Credentials `yaml:"credentials"`

	// SSH Commands that can be run
	Commands []Command `yaml:"commands"`

	// Hosts to run the commands against
	Hosts []Host `yaml:"hosts"`
}

// Validate validates the conf2el configuration file and returns the
// first error that it found.
func (conf *Config) Validate() error {
	if len(conf.Creds) < 1 {
		return fmt.Errorf("no SSH credentials defined")
	}
	if len(conf.Commands) < 1 {
		return fmt.Errorf("no SSH commands defined")
	}
	if len(conf.Hosts) < 1 {
		return fmt.Errorf("no SSH hosts defined")
	}

	for n, cred := range conf.Creds {
		if err := cred.Validate(); err != nil {
			return fmt.Errorf("failed to validate credential #%d: %v", n+1, err)
		}
	}

	for n, cmd := range conf.Commands {
		if err := cmd.Validate(); err != nil {
			return fmt.Errorf("failed to validate command #%d: %v", n+1, err)
		}
	}

	for n, host := range conf.Hosts {
		if err := host.Validate(); err != nil {
			return fmt.Errorf("failed to validate host #%d: %v", n+1, err)
		}
	}

	return nil
}
