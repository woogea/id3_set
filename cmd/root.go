package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/woogea/id3_set/id3"
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "id3_set",
	Short: "set id3 to mp3",
	Long:  "Set id3 tags from command line to mp3",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
	setTitleCmd.Flags().StringVarP(&filename, "filename", "f", "", "file")
	setTitleCmd.Flags().StringVarP(&title, "title", "t", "", "title")
	setTitleCmd.Flags().StringVarP(&artist, "artist", "a", "", "artist")
	setTitleCmd.Flags().StringVarP(&album, "album", "l", "", "album")
	setTitleCmd.MarkFlagRequired("filename")
	RootCmd.AddCommand(setTitleCmd)
}

var filename string
var title string
var album string
var artist string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("id3_set v0.01")
	},
}

var setTitleCmd = &cobra.Command{
	Use:   "setTitle",
	Short: "Get Title from mp3 idtag",
	Long:  "Get title from mp3 id3tag",
	Run: func(cmd *cobra.Command, args []string) {
		id3.SetTagElements(filename, title, artist, album)
	},
}
