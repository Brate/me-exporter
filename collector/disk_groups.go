package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type diskGroups struct {
	meSession                 *MeMetrics
	blocksize                 descMétrica
	size                      descMétrica
	freespace                 descMétrica
	rawSize                   descMétrica
	storageType               descMétrica
	storageTier               descMétrica
	totalPages                descMétrica
	allocatedPages            descMétrica
	availablePages            descMétrica
	poolPercentage            descMétrica
	performanceRank           descMétrica
	owner                     descMétrica
	preferredOwner            descMétrica
	raidtype                  descMétrica
	diskcount                 descMétrica
	sparecount                descMétrica
	status                    descMétrica
	minDriveSize              descMétrica
	createDate                descMétrica
	currentJob                descMétrica
	numArrayPartitions        descMétrica
	largestFreePartitionSpace descMétrica
	numDrivesPerLowLevelArray descMétrica
	numExpansionPartitions    descMétrica
	numPartitionSegments      descMétrica
	newPartitionLba           descMétrica
	arrayDriveType            descMétrica
	diskDescription           descMétrica
	isJobAutoAbortable        descMétrica
	blocks                    descMétrica
	diskDsdEnableVdisk        descMétrica
	diskDsdDelayVdisk         descMétrica
	scrubDurationGoal         descMétrica
	adaptTargetSpareCapacity  descMétrica
	adaptActualSpareCapacity  descMétrica
	adaptCriticalCapacity     descMétrica
	adaptDegradedCapacity     descMétrica
	adaptLinearVolumeBoundary descMétrica
	poolSectorFormat          descMétrica
	health                    descMétrica
	logger                    log.Logger
}

func init() {
	registerCollector("disk_groups", NewDiskGroupsCollector)
}

func NewDiskGroupsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &diskGroups{
		meSession: me,
		blocksize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "blocksize"),
				"Block size of the disk group", []string{"disk_group"}),
		},
		size: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "size"),
				"Size of the disk group", []string{"disk_group"}),
		},
		freespace: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "free_space"),
				"Free space of the disk group", []string{"disk_group"}),
		},
		rawSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "raw_size"),
				"Raw size of the disk group", []string{"disk_group"}),
		},
		storageType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "storage_type"),
				"Storage type of the disk group", []string{"disk_group"}),
		},
		storageTier: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "storage_tier"),
				"Storage tier of the disk group", []string{"disk_group"}),
		},
		totalPages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "total_pages"),
				"Total pages of the disk group", []string{"disk_group"}),
		},
		allocatedPages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "allocated_pages"),
				"Allocated pages of the disk group", []string{"disk_group"}),
		},
		availablePages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "available_pages"),
				"Available pages of the disk group", []string{"disk_group"}),
		},
		poolPercentage: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "pool_percentage"),
				"Pool percentage of the disk group", []string{"disk_group"}),
		},
		performanceRank: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "performance_rank"),
				"Performance rank of the disk group", []string{"disk_group"}),
		},
		owner: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "owner"),
				"Owner of the disk group", []string{"disk_group"}),
		},
		preferredOwner: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "preferred_owner"),
				"Preferred owner of the disk group", []string{"disk_group"}),
		},
		raidtype: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "raidtype"),
				"Raid type of the disk group", []string{"disk_group"}),
		},
		diskcount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_groups", "diskcount"),
				"Disk count of the disk group", []string{"disk_group"}),
		},
		sparecount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_groups", "sparecount"),
				"Spare count of the disk group", []string{"disk_group"}),
		},
		status: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "status"),
				"Status of the disk group", []string{"disk_group"}),
		},
		minDriveSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "min_drive_size"),
				"Minimum drive size of the disk group", []string{"disk_group"}),
		},
		createDate: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "create_date"),
				"Creation date of the disk group", []string{"disk_group"}),
		},
		currentJob: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_groups", "current_job"),
				"Current job of the disk group", []string{"disk_group"}),
		},
		numArrayPartitions: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "num_array_partitions"),
				"Number of array partitions of the disk group", []string{"disk_group"}),
		},
		largestFreePartitionSpace: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "largest_free_partition_space"),
				"Largest free partition space of the disk group", []string{"disk_group"}),
		},
		numDrivesPerLowLevelArray: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "num_drives_per_low_level_array"),
				"Number of drives per low level array of the disk group", []string{"disk_group"}),
		},
		numExpansionPartitions: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "num_expansion_partitions"),
				"Number of expansion partitions of the disk group", []string{"disk_group"}),
		},
		numPartitionSegments: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "num_partition_segments"),
				"Number of partition segments of the disk group", []string{"disk_group"}),
		},
		newPartitionLba: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "new_partition_lba"),
				"New partition LBA of the disk group", []string{"disk_group"}),
		},
		arrayDriveType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "array_drive_type"),
				"Array drive type of the disk group", []string{"disk_group"}),
		},
		diskDescription: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "disk_description"),
				"Disk description of the disk group", []string{"disk_group"}),
		},
		isJobAutoAbortable: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "is_job_auto_abortable"),
				"Is job auto abortable of the disk group", []string{"disk_group"}),
		},
		blocks: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_groups", "blocks"),
				"Blocks of the disk group", []string{"disk_group"}),
		},
		diskDsdEnableVdisk: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "disk_dsd_enable_vdisk"),
				"Disk DSD enable vdisk of the disk group", []string{"disk_group"}),
		},
		diskDsdDelayVdisk: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "disk_dsd_delay_vdisk"),
				"Disk DSD delay vdisk of the disk group", []string{"disk_group"}),
		},
		scrubDurationGoal: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "scrub_duration_goal"),
				"Scrub duration goal of the disk group", []string{"disk_group"}),
		},
		adaptTargetSpareCapacity: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "adapt_target_spare_capacity"),
				"Adapt target spare capacity of the disk group", []string{"disk_group"}),
		},
		adaptActualSpareCapacity: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "adapt_actual_spare_capacity"),
				"Adapt actual spare capacity of the disk group", []string{"disk_group"}),
		},
		adaptCriticalCapacity: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "adapt_critical_capacity"),
				"Adapt critical capacity of the disk group", []string{"disk_group"}),
		},
		adaptDegradedCapacity: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "adapt_degraded_capacity"),
				"Adapt degraded capacity of the disk group", []string{"disk_group"}),
		},
		adaptLinearVolumeBoundary: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "adapt_linear_volume_boundary"),
				"Adapt linear volume boundary of the disk group", []string{"disk_group"}),
		},
		poolSectorFormat: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "pool_sector_format"),
				"Pool sector format of the disk group", []string{"disk_group"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_groups", "health"),
				"Health of the disk group", []string{"disk_group"}),
		},
		logger: logger,
	}, nil
}

