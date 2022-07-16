package main

import (
	"log"

	"github.com/macduyhai/golang-design-pattern/Behavioural-pattern/observer"
)

func main() {
	var ytChannel observer.Subject = &observer.YoutubeChannel{}
	ui1 := observer.NewUserInterface("Hai")
	ui2 := observer.NewUserInterface("Simon")
	ytChannel.RegisterObserver(ui1)
	ytChannel.RegisterObserver(ui2)

	ytChannel.(*observer.YoutubeChannel).ReleaseNewvideo(&observer.Video{Title: "Avatar 2 trailer"})
	ytChannel.(*observer.YoutubeChannel).ReleaseNewvideo(&observer.Video{Title: "Minion trailer"})
	// remove observer ui1
	log.Println("remove observer ui1")
	ytChannel.RemoveObserver(ui1)
	ytChannel.NotifyObservers()
}
