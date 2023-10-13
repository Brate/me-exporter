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
	activeVersion        descMétrica
	linkSpeed            descMétrica
	duplexMode           descMétrica
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
	enclosureID             descMétrica
	expanderPortsController descMétrica
	sasPortType             descMétrica
	sasPortIndex            descMétrica
	expanderPortsStatus     descMétrica
	expanderPortsHealth     descMétrica

	//Controllers CompactFlash
	compactFlashStatus descMétrica
	cacheFlush         descMétrica
	compactFlashHealth descMétrica

	//Controllers Expanders
	pathID          descMétrica
	expandersStatus descMétrica
	expandersHealth descMétrica

	//PowerSupplies
	powerSuppliesEnclosureID descMétrica
	powerSuppliesStatus      descMétrica
	powerSuppliesPosition    descMétrica
	powerSuppliesHealth      descMétrica

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
				NomeMetrica("enclosure", "up"),
				"Up", []string{"durable_id", "up", "enclosure_wwn", "location", "rack_number", "rack_position"}),
		},
		numberOfCoolingsElements: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "coolings_elements_count"),
				"Number of coolings elements", []string{"durable_id"}),
		},
		numberOfDisks: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "disks_count"),
				"Number of disks", []string{"durable_id"}),
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
				NomeMetrica("enclosure", "disks"),
				"Disks", []string{"durable_id"}),
		},
		numberOfStoragePools: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("enclosure", "number_of_storage_pools"),
				"Number of storage pools", []string{"durable_id"}),
		},
		virtualDisks: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("enclosure", "virtual_disks"),
				"Virtual disks", []string{"durable_id"}),
		},
		cacheMemorySize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "cache_memory_size"),
				"Cache memory size in MB", []string{"durable_id"}),
		},
		systemMemorySize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "system_memory_size"),
				"System memory size in MB", []string{"durable_id"}),
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
				NomeMetrica("enclosure", "write_policy"),
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
				"Redundancy mode", []string{"durable_id", "controller_id", "redundancy_mode"}),
		},
		redundancyStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "redundancy_status"),
				"Redundancy status", []string{"durable_id", "controller_id", "redundancy_status"}),
		},
		controllerHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller_health"),
				"Controller health", []string{"durable_id", "controller_id", "controller_health_reason", "controller_health_recommendation"}),
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
		duplexMode: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "duplex_mode"),
				"Duplex mode", []string{"durable_id", "controller_id", "duplex_mode"}),
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
		controller: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "controller"),
				"Controller", []string{"durable_id", "controller_id", "controller", "port"}),
		},
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
				"Configured speed", []string{"durable_id", "controller_id", "configured_speed", "port"}),
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
				"SFP status", []string{"durable_id", "controller_id", "sfp_status", "port"}),
		},
		sfpPresent: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "sfp_present"),
				"SFP present", []string{"durable_id", "controller_id", "sfp_present", "port"}),
		},
		//Controllers ExpanderPorts
		enclosureID: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "enclosure_id"),
				"Enclosure ID", []string{"durable_id", "name", "controller"}),
		},
		expanderPortsController: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "expander_ports_controller"),
				"Expander ports controller", []string{"durable_id", "name", "controller", "expander_ports_controller"}),
		},
		sasPortType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "sas_port_type"),
				"SAS port type", []string{"durable_id", "controller", "sas_port_type", "name"}),
		},
		sasPortIndex: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "sas_port_index"),
				"SAS port index", []string{"durable_id", "controller", "name"}),
		},
		expanderPortsStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "expander_ports_status"),
				"Expander ports status", []string{"durable_id", "controller", "expander_ports_status", "name"}),
		},
		expanderPortsHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "expander_ports_health"),
				"Expander ports health", []string{"durable_id", "controller", "expander_ports_health", "health_reason", "health_recommendation", "name"}),
		},
		//Controllers CompactFlash
		compactFlashStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "compact_flash_status"),
				"Compact flash status", []string{"durable_id", "controller", "compact_flash_status", "compact_flash_name"}),
		},
		cacheFlush: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "cache_flush"),
				"Cache flush", []string{"durable_id", "controller", "cache_flush", "compact_flash_name"}),
		},
		compactFlashHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "compact_flash_health"),
				"Compact flash health", []string{"durable_id", "controller", "compact_flash_name", "compact_flash_health",
					"compact_flash_health_reason", "compact_flash_health_recommendation"}),
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
		powerSuppliesEnclosureID: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_enclosure_id"),
				"Power supplies enclosure ID", []string{"durable_id", "name"}),
		},
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
				NomeMetrica("enclosure", "status_ses"),
				"Status SES", []string{"durable_id", "name", "status_ses"}),
		},
		powerSuppliesFansExtendedStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fans_extended_status"),
				"Power supplies fans extended status", []string{"durable_id", "name", "power_supplies_fans_extended_status"}),
		},
		powerSuppliesFansStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fans_status"),
				"Power supplies fans status", []string{"durable_id", "name", "power_supplies_fans_status"}),
		},
		speed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "speed"),
				"Speed", []string{"durable_id", "name"}),
		},
		powerSuppliesFansPosition: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fans_position"),
				"Power supplies fans position", []string{"durable_id", "name", "power_supplies_fans_position"}),
		},
		powerSuppliesFansHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("enclosure", "power_supplies_fans_health"),
				"Power supplies fans health", []string{"durable_id", "name", "power_supplies_fans_health",
					"power_supplies_fans_health_reason", "power_supplies_fans_health_recommendation"}),
		},
		logger: logger,
	}, nil
}

