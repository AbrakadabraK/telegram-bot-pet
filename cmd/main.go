package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AbrakadabraK/telegram-bot-pet/config"
	"github.com/AbrakadabraK/telegram-bot-pet/internal/clients"
)

func main() {
	appCtx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}
	token, err := config.GetToken(cfg, "token")
	if err != nil {
		log.Fatalf("Ошибка получения токена: %v", err)
	}

	telegramHost, err := config.GetToken(cfg, "tgBotHost")
	if err != nil {
		log.Fatalf("Ошибка получения хоста у ТГ : %v", err)
	}

	tgClient := clients.New(telegramHost, token)
	fmt.Println(token)
	fmt.Println(tgClient)

	up, err := tgClient.GetUpdates(appCtx, 0, 100)
	if err != nil {
		log.Fatalf("Ошибка получения Updates : %v", err)
	}

	for _, r := range up {
		fmt.Println(r.Message.Text)
		fmt.Println("\n")
	}

}
