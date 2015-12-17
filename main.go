package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"
    "github.com/rockneurotiko/go-tgbot"
    "github.com/joho/godotenv"
)

func echoHandler(bot tgbot.TgBot, msg tgbot.Message, vals []string, kvals map[string]string) *string {
    newmsg := fmt.Sprintf("[Echoed]: %s", vals[1])

    return &newmsg
}

func testeHandler(bot tgbot.TgBot, msg tgbot.Message, text string) *string {
    replies := []string{
        "Peste",
        "Teste",
        "Leste",
        "Oeste",
        "Veste",
    }

    reply := fmt.Sprintf(replies[rand.Intn(len(replies))])

    return &reply
}

func instagramHandler(bot tgbot.TgBot, msg tgbot.Message, text string) *string {
    bot.Answer(msg).Text(">instagram").ReplyToMessage(msg.ID).End()

    return nil
}

func anyHandler(bot tgbot.TgBot, msg tgbot.Message) {
    rand.Seed(time.Now().Unix())

    if rand.Intn(100) == 1 {
        bot.Answer(msg).Text("Isso!").ReplyToMessage(msg.ID).End()
    }
}

func main() {
    godotenv.Load()
    token := os.Getenv("TELEGRAM_KEY")

    bot := tgbot.NewTgBot(token)

    bot.CommandFn(`echo (.+)`, echoHandler)

    bot.SimpleRegexFn(`^(?i)teste$`, testeHandler)
    bot.SimpleRegexFn(`(?i)instagram`, instagramHandler)

    bot.AnyMsgFn(anyHandler)

    bot.SimpleStart()
}
