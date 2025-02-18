package cmd

import (
	"errors"
	"log"

	"github.com/ludviglundgren/qbittorrent-cli/internal/config"
	"github.com/ludviglundgren/qbittorrent-cli/pkg/qbittorrent"

	"github.com/spf13/cobra"
)

// RunAdd cmd to add torrents
func RunAdd() *cobra.Command {
	var (
		magnet        bool
		dry           bool
		paused        bool
		skipHashCheck bool
		savePath      string
		category      string
	)

	var command = &cobra.Command{
		Use:   "add",
		Short: "Add torrent",
		Long:  `Add new torrent to qBittorrent from file`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a torrent file as first argument")
			}

			return nil
		},
	}
	command.Flags().BoolVar(&magnet, "magnet", false, "Add magnet link instead of torrent file")
	command.Flags().BoolVar(&dry, "dry-run", false, "Run without doing anything")
	command.Flags().BoolVar(&paused, "paused", false, "Add torrent in paused state")
	command.Flags().BoolVar(&skipHashCheck, "skip-hash-check", false, "Skip hash check")
	command.Flags().StringVar(&savePath, "save-path", "", "Add torrent to the specified path")
	command.Flags().StringVar(&category, "category", "", "Add torrent to the specified category")

	command.Run = func(cmd *cobra.Command, args []string) {
		// args
		// first arg is path to torrent file
		filePath := args[0]

		if !dry {
			qbtSettings := qbittorrent.Settings{
				Hostname: config.Qbit.Host,
				Port:     config.Qbit.Port,
				Username: config.Qbit.Login,
				Password: config.Qbit.Password,
			}
			qb := qbittorrent.NewClient(qbtSettings)
			err := qb.Login()
			if err != nil {
				log.Fatalf("connection failed %v", err)
			}

			options := map[string]string{}
			if paused != false {
				options["paused"] = "true"
			}
			if skipHashCheck != false {
				options["skip_checking"] = "true"
			}
			if savePath != "" {
				options["savepath"] = savePath
				options["autoTMM"] = "false"
			}
			if category != "" {
				options["category"] = category
			}

			var res string
			if magnet {
				res, err = qb.AddTorrentFromMagnet(filePath, options)
			} else {
				res, err = qb.AddTorrentFromFile(filePath, options)
			}
			if err != nil {
				log.Fatalf("adding torrent failed: %v", err)
			}

			log.Printf("torrent successfully added: %v", res)
		} else {
			log.Println("dry-run: torrent successfully added!")
		}
	}

	return command
}
