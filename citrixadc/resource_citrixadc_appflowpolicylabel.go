package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppflowpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppflowpolicylabelFunc,
		ReadContext:   readAppflowpolicylabelFunc,
		DeleteContext: deleteAppflowpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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
				ForceNew: true,
			},
		},
	}
}

func createAppflowpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppflowpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicylabelName := d.Get("labelname").(string)

	appflowpolicylabel := appflow.Appflowpolicylabel{
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Appflowpolicylabel.Type(), appflowpolicylabelName, &appflowpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appflowpolicylabelName)

	return readAppflowpolicylabelFunc(ctx, d, meta)
}

func readAppflowpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppflowpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appflowpolicylabel state %s", appflowpolicylabelName)
	data, err := client.FindResource(service.Appflowpolicylabel.Type(), appflowpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appflowpolicylabel state %s", appflowpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func deleteAppflowpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppflowpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicylabelName := d.Id()
	err := client.DeleteResource(service.Appflowpolicylabel.Type(), appflowpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
