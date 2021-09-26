package etcd

import (
	"context"
	"log"

	"github.com/nadilas/moar/internal"
	"go.etcd.io/etcd/client/v3"
)

type Storage struct {
	etcd *clientv3.Client
}

func New() *Storage {
	client, err := clientv3.New(clientv3.Config{})
	if err != nil {
		log.Fatalln("failed to connect to etcd: ", err)
	}
	return &Storage{etcd: client}
}

func (d *Storage) UriForModule(ctx context.Context, module *internal.Module) (string, error) {
	panic("implement me")
}

func (d *Storage) PutModule(module *internal.Module) error {
	panic("implement me")
}

func (d *Storage) GetModule(ctx context.Context, name string) (*internal.Module, error) {
	panic("implement me")
}

func (d *Storage) Shutdown() {
	err := d.etcd.Close()
	if err != nil {
		panic(err)
	}
}
