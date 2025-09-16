package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcSnmpmanager() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpmanagerFunc,
		Read:          readSnmpmanagerFunc,
		Update:        updateSnmpmanagerFunc,
		Delete:        deleteSnmpmanagerFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domainresolveretry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func createSnmpmanagerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Get("ipaddress").(string)

	snmpmanager := snmp.Snmpmanager{
		Domainresolveretry: d.Get("domainresolveretry").(int),
		Ipaddress:          d.Get("ipaddress").(string),
		Netmask:            d.Get("netmask").(string),
	}

	_, err := client.AddResource(service.Snmpmanager.Type(), snmpmanagerName, &snmpmanager)
	if err != nil {
		return err
	}

	d.SetId(snmpmanagerName)

	err = readSnmpmanagerFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpmanager but we can't read it ?? %s", snmpmanagerName)
		return nil
	}
	return nil
}

func readSnmpmanagerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpmanager state %s", snmpmanagerName)

	dataArr, err := client.FindAllResources(service.Snmpmanager.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpmanager state %s", snmpmanagerName)
		d.SetId("")
		return nil
	}

	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: snmpmanager does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == snmpmanagerName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceAllResources snmpgroup not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing snmpmanager state %s", snmpmanagerName)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	d.Set("domainresolveretry", data["domainresolveretry"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("netmask", data["netmask"])

	return nil

}

func updateSnmpmanagerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Id()

	snmpmanager := snmp.Snmpmanager{
		Ipaddress: snmpmanagerName,
	}
	hasChange := false
	if d.HasChange("domainresolveretry") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainresolveretry has changed for snmpmanager %s, starting update", snmpmanagerName)
		snmpmanager.Domainresolveretry = d.Get("domainresolveretry").(int)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for snmpmanager %s, starting update", snmpmanagerName)
		snmpmanager.Netmask = d.Get("netmask").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpmanager.Type(), &snmpmanager)
		if err != nil {
			return fmt.Errorf("Error updating snmpmanager %s", snmpmanagerName)
		}
	}
	return readSnmpmanagerFunc(d, meta)
}

func deleteSnmpmanagerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpmanagerFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpmanagerName := d.Id()

	args := make([]string, 0)

	args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(d.Get("netmask").(string))))

	err := client.DeleteResourceWithArgs(service.Snmpmanager.Type(), snmpmanagerName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
