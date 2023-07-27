package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

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
		CustomizeDiff: customizeDiff,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bundle": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certkey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"expirymonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fipskey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hsmkey": {
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
			"linkcertkeyname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
			"nodomaincheck": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"notificationperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocspstaplingcache": {
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
		Bundle:        d.Get("bundle").(string),
		Cert:          d.Get("cert").(string),
		Certkey:       d.Get("certkey").(string),
		Expirymonitor: d.Get("expirymonitor").(string),
		Fipskey:       d.Get("fipskey").(string),
		Hsmkey:        d.Get("hsmkey").(string),
		Inform:        d.Get("inform").(string),
		Key:           d.Get("key").(string),
		// This is always set to false on creation which effectively excludes it from the request JSON
		// Nodomaincheck is not an object attribute but a flag for the change operation
		// of the resource
		Nodomaincheck:      false,
		Notificationperiod: d.Get("notificationperiod").(int),
		Ocspstaplingcache:  d.Get("ocspstaplingcache").(bool),
		Passplain:          d.Get("passplain").(string),
		Password:           d.Get("password").(bool),
	}

	_, err := client.AddResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkey)
	if err != nil {
		return err
	}

	d.SetId(sslcertkeyName)

	if err := handleLinkedCertificate(d, client); err != nil {
		log.Printf("Error linking certificate during creation\n")
		err2 := deleteSslcertkeyFunc(d, meta)
		if err2 != nil {
			return fmt.Errorf("Delete error:%s while handling linked certificate error: %s", err2.Error(), err.Error())
		}
		return err
	}

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
	data, err := client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing sslcertkey state %s", sslcertkeyName)
		d.SetId("")
		return nil
	}
	d.Set("certkey", data["certkey"])
	d.Set("bundle", data["bundle"])
	d.Set("cert", data["cert"])
	d.Set("certkey", data["certkey"])
	// d.Set("expirymonitor", data["expirymonitor"])
	d.Set("fipskey", data["fipskey"])
	d.Set("hsmkey", data["hsmkey"])
	d.Set("inform", data["inform"])
	d.Set("key", data["key"])
	d.Set("linkcertkeyname", data["linkcertkeyname"])
	d.Set("nodomaincheck", data["nodomaincheck"])
	// d.Set("notificationperiod", data["notificationperiod"])
	d.Set("ocspstaplingcache", data["ocspstaplingcache"])
	// `passplain` and `password` are not returned by NITRO request
	// commenting out to avoid perpetual divergence between local and remote state
	//d.Set("passplain", data["passplain"])
	//d.Set("password", data["password"])

	return nil

}

func updateSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In updateSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Get("certkey").(string)

	sslcertkeyUpdate := ssl.Sslcertkey{
		Certkey: d.Get("certkey").(string),
	}
	sslcertkeyChange := ssl.Sslcertkey{
		Certkey: d.Get("certkey").(string),
	}
	hasUpdate := false //depending on which field changed, we have to use Update or Change API
	hasChange := false
	if d.HasChange("expirymonitor") {
		log.Printf("[DEBUG] netscaler-provider:  Expirymonitor has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyUpdate.Expirymonitor = d.Get("expirymonitor").(string)
		hasUpdate = true
	}
	if d.HasChange("notificationperiod") {
		log.Printf("[DEBUG] netscaler-provider:  Notificationperiod has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyUpdate.Notificationperiod = d.Get("notificationperiod").(int)
		hasUpdate = true
	}
	if d.HasChange("cert") {
		log.Printf("[DEBUG] netscaler-provider:  cert has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Cert = d.Get("cert").(string)
		hasChange = true
	}
	if d.HasChange("key") {
		log.Printf("[DEBUG] netscaler-provider:  key has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Key = d.Get("key").(string)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG] netscaler-provider:  password has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Password = d.Get("password").(bool)
		hasChange = true
	}
	if d.HasChange("fipskey") {
		log.Printf("[DEBUG] netscaler-provider:  fipskey has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Fipskey = d.Get("fipskey").(string)
		hasChange = true
	}
	if d.HasChange("hsmkey") {
		log.Printf("[DEBUG]  netscaler-provider: Hsmkey has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Hsmkey = d.Get("hsmkey").(string)
		hasChange = true
	}
	if d.HasChange("inform") {
		log.Printf("[DEBUG] netscaler-provider:  inform has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Inform = d.Get("inform").(string)
		hasChange = true
	}
	if d.HasChange("passplain") {
		log.Printf("[DEBUG] netscaler-provider:  passplain has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Passplain = d.Get("passplain").(string)
		hasChange = true
	}
	if d.HasChange("ocspstaplingcache") {
		log.Printf("[DEBUG]  netscaler-provider: Ocspstaplingcache has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyChange.Ocspstaplingcache = d.Get("ocspstaplingcache").(bool)
		hasChange = true
	}

	if hasUpdate {
		sslcertkeyUpdate.Expirymonitor = d.Get("expirymonitor").(string) //always expected by NITRO API
		_, err := client.UpdateResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkeyUpdate)
		if err != nil {
			return fmt.Errorf("Error updating sslcertkey %s", sslcertkeyName)
		}
	}
	// nodomaincheck is a flag for the change operation
	// therefore its value is always used for the operation
	sslcertkeyChange.Nodomaincheck = d.Get("nodomaincheck").(bool)
	if hasChange {

		_, err := client.ChangeResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkeyChange)
		if err != nil {
			return fmt.Errorf("Error changing sslcertkey %s", sslcertkeyName)
		}
	}

	if err := handleLinkedCertificate(d, client); err != nil {
		log.Printf("Error linking certificate during update\n")
		return err
	}

	return readSslcertkeyFunc(d, meta)
}

func handleLinkedCertificate(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG] netscaler-provider:  In handleLinkedCertificate")
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
	if err := unlinkCertificate(d, client); err != nil {
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

func unlinkCertificate(d *schema.ResourceData, client *service.NitroClient) error {
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

func deleteSslcertkeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In deleteSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client

	if err := unlinkCertificate(d, client); err != nil {
		return err
	}
	sslcertkeyName := d.Id()
	err := client.DeleteResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func customizeDiff(diff *schema.ResourceDiff, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In customizeDiff")
	o := diff.GetChangedKeysPrefix("")

	if len(o) == 1 && o[0] == "nodomaincheck" {
		log.Printf("Only nodomaincheck in diff")
		diff.Clear("nodomaincheck")
	}
	return nil
}
