package BotUtils

import "github.com/Syfaro/telegram-bot-api"

func CreateKeyBoardHi(Bot *tgbotapi.BotAPI, Updates tgbotapi.Update, fields []string)tgbotapi.ReplyKeyboardMarkup {
	/*Creates a new keyboard with such buttons
	:: "Розпочати пошук" - to start a research*/
	
	markup := tgbotapi.ReplyKeyboardMarkup{ResizeKeyboard: true}
	var row []tgbotapi.KeyboardButton
	for _, value := range(fields){
		row = append(row, tgbotapi.KeyboardButton{Text: value})	
	}
	markup.Keyboard = append(markup.Keyboard, row)
	return markup
}