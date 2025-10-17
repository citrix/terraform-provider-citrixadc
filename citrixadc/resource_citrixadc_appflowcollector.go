package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppflowcollector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppflowcollectorFunc,
		ReadContext:   readAppflowcollectorFunc,
		UpdateContext: updateAppflowcollectorFunc,
		DeleteContext: deleteAppflowcollectorFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppflowcollectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Get("name").(string)

	appflowcollector := appflow.Appflowcollector{
		Ipaddress:  d.Get("ipaddress").(string),
		Name:       d.Get("name").(string),
		Netprofile: d.Get("netprofile").(string),
		Transport:  d.Get("transport").(string),
	}

	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		appflowcollector.Port = intPtr(d.Get("port").(int))
	}

	_, err := client.AddResource(service.Appflowcollector.Type(), appflowcollectorName, &appflowcollector)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appflowcollectorName)

	return readAppflowcollectorFunc(ctx, d, meta)
}

func readAppflowcollectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appflowcollector state %s", appflowcollectorName)
	data, err := client.FindResource(service.Appflowcollector.Type(), appflowcollectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appflowcollector state %s", appflowcollectorName)
		d.SetId("")
		return nil
	}
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	setToInt("port", d, data["port"])
	d.Set("transport", data["transport"])

	return nil

}

func updateAppflowcollectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Get("name").(string)

	appflowcollector := appflow.Appflowcollector{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for appflowcollector %s, starting update", appflowcollectorName)
		appflowcollector.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for appflowcollector %s, starting update", appflowcollectorName)
		appflowcollector.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for appflowcollector %s, starting update", appflowcollectorName)
		appflowcollector.Port = intPtr(d.Get("port").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appflowcollector.Type(), appflowcollectorName, &appflowcollector)
		if err != nil {
			return diag.Errorf("Error updating appflowcollector %s", appflowcollectorName)
		}
	}
	return readAppflowcollectorFunc(ctx, d, meta)
}

func deleteAppflowcollectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Id()
	err := client.DeleteResource(service.Appflowcollector.Type(), appflowcollectorName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
