package types

import (
	"arashrasoulzadeh/deepzy/structs"
	"log"
	"time"

	"github.com/briandowns/spinner"
)

type Runner func(structs.ExecStruct)

func Run(e structs.ExecStruct) {
	s := spinner.New(spinner.CharSets[11], 50*time.Millisecond) // Build our new spinner
	s.Start()                                                   // Start the spinner

	switch e.Type {
	case "bash":
		runBash(e, s)
		log.Println()
	case "nginx":
		runNginx(e, s)
		log.Println()
	case "mysql":
		runMysql(e, s)
		log.Println()
	case "maria":
		runMaria(e, s)
		log.Println()
	case "folder":
		runFolder(e, s)
		log.Println()
	case "git":
		runGit(e, s)
		log.Println()
	case "pm2":
		runPm2(e, s)
		log.Println()
	case "npm":
		runNpm(e, s)
		log.Println()
	case "composer":
		runComposer(e, s)
		log.Println()
	case "service":
		runService(e, s)
		log.Println()
	case "httpd":
		runHttpd(e, s)
		log.Println()
	case "nuxt":
		runNuxt(e, s)
		log.Println()
	}
}
