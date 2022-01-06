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

func TestHost_Validate(t *testing.T) {
	type fields struct {
		Name        string
		Hostname    string
		Credentials string
		Commands    []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Invalid name",
			fields: fields{
				Name:        "not valid",
				Hostname:    "example.com",
				Credentials: "a",
				Commands:    []string{"b"},
			},
			wantErr: true,
		},
		{
			name: "Valid hostname",
			fields: fields{
				Name:        "all-good",
				Hostname:    "example.org",
				Credentials: "a",
				Commands:    []string{"b"},
			},
			wantErr: false,
		},
		{
			name: "Valid IPv6",
			fields: fields{
				Name:        "all-good-6",
				Hostname:    "2001:db8::5",
				Credentials: "a",
				Commands:    []string{"b"},
			},
			wantErr: false,
		},
		{
			name: "Valid IPv4",
			fields: fields{
				Name:        "all-good-4",
				Hostname:    "192.0.2.5",
				Credentials: "a",
				Commands:    []string{"b"},
			},
			wantErr: false,
		},
		{
			name: "No credentials",
			fields: fields{
				Name:     "should-fail",
				Hostname: "2001:db8::5",
				Commands: []string{"b"},
			},
			wantErr: true,
		},
		{
			name: "No commands",
			fields: fields{
				Name:        "should-fail",
				Hostname:    "2001:db8::5",
				Credentials: "secure",
			},
			wantErr: true,
		},
		{
			name: "Invalid credentials",
			fields: fields{
				Name:        "should-fail",
				Hostname:    "2001:db8::5",
				Credentials: "not valid",
			},
			wantErr: true,
		},
		{
			name: "Invalid commands",
			fields: fields{
				Name:     "should-fail",
				Hostname: "2001:db8::5",
				Commands: []string{"this-is-valid", "this is not"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host := &Host{
				Name:        tt.fields.Name,
				Hostname:    tt.fields.Hostname,
				Credentials: tt.fields.Credentials,
				Commands:    tt.fields.Commands,
			}
			if err := host.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Host.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
