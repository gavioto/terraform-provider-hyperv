package hyperv

import (
	"fmt"
	"log"
	"github.com/taliesins/terraform-provider-hyperv/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHyperVMachineInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceHyperVMachineInstanceCreate,
		Read:   resourceHyperVMachineInstanceRead,
		Update: resourceHyperVMachineInstanceUpdate,
		Delete: resourceHyperVMachineInstanceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"generation": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ForceNew: true,
			},

			"automatic_critical_error_action": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      api.CriticalErrorAction_name[api.CriticalErrorAction_Pause],
				ValidateFunc: stringKeyInMap(api.CriticalErrorAction_value, true),
			},

			"automatic_critical_error_action_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},

			"automatic_start_action": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      api.StartAction_name[api.StartAction_StartIfRunning],
				ValidateFunc: stringKeyInMap(api.StartAction_value, true),
			},

			"automatic_start_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			"automatic_stop_action": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      api.StopAction_name[api.StopAction_Save],
				ValidateFunc: stringKeyInMap(api.StopAction_value, true),
			},

			"checkpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      api.CheckpointType_name[api.CheckpointType_Production],
				ValidateFunc: stringKeyInMap(api.CheckpointType_value, true),
			},

			"dynamic_memory": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"guest_controlled_cache_types": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"high_memory_mapped_io_space": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  536870912,
			},

			"lock_on_disconnect": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      api.OnOffState_name[api.OnOffState_Off],
				ValidateFunc: stringKeyInMap(api.OnOffState_value, true),
			},

			"low_memory_mapped_io_space": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  134217728,
			},

			"memory_maximum_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1099511627776,
			},

			"memory_minimum_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  536870912,
			},

			"memory_startup_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  536870912,
			},

			"notes": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"processor_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},

			"smart_paging_file_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  `C:\ProgramData\Microsoft\Windows\Hyper-V`,
			},

			"snapshot_file_location": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  `C:\ProgramData\Microsoft\Windows\Hyper-V`,
			},

			"static_memory": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"integration_services": {
				Type:     schema.TypeMap,
				Optional: true,
				DefaultFunc: api.DefaultVmIntegrationServices,
				DiffSuppressFunc: api.DiffSuppressVmIntegrationServices,
				Elem:     schema.TypeBool,
			},

			"network_adaptors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"switch_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
							ForceNew: true,
						},
						"management_os": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"is_legacy": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: true,
						},
						"dynamic_mac_address": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"static_mac_address": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"dhcp_guard": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.OnOffState_name[api.OnOffState_Off],
							ValidateFunc: stringKeyInMap(api.OnOffState_value, true),
						},
						"router_guard": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.OnOffState_name[api.OnOffState_Off],
							ValidateFunc: stringKeyInMap(api.OnOffState_value, true),
						},
						"port_mirroring": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.PortMirroring_name[api.PortMirroring_None],
							ValidateFunc: stringKeyInMap(api.PortMirroring_value, true),
						},
						"ieee_priority_tag": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.OnOffState_name[api.OnOffState_Off],
							ValidateFunc: stringKeyInMap(api.OnOffState_value, true),
						},
						"vmq_weigth": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"iov_queue_pairs_requested": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"iov_interrupt_moderation": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.IovInterruptModerationValue_name[api.IovInterruptModerationValue_Off],
							ValidateFunc: stringKeyInMap(api.IovInterruptModerationValue_value, true),
						},
						"iov_weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"ipsec_offload_maximum_security_association": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"maximum_bandwidth": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"minimum_bandwidth_absolute": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"minimum_bandwidth_weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"mandatory_feature_id": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"resource_pool_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"test_replica_pool_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"test_replica_switch_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"virtual_subnet_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"allow_teaming": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.OnOffState_name[api.OnOffState_Off],
							ValidateFunc: stringKeyInMap(api.OnOffState_value, true),
						},
						"not_monitored_in_cluster": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"storm_limit": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"dynamic_ip_address_limit": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"device_naming": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.OnOffState_name[api.OnOffState_Off],
							ValidateFunc: stringKeyInMap(api.OnOffState_value, true),
						},
						"fix_speed_10g": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.OnOffState_name[api.OnOffState_Off],
							ValidateFunc: stringKeyInMap(api.OnOffState_value, true),
						},
						"packet_direct_num_procs": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"packet_direct_moderation_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"packet_direct_moderation_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"vrss_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"vmmq_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"vmmq_queue_pairs": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},

			"dvd_drives": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller_number": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"controller_location": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"resource_pool_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
					},
				},
			},

			"hard_disk_drives": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller_type": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.ControllerType_name[api.ControllerType_Ide],
							ValidateFunc: stringKeyInMap(api.ControllerType_value, true),
						},
						"controller_number": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"controller_location": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"disk_number": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"resource_pool_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"support_persistent_reservations": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"maximum_iops": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"minimum_iops": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"qos_policy_id": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"override_cache_attributes": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      api.CacheAttributes_name[api.CacheAttributes_Default],
							ValidateFunc: stringKeyInMap(api.CacheAttributes_value, true),
						},
					},
				},
			},
		},
	}
}

