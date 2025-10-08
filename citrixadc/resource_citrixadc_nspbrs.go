package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNspbrs() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNspbrsFunc,
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

func createNspbrsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.Errorf("Invalid value for action %s. Supported values of action are `apply`, `clear` or `renumber`", nspbrsAction)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nspbrsName)

	return nil
}
