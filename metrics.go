package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricSmartctlVersion = prometheus.NewDesc(
		"smartctl_version",
		"smartctl version",
		[]string{
			"json_format_version",
			"smartctl_version",
			"svn_revision",
			"build_info",
		},
		nil,
	)
	metricDeviceModel = prometheus.NewDesc(
		"smartctl_device",
		"Device info",
		[]string{
			"device",
			"interface",
			"protocol",
			"model_family",
			"model_name",
			"serial_number",
			"ata_additional_product_id",
			"firmware_version",
			"ata_version",
			"sata_version",
		},
		nil,
	)
	metricDeviceCapacityBlocks = prometheus.NewDesc(
		"smartctl_device_capacity_blocks",
		"Device capacity in blocks",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceCapacityBytes = prometheus.NewDesc(
		"smartctl_device_capacity_bytes",
		"Device capacity in bytes",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceBlockSize = prometheus.NewDesc(
		"smartctl_device_block_size",
		"Device block size",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"blocks_type",
		},
		nil,
	)
	metricDeviceInterfaceSpeed = prometheus.NewDesc(
		"smartctl_device_interface_speed",
		"Device interface speed, bits per second",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"speed_type",
		},
		nil,
	)
	metricDeviceAttribute = prometheus.NewDesc(
		"smartctl_device_attribute",
		"Device attributes",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"attribute_name",
			"attribute_flags_short",
			"attribute_flags_long",
			"attribute_value_type",
			"attribute_id",
		},
		nil,
	)
	metricDevicePowerOnSeconds = prometheus.NewDesc(
		"smartctl_device_power_on_seconds",
		"Device power on seconds",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceRotationRate = prometheus.NewDesc(
		"smartctl_device_rotation_rate",
		"Device rotation rate",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceTemperature = prometheus.NewDesc(
		"smartctl_device_temperature",
		"Device temperature celsius",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"temperature_type",
		},
		nil,
	)
	metricDevicePowerCycleCount = prometheus.NewDesc(
		"smartctl_device_power_cycle_count",
		"Device power cycle count",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceExitStatus = prometheus.NewDesc(
		"smartctl_device_smartctl_exit_status",
		"Exit status of smartctl on device",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceStatistics = prometheus.NewDesc(
		"smartctl_device_statistics",
		"Device statistics",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"statistic_table",
			"statistic_name",
			"statistic_flags_short",
			"statistic_flags_long",
		},
		nil,
	)
	metricCriticalWarning = prometheus.NewDesc(
		"smartctl_device_critical_warning",
		"Critical warning counter",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricAvailableSpare = prometheus.NewDesc(
		"smartctl_device_available_spare",
		"Available spare",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricMediaErrors = prometheus.NewDesc(
		"smartctl_device_media_errors",
		"Media errors counter",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricPercentageUsed = prometheus.NewDesc(
			"smartctl_device_percentage_used",
			"Percentage Used",
			[]string{
				"device",
				"model_family",
				"model_name",
				"serial_number",
			},
			nil,
		)
	metricSmartStatus = prometheus.NewDesc(
		"smartctl_device_smart_status",
		"Smart status",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
)
