package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	_ "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcClusterfilesSyncer() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createClusterfilessyncerFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"timestamp": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mode": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createClusterfilessyncerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusterfilesFunc")
	client := meta.(*NetScalerNitroClient).client
	timestamp := d.Get("timestamp").(string)
	clusterfiles := cluster.Clusterfiles{
		Mode: toStringList(d.Get("mode").(*schema.Set).List()),
	}

	err := client.ActOnResource(service.Clusterfiles.Type(), &clusterfiles, "sync")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(timestamp)

	return nil
}
