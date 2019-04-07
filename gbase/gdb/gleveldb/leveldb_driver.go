package gleveldb

import (
	"github.com/gosrv/gcluster/gbase/gdb"
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/goioc"
	"github.com/syndtr/goleveldb/leveldb"
	"reflect"
)

type ILevelDBDriver interface {
	LevelDb()
	Get(key string) (string, error)
	Set(key string, value string) error
	GetAttributeGroup(group, id string) gdb.IDBAttributeGroup
}

var ILevelDBDriverType = reflect.TypeOf((*ILevelDBDriver)(nil)).Elem()

type levelDBDriver struct {
	dbname string
	db     *leveldb.DB
}

func NewLevelDBDriver(url string, dbname string) *levelDBDriver {
	db, err := leveldb.OpenFile(url, nil)
	if err != nil {
		gl.Panic("level db [%v] open failed %v", url, err)
	}
	return &levelDBDriver{
		dbname: dbname,
		db:     db,
	}
}

func (this *levelDBDriver) wrapKey(key string) string {
	return this.dbname + ":" + key
}

func (this *levelDBDriver) LevelDb() {
}

func (this *levelDBDriver) Get(key string) (string, error) {
	data, err := this.db.Get([]byte(this.wrapKey(key)), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (this *levelDBDriver) Set(key string, value string) error {
	return this.db.Put([]byte(this.wrapKey(key)), []byte(value), nil)
}

func (this *levelDBDriver) GetAttributeGroup(group, id string) gdb.IDBAttributeGroup {
	return NewLevelDBAttributeGroup(this, id)
}

func (this *levelDBDriver) GetPriority() int {
	return gioc.PriorityLow - 1
}
