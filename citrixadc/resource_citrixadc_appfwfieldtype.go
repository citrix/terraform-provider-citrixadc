package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwfieldtype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwfieldtypeFunc,
		ReadContext:   readAppfwfieldtypeFunc,
		UpdateContext: updateAppfwfieldtypeFunc,
		DeleteContext: deleteAppfwfieldtypeFunc,
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
			"nocharmaps": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"regex": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAppfwfieldtypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Get("name").(string)
	appfwfieldtype := appfw.Appfwfieldtype{
		Comment:    d.Get("comment").(string),
		Name:       appfwfieldtypeName,
		Nocharmaps: d.Get("nocharmaps").(bool),
		Regex:      d.Get("regex").(string),
	}

	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		appfwfieldtype.Priority = intPtr(d.Get("priority").(int))
	}

	_, err := client.AddResource(service.Appfwfieldtype.Type(), appfwfieldtypeName, &appfwfieldtype)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appfwfieldtypeName)

	return readAppfwfieldtypeFunc(ctx, d, meta)
}

func readAppfwfieldtypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwfieldtype state %s", appfwfieldtypeName)
	data, err := client.FindResource(service.Appfwfieldtype.Type(), appfwfieldtypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwfieldtype state %s", appfwfieldtypeName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("nocharmaps", data["nocharmaps"])
	setToInt("priority", d, data["priority"])
	d.Set("regex", data["regex"])

	return nil

}

func updateAppfwfieldtypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Get("name").(string)

	appfwfieldtype := appfw.Appfwfieldtype{
		Name:  d.Get("name").(string),
		Regex: d.Get("regex").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appfwfieldtype %s, starting update", appfwfieldtypeName)
		appfwfieldtype.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("nocharmaps") {
		log.Printf("[DEBUG]  citrixadc-provider: Nocharmaps has changed for appfwfieldtype %s, starting update", appfwfieldtypeName)
		appfwfieldtype.Nocharmaps = d.Get("nocharmaps").(bool)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for appfwfieldtype %s, starting update", appfwfieldtypeName)
		appfwfieldtype.Priority = intPtr(d.Get("priority").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appfwfieldtype.Type(), appfwfieldtypeName, &appfwfieldtype)
		if err != nil {
			return diag.Errorf("Error updating appfwfieldtype %s", appfwfieldtypeName)
		}
	}
	return readAppfwfieldtypeFunc(ctx, d, meta)
}

func deleteAppfwfieldtypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Id()
	err := client.DeleteResource(service.Appfwfieldtype.Type(), appfwfieldtypeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
