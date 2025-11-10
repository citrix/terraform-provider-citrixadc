package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsnappsattributes() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnappsattributesFunc,
		ReadContext:   readLsnappsattributesFunc,
		UpdateContext: updateLsnappsattributesFunc,
		DeleteContext: deleteLsnappsattributesFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"sessiontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transportprotocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnappsattributesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Get("name").(string)
	lsnappsattributes := lsn.Lsnappsattributes{
		Name:              d.Get("name").(string),
		Port:              d.Get("port").(string),
		Transportprotocol: d.Get("transportprotocol").(string),
	}

	if raw := d.GetRawConfig().GetAttr("sessiontimeout"); !raw.IsNull() {
		lsnappsattributes.Sessiontimeout = intPtr(d.Get("sessiontimeout").(int))
	}

	_, err := client.AddResource("lsnappsattributes", lsnappsattributesName, &lsnappsattributes)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnappsattributesName)

	return readLsnappsattributesFunc(ctx, d, meta)
}

func readLsnappsattributesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnappsattributes state %s", lsnappsattributesName)
	data, err := client.FindResource("lsnappsattributes", lsnappsattributesName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnappsattributes state %s", lsnappsattributesName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("port", data["port"])
	setToInt("sessiontimeout", d, data["sessiontimeout"])
	d.Set("transportprotocol", data["transportprotocol"])

	return nil

}

func updateLsnappsattributesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Get("name").(string)

	lsnappsattributes := lsn.Lsnappsattributes{
		Name: d.Get("name").(string),
	}
	hasChange := false

	if d.HasChange("sessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessiontimeout has changed for lsnappsattributes %s, starting update", lsnappsattributesName)
		lsnappsattributes.Sessiontimeout = intPtr(d.Get("sessiontimeout").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnappsattributes", &lsnappsattributes)
		if err != nil {
			return diag.Errorf("Error updating lsnappsattributes %s", lsnappsattributesName)
		}
	}
	return readLsnappsattributesFunc(ctx, d, meta)
}

func deleteLsnappsattributesFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Id()
	err := client.DeleteResource("lsnappsattributes", lsnappsattributesName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
