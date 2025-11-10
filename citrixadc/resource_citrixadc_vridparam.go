package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVridparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVridparamFunc,
		ReadContext:   readVridparamFunc,
		UpdateContext: updateVridparamFunc,
		DeleteContext: deleteVridparamFunc,
		Schema: map[string]*schema.Schema{
			"deadinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hellointerval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sendtomaster": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVridparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVridparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var vridparamName string
	// there is no primary key in vridparam resource. Hence generate one for terraform state maintenance
	vridparamName = resource.PrefixedUniqueId("tf-vridparam-")
	vridparam := network.Vridparam{
		Sendtomaster: d.Get("sendtomaster").(string),
	}

	if raw := d.GetRawConfig().GetAttr("deadinterval"); !raw.IsNull() {
		vridparam.Deadinterval = intPtr(d.Get("deadinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("hellointerval"); !raw.IsNull() {
		vridparam.Hellointerval = intPtr(d.Get("hellointerval").(int))
	}

	err := client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vridparamName)

	return readVridparamFunc(ctx, d, meta)
}

func readVridparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVridparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading vridparam state")
	data, err := client.FindResource(service.Vridparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vridparam state ")
		d.SetId("")
		return nil
	}
	setToInt("deadinterval", d, data["deadinterval"])
	setToInt("hellointerval", d, data["hellointerval"])
	d.Set("sendtomaster", data["sendtomaster"])

	return nil

}

func updateVridparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVridparamFunc")
	client := meta.(*NetScalerNitroClient).client
	vridparam := network.Vridparam{}
	hasChange := false
	if d.HasChange("deadinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Deadinterval has changed for vridparam, starting update")
		vridparam.Deadinterval = intPtr(d.Get("deadinterval").(int))
		hasChange = true
	}
	if d.HasChange("hellointerval") {
		log.Printf("[DEBUG]  citrixadc-provider: Hellointerval has changed for vridparam, starting update")
		vridparam.Hellointerval = intPtr(d.Get("hellointerval").(int))
		hasChange = true
	}
	if d.HasChange("sendtomaster") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendtomaster has changed for vridparam, starting update")
		vridparam.Sendtomaster = d.Get("sendtomaster").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
		if err != nil {
			return diag.Errorf("Error updating vridparam")
		}
	}
	return readVridparamFunc(ctx, d, meta)
}

func deleteVridparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVridparamFunc")

	d.SetId("")

	return nil
}
