package main

import (
	"flag"
	"fmt"
	"crossword/internal/board"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Logging at a Debug level")
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Logging at an Info level")
	}

	board.Set(0, 0, 'q')
	let, err := board.Get(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Output: %c", let)
}
