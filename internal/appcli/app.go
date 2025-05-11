package appcli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type Cli struct {
}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) Run() {
	app := &cli.App{
		Name:  "neuro-cli",
		Usage: "CLI для работы с нейросетью",
		Commands: []*cli.Command{
			{
				Name:  "train",
				Usage: "Обучить модель",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Usage:    "Путь к папке с изображениями",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "output",
						Usage: "Путь для сохранения модели",
						//Value: "model.json",
					},
				},
				Action: func(c *cli.Context) error {
					input := c.String("input")
					output := c.String("output")
					return trainAndSaveModel(input, output)
				},
			},
			{
				Name:  "predict",
				Usage: "Распознать изображение",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Usage:    "Путь к изображению",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "model",
						Usage: "Путь к файлу модели",
						Value: "model.json",
					},
				},
				Action: func(c *cli.Context) error {
					input := c.String("input")
					modelPath := c.String("model")
					return predictImage(modelPath, input)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
