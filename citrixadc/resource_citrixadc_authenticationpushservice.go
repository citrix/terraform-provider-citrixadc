package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationpushservice() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationpushserviceFunc,
		Read:          readAuthenticationpushserviceFunc,
		Update:        updateAuthenticationpushserviceFunc,
		Delete:        deleteAuthenticationpushserviceFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"clientid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientsecret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customerid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refreshinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationpushserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Get("name").(string)
	authenticationpushservice := authentication.Authenticationpushservice{
		Clientid:        d.Get("clientid").(string),
		Clientsecret:    d.Get("clientsecret").(string),
		Customerid:      d.Get("customerid").(string),
		Name:            d.Get("name").(string),
		Refreshinterval: d.Get("refreshinterval").(int),
	}

	_, err := client.AddResource("authenticationpushservice", authenticationpushserviceName, &authenticationpushservice)
	if err != nil {
		return err
	}

	d.SetId(authenticationpushserviceName)

	err = readAuthenticationpushserviceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationpushservice but we can't read it ?? %s", authenticationpushserviceName)
		return nil
	}
	return nil
}

func readAuthenticationpushserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationpushservice state %s", authenticationpushserviceName)
	data, err := client.FindResource("authenticationpushservice", authenticationpushserviceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationpushservice state %s", authenticationpushserviceName)
		d.SetId("")
		return nil
	}
	d.Set("clientid", data["clientid"])
	//d.Set("clientsecret", data["clientsecret"]) different value is received each time
	d.Set("customerid", data["customerid"])
	d.Set("name", data["name"])
	d.Set("refreshinterval", data["refreshinterval"])

	return nil

}

func updateAuthenticationpushserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Get("name").(string)

	authenticationpushservice := authentication.Authenticationpushservice{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientid") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientid has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Clientid = d.Get("clientid").(string)
		hasChange = true
	}
	if d.HasChange("clientsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecret has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Clientsecret = d.Get("clientsecret").(string)
		hasChange = true
	}
	if d.HasChange("customerid") {
		log.Printf("[DEBUG]  citrixadc-provider: Customerid has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Customerid = d.Get("customerid").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("refreshinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Refreshinterval has changed for authenticationpushservice %s, starting update", authenticationpushserviceName)
		authenticationpushservice.Refreshinterval = d.Get("refreshinterval").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationpushservice", authenticationpushserviceName, &authenticationpushservice)
		if err != nil {
			return fmt.Errorf("Error updating authenticationpushservice %s", authenticationpushserviceName)
		}
	}
	return readAuthenticationpushserviceFunc(d, meta)
}

func deleteAuthenticationpushserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationpushserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpushserviceName := d.Id()
	err := client.DeleteResource("authenticationpushservice", authenticationpushserviceName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
