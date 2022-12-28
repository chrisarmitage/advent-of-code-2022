package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"advent-2022/day01"

	"github.com/urfave/cli"
)

func  main()  {
	app := cli.NewApp()
	app.Name = "Advent of Code 2022 runner"
	app.Usage = "Run the various challenges"

	app.Commands = []cli.Command{
		{
			Name:        "run",
			HelpName:    "run",
			Action:      RunAction,
			ArgsUsage:   ` `,
			Usage:       `Runs a challenge`,
			Description: `Runs a challenge`,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:  "day",
					Usage: "Advent day to run",
				},
				&cli.BoolFlag{
					Name:  "sample",
					Usage: "`sample`",
				},
			},
		},
    }
  
    err := app.Run(os.Args)
    if err !=  nil  {
        log.Fatal(err)
    }
}

func RunAction(c *cli.Context) error {
	if !c.IsSet("day") {
		return errors.New("advent day not set")
	}

	result, err := calculateChallange(c.Int("day"), c.Bool("sample"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

	return nil
}

func calculateChallange(day int, sample bool) (string, error) {
	fmt.Printf("Running day %02d", day)
	if sample {
		fmt.Printf(" against the sample input")
	}
	fmt.Println()
	switch day {
	case 1:
		return day01.Run(sample)
	default:
		return "", fmt.Errorf("day %02d not available", day)
	}

	return "fakeResult", nil
}
