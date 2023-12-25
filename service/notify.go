package service

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
)

func NotifyOnline() {
	message := fmt.Sprintf("User is online now! %s", time.Now().Format("2006-01-02 15:04:05"))

	err := beeep.Notify("User Online", message, "")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Printf("Notification sent")

}
