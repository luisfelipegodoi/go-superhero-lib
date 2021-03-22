package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

func SendSlackNotification(webHookUrl, message string) error {

	slackBody, err := json.Marshal(SlackRequestBody{Text: message})
	if err != nil {
		return errors.New("Erro ao serializar mensagem para JSON")
	}

	req, err := http.NewRequest(http.MethodPost, webHookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return errors.New("Erro ao fazer request para o slack")
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Erro ao instanciar client para request")
	}

	buff := new(bytes.Buffer)
	buff.ReadFrom(resp.Body)
	if buff.String() != "ok" {
		return errors.New("A mensagem não pode ser enviada para o slack")
	}

	return nil
}
