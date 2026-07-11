package auxiliar

import (
	"fmt"

	"github.com/google/uuid"
)

func escrevendo3() {
	uuid := uuid.New().String()

	fmt.Println(uuid)
}
