package main

import (
	"fmt"
	"os"
	"github.com/bwmarrin/discordgo"
)

func main() {
	Token := "TokenGoesHere"
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
