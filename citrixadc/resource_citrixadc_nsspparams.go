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

func resourceCitrixAdcNsspparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsspparamsFunc,
		ReadContext:   readNsspparamsFunc,
		UpdateContext: updateNsspparamsFunc,
		DeleteContext: deleteNsspparamsFunc,
		Schema: map[string]*schema.Schema{
			"basethreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"throttle": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsspparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsspparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsspparamsName string
	// there is no primary key in nsspparams resource. Hence generate one for terraform state maintenance
	nsspparamsName = resource.PrefixedUniqueId("tf-nsspparams-")
	nsspparams := ns.Nsspparams{
		Basethreshold: d.Get("basethreshold").(int),
		Throttle:      d.Get("throttle").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsspparams.Type(), &nsspparams)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsspparamsName)

	return readNsspparamsFunc(ctx, d, meta)
}

func readNsspparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsspparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsspparams state")
	data, err := client.FindResource(service.Nsspparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsspparams state")
		d.SetId("")
		return nil
	}
	setToInt("basethreshold", d, data["basethreshold"])
	d.Set("throttle", data["throttle"])

	return nil

}

func updateNsspparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsspparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	nsspparams := ns.Nsspparams{}
	hasChange := false
	if d.HasChange("basethreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Basethreshold has changed for nsspparams , starting update")
		nsspparams.Basethreshold = d.Get("basethreshold").(int)
		hasChange = true
	}
	if d.HasChange("throttle") {
		log.Printf("[DEBUG]  citrixadc-provider: Throttle has changed for nsspparams , starting update")
		nsspparams.Throttle = d.Get("throttle").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsspparams.Type(), &nsspparams)
		if err != nil {
			return diag.Errorf("Error updating nsspparams")
		}
	}
	return readNsspparamsFunc(ctx, d, meta)
}

func deleteNsspparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsspparamsFunc")
	// nsspparams do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