func (e enclosures) Update(ch chan<- prometheus.Metric) error {
	if err := e.meSession.Enclosures(); err != nil {
		return err
	}

	for _, enc := range e.meSession.enclosures {
		ch <- prometheus.MustNewConstMetric(e.up.desc, e.up.tipo, 1, enc.DurableID, enc.EnclosureWwn, enc.Location, h.IntToString(enc.RackNumber), h.IntToString(enc.RackPosition))
		ch <- prometheus.MustNewConstMetric(e.numberOfCoolingsElements.desc, e.numberOfCoolingsElements.tipo, float64(enc.NumberOfCoolingsElements), enc.DurableID)
		ch <- prometheus.MustNewConstMetric(e.numberOfDisks.desc, e.numberOfDisks.tipo, float64(enc.NumberOfDisks), enc.DurableID)
		ch <- prometheus.MustNewConstMetric(e.numberOfPowerSupplies.desc, e.numberOfPowerSupplies.tipo, float64(enc.NumberOfPowerSupplies), enc.DurableID)
		ch <- prometheus.MustNewConstMetric(e.status.desc, e.status.tipo, float64(enc.StatusNumeric), enc.DurableID, enc.Status)
		ch <- prometheus.MustNewConstMetric(e.slots.desc, e.slots.tipo, float64(enc.Slots), enc.DurableID)
		ch <- prometheus.MustNewConstMetric(e.enclosurePower.desc, e.enclosurePower.tipo, h.StringToFloat(enc.EnclosurePower), enc.DurableID)
		ch <- prometheus.MustNewConstMetric(e.health.desc, e.health.tipo, float64(enc.HealthNumeric), enc.DurableID, enc.Health, enc.HealthReason, enc.HealthRecommendation)

		//Controllers
		for _, controller := range enc.Controllers {
			ch <- prometheus.MustNewConstMetric(e.controllerID.desc, e.controllerID.tipo, float64(controller.ControllerIDNumeric), enc.DurableID, controller.ControllerID)
			ch <- prometheus.MustNewConstMetric(e.controllerUp.desc, e.controllerUp.tipo, 1, enc.DurableID, controller.ControllerID, controller.Vendor, controller.Model, controller.Revision)
			ch <- prometheus.MustNewConstMetric(e.disks.desc, e.disks.tipo, float64(controller.Disks), enc.DurableID)
			ch <- prometheus.MustNewConstMetric(e.numberOfStoragePools.desc, e.numberOfStoragePools.tipo, float64(controller.NumberOfStoragePools), enc.DurableID)
			ch <- prometheus.MustNewConstMetric(e.virtualDisks.desc, e.virtualDisks.tipo, float64(controller.VirtualDisks), enc.DurableID)
			ch <- prometheus.MustNewConstMetric(e.cacheMemorySize.desc, e.cacheMemorySize.tipo, float64(controller.CacheMemorySize), enc.DurableID)
			ch <- prometheus.MustNewConstMetric(e.systemMemorySize.desc, e.systemMemorySize.tipo, float64(controller.SystemMemorySize), enc.DurableID)
			ch <- prometheus.MustNewConstMetric(e.controllerStatus.desc, e.controllerStatus.tipo, float64(controller.StatusNumeric), enc.DurableID, controller.ControllerID, controller.Status)
			ch <- prometheus.MustNewConstMetric(e.failedOver.desc, e.failedOver.tipo, float64(controller.FailedOverNumeric), enc.DurableID, controller.ControllerID, controller.FailedOver)
			ch <- prometheus.MustNewConstMetric(e.failOverReason.desc, e.failOverReason.tipo, float64(controller.FailOverReasonNumeric), enc.DurableID, controller.ControllerID, controller.FailOverReason)
			ch <- prometheus.MustNewConstMetric(e.controllersStatus.desc, e.controllersStatus.tipo, float64(controller.StatusNumeric), enc.DurableID, controller.ControllerID, controller.Status)
			ch <- prometheus.MustNewConstMetric(e.cacheLock.desc, e.cacheLock.tipo, float64(controller.CacheLockNumeric), enc.DurableID, controller.ControllerID, controller.CacheLock)
			ch <- prometheus.MustNewConstMetric(e.writePolicy.desc, e.writePolicy.tipo, float64(controller.WritePolicyNumeric), enc.DurableID, controller.ControllerID, controller.WritePolicy)
			ch <- prometheus.MustNewConstMetric(e.position.desc, e.position.tipo, float64(controller.PositionNumeric), enc.DurableID, controller.ControllerID, controller.Position)
			ch <- prometheus.MustNewConstMetric(e.redundancyMode.desc, e.redundancyMode.tipo, float64(controller.RedundancyModeNumeric), enc.DurableID, controller.ControllerID, controller.RedundancyMode)
			ch <- prometheus.MustNewConstMetric(e.redundancyStatus.desc, e.redundancyStatus.tipo, float64(controller.RedundancyStatusNumeric), enc.DurableID, controller.ControllerID, controller.RedundancyStatus)
			ch <- prometheus.MustNewConstMetric(e.controllerHealth.desc, e.controllerHealth.tipo, float64(controller.HealthNumeric), enc.DurableID, controller.ControllerID, controller.HealthReason, controller.HealthRecommendation)

			//Controllers NetworkParameters
			for _, networkParameters := range controller.NetworkParameters {
				ch <- prometheus.MustNewConstMetric(e.activeVersion.desc, e.activeVersion.tipo, float64(networkParameters.ActiveVersion), enc.DurableID, controller.ControllerID)
				ch <- prometheus.MustNewConstMetric(e.linkSpeed.desc, e.linkSpeed.tipo, float64(networkParameters.LinkSpeedNumeric), enc.DurableID, controller.ControllerID, networkParameters.LinkSpeed)
				ch <- prometheus.MustNewConstMetric(e.duplexMode.desc, e.duplexMode.tipo, float64(networkParameters.DuplexModeNumeric), enc.DurableID, controller.ControllerID, networkParameters.DuplexMode)
				ch <- prometheus.MustNewConstMetric(e.networkHealth.desc, e.networkHealth.tipo, float64(networkParameters.HealthNumeric), enc.DurableID, controller.ControllerID, networkParameters.Health, networkParameters.HealthRecommendation)
				ch <- prometheus.MustNewConstMetric(e.networkPingBroadcast.desc, e.networkPingBroadcast.tipo, float64(networkParameters.PingBroadcastNumeric), enc.DurableID, controller.ControllerID, networkParameters.PingBroadcast)
			}

			//Controllers Port
			for _, port := range controller.Port {
				ch <- prometheus.MustNewConstMetric(e.controller.desc, e.controller.tipo, float64(port.ControllerNumeric), enc.DurableID, controller.ControllerID, port.Controller)
				ch <- prometheus.MustNewConstMetric(e.portType.desc, e.portType.tipo, float64(port.PortTypeNumeric), enc.DurableID, controller.ControllerID, port.PortType, port.Port)
				ch <- prometheus.MustNewConstMetric(e.portStatus.desc, e.portStatus.tipo, float64(port.StatusNumeric), enc.DurableID, controller.ControllerID, port.Status, port.Port)
				ch <- prometheus.MustNewConstMetric(e.actualSpeed.desc, e.actualSpeed.tipo, float64(port.ActualSpeedNumeric), enc.DurableID, controller.ControllerID, port.ActualSpeed, port.Port)
				ch <- prometheus.MustNewConstMetric(e.configuredSpeed.desc, e.configuredSpeed.tipo, float64(port.ConfiguredSpeedNumeric), enc.DurableID, controller.ControllerID, port.ConfiguredSpeed, port.Port)
				ch <- prometheus.MustNewConstMetric(e.portHealth.desc, e.portHealth.tipo, float64(port.HealthNumeric), enc.DurableID, controller.ControllerID, port.Health, port.HealthReason, port.HealthRecommendation, port.Port)

				//Controllers Port IscsiPort
				for _, iscsiPort := range port.IscsiPort {
					ch <- prometheus.MustNewConstMetric(e.sfpStatus.desc, e.sfpStatus.tipo, float64(iscsiPort.SfpStatusNumeric), enc.DurableID, controller.ControllerID, iscsiPort.SfpStatus, port.Port)
					ch <- prometheus.MustNewConstMetric(e.sfpPresent.desc, e.sfpPresent.tipo, float64(iscsiPort.SfpPresentNumeric), enc.DurableID, controller.ControllerID, iscsiPort.SfpPresent, port.Port)
				}

			}

			//Controllers ExpanderPorts
			for _, expanderPorts := range controller.ExpanderPorts {
				ch <- prometheus.MustNewConstMetric(e.enclosureID.desc, e.enclosureID.tipo, float64(expanderPorts.EnclosureID), enc.DurableID, expanderPorts.Name, controller.ControllerID)
				ch <- prometheus.MustNewConstMetric(e.expanderPortsController.desc, e.expanderPortsController.tipo, float64(expanderPorts.ControllerNumeric), enc.DurableID, expanderPorts.Name, controller.ControllerID, expanderPorts.Controller)
				ch <- prometheus.MustNewConstMetric(e.sasPortType.desc, e.sasPortType.tipo, float64(expanderPorts.SasPortTypeNumeric), enc.DurableID, controller.ControllerID, expanderPorts.SasPortType, expanderPorts.Name)
				ch <- prometheus.MustNewConstMetric(e.sasPortIndex.desc, e.sasPortIndex.tipo, float64(expanderPorts.SasPortIndex), enc.DurableID, controller.ControllerID, expanderPorts.Name)
				ch <- prometheus.MustNewConstMetric(e.expanderPortsStatus.desc, e.expanderPortsStatus.tipo, float64(expanderPorts.StatusNumeric), enc.DurableID, controller.ControllerID, expanderPorts.Status, expanderPorts.Name)
				ch <- prometheus.MustNewConstMetric(e.expanderPortsHealth.desc, e.expanderPortsHealth.tipo, float64(expanderPorts.HealthNumeric), enc.DurableID, controller.ControllerID, expanderPorts.Health, expanderPorts.HealthReason, expanderPorts.HealthRecommendation, expanderPorts.Name)
			}

			//Controllers CompactFlash
			for _, compactFlash := range controller.CompactFlash {
				ch <- prometheus.MustNewConstMetric(e.compactFlashStatus.desc, e.compactFlashStatus.tipo, float64(compactFlash.StatusNumeric), enc.DurableID, controller.ControllerID, compactFlash.Status, compactFlash.Name)
				ch <- prometheus.MustNewConstMetric(e.cacheFlush.desc, e.cacheFlush.tipo, float64(compactFlash.CacheFlushNumeric), enc.DurableID, controller.ControllerID, compactFlash.CacheFlush, compactFlash.Name)
				ch <- prometheus.MustNewConstMetric(e.compactFlashHealth.desc, e.compactFlashHealth.tipo, float64(compactFlash.HealthNumeric), enc.DurableID, controller.ControllerID, compactFlash.Name, compactFlash.Health, compactFlash.HealthReason, compactFlash.HealthRecommendation)
			}

			//Controllers Expanders
			for _, expanders := range controller.Expanders {
				ch <- prometheus.MustNewConstMetric(e.pathID.desc, e.pathID.tipo, float64(expanders.PathIDNumeric), enc.DurableID, controller.ControllerID, expanders.PathID, expanders.Name)
				ch <- prometheus.MustNewConstMetric(e.expandersStatus.desc, e.expandersStatus.tipo, float64(expanders.StatusNumeric), enc.DurableID, controller.ControllerID, expanders.Status, expanders.Name)
				ch <- prometheus.MustNewConstMetric(e.expandersHealth.desc, e.expandersHealth.tipo, float64(expanders.HealthNumeric), enc.DurableID, controller.ControllerID, expanders.Health, expanders.HealthReason, expanders.HealthRecommendation, expanders.Name)
			}

		}

		e.collectPowerSupplies(ch, enc)
	}

	return nil
}

