package main

// MjY0ODI2ODI2MzUyODg1NzYz.G2PdSz.mcBpTZsekaerqnVfXhnz4ynsExwkD2tOSIW0is
import(
	"fmt"
	"github.com/fei-felicia-chen/myBot/config"
	"github.com/fei-felicia-chen/myBot/bot"	
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}