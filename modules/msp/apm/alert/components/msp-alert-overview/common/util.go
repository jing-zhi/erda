// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	stdtime "time"
)

func GetInterval(startTimeMs, endTimeMs int64, minInterval stdtime.Duration, preferredPoints int64) string {
	interval := stdtime.Duration((endTimeMs - startTimeMs) / preferredPoints / 1e3 * 1e9)
	if interval < minInterval {
		interval = minInterval
	}
	return interval.String()
}

func ToInterface(value []int64) []interface{} {
	arr := make([]interface{}, 0)
	for _, v := range value {
		arr = append(arr, v)
	}
	return arr
}