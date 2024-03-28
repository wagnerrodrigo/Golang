package main

import (
	"fmt"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

func main() {
	speech := htgotts.Speech{
		Folder:   "audio",
		Language: voices.Portuguese,
	}

	speech.Speak("Esse texto vai virar um áudio.")
	fmt.Println("Finalizando..")

}
