/*
Copyright (c) 2017 GigaSpaces Technologies Ltd. All rights reserved

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudifyutils

import (
	"reflect"
	"testing"
)

type testpair struct {
	input  []string
	output [2][]string
}

var tests = []testpair{
	{[]string{"cfy-go", "status"}, [2][]string{{"cfy-go", "status"}, {}}},
	{[]string{"cfy-go", "status", "-user", "admin"}, [2][]string{{"cfy-go", "status"}, {"-user", "admin"}}},
}

func TestCliArgumentsList(t *testing.T) {
	for _, pair := range tests {
		args, options := CliArgumentsList(pair.input)

		if !reflect.DeepEqual(args, pair.output[0]) {
			t.Error(
				"For", pair.input,
				"expected", pair.output[0],
				"got", args,
			)
		}

		if !reflect.DeepEqual(options, pair.output[1]) {
			t.Error(
				"For", pair.input,
				"expected", pair.output[1],
				"got", options,
			)
		}

	}
}
