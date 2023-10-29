package random

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fkondej/gocli/generate"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type ServerDataArgs struct {
	*RandomArgs
	ServerName  bool
	Avatar      bool
	AboutURL    bool
	CountryCode bool
	PrintJSON   bool
}

var serverDataArgs ServerDataArgs

var serverDataCmd = &cobra.Command{
	Use:   "server-data",
	Short: "Create ServerData wallet",
	Long:  `Create ServerData wallet`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunServerData(serverDataArgs); err != nil {
			serverDataArgs.Logger.Error("Error", zap.Error(err))
			os.Exit(1)
		}
	},
}

func init() {
	serverDataArgs.RandomArgs = &randomArgs
	RandomCmd.AddCommand(serverDataCmd)

	serverDataCmd.PersistentFlags().BoolVar(&serverDataArgs.ServerName, "server-name", false, "Random server name")
	serverDataCmd.PersistentFlags().BoolVar(&serverDataArgs.Avatar, "avatar", false, "Random avatar URL")
	serverDataCmd.PersistentFlags().BoolVar(&serverDataArgs.AboutURL, "about-url", false, "Random about URL (random wiki page)")
	serverDataCmd.PersistentFlags().BoolVar(&serverDataArgs.CountryCode, "country-code", false, "Random country code")

	serverDataCmd.PersistentFlags().BoolVar(&serverDataArgs.PrintJSON, "json", false, "Print result in json format")
}

func RunServerData(args ServerDataArgs) error {
	var (
		randomAll = true
		result    = map[string]string{}
		logger    = args.Logger
		err       error
		errMsg    = "failed to generate a random server data, %w"
	)
	if args.ServerName || args.Avatar || args.AboutURL || args.CountryCode {
		randomAll = false
	}

	if args.ServerName || randomAll {
		result["server-name"], err = generate.GenerateName()
		if err != nil {
			return fmt.Errorf(errMsg, err)
		}
		logger.Debug("generated random server name", zap.String("server-name", result["server-name"]))
	}

	if args.Avatar || randomAll {
		result["avatar"], err = generate.GenerateAvatarURL()
		if err != nil {
			return fmt.Errorf(errMsg, err)
		}
		logger.Debug("generated random avatar", zap.String("avatar", result["avatar"]))
	}

	if args.AboutURL || randomAll {
		result["about-url"], err = generate.GenerateRandomWikiURL()
		if err != nil {
			return fmt.Errorf(errMsg, err)
		}
		logger.Debug("generated random about URL (wiki url)", zap.String("about-url", result["about-url"]))
	}

	if args.CountryCode || randomAll {
		result["country-code"], err = generate.GenerateCountryCode()
		if err != nil {
			return fmt.Errorf(errMsg, err)
		}
		logger.Debug("generated random country code", zap.String("country-code", result["country-code"]))
	}

	if args.PrintJSON {
		resultJSON, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			return fmt.Errorf(errMsg, err)
		}

		fmt.Println(string(resultJSON))
	} else {
		for name, value := range result {
			fmt.Printf(" - %s: %s\n", name, value)
		}
	}

	return nil
}
