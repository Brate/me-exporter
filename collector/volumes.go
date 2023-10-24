package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type volumes struct {
	meSession                 *MeMetrics
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
	volumeQualifier           descMétrica
	blocks                    descMétrica
	progress                  descMétrica
	largeVirtualExtents       descMétrica
	raidtype                  descMétrica
	csCopyDest                descMétrica
	csCopySrc                 descMétrica
	csPrimary                 descMétrica
	csSecondary               descMétrica
	health                    descMétrica
	logger                    log.Logger
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
				"Volume is up", []string{"volume name", "virtual disk name", "storage pool name", "volume name", "snapshot", "wwn", "container name", "container serial", "virtual disk serial", "capabilities", "volume parent", "snap pool", "cs replication role", "volume group", "group key"}),
		},
		size: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "size"),
				"Volume size in GB", []string{"volume name", "size"}),
		},
		totalSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "total size"),
				"Volume total size in GB", []string{"volume name", "total size"}),
		},
		allocatedSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "allocated size"),
				"Volume allocated size in GB", []string{"volume name", "allocated size"}),
		},
		storageType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "storage type"),
				"Volume storage type", []string{"volume name", "storage type"}),
		},
		preferredOwner: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "preferred owner"),
				"Volume preferred owner", []string{"volume name", "preferred owner"}),
		},
		owner: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "owner"),
				"Volume owner", []string{"volume name", "owner"}),
		},
		writePolicy: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "write policy"),
				"Volume write policy", []string{"volume name", "write policy"}),
		},
		cacheOptimization: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "cache optimization"),
				"Volume cache optimization", []string{"volume name", "cache optimization"}),
		},
		readAheadSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "read ahead size"),
				"Volume read ahead size", []string{"volume name", "read ahead size"}),
		},
		volumeType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "volume type"),
				"Volume type", []string{"volume name", "volume type"}),
		},
		volumeClass: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "volume class"),
				"Volume class", []string{"volume name", "volume class"}),
		},
		tierAffinity: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "tier affinity"),
				"Volume tier affinity", []string{"volume name", "tier affinity"}),
		},
		snapshotRetentionPriority: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "snapshot retention priority"),
				"Volume snapshot retention priority", []string{"volume name", "snapshot retention priority"}),
		},
		volumeQualifier: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "volume qualifier"),
				"Volume qualifier", []string{"volume name", "volume qualifier"}),
		},
		blocks: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "blocks"),
				"Volume blocks", []string{"volume name"}),
		},
		progress: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "progress"),
				"Volume progress", []string{"volume name", "progress"}),
		},
		largeVirtualExtents: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "large virtual extents"),
				"Volume large virtual extents", []string{"volume name", "large virtual extents"}),
		},
		raidtype: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "raidType"),
				"Volume raid type", []string{"volume name", "raidType"}),
		},
		csCopyDest: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "cs copy dest"),
				"Volume cs copy dest", []string{"volume name", "cs copy dest"}),
		},
		csCopySrc: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "cs copy src"),
				"Volume cs copy src", []string{"volume name", "cs copy src"}),
		},
		csPrimary: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "cs primary"),
				"Volume cs primary", []string{"volume name", "cs primary"}),
		},
		csSecondary: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "cs secondary"),
				"Volume cs secondary", []string{"volume name", "cs secondary"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "health"),
				"Volume health", []string{"volume name", "health", "health reason"}),
		},
		logger: logger,
	}, nil
}

func (v volumes) Update(ch chan<- prometheus.Metric) error {
	if err := v.meSession.Volumes(); err != nil {
		return err
	}

	for _, volume := range v.meSession.volumes {
		ch <- v.up.constMetric(1, volume.VolumeName, volume.VirtualDiskName, volume.StoragePoolName, volume.VolumeName, volume.Snapshot, volume.Wwn, volume.ContainerName, volume.ContainerSerial, volume.VirtualDiskSerial, volume.Capabilities, volume.VolumeParent, volume.SnapPool, volume.CsReplicationRole, volume.VolumeGroup, volume.GroupKey)
		ch <- v.size.constMetric(float64(volume.SizeNumeric), volume.VolumeName, volume.Size)
		ch <- v.totalSize.constMetric(float64(volume.TotalSizeNumeric), volume.VolumeName, volume.TotalSize)
		ch <- v.allocatedSize.constMetric(float64(volume.AllocatedSizeNumeric), volume.VolumeName, volume.AllocatedSize)
		ch <- v.storageType.constMetric(float64(volume.StorageTypeNumeric), volume.VolumeName, volume.StorageType)
		ch <- v.preferredOwner.constMetric(float64(volume.PreferredOwnerNumeric), volume.VolumeName, volume.PreferredOwner)
		ch <- v.owner.constMetric(float64(volume.OwnerNumeric), volume.VolumeName, volume.Owner)
		ch <- v.writePolicy.constMetric(float64(volume.WritePolicyNumeric), volume.VolumeName, volume.WritePolicy)
		ch <- v.cacheOptimization.constMetric(float64(volume.CacheOptimizationNumeric), volume.VolumeName, volume.CacheOptimization)
		ch <- v.readAheadSize.constMetric(float64(volume.ReadAheadSizeNumeric), volume.VolumeName, volume.ReadAheadSize)
		ch <- v.volumeType.constMetric(float64(volume.VolumeTypeNumeric), volume.VolumeName, volume.VolumeType)
		ch <- v.volumeClass.constMetric(float64(volume.VolumeClassNumeric), volume.VolumeName, volume.VolumeClass)
		ch <- v.tierAffinity.constMetric(float64(volume.TierAffinityNumeric), volume.VolumeName, volume.TierAffinity)
		ch <- v.snapshotRetentionPriority.constMetric(float64(volume.SnapshotRetentionPriorityNumeric), volume.VolumeName, volume.SnapshotRetentionPriority)
		ch <- v.volumeQualifier.constMetric(float64(volume.VolumeQualifierNumeric), volume.VolumeName, volume.VolumeQualifier)
		ch <- v.blocks.constMetric(float64(volume.Blocks), volume.VolumeName)
		ch <- v.progress.constMetric(float64(volume.ProgressNumeric), volume.VolumeName, volume.Progress)
		ch <- v.largeVirtualExtents.constMetric(float64(volume.LargeVirtualExtentsNumeric), volume.VolumeName, volume.LargeVirtualExtents)
		ch <- v.raidtype.constMetric(float64(volume.RaidtypeNumeric), volume.VolumeName, volume.Raidtype)
		ch <- v.csCopyDest.constMetric(float64(volume.CsCopyDestNumeric), volume.VolumeName, volume.CsCopyDest)
		ch <- v.csCopySrc.constMetric(float64(volume.CsCopySrcNumeric), volume.VolumeName, volume.CsCopySrc)
		ch <- v.csPrimary.constMetric(float64(volume.CsPrimaryNumeric), volume.VolumeName, volume.CsPrimary)
		ch <- v.csSecondary.constMetric(float64(volume.CsSecondaryNumeric), volume.VolumeName, volume.CsSecondary)
		ch <- v.health.constMetric(float64(volume.HealthNumeric), volume.VolumeName, volume.Health, volume.HealthReason)
	}

	return nil
}
