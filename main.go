package main

import (
	"os"
	"fmt"
	"sync"
	"flag"
	"time"
	"github.com/bwmarrin/discordgo"
)

var delayFlag = flag.Int("d", 1, "Delay between requests (seconds)")

func Usage() {
	fmt.Printf("Usage: %s [OPTIONS] [TOKEN...]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var wg sync.WaitGroup
	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	} else {
		for _, token := range flag.Args() {
			wg.Add(1)
			time.Sleep(time.Duration(*delayFlag) * time.Second)
			go ValidateToken(&wg, token)
		}
	}
	wg.Wait()
}

func ValidateToken(wg *sync.WaitGroup, token string) {
	defer wg.Done()
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Printf("ERR %s (%s)\n", token, err)
		return
	}
	user, err := discord.User("@me")
	if err != nil {
		fmt.Printf("ERR %s (%s)\n", token, err)
		return
	}
	
	fmt.Printf("OK %s (%s)\n", token, user.String())

	discord.Logout()
}