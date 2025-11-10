package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsencryptionkey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsencryptionkeyFunc,
		ReadContext:   readNsencryptionkeyFunc,
		UpdateContext: updateNsencryptionkeyFunc,
		DeleteContext: deleteNsencryptionkeyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"method": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyvalue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"padding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsencryptionkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsencryptionkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nsencryptionkeyName := d.Get("name").(string)
	nsencryptionkey := ns.Nsencryptionkey{
		Comment:  d.Get("comment").(string),
		Iv:       d.Get("iv").(string),
		Keyvalue: d.Get("keyvalue").(string),
		Method:   d.Get("method").(string),
		Name:     d.Get("name").(string),
		Padding:  d.Get("padding").(string),
	}

	_, err := client.AddResource("nsencryptionkey", nsencryptionkeyName, &nsencryptionkey)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsencryptionkeyName)

	return readNsencryptionkeyFunc(ctx, d, meta)
}

func readNsencryptionkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsencryptionkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nsencryptionkeyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsencryptionkey state %s", nsencryptionkeyName)
	data, err := client.FindResource("nsencryptionkey", nsencryptionkeyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsencryptionkey state %s", nsencryptionkeyName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("iv", data["iv"])
	// d.Set("keyvalue", data["keyvalue"])
	d.Set("method", data["method"])
	d.Set("name", data["name"])
	d.Set("padding", data["padding"])

	return nil

}

func updateNsencryptionkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsencryptionkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nsencryptionkeyName := d.Get("name").(string)

	nsencryptionkey := ns.Nsencryptionkey{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for nsencryptionkey %s, starting update", nsencryptionkeyName)
		nsencryptionkey.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("iv") {
		log.Printf("[DEBUG]  citrixadc-provider: Iv has changed for nsencryptionkey %s, starting update", nsencryptionkeyName)
		nsencryptionkey.Iv = d.Get("iv").(string)
		hasChange = true
	}
	if d.HasChange("keyvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Keyvalue has changed for nsencryptionkey %s, starting update", nsencryptionkeyName)
		nsencryptionkey.Keyvalue = d.Get("keyvalue").(string)
		hasChange = true
	}
	if d.HasChange("padding") {
		log.Printf("[DEBUG]  citrixadc-provider: Padding has changed for nsencryptionkey %s, starting update", nsencryptionkeyName)
		nsencryptionkey.Padding = d.Get("padding").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("nsencryptionkey", nsencryptionkeyName, &nsencryptionkey)
		if err != nil {
			return diag.Errorf("Error updating nsencryptionkey %s", nsencryptionkeyName)
		}
	}
	return readNsencryptionkeyFunc(ctx, d, meta)
}

func deleteNsencryptionkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsencryptionkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nsencryptionkeyName := d.Id()
	err := client.DeleteResource("nsencryptionkey", nsencryptionkeyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