func (dg *diskGroups) Update(ch chan<- prometheus.Metric) error {
	if err := dg.meSession.DiskGroups(); err != nil {
		return err
	}

	for _, dkg := range dg.meSession.diskGroups {
		ch <- prometheus.MustNewConstMetric(dg.blocksize.desc, dg.blocksize.tipo, float64(dkg.Blocksize), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.size.desc, dg.size.tipo, float64(dkg.SizeNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.freespace.desc, dg.freespace.tipo, float64(dkg.FreespaceNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.rawSize.desc, dg.rawSize.tipo, float64(dkg.RawSizeNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.storageType.desc, dg.storageType.tipo, float64(dkg.StorageTypeNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.storageTier.desc, dg.storageTier.tipo, float64(dkg.StorageTierNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.totalPages.desc, dg.totalPages.tipo, float64(dkg.TotalPages), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.allocatedPages.desc, dg.allocatedPages.tipo, float64(dkg.AllocatedPages), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.availablePages.desc, dg.availablePages.tipo, float64(dkg.AvailablePages), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.poolPercentage.desc, dg.poolPercentage.tipo, float64(dkg.PoolPercentage), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.performanceRank.desc, dg.performanceRank.tipo, float64(dkg.PerformanceRank), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.owner.desc, dg.owner.tipo, float64(dkg.OwnerNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.preferredOwner.desc, dg.preferredOwner.tipo, float64(dkg.PreferredOwnerNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.raidtype.desc, dg.raidtype.tipo, float64(dkg.RaidtypeNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.diskcount.desc, dg.diskcount.tipo, float64(dkg.Diskcount), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.sparecount.desc, dg.sparecount.tipo, float64(dkg.Sparecount), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.status.desc, dg.status.tipo, float64(dkg.StatusNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.minDriveSize.desc, dg.minDriveSize.tipo, float64(dkg.MinDriveSizeNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.createDate.desc, dg.createDate.tipo, float64(dkg.CreateDateNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.currentJob.desc, dg.currentJob.tipo, float64(dkg.CurrentJobNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.numArrayPartitions.desc, dg.numArrayPartitions.tipo, float64(dkg.NumArrayPartitions), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.largestFreePartitionSpace.desc, dg.largestFreePartitionSpace.tipo, float64(dkg.LargestFreePartitionSpaceNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.numDrivesPerLowLevelArray.desc, dg.numDrivesPerLowLevelArray.tipo, float64(dkg.NumDrivesPerLowLevelArray), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.numExpansionPartitions.desc, dg.numExpansionPartitions.tipo, float64(dkg.NumExpansionPartitions), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.numPartitionSegments.desc, dg.numPartitionSegments.tipo, float64(dkg.NumPartitionSegments), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.newPartitionLba.desc, dg.newPartitionLba.tipo, float64(dkg.NewPartitionLbaNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.arrayDriveType.desc, dg.arrayDriveType.tipo, float64(dkg.ArrayDriveTypeNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.diskDescription.desc, dg.diskDescription.tipo, float64(dkg.DiskDescriptionNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.isJobAutoAbortable.desc, dg.isJobAutoAbortable.tipo, float64(dkg.IsJobAutoAbortableNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.blocks.desc, dg.blocks.tipo, float64(dkg.Blocks), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.diskDsdEnableVdisk.desc, dg.diskDsdEnableVdisk.tipo, float64(dkg.DiskDsdEnableVdiskNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.diskDsdDelayVdisk.desc, dg.diskDsdDelayVdisk.tipo, float64(dkg.DiskDsdDelayVdisk), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.scrubDurationGoal.desc, dg.scrubDurationGoal.tipo, float64(dkg.ScrubDurationGoal), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.adaptTargetSpareCapacity.desc, dg.adaptTargetSpareCapacity.tipo, float64(dkg.AdaptTargetSpareCapacityNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.adaptActualSpareCapacity.desc, dg.adaptActualSpareCapacity.tipo, float64(dkg.AdaptActualSpareCapacityNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.adaptCriticalCapacity.desc, dg.adaptCriticalCapacity.tipo, float64(dkg.AdaptCriticalCapacityNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.adaptDegradedCapacity.desc, dg.adaptDegradedCapacity.tipo, float64(dkg.AdaptDegradedCapacityNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.adaptLinearVolumeBoundary.desc, dg.adaptLinearVolumeBoundary.tipo, float64(dkg.AdaptLinearVolumeBoundary), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.poolSectorFormat.desc, dg.poolSectorFormat.tipo, float64(dkg.PoolSectorFormatNumeric), dkg.Name)
		ch <- prometheus.MustNewConstMetric(dg.health.desc, dg.health.tipo, float64(dkg.HealthNumeric), dkg.Name)
	}

	return nil
}
