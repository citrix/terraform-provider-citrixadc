package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSnmpengineid() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSnmpengineidFunc,
		ReadContext:   readSnmpengineidFunc,
		UpdateContext: updateSnmpengineidFunc,
		DeleteContext: deleteSnmpengineidFunc, // Thought snmpengineid resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"engineid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
func createSnmpengineidFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpengineidFunc")
	client := meta.(*NetScalerNitroClient).client

	snmpengineidName := resource.PrefixedUniqueId("tf-snmpengineid-")

	snmpengineid := snmp.Snmpengineid{
		Engineid:  d.Get("engineid").(string),
		Ownernode: d.Get("ownernode").(int),
	}

	err := client.UpdateUnnamedResource(service.Snmpengineid.Type(), &snmpengineid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(snmpengineidName)

	return readSnmpengineidFunc(ctx, d, meta)
}

func readSnmpengineidFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpengineidFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpengineid state")
	data, err := client.FindResource(service.Snmpengineid.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpengineid state")
		d.SetId("")
		return nil
	}
	d.Set("engineid", data["engineid"])
	setToInt("ownernode", d, data["ownernode"])

	return nil

}

func updateSnmpengineidFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpengineidFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpengineid := snmp.Snmpengineid{}

	hasChange := false
	if d.HasChange("engineid") {
		log.Printf("[DEBUG]  citrixadc-provider: Engineid has changed for snmpengineid, starting update")
		snmpengineid.Engineid = d.Get("engineid").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for snmpengineid, starting update")
		snmpengineid.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpengineid.Type(), &snmpengineid)
		if err != nil {
			return diag.Errorf("Error updating snmpengineid")
		}
	}
	return readSnmpengineidFunc(ctx, d, meta)
}

func deleteSnmpengineidFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpengineidFunc")
	// snmpenigneid  does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
