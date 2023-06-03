package main

import (
	"github.com/Mau005/KrayManagerUpdate/schema"
)

func main() {
	querys := schema.NewQuerys("data/config.json")
	querys.ChangePosition(3552, 3650, 7)

}
