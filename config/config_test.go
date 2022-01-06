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

import "testing"

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Creds    []Credentials
		Commands []Command
		Hosts    []Host
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Empty Credentials",
			fields: fields{
				Commands: []Command{{Name: "foo", Shell: "/bin/ls"}},
				Hosts:    []Host{{Name: "bar", Credentials: "none", Commands: []string{"none"}}},
			},
			wantErr: true,
		},
		{
			name: "Empty Commands",
			fields: fields{
				Creds: []Credentials{{Name: "foo", Username: "root", Password: "toor"}},
				Hosts: []Host{{Name: "bar", Credentials: "none", Commands: []string{"none"}}},
			},
			wantErr: true,
		},
		{
			name: "Empty Hosts",
			fields: fields{
				Creds:    []Credentials{{Name: "foo", Username: "root", Password: "toor"}},
				Commands: []Command{{Name: "bar", Shell: "/bin/cp /dev/sda /dev/sdb"}},
			},
			wantErr: true,
		},
		{
			name: "At least one",
			fields: fields{
				Creds:    []Credentials{{Name: "foo", Username: "root", Password: "toor"}},
				Commands: []Command{{Name: "test", Shell: "/bin/rm /bin/rm"}},
				Hosts:    []Host{{Name: "bar", Credentials: "none", Commands: []string{"none"}}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Creds:    tt.fields.Creds,
				Commands: tt.fields.Commands,
				Hosts:    tt.fields.Hosts,
			}
			if err := conf.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
