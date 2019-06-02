package main

import (
	"OmiseChallenges/app/usecase"
	"log"
	"os"
)

func main() {
	csvfile, err := os.Open("../data/fng.1000.csv.rot128")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	useCase, err := usecase.BuildTumboonUseCase(csvfile)

	if err != nil {
		panic(err)
		return
	}

	err = useCase.DoTumboonFromList()

	if err != nil {
		panic(err)
		return
	}


}
