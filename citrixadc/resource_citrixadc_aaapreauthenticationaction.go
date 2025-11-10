package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAaapreauthenticationaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaapreauthenticationactionFunc,
		ReadContext:   readAaapreauthenticationactionFunc,
		UpdateContext: updateAaapreauthenticationactionFunc,
		DeleteContext: deleteAaapreauthenticationactionFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"defaultepagroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deletefiles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"killprocess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preauthenticationaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaapreauthenticationactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Get("name").(string)

	aaapreauthenticationaction := aaa.Aaapreauthenticationaction{
		Defaultepagroup:         d.Get("defaultepagroup").(string),
		Deletefiles:             d.Get("deletefiles").(string),
		Killprocess:             d.Get("killprocess").(string),
		Name:                    d.Get("name").(string),
		Preauthenticationaction: d.Get("preauthenticationaction").(string),
	}

	_, err := client.AddResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName, &aaapreauthenticationaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaapreauthenticationactionName)

	return readAaapreauthenticationactionFunc(ctx, d, meta)
}

func readAaapreauthenticationactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading aaapreauthenticationaction state %s", aaapreauthenticationactionName)
	data, err := client.FindResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaapreauthenticationaction state %s", aaapreauthenticationactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("defaultepagroup", data["defaultepagroup"])
	d.Set("deletefiles", data["deletefiles"])
	d.Set("killprocess", data["killprocess"])
	d.Set("name", data["name"])
	d.Set("preauthenticationaction", data["preauthenticationaction"])

	return nil

}

func updateAaapreauthenticationactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Get("name").(string)

	aaapreauthenticationaction := aaa.Aaapreauthenticationaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultepagroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultepagroup has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Defaultepagroup = d.Get("defaultepagroup").(string)
		hasChange = true
	}
	if d.HasChange("deletefiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Deletefiles has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Deletefiles = d.Get("deletefiles").(string)
		hasChange = true
	}
	if d.HasChange("killprocess") {
		log.Printf("[DEBUG]  citrixadc-provider: Killprocess has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Killprocess = d.Get("killprocess").(string)
		hasChange = true
	}
	if d.HasChange("preauthenticationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Preauthenticationaction has changed for aaapreauthenticationaction %s, starting update", aaapreauthenticationactionName)
		aaapreauthenticationaction.Preauthenticationaction = d.Get("preauthenticationaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName, &aaapreauthenticationaction)
		if err != nil {
			return diag.Errorf("Error updating aaapreauthenticationaction %s", aaapreauthenticationactionName)
		}
	}
	return readAaapreauthenticationactionFunc(ctx, d, meta)
}

func deleteAaapreauthenticationactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaapreauthenticationactionFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationactionName := d.Id()
	err := client.DeleteResource(service.Aaapreauthenticationaction.Type(), aaapreauthenticationactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
