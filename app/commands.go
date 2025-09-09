package app

import (
	"flag"
	"fmt"
	"yanying-x/biz"
)

func RunCommand(command string) {
	switch command {
	case "create":
		var count = flag.Uint("count", 1, "数量")
		var domain = flag.String("domain", "person", "数量")
		flag.Parse()
		batchCreateX(*domain, count)

	}
}

func batchCreateX(domain string, count *uint) {
	fmt.Printf("创建【%s】%v个", domain, count)
	for _ = range *count {
		applyCreateX(domain)
	}

}

func applyCreateX(domain string) {
	switch domain {
	case "person":
		biz.NewPerson()
	}
}
