package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslcert() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcertFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"certfile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"reqfile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"certtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"cacert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cacertform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cakey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cakeyform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"caserial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keyfile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keyform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pempassphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subjectaltname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcertFunc(d *schema.ResourceData, meta interface{}) error {
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
		Days:           d.Get("days").(int),
		Keyfile:        d.Get("keyfile").(string),
		Keyform:        d.Get("keyform").(string),
		Pempassphrase:  d.Get("pempassphrase").(string),
		Reqfile:        d.Get("reqfile").(string),
		Subjectaltname: d.Get("subjectaltname").(string),
	}

	err := client.ActOnResource(service.Sslcert.Type(), &sslcert, "create")
	if err != nil {
		return err
	}

	d.SetId(sslcertName)

	return nil
}