func resourceHyperVMachineInstanceCreate(data *schema.ResourceData, meta interface{}) (err error) {
	log.Printf("[INFO][hyperv][create] creating hyperv machine: %#v", data)
	client := meta.(*api.HypervClient)

	name := ""

	if v, ok := data.GetOk("name"); ok {
		name = v.(string)
	} else {
		return fmt.Errorf("[ERROR][hyperv][create] name argument is required")
	}

	generation := (data.Get("generation")).(int)
	automaticCriticalErrorAction := api.ToCriticalErrorAction((data.Get("automatic_critical_error_action")).(string))
	automaticCriticalErrorActionTimeout := int32((data.Get("automatic_critical_error_action_timeout")).(int))
	automaticStartAction := api.ToStartAction((data.Get("automatic_start_action")).(string))
	automaticStartDelay := int32((data.Get("automatic_start_delay")).(int))
	automaticStopAction := api.ToStopAction((data.Get("automatic_stop_action")).(string))
	checkpointType := api.ToCheckpointType((data.Get("checkpoint_type")).(string))
	dynamicMemory := (data.Get("dynamic_memory")).(bool)
	guestControlledCacheTypes := (data.Get("guest_controlled_cache_types")).(bool)
	highMemoryMappedIoSpace := int64((data.Get("high_memory_mapped_io_space")).(int))
	lockOnDisconnect := api.ToOnOffState((data.Get("lock_on_disconnect")).(string))
	lowMemoryMappedIoSpace := int32((data.Get("low_memory_mapped_io_space")).(int))
	memoryMaximumBytes := int64((data.Get("memory_maximum_bytes")).(int))
	memoryMinimumBytes := int64((data.Get("memory_minimum_bytes")).(int))
	memoryStartupBytes := int64((data.Get("memory_startup_bytes")).(int))
	notes := (data.Get("notes")).(string)
	processorCount := int64((data.Get("processor_count")).(int))
	smartPagingFilePath := (data.Get("smart_paging_file_path")).(string)
	snapshotFileLocation := (data.Get("snapshot_file_location")).(string)
	staticMemory := (data.Get("static_memory")).(bool)

	if dynamicMemory && staticMemory {
		return fmt.Errorf("[ERROR][hyperv][create] Dynamic and static can't be both selected at the same time")
	}

	if !dynamicMemory && !staticMemory {
		return fmt.Errorf("[ERROR][hyperv][create] Either dynamic or static must be selected")
	}

	integrationServices, err := api.ExpandIntegrationServices(data)
	if err != nil {
		return err
	}

	networkAdapters, err := api.ExpandNetworkAdapters(data)
	if err != nil {
		return err
	}

	dvdDrives, err := api.ExpandDvdDrives(data)
	if err != nil {
		return err
	}

	hardDiskDrives, err := api.ExpandHardDiskDrives(data)
	if err != nil {
		return err
	}

	err = client.CreateVM(name, generation, automaticCriticalErrorAction, automaticCriticalErrorActionTimeout, automaticStartAction, automaticStartDelay, automaticStopAction, checkpointType, dynamicMemory, guestControlledCacheTypes, highMemoryMappedIoSpace, lockOnDisconnect, lowMemoryMappedIoSpace, memoryMaximumBytes, memoryMinimumBytes, memoryStartupBytes, notes, processorCount, smartPagingFilePath, snapshotFileLocation, staticMemory)
	if err != nil {
		return err
	}

	err = client.CreateOrUpdateVMNetworkAdapters(name, networkAdapters)
	if err != nil {
		return err
	}

	err = client.CreateOrUpdateVMIntegrationServices(name, integrationServices)
	if err != nil {
		return err
	}

	err = client.CreateOrUpdateVMDvdDrives(name, dvdDrives)
	if err != nil {
		return err
	}

	err = client.CreateOrUpdateVMHardDiskDrives(name, hardDiskDrives)
	if err != nil {
		return err
	}

	data.SetId(name)
	log.Printf("[INFO][hyperv][create] created hyperv machine: %#v", data)

	return resourceHyperVMachineInstanceRead(data, meta)
}

