package clients

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/AbrakadabraK/telegram-bot-pet/model"
	"github.com/pkg/errors"
)

const (
	getUpd = "getUpdates"
)

func (tgc *TelegramClient) GetUpdates(ctx context.Context, offset int, limit int) ([]*model.Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := tgc.doRequest(ctx, getUpd, q)
	if err != nil {
		return nil, errors.Errorf("error do req for get updates : %v", err)
	}

	var res model.UpdateResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, errors.Errorf("error unmarshaling update response : %v", err)
	}
	return res.Updates, nil
}
