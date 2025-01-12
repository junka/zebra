// Copyright 2017 zebra project
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

package linux

import (
	"fmt"

	"github.com/junka/zebra/pkg/fea"
)

func VrfAdd(vrfName string, index uint32) error {
	fmt.Println("Linux VrfAdd")
	return nil
}

func Init() {
	fea.VrfAdd = VrfAdd

	ForwardingIPv4Set(true)
	fmt.Println("Forwarding IPv4", ForwardingIPv4Get())
	ForwardingIPv6Set(true)
	fmt.Println("Forwarding IPv6", ForwardingIPv6Get())
}
