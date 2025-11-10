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

func resourceCitrixAdcAppalgparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppalgparamFunc,
		ReadContext:   readAppalgparamFunc,
		UpdateContext: updateAppalgparamFunc,
		DeleteContext: deleteAppalgparamFunc,
		Schema: map[string]*schema.Schema{
			"pptpgreidletimeout": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAppalgparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppalgparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var appalgparamName string
	// there is no primary key in appalgparam resource. Hence generate one for terraform state maintenance
	appalgparamName = resource.PrefixedUniqueId("tf-appalgparam-")
	appalgparam := network.Appalgparam{
		Pptpgreidletimeout: intPtr(d.Get("pptpgreidletimeout").(int)),
	}

	err := client.UpdateUnnamedResource(service.Appalgparam.Type(), &appalgparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appalgparamName)

	return readAppalgparamFunc(ctx, d, meta)
}

func readAppalgparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppalgparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading appalgparam state")
	data, err := client.FindResource(service.Appalgparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appalgparam state")
		d.SetId("")
		return nil
	}
	val, _ := strconv.Atoi(data["pptpgreidletimeout"].(string))
	d.Set("pptpgreidletimeout", val)

	return nil

}

func updateAppalgparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppalgparamFunc")
	client := meta.(*NetScalerNitroClient).client

	appalgparam := network.Appalgparam{}
	hasChange := false
	if d.HasChange("pptpgreidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Pptpgreidletimeout has changed for appalgparam, starting update")
		appalgparam.Pptpgreidletimeout = intPtr(d.Get("pptpgreidletimeout").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appalgparam.Type(), &appalgparam)
		if err != nil {
			return diag.Errorf("Error updating appalgparam")
		}
	}
	return readAppalgparamFunc(ctx, d, meta)
}

func deleteAppalgparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppalgparamFunc")

	d.SetId("")

	return nil
}
