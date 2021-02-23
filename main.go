package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"vnfco.ir/rabbit/fail"
	"vnfco.ir/rabbit/rabbit"
)

func init() {
	err := godotenv.Load(".env")
	fail.FailOnError(err, "Can't Load ENV file")
}

func main() {
	_, channel := rabbit.ConnectToAMPQServerAndCreateChannel(os.Getenv("AMPQ_HOST"), os.Getenv("AMPQ_PORT"), os.Getenv("AMPQ_USERNAME"), os.Getenv("AMPQ_PASSWORD"))
	queueHello := rabbit.CreateOrJoinSimpleQueue(channel, "hi")
	queueReply := rabbit.CreateOrJoinSimpleQueue(channel, "reply")
	msgs := rabbit.Listen(channel, queueHello)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s with Headers : %s", d.Body, d.Headers)
			BodyMap := map[string]string{
				"title": "Hello baby I love you",
				"body":  "I love Your More Than You Ever Think ",
			}
			jsonBody, err := json.Marshal(BodyMap)
			fail.FailOnError(err, "Can Not Stringify Map")
			rabbit.Publish(channel, queueReply, string(jsonBody), "application/json")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
