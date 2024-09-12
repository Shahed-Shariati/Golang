package main

import (
	usr "exception/User"
	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout)
	logger.Trace().Msg("Trace message")
	result, _ := usr.Sum(1, 5)
	println(result)
	resultErr, err := usr.Sum(0, 5)
	println(resultErr, err.Error())
}
