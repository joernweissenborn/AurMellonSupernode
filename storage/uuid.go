package storage
import (
	uuid "github.com/nu7hatch/gouuid"
	"log"
)

func GenerateUuid() string {
	Uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal("Failed to generate UUID")
		return ""
	}
	return Uuid.String()
}

