package segring

import (
	"context"
	"fmt"
	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/db/segring/lib"
	"github.com/pingcap/go-ycsb/pkg/util"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
)

type tsw struct {
	db      *lib.Client
	bufPool *util.BufPool
	r       *util.RowCodec
}

func (db *tsw) Close() error {
	return nil
}

func (db *tsw) InitThread(ctx context.Context, threadID int, threadCount int) context.Context {
	return ctx
}

func (db *tsw) CleanupThread(_ context.Context) {
}

func (db *tsw) getRowKey(table string, key string) string {
	return fmt.Sprintf("%s:%s", table, key)
}

func (db *tsw) Insert(ctx context.Context, table string, key string, values map[string][]byte) error {
	rowKey := db.getRowKey(table, key)

	buf := db.bufPool.Get()
	defer func() {
		db.bufPool.Put(buf)
	}()

	buf, err := db.r.Encode(buf, values)
	if err != nil {
		return err
	}
	db.db.Add(rowKey, buf)

	return nil
}

func (db *tsw) Read(ctx context.Context, table string, key string, fields []string) (map[string][]byte, error) {
	var m map[string][]byte
	rowKey := db.getRowKey(table, key)
	row := db.db.Read(rowKey)
	m, _ = db.r.Decode(row, fields)
	return m, nil
}

func (db *tsw) Scan(ctx context.Context, table string, startKey string, count int, fields []string) ([]map[string][]byte, error) {
	res := make([]map[string][]byte, count)
	rowStartKey := db.getRowKey(table, startKey)

	items := db.db.Scan(rowStartKey, count)
	i := 0
	for _, item := range items {
		m, err := db.r.Decode(item, fields)
		if err != nil {
			return nil, err
		}
		res[i] = m
	}

	return res, nil
}

func (db *tsw) Update(ctx context.Context, table string, key string, values map[string][]byte) error {
	rowKey := db.getRowKey(table, key)

	buf := db.bufPool.Get()
	defer func() {
		db.bufPool.Put(buf)
	}()

	buf, err := db.r.Encode(buf, values)
	if err != nil {
		return err
	}
	db.db.Add(rowKey, buf)

	return nil
}

func (db *tsw) Delete(ctx context.Context, table string, key string) error {
	rowKey := db.getRowKey(table, key)
	db.db.Delete(rowKey)
	return nil
}

func init() {
	ycsb.RegisterDBCreator("segring", tswCreator{})
}

type tswCreator struct {
}

func (t tswCreator) Create(p *properties.Properties) (ycsb.DB, error) {
	return &tsw{
		db:      lib.NewClient(),
		r:       util.NewRowCodec(p),
		bufPool: util.NewBufPool(),
	}, nil
}
