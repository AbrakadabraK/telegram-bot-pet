package clients

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/AbrakadabraK/telegram-bot-pet/internal/service"
	"github.com/pkg/errors"
)

type TelegramClient struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) service.TelegramManager {
	return &TelegramClient{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (tgc *TelegramClient) doRequest(ctx context.Context, method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   tgc.host,
		Path:   path.Join(tgc.basePath, method),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Errorf("start new reques err : %v", err)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := tgc.client.Do(req)
	if err != nil {
		return nil, errors.Errorf("error do request : %v", err)
	}

	defer func() { _ = resp.Body.Close() }()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("error read body : %v", err)
	}
	return res, nil
}
