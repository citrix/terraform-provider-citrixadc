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

func resourceCitrixAdcSslcert() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcertFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"certfile": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"reqfile": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"certtype": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"cacert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cacertform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cakey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cakeyform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"caserial": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keyfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keyform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pempassphrase": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subjectaltname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcertFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcertFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslcertName string

	sslcertName = resource.PrefixedUniqueId("tf-sslcert-")

	sslcert := ssl.Sslcert{
		Cacert:         d.Get("cacert").(string),
		Cacertform:     d.Get("cacertform").(string),
		Cakey:          d.Get("cakey").(string),
		Cakeyform:      d.Get("cakeyform").(string),
		Caserial:       d.Get("caserial").(string),
		Certfile:       d.Get("certfile").(string),
		Certform:       d.Get("certform").(string),
		Certtype:       d.Get("certtype").(string),
		Keyfile:        d.Get("keyfile").(string),
		Keyform:        d.Get("keyform").(string),
		Pempassphrase:  d.Get("pempassphrase").(string),
		Reqfile:        d.Get("reqfile").(string),
		Subjectaltname: d.Get("subjectaltname").(string),
	}

	if raw := d.GetRawConfig().GetAttr("days"); !raw.IsNull() {
		sslcert.Days = intPtr(d.Get("days").(int))
	}

	err := client.ActOnResource(service.Sslcert.Type(), &sslcert, "create")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslcertName)

	return nil
}
