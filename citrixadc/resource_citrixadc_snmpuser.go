package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSnmpuser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpuserFunc,
		Read:          readSnmpuserFunc,
		Update:        updateSnmpuserFunc,
		Delete:        deleteSnmpuserFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"authpasswd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privpasswd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSnmpuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Get("name").(string)

	snmpuser := snmp.Snmpuser{
		Authpasswd: d.Get("authpasswd").(string),
		Authtype:   d.Get("authtype").(string),
		Group:      d.Get("group").(string),
		Name:       d.Get("name").(string),
		Privpasswd: d.Get("privpasswd").(string),
		Privtype:   d.Get("privtype").(string),
	}

	_, err := client.AddResource(service.Snmpuser.Type(), snmpuserName, &snmpuser)
	if err != nil {
		return err
	}

	d.SetId(snmpuserName)

	err = readSnmpuserFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpuser but we can't read it ?? %s", snmpuserName)
		return nil
	}
	return nil
}

func readSnmpuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpuser state %s", snmpuserName)
	data, err := client.FindResource(service.Snmpuser.Type(), snmpuserName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpuser state %s", snmpuserName)
		d.SetId("")
		return nil
	}
	log.Printf("DATA: %v", data)
	d.Set("name", data["name"])
	//d.Set("authpasswd", data["authpasswd"])
	d.Set("authtype", data["authtype"])
	d.Set("group", data["group"])
	//d.Set("privpasswd", data["privpasswd"])
	d.Set("privtype", data["privtype"])

	return nil

}

func updateSnmpuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Get("name").(string)

	snmpuser := snmp.Snmpuser{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("authpasswd") {
		log.Printf("[DEBUG]  citrixadc-provider: Authpasswd has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Authpasswd = d.Get("authpasswd").(string)
		snmpuser.Authtype = d.Get("authtype").(string)
		hasChange = true
	}
	if d.HasChange("authtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtype has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Authtype = d.Get("authtype").(string)
		snmpuser.Authpasswd = d.Get("authpasswd").(string)
		hasChange = true
	}
	if d.HasChange("group") {
		log.Printf("[DEBUG]  citrixadc-provider: Group has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Group = d.Get("group").(string)
		hasChange = true
	}
	if d.HasChange("privpasswd") {
		log.Printf("[DEBUG]  citrixadc-provider: Privpasswd has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Privpasswd = d.Get("privpasswd").(string)
		snmpuser.Privtype = d.Get("privtype").(string)
		hasChange = true
	}
	if d.HasChange("privtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Privtype has changed for snmpuser %s, starting update", snmpuserName)
		snmpuser.Privtype = d.Get("privtype").(string)
		snmpuser.Privpasswd = d.Get("privpasswd").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Snmpuser.Type(), snmpuserName, &snmpuser)
		if err != nil {
			return fmt.Errorf("Error updating snmpuser %s", snmpuserName)
		}
	}
	return readSnmpuserFunc(d, meta)
}

func deleteSnmpuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpuserFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpuserName := d.Id()
	err := client.DeleteResource(service.Snmpuser.Type(), snmpuserName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
