package mqclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mheers/clipboard-sync/config"
)

func TestInit(t *testing.T) {
	cfg := config.GetFakeConfig()

	mqClient, err := Init(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, mqClient)
}
