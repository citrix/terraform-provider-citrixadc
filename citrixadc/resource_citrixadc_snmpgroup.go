package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSnmpgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpgroupFunc,
		Read:          readSnmpgroupFunc,
		Update:        updateSnmpgroupFunc,
		Delete:        deleteSnmpgroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"securitylevel": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"readviewname": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createSnmpgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Get("name").(string)
	snmpgroup := snmp.Snmpgroup{
		Name:          d.Get("name").(string),
		Readviewname:  d.Get("readviewname").(string),
		Securitylevel: d.Get("securitylevel").(string),
	}

	_, err := client.AddResource(service.Snmpgroup.Type(), snmpgroupName, &snmpgroup)
	if err != nil {
		return err
	}

	d.SetId(snmpgroupName)

	err = readSnmpgroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpgroup but we can't read it ?? %s", snmpgroupName)
		return nil
	}
	return nil
}

func readSnmpgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpgroup state %s", snmpgroupName)

	dataArr, err := client.FindAllResources(service.Snmpgroup.Type())

	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: snmpgroup does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["name"].(string) == d.Get("name") {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceAllResources snmpgroup not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing snmpgroup state %s", snmpgroupName)
		d.SetId("")
		return nil
	}

	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpgroup state %s", snmpgroupName)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	d.Set("name", data["name"])
	d.Set("readviewname", data["readviewname"])
	d.Set("securitylevel", data["securitylevel"])

	return nil

}

func updateSnmpgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Get("name").(string)

	snmpgroup := snmp.Snmpgroup{
		Name:          d.Get("name").(string),
		Securitylevel: d.Get("securitylevel").(string),
		Readviewname:  d.Get("readviewname").(string),
	}

	hasChange := false
	if d.HasChange("readviewname") {
		log.Printf("[DEBUG]  citrixadc-provider: Readviewname has changed for snmpgroup %s, starting update", snmpgroupName)
		//snmpgroup.Readviewname = d.Get("readviewname").(string)
		hasChange = true
	}
	if d.HasChange("securitylevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Securitylevel has changed for snmpgroup %s, starting update", snmpgroupName)
		//snmpgroup.Securitylevel = d.Get("securitylevel").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpgroup.Type(), &snmpgroup)
		if err != nil {
			return fmt.Errorf("Error updating snmpgroup %s", snmpgroupName)
		}
	}
	return readSnmpgroupFunc(d, meta)
}

func deleteSnmpgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpgroupName := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("securitylevel:%s", d.Get("securitylevel").(string)))

	err := client.DeleteResourceWithArgs(service.Snmpgroup.Type(), snmpgroupName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
