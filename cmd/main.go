package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	gw "github.com/YukiHime23/game-wallpaper"
	"github.com/YukiHime23/game-wallpaper/aethergazer"
	"github.com/YukiHime23/game-wallpaper/arknight"
	"github.com/YukiHime23/game-wallpaper/azurlane"
	"github.com/YukiHime23/game-wallpaper/majhongsoul"
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
		var path string
		if *pathP == "" {
			path = "AetherGazer_Wallpaper"
		} else {
			path = *pathP
		}
		// Create subdirectories for different image types
		contentImgPath, err := gw.CreateFolder(filepath.Join(path, "contentImg"))
		if err != nil {
			log.Fatalf("Failed to create contentImg folder: %v", err)
		}
		mobileContentImgPath, err := gw.CreateFolder(filepath.Join(path, "mobileContentImg"))
		if err != nil {
			log.Fatalf("Failed to create mobileContentImg folder: %v", err)
		}

		fmt.Println("Downloading Aether Gazer wallpapers...")
		aethergazer.DownloadAetherGazerImages(contentImgPath, mobileContentImgPath)

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
		var path string
		if *pathP == "" {
			path = "MahjongSoul_Wallpaper"
		} else {
			path = *pathP
		}
		newPath, err := gw.CreateFolder(path)
		if err != nil {
			log.Fatalf("Failed to create folder for Mahjong Soul: %v", err)
		}

		fmt.Println("Downloading Mahjong Soul wallpapers...")
		majhongsoul.DownloadMahjongSoulImages(newPath)

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
