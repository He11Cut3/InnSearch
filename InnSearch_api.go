package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Suggestion struct {
	Suggestions []Model `json:"suggestions"`
}
type Model struct {
	Value              string `json:"value"`
	Unrestricted_value string `json:"unrestricted_value"`
	Data               Data   `json:"data"`
}
type Data struct {
	KPP        string     `json:"kpp"`
	Management Management `json:"management"`
}
type Management struct {
	Name string `json:"name"`
	Post string `json:"post"`
}

func Search() {
	fmt.Printf("Введите ИНН: ")
	var inn string
	_, err := fmt.Scanf("%s", &inn)
	if err != nil {
		fmt.Printf("Не правильно введён ИНН.")
	} else {
		token := "(api_token)"
		url_api := "http://suggestions.dadata.ru/suggestions/api/4_1/rs/findById/party"
		query := fmt.Sprintf(`{"query":"%s"}`, inn)
		reqs, err := http.NewRequest("POST", url_api, bytes.NewBuffer([]byte(query)))
		if err != nil {
			fmt.Printf("Ошибка в создании запроса. Повторите попытку.")
		} else {
			reqs.Header.Add("Authorization", "Token "+token)
			reqs.Header.Add("Accept", "application/json")
			reqs.Header.Add("Content-Type", "application/json")
			client := http.Client{}
			send, err := client.Do(reqs)
			defer reqs.Body.Close()
			if err != nil {
				fmt.Printf("Ошибка при отправке запроса. Повторите попытку.")
			} else {
				body, err := io.ReadAll(send.Body)
				if err != nil {
					log.Panic(err)
				} else {
					var (
						suggestions Suggestion
					)
					_ = json.Unmarshal(body, &suggestions)
					for _, suggestion := range suggestions.Suggestions {
						fmt.Printf("Наименование: %s\nКПП: %s\nСотрудник: %s\nДолжность: %s\n", suggestion.Value, suggestion.Data.KPP, suggestion.Data.Management.Name, suggestion.Data.Management.Post)
					}
				}
			}
		}
	}
}
func main() {
	Search()
}