func resourceHyperVMachineInstanceRead(data *schema.ResourceData, meta interface{}) (err error) {
	log.Printf("[INFO][hyperv][read] reading hyperv machine: %#v", data)
	client := meta.(*api.HypervClient)

	name := data.Id()

	vm, err := client.GetVM(name)
	if err != nil {
		return err
	}

	integrationServices, err := client.GetVMIntegrationServices(name)
	if err != nil {
		return err
	}

	networkAdapters, err := client.GetVMNetworkAdapters(name)
	if err != nil {
		return err
	}

	dvdDrives, err := client.GetVMDvdDrives(name)
	if err != nil {
		return err
	}

	hardDiskDrives, err := client.GetVMHardDiskDrives(name)
	if err != nil {
		return err
	}

	log.Printf("[INFO][hyperv][read] retrieved vm: %+v", vm)

	if vm.Name != name {
		data.SetId("")
		log.Printf("[INFO][hyperv][read] unable to read hyperv machine as it does not exist: %#v", name)
		return nil
	}

	if vm.DynamicMemory && vm.StaticMemory {
		return fmt.Errorf("[ERROR][hyperv][read] Dynamic and static can't be both selected at the same time")
	}

	if !vm.DynamicMemory && !vm.StaticMemory {
		return fmt.Errorf("[ERROR][hyperv][read] Either dynamic or static must be selected")
	}

	flattenedNetworkAdapters := api.FlattenNetworkAdapters(&networkAdapters)
	if err := data.Set("network_adaptors", flattenedNetworkAdapters); err != nil {
		return fmt.Errorf("[DEBUG] Error setting network_adaptors error: %v", err)
	}

	flattenedIntegrationServices := api.FlattenIntegrationServices(&integrationServices)
	if err := data.Set("integration_services", flattenedIntegrationServices); err != nil {
		return fmt.Errorf("[DEBUG] Error setting integration_services error: %v", err)
	}
	log.Printf("[INFO][hyperv][read] flattenedIntegrationServices: %v", flattenedIntegrationServices)

	flattenedDvdDrives := api.FlattenDvdDrives(&dvdDrives)
	if err := data.Set("dvd_drives", flattenedDvdDrives); err != nil {
		return fmt.Errorf("[DEBUG] Error setting dvd_drives error: %v", err)
	}

	flattenedHardDiskDrives := api.FlattenHardDiskDrives(&hardDiskDrives)
	if err := data.Set("hard_disk_drives", flattenedHardDiskDrives); err != nil {
		return fmt.Errorf("[DEBUG] Error setting hard_disk_drives error: %v", err)
	}

	data.Set("generation", vm.Generation)
	data.Set("automatic_critical_error_action", vm.AutomaticCriticalErrorAction.String())
	data.Set("automatic_critical_error_action_timeout", vm.AutomaticCriticalErrorActionTimeout)
	data.Set("automatic_start_action", vm.AutomaticStartAction.String())
	data.Set("automatic_start_delay", vm.AutomaticStartDelay)
	data.Set("automatic_stop_action", vm.AutomaticStopAction.String())
	data.Set("checkpoint_type", vm.CheckpointType.String())
	data.Set("dynamic_memory", vm.DynamicMemory)
	data.Set("guest_controlled_cache_types", vm.GuestControlledCacheTypes)
	data.Set("high_memory_mapped_io_space", vm.HighMemoryMappedIoSpace)
	data.Set("lock_on_disconnect", vm.LockOnDisconnect.String())
	data.Set("low_memory_mapped_io_space", vm.LowMemoryMappedIoSpace)
	data.Set("memory_maximum_bytes", vm.MemoryMaximumBytes)
	data.Set("memory_minimum_bytes", vm.MemoryMinimumBytes)
	data.Set("memory_startup_bytes", vm.MemoryStartupBytes)
	data.Set("notes", vm.Notes)
	data.Set("processor_count", vm.ProcessorCount)
	data.Set("smart_paging_file_path", vm.SmartPagingFilePath)
	data.Set("snapshot_file_location", vm.SnapshotFileLocation)
	data.Set("static_memory", vm.StaticMemory)

	log.Printf("[INFO][hyperv][read] read hyperv machine: %#v", data)

	return nil
}

