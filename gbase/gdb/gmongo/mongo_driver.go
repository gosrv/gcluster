package gmongo

import (
	"github.com/globalsign/mgo"
	"github.com/gosrv/gcluster/gbase/gdb"
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/goioc"
	"reflect"
)

type IMongoDBDriver interface {
	Name() string
	MongoDb()
	GetCollection(name string) *mgo.Collection
	GetAttributeGroup(group, id string) gdb.IDBAttributeGroup
}

var IMongoDBDriverType = reflect.TypeOf((*IMongoDBDriver)(nil)).Elem()

type mongoDBDriver struct {
	name    string
	dbname  string
	session *mgo.Session
}

func (this *mongoDBDriver) Name() string {
	return this.name
}

func NewMongoDriver(name string, url string, dbname string) *mongoDBDriver {
	session, err := mgo.Dial(url)
	if err != nil {
		glog.Panic("mongo dial failed %v", err)
	}
	return &mongoDBDriver{
		name:    name,
		dbname:  dbname,
		session: session,
	}
}

func (this *mongoDBDriver) MongoDb() {
}

func (this *mongoDBDriver) GetCollection(name string) *mgo.Collection {
	return this.session.DB(this.dbname).C(name)
}

func (this *mongoDBDriver) GetAttributeGroup(group, id string) gdb.IDBAttributeGroup {
	return NewMongoAttributeGroup(this.GetCollection(group), id)
}

func (this *mongoDBDriver) GetPriority() int {
	return gioc.PriorityMiddle
}

/*
func (this *mongoDBDriver) GetMessageQueue(queue string) gdb.IMessageQueue {
	return NewMongoMessageQueue(queue, session)
}
*/
