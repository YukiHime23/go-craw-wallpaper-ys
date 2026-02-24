package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	gw "github.com/YukiHime23/game-wallpaper"
	"github.com/YukiHime23/game-wallpaper/arknight"
	"github.com/YukiHime23/game-wallpaper/azurlane"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	game := os.Args[1]

	if game == "help" || game == "-help" || game == "--help" {
		printHelp()
		return
	}

	pathP := flag.String("path", "", "Path to the directory where wallpapers should be saved")
	flag.CommandLine.Parse(os.Args[2:])

	switch game {
	case "arknight":
		var path string
		if *pathP == "" {
			path = "Arknight_Wallpaper"
		} else {
			path = *pathP
		}
		newPath, err := gw.CreateFolder(path)
		if err != nil {
			log.Fatalf("Failed to create folder for Arknights: %v", err)
		}
		fmt.Println("Downloading Arknights wallpapers...")
		arknight.DownloadArknightImages(newPath)

	case "endfield":
		fmt.Println("Downloading Arknights Endfield wallpapers...")
		arknight.DownloadEndfieldImages()

	case "aethergazer":
		fmt.Println("Aether Gazer download not yet integrated in main command.")

	case "azurlane":
		var path string
		if *pathP == "" {
			path = "AzurLane_Wallpaper"
		} else {
			path = *pathP
		}
		newPath, err := gw.CreateFolder(path)
		if err != nil {
			log.Fatalf("Failed to create folder for Azur Lane: %v", err)
		}

		fmt.Println("Downloading Azur Lane wallpapers...")
		azurlane.DownloadAzurLaneImages(newPath)

	case "majhongsoul":
		fmt.Println("Mahjong Soul download not yet integrated in main command.")

	default:
		fmt.Fprintf(os.Stderr, "Unknown game: %s\n\n", game)
		printHelp()
		os.Exit(1)
	}

	fmt.Println("Download completed!")
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage: %s <game> [OPTIONS]\n\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Game Wallpaper Downloader - Download wallpapers from various games")
	fmt.Fprintln(os.Stderr, "\nAvailable games:")
	fmt.Fprintln(os.Stderr, "  arknight      - Download Arknights wallpapers")
	fmt.Fprintln(os.Stderr, "  endfield      - Download Arknights Endfield wallpapers")
	fmt.Fprintln(os.Stderr, "  aethergazer   - Download Aether Gazer wallpapers")
	fmt.Fprintln(os.Stderr, "  azurlane      - Download Azur Lane wallpapers")
	fmt.Fprintln(os.Stderr, "  majhongsoul   - Download Mahjong Soul wallpapers")
	fmt.Fprintln(os.Stderr, "\nOptions:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nExamples:")
	fmt.Fprintf(os.Stderr, "  %s arknight\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s arknight -path ./wallpapers\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s endfield -path ~/Downloads\n", os.Args[0])
}
