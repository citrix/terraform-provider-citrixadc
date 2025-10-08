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

func resourceCitrixAdcNsencryptionparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsencryptionparamsFunc,
		ReadContext:   readNsencryptionparamsFunc,
		UpdateContext: updateNsencryptionparamsFunc,
		DeleteContext: deleteNsencryptionparamsFunc,
		Schema: map[string]*schema.Schema{
			"method": {
				Type:     schema.TypeString,
				Required: true,
			},
			"keyvalue": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createNsencryptionparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsencryptionparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	nsencryptionparamsName := resource.PrefixedUniqueId("tf-nsencryptionparams-")
	nsencryptionparams := ns.Nsencryptionparams{
		Keyvalue: d.Get("keyvalue").(string),
		Method:   d.Get("method").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsencryptionparams.Type(), &nsencryptionparams)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsencryptionparamsName)

	return readNsencryptionparamsFunc(ctx, d, meta)
}

func readNsencryptionparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsencryptionparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsencryptionparams state")
	data, err := client.FindResource(service.Nsencryptionparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsencryptionparams state")
		d.SetId("")
		return nil
	}
	d.Set("keyvalue", data["keyvalue"])
	d.Set("method", data["method"])

	return nil

}

func updateNsencryptionparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsencryptionparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	nsencryptionparams := ns.Nsencryptionparams{}
	hasChange := false

	if d.HasChange("keyvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Keyvalue has changed for nsencryptionparams, starting update")
		nsencryptionparams.Keyvalue = d.Get("keyvalue").(string)
		hasChange = true
	}
	if d.HasChange("method") {
		log.Printf("[DEBUG]  citrixadc-provider: Method has changed for nsencryptionparams, starting update")
		nsencryptionparams.Method = d.Get("method").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsencryptionparams.Type(), &nsencryptionparams)
		if err != nil {
			return diag.Errorf("Error updating nsencryptionparams")
		}
	}
	return readNsencryptionparamsFunc(ctx, d, meta)
}

func deleteNsencryptionparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsencryptionparamsFunc")
	// nsencryption does not support delete operation
	d.SetId("")

	return nil
}
