// Package livestreaming https://doc.yunxin.163.com/docs/jEyODA1OTE/zkzMDA3NTU?platformId=110057
package livestreaming

type ChannelType int

const (
	ChannelTypeRTMP ChannelType = 0
)

type CreateChannelRequest struct {
	Name string      `json:"name"`
	Type ChannelType `json:"type"`
}
type CreateChannelResponse struct {
	ChannelID   string `json:"cid"`
	CreateTime  int64  `json:"ctime"`
	Name        string `json:"name"`
	PushURL     string `json:"pushUrl"`
	HTTPPullURL string `json:"httpPullUrl"`
	HLSPullURL  string `json:"hlsPullURL"`
	RTMPPullURL string `json:"rtmpPullUrl"`
	RTSPullURL  string `json:"rtsPullURL"`
}

func (c *Client) CreateChannel(req *CreateChannelRequest) (*Response[CreateChannelResponse], error) {
	res := &Response[CreateChannelResponse]{
		Ret: CreateChannelResponse{},
	}
	err := c.doRequest("/app/channel/create", req, res)
	return res, err
}

type DeleteChannelRequest struct {
	ChannelID string `json:"cid"`
}
type DeleteChannelResponse struct {
}

func (c *Client) DeleteChannel(req *DeleteChannelRequest) (*Response[DeleteChannelResponse], error) {
	res := &Response[DeleteChannelResponse]{
		Ret: DeleteChannelResponse{},
	}
	err := c.doRequest("/app/channel/delete", req, res)
	return res, err
}

type GetChannelAddressRequest struct {
	ChannelID string `json:"cid"`
}
type GetChannelAddressResponse struct {
	PushURL     string `json:"pushUrl"`
	HTTPPullURL string `json:"httpPullUrl"`
	HLSPullURL  string `json:"hlsPullURL"`
	RTMPPullURL string `json:"rtmpPullUrl"`
	RTSPullURL  string `json:"rtsPullURL"`
}

func (c *Client) GetChannelAddress(req *GetChannelAddressRequest) (*Response[GetChannelAddressResponse], error) {
	res := &Response[GetChannelAddressResponse]{
		Ret: GetChannelAddressResponse{},
	}
	err := c.doRequest("/app/address", req, res)
	return res, err
}
