package capture

import (
	"OmiseChallenges/pkg/cipher"
	"fmt"
	"io"
	"log"
	"os"
)

func Sample() {
	csvfile, err := os.Open("fng.1000.csv.rot128")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	rot128Reader, err := cipher.NewRot128Reader(csvfile)

	done := Capture()
	io.Copy(os.Stdout, rot128Reader)
	capturedOutput, err := done()
	if err != nil {
		log.Fatalln("Couldn't capture output", err)
	}

	fmt.Printf("%q\n", err, capturedOutput)
}
