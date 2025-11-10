package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcArpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createArpparamFunc,
		ReadContext:   readArpparamFunc,
		UpdateContext: updateArpparamFunc,
		DeleteContext: deleteArpparamFunc,
		Schema: map[string]*schema.Schema{
			"spoofvalidation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createArpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createArpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var arpparamName string
	// there is no primary key in arpparam resource. Hence generate one for terraform state maintenance
	arpparamName = resource.PrefixedUniqueId("tf-arpparam-")
	arpparam := network.Arpparam{
		Spoofvalidation: d.Get("spoofvalidation").(string),
	}

	if raw := d.GetRawConfig().GetAttr("timeout"); !raw.IsNull() {
		arpparam.Timeout = intPtr(d.Get("timeout").(int))
	}

	err := client.UpdateUnnamedResource(service.Arpparam.Type(), &arpparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(arpparamName)

	return readArpparamFunc(ctx, d, meta)
}

func readArpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readArpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading arpparam state")
	data, err := client.FindResource(service.Arpparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing arpparam state")
		d.SetId("")
		return nil
	}
	d.Set("spoofvalidation", data["spoofvalidation"])
	val, _ := strconv.Atoi(data["timeout"].(string))
	d.Set("timeout", val)

	return nil

}

func updateArpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateArpparamFunc")
	client := meta.(*NetScalerNitroClient).client

	arpparam := network.Arpparam{}
	hasChange := false
	if d.HasChange("spoofvalidation") {
		log.Printf("[DEBUG]  citrixadc-provider: Spoofvalidation has changed for arpparam, starting update")
		arpparam.Spoofvalidation = d.Get("spoofvalidation").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for arpparam, starting update")
		arpparam.Timeout = intPtr(d.Get("timeout").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Arpparam.Type(), &arpparam)
		if err != nil {
			return diag.Errorf("Error updating arpparam")
		}
	}
	return readArpparamFunc(ctx, d, meta)
}

func deleteArpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteArpparamFunc")

	d.SetId("")

	return nil
}
