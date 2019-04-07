package gredis

import (
	"github.com/go-redis/redis"
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
	"net/url"
	"reflect"
	"strings"
)

const (
	redisConfigUrlValue        = "redis.url"
	redisConfigStandaloneValue = "redis.standalone"
	redisConfigClusterValue    = "redis.cluster"
)

type AutoConfigReids struct {
	appGroup string `cfg:"app.group"`
	// 启动条件
	gioc.IBeanCondition
	gioc.IConfigBase
	url              string                `cfg.d:"redis.url"`
	standaloneOption *redis.Options        `cfg.d:"redis.standalone"`
	clusterOption    *redis.ClusterOptions `cfg.d:"redis.cluster"`
	IRedisDriver
	tagProcessor *redisTagProcessor
	domain       string
}

var _ gioc.ITagProcessor = (*AutoConfigReids)(nil)

func (this *AutoConfigReids) TagProcessorName() string {
	return this.tagProcessor.TagProcessorName()
}

func (this *AutoConfigReids) TagProcess(bean interface{}, field reflect.Value, tags map[string]string) {
	this.tagProcessor.TagProcess(bean, field, tags)
}

func NewAutoConfigReids(cfgBase, domain string) *AutoConfigReids {
	return &AutoConfigReids{
		IBeanCondition: gioc.NewConditionOnValue(cfgBase, true),
		IConfigBase:    gioc.NewConfigBase(cfgBase),
		domain:         domain,
	}
}

func (this *AutoConfigReids) PrepareProcess() {
	util.Assert(this.IRedisDriver == nil, "")

	initNum := 0
	if len(this.url) > 0 {
		initNum++
	}
	if this.standaloneOption != nil {
		initNum++
	}
	if this.clusterOption != nil {
		initNum++
	}
	if initNum != 1 {
		gl.Panic("redis config init ambiguous")
	}
	if len(this.url) > 0 {
		redisUrl, err := url.Parse(this.url)
		if err != nil {
			gl.Panic("redis url [%v] parse error %v", this.url, err)
		}
		if len(redisUrl.Host) == 0 {
			gl.Panic("redis url [%v] no server host find", this.url)
		}
		queryCluster := redisUrl.Query()["cluster"]
		hosts := strings.Split(redisUrl.Host, ",")
		isCluster := len(hosts) > 1 || (len(queryCluster) > 0 && queryCluster[0] == "true")
		passwd, _ := redisUrl.User.Password()
		if !isCluster {
			this.IRedisDriver = NewRedisDriverStandalone(this.domain, this.appGroup, ":",
				&redis.Options{
					Addr:     hosts[0],
					Password: passwd, // no password set
					DB:       0,      // use default DB
				})
		} else {
			this.IRedisDriver = NewRedisDriverCluster(this.domain, this.appGroup, ":",
				&redis.ClusterOptions{
					Addrs:    hosts,
					Password: passwd, // no password set
				})
		}
	} else if this.standaloneOption != nil {
		this.IRedisDriver = NewRedisDriverStandalone(this.domain, this.appGroup, ":", this.standaloneOption)
	} else if this.clusterOption != nil {
		this.IRedisDriver = NewRedisDriverCluster(this.domain, this.appGroup, ":", this.clusterOption)
	}
	this.tagProcessor = NewRedisTagProcessor(this.IRedisDriver)
}
