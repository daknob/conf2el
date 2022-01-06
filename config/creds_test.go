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
	"testing"
)

func TestCredentials_Validate(t *testing.T) {
	type fields struct {
		Name       string
		Username   string
		SSHKeyPath string
		Password   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "No Name",
			fields: fields{
				Username: "root",
				Password: "toor",
			},
			wantErr: true,
		},
		{
			name: "No Username",
			fields: fields{
				Name:     "should-fail",
				Password: "toor",
			},
			wantErr: true,
		},
		{
			name: "No Password or SSH Key",
			fields: fields{
				Name:     "must-fail",
				Username: "root",
			},
			wantErr: true,
		},
		{
			name: "Just password",
			fields: fields{
				Name:     "must-pass",
				Username: "root",
				Password: "toor",
			},
			wantErr: false,
		},
		{
			name: "Both password and SSH key",
			fields: fields{
				Name:       "should-fail",
				Username:   "root",
				Password:   "toor",
				SSHKeyPath: "/tmp/test.txt",
			},
			wantErr: true,
		},
		{
			name: "Name contains underscore",
			fields: fields{
				Name:     "should_fail",
				Username: "root",
				Password: "toor",
			},
			wantErr: true,
		},
		{
			name: "Name is over 64 characters",
			fields: fields{
				Name:     "12345678901234567890123456789012345678901234567890123456789012345",
				Username: "root",
				Password: "toor",
			},
			wantErr: true,
		},
		{
			name: "Name is valid",
			fields: fields{
				Name:     "-test-1234-",
				Username: "root",
				Password: "toor",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cred := &Credentials{
				Name:       tt.fields.Name,
				Username:   tt.fields.Username,
				SSHKeyPath: tt.fields.SSHKeyPath,
				Password:   tt.fields.Password,
			}
			if err := cred.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Credentials.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
