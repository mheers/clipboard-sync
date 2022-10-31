package syncer

import (
	"fmt"

	"github.com/mheers/clipboard-sync/config"
	"github.com/mheers/clipboard-sync/helpers"
	"github.com/mheers/clipboard-sync/mqclient/models"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var latestPublishedMessage = ""
var latestReceivedMessage = ""

type Syncer struct {
	cfg          *config.Config
	mqConnection *models.MQClient
	machineID    string
}

func NewSyncer(cfg *config.Config) (*Syncer, error) {
	mqConnection, err := models.NewMQClient(cfg)
	if err != nil {
		return nil, err
	}

	machineID, err := helpers.GetMachineID()
	if err != nil {
		return nil, err
	}

	return &Syncer{
		cfg:          cfg,
		mqConnection: mqConnection,
		machineID:    machineID,
	}, nil
}

func (r *Syncer) Start() error {
	err := r.addSubscriptions()
	if err != nil {
		return err
	}

	err = r.StartClipboard()
	if err != nil {
		return err
	}

	// loop forever
	select {}
}

func (r *Syncer) addSubscriptions() error {
	err := r.addClipboardSubscription()
	if err != nil {
		return err
	}
	return nil
}

// addClipboardSubscription adds a mq subscription to the instance topic
func (r *Syncer) addClipboardSubscription() error {
	subscription, err := r.mqConnection.Connection.Subscribe("instance.clipboard.*", r.handleClipboardMessage)
	if err != nil {
		return err
	}
	r.mqConnection.Subscriptions["instance"] = subscription
	fmt.Println("MQ: Subscribed to instance.clipboard.*")
	return nil
}

func (r *Syncer) publishClipboardMessage(data string) {
	if data == latestPublishedMessage {
		return
	}
	latestPublishedMessage = data
	fmt.Printf("Publishing new clipboard: %s\n", data)
	err := r.mqConnection.Connection.Publish(fmt.Sprintf("instance.clipboard.%s", r.machineID), []byte(data))
	if err != nil {
		logrus.Error(err)
	}
}

func (r *Syncer) handleClipboardMessage(msg *nats.Msg) {
	logrus.Infof("MQ: Message on subject %s: %s", msg.Subject, msg.Data)
	data := string(msg.Data)
	if data == latestReceivedMessage {
		return
	}
	latestReceivedMessage = data
	err := r.writeClipboard(data)
	if err != nil {
		logrus.Errorf("Error writing to clipboard: %s", err)
	}
}
