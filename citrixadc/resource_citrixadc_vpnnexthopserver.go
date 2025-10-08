package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpnnexthopserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnnexthopserverFunc,
		ReadContext:   readVpnnexthopserverFunc,
		DeleteContext: deleteVpnnexthopserverFunc,
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
			"nexthopport": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"nexthopfqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nexthopip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resaddresstype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"secure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnnexthopserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnnexthopserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnnexthopserverName := d.Get("name").(string)
	vpnnexthopserver := vpn.Vpnnexthopserver{
		Name:           d.Get("name").(string),
		Nexthopfqdn:    d.Get("nexthopfqdn").(string),
		Nexthopip:      d.Get("nexthopip").(string),
		Nexthopport:    d.Get("nexthopport").(int),
		Resaddresstype: d.Get("resaddresstype").(string),
		Secure:         d.Get("secure").(string),
	}

	_, err := client.AddResource("vpnnexthopserver", vpnnexthopserverName, &vpnnexthopserver)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnnexthopserverName)

	return readVpnnexthopserverFunc(ctx, d, meta)
}

func readVpnnexthopserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnnexthopserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnnexthopserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnnexthopserver state %s", vpnnexthopserverName)
	data, err := client.FindResource(service.Vpnnexthopserver.Type(), vpnnexthopserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnnexthopserver state %s", vpnnexthopserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("nexthopfqdn", data["nexthopfqdn"])
	d.Set("nexthopip", data["nexthopip"])
	setToInt("nexthopport", d, data["nexthopport"])
	d.Set("resaddresstype", data["resaddresstype"])
	d.Set("secure", data["secure"])

	return nil

}

func deleteVpnnexthopserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnnexthopserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnnexthopserverName := d.Id()
	err := client.DeleteResource(service.Vpnnexthopserver.Type(), vpnnexthopserverName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
