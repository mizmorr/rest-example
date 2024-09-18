package redis

import (
	"context"
	"encoding/json"

	"github.com/mizmorr/rest-example/config"
	"github.com/redis/go-redis/v9"
)

type cache struct {
	*redis.Client
}

type data struct {
	Github      string `json:"github"`
	Pgaddress   string `json:"pgaddress"`
	Httpaddress string `json:"httpaddress"`
	Greetings   string `json:"greetings"`
}

func (d data) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *data) UnmarshalBinary(res []byte) error {
	return json.Unmarshal(res, &d)
}

func New(dbnum int, addr, pasw string) *cache {

	return &cache{
		redis.NewClient(
			&redis.Options{
				Addr:     addr,
				DB:       dbnum,
				Password: pasw,
			}),
	}
}

func (c *cache) Setup(ctx context.Context) error {

	cfg := config.Get()

	err := c.FlushDB(ctx).Err()

	if err != nil {
		return err
	}

	curData := data{
		Github:      "github.com/mizmorr",
		Pgaddress:   cfg.PgURL,
		Httpaddress: cfg.HTTPAddress,
		Greetings:   "Hello everyone!",
	}

	if err := c.Set(ctx, "1", curData, 0).Err(); err != nil {
		return err
	}
	return nil

}

func (c *cache) Take(ctx context.Context) interface{} {
	var d data

	str, _ := c.Get(ctx, "1").Result()

	err := d.UnmarshalBinary([]byte(str))
	if err != nil {
		return nil
	}

	return &d
}
