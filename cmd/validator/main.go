package validator

import (
	"flag"
	"fmt"
	"os"

	"code.ysitd.cloud/component/aviation/tower/validate"
	"sync"
)

var target string

func init() {
	flag.StringVar(&target, "target", "airline", "Type to test: flyer, airline (default: airline)")
	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		fmt.Printf("No file is provided")
		os.Exit(1)
		return
	}

	if target != "airline" && target != "flyer" {
		fmt.Printf("No unknown type")
		os.Exit(1)
		return
	}

	validation := func(path string, wg *sync.WaitGroup) {
		var err error
		switch target {
		case "airline":
			_, err = validate.ValidateAirline(path)
			break
		case "flyer":
			_, err = validate.ValidateFlyer(path)
		}
		if err != nil {
			fmt.Printf("Error in %s: %s", path, err.Error())
		}

		wg.Done()
	}

	var wait sync.WaitGroup

	for _, file := range flag.Args() {
		wait.Add(1)
		go validation(file, &wait)
	}

	wait.Wait()
}
