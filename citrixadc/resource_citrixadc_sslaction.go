package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcSslaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslactionFunc,
		Read:          readSslactionFunc,
		Delete:        deleteSslactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cacertgrpname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certfingerprintdigest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certfingerprintheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certhashheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certissuerheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certnotafterheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certnotbeforeheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certserialheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certsubjectheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cipherheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcertfingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcerthash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcertissuer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcertnotafter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcertnotbefore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcertserialnumber": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcertsubject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clientcertverification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"owasupport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sessionid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sessionidheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ssllogprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslactionFunc")
	client := meta.(*NetScalerNitroClient).client
	sslactionName := d.Get("name").(string)

	sslaction := ssl.Sslaction{
		Name:                   sslactionName,
		Cacertgrpname:          d.Get("cacertgrpname").(string),
		Certfingerprintdigest:  d.Get("certfingerprintdigest").(string),
		Certfingerprintheader:  d.Get("certfingerprintheader").(string),
		Certhashheader:         d.Get("certhashheader").(string),
		Certheader:             d.Get("certheader").(string),
		Certissuerheader:       d.Get("certissuerheader").(string),
		Certnotafterheader:     d.Get("certnotafterheader").(string),
		Certnotbeforeheader:    d.Get("certnotbeforeheader").(string),
		Certserialheader:       d.Get("certserialheader").(string),
		Certsubjectheader:      d.Get("certsubjectheader").(string),
		Cipher:                 d.Get("cipher").(string),
		Cipherheader:           d.Get("cipherheader").(string),
		Clientauth:             d.Get("clientauth").(string),
		Clientcert:             d.Get("clientcert").(string),
		Clientcertfingerprint:  d.Get("clientcertfingerprint").(string),
		Clientcerthash:         d.Get("clientcerthash").(string),
		Clientcertissuer:       d.Get("clientcertissuer").(string),
		Clientcertnotafter:     d.Get("clientcertnotafter").(string),
		Clientcertnotbefore:    d.Get("clientcertnotbefore").(string),
		Clientcertserialnumber: d.Get("clientcertserialnumber").(string),
		Clientcertsubject:      d.Get("clientcertsubject").(string),
		Clientcertverification: d.Get("clientcertverification").(string),
		Forward:                d.Get("forward").(string),
		Owasupport:             d.Get("owasupport").(string),
		Sessionid:              d.Get("sessionid").(string),
		Sessionidheader:        d.Get("sessionidheader").(string),
		Ssllogprofile:          d.Get("ssllogprofile").(string),
	}

	_, err := client.AddResource(service.Sslaction.Type(), sslactionName, &sslaction)
	if err != nil {
		return err
	}

	d.SetId(sslactionName)

	err = readSslactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslaction but we can't read it ?? %s", sslactionName)
		return nil
	}
	return nil
}

func readSslactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslactionFunc")
	client := meta.(*NetScalerNitroClient).client
	sslactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslaction state %s", sslactionName)
	data, err := client.FindResource(service.Sslaction.Type(), sslactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslaction state %s", sslactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("cacertgrpname", data["cacertgrpname"])
	d.Set("certfingerprintdigest", data["certfingerprintdigest"])
	d.Set("certfingerprintheader", data["certfingerprintheader"])
	d.Set("certhashheader", data["certhashheader"])
	d.Set("certheader", data["certheader"])
	d.Set("certissuerheader", data["certissuerheader"])
	d.Set("certnotafterheader", data["certnotafterheader"])
	d.Set("certnotbeforeheader", data["certnotbeforeheader"])
	d.Set("certserialheader", data["certserialheader"])
	d.Set("certsubjectheader", data["certsubjectheader"])
	d.Set("cipher", data["cipher"])
	d.Set("cipherheader", data["cipherheader"])
	d.Set("clientauth", data["clientauth"])
	d.Set("clientcert", data["clientcert"])
	d.Set("clientcertfingerprint", data["clientcertfingerprint"])
	d.Set("clientcerthash", data["clientcerthash"])
	d.Set("clientcertissuer", data["clientcertissuer"])
	d.Set("clientcertnotafter", data["clientcertnotafter"])
	d.Set("clientcertnotbefore", data["clientcertnotbefore"])
	d.Set("clientcertserialnumber", data["clientcertserialnumber"])
	d.Set("clientcertsubject", data["clientcertsubject"])
	d.Set("clientcertverification", data["clientcertverification"])
	d.Set("forward", data["forward"])
	d.Set("name", data["name"])
	d.Set("owasupport", data["owasupport"])
	d.Set("sessionid", data["sessionid"])
	d.Set("sessionidheader", data["sessionidheader"])
	d.Set("ssllogprofile", data["ssllogprofile"])

	return nil

}

func deleteSslactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslactionFunc")
	client := meta.(*NetScalerNitroClient).client
	sslactionName := d.Id()
	err := client.DeleteResource(service.Sslaction.Type(), sslactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
