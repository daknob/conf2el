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

func TestCommand_Validate(t *testing.T) {
	type fields struct {
		Name     string
		Shell    string
		Compress bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "All good",
			fields: fields{
				Name:  "-a-ValiD-n4m3-",
				Shell: "/bin/ls",
			},
			wantErr: false,
		},
		{
			name: "Empty Name",
			fields: fields{
				Shell: "/bin/cp /dev/sda /dev/sdb",
			},
			wantErr: true,
		},
		{
			name: "Empty Shell",
			fields: fields{
				Name: "this-should-fail",
			},
			wantErr: true,
		},
		{
			name: "Invalid Name",
			fields: fields{
				Name:  "this_is invalid",
				Shell: "/bin/rm /bin/rm",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &Command{
				Name:     tt.fields.Name,
				Shell:    tt.fields.Shell,
				Compress: tt.fields.Compress,
			}
			if err := cmd.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Command.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
