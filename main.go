package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	client := createClient(nil)
	hashes := getRegistrationHashes(client)
	hashes["eventId"], hashes["inputId"], hashes["markerId"], hashes["instanceId"] = newRegistrationEvent(client, hashes)

	account := Account{
		Email: "dqeri3nf9en2few@gmail.com",
	}
	hashes["regContext"], hashes["blokMachineId"] = submitEmail(client, hashes, account)
	eventId, inputId, marketId, instanceId := sendConfirmationCode(client, hashes, account)
	
	fmt.Println(eventId, inputId, marketId, instanceId)
}