func (m *descMétrica) constMetric(value float64, labels ...string) prometheus.Metric {
	return prometheus.MustNewConstMetric(m.desc, m.tipo, value, labels...)
}

func (e enclosures) collectPowerSupplies(ch chan<- prometheus.Metric, enc Me.Enclosure) {
	for _, powerSupplies := range enc.PowerSupplies {
		ch <- e.powerSuppliesEnclosureID.constMetric(float64(powerSupplies.EnclosureID), enc.DurableID, powerSupplies.Name)
		ch <- e.powerSuppliesStatus.constMetric(float64(powerSupplies.StatusNumeric), enc.DurableID, powerSupplies.Name, powerSupplies.Status)
		ch <- e.powerSuppliesPosition.constMetric(float64(powerSupplies.PositionNumeric), enc.DurableID, powerSupplies.Name, powerSupplies.Position)
		ch <- e.powerSuppliesHealth.constMetric(float64(powerSupplies.HealthNumeric), enc.DurableID, powerSupplies.Name, powerSupplies.Health, powerSupplies.HealthReason, powerSupplies.HealthRecommendation)

		// Fans
		for _, powerSuppliesFans := range powerSupplies.Fan {
			ch <- e.statusSes.constMetric(float64(powerSuppliesFans.StatusSesNumeric), enc.DurableID, powerSupplies.Name, powerSuppliesFans.StatusSes)
			ch <- e.powerSuppliesFansExtendedStatus.constMetric(1, enc.DurableID, powerSupplies.Name, powerSuppliesFans.ExtendedStatus)
			ch <- e.powerSuppliesFansStatus.constMetric(float64(powerSuppliesFans.StatusNumeric), enc.DurableID, powerSupplies.Name, powerSuppliesFans.Status)
			ch <- e.speed.constMetric(float64(powerSuppliesFans.Speed), enc.DurableID, powerSupplies.Name)
			ch <- e.powerSuppliesFansPosition.constMetric(float64(powerSuppliesFans.PositionNumeric), enc.DurableID, powerSupplies.Name, powerSuppliesFans.Position)
			ch <- e.powerSuppliesFansHealth.constMetric(float64(powerSuppliesFans.HealthNumeric), enc.DurableID, powerSupplies.Name, powerSuppliesFans.Health, powerSuppliesFans.HealthReason, powerSuppliesFans.HealthRecommendation)
		}
	}
}
