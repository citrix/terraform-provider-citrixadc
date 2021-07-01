package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	_ "fmt"
	"log"
)

func resourceCitrixAdcClusterfilesSyncer() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusterfilessyncerFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createClusterfilessyncerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusterfilesFunc")
	client := meta.(*NetScalerNitroClient).client
	timestamp := d.Get("timestamp").(string)
	clusterfiles := cluster.Clusterfiles{
		Mode: toStringList(d.Get("mode").(*schema.Set).List()),
	}

	err := client.ActOnResource(service.Clusterfiles.Type(), &clusterfiles, "sync")
	if err != nil {
		return err
	}

	d.SetId(timestamp)

	return nil
}
