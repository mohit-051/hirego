package cronjobs

import (
	"fmt"
)

func Cronjobs() {
	fmt.Println("Running cron job every minutes. Hello world I am being invoked.")

	SendSlackMessage("", "Hey there this message is generated through cron jobs from golang service. This is for testing.!")
}
