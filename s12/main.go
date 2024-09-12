package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/rs/zerolog/log"
	usr "myapp/User"
	"os"
)

func main() {
	log.Info().Msg("Start...")
	logger := zerolog.New(os.Stdout)
	logger.Trace().Msg("Trace message")
	fmt.Println(usr.Sum(5, 9))
	fmt.Println(usr.Sum("5", "9"))
}
