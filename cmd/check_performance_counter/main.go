// +build windows

package main

import (
	"github.com/ncr-devops-platform/nagiosfoundation/lib/app/nagiosfoundation"
)

func main() {
	nagiosfoundation.CheckExecutableVersion()
	nagiosfoundation.CheckPerformanceCounter()
}
