package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ssl"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSslcertkey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcertkeyFunc,
		Read:          readSslcertkeyFunc,
		Update:        updateSslcertkeyFunc,
		Delete:        deleteSslcertkeyFunc,
		Schema: map[string]*schema.Schema{
			"bundle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deletefromdevice": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"expirymonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fipskey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hsmkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linkcertkeyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodomaincheck": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"notificationperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocspstaplingcache": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"passplain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslcertkeyName string
	if v, ok := d.GetOk("certkey"); ok {
		sslcertkeyName = v.(string)
	} else {
		sslcertkeyName = resource.PrefixedUniqueId("tf-sslcertkey-")
		d.Set("certkey", sslcertkeyName)
	}
	sslcertkey := ssl.Sslcertkey{
		Bundle:             d.Get("bundle").(string),
		Cert:               d.Get("cert").(string),
		Certkey:            d.Get("certkey").(string),
		Deletefromdevice:   d.Get("deletefromdevice").(bool),
		Expirymonitor:      d.Get("expirymonitor").(string),
		Fipskey:            d.Get("fipskey").(string),
		Hsmkey:             d.Get("hsmkey").(string),
		Inform:             d.Get("inform").(string),
		Key:                d.Get("key").(string),
		Linkcertkeyname:    d.Get("linkcertkeyname").(string),
		Nodomaincheck:      d.Get("nodomaincheck").(bool),
		Notificationperiod: d.Get("notificationperiod").(int),
		Ocspstaplingcache:  d.Get("ocspstaplingcache").(bool),
		Passplain:          d.Get("passplain").(string),
		Password:           d.Get("password").(bool),
	}

	_, err := client.AddResource(netscaler.Sslcertkey.Type(), sslcertkeyName, &sslcertkey)
	if err != nil {
		return err
	}

	d.SetId(sslcertkeyName)

	err = readSslcertkeyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslcertkey but we can't read it ?? %s", sslcertkeyName)
		return nil
	}
	return nil
}

func readSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslcertkey state %s", sslcertkeyName)
	data, err := client.FindResource(netscaler.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslcertkey state %s", sslcertkeyName)
		d.SetId("")
		return nil
	}
	d.Set("certkey", data["certkey"])
	d.Set("bundle", data["bundle"])
	d.Set("cert", data["cert"])
	d.Set("certkey", data["certkey"])
	d.Set("deletefromdevice", data["deletefromdevice"])
	d.Set("expirymonitor", data["expirymonitor"])
	d.Set("fipskey", data["fipskey"])
	d.Set("hsmkey", data["hsmkey"])
	d.Set("inform", data["inform"])
	d.Set("key", data["key"])
	d.Set("linkcertkeyname", data["linkcertkeyname"])
	d.Set("nodomaincheck", data["nodomaincheck"])
	d.Set("notificationperiod", data["notificationperiod"])
	d.Set("ocspstaplingcache", data["ocspstaplingcache"])
	d.Set("passplain", data["passplain"])
	d.Set("password", data["password"])

	return nil

}

func updateSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Get("certkey").(string)

	sslcertkey := ssl.Sslcertkey{
		Certkey: d.Get("certkey").(string),
	}
	hasChange := false
	if d.HasChange("bundle") {
		log.Printf("[DEBUG]  citrixadc-provider: Bundle has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Bundle = d.Get("bundle").(string)
		hasChange = true
	}
	if d.HasChange("cert") {
		log.Printf("[DEBUG]  citrixadc-provider: Cert has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Cert = d.Get("cert").(string)
		hasChange = true
	}
	if d.HasChange("certkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Certkey has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Certkey = d.Get("certkey").(string)
		hasChange = true
	}
	if d.HasChange("deletefromdevice") {
		log.Printf("[DEBUG]  citrixadc-provider: Deletefromdevice has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Deletefromdevice = d.Get("deletefromdevice").(bool)
		hasChange = true
	}
	if d.HasChange("expirymonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Expirymonitor has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Expirymonitor = d.Get("expirymonitor").(string)
		hasChange = true
	}
	if d.HasChange("fipskey") {
		log.Printf("[DEBUG]  citrixadc-provider: Fipskey has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Fipskey = d.Get("fipskey").(string)
		hasChange = true
	}
	if d.HasChange("hsmkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Hsmkey has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Hsmkey = d.Get("hsmkey").(string)
		hasChange = true
	}
	if d.HasChange("inform") {
		log.Printf("[DEBUG]  citrixadc-provider: Inform has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Inform = d.Get("inform").(string)
		hasChange = true
	}
	if d.HasChange("key") {
		log.Printf("[DEBUG]  citrixadc-provider: Key has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Key = d.Get("key").(string)
		hasChange = true
	}
	if d.HasChange("linkcertkeyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Linkcertkeyname has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Linkcertkeyname = d.Get("linkcertkeyname").(string)
		hasChange = true
	}
	if d.HasChange("nodomaincheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodomaincheck has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Nodomaincheck = d.Get("nodomaincheck").(bool)
		hasChange = true
	}
	if d.HasChange("notificationperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Notificationperiod has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Notificationperiod = d.Get("notificationperiod").(int)
		hasChange = true
	}
	if d.HasChange("ocspstaplingcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Ocspstaplingcache has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Ocspstaplingcache = d.Get("ocspstaplingcache").(bool)
		hasChange = true
	}
	if d.HasChange("passplain") {
		log.Printf("[DEBUG]  citrixadc-provider: Passplain has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Passplain = d.Get("passplain").(string)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Password = d.Get("password").(bool)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Sslcertkey.Type(), sslcertkeyName, &sslcertkey)
		if err != nil {
			return fmt.Errorf("Error updating sslcertkey %s", sslcertkeyName)
		}
	}
	return readSslcertkeyFunc(d, meta)
}

func deleteSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Id()
	err := client.DeleteResource(netscaler.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
