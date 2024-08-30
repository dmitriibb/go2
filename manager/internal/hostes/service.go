package hostes

import (
	"fmt"
	"math/rand"
)

func generateIdForClient(clientName string) string {
	id := fmt.Sprintf("%v-%v", clientName, rand.Intn(10000))
	logger.Debug("assign id %v to client '%v'", id, clientName)
	// TODO save id to the database
	return id
}

func getAvailableTableNumber(clientId string) int {
	id := rand.Intn(200)
	logger.Debug("assign table %v to client '%v'", id, clientId)
	// TODO manage available tables
	return id
}
