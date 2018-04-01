package main

import (
	"github.com/nlopes/slack"
    "os"
    "github.com/joho/godotenv"
    "log"
)

func Env_load() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    Env_load()

    // API Clientを作成する
	api := slack.New(os.Getenv("BOT_USER_TOKEN"))

    // WebSocketでSlack RTM APIに接続する
	rtm := api.NewRTM()

    // goroutineで並列化する
	go rtm.ManageConnection()

    // イベントを取得する
	for msg := range rtm.IncomingEvents {
		// 型swtichで型を比較する
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// MessageEventだったら応答する
			rtm.SendMessage(rtm.NewOutgoingMessage("えらい！", ev.Channel))
		}
	}
}
