package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	// for security purpose you can use .env file and load it here by using the line below.
	//  godotenv.Load(".env")   Don't forget to import godotenv

	os.Setenv("SLACK_BOT_TOKEN", "// slack bot token")
	os.Setenv("CHANNEL_ID", "C04U2MUMARJ")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"//Your File"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
