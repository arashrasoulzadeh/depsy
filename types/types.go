package types

import (
	"arashrasoulzadeh/deepzy/structs"
	"log"
)

type Runner func(structs.ExecStruct)

func Run(e structs.ExecStruct) {
	switch e.Type {
	case "bash":
		runBash(e)
		log.Println()
	case "nginx":
		runNginx(e)
		log.Println()
	case "mysql":
		runMysql(e)
		log.Println()
	case "maria":
		runMaria(e)
		log.Println()
	case "folder":
		runFolder(e)
		log.Println()
	}

}
