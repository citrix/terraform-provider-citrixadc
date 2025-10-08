package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSslfipskey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslfipskeyFunc,
		ReadContext:   readSslfipskeyFunc,
		DeleteContext: deleteSslfipskeyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"curve": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exponent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fipskeyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"inform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"iv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keytype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"modulus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"wrapkeyname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslfipskeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslfipskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslfipskeyName = d.Get("fipskeyname").(string)

	sslfipskey := ssl.Sslfipskey{
		Curve:       d.Get("curve").(string),
		Exponent:    d.Get("exponent").(string),
		Fipskeyname: sslfipskeyName,
		Inform:      d.Get("inform").(string),
		Iv:          d.Get("iv").(string),
		Key:         d.Get("key").(string),
		Keytype:     d.Get("keytype").(string),
		Modulus:     d.Get("modulus").(int),
		Wrapkeyname: d.Get("wrapkeyname").(string),
	}

	err := client.ActOnResource(service.Sslfipskey.Type(), &sslfipskey, "create")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslfipskeyName)

	return readSslfipskeyFunc(ctx, d, meta)
}

func readSslfipskeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslfipskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslfipskeyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslfipskey state %s", sslfipskeyName)
	data, err := client.FindResource(service.Sslfipskey.Type(), sslfipskeyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslfipskey state %s", sslfipskeyName)
		d.SetId("")
		return nil
	}
	d.Set("fipskeyname", data["fipskeyname"])
	d.Set("curve", data["curve"])
	d.Set("exponent", data["exponent"])
	d.Set("fipskeyname", data["fipskeyname"])
	d.Set("inform", data["inform"])
	d.Set("iv", data["iv"])
	d.Set("key", data["key"])
	d.Set("keytype", data["keytype"])
	setToInt("modulus", d, data["modulus"])
	d.Set("wrapkeyname", data["wrapkeyname"])

	return nil

}

func deleteSslfipskeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslfipskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslfipskeyName := d.Id()
	err := client.DeleteResource(service.Sslfipskey.Type(), sslfipskeyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
