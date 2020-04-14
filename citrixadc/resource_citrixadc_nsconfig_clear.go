package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcNsconfigClear() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsconfigClearFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"force": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"rbaconfig": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createNsconfigClearFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsconfigClearFunc")
	client := meta.(*NetScalerNitroClient).client
	timestamp := d.Get("timestamp").(string)
	log.Printf("[DEBUG]  citrixadc-provider: timestamp %s", timestamp)

	nsconfig := ns.Nsconfig{
		Force:     d.Get("force").(bool),
		Level:     d.Get("level").(string),
		Rbaconfig: d.Get("rbaconfig").(string),
	}

	err := client.ActOnResource(netscaler.Nsconfig.Type(), &nsconfig, "clear")
	if err != nil {
		return err
	}

	d.SetId(timestamp)

	return nil
}
