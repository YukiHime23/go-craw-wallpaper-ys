package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"strings"

	downloadAL "github.com/YukiHime23/go-craw-al"
	"github.com/YukiHime23/go-craw-al/models"
)

var (
	ApiListWallpaperAzurLane    = "https://azurlane.yo-star.com/api/admin/special/public-list?page_index=1&page_num=1200&type=1"
	DomainLoadWallpaperAzurLane = "https://webusstatic.yo-star.com/"
)

func main() {
	var pathFile string
	pathP := flag.String("path", "", "Path to the directory where wallpapers should be saved.")
	flag.Parse()
	if pathP == nil || *pathP == "" {
		pathFile = "AzurLane_Wallpaper"
	} else {
		pathFile = *pathP
	}

	newPath, err := downloadAL.CreateFolder(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Get(ApiListWallpaperAzurLane)
	if err != nil {
		log.Fatal("call api error: ", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("read body error: ", err)
	}

	var resApi models.ResponseApi
	if err = json.Unmarshal(resBody, &resApi); err != nil {
		log.Fatal("json Unmarshal error: ", err)
	}

	db := downloadAL.GetSqliteDb()

	var idExist []int
	// get id exist
	ids, err := db.Query("SELECT id_wallpaper FROM azur_lane")
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("select id error: ", err)
	}
	defer ids.Close()

	var id int
	for ids.Next() {
		ids.Scan(&id)
		idExist = append(idExist, id)
	}

  listWallpp := make([]models.AzurLane, 0)
	for _, row := range resApi.Data.Rows {
		if downloadAL.IntInArray(idExist, row.ID) {
			continue
		}

		var al models.AzurLane
		al.Url = DomainLoadWallpaperAzurLane + row.Works
		al.FileName = strings.ReplaceAll(row.Title+" ("+row.Artist+").jpeg", "/", "-")
		al.IdWallpaper = row.ID

    listWallpp = append(listWallpp, al)
  }
  
  queue := downloadAL.StartCraw(listWallpp)
  
  for _, v := range listWallpp {  
		if err = downloadAL.DownloadFile(al.Url, al.FileName, newPath); err != nil {
			log.Fatal("download file error: ", err)
		}
		insertData := "INSERT INTO azur_lane VALUES (?, ?, ?)"
		_, err = db.Exec(insertData, al.IdWallpaper, al.FileName, al.Url)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer db.Close()
}
