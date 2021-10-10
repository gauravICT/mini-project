package pulsar

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// PulsarProducer : for producing msg into topic
func PulsarProducer(ctx context.Context, msg string) error {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  5 * time.Second,
		ConnectionTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "employeeInfo",
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msg),
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")

	return nil

}

// PulsarConsumer : For consuming mas from topic
func PulsarConsumer(ctx context.Context, topic string) error {
	client, err := pulsar.NewClient(pulsar.ClientOptions{URL: "pulsar://localhost:6650"})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	reader, err := client.CreateReader(pulsar.ReaderOptions{
		Topic:          "persistent://public/default/" + topic,
		StartMessageID: pulsar.EarliestMessageID(),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	for reader.HasNext() {
		msg, err := reader.Next(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		// writing into file
		f, err := os.OpenFile("text.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		//defer f.Close()
		if _, err := f.WriteString(string(msg.Payload())); err != nil {
			log.Println(err)
		}

	}
	return nil
}
