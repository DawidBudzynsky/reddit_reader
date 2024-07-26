package texttospeach

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	apiURL       = "https://tiktok-tts.weilnet.workers.dev/api/generation"
	apiCharLimit = 300
)

type RequestPayload struct {
	Text   string `json:"text"`
	Voice  string `json:"voice"`
	Base64 bool   `json:"base64"`
}

type DataResponse struct {
	Data string `json:"data"`
}

type TiktokTTS struct {
	apiURL          string
	voice           Voice
	fileDestination string
}

func NewTikTokTTS(voice Voice, fileDestination string) *TiktokTTS {
	return &TiktokTTS{
		apiURL:          apiURL,
		voice:           voice,
		fileDestination: fileDestination,
	}
}

// sets the file destination path for saving the MP3 file.
func (t *TiktokTTS) SetDestinationPath(destination string) {
	t.fileDestination = destination
}

// converts the provided text to MP3 and saves it to the specified file destination.
func (t *TiktokTTS) TextToMp3(text string) {
	mp3Data, err := t.createMP3Data(RequestPayload{
		Text:   text,
		Voice:  string(t.voice),
		Base64: true,
	})
	if err != nil {
		log.Fatalf("Couldn't create a request to api: %s\nerror: %v", t.apiURL, err)
	}

	if err = t.saveAsMp3(mp3Data, t.fileDestination); err != nil {
		log.Fatal("Couldn't save as mp3")
	}
}

// sets the voice for the TTS service.
func (t *TiktokTTS) chooseVoice(voice Voice) {
	t.voice = voice
}

// creates the MP3 data by calling the TTS API and combining the results.
func (t *TiktokTTS) createMP3Data(payload RequestPayload) (string, error) {
	var chunks []string
	if len(payload.Text) > 300 {
		chunks = splitText(payload.Text, apiCharLimit)
	} else {
		chunks = append(chunks, payload.Text)
	}

	bytesData, err := t.fetchAudio(chunks)
	if err != nil {
		return "", err
	}
	combinedBytesData, err := combineAudio(bytesData)
	if err != nil {
		return "", nil
	}
	return combinedBytesData, nil
}

// creates an HTTP POST request to the TTS API with the provided payload.
func (t *TiktokTTS) createRequest(payload RequestPayload) ([]byte, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", t.apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// sends requests for each chunk of text to the TTS API concurrently and collects the responses.
// It ensures that the returned data is in the correct order.
func (t *TiktokTTS) fetchAudio(chunks []string) ([][]byte, error) {
	var wg sync.WaitGroup
	responses := make([][]byte, len(chunks))
	errChan := make(chan error, len(chunks))

	for i, chunk := range chunks {
		wg.Add(1)
		go func(i int, chunk string) {
			defer wg.Done()
			payload := RequestPayload{
				Text:   chunk,
				Voice:  string(t.voice),
				Base64: true,
			}
			body, err := t.createRequest(payload)
			if err != nil {
				errChan <- fmt.Errorf("request error: %v", err)
				return
			}
			responses[i] = body
		}(i, chunk)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, <-errChan
	}
	return responses, nil
}

// saves the provided base64-encoded MP3 data to a file.
func (t *TiktokTTS) saveAsMp3(data string, fileName string) error {
	buf, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, buf, 0644)
	if err != nil {
		log.Fatal("write file failed")
		return err
	}
	return nil
}

// splits the text into chunks of the specified maximum length.
func splitText(text string, maxLen int) []string {
	var chunks []string
	for len(text) > maxLen {
		splitPos := strings.LastIndex(text[:maxLen], " ")
		if splitPos == -1 {
			splitPos = maxLen
		}
		chunks = append(chunks, text[:splitPos])
		text = text[splitPos:]
	}
	chunks = append(chunks, text)
	return chunks
}

// combines the base64-encoded audio data from each chunk into a single string.
func combineAudio(chunks [][]byte) (string, error) {
	var returnString string
	for _, chunk := range chunks {
		var dataResponse DataResponse
		if err := json.Unmarshal(chunk, &dataResponse); err != nil {
			return "", err
		}
		returnString = returnString + dataResponse.Data
	}
	return returnString, nil
}
