package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	client := createClient(nil)
	hashes := getRegistrationHashes(client)
	fmt.Println(hashes)
	hashes["eventId"], hashes["inputId"], hashes["markerId"], hashes["instanceId"] = newRegistrationEvent(client, hashes)

	account := Account{
		Email: newEmail(client),
	}
	fmt.Println(hashes, account.Email)
	hashes["regContext"], hashes["blokMachineId"] = submitEmail(client, hashes, account)
	eventId, inputId, marketId, instanceId := sendConfirmationCode(client, hashes, account)
	
	fmt.Println(eventId, inputId, marketId, instanceId)
}