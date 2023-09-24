package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	TelegramBotToken string `json:"telegram_bot_token"`
}

type SensorData struct {
	AHumidity    string `json:"a_humidity"`
	ATemperature string `json:"a_temperature"`
}

func fetchSensorData() (*SensorData, error) {
	url := "http://10.5.4.7/temperatureactual"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data SensorData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		return nil, err
	}
	return &data, nil
}

func main() {

	fmt.Println("iniciando...")
	// Lectura de datos en API

	// Leer el archivo de configuracion
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error al abrir el archivo de configuracion", err)
	}

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal("Error al decodificar el archivo de configuracion", err)
	}

	// Acceder al token
	token := config.TelegramBotToken

	// instancia de BotAPI
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// habilitar el modo depuracion
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "ayuda":
			chatID := int64(5662534540)
			fileID := "CAACAgIAAxkBAANSZQ96opB0miPq8Y5q-HJut-zboswAAvoQAAKhxyhIOWV265NYB6MwBA"

			StickerMessage := tgbotapi.NewSticker(chatID, tgbotapi.FileID(fileID))
			_, err = bot.Send(StickerMessage)

			msg.Text = "Los comandos disponibles son: \n\n  /temperatura \n\n /humedad."
			if err != nil {
				log.Panic(err)
			}

		case "temperatura":

			// Obtener la temperatura
			data, err := fetchSensorData()
			if err != nil {
				log.Println("Error al obtener los datos de la API")
				msg.Text = "Error al obtener los datos de la API"
			} else {
				msg.Text = fmt.Sprintf("🌡 La temperatura actual en la incubadora es: %s °C", data.ATemperature)
			}

		case "humedad":

			// Obtener la humedad
			data, err := fetchSensorData()
			if err != nil {
				log.Println("Error al obtener los datos de la API")
				msg.Text = "Error al obtener los datos de la API"
			} else {
				msg.Text = fmt.Sprintf("💧 La humedad actual en la incubadora es: %s %%", data.AHumidity)
			}

		default:

			chatID := int64(5662534540)
			fileID := "CAACAgIAAxkBAANJZQ93-VLDn67zWDNqB3kOjrIzns0AAkUYAAIUqPBIVd-bm1T8xSswBA"

			StickerMessage := tgbotapi.NewSticker(chatID, tgbotapi.FileID(fileID))
			_, err = bot.Send(StickerMessage)
			if err != nil {
				log.Panic(err)
			}

			msg := "No conozco ese comando \n\n Usá /ayuda para ver los comandos disponibles"
			textMessage := tgbotapi.NewMessage(chatID, msg)
			_, err = bot.Send(textMessage)
			if err != nil {
				log.Panic(err)
			}

		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

}
