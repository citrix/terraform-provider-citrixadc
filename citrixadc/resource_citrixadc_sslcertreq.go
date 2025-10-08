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

func resourceCitrixAdcSslcertreq() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcertreqFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"reqfile": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"countryname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"organizationname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"statename": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"challengepassword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commonname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"companyname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digestmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fipskeyname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localityname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organizationunitname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pempassphrase": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subjectaltname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslcertreqFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcertreqFunc")
	client := meta.(*NetScalerNitroClient).client

	sslcertreqName := resource.PrefixedUniqueId("tf-sslcertreq-")
	sslcertreq := ssl.Sslcertreq{
		Challengepassword:    d.Get("challengepassword").(string),
		Commonname:           d.Get("commonname").(string),
		Companyname:          d.Get("companyname").(string),
		Countryname:          d.Get("countryname").(string),
		Digestmethod:         d.Get("digestmethod").(string),
		Emailaddress:         d.Get("emailaddress").(string),
		Fipskeyname:          d.Get("fipskeyname").(string),
		Keyfile:              d.Get("keyfile").(string),
		Keyform:              d.Get("keyform").(string),
		Localityname:         d.Get("localityname").(string),
		Organizationname:     d.Get("organizationname").(string),
		Organizationunitname: d.Get("organizationunitname").(string),
		Pempassphrase:        d.Get("pempassphrase").(string),
		Reqfile:              d.Get("reqfile").(string),
		Statename:            d.Get("statename").(string),
		Subjectaltname:       d.Get("subjectaltname").(string),
	}

	err := client.ActOnResource(service.Sslcertreq.Type(), &sslcertreq, "create")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslcertreqName)

	return nil
}
