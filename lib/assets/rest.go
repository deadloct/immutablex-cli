package assets

import (
	"context"
	"errors"
	"net/http"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

type RESTClientConfig struct {
	url string
}

type RESTClient struct {
	client *http.Client
	url    string
}

func NewRESTClient(cfg RESTClientConfig) *RESTClient {
	return &RESTClient{client: &http.Client{}, url: cfg.url}
}

func (c *RESTClient) Start() error {
	return nil
}

func (c *RESTClient) Stop() {}

func (c *RESTClient) GetAsset(ctx context.Context, tokenAddress, tokenID string, includeFees bool) (*api.Asset, error) {
	return nil, errors.New("not yet created")
}

func (c *RESTClient) ListAssets(context.Context, ListAssetsConfig) ([]api.AssetWithOrders, error) {
	return nil, errors.New("not yet created")
}
