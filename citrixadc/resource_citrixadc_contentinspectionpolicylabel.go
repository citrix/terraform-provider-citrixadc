package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcContentinspectionpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createContentinspectionpolicylabelFunc,
		ReadContext:   readContentinspectionpolicylabelFunc,
		DeleteContext: deleteContentinspectionpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createContentinspectionpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicylabelName := d.Get("labelname").(string)
	contentinspectionpolicylabel := contentinspection.Contentinspectionpolicylabel{
		Comment:   d.Get("comment").(string),
		Labelname: d.Get("labelname").(string),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource("contentinspectionpolicylabel", contentinspectionpolicylabelName, &contentinspectionpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(contentinspectionpolicylabelName)

	return readContentinspectionpolicylabelFunc(ctx, d, meta)
}

func readContentinspectionpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionpolicylabel state %s", contentinspectionpolicylabelName)
	data, err := client.FindResource("contentinspectionpolicylabel", contentinspectionpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionpolicylabel state %s", contentinspectionpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("type", data["type"])

	return nil

}

func deleteContentinspectionpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicylabelName := d.Id()
	err := client.DeleteResource("contentinspectionpolicylabel", contentinspectionpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
