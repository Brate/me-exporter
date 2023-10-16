package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"me_exporter/Me"
	h "me_exporter/app/helpers"
)

type enclosures struct {
	//All used durable-id in your labels
	meSession                *MeMetrics
	up                       descMétrica
	numberOfCoolingsElements descMétrica
	numberOfDisks            descMétrica
	numberOfPowerSupplies    descMétrica
	status                   descMétrica
	slots                    descMétrica
	enclosurePower           descMétrica
	health                   descMétrica

	//Controllers
	controllerID         descMétrica
	controllerUp         descMétrica
	disks                descMétrica
	numberOfStoragePools descMétrica
	virtualDisks         descMétrica
	cacheMemorySize      descMétrica
	systemMemorySize     descMétrica
	controllerStatus     descMétrica
	failedOver           descMétrica
	failOverReason       descMétrica
	controllersStatus    descMétrica
	cacheLock            descMétrica
	writePolicy          descMétrica
	position             descMétrica
	redundancyMode       descMétrica
	redundancyStatus     descMétrica
	controllerHealth     descMétrica

	//Controllers NetworkParameters
	activeVersion descMétrica
	linkSpeed     descMétrica
	//duplexMode           descMétrica
	networkHealth        descMétrica
	networkPingBroadcast descMétrica

	//Controllers Port
	controller      descMétrica
	portType        descMétrica
	portStatus      descMétrica
	actualSpeed     descMétrica
	configuredSpeed descMétrica
	portHealth      descMétrica

	//Controllers Port IscsiPort
	sfpStatus  descMétrica
	sfpPresent descMétrica

	//Controllers ExpanderPorts
	//enclosureID             descMétrica
	//expanderPortsController descMétrica
	//sasPortType             descMétrica
	//sasPortIndex            descMétrica
	//expanderPortsStatus     descMétrica
	//expanderPortsHealth     descMétrica

	//Controllers CompactFlash
	compactFlashStatus descMétrica
	cacheFlush         descMétrica
	compactFlashHealth descMétrica

	//Controllers Expanders
	pathID          descMétrica
	expandersStatus descMétrica
	expandersHealth descMétrica

	//PowerSupplies
	//powerSuppliesEnclosureID descMétrica
	powerSuppliesStatus   descMétrica
	powerSuppliesPosition descMétrica
	powerSuppliesHealth   descMétrica

	//PowerSupplies Fans
	statusSes                       descMétrica
	powerSuppliesFansExtendedStatus descMétrica
	powerSuppliesFansStatus         descMétrica
	speed                           descMétrica
	powerSuppliesFansPosition       descMétrica
	powerSuppliesFansHealth         descMétrica

	logger log.Logger
}

func init() {
	registerCollector("enclosure", NewEnclosuresCollector)
}

func NewEnclosuresCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &enclosures{
		meSession: me,
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_up"),
				"Up", []string{"durable_id", "enclosure_wwn", "location",
					"rack_number", "rack_position"}),
		},
		numberOfCoolingsElements: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "coolings_elements_count"),
				"Number of coolings elements", []string{"durable_id"}),
		},
		numberOfDisks: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "disks_count"),
				"Number of installed disks", []string{"durable_id"}),
		},
		numberOfPowerSupplies: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_count"),
				"Number of power supplies", []string{"durable_id"}),
		},
		status: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "status"),
				"Status", []string{"durable_id", "status"}),
		},
		slots: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "slots"),
				"Slots", []string{"durable_id"}),
		},
		enclosurePower: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_watts"),
				"Enclosure power", []string{"durable_id"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "health"),
				"Health", []string{"durable_id", "health", "health_reason", "health_recommendation"}),
		},

		//Controllers
		controllerID: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_id"),
				"Controller ID", []string{"durable_id", "controller_id"}),
		},
		controllerUp: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "up"),
				"Up", []string{"durable_id", "controller_id", "vendor", "model", "revision"}),
		},
		disks: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "disks_count"),
				"Number of installed disks", []string{"durable_id", "controller"}),
		},
		numberOfStoragePools: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("enclosure", "storage_pools_count"),
				"Number of storage pools", []string{"durable_id"}),
		},
		virtualDisks: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("enclosure", "virtual_disks"),
				"Virtual disks", []string{"durable_id"}),
		},
		cacheMemorySize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_cache_memory_size"),
				"Cache memory size in MB", []string{"durable_id", "controller_id"}),
		},
		systemMemorySize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_memory_size"),
				"System memory size in MB", []string{"durable_id", "controller_id"}),
		},
		controllerStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_status"),
				"Controller status", []string{"durable_id", "controller_id", "status"}),
		},
		failedOver: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "failed_over"),
				"Failed over", []string{"durable_id", "controller_id", "failed_over"}),
		},
		failOverReason: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "fail_over_reason"),
				"Fail over reason", []string{"durable_id", "controller_id", "fail_over_reason"}),
		},
		controllersStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controllers_status"),
				"Controllers status", []string{"durable_id", "controller_id", "status"}),
		},
		cacheLock: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "cache_lock"),
				"Cache lock", []string{"durable_id", "controller_id", "cache_lock"}),
		},
		writePolicy: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_write_policy"),
				"Write policy", []string{"durable_id", "controller_id", "write_policy"}),
		},
		position: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "position"),
				"Position", []string{"durable_id", "controller_id", "position"}),
		},
		redundancyMode: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "redundancy_mode"),
				"Redundancy mode", []string{"durable_id", "controller_id", "mode"}),
		},
		redundancyStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "redundancy_status"),
				"Redundancy status", []string{"durable_id", "controller_id", "status"}),
		},
		controllerHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_health"),
				"Controller health", []string{"durable_id", "controller_id", "reason", "recommendation"}),
		},

		//Controllers NetworkParameters
		activeVersion: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "active_version"),
				"Active version", []string{"durable_id", "controller_id"}),
		},
		linkSpeed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "link_speed"),
				"Link speed in Mbps", []string{"durable_id", "controller_id", "link_speed"}),
		},
		networkHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "network_health"),
				"Network health", []string{"durable_id", "controller_id", "network_health", "network_health_recommendation"}),
		},
		networkPingBroadcast: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "network_ping_broadcast"),
				"Network ping broadcast", []string{"durable_id", "controller_id", "network_ping_broadcast"}),
		},
		//Controllers Port
		//controller: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "controller_port"),
		//		"Controller", []string{"durable_id", "controller_id", "port"}),
		//},
		portType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "port_type"),
				"Port type", []string{"durable_id", "controller_id", "port_type", "port"}),
		},
		portStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "port_status"),
				"Port status", []string{"durable_id", "controller_id", "port_status", "port"}),
		},
		actualSpeed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "actual_speed"),
				"Actual speed in GB", []string{"durable_id", "controller_id", "actual_speed", "port"}),
		},
		configuredSpeed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "configured_speed"),
				"Configured speed", []string{"durable_id", "controller_id", "speed", "port"}),
		},
		portHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "port_health"),
				"Port health", []string{"durable_id", "controller_id", "port_health", "health_reason", "health_recommendation", "port"}),
		},

		//Controllers Port IscsiPort
		sfpStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "sfp_status"),
				"SFP status", []string{"durable_id", "controller_id", "port", "sfp_status"}),
		},
		sfpPresent: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "sfp_present"),
				"SFP present", []string{"durable_id", "controller_id", "port", "status"}),
		},

		//Controllers ExpanderPorts
		//
		//enclosureID: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "enclosure_id"),
		//		"Enclosure ID", []string{"durable_id", "name", "controller"}),
		//},
		//expanderPortsController: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "expander_ports_controller"),
		//		"Expander ports controller", []string{"durable_id", "name", "controller", "expander_ports_controller"}),
		//},
		//sasPortType: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "sas_port_type"),
		//		"SAS port type", []string{"durable_id", "controller", "sas_port_type", "name"}),
		//},
		//sasPortIndex: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "sas_port_index"),
		//		"SAS port index", []string{"durable_id", "controller", "name"}),
		//},
		//expanderPortsStatus: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "expander_ports_status"),
		//		"Expander ports status", []string{"durable_id", "controller", "expander_ports_status", "name"}),
		//},
		//expanderPortsHealth: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "expander_ports_health"),
		//		"Expander ports health", []string{"durable_id", "controller", "expander_ports_health", "health_reason", "health_recommendation", "name"}),
		//},

		//Controllers CompactFlash
		compactFlashStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "compact_flash_status"),
				"Compact flash status", []string{"durable_id", "controller", "status", "name"}),
		},
		cacheFlush: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "cache_flush"),
				"Cache flush", []string{"durable_id", "controller", "cache_flush", "compact_flash_name"}),
		},
		compactFlashHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "compact_flash_health"),
				"Compact flash health", []string{"durable_id", "controller", "name", "health",
					"health_reason", "health_recommendation"}),
		},
		//Controllers Expanders
		pathID: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "path_id"),
				"Path ID", []string{"durable_id", "controller_id", "path_id", "expanders_name"}),
		},
		expandersStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "expanders_status"),
				"Expanders status", []string{"durable_id", "controller_id", "expanders_status", "expanders_name"}),
		},
		expandersHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "expanders_health"),
				"Expanders health", []string{"durable_id", "controller_id", "expanders_health", "expanders_health_reason", "expanders_health_recommendation", "expanders_name"}),
		},
		//PowerSupplies
		//powerSuppliesEnclosureID: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("enclosure", "power_supplies_enclosure_id"),
		//		"Power supplies enclosure ID", []string{"durable_id", "name"}),
		//},
		powerSuppliesStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_status"),
				"Power supplies status", []string{"durable_id", "name", "power_supplies_status"}),
		},
		powerSuppliesPosition: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_position"),
				"Power supplies position", []string{"durable_id", "name", "power_supplies_position"}),
		},
		powerSuppliesHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_health"),
				"Power supplies health", []string{"durable_id", "name", "power_supplies_health",
					"power_supplies_health_reason", "power_supplies_health_recommendation"}),
		},
		//PowerSupplies Fans
		statusSes: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "psu_fan_ses_status"),
				"Status SES", []string{"durable_id", "name", "status"}),
		},
		powerSuppliesFansExtendedStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fan_extended_status"),
				"Power supplies fans extended status", []string{"durable_id", "name", "status"}),
		},
		powerSuppliesFansStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fan_status"),
				"Power supplies fan status", []string{"durable_id", "name", "status"}),
		},
		speed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "psu_fan_speed_rpm"),
				"Speed", []string{"durable_id", "name"}),
		},
		powerSuppliesFansPosition: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fan_position"),
				"Power supplies fans position", []string{"durable_id", "name", "position"}),
		},
		powerSuppliesFansHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fan_health"),
				"Power supplies fans health", []string{"durable_id", "name", "health",
					"health_reason"}),
		},
		logger: logger,
	}, nil
}

