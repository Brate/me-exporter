package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type volumes struct {
	meSession *MeMetrics

	up                        descMétrica
	size                      descMétrica
	totalSize                 descMétrica
	allocatedSize             descMétrica
	storageType               descMétrica
	preferredOwner            descMétrica
	owner                     descMétrica
	writePolicy               descMétrica
	cacheOptimization         descMétrica
	readAheadSize             descMétrica
	volumeType                descMétrica
	volumeClass               descMétrica
	tierAffinity              descMétrica
	snapshotRetentionPriority descMétrica
	blocks                    descMétrica
	blockSize                 descMétrica
	raidType                  descMétrica
	health                    descMétrica

	logger log.Logger
}

func init() {
	registerCollector("volumes", NewVolumesCollector)
}

func NewVolumesCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &volumes{
		meSession: me,
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "up"),
				"Volume literals", []string{"name", "wwn", "virtual_disk",
					"snapshot", "volume_parent", "snap_pool",
					"volume_group", "vg_key"}),
		},
		size: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "size_gigabytes"),
				"Volume size in GB", []string{"name"}),
		},
		totalSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "total_size_blocks"),
				"Volume total size in blocks", []string{"name"}),
		},
		allocatedSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "allocated_size_gigabytes"),
				"Volume allocated size in GB", []string{"name"}),
		},
		storageType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "storage_type"),
				"Volume storage type", []string{"name", "type"}),
		},
		preferredOwner: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "preferred_owner"),
				"Volume preferred owner", []string{"name", "preferred"}),
		},
		owner: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "owner"),
				"Volume owner", []string{"name", "owner"}),
		},
		writePolicy: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "write_policy"),
				"Volume write policy", []string{"name", "policy"}),
		},
		cacheOptimization: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "cache_optimization"),
				"Volume cache optimization", []string{"name", "optimization"}),
		},
		readAheadSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "read_ahead_size_bytes"),
				"Volume read ahead size", []string{"name", "size"}),
		},
		volumeType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "type"),
				"Volume type", []string{"name", "type"}),
		},
		volumeClass: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "class"),
				"Volume class", []string{"name", "class"}),
		},
		tierAffinity: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "tier_affinity"),
				"Volume tier affinity", []string{"name", "affinity"}),
		},
		snapshotRetentionPriority: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "snapshot_retention_priority"),
				"Volume snapshot retention priority", []string{"name", "priority"}),
		},
		blocks: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "blocks"),
				"Volume blocks", []string{"name"}),
		},
		blockSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "block_size_bytes"),
				"Volume blocks", []string{"name"}),
		},
		raidType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "raid_type"),
				"Volume raid type", []string{"name", "type"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "health"),
				"Volume health", []string{"name", "health", "reason"}),
		},
		logger: logger,
	}, nil
}

func (v volumes) Update(ch chan<- prometheus.Metric) error {
	if err := v.meSession.Volumes(); err != nil {
		return err
	}

	for _, vol := range v.meSession.volumes {
		ch <- v.up.constMetric(1, vol.VolumeName, vol.Wwn, vol.VirtualDiskName,
			vol.Snapshot, vol.VolumeParent, vol.SnapPool,
			vol.VolumeGroup, vol.GroupKey)
		ch <- v.size.constMetric(float64(vol.SizeNumeric), vol.VolumeName)
		ch <- v.totalSize.constMetric(float64(vol.TotalSizeNumeric), vol.VolumeName)
		ch <- v.allocatedSize.constMetric(float64(vol.AllocatedSizeNumeric), vol.VolumeName)
		ch <- v.storageType.constMetric(float64(vol.StorageTypeNumeric), vol.VolumeName, vol.StorageType)
		ch <- v.preferredOwner.constMetric(float64(vol.PreferredOwnerNumeric), vol.VolumeName, vol.PreferredOwner)
		ch <- v.owner.constMetric(float64(vol.OwnerNumeric), vol.VolumeName, vol.Owner)
		ch <- v.writePolicy.constMetric(float64(vol.WritePolicyNumeric), vol.VolumeName, vol.WritePolicy)
		ch <- v.cacheOptimization.constMetric(float64(vol.CacheOptimizationNumeric), vol.VolumeName, vol.CacheOptimization)
		ch <- v.readAheadSize.constMetric(float64(vol.ReadAheadSizeNumeric), vol.VolumeName, vol.ReadAheadSize)
		ch <- v.volumeType.constMetric(float64(vol.VolumeTypeNumeric), vol.VolumeName, vol.VolumeType)
		ch <- v.volumeClass.constMetric(float64(vol.VolumeClassNumeric), vol.VolumeName, vol.VolumeClass)
		ch <- v.tierAffinity.constMetric(float64(vol.TierAffinityNumeric), vol.VolumeName, vol.TierAffinity)
		ch <- v.snapshotRetentionPriority.constMetric(float64(vol.SnapshotRetentionPriorityNumeric), vol.VolumeName, vol.SnapshotRetentionPriority)
		ch <- v.blocks.constMetric(float64(vol.Blocks), vol.VolumeName)
		ch <- v.blockSize.constMetric(float64(vol.Blocksize), vol.VolumeName)
		ch <- v.raidType.constMetric(float64(vol.RaidtypeNumeric), vol.VolumeName, vol.Raidtype)
		ch <- v.health.constMetric(float64(vol.HealthNumeric), vol.VolumeName, vol.Health, vol.HealthReason)
	}

	return nil
}
