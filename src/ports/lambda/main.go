package main

import (
	"os"
	"strings"
)

func main() {
	handlerPathElements := strings.Split(os.Getenv("_HANDLER"), "/")

	if len(handlerPathElements) == 0 {
		panic("invalid handler path")
	}

	handlerName := handlerPathElements[len(handlerPathElements)-1]

	switch handlerName {
	case "trackEvent":
		break
	case "getTracking":
		break
	default:
		panic("unsupported handler")
	}
}
