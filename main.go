/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"os"

	"github.com/apex/log"
	"github.com/sniperkit/snk.fork.myke/cmd"
)

func main() {
	if err := cmd.Exec(os.Args[1:]); err != nil {
		log.WithError(err).Error("error")
		os.Exit(1)
	}
}
