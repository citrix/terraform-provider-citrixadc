package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSnmpview() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpviewFunc,
		Read:          readSnmpviewFunc,
		Update:        updateSnmpviewFunc,
		Delete:        deleteSnmpviewFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subtree": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSnmpviewFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpviewFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpviewName := d.Get("name").(string)
	snmpview := snmp.Snmpview{
		Name:    d.Get("name").(string),
		Subtree: d.Get("subtree").(string),
		Type:    d.Get("type").(string),
	}

	_, err := client.AddResource(service.Snmpview.Type(), snmpviewName, &snmpview)
	if err != nil {
		return err
	}

	d.SetId(snmpviewName)

	err = readSnmpviewFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpview but we can't read it ?? %s", snmpviewName)
		return nil
	}
	return nil
}

func readSnmpviewFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpviewFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpviewName := d.Id()
	subtree := d.Get("subtree").(string)
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpview state %s", snmpviewName)
	dataArr, err := client.FindAllResources(service.Snmpview.Type())

	foundIndex := -1
	for i, v := range dataArr {
		if v["name"].(string) == snmpviewName && v["subtree"].(string) == subtree {
			foundIndex = i
			break
		}
	}
	log.Printf("[DEBUG] citrixadc-provider: dataArr: %v", dataArr)
	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindAllResources %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllResources returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing snmpview %s", snmpviewName)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpview state %s", snmpviewName)
		d.SetId("")
		return nil
	}
	data := dataArr[foundIndex]
	d.Set("name", data["name"])
	d.Set("subtree", data["subtree"])
	d.Set("type", data["type"])

	return nil

}

func updateSnmpviewFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpviewFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpviewName := d.Get("name").(string)

	snmpview := snmp.Snmpview{
		Name: d.Get("name").(string),
		Subtree: d.Get("subtree").(string),
		Type:    d.Get("type").(string),
		
	}
	hasChange := false
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for snmpview %s, starting update", snmpviewName)
		snmpview.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpview.Type(), &snmpview)
		if err != nil {
			return fmt.Errorf("Error updating snmpview %s", snmpviewName)
		}
	}
	return readSnmpviewFunc(d, meta)
}

func deleteSnmpviewFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpviewFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpviewName := d.Id()
	subtree := d.Get("subtree").(string)

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("subtree:%s", subtree))

	err := client.DeleteResourceWithArgs(service.Snmpview.Type(), snmpviewName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
