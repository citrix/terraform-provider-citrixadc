package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/audit"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAuditmessageaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuditmessageactionFunc,
		ReadContext:   readAuditmessageactionFunc,
		UpdateContext: updateAuditmessageactionFunc,
		DeleteContext: deleteAuditmessageactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"bypasssafetycheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loglevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtonewnslog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stringbuilderexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuditmessageactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client

	auditmessageactionName := d.Get("name").(string)
	d.Set("name", auditmessageactionName)

	auditmessageaction := audit.Auditmessageaction{
		Bypasssafetycheck: d.Get("bypasssafetycheck").(string),
		Loglevel:          d.Get("loglevel").(string),
		Logtonewnslog:     d.Get("logtonewnslog").(string),
		Name:              d.Get("name").(string),
		Stringbuilderexpr: d.Get("stringbuilderexpr").(string),
	}

	_, err := client.AddResource(service.Auditmessageaction.Type(), auditmessageactionName, &auditmessageaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(auditmessageactionName)

	return readAuditmessageactionFunc(ctx, d, meta)
}

func readAuditmessageactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditmessageactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading auditmessageaction state %s", auditmessageactionName)
	data, err := client.FindResource(service.Auditmessageaction.Type(), auditmessageactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditmessageaction state %s", auditmessageactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("bypasssafetycheck", data["bypasssafetycheck"])
	d.Set("loglevel", data["loglevel1"])
	d.Set("logtonewnslog", data["logtonewnslog"])
	d.Set("name", data["name"])
	d.Set("stringbuilderexpr", data["stringbuilderexpr"])

	return nil

}

func updateAuditmessageactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditmessageactionName := d.Get("name").(string)

	auditmessageaction := audit.Auditmessageaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bypasssafetycheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Bypasssafetycheck has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Bypasssafetycheck = d.Get("bypasssafetycheck").(string)
		hasChange = true
	}
	if d.HasChange("loglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Loglevel has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Loglevel = d.Get("loglevel").(string)
		hasChange = true
	}
	if d.HasChange("logtonewnslog") {
		log.Printf("[DEBUG]  citrixadc-provider: Logtonewnslog has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Logtonewnslog = d.Get("logtonewnslog").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("stringbuilderexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Stringbuilderexpr has changed for auditmessageaction %s, starting update", auditmessageactionName)
		auditmessageaction.Stringbuilderexpr = d.Get("stringbuilderexpr").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Auditmessageaction.Type(), auditmessageactionName, &auditmessageaction)
		if err != nil {
			return diag.Errorf("Error updating auditmessageaction %s", auditmessageactionName)
		}
	}
	return readAuditmessageactionFunc(ctx, d, meta)
}

func deleteAuditmessageactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditmessageactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditmessageactionName := d.Id()
	err := client.DeleteResource(service.Auditmessageaction.Type(), auditmessageactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
