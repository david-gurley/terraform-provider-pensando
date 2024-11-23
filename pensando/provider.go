package pensando

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Penando Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"grpc_dscs": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: descriptions["dsc_address"],
						},
						"port": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     11357,
							Description: descriptions["dsc_grpc_port"],
						},
					},
				},
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"pensando_device_status": dataSourceDeviceStatus(),
		},
		ResourcesMap:         map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"dsc_address":   "IP address or FQDN for DSC management interface.",
		"dsc_grpc_port": "TCP port for DSC gRPC management interface. Default 11357.",
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	dscAddress := d.Get("dsc_address").(string)
	dscPort := d.Get("dsc_port").(string)
	config := Config{
		DSCAddress: dscAddress,
		DSCPort:    dscPort,
	}
	client, err := config.Client()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Pensando client",
			Detail:   "Unable to create Pensando client",
		})
		return nil, diags
	}
	return client, diags
}
