/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package threagile

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/threagile/threagile/pkg/common"
	"github.com/threagile/threagile/pkg/docs"
	"github.com/threagile/threagile/pkg/macros"
	"github.com/threagile/threagile/pkg/model"
)

func (what *Threagile) initMacros() *Threagile {
	what.rootCmd.AddCommand(&cobra.Command{
		Use:   common.ListModelMacrosCommand,
		Short: "Print model macros",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(docs.Logo + "\n\n" + fmt.Sprintf(docs.VersionText, what.buildTimestamp))
			cmd.Println("The following model macros are available (can be extended via custom model macros):")
			cmd.Println()
			/* TODO finish plugin stuff
			cmd.Println("Custom model macros:")
			for _, macros := range macros.ListCustomMacros() {
				details := macros.GetMacroDetails()
				cmd.Println(details.ID, "-->", details.Title)
			}
			cmd.Println()
			*/
			cmd.Println("----------------------")
			cmd.Println("Built-in model macros:")
			cmd.Println("----------------------")
			for _, macros := range macros.ListBuiltInMacros() {
				details := macros.GetMacroDetails()
				cmd.Println(details.ID, "-->", details.Title)
			}
			cmd.Println()
		},
	})

	what.rootCmd.AddCommand(&cobra.Command{
		Use:   common.ExplainModelMacrosCommand,
		Short: "Explain model macros",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(docs.Logo + "\n\n" + fmt.Sprintf(docs.VersionText, what.buildTimestamp))
			cmd.Println("Explanation for the model macros:")
			cmd.Println()
			/* TODO finish plugin stuff
			cmd.Println("Custom model macros:")
			for _, macros := range macros.ListCustomMacros() {
				details := macros.GetMacroDetails()
				cmd.Println(details.ID, "-->", details.Title)
			}
			cmd.Println()
			*/
			cmd.Println("----------------------")
			cmd.Println("Built-in model macros:")
			cmd.Println("----------------------")
			for _, macros := range macros.ListBuiltInMacros() {
				details := macros.GetMacroDetails()
				cmd.Printf("%v: %v\n", details.ID, details.Title)
			}

			cmd.Println()
		},
	})

	what.rootCmd.AddCommand(&cobra.Command{
		Use:   "execute-model-macro",
		Short: "Execute model macro",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := what.readConfig(cmd, what.buildTimestamp)
			progressReporter := common.DefaultProgressReporter{Verbose: cfg.Verbose}

			r, err := model.ReadAndAnalyzeModel(*cfg, progressReporter)
			if err != nil {
				return fmt.Errorf("unable to read and analyze model: %v", err)
			}

			macrosId := args[0]
			err = macros.ExecuteModelMacro(r.ModelInput, cfg.InputFile, r.ParsedModel, macrosId)
			if err != nil {
				return fmt.Errorf("unable to execute model macro: %v", err)
			}
			return nil
		},
	})

	return what
}
