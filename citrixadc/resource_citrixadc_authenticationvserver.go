package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationvserverFunc,
		Read:          readAuthenticationvserverFunc,
		Update:        updateAuthenticationvserverFunc,
		Delete:        deleteAuthenticationvserverFunc,
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
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationdomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certkeynames": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failedlogintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxloginattempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"samesite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationvserverName := d.Get("name").(string)
	authenticationvserver := authentication.Authenticationvserver{
		Appflowlog:           d.Get("appflowlog").(string),
		Authentication:       d.Get("authentication").(string),
		Authenticationdomain: d.Get("authenticationdomain").(string),
		Certkeynames:         d.Get("certkeynames").(string),
		Comment:              d.Get("comment").(string),
		Failedlogintimeout:   d.Get("failedlogintimeout").(int),
		Ipv46:                d.Get("ipv46").(string),
		Maxloginattempts:     d.Get("maxloginattempts").(int),
		Name:                 d.Get("name").(string),
		Newname:              d.Get("newname").(string),
		Port:                 d.Get("port").(int),
		Range:                d.Get("range").(int),
		Samesite:             d.Get("samesite").(string),
		Servicetype:          d.Get("servicetype").(string),
		State:                d.Get("state").(string),
		Td:                   d.Get("td").(int),
	}

	_, err := client.AddResource(service.Authenticationvserver.Type(), authenticationvserverName, &authenticationvserver)
	if err != nil {
		return err
	}

	d.SetId(authenticationvserverName)

	err = readAuthenticationvserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationvserver but we can't read it ?? %s", authenticationvserverName)
		return nil
	}
	return nil
}

func readAuthenticationvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationvserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationvserver state %s", authenticationvserverName)
	data, err := client.FindResource(service.Authenticationvserver.Type(), authenticationvserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationvserver state %s", authenticationvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("authentication", data["authentication"])
	d.Set("authenticationdomain", data["authenticationdomain"])
	d.Set("certkeynames", data["certkeynames"])
	d.Set("comment", data["comment"])
	d.Set("failedlogintimeout", data["failedlogintimeout"])
	d.Set("ipv46", data["ipv46"])
	d.Set("maxloginattempts", data["maxloginattempts"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("port", data["port"])
	d.Set("range", data["range"])
	d.Set("samesite", data["samesite"])
	d.Set("servicetype", data["servicetype"])
	d.Set("td", data["td"])

	return nil

}

func updateAuthenticationvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationvserverName := d.Get("name").(string)

	authenticationvserver := authentication.Authenticationvserver{
		Name: d.Get("name").(string),
	}
	hasChange := false
	stateChange := false

	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authenticationdomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationdomain has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Authenticationdomain = d.Get("authenticationdomain").(string)
		hasChange = true
	}
	if d.HasChange("certkeynames") {
		log.Printf("[DEBUG]  citrixadc-provider: Certkeynames has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Certkeynames = d.Get("certkeynames").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("failedlogintimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Failedlogintimeout has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Failedlogintimeout = d.Get("failedlogintimeout").(int)
		hasChange = true
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv46 has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Ipv46 = d.Get("ipv46").(string)
		hasChange = true
	}
	if d.HasChange("maxloginattempts") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxloginattempts has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Maxloginattempts = d.Get("maxloginattempts").(int)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG]  citrixadc-provider: Range has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Range = d.Get("range").(int)
		hasChange = true
	}
	if d.HasChange("samesite") {
		log.Printf("[DEBUG]  citrixadc-provider: Samesite has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Samesite = d.Get("samesite").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicetype has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for authenticationvserver %s, starting update", authenticationvserverName)
		stateChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for authenticationvserver %s, starting update", authenticationvserverName)
		authenticationvserver.Td = d.Get("td").(int)
		hasChange = true
	}
	if stateChange {
		err := doAuthenticationvserverStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling cs vserver %s", authenticationvserverName)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Authenticationvserver.Type(), authenticationvserverName, &authenticationvserver)
		if err != nil {
			return fmt.Errorf("Error updating authenticationvserver %s", authenticationvserverName)
		}
	}
	return readAuthenticationvserverFunc(d, meta)
}

func deleteAuthenticationvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationvserverName := d.Id()
	err := client.DeleteResource(service.Authenticationvserver.Type(), authenticationvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
func doAuthenticationvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doLbvserverStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	authenticationvserver := authentication.Authenticationvserver{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Authenticationvserver.Type(), authenticationvserver, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Authenticationvserver.Type(), authenticationvserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}