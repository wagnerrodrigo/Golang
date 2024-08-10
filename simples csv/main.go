package main

import (
	"encoding/csv"
	"os"
)

// https://www.youtube.com/watch?v=gXmznGEW9vo
func main() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"1", "teste", "hoje"})

	writer.Write([]string{"2", "teste 2", "Amanha"})

	writer.Flush()
}
