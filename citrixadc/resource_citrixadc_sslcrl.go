package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslcrl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcrlFunc,
		ReadContext:   readSslcrlFunc,
		UpdateContext: updateSslcrlFunc,
		DeleteContext: deleteSslcrlFunc,
		Schema: map[string]*schema.Schema{
			"basedn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"binary": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"binddn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacertfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cakeyfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crlname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"crlpath": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"day": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gencrl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"indexfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"refresh": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revoke": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslcrlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcrlFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslcrlName string
	if v, ok := d.GetOk("crlname"); ok {
		sslcrlName = v.(string)
	} else {
		sslcrlName = resource.PrefixedUniqueId("tf-sslcrl-")
		d.Set("crlname", sslcrlName)
	}
	sslcrl := ssl.Sslcrl{
		Basedn:     d.Get("basedn").(string),
		Binary:     d.Get("binary").(string),
		Binddn:     d.Get("binddn").(string),
		Cacert:     d.Get("cacert").(string),
		Cacertfile: d.Get("cacertfile").(string),
		Cakeyfile:  d.Get("cakeyfile").(string),
		Crlname:    d.Get("crlname").(string),
		Crlpath:    d.Get("crlpath").(string),
		Day:        d.Get("day").(int),
		Gencrl:     d.Get("gencrl").(string),
		Indexfile:  d.Get("indexfile").(string),
		Inform:     d.Get("inform").(string),
		Interval:   d.Get("interval").(string),
		Method:     d.Get("method").(string),
		Password:   d.Get("password").(string),
		Port:       d.Get("port").(int),
		Refresh:    d.Get("refresh").(string),
		Revoke:     d.Get("revoke").(string),
		Scope:      d.Get("scope").(string),
		Server:     d.Get("server").(string),
		Time:       d.Get("time").(string),
		Url:        d.Get("url").(string),
	}

	_, err := client.AddResource(service.Sslcrl.Type(), sslcrlName, &sslcrl)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslcrlName)

	return readSslcrlFunc(ctx, d, meta)
}

func readSslcrlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcrlFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcrlName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslcrl state %s", sslcrlName)
	data, err := client.FindResource(service.Sslcrl.Type(), sslcrlName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslcrl state %s", sslcrlName)
		d.SetId("")
		return nil
	}
	d.Set("crlname", data["crlname"])
	d.Set("basedn", data["basedn"])
	d.Set("binary", data["binary"])
	d.Set("binddn", data["binddn"])
	d.Set("cacert", data["cacert"])
	d.Set("cacertfile", data["cacertfile"])
	d.Set("cakeyfile", data["cakeyfile"])
	d.Set("crlname", data["crlname"])
	d.Set("crlpath", data["crlpath"])
	setToInt("day", d, data["day"])
	d.Set("gencrl", data["gencrl"])
	d.Set("indexfile", data["indexfile"])
	d.Set("inform", data["inform"])
	d.Set("interval", data["interval"])
	d.Set("method", data["method"])
	d.Set("password", data["password"])
	setToInt("port", d, data["port"])
	d.Set("refresh", data["refresh"])
	d.Set("revoke", data["revoke"])
	d.Set("scope", data["scope"])
	d.Set("server", data["server"])
	d.Set("time", data["time"])
	d.Set("url", data["url"])

	return nil

}

func updateSslcrlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslcrlFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcrlName := d.Get("crlname").(string)

	sslcrl := ssl.Sslcrl{
		Crlname: d.Get("crlname").(string),
	}
	hasChange := false
	if d.HasChange("basedn") {
		log.Printf("[DEBUG]  citrixadc-provider: Basedn has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Basedn = d.Get("basedn").(string)
		hasChange = true
	}
	if d.HasChange("binary") {
		log.Printf("[DEBUG]  citrixadc-provider: Binary has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Binary = d.Get("binary").(string)
		hasChange = true
	}
	if d.HasChange("binddn") {
		log.Printf("[DEBUG]  citrixadc-provider: Binddn has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Binddn = d.Get("binddn").(string)
		hasChange = true
	}
	if d.HasChange("cacert") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacert has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Cacert = d.Get("cacert").(string)
		hasChange = true
	}
	if d.HasChange("cacertfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacertfile has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Cacertfile = d.Get("cacertfile").(string)
		hasChange = true
	}
	if d.HasChange("cakeyfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Cakeyfile has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Cakeyfile = d.Get("cakeyfile").(string)
		hasChange = true
	}
	if d.HasChange("crlname") {
		log.Printf("[DEBUG]  citrixadc-provider: Crlname has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Crlname = d.Get("crlname").(string)
		hasChange = true
	}
	if d.HasChange("crlpath") {
		log.Printf("[DEBUG]  citrixadc-provider: Crlpath has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Crlpath = d.Get("crlpath").(string)
		hasChange = true
	}
	if d.HasChange("day") {
		log.Printf("[DEBUG]  citrixadc-provider: Day has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Day = d.Get("day").(int)
		hasChange = true
	}
	if d.HasChange("gencrl") {
		log.Printf("[DEBUG]  citrixadc-provider: Gencrl has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Gencrl = d.Get("gencrl").(string)
		hasChange = true
	}
	if d.HasChange("indexfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Indexfile has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Indexfile = d.Get("indexfile").(string)
		hasChange = true
	}
	if d.HasChange("inform") {
		log.Printf("[DEBUG]  citrixadc-provider: Inform has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Inform = d.Get("inform").(string)
		hasChange = true
	}
	if d.HasChange("interval") {
		log.Printf("[DEBUG]  citrixadc-provider: Interval has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Interval = d.Get("interval").(string)
		hasChange = true
	}
	if d.HasChange("method") {
		log.Printf("[DEBUG]  citrixadc-provider: Method has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Method = d.Get("method").(string)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("refresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Refresh has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Refresh = d.Get("refresh").(string)
		hasChange = true
	}
	if d.HasChange("revoke") {
		log.Printf("[DEBUG]  citrixadc-provider: Revoke has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Revoke = d.Get("revoke").(string)
		hasChange = true
	}
	if d.HasChange("scope") {
		log.Printf("[DEBUG]  citrixadc-provider: Scope has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Scope = d.Get("scope").(string)
		hasChange = true
	}
	if d.HasChange("server") {
		log.Printf("[DEBUG]  citrixadc-provider: Server has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Server = d.Get("server").(string)
		hasChange = true
	}
	if d.HasChange("time") {
		log.Printf("[DEBUG]  citrixadc-provider: Time has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Time = d.Get("time").(string)
		hasChange = true
	}
	if d.HasChange("url") {
		log.Printf("[DEBUG]  citrixadc-provider: Url has changed for sslcrl %s, starting update", sslcrlName)
		sslcrl.Url = d.Get("url").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Sslcrl.Type(), sslcrlName, &sslcrl)
		if err != nil {
			return diag.Errorf("Error updating sslcrl %s", sslcrlName)
		}
	}
	return readSslcrlFunc(ctx, d, meta)
}

func deleteSslcrlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcrlFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcrlName := d.Id()
	err := client.DeleteResource(service.Sslcrl.Type(), sslcrlName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
