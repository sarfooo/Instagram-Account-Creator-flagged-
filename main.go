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
		Email: "feiofj22239fjei@gmail.com",
	}
	hashes["regContext"], hashes["blokMachineId"] = submitEmailEvent(client, hashes, account)
	
}