package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLbaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbactionFunc,
		ReadContext:   readLbactionFunc,
		UpdateContext: updateLbactionFunc,
		DeleteContext: deleteLbactionFunc,
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
				Computed: false,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbactionFunc")
	client := meta.(*NetScalerNitroClient).client

	lbactionName := d.Get("name").(string)

	lbaction := lb.Lbaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
		Value:   toIntegerList(d.Get("value").([]interface{})),
	}

	_, err := client.AddResource("lbaction", lbactionName, &lbaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lbactionName)

	return readLbactionFunc(ctx, d, meta)
}

func readLbactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbactionFunc")
	client := meta.(*NetScalerNitroClient).client
	lbactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbaction state %s", lbactionName)
	data, err := client.FindResource("lbaction", lbactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbaction state %s", lbactionName)
		d.SetId("")
		return nil
	}

	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("type", data["type"])
	d.Set("value", stringListToIntList(data["value"].([]interface{})))

	return nil

}

func updateLbactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbactionFunc")
	client := meta.(*NetScalerNitroClient).client
	lbactionName := d.Get("name").(string)

	lbaction := lb.Lbaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for lbaction %s, starting update", lbactionName)
		lbaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for lbaction %s, starting update", lbactionName)
		lbaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for lbaction %s, starting update", lbactionName)
		lbaction.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("value") {
		log.Printf("[DEBUG]  citrixadc-provider: Value has changed for lbaction %s, starting update", lbactionName)
		lbaction.Value = toIntegerList(d.Get("value").([]interface{}))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lbaction", lbactionName, &lbaction)
		if err != nil {
			return diag.Errorf("Error updating lbaction %s", lbactionName)
		}
	}
	return readLbactionFunc(ctx, d, meta)
}

func deleteLbactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbactionFunc")
	client := meta.(*NetScalerNitroClient).client
	lbactionName := d.Id()
	err := client.DeleteResource("lbaction", lbactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
