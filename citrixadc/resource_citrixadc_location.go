package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcLocation() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLocationFunc,
		ReadContext:   readLocationFunc,
		DeleteContext: deleteLocationFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipfrom": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ipto": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"preferredlocation": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"latitude": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"longitude": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLocationFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationFunc")
	client := meta.(*NetScalerNitroClient).client
	locationName := d.Get("ipfrom").(string)
	location := basic.Location{
		Ipfrom:            d.Get("ipfrom").(string),
		Ipto:              d.Get("ipto").(string),
		Preferredlocation: d.Get("preferredlocation").(string),
	}

	if raw := d.GetRawConfig().GetAttr("latitude"); !raw.IsNull() {
		location.Latitude = intPtr(d.Get("latitude").(int))
	}
	if raw := d.GetRawConfig().GetAttr("longitude"); !raw.IsNull() {
		location.Longitude = intPtr(d.Get("longitude").(int))
	}

	_, err := client.AddResource(service.Location.Type(), locationName, &location)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(locationName)

	return readLocationFunc(ctx, d, meta)
}

func readLocationFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLocationFunc")
	client := meta.(*NetScalerNitroClient).client
	locationName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading location state %s", locationName)
	data, err := client.FindResource(service.Location.Type(), locationName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing location state %s", locationName)
		d.SetId("")
		return nil
	}
	d.Set("ipfrom", data["ipfrom"])
	d.Set("ipto", data["ipto"])
	setToInt("latitude", d, data["latitude"])
	setToInt("longitude", d, data["longitude"])
	//d.Set("preferredlocation", data["preferredlocation"])

	return nil

}

func deleteLocationFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLocationFunc")
	client := meta.(*NetScalerNitroClient).client

	argsMap := make(map[string]string)
	// Only the ipfrom and ipto properties are required for deletion
	argsMap["ipfrom"] = url.QueryEscape(d.Get("ipfrom").(string))
	argsMap["ipto"] = url.QueryEscape(d.Get("ipto").(string))
	err := client.DeleteResourceWithArgsMap(service.Location.Type(), "", argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
