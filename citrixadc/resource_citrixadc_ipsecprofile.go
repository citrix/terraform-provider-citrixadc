package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ipsec"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcIpsecprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIpsecprofileFunc,
		ReadContext:   readIpsecprofileFunc,
		DeleteContext: deleteIpsecprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"encalgo": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hashalgo": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ikeretryinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ikeversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"livenesscheckinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peerpublickey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"perfectforwardsecrecy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"privatekey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"psk": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publickey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"replaywindowsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retransmissiontime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createIpsecprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsecprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecprofileName := d.Get("name").(string)

	ipsecprofile := ipsec.Ipsecprofile{
		Encalgo:               toStringList(d.Get("encalgo").([]interface{})),
		Hashalgo:              toStringList(d.Get("hashalgo").([]interface{})),
		Ikeversion:            d.Get("ikeversion").(string),
		Name:                  d.Get("name").(string),
		Peerpublickey:         d.Get("peerpublickey").(string),
		Perfectforwardsecrecy: d.Get("perfectforwardsecrecy").(string),
		Privatekey:            d.Get("privatekey").(string),
		Psk:                   d.Get("psk").(string),
		Publickey:             d.Get("publickey").(string),
	}

	if raw := d.GetRawConfig().GetAttr("ikeretryinterval"); !raw.IsNull() {
		ipsecprofile.Ikeretryinterval = intPtr(d.Get("ikeretryinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("lifetime"); !raw.IsNull() {
		ipsecprofile.Lifetime = intPtr(d.Get("lifetime").(int))
	}
	if raw := d.GetRawConfig().GetAttr("livenesscheckinterval"); !raw.IsNull() {
		ipsecprofile.Livenesscheckinterval = intPtr(d.Get("livenesscheckinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("replaywindowsize"); !raw.IsNull() {
		ipsecprofile.Replaywindowsize = intPtr(d.Get("replaywindowsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("retransmissiontime"); !raw.IsNull() {
		ipsecprofile.Retransmissiontime = intPtr(d.Get("retransmissiontime").(int))
	}

	_, err := client.AddResource(service.Ipsecprofile.Type(), ipsecprofileName, &ipsecprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ipsecprofileName)

	return readIpsecprofileFunc(ctx, d, meta)
}

func readIpsecprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsecprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ipsecprofile state %s", ipsecprofileName)
	data, err := client.FindResource(service.Ipsecprofile.Type(), ipsecprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipsecprofile state %s", ipsecprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("encalgo", data["encalgo"])
	d.Set("hashalgo", data["hashalgo"])
	setToInt("ikeretryinterval", d, data["ikeretryinterval"])
	d.Set("ikeversion", data["ikeversion"])
	setToInt("lifetime", d, data["lifetime"])
	setToInt("livenesscheckinterval", d, data["livenesscheckinterval"])
	d.Set("name", data["name"])
	d.Set("peerpublickey", data["peerpublickey"])
	//d.Set("perfectforwardsecrecy", data["perfectforwardsecrecy"])
	d.Set("privatekey", data["privatekey"])
	//d.Set("psk", data["psk"])
	d.Set("publickey", data["publickey"])
	setToInt("replaywindowsize", d, data["replaywindowsize"])
	setToInt("retransmissiontime", d, data["retransmissiontime"])

	return nil

}

func deleteIpsecprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsecprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecprofileName := d.Id()
	err := client.DeleteResource(service.Ipsecprofile.Type(), ipsecprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
