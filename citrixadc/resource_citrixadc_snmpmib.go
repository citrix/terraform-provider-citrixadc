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

func resourceCitrixAdcSnmpmib() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSnmpmibFunc,
		ReadContext:   readSnmpmibFunc,
		UpdateContext: updateSnmpmibFunc,
		DeleteContext: deleteSnmpmibFunc, // Thought snmpmib resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"contact": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSnmpmibFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpmibFunc")
	client := meta.(*NetScalerNitroClient).client

	// there is no primary key in snmpmib resource. Hence generate one for terraform state maintenance
	snmpmibName := resource.PrefixedUniqueId("tf-snmpmib-")

	snmpmib := snmp.Snmpmib{
		Contact:  d.Get("contact").(string),
		Customid: d.Get("customid").(string),
		Location: d.Get("location").(string),
		Name:     d.Get("name").(string),
	}

	if raw := d.GetRawConfig().GetAttr("ownernode"); !raw.IsNull() {
		snmpmib.Ownernode = intPtr(d.Get("ownernode").(int))
	}

	err := client.UpdateUnnamedResource(service.Snmpmib.Type(), &snmpmib)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(snmpmibName)

	return readSnmpmibFunc(ctx, d, meta)
}

func readSnmpmibFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpmibFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpmib state")
	data, err := client.FindResource(service.Snmpmib.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpmib state")
		d.SetId("")
		return nil
	}
	d.Set("contact", data["contact"])
	d.Set("customid", data["customid"])
	d.Set("location", data["location"])
	d.Set("name", data["name"])
	setToInt("ownernode", d, data["ownernode"])

	return nil

}

func updateSnmpmibFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpmibFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmib := snmp.Snmpmib{}

	hasChange := false
	if d.HasChange("contact") {
		log.Printf("[DEBUG]  citrixadc-provider: Contact has changed for snmpmib, starting update")
		snmpmib.Contact = d.Get("contact").(string)
		hasChange = true
	}
	if d.HasChange("customid") {
		log.Printf("[DEBUG]  citrixadc-provider: Customid has changed for snmpmib, starting update")
		snmpmib.Customid = d.Get("customid").(string)
		hasChange = true
	}
	if d.HasChange("location") {
		log.Printf("[DEBUG]  citrixadc-provider: Location has changed for snmpmib, starting update")
		snmpmib.Location = d.Get("location").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for snmpmib, starting update")
		snmpmib.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for snmpmib, starting update")
		snmpmib.Ownernode = intPtr(d.Get("ownernode").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpmib.Type(), &snmpmib)
		if err != nil {
			return diag.Errorf("Error updating snmpmib")
		}
	}
	return readSnmpmibFunc(ctx, d, meta)
}

func deleteSnmpmibFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpmibFunc")
	// snmpmib do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
