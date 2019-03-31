package gleveldb

type LevelDBAttributeGroup struct {
	leveldb *levelDBDriver
	id      string
}

func NewLevelDBAttributeGroup(leveldb *levelDBDriver, id string) *LevelDBAttributeGroup {
	return &LevelDBAttributeGroup{leveldb: leveldb, id: id}
}

func (this *LevelDBAttributeGroup) wrapKey(key string) string {
	return this.id + ":" + key
}

func (this *LevelDBAttributeGroup) CasSetAttribute(key string, oldValue string, newValue string) bool {
	data, err := this.leveldb.Get(this.wrapKey(key))
	if err != nil || data != oldValue {
		return false
	}
	return this.leveldb.Set(this.wrapKey(key), newValue) == nil
}

func (this *LevelDBAttributeGroup) GetAttribute(key string) (string, error) {
	return this.leveldb.Get(this.wrapKey(key))
}

func (this *LevelDBAttributeGroup) SetAttribute(key string, value string) error {
	return this.leveldb.Set(this.wrapKey(key), value)
}

func (this *LevelDBAttributeGroup) SetAttributes(values map[string]interface{}) error {
	var err error
	for k, v := range values {
		cerr := this.SetAttribute(k, v.(string))
		if cerr != nil && err == nil {
			err = cerr
		}
	}
	return err
}
