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

func resourceCitrixAdcSslocspresponder() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslocspresponderFunc,
		ReadContext:   readSslocspresponderFunc,
		UpdateContext: updateSslocspresponderFunc,
		DeleteContext: deleteSslocspresponderFunc,
		Schema: map[string]*schema.Schema{
			"batchingdelay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"batchingdepth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cache": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertclientcert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ocspurlresolvetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"producedattimeskew": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"respondercert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resptimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signingcert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"trustresponder": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usenonce": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslocspresponderFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslocspresponderFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslocspresponderName string
	if v, ok := d.GetOk("name"); ok {
		sslocspresponderName = v.(string)
	} else {
		sslocspresponderName = resource.PrefixedUniqueId("tf-sslocspresponder-")
		d.Set("name", sslocspresponderName)
	}
	sslocspresponder := ssl.Sslocspresponder{
		Cache:            d.Get("cache").(string),
		Httpmethod:       d.Get("httpmethod").(string),
		Insertclientcert: d.Get("insertclientcert").(string),
		Name:             d.Get("name").(string),
		Respondercert:    d.Get("respondercert").(string),
		Signingcert:      d.Get("signingcert").(string),
		Trustresponder:   d.Get("trustresponder").(bool),
		Url:              d.Get("url").(string),
		Usenonce:         d.Get("usenonce").(string),
	}

	if raw := d.GetRawConfig().GetAttr("batchingdelay"); !raw.IsNull() {
		sslocspresponder.Batchingdelay = intPtr(d.Get("batchingdelay").(int))
	}
	if raw := d.GetRawConfig().GetAttr("batchingdepth"); !raw.IsNull() {
		sslocspresponder.Batchingdepth = intPtr(d.Get("batchingdepth").(int))
	}
	if raw := d.GetRawConfig().GetAttr("cachetimeout"); !raw.IsNull() {
		sslocspresponder.Cachetimeout = intPtr(d.Get("cachetimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ocspurlresolvetimeout"); !raw.IsNull() {
		sslocspresponder.Ocspurlresolvetimeout = intPtr(d.Get("ocspurlresolvetimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("producedattimeskew"); !raw.IsNull() {
		sslocspresponder.Producedattimeskew = intPtr(d.Get("producedattimeskew").(int))
	}
	if raw := d.GetRawConfig().GetAttr("resptimeout"); !raw.IsNull() {
		sslocspresponder.Resptimeout = intPtr(d.Get("resptimeout").(int))
	}

	_, err := client.AddResource(service.Sslocspresponder.Type(), sslocspresponderName, &sslocspresponder)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslocspresponderName)

	return readSslocspresponderFunc(ctx, d, meta)
}

func readSslocspresponderFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslocspresponderFunc")
	client := meta.(*NetScalerNitroClient).client
	sslocspresponderName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslocspresponder state %s", sslocspresponderName)
	data, err := client.FindResource(service.Sslocspresponder.Type(), sslocspresponderName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslocspresponder state %s", sslocspresponderName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	setToInt("batchingdelay", d, data["batchingdelay"])
	setToInt("batchingdepth", d, data["batchingdepth"])
	d.Set("cache", data["cache"])
	setToInt("cachetimeout", d, data["cachetimeout"])
	d.Set("httpmethod", data["httpmethod"])
	d.Set("insertclientcert", data["insertclientcert"])
	d.Set("name", data["name"])
	setToInt("ocspurlresolvetimeout", d, data["ocspurlresolvetimeout"])
	setToInt("producedattimeskew", d, data["producedattimeskew"])
	d.Set("respondercert", data["respondercert"])
	setToInt("resptimeout", d, data["resptimeout"])
	d.Set("signingcert", data["signingcert"])
	d.Set("trustresponder", data["trustresponder"])
	d.Set("url", data["url"])
	d.Set("usenonce", data["usenonce"])

	return nil

}

func updateSslocspresponderFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslocspresponderFunc")
	client := meta.(*NetScalerNitroClient).client
	sslocspresponderName := d.Get("name").(string)

	sslocspresponder := ssl.Sslocspresponder{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("batchingdelay") {
		log.Printf("[DEBUG]  citrixadc-provider: Batchingdelay has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Batchingdelay = intPtr(d.Get("batchingdelay").(int))
		hasChange = true
	}
	if d.HasChange("batchingdepth") {
		log.Printf("[DEBUG]  citrixadc-provider: Batchingdepth has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Batchingdepth = intPtr(d.Get("batchingdepth").(int))
		hasChange = true
	}
	if d.HasChange("cache") {
		log.Printf("[DEBUG]  citrixadc-provider: Cache has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Cache = d.Get("cache").(string)
		hasChange = true
	}
	if d.HasChange("cachetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachetimeout has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Cachetimeout = intPtr(d.Get("cachetimeout").(int))
		hasChange = true
	}
	if d.HasChange("httpmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpmethod has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Httpmethod = d.Get("httpmethod").(string)
		hasChange = true
	}
	if d.HasChange("insertclientcert") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertclientcert has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Insertclientcert = d.Get("insertclientcert").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("ocspurlresolvetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Ocspurlresolvetimeout has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Ocspurlresolvetimeout = intPtr(d.Get("ocspurlresolvetimeout").(int))
		hasChange = true
	}
	if d.HasChange("producedattimeskew") {
		log.Printf("[DEBUG]  citrixadc-provider: Producedattimeskew has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Producedattimeskew = intPtr(d.Get("producedattimeskew").(int))
		hasChange = true
	}
	if d.HasChange("respondercert") {
		log.Printf("[DEBUG]  citrixadc-provider: Respondercert has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Respondercert = d.Get("respondercert").(string)
		hasChange = true
	}
	if d.HasChange("resptimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Resptimeout has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Resptimeout = intPtr(d.Get("resptimeout").(int))
		hasChange = true
	}
	if d.HasChange("signingcert") {
		log.Printf("[DEBUG]  citrixadc-provider: Signingcert has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Signingcert = d.Get("signingcert").(string)
		hasChange = true
	}
	if d.HasChange("trustresponder") {
		log.Printf("[DEBUG]  citrixadc-provider: Trustresponder has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Trustresponder = d.Get("trustresponder").(bool)
		hasChange = true
	}
	if d.HasChange("url") {
		log.Printf("[DEBUG]  citrixadc-provider: Url has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Url = d.Get("url").(string)
		hasChange = true
	}
	if d.HasChange("usenonce") {
		log.Printf("[DEBUG]  citrixadc-provider: Usenonce has changed for sslocspresponder %s, starting update", sslocspresponderName)
		sslocspresponder.Usenonce = d.Get("usenonce").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Sslocspresponder.Type(), sslocspresponderName, &sslocspresponder)
		if err != nil {
			return diag.Errorf("Error updating sslocspresponder %s", sslocspresponderName)
		}
	}
	return readSslocspresponderFunc(ctx, d, meta)
}

func deleteSslocspresponderFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslocspresponderFunc")
	client := meta.(*NetScalerNitroClient).client
	sslocspresponderName := d.Id()
	err := client.DeleteResource(service.Sslocspresponder.Type(), sslocspresponderName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