func resourceHyperVMachineInstanceUpdate(data *schema.ResourceData, meta interface{}) (err error) {
	log.Printf("[INFO][hyperv][update] updating hyperv machine: %#v", data)
	client := meta.(*api.HypervClient)

	name := data.Id()

	if data.HasChange("automatic_critical_error_action") ||
		data.HasChange("automatic_critical_error_action_timeout") ||
		data.HasChange("automatic_start_action") ||
		data.HasChange("automatic_start_delay") ||
		data.HasChange("automatic_stop_action") ||
		data.HasChange("checkpoint_type") ||
		data.HasChange("dynamic_memory") ||
		data.HasChange("guest_controlled_cache_types") ||
		data.HasChange("high_memory_mapped_io_space") ||
		data.HasChange("lock_on_disconnect") ||
		data.HasChange("low_memory_mapped_io_space") ||
		data.HasChange("memory_maximum_bytes") ||
		data.HasChange("memory_minimum_bytes") ||
		data.HasChange("memory_startup_bytes") ||
		data.HasChange("notes") ||
		data.HasChange("processor_count") ||
		data.HasChange("smart_paging_file_path") ||
		data.HasChange("snapshot_file_location") ||
		data.HasChange("static_memory") {
		//generation := (d.Get("generation")).(int)
		automaticCriticalErrorAction := api.ToCriticalErrorAction((data.Get("automatic_critical_error_action")).(string))
		automaticCriticalErrorActionTimeout := int32((data.Get("automatic_critical_error_action_timeout")).(int))
		automaticStartAction := api.ToStartAction((data.Get("automatic_start_action")).(string))
		automaticStartDelay := int32((data.Get("automatic_start_delay")).(int))
		automaticStopAction := api.ToStopAction((data.Get("automatic_stop_action")).(string))
		checkpointType := api.ToCheckpointType((data.Get("checkpoint_type")).(string))
		dynamicMemory := (data.Get("dynamic_memory")).(bool)
		guestControlledCacheTypes := (data.Get("guest_controlled_cache_types")).(bool)
		highMemoryMappedIoSpace := int64((data.Get("high_memory_mapped_io_space")).(int))
		lockOnDisconnect := api.ToOnOffState((data.Get("lock_on_disconnect")).(string))
		lowMemoryMappedIoSpace := int32((data.Get("low_memory_mapped_io_space")).(int))
		memoryMaximumBytes := int64((data.Get("memory_maximum_bytes")).(int))
		memoryMinimumBytes := int64((data.Get("memory_minimum_bytes")).(int))
		memoryStartupBytes := int64((data.Get("memory_startup_bytes")).(int))
		notes := (data.Get("notes")).(string)
		processorCount := int64((data.Get("processor_count")).(int))
		smartPagingFilePath := (data.Get("smart_paging_file_path")).(string)
		snapshotFileLocation := (data.Get("snapshot_file_location")).(string)
		staticMemory := (data.Get("static_memory")).(bool)

		if dynamicMemory && staticMemory {
			return fmt.Errorf("[ERROR][hyperv][update] Dynamic and static can't be both selected at the same time")
		}

		if !dynamicMemory && !staticMemory {
			return fmt.Errorf("[ERROR][hyperv][update] Either dynamic or static must be selected")
		}

		err = client.UpdateVM(name, automaticCriticalErrorAction, automaticCriticalErrorActionTimeout, automaticStartAction, automaticStartDelay, automaticStopAction, checkpointType, dynamicMemory, guestControlledCacheTypes, highMemoryMappedIoSpace, lockOnDisconnect, lowMemoryMappedIoSpace, memoryMaximumBytes, memoryMinimumBytes, memoryStartupBytes, notes, processorCount, smartPagingFilePath, snapshotFileLocation, staticMemory)
		if err != nil {
			return err
		}
	}

	if data.HasChange("integration_services") {
		integrationServices, err := api.ExpandIntegrationServices(data)
		if err != nil {
			return err
		}

		changedIntegrationServices := api.GetChangedIntegrationServices(integrationServices, data)

		err = client.CreateOrUpdateVMIntegrationServices(name, changedIntegrationServices)
		if err != nil {
			return err
		}
	}

	if data.HasChange("network_adaptors") {
		networkAdapters, err := api.ExpandNetworkAdapters(data)
		if err != nil {
			return err
		}

		err = client.CreateOrUpdateVMNetworkAdapters(name, networkAdapters)
		if err != nil {
			return err
		}
	}

	if data.HasChange("dvd_drives") {
		dvdDrives, err := api.ExpandDvdDrives(data)
		if err != nil {
			return err
		}

		err = client.CreateOrUpdateVMDvdDrives(name, dvdDrives)
		if err != nil {
			return err
		}
	}

	if data.HasChange("hard_disk_drives") {
		hardDiskDrives, err := api.ExpandHardDiskDrives(data)
		if err != nil {
			return err
		}

		err = client.CreateOrUpdateVMHardDiskDrives(name, hardDiskDrives)
		if err != nil {
			return err
		}
	}

	log.Printf("[INFO][hyperv][update] updated hyperv machine: %#v", data)

	return resourceHyperVMachineInstanceRead(data, meta)
}

func resourceHyperVMachineInstanceDelete(data *schema.ResourceData, meta interface{}) (err error) {
	log.Printf("[INFO][hyperv][delete] deleting hyperv machine: %#v", data)

	client := meta.(*api.HypervClient)

	name := data.Id()

	err = client.DeleteVM(name)

	if err != nil {
		return err
	}

	log.Printf("[INFO][hyperv][delete] deleted hyperv machine: %#v", data)
	return nil
}