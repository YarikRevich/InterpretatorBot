package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"Bot/States"
	"Bot/BotUtils"
	"Bot/Utils"
	"Bot/Parcer"
)


const (
	SiteWithExpl string = "http://sum.in.ua/?swrd="
)

var states = States.States{
	Wait: false, 
	Question: false}

var info = States.Info{
	Trials: 0,
	WordsToIgnore: []string{"🌐Розпочати пошук"}}


func main(){

	// используя токен создаем новый инстанс бота
	Bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	
	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := Bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	//Checks all the updates 
	for Update := range updates {

		reply := "👌🏿Натисни будь-яку кнопку для того, щоб розпочати розмову"
		keyboard := BotUtils.CreateKeyBoardHi(Bot, Update, []string{""})

		if Update.Message == nil {
			continue
		}

		//Checks whether answer mode is not active if it does it sends welcome message and states mode active
		if states.Wait != true{
			switch Update.Message.Text{
			case "🌐Розпочати пошук":
				reply = "🖊Введіть слово для тлумачення"
				states.Wait = true
			}
		}else{
			//Checks whether user hasn't made more than 3 trials if he does it sends warning message
			if info.Trials <= 3{
				//If sent message is in exception list it sends equal message
				if Utils.CheckWordExistance(Update.Message.Text, info.WordsToIgnore){
					states.Wait = false
					reply = "⚠️Ви призупинили пошук тлумачень"
				}else{
					//Gets explanation of the written word and if it's giant it sends special link to check it in browser
					reply = Parcer.GetWordDescription(Update.Message.Text, SiteWithExpl)
					if len(reply) > 4096{
						reply = "❗️Тлумачення виявилося занадто велике.\nПерейди за посиланням нижче, щоб подивитися тлумачення\n" + SiteWithExpl + Update.Message.Text 
					}
					info.Trials++	
				}	
			}else{
				reply = "😓Ти зробив багато спроб, дай перепочинути"
				info.Trials = 0
				states.Wait = false
			}
		}

		//Processes comands
		switch Update.Message.Command() {
		case "start":
			states.Wait = false
			reply = "😛Привіт, я бот для тлумачення слів"
			keyboard = BotUtils.CreateKeyBoardHi(Bot, Update, []string{"🌐Розпочати пошук"})
		}

		//Sends messages written in the 'reply' variable
		msg := tgbotapi.NewMessage(Update.Message.Chat.ID, reply)
		msg.ReplyMarkup = keyboard
		Bot.Send(msg)
	}
}

