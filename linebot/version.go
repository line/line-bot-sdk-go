// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linebot

import (
	"runtime/debug"
	"strings"
	"sync"
)

var (
	sdkVersion     = "8.unknown"
	sdkVersionOnce sync.Once
)

func GetVersion() string {
	// getting the version of line-bot-sdk-go should be done only once. Computing it repeatedly is meaningless.
	sdkVersionOnce.Do(func() {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}
		for _, dep := range info.Deps {
			if strings.Contains(dep.Path, "github.com/line/line-bot-sdk-go") {
				sdkVersion = strings.TrimPrefix(dep.Version, "v")
				break
			}
		}
	})
	return sdkVersion
}
