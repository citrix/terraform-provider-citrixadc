package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslcertkeyUpdate() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcertkeyUpdateFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"certkey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fipskey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"inform": {
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
			"nodomaincheck": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passplain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"linkcertkeyname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcertkeyUpdateFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In createSslcertkeyUpdateFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Get("certkey").(string)

	sslcertkey := ssl.Sslcertkey{
		Cert:          d.Get("cert").(string),
		Certkey:       d.Get("certkey").(string),
		Fipskey:       d.Get("fipskey").(string),
		Inform:        d.Get("inform").(string),
		Key:           d.Get("key").(string),
		Nodomaincheck: true,
		Passplain:     d.Get("passplain").(string),
		Password:      d.Get("password").(bool),
	}

	err := client.ActOnResource(service.Sslcertkey.Type(), &sslcertkey, "update")
	if err != nil {
		return err
	}

	d.SetId(sslcertkeyName)

	if _, ok := d.GetOk("linkcertkeyname"); ok {
		if err := handleLinkedCertificate_update(d, client); err != nil {
			return err
		}
	}

	return nil
}

func handleLinkedCertificate_update(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG] netscaler-provider:  In handleLinkedCertificate_update")
	sslcertkeyName := d.Get("certkey").(string)
	data, err := client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: Clearing sslcertkey state %s", sslcertkeyName)
		d.SetId("")
		return err
	}
	actualLinkedCertKeyname := data["linkcertkeyname"]
	configuredLinkedCertKeyname := d.Get("linkcertkeyname")

	// Check for noop conditions
	if actualLinkedCertKeyname == configuredLinkedCertKeyname {
		log.Printf("[DEBUG] netscaler-provider: actual and configured linked certificates identical \"%s\"", actualLinkedCertKeyname)
		return nil
	}

	if actualLinkedCertKeyname == nil && configuredLinkedCertKeyname == "" {
		log.Printf("[DEBUG] netscaler-provider: actual and configured linked certificates both empty ")
		return nil
	}

	// Fallthrough to rest of execution
	if err := unlinkCertificate_update(d, client); err != nil {
		return err
	}

	if configuredLinkedCertKeyname != "" {
		log.Printf("[DEBUG] netscaler-provider: Linking certkey \"%s\"", configuredLinkedCertKeyname)
		sslCertkey := ssl.Sslcertkey{
			Certkey:         data["certkey"].(string),
			Linkcertkeyname: configuredLinkedCertKeyname.(string),
		}
		if err := client.ActOnResource(service.Sslcertkey.Type(), &sslCertkey, "link"); err != nil {
			log.Printf("[ERROR] netscaler-provider: Error linking certificate \"%v\"", err)
			return err
		}
	} else {
		log.Printf("[DEBUG] netscaler-provider: configured linked certkey is empty, nothing to do")
	}
	return nil
}

func unlinkCertificate_update(d *schema.ResourceData, client *service.NitroClient) error {
	sslcertkeyName := d.Get("certkey").(string)
	data, err := client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: Clearing sslcertkey state %s", sslcertkeyName)
		d.SetId("")
		return err
	}

	actualLinkedCertKeyname := data["linkcertkeyname"]

	if actualLinkedCertKeyname != nil {
		log.Printf("[DEBUG] netscaler-provider: Unlinking certkey \"%s\"", actualLinkedCertKeyname)

		sslCertkey := ssl.Sslcertkey{
			Certkey: data["certkey"].(string),
		}
		if err := client.ActOnResource(service.Sslcertkey.Type(), &sslCertkey, "unlink"); err != nil {
			log.Printf("[ERROR] netscaler-provider: Error unlinking certificate \"%v\"", err)
			return err
		}
	} else {
		log.Printf("[DEBUG] netscaler-provider: actual linked certkey is nil, nothing to do")
	}
	return nil
}
