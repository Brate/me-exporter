package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type disks struct {
	//all labels are with location
	meSession        *MeMetrics
	slot             descMétrica
	port             descMétrica
	scsiID           descMétrica
	blocksize        descMétrica
	up               descMétrica // Serial number, vendor, model, revision
	secondaryChannel descMétrica
	//containerIndex   descMétrica
	//memberIndex      descMétrica
	description   descMétrica
	architecture  descMétrica
	diskInterface descMétrica
	//singlePorted     descMétrica
	diskType   descMétrica
	usage      descMétrica
	jobRunning descMétrica
	//blink            descMétrica
	//locatorLed       descMétrica
	speed    descMétrica
	smart    descMétrica
	dualPort descMétrica
	//error         descMétrica
	driveDownCode descMétrica
	owner         descMétrica
	index         descMétrica
	rpm           descMétrica
	size          descMétrica
	sectorFormat  descMétrica
	transferRate  descMétrica
	//attributes    descMétrica
	//reconState descMétrica

	// TODO: Testar no ME4
	storageTier descMétrica //Me5?
	ssdLifeLeft descMétrica //Me5?
	//ledStatus            descMétrica //Me5?
	diskDsdCount         descMétrica
	spunDown             descMétrica
	IOCount              descMétrica
	totalDataTransferred descMétrica
	avgRspTime           descMétrica
	fdeState             descMétrica
	fdeConfigTime        descMétrica
	temperature          descMétrica
	temperatureStatus    descMétrica
	powerOnHours         descMétrica
	extendedStatus       descMétrica
	health               descMétrica
	logger               log.Logger
	//FcPChannel                  descMétrica //Fc P1 and P2
	//FcPDeviceID                 descMétrica //Fc P1 and P2
	//FcPUnitNumber               descMétrica //Fc P1 and P2
}

func identifyCodes(mask uint64) map[uint64]string {

	var codeMap = map[uint64]string{
		0x00000000:  "Ok",
		0x00000001:  "Single-pathed, A down",
		0x00000002:  "SSD exhausted",
		0x00000004:  "Degraded warning",
		0x00000008:  "Spun down",
		0x00000010:  "Downed by user",
		0x00000020:  "Reconstruction failed",
		0x00000040:  "Leftover, no reason",
		0x00000080:  "Previously missing",
		0x00000100:  "Medium error",
		0x00000200:  "SMART event",
		0x00000400:  "Hardware failure",
		0x00000800:  "Foreign disk unlocked",
		0x00001000:  "Non-FDE disk",
		0x00002000:  "FDE protocol failure",
		0x00004000:  "Using alternate path",
		0x00008000:  "Initialization failed",
		0x00010000:  "Unsupported type",
		0x00040000:  "Recovered errors",
		0x00080000:  "Unexpected leftover",
		0x00100000:  "Not auto-secured",
		0x00200000:  "SSD nearly exhausted",
		0x00400000:  "Degraded critical",
		0x00800000:  "Single-pathed, B down",
		0x01000000:  "Foreign disk secured",
		0x02000000:  "Foreign disk secured and locked",
		0x04000000:  "Unexpected usage",
		0x08000000:  "Enclosure fault sensed",
		0x10000000:  "Unsupported block size",
		0x20000000:  "Unsupported vendor",
		0x40000000:  "Timed-out",
		0x200000000: "Preemptive pending degraded",
	}

	result := make(map[uint64]string)

	for code, errorDesc := range codeMap {
		if mask&code != 0 {
			result[code] = errorDesc
		}
	}

	return result
}

func init() {
	registerCollector("disks", NewDisks)
}

