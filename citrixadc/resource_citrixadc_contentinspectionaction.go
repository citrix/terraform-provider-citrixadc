package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcContentinspectionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createContentinspectionactionFunc,
		ReadContext:   readContentinspectionactionFunc,
		UpdateContext: updateContentinspectionactionFunc,
		DeleteContext: deleteContentinspectionactionFunc,
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
			"icapprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ifserverdown": {
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

func createContentinspectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Get("name").(string)
	contentinspectionaction := contentinspection.Contentinspectionaction{
		Icapprofilename: d.Get("icapprofilename").(string),
		Ifserverdown:    d.Get("ifserverdown").(string),
		Name:            d.Get("name").(string),
		Serverip:        d.Get("serverip").(string),
		Servername:      d.Get("servername").(string),
		Type:            d.Get("type").(string),
	}

	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		contentinspectionaction.Serverport = intPtr(d.Get("serverport").(int))
	}

	_, err := client.AddResource("contentinspectionaction", contentinspectionactionName, &contentinspectionaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(contentinspectionactionName)

	return readContentinspectionactionFunc(ctx, d, meta)
}

func readContentinspectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionaction state %s", contentinspectionactionName)
	data, err := client.FindResource("contentinspectionaction", contentinspectionactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionaction state %s", contentinspectionactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("icapprofilename", data["icapprofilename"])
	d.Set("ifserverdown", data["ifserverdown"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])
	setToInt("serverport", d, data["serverport"])
	d.Set("type", data["type"])

	return nil

}

func updateContentinspectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Get("name").(string)

	contentinspectionaction := contentinspection.Contentinspectionaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("icapprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Icapprofilename has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Icapprofilename = d.Get("icapprofilename").(string)
		hasChange = true
	}
	if d.HasChange("ifserverdown") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifserverdown has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Ifserverdown = d.Get("ifserverdown").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for contentinspectionaction %s, starting update", contentinspectionactionName)
		contentinspectionaction.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectionaction", &contentinspectionaction)
		if err != nil {
			return diag.Errorf("Error updating contentinspectionaction %s", contentinspectionactionName)
		}
	}
	return readContentinspectionactionFunc(ctx, d, meta)
}

func deleteContentinspectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionactionName := d.Id()
	err := client.DeleteResource("contentinspectionaction", contentinspectionactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
