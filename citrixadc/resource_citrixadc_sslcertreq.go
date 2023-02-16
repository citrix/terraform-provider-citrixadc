package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslcertreq() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcertreqFunc,
		Read:          schema.Noop,
		Delete:        deleteSslcertreqFunc,
		Schema: map[string]*schema.Schema{
			"reqfile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"countryname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"organizationname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"statename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"challengepassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commonname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"companyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digestmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fipskeyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyfile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localityname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organizationunitname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pempassphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subjectaltname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslcertreqFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(sslcertreqName)

	return nil
}

func deleteSslcertreqFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcertreqFunc")

	d.SetId("")

	return nil
}