func NewDisks(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &disks{
		meSession: me,
		slot: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "slot"),
				"Slot of the disk", []string{"disk"}),
		},
		port: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "port"),
				"Port of the disk", []string{"disk"}),
		},
		scsiID: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "scsi_id"),
				"SCSI ID of the disk on primary channel", []string{"disk"}),
		},
		secondaryChannel: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "scsi_id_secondary_channel"),
				"SCSI ID of the disk on Secondary channel", []string{"disk"}),
		},
		blocksize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "block_size"),
				"Block size of the disk", []string{"disk"}),
		},
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "up"),
				"Up of the disk", []string{"disk", "vendor", "model",
					"virtual_disk", "disk_group", "storage_pool"}),
		},
		//containerIndex: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "container_index"),
		//		"Container index of the disk", []string{"disk"}),
		//},
		//memberIndex: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "member_index"),
		//		"Member index of the disk", []string{"disk"}),
		//},
		description: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "description"),
				"Description of the disk", []string{"disk", "description"}),
		},
		architecture: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "architecture"),
				"Disk architecture", []string{"disk", "architecture"}),
		},
		diskInterface: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "interface"),
				"Disk interface type", []string{"disk"}),
		},
		//singlePorted: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "single_ported"),
		//		"Is disk single ported ?", []string{"disk"}),
		//},
		diskType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "type"),
				"Disk type", []string{"disk"}),
		},
		usage: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "usage"),
				"Usage of the disk", []string{"disk", "usage"}),
		},
		jobRunning: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "job_running"),
				"Job running of the disk", []string{"disk", "job_running", "status"}),
		},
		//blink: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "blink"),
		//		"Blink of the disk", []string{"disk"}),
		//},
		//locatorLed: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "locator_led"),
		//		"Locator led of the disk", []string{"disk"}),
		//},
		speed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "speed"),
				"Speed of the disk", []string{"disk"}),
		},
		smart: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "smart"),
				"Smart of the disk", []string{"disk", "status"}),
		},
		dualPort: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "dual_port"),
				"Dual port of the disk", []string{"disk"}),
		},
		//error: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "error"),
		//		"Error of the disk", []string{"disk"}),
		//},
		driveDownCode: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "drive_down_code"),
				"Drive down code of the disk", []string{"disk"}),
		},
		owner: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "owner"),
				"Owner of the disk", []string{"disk", "owner"}),
		},
		index: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "index"),
				"Index of the disk", []string{"disk"}),
		},
		rpm: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "rpm"),
				"Rpm of the disk", []string{"disk"}),
		},
		size: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "size"),
				"Size of the disk", []string{"disk"}),
		},
		sectorFormat: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "sector_format"),
				"Sector format of the disk", []string{"disk", "format"}),
		},
		transferRate: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "transfer_rate"),
				"Transfer rate of the disk", []string{"disk", "rate"}),
		},
		//attributes: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "attributes"),
		//		"Attributes of the disk", []string{"disk", "attributes"}),
		//},
		//reconState: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "recon_state"),
		//		"Recon state of the disk", []string{"disk", "recon_state"}),
		//},

		storageTier: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "storage_tier"),
				"Storage tier of the disk", []string{"disk", "storage_tier"}),
		},
		ssdLifeLeft: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "ssd_life_left"),
				"Ssd life left of the disk", []string{"disk", "percentage_left"}),
		},
		//ledStatus: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disks", "led_status"),
		//		"Led status of the disk", []string{"disk", "led_status"}),
		//},
		diskDsdCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disks", "disk_dsd_count"),
				"Disk dsd count of the disk", []string{"disk"}),
		},
		spunDown: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "spun_down"),
				"Spun down of the disk", []string{"disk"}),
		},
		IOCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disks", "io_count"),
				"IO counter of the disk", []string{"disk"}),
		},
		totalDataTransferred: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disks", "total_data_transferred"),
				"Total data transferred of the disk", []string{"disk"}),
		},
		avgRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "avg_rsp_time_microseconds"),
				"Avg rsp time of the disk", []string{"disk"}),
		},
		fdeState: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "fde_state"),
				"Fde state of the disk", []string{"disk", "state"}),
		},
		fdeConfigTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "fde_config_time_epoch"),
				"Fde config time of the disk", []string{"disk"}),
		},
		temperature: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "temperature"),
				"Temperature of the disk", []string{"disk"}),
		},
		temperatureStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "temperature_status"),
				"Temperature status of the disk", []string{"disk", "temperature_status"}),
		},
		powerOnHours: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "power_on_hours"),
				"Power on hours of the disk", []string{"disk"}),
		},
		extendedStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "extended_status"),
				"Extended status of the disk", []string{"disk", "status"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disks", "health"),
				"Health of the disk", []string{"disk", "health"}),
		},
		logger: logger,
	}, nil
}

