package livestreaming

import "testing"

const (
	testAppKey    = ""
	testAppSecret = ""
)

var (
	c = NewClient(API_GATEWAY_GLOBAL, testAppKey, testAppSecret)
)

func TestCreateChannel(t *testing.T) {
	res, err := c.CreateChannel(&CreateChannelRequest{
		Name: "test-channel",
	})
	t.Logf("err=%v, res=%+v", err, res)
}

func TestDeleteChannel(t *testing.T) {
	res, err := c.DeleteChannel(&DeleteChannelRequest{
		ChannelID: "channel-id",
	})
	t.Logf("err=%v, res=%+v", err, res)
}

func TestGetChannelAddress(t *testing.T) {
	res, err := c.GetChannelAddress(&GetChannelAddressRequest{
		ChannelID: "channel-id",
	})
	t.Logf("err=%v, res=%+v", err, res)
}
