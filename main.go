/*
#######
##                                   __   __
##        ____    _____ ____ _      / /  / /__  ______ _____
##       (_-< |/|/ / _ `/ _ `/ _   / _ \/ / _ \/ __/ // (_-<
##      /___/__,__/\_,_/\_, / (_) /_.__/_/\___/\__/\_,_/___/
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package main

import (
	"os"

	"github.com/arnumina/swag.blocus/cmd/blocus"
)

var version, builtAt string

func main() {
	if blocus.Run(version, builtAt) != nil {
		os.Exit(-1)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
