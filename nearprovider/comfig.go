package nearprovider

import (
	"context"
	"github.com/rarimo/near-go/nearclient"
	"github.com/rarimo/near-go/nearprovider/s3"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

type RPCs struct {
	Main    *nearclient.Client `fig:"main,required"`
	History *nearclient.Client `fig:"history,required"`
}

type Config struct {
	S3  s3.Config `fig:"s3,required"`
	RPC RPCs      `fig:"rpc,required"`
}

type Nearer interface {
	Near() Provider
}

type nearer struct {
	getter kv.Getter
	once   comfig.Once
	log    *logan.Entry
}

func NewNearer(getter kv.Getter, log *logan.Entry) Nearer {
	return &nearer{
		getter: getter,
		log:    log,
	}
}

func (s *nearer) Near() Provider {
	return s.once.Do(func() interface{} {
		var cfg Config
		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(s.getter, "near_provider")).
			With(figure.BaseHooks, NearHooks).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out near provider config"))
		}

		s3conn, err := s3.New(cfg.S3)
		if err != nil {
			panic(errors.Wrap(err, "failed to create s3 client"))
		}

		_, err = cfg.RPC.Main.NetworkInfo(context.TODO())
		if err != nil {
			panic(errors.Wrap(err, "failed to dial main rpc"))
		}

		_, err = cfg.RPC.History.NetworkInfo(context.TODO())
		if err != nil {
			panic(errors.Wrap(err, "failed to dial history rpc"))
		}

		return New(cfg.RPC.Main, cfg.RPC.History, s3conn, s.log.WithField("who", "near_provider"))
	}).(Provider)
}

var NearHooks = figure.Hooks{
	"*nearclient.Client": func(raw interface{}) (reflect.Value, error) {
		v, err := cast.ToStringE(raw)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "expected string")
		}

		client, err := nearclient.New(v)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to create near rpc client")
		}

		return reflect.ValueOf(client), nil
	},
}
