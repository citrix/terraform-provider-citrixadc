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

func resourceCitrixAdcSslcertkey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcertkeyFunc,
		ReadContext:   readSslcertkeyFunc,
		UpdateContext: updateSslcertkeyFunc,
		DeleteContext: deleteSslcertkeyFunc,
		CustomizeDiff: customizeDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
			"passplain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslcertkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		Nodomaincheck:     false,
		Ocspstaplingcache: d.Get("ocspstaplingcache").(bool),
		Passplain:         d.Get("passplain").(string),
		Password:          d.Get("password").(bool),
	}

	if raw := d.GetRawConfig().GetAttr("notificationperiod"); !raw.IsNull() {
		sslcertkey.Notificationperiod = intPtr(d.Get("notificationperiod").(int))
	}

	_, err := client.AddResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkey)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslcertkeyName)

	if _, ok := d.GetOk("linkcertkeyname"); ok {
		if err := handleLinkedCertificate(d, client); err != nil {
			log.Printf("Error linking certificate during creation\n")
			err2 := deleteSslcertkeyFunc(ctx, d, meta)
			if err2.HasError() {

				for _, d := range err2 {
					if d.Severity == diag.Error {
						return diag.Errorf("Delete error:%s while handling linked certificate error: %s", d.Summary, err.Error())
					}
				}
			}
			return diag.FromErr(err)
		}
	}

	return readSslcertkeyFunc(ctx, d, meta)
}

func readSslcertkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	// setToInt("notificationperiod", d, data["notificationperiod"])
	d.Set("ocspstaplingcache", data["ocspstaplingcache"])
	// `passplain` and `password` are not returned by NITRO request
	// commenting out to avoid perpetual divergence between local and remote state
	//d.Set("passplain", data["passplain"])
	//d.Set("password", data["password"])

	return nil

}

func updateSslcertkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In updateSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Get("certkey").(string)

	sslcertkeyUpdate := ssl.Sslcertkey{
		Certkey: d.Get("certkey").(string),
	}
	sslcertkeyChange := ssl.Sslcertkey{
		Certkey: d.Get("certkey").(string),
	}
	sslcertkeyClear := ssl.Sslcertkey{
		Certkey: d.Get("certkey").(string),
	}
	hasUpdate := false //depending on which field changed, we have to use Update or Change API
	hasChange := false
	hasClear := false
	if d.HasChange("expirymonitor") {
		log.Printf("[DEBUG] netscaler-provider:  Expirymonitor has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyUpdate.Expirymonitor = d.Get("expirymonitor").(string)
		hasUpdate = true
	}
	if d.HasChange("notificationperiod") {
		log.Printf("[DEBUG] netscaler-provider:  Notificationperiod has changed for sslcertkey %s, starting update", sslcertkeyName)
		sslcertkeyUpdate.Notificationperiod = intPtr(d.Get("notificationperiod").(int))
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
		sslcertkeyClear.Ocspstaplingcache = d.Get("ocspstaplingcache").(bool)
		hasClear = true
	}

	if hasUpdate {
		sslcertkeyUpdate.Expirymonitor = d.Get("expirymonitor").(string) //always expected by NITRO API
		_, err := client.UpdateResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkeyUpdate)
		if err != nil {
			return diag.Errorf("Error updating sslcertkey %s", sslcertkeyName)
		}
	}
	// nodomaincheck is a flag for the change operation
	// therefore its value is always used for the operation
	sslcertkeyChange.Nodomaincheck = d.Get("nodomaincheck").(bool)
	if hasChange {

		_, err := client.ChangeResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkeyChange)
		if err != nil {
			return diag.Errorf("Error changing sslcertkey %s", sslcertkeyName)
		}
	}

	if hasClear {

		err := client.ActOnResource(service.Sslcertkey.Type(), &sslcertkeyClear, "clear")
		if err != nil {
			return diag.Errorf("Error clearing sslcertkey %s", sslcertkeyName)
		}
	}

	if err := handleLinkedCertificate(d, client); err != nil {
		log.Printf("Error linking certificate during update\n")
		return diag.FromErr(err)
	}

	return readSslcertkeyFunc(ctx, d, meta)
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

func deleteSslcertkeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In deleteSslcertkeyFunc")
	client := meta.(*NetScalerNitroClient).client

	if err := unlinkCertificate(d, client); err != nil {
		return diag.FromErr(err)
	}
	sslcertkeyName := d.Id()
	err := client.DeleteResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func customizeDiff(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In customizeDiff")
	o := diff.GetChangedKeysPrefix("")

	if len(o) == 1 && o[0] == "nodomaincheck" {
		log.Printf("Only nodomaincheck in diff")
		diff.Clear("nodomaincheck")
	}
	return nil
}
