package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcExtendedmemoryparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createExtendedmemoryparamFunc,
		ReadContext:   readExtendedmemoryparamFunc,
		UpdateContext: updateExtendedmemoryparamFunc,
		DeleteContext: deleteExtendedmemoryparamFunc,
		Schema: map[string]*schema.Schema{
			"memlimit": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createExtendedmemoryparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createExtendedmemoryparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var extendedmemoryparamName string
	extendedmemoryparamName = resource.PrefixedUniqueId("tf-extendedmemoryparam-")
	extendedmemoryparam := basic.Extendedmemoryparam{}

	if raw := d.GetRawConfig().GetAttr("memlimit"); !raw.IsNull() {
		extendedmemoryparam.Memlimit = intPtr(d.Get("memlimit").(int))
	}

	err := client.UpdateUnnamedResource(service.Extendedmemoryparam.Type(), &extendedmemoryparam)
	if err != nil {
		return diag.Errorf("Error updating extendedmemoryparam")
	}

	d.SetId(extendedmemoryparamName)

	return readExtendedmemoryparamFunc(ctx, d, meta)
}

func readExtendedmemoryparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readExtendedmemoryparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading extendedmemoryparam state")
	data, err := client.FindResource(service.Extendedmemoryparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing extendedmemoryparam state")
		d.SetId("")
		return nil
	}
	setToInt("memlimit", d, data["memlimit"])

	return nil

}

func updateExtendedmemoryparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateExtendedmemoryparamFunc")
	client := meta.(*NetScalerNitroClient).client

	extendedmemoryparam := basic.Extendedmemoryparam{}
	hasChange := false
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for extendedmemoryparam , starting update")
		extendedmemoryparam.Memlimit = intPtr(d.Get("memlimit").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Extendedmemoryparam.Type(), &extendedmemoryparam)
		if err != nil {
			return diag.Errorf("Error updating extendedmemoryparam %s", err.Error())
		}
	}
	return readExtendedmemoryparamFunc(ctx, d, meta)
}

func deleteExtendedmemoryparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteExtendedmemoryparamFunc")

	d.SetId("")

	return nil
}