func (d disks) Update(ch chan<- prometheus.Metric) error {
	if err := d.meSession.Disks(); err != nil {
		return err
	}

	for _, disk := range d.meSession.disks {
		ch <- prometheus.MustNewConstMetric(d.up.desc, d.up.tipo, 1, disk.Location,
			disk.Vendor, disk.Model, disk.VirtualDiskSerial, disk.DiskGroup, disk.StoragePoolName)
		ch <- prometheus.MustNewConstMetric(d.slot.desc, d.slot.tipo, float64(disk.Slot), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.port.desc, d.port.tipo, float64(disk.Port), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.scsiID.desc, d.scsiID.tipo, float64(disk.ScsiID), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.secondaryChannel.desc, d.secondaryChannel.tipo, float64(disk.SecondaryChannel), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.blocksize.desc, d.blocksize.tipo, float64(disk.Blocksize), disk.Location)
		//ch <- prometheus.MustNewConstMetric(d.containerIndex.desc, d.containerIndex.tipo, float64(disk.ContainerIndex), disk.Location)
		//ch <- prometheus.MustNewConstMetric(d.memberIndex.desc, d.memberIndex.tipo, float64(disk.MemberIndex), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.description.desc, d.description.tipo, float64(disk.DescriptionNumeric), disk.Location, disk.Description)
		ch <- prometheus.MustNewConstMetric(d.architecture.desc, d.architecture.tipo, float64(disk.ArchitectureNumeric), disk.Location, disk.Architecture)
		ch <- prometheus.MustNewConstMetric(d.diskInterface.desc, d.diskInterface.tipo, float64(disk.InterfaceNumeric), disk.Location)
		//ch <- prometheus.MustNewConstMetric(d.singlePorted.desc, d.singlePorted.tipo, float64(disk.SinglePortedNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.diskType.desc, d.diskType.tipo, float64(disk.TypeNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.usage.desc, d.usage.tipo, float64(disk.UsageNumeric), disk.Location, disk.Usage)
		ch <- prometheus.MustNewConstMetric(d.jobRunning.desc, d.jobRunning.tipo, float64(disk.JobRunningNumeric), disk.Location, disk.JobRunning, disk.CurrentJobCompletion)
		//ch <- prometheus.MustNewConstMetric(d.blink.desc, d.blink.tipo, float64(disk.Blink), disk.Location)
		//ch <- prometheus.MustNewConstMetric(d.locatorLed.desc, d.locatorLed.tipo, float64(disk.LocatorLedNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.speed.desc, d.speed.tipo, float64(disk.Speed), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.smart.desc, d.smart.tipo, float64(disk.SmartNumeric), disk.Location, disk.Smart)
		ch <- prometheus.MustNewConstMetric(d.dualPort.desc, d.dualPort.tipo, float64(disk.DualPort), disk.Location)
		// TODO: Incluir métricas FcP1Channel, FcP1DeviceID, FcP1NodeWwn, FcP1PortWwn, FcP1UnitNumber
		//ch <- prometheus.MustNewConstMetric(d.error.desc, d.error.tipo, float64(disk.Error), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.driveDownCode.desc, d.driveDownCode.tipo, float64(disk.DriveDownCode), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.owner.desc, d.owner.tipo, float64(disk.OwnerNumeric), disk.Location, disk.Owner)
		ch <- prometheus.MustNewConstMetric(d.index.desc, d.index.tipo, float64(disk.Index), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.rpm.desc, d.rpm.tipo, float64(disk.Rpm), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.size.desc, d.size.tipo, float64(disk.Blocks), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.sectorFormat.desc, d.sectorFormat.tipo, float64(disk.SectorFormatNumeric), disk.Location, disk.SectorFormat)
		ch <- prometheus.MustNewConstMetric(d.transferRate.desc, d.transferRate.tipo, float64(disk.Speed), disk.Location, disk.TransferRate)
		//ch <- prometheus.MustNewConstMetric(d.attributes.desc, d.attributes.tipo, float64(disk.AttributesNumeric), disk.Location, disk.Attributes)
		//ch <- prometheus.MustNewConstMetric(d.reconState.desc, d.reconState.tipo, float64(disk.ReconStateNumeric), disk.Location, disk.ReconState)
		ch <- prometheus.MustNewConstMetric(d.storageTier.desc, d.storageTier.tipo, float64(disk.StorageTierNumeric), disk.Location, disk.StorageTier)
		ch <- prometheus.MustNewConstMetric(d.ssdLifeLeft.desc, d.ssdLifeLeft.tipo, float64(disk.SsdLifeLeftNumeric), disk.Location, disk.SsdLifeLeft)
		//ch <- prometheus.MustNewConstMetric(d.ledStatus.desc, d.ledStatus.tipo, float64(disk.LedStatusNumeric), disk.Location, disk.LedStatus)

		ch <- prometheus.MustNewConstMetric(d.diskDsdCount.desc, d.diskDsdCount.tipo, float64(disk.DiskDsdCount), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.spunDown.desc, d.spunDown.tipo, float64(disk.SpunDown), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.IOCount.desc, d.IOCount.tipo, float64(disk.NumberOfIos), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.totalDataTransferred.desc, d.totalDataTransferred.tipo, float64(disk.TotalDataTransferredNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.avgRspTime.desc, d.avgRspTime.tipo, float64(disk.AvgRspTime), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.fdeState.desc, d.fdeState.tipo, float64(disk.FdeStateNumeric), disk.Location, disk.FdeState)
		ch <- prometheus.MustNewConstMetric(d.fdeConfigTime.desc, d.fdeConfigTime.tipo, float64(disk.FdeConfigTimeNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.temperature.desc, d.temperature.tipo, float64(disk.TemperatureNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.temperatureStatus.desc, d.temperatureStatus.tipo, float64(disk.TemperatureStatusNumeric), disk.Location, disk.TemperatureStatus)
		ch <- prometheus.MustNewConstMetric(d.powerOnHours.desc, d.powerOnHours.tipo, float64(disk.PowerOnHours), disk.Location)
		ch <- prometheus.MustNewConstMetric(d.health.desc, d.health.tipo, float64(disk.HealthNumeric), disk.Location, disk.Health)

		statusCodes := identifyCodes(uint64(disk.ExtendedStatus))
		for _, errorDesc := range statusCodes {
			ch <- prometheus.MustNewConstMetric(d.extendedStatus.desc, d.extendedStatus.tipo, float64(disk.ExtendedStatus), disk.Location, errorDesc)
		}
	}
	return nil
}
