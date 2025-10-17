package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnskey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnskeyFunc,
		ReadContext:   readDnskeyFunc,
		UpdateContext: updateDnskeyFunc,
		DeleteContext: deleteDnskeyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expires": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"filenameprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"keysize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"keytype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notificationperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privatekey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"publickey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"src": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"units1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zonename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnskeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnskeyName := d.Get("keyname").(string)
	dnskey := dns.Dnskey{
		Algorithm:      d.Get("algorithm").(string),
		Filenameprefix: d.Get("filenameprefix").(string),
		Keyname:        dnskeyName,
		Keytype:        d.Get("keytype").(string),
		Password:       d.Get("password").(string),
		Privatekey:     d.Get("privatekey").(string),
		Publickey:      d.Get("publickey").(string),
		Src:            d.Get("src").(string),
		Units1:         d.Get("units1").(string),
		Units2:         d.Get("units2").(string),
		Zonename:       d.Get("zonename").(string),
	}

	if raw := d.GetRawConfig().GetAttr("expires"); !raw.IsNull() {
		dnskey.Expires = intPtr(d.Get("expires").(int))
	}
	if raw := d.GetRawConfig().GetAttr("keysize"); !raw.IsNull() {
		dnskey.Keysize = intPtr(d.Get("keysize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("notificationperiod"); !raw.IsNull() {
		dnskey.Notificationperiod = intPtr(d.Get("notificationperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		dnskey.Ttl = intPtr(d.Get("ttl").(int))
	}

	_, err := client.AddResource(service.Dnskey.Type(), dnskeyName, &dnskey)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnskeyName)

	return readDnskeyFunc(ctx, d, meta)
}

func readDnskeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnskeyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnskey state %s", dnskeyName)
	data, err := client.FindResource(service.Dnskey.Type(), dnskeyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnskey state %s", dnskeyName)
		d.SetId("")
		return nil
	}
	d.Set("algorithm", data["algorithm"])
	setToInt("expires", d, data["expires"])
	d.Set("filenameprefix", data["filenameprefix"])
	d.Set("keyname", data["keyname"])
	setToInt("keysize", d, data["keysize"])
	d.Set("keytype", data["keytype"])
	setToInt("notificationperiod", d, data["notificationperiod"])
	d.Set("password", data["password"])
	d.Set("privatekey", data["privatekey"])
	d.Set("publickey", data["publickey"])
	d.Set("src", data["src"])
	setToInt("ttl", d, data["ttl"])
	d.Set("units1", data["units1"])
	d.Set("units2", data["units2"])
	d.Set("zonename", data["zonename"])

	return nil

}

func updateDnskeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnskeyName := d.Get("keyname").(string)

	dnskey := dns.Dnskey{
		Keyname: d.Get("keyname").(string),
	}
	hasChange := false
	if d.HasChange("algorithm") {
		log.Printf("[DEBUG]  citrixadc-provider: Algorithm has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Algorithm = d.Get("algorithm").(string)
		hasChange = true
	}
	if d.HasChange("expires") {
		log.Printf("[DEBUG]  citrixadc-provider: Expires has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Expires = intPtr(d.Get("expires").(int))
		hasChange = true
	}
	if d.HasChange("filenameprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Filenameprefix has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Filenameprefix = d.Get("filenameprefix").(string)
		hasChange = true
	}
	if d.HasChange("keysize") {
		log.Printf("[DEBUG]  citrixadc-provider: Keysize has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Keysize = intPtr(d.Get("keysize").(int))
		hasChange = true
	}
	if d.HasChange("keytype") {
		log.Printf("[DEBUG]  citrixadc-provider: Keytype has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Keytype = d.Get("keytype").(string)
		hasChange = true
	}
	if d.HasChange("notificationperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Notificationperiod has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Notificationperiod = intPtr(d.Get("notificationperiod").(int))
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("src") {
		log.Printf("[DEBUG]  citrixadc-provider: Src has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Src = d.Get("src").(string)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Ttl = intPtr(d.Get("ttl").(int))
		hasChange = true
	}
	if d.HasChange("units1") {
		log.Printf("[DEBUG]  citrixadc-provider: Units1 has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Units1 = d.Get("units1").(string)
		hasChange = true
	}
	if d.HasChange("units2") {
		log.Printf("[DEBUG]  citrixadc-provider: Units2 has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Units2 = d.Get("units2").(string)
		hasChange = true
	}
	if d.HasChange("zonename") {
		log.Printf("[DEBUG]  citrixadc-provider: Zonename has changed for dnskey %s, starting update", dnskeyName)
		dnskey.Zonename = d.Get("zonename").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnskey.Type(), dnskeyName, &dnskey)
		if err != nil {
			return diag.Errorf("Error updating dnskey %s", dnskeyName)
		}
	}
	return readDnskeyFunc(ctx, d, meta)
}

func deleteDnskeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnskeyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnskeyName := d.Id()
	err := client.DeleteResource(service.Dnskey.Type(), dnskeyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
