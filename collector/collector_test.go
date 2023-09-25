package collector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go_prometheus/config"
	"os"
	"testing"
)

func Test_me4Metrics_Login(t *testing.T) {
	mydir, _ := os.Getwd()
	fmt.Printf("path: %v \n", mydir)

	instance := config.AuthEntry{Instance: "1.2.3.4", Hash: config.Hash}
	me := NewMeMetrics()

	if err := me.Login(instance); err != nil {
		//t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
	}

	err := me.ServiceTag()
	err2 := me.CacheSettings()
	err3 := me.DiskGroupStatistics()
	err4 := me.DiskGroups()
	err5 := me.DiskStatistics()
	err6 := me.Disks()
	err7 := me.Enclosures()
	err8 := me.ExpanderStatus()
	err9 := me.Fans()
	err10 := me.Frus()
	errPool := me.Pools()
	errPoolStat := me.PoolsStatistics()
	err12 := me.Ports()
	err13 := me.SensorStatus()
	err14 := me.ControllerStatistics()
	err15 := me.Volumes()
	err16 := me.VolumeStatistics()
	err17 := me.Tiers()
	err18 := me.TierStatistics()
	err19 := me.UnwritableCache()

	assert.Nil(t, err)
	assert.Equal(t, me.serviceTag.ServiceTag, "9YHT4Z2")

	assert.Nil(t, err2)
	assert.Equal(t, me.cacheSettings.Meta, "/meta/cache-settings")

	assert.Nil(t, err3)
	assert.Equal(t, me.diskGroupsStatistics.Meta, "/meta/disk-group-statistics")

	assert.Nil(t, err4)
	assert.Equal(t, me.diskGroups.Meta, "/meta/disk-groups")

	assert.Nil(t, err5)
	assert.Equal(t, me.diskStatistic.Meta, "/meta/disk-statistics")

	assert.Nil(t, err6)
	assert.Equal(t, me.disks.Meta, "/meta/drives")

	assert.Nil(t, err7)
	assert.Equal(t, me.enclosures.Meta, "/meta/enclosures")

	assert.Nil(t, err8)
	assert.Equal(t, me.expanderStatus.Meta, "/meta/sas-status-controller-a")

	assert.Nil(t, err9)
	assert.Equal(t, me.fans.Meta, "/meta/fan")

	assert.Nil(t, err10)
	assert.Equal(t, me.frus.Meta, "/meta/enclosure-fru")

	assert.Nil(t, errPool)
	assert.Equal(t, me.Pools(), "/meta/pools")

	assert.Nil(t, errPoolStat)
	assert.Equal(t, me.PoolsStatistics(), "/meta/pool-statistics")

	assert.Nil(t, err12)
	assert.Equal(t, me.ports.Meta, "/meta/port")

	assert.Nil(t, err13)
	assert.Equal(t, me.sensorStatus.Meta, "/meta/sensors")

	assert.Nil(t, err14)
	assert.Equal(t, me.controllerStatistics.Meta, "/meta/controller-statistics")

	assert.Nil(t, err15)
	assert.Equal(t, me.volumes.Meta, "/meta/volumes")

	assert.Nil(t, err16)
	assert.Equal(t, me.volumeStatistics.Meta, "/meta/volume-statistics")

	assert.Nil(t, err17)
	assert.Equal(t, me.tiers.Meta, "/meta/tiers")

	assert.Nil(t, err18)
	assert.Equal(t, me.tierStatistics.Meta, "/meta/tier-statistics")

	assert.Nil(t, err19)
	assert.Equal(t, me.unwritableCache.Meta, "/meta/unwritable-cache")

	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		_ = &MeMetrics{
	//			sessionKey:    tt.metrics.sessionKey,
	//			serviceTag:    tt.metrics.serviceTag,
	//			cacheSettings: tt.metrics.cacheSettings,
	//		}
	//	})
	//}
}
