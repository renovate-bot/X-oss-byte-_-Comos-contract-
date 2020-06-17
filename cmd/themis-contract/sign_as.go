package main

import (
	"os"

	contract "github.com/informalsystems/themis-contract/pkg/themis-contract"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func signAsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sign-as [signatory-id] [contract]",
		Short: "Sign a contract as a particular signatory",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			contractPath := defaultContractPath
			if len(args) > 1 {
				contractPath = args[1]
			}
			ctx, err := contract.InitContext(themisContractHome())
			if err != nil {
				log.Error().Err(err).Msg("Failed to initialize context")
			}
			c, err := contract.Load(contractPath, ctx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to load contract")
				os.Exit(1)
			}
			err = c.SignAs(themisContractHome(), args[0], ctx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to sign contract")
				os.Exit(1)
			}
		},
	}
}
