package gmongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	MONGO_KEY_ID    = "_id"
	MONGO_OPT_SET   = "$set"
	MONGO_OPT_UNSET = "$unset"
)

type MongoAttributeGroup struct {
	collection *mgo.Collection
	id         string
}

func NewMongoAttributeGroup(collection *mgo.Collection, id string) *MongoAttributeGroup {
	return &MongoAttributeGroup{collection: collection, id: id}
}

func (this *MongoAttributeGroup) CasSetAttribute(key string, oldValue string, newValue string) bool {
	var findKey map[string]interface{}
	if len(oldValue) > 0 {
		findKey = bson.M{MONGO_KEY_ID: this.id, key: oldValue}
	} else {
		findKey = bson.M{MONGO_KEY_ID: this.id, key: nil}
	}
	var setter map[string]interface{}
	if len(newValue) > 0 {
		setter = bson.M{MONGO_OPT_SET: bson.M{key: newValue}}
	} else {
		setter = bson.M{MONGO_OPT_UNSET: bson.M{key: ""}}
	}
	change := mgo.Change{
		Update:    setter,
		ReturnNew: true,
		Upsert:    true,
	}
	res := make(map[string]interface{})
	this.collection.Find(findKey).Apply(change, &res)
	return res[key] == newValue
}

func (this *MongoAttributeGroup) GetAttribute(key string) (string, error) {
	res := make(map[string]string)
	err := this.collection.Find(bson.M{MONGO_KEY_ID: this.id}).Select(bson.M{key: 1}).One(&res)
	return res[key], err
}

func (this *MongoAttributeGroup) SetAttribute(key string, value string) error {
	_, err := this.collection.Upsert(bson.M{MONGO_KEY_ID: this.id}, bson.M{MONGO_OPT_SET: bson.M{key: value}})
	return err
}

func (this *MongoAttributeGroup) SetAttributes(values map[string]interface{}) error {
	_, err := this.collection.Upsert(bson.M{MONGO_KEY_ID: this.id}, bson.M{MONGO_OPT_SET: values})
	return err
}
