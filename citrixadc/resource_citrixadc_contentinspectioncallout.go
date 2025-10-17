package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcContentinspectioncallout() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createContentinspectioncalloutFunc,
		ReadContext:   readContentinspectioncalloutFunc,
		UpdateContext: updateContentinspectioncalloutFunc,
		DeleteContext: deleteContentinspectioncalloutFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resultexpr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"returntype": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createContentinspectioncalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Get("name").(string)
	contentinspectioncallout := contentinspection.Contentinspectioncallout{
		Comment:     d.Get("comment").(string),
		Name:        d.Get("name").(string),
		Profilename: d.Get("profilename").(string),
		Resultexpr:  d.Get("resultexpr").(string),
		Returntype:  d.Get("returntype").(string),
		Serverip:    d.Get("serverip").(string),
		Servername:  d.Get("servername").(string),
		Type:        d.Get("type").(string),
	}

	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		contentinspectioncallout.Serverport = intPtr(d.Get("serverport").(int))
	}

	_, err := client.AddResource("contentinspectioncallout", contentinspectioncalloutName, &contentinspectioncallout)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(contentinspectioncalloutName)

	return readContentinspectioncalloutFunc(ctx, d, meta)
}

func readContentinspectioncalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectioncallout state %s", contentinspectioncalloutName)
	data, err := client.FindResource("contentinspectioncallout", contentinspectioncalloutName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectioncallout state %s", contentinspectioncalloutName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("profilename", data["profilename"])
	d.Set("resultexpr", data["resultexpr"])
	d.Set("returntype", data["returntype"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])
	setToInt("serverport", d, data["serverport"])
	d.Set("type", data["type"])

	return nil

}

func updateContentinspectioncalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Get("name").(string)

	contentinspectioncallout := contentinspection.Contentinspectioncallout{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("resultexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Resultexpr has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Resultexpr = d.Get("resultexpr").(string)
		hasChange = true
	}
	if d.HasChange("returntype") {
		log.Printf("[DEBUG]  citrixadc-provider: Returntype has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Returntype = d.Get("returntype").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for contentinspectioncallout %s, starting update", contentinspectioncalloutName)
		contentinspectioncallout.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectioncallout", &contentinspectioncallout)
		if err != nil {
			return diag.Errorf("Error updating contentinspectioncallout %s", contentinspectioncalloutName)
		}
	}
	return readContentinspectioncalloutFunc(ctx, d, meta)
}

func deleteContentinspectioncalloutFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectioncalloutFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectioncalloutName := d.Id()
	err := client.DeleteResource("contentinspectioncallout", contentinspectioncalloutName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
