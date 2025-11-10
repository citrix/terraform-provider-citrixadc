package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCsaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCsactionFunc,
		ReadContext:   readCsactionFunc,
		UpdateContext: updateCsactionFunc,
		DeleteContext: deleteCsactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"targetlbvserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"targetvserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"targetvserverexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	csactionName := d.Get("name").(string)

	csaction := cs.Csaction{
		Comment:           d.Get("comment").(string),
		Name:              d.Get("name").(string),
		Targetlbvserver:   d.Get("targetlbvserver").(string),
		Targetvserver:     d.Get("targetvserver").(string),
		Targetvserverexpr: d.Get("targetvserverexpr").(string),
	}

	_, err := client.AddResource(service.Csaction.Type(), csactionName, &csaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(csactionName)

	return readCsactionFunc(ctx, d, meta)
}

func readCsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	csactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading csaction state %s", csactionName)
	data, err := client.FindResource(service.Csaction.Type(), csactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing csaction state %s", csactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("targetlbvserver", data["targetlbvserver"])
	d.Set("targetvserver", data["targetvserver"])
	d.Set("targetvserverexpr", data["targetvserverexpr"])

	return nil

}

func updateCsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	csactionName := d.Get("name").(string)

	csaction := cs.Csaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for csaction %s, starting update", csactionName)
		csaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for csaction %s, starting update", csactionName)
		csaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("targetlbvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Targetlbvserver has changed for csaction %s, starting update", csactionName)
		csaction.Targetlbvserver = d.Get("targetlbvserver").(string)
		hasChange = true
	}
	if d.HasChange("targetvserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Targetvserver has changed for csaction %s, starting update", csactionName)
		csaction.Targetvserver = d.Get("targetvserver").(string)
		hasChange = true
	}
	if d.HasChange("targetvserverexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Targetvserverexpr has changed for csaction %s, starting update", csactionName)
		csaction.Targetvserverexpr = d.Get("targetvserverexpr").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Csaction.Type(), csactionName, &csaction)
		if err != nil {
			return diag.Errorf("Error updating csaction %s", csactionName)
		}
	}
	return readCsactionFunc(ctx, d, meta)
}

func deleteCsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	csactionName := d.Id()
	err := client.DeleteResource(service.Csaction.Type(), csactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
