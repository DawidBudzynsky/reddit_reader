package texttospeach

type TextToSpeach interface {
	TextToMp3(inputText string)
	SetDestinationPath(destination string)
}
