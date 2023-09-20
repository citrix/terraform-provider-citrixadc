package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNspbrs() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNspbrsFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNspbrsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspbrsFunc")
	client := meta.(*NetScalerNitroClient).client
	nspbrsName := resource.PrefixedUniqueId("tf-nspbrs-")
	nspbrs := ns.Nspbrs{}

	var err error
	nspbrsAction := d.Get("action").(string)
	if nspbrsAction == "apply" {
		err = client.ActOnResource(service.Nspbrs.Type(), &nspbrs, "apply")
	} else if nspbrsAction == "clear" {
		err = client.ActOnResource(service.Nspbrs.Type(), &nspbrs, "clear")
	} else if nspbrsAction == "renumber" {
		err = client.ActOnResource(service.Nspbrs.Type(), &nspbrs, "renumber")
	} else {
		return fmt.Errorf("Invalid value for action %s. Supported values of action are `apply`, `clear` or `renumber`", nspbrsAction)
	}

	if err != nil {
		return err
	}

	d.SetId(nspbrsName)

	return nil
}
