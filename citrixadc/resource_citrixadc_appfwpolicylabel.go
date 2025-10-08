package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwpolicylabelFunc,
		ReadContext:   readAppfwpolicylabelFunc,
		DeleteContext: deleteAppfwpolicylabelFunc,
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policylabeltype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicylabelName := d.Get("labelname").(string)
	appfwpolicylabel := appfw.Appfwpolicylabel{
		Labelname:       appfwpolicylabelName,
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Appfwpolicylabel.Type(), appfwpolicylabelName, &appfwpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appfwpolicylabelName)

	return readAppfwpolicylabelFunc(ctx, d, meta)
}

func readAppfwpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwpolicylabel state %s", appfwpolicylabelName)
	data, err := client.FindResource(service.Appfwpolicylabel.Type(), appfwpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwpolicylabel state %s", appfwpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func deleteAppfwpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicylabelName := d.Id()
	err := client.DeleteResource(service.Appfwpolicylabel.Type(), appfwpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
