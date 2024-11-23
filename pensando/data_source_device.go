package pensando

import (
	"context"
	"strconv"

	"github.com/david-gurley/gopen/pds"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceStatusRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_mac_address": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sku": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"firmware_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"manufacturing_date": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"chip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardware_revision": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_specification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"soc_os_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"soc_disk_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pcie_specification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pcie_bus_info": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"num_pcie_ports": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"num_ports": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vendor_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pxe_version": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"uefi_version": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"num_host_if": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"firmware_description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"firmware_build_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDeviceStatusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*PensandoClient).DSCClient

	deviceStatus, err := client.GetDeviceStatus()
	if err != nil {
		return diag.FromErr(err)
	}

	return deviceStatusAttributes(d, deviceStatus, meta)
}

func deviceStatusAttributes(d *schema.ResourceData, deviceStatus *pds.DeviceStatus, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId(strconv.FormatUint(deviceStatus.SystemMACAddress, 16))
	d.Set("system_mac_address", deviceStatus.SystemMACAddress)
	d.Set("serial_number", deviceStatus.SerialNumber)
	d.Set("sku", deviceStatus.Sku)
	d.Set("firmware_version", deviceStatus.FirmwareVersion)
	d.Set("memory", deviceStatus.Memory)
	d.Set("product_name", deviceStatus.ProductName)
	d.Set("manufacturing_date", deviceStatus.ManufacturingDate)
	d.Set("desciption", deviceStatus.Description)
	d.Set("vendor_id", deviceStatus.VendorID)
	d.Set("chip_type", deviceStatus.ChipType)
	d.Set("hardware_revision", deviceStatus.HardwareRevision)
	d.Set("cpu_vendor", deviceStatus.CpuVendor)
	d.Set("cpu_specification", deviceStatus.CpuSpecification)
	d.Set("soc_os_version", deviceStatus.SocOSVersion)
	d.Set("soc_disk_size", deviceStatus.SocDiskSize)
	d.Set("pcie_specification", deviceStatus.PCIeSpecification)
	d.Set("pcie_bus_info", deviceStatus.PCIeBusInfo)
	d.Set("num_pcie_ports", deviceStatus.NumPCIePorts)
	d.Set("num_ports", deviceStatus.NumPorts)
	d.Set("vendor_name", deviceStatus.VendorName)
	d.Set("pxe_version", deviceStatus.PXEVersion)
	d.Set("uefi_version", deviceStatus.UEFIVersion)
	d.Set("num_host_if", deviceStatus.NumHostIf)
	d.Set("firmware_description", deviceStatus.FirmwareDescription)
	d.Set("firmware_build_time", deviceStatus.FirmwareBuildTime)
	return diags
}