func (e enclosures) Update(ch chan<- prometheus.Metric) error {
	if err := e.meSession.Enclosures(); err != nil {
		return err
	}

	is := h.IntToString
	sf := h.StringToFloat

	for _, enc := range e.meSession.enclosures {
		ch <- e.up.constMetric(1, enc.DurableID, enc.EnclosureWwn, enc.Location, is(enc.RackNumber), is(enc.RackPosition))
		ch <- e.numberOfCoolingsElements.constMetric(float64(enc.NumberOfCoolingsElements), enc.DurableID)
		ch <- e.numberOfDisks.constMetric(float64(enc.NumberOfDisks), enc.DurableID)
		ch <- e.numberOfPowerSupplies.constMetric(float64(enc.NumberOfPowerSupplies), enc.DurableID)
		ch <- e.status.constMetric(float64(enc.StatusNumeric), enc.DurableID, enc.Status)
		ch <- e.slots.constMetric(float64(enc.Slots), enc.DurableID)
		ch <- e.enclosurePower.constMetric(sf(enc.EnclosurePower), enc.DurableID)
		ch <- e.health.constMetric(float64(enc.HealthNumeric), enc.DurableID, enc.Health, enc.HealthReason, enc.HealthRecommendation)
		e.collectControllers(ch, enc)
		e.collectPowerSupplies(ch, enc)
	}

	return nil
}

func (d *descMétrica) constMetric(value float64, labels ...string) prometheus.Metric {
	return prometheus.MustNewConstMetric(d.desc, d.tipo, value, labels...)
}

func (e enclosures) collectControllers(ch chan<- prometheus.Metric, enc Me.Enclosure) {
	for _, controller := range enc.Controllers {
		ch <- e.controllerID.constMetric(float64(controller.ControllerIDNumeric), enc.DurableID, controller.ControllerID)
		ch <- e.controllerUp.constMetric(1, enc.DurableID, controller.ControllerID, controller.Vendor, controller.Model, controller.Revision)
		ch <- e.disks.constMetric(float64(controller.Disks), enc.DurableID, controller.ControllerID)
		ch <- e.numberOfStoragePools.constMetric(float64(controller.NumberOfStoragePools), enc.DurableID)
		ch <- e.virtualDisks.constMetric(float64(controller.VirtualDisks), enc.DurableID)
		ch <- e.cacheMemorySize.constMetric(float64(controller.CacheMemorySize), enc.DurableID, controller.ControllerID)
		ch <- e.systemMemorySize.constMetric(float64(controller.SystemMemorySize), enc.DurableID, controller.ControllerID)
		ch <- e.controllerStatus.constMetric(float64(controller.StatusNumeric), enc.DurableID, controller.ControllerID, controller.Status)
		ch <- e.failedOver.constMetric(float64(controller.FailedOverNumeric), enc.DurableID, controller.ControllerID, controller.FailedOver)
		ch <- e.failOverReason.constMetric(float64(controller.FailOverReasonNumeric), enc.DurableID, controller.ControllerID, controller.FailOverReason)
		ch <- e.controllersStatus.constMetric(float64(controller.StatusNumeric), enc.DurableID, controller.ControllerID, controller.Status)
		ch <- e.cacheLock.constMetric(float64(controller.CacheLockNumeric), enc.DurableID, controller.ControllerID, controller.CacheLock)
		ch <- e.writePolicy.constMetric(float64(controller.WritePolicyNumeric), enc.DurableID, controller.ControllerID, controller.WritePolicy)
		ch <- e.position.constMetric(float64(controller.PositionNumeric), enc.DurableID, controller.ControllerID, controller.Position)
		ch <- e.redundancyMode.constMetric(float64(controller.RedundancyModeNumeric), enc.DurableID, controller.ControllerID, controller.RedundancyMode)
		ch <- e.redundancyStatus.constMetric(float64(controller.RedundancyStatusNumeric), enc.DurableID, controller.ControllerID, controller.RedundancyStatus)
		ch <- e.controllerHealth.constMetric(float64(controller.HealthNumeric), enc.DurableID, controller.ControllerID, controller.HealthReason, controller.HealthRecommendation)

		// Helpers
		e.collectControllerNetworkParameters(ch, enc, controller)
		e.collectControllerPorts(ch, enc, controller)
		//e.collectControllerExpanderPorts(ch, enc, controller)
		e.collectControllerCompactFlash(ch, enc, controller)
		e.collectControllerExpanders(ch, enc, controller)

	}
}

// Helpers
func (e enclosures) collectPowerSupplies(ch chan<- prometheus.Metric, enc Me.Enclosure) {
	for _, psu := range enc.PowerSupplies {
		ch <- e.powerSuppliesStatus.constMetric(float64(psu.StatusNumeric), enc.DurableID, psu.Name, psu.Status)
		ch <- e.powerSuppliesPosition.constMetric(float64(psu.PositionNumeric), enc.DurableID, psu.Name, psu.Position)
		ch <- e.powerSuppliesHealth.constMetric(float64(psu.HealthNumeric), enc.DurableID, psu.Name, psu.Health, psu.HealthReason, psu.HealthRecommendation)

		// Fans
		for _, fan := range psu.Fan {
			ch <- e.statusSes.constMetric(float64(fan.StatusSesNumeric), enc.DurableID, psu.Name, fan.StatusSes)
			ch <- e.powerSuppliesFansExtendedStatus.constMetric(1, enc.DurableID, psu.Name, fan.ExtendedStatus)
			ch <- e.powerSuppliesFansStatus.constMetric(float64(fan.StatusNumeric), enc.DurableID, psu.Name, fan.Status)
			ch <- e.speed.constMetric(float64(fan.Speed), enc.DurableID, psu.Name)
			ch <- e.powerSuppliesFansPosition.constMetric(float64(fan.PositionNumeric), enc.DurableID, psu.Name, fan.Position)
			ch <- e.powerSuppliesFansHealth.constMetric(float64(fan.HealthNumeric), enc.DurableID, psu.Name, fan.Health, fan.HealthReason)
		}
	}
}
func (e enclosures) collectControllerNetworkParameters(ch chan<- prometheus.Metric, enc Me.Enclosure, controller Me.Controllers) {
	for _, networkParameters := range controller.NetworkParameters {
		ch <- e.activeVersion.constMetric(float64(networkParameters.ActiveVersion), enc.DurableID, controller.ControllerID)
		ch <- e.linkSpeed.constMetric(float64(networkParameters.LinkSpeedNumeric), enc.DurableID, controller.ControllerID, networkParameters.LinkSpeed)
		//ch <- e.duplexMode.constMetric(float64(networkParameters.DuplexModeNumeric), enc.DurableID, controller.ControllerID, networkParameters.DuplexMode)
		ch <- e.networkHealth.constMetric(float64(networkParameters.HealthNumeric), enc.DurableID, controller.ControllerID, networkParameters.Health, networkParameters.HealthRecommendation)
		ch <- e.networkPingBroadcast.constMetric(float64(networkParameters.PingBroadcastNumeric), enc.DurableID, controller.ControllerID, networkParameters.PingBroadcast)
	}
}
func (e enclosures) collectControllerPorts(ch chan<- prometheus.Metric, enc Me.Enclosure, controller Me.Controllers) {
	for _, port := range controller.Port {
		//ch <- e.controller.constMetric(float64(port.ControllerNumeric), enc.DurableID, controller.ControllerID, port.Port)
		ch <- e.portType.constMetric(float64(port.PortTypeNumeric), enc.DurableID, controller.ControllerID, port.PortType, port.Port)
		ch <- e.portStatus.constMetric(float64(port.StatusNumeric), enc.DurableID, controller.ControllerID, port.Status, port.Port)
		ch <- e.actualSpeed.constMetric(float64(port.ActualSpeedNumeric), enc.DurableID, controller.ControllerID, port.ActualSpeed, port.Port)
		ch <- e.configuredSpeed.constMetric(float64(port.ConfiguredSpeedNumeric), enc.DurableID, controller.ControllerID, port.ConfiguredSpeed, port.Port)
		ch <- e.portHealth.constMetric(float64(port.HealthNumeric), enc.DurableID, controller.ControllerID, port.Health, port.HealthReason, port.HealthRecommendation, port.Port)

		//Controllers Port IscsiPort
		for _, iscsiPort := range port.IscsiPort {
			ch <- e.sfpStatus.constMetric(float64(iscsiPort.SfpStatusNumeric), enc.DurableID, controller.ControllerID, port.Port, iscsiPort.SfpStatus)
			ch <- e.sfpPresent.constMetric(float64(iscsiPort.SfpPresentNumeric), enc.DurableID, controller.ControllerID, port.Port, iscsiPort.SfpPresent)
		}
	}
}

//	func (e enclosures) collectControllerExpanderPorts(ch chan<- prometheus.Metric, enc Me.Enclosure, controller Me.Controllers) {
//		for _, expanderPorts := range controller.ExpanderPorts {
//			ch <- e.enclosureID.constMetric(float64(expanderPorts.EnclosureID), enc.DurableID, expanderPorts.Name, controller.ControllerID)
//			ch <- e.expanderPortsController.constMetric(float64(expanderPorts.ControllerNumeric), enc.DurableID, expanderPorts.Name, controller.ControllerID, expanderPorts.Controller)
//			ch <- e.sasPortType.constMetric(float64(expanderPorts.SasPortTypeNumeric), enc.DurableID, controller.ControllerID, expanderPorts.SasPortType, expanderPorts.Name)
//			ch <- e.sasPortIndex.constMetric(float64(expanderPorts.SasPortIndex), enc.DurableID, controller.ControllerID, expanderPorts.Name)
//			ch <- e.expanderPortsStatus.constMetric(float64(expanderPorts.StatusNumeric), enc.DurableID, controller.ControllerID, expanderPorts.Status, expanderPorts.Name)
//			ch <- e.expanderPortsHealth.constMetric(float64(expanderPorts.HealthNumeric), enc.DurableID, controller.ControllerID, expanderPorts.Health, expanderPorts.HealthReason, expanderPorts.HealthRecommendation, expanderPorts.Name)
//		}
//	}
func (e enclosures) collectControllerCompactFlash(ch chan<- prometheus.Metric, enc Me.Enclosure, controller Me.Controllers) {
	for _, compactFlash := range controller.CompactFlash {
		ch <- e.compactFlashStatus.constMetric(float64(compactFlash.StatusNumeric), enc.DurableID, controller.ControllerID, compactFlash.Status, compactFlash.Name)
		ch <- e.cacheFlush.constMetric(float64(compactFlash.CacheFlushNumeric), enc.DurableID, controller.ControllerID, compactFlash.CacheFlush, compactFlash.Name)
		ch <- e.compactFlashHealth.constMetric(float64(compactFlash.HealthNumeric), enc.DurableID, controller.ControllerID, compactFlash.Name, compactFlash.Health, compactFlash.HealthReason, compactFlash.HealthRecommendation)
	}
}
func (e enclosures) collectControllerExpanders(ch chan<- prometheus.Metric, enc Me.Enclosure, controller Me.Controllers) {
	for _, expanders := range controller.Expanders {
		ch <- e.pathID.constMetric(float64(expanders.PathIDNumeric), enc.DurableID, controller.ControllerID, expanders.PathID, expanders.Name)
		ch <- e.expandersStatus.constMetric(float64(expanders.StatusNumeric), enc.DurableID, controller.ControllerID, expanders.Status, expanders.Name)
		ch <- e.expandersHealth.constMetric(float64(expanders.HealthNumeric), enc.DurableID, controller.ControllerID, expanders.Health, expanders.HealthReason, expanders.HealthRecommendation, expanders.Name)
	}
}
