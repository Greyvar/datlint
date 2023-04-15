package main

import (
	log "github.com/sirupsen/logrus"

	"errors"

	datlib "github.com/greyvar/datlib/common"
	"github.com/greyvar/datlib/entdefs"
	grids "github.com/greyvar/datlib/gridfiles"
	"path/filepath"
)

var errorCount = 0

func main() {
	log.Infof("datlint")
	log.Infof("datdir: %v", datlib.DatDir()) 

	lintEntdefs()
	gridfiles() 

	if errorCount == 0 {
		log.Infof("Finished. All good.")
	} else {
		log.Errorf("Errors!: %v", errorCount)
	}
}

func lintEntdefs() {
	matches, err := filepath.Glob(datlib.DatDir() + "/entdefs/*.yml")

	if err != nil {
		log.Errorf("%v", err)
	}

	for _, match := range matches {
		log.Infof("entdef: %v", match)

		entdef, err := entdefs.ReadEntdefFile(match)

		lintEntdefStateFrames(entdef)

		errcheck(err)
	}
}

func lintEntdefStateFrames(entdef *entdefs.EntityDefinition) {
	for name, state := range entdef.States {
		if len(state.Frames) == 0 {
			errcheck(errors.New(entdef.Title + " has no frames for state " + name))
		}
	}
}

func gridfiles() {
	matches, err := filepath.Glob(datlib.DatDir() + "/worlds/**/grids/*.grid")

	if err != nil {
		log.Errorf("%v", err)
	}

	for _, match := range matches {
		log.Infof("gridfile: %v", match)

		_, err := grids.ReadGrid(match)

		errcheck(err)
	}

}

func errcheck(err error) {
	if err != nil {
		log.Errorf("%v", err)
		errorCount++
	}
}
