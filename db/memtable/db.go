package memtable

import (
	"context"
	"fmt"
	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/db/memtable/rest"
	"github.com/pingcap/go-ycsb/pkg/util"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
)

type Memtable struct {
	db      *rest.Client
	bufPool *util.BufPool
	r       *util.RowCodec
}

func (tbl *Memtable) Close() error {
	tbl.db.Close()
	return nil
}

func (tbl *Memtable) InitThread(ctx context.Context, threadID int, threadCount int) context.Context {
	return ctx
}

func (tbl *Memtable) CleanupThread(_ context.Context) {
}

func (tbl *Memtable) getRowKey(table string, key string) string {
	return fmt.Sprintf("%s:%s", table, key)
}

func (tbl *Memtable) Insert(ctx context.Context, table string, key string, values map[string][]byte) error {
	rowKey := tbl.getRowKey(table, key)

	buf := tbl.bufPool.Get()
	defer func() {
		tbl.bufPool.Put(buf)
	}()

	buf, err := tbl.r.Encode(buf, values)
	if err != nil {
		return err
	}
	tbl.db.Put(rowKey, buf)

	return nil
}

func (tbl *Memtable) Read(ctx context.Context, table string, key string, fields []string) (map[string][]byte, error) {
	var m map[string][]byte
	rowKey := tbl.getRowKey(table, key)
	row := tbl.db.Get(rowKey)
	m, _ = tbl.r.Decode(row, fields)
	return m, nil
}

func (tbl *Memtable) Scan(ctx context.Context, table string, startKey string, count int, fields []string) ([]map[string][]byte, error) {
	res := make([]map[string][]byte, count)
	rowStartKey := tbl.getRowKey(table, startKey)

	items := tbl.db.Scan(rowStartKey, count)
	i := 0
	for _, item := range items {
		m, err := tbl.r.Decode(item, fields)
		if err != nil {
			return nil, err
		}
		res[i] = m
	}

	return res, nil
}

func (tbl *Memtable) Update(ctx context.Context, table string, key string, values map[string][]byte) error {
	rowKey := tbl.getRowKey(table, key)

	buf := tbl.bufPool.Get()
	defer func() {
		tbl.bufPool.Put(buf)
	}()

	buf, err := tbl.r.Encode(buf, values)
	if err != nil {
		return err
	}
	tbl.db.Put(rowKey, buf)

	return nil
}

func (tbl *Memtable) Delete(ctx context.Context, table string, key string) error {
	rowKey := tbl.getRowKey(table, key)
	tbl.db.Delete(rowKey)
	return nil
}

func init() {
	ycsb.RegisterDBCreator("memtable", Creator{})
}

type Creator struct {
}

func (c Creator) Create(p *properties.Properties) (ycsb.DB, error) {
	return &Memtable{
		db:      rest.NewClient(),
		r:       util.NewRowCodec(p),
		bufPool: util.NewBufPool(),
	}, nil
}
