package netscaler

import (
	"github.com/chiradeep/go-nitro/config/ssl"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerSslcertkey() *schema.Resource {
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
	log.Printf("[DEBUG] netscaler-provider:  In createSslcertkeyFunc")
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
		Expirymonitor:      d.Get("expirymonitor").(string),
		Fipskey:            d.Get("fipskey").(string),
		Inform:             d.Get("inform").(string),
		Key:                d.Get("key").(string),
		Linkcertkeyname:    d.Get("linkcertkeyname").(string),
		Nodomaincheck:      d.Get("nodomaincheck").(bool),
		Notificationperiod: d.Get("notificationperiod").(int),
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
	log.Printf("[DEBUG] netscaler-provider:  In readSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading sslcertkey state %s", sslcertkeyName)
	data, err := client.FindResource(netscaler.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing sslcertkey state %s", sslcertkeyName)
		d.SetId("")
		return nil
	}
	d.Set("certkey", data["certkey"])
	d.Set("bundle", data["bundle"])
	d.Set("cert", data["cert"])
	d.Set("certkey", data["certkey"])
	d.Set("expirymonitor", data["expirymonitor"])
	d.Set("fipskey", data["fipskey"])
	d.Set("inform", data["inform"])
	d.Set("key", data["key"])
	d.Set("linkcertkeyname", data["linkcertkeyname"])
	d.Set("nodomaincheck", data["nodomaincheck"])
	d.Set("notificationperiod", data["notificationperiod"])
	d.Set("passplain", data["passplain"])
	d.Set("password", data["password"])

	return nil

}

func updateSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In updateSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Get("certkey").(string)

	sslcertkey := ssl.Sslcertkey{
		Certkey: d.Get("certkey").(string),
	}
	hasChange := false
	if d.HasChange("expirymonitor") {
		log.Printf("[DEBUG] netscaler-provider:  Expirymonitor has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Expirymonitor = d.Get("expirymonitor").(string)
		hasChange = true
	}
	if d.HasChange("notificationperiod") {
		log.Printf("[DEBUG] netscaler-provider:  Notificationperiod has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkey.Notificationperiod = d.Get("notificationperiod").(int)
		hasChange = true
	}
	sslcertkey.Expirymonitor = d.Get("expirymonitor").(string) //always expected by NITRO API

	if hasChange {
		_, err := client.UpdateResource(netscaler.Sslcertkey.Type(), sslcertkeyName, &sslcertkey)
		if err != nil {
			return fmt.Errorf("Error updating sslcertkey %s", sslcertkeyName)
		}
	}
	return readSslcertkeyFunc(d, meta)
}

func deleteSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In deleteSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Id()
	err := client.DeleteResource(netscaler.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
