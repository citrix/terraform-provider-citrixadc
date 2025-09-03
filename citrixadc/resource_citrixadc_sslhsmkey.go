package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslhsmkey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslhsmkeyFunc,
		Read:          readSslhsmkeyFunc,
		Delete:        deleteSslhsmkeyFunc,
		Schema: map[string]*schema.Schema{
			"hsmkeyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hsmtype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"keystore": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"serialnum": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSslhsmkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslhsmkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslhsmkeyName = d.Get("hsmkeyname").(string)
	sslhsmkey := ssl.Sslhsmkey{
		Hsmkeyname: sslhsmkeyName,
		Hsmtype:    d.Get("hsmtype").(string),
		Key:        d.Get("key").(string),
		Keystore:   d.Get("keystore").(string),
		Password:   d.Get("password").(string),
		Serialnum:  d.Get("serialnum").(string),
	}

	_, err := client.AddResource(service.Sslhsmkey.Type(), sslhsmkeyName, &sslhsmkey)
	if err != nil {
		return err
	}

	d.SetId(sslhsmkeyName)

	err = readSslhsmkeyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslhsmkey but we can't read it ?? %s", sslhsmkeyName)
		return nil
	}
	return nil
}

func readSslhsmkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslhsmkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslhsmkeyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslhsmkey state %s", sslhsmkeyName)
	data, err := client.FindResource(service.Sslhsmkey.Type(), sslhsmkeyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslhsmkey state %s", sslhsmkeyName)
		d.SetId("")
		return nil
	}
	d.Set("hsmkeyname", data["hsmkeyname"])
	d.Set("hsmtype", data["hsmtype"])
	d.Set("key", data["key"])
	d.Set("keystore", data["keystore"])
	d.Set("password", d.Get("password").(string))
	d.Set("serialnum", data["serialnum"])

	return nil

}

func deleteSslhsmkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslhsmkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslhsmkeyName := d.Id()

	args := make(map[string]string)

	if val, ok := d.GetOk("hsmtype"); ok {
		args["hsmtype"] = val.(string)
	}
	if val, ok := d.GetOk("key"); ok {
		args["key"] = val.(string)
	}
	if val, ok := d.GetOk("keystore"); ok {
		args["keystore"] = val.(string)
	}
	if val, ok := d.GetOk("password"); ok {
		args["password"] = val.(string)
	}
	if val, ok := d.GetOk("serialnum"); ok {
		args["serialnum"] = val.(string)
	}

	err := client.DeleteResourceWithArgsMap(service.Sslhsmkey.Type(), sslhsmkeyName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
