package mqclient

import (
	"sync"

	"github.com/mheers/clipboard-sync/config"
	mqmodels "github.com/mheers/clipboard-sync/mqclient/models"
)

var MQClient *mqmodels.MQClient

var once sync.Once

// Init initializes a message queue client
func Init(appConfig *config.Config) (*mqmodels.MQClient, error) {
	var err error
	once.Do(func() {

		mqClient, err := mqmodels.NewMQClient(appConfig)
		if err != nil {
			return
		}

		MQClient = mqClient

	})
	return MQClient, err
}
