package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSnmpoption() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpoptionFunc,
		Read:          readSnmpoptionFunc,
		Update:        updateSnmpoptionFunc,
		Delete:        deleteSnmpoptionFunc,  // Thought snmpoption resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"partitionnameintrap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmpset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmptraplogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmptraplogginglevel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSnmpoptionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpoptionFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpoptionName := resource.PrefixedUniqueId("tf-snmpoption-")
	
	snmpoption := snmp.Snmpoption{
		Partitionnameintrap:  d.Get("partitionnameintrap").(string),
		Snmpset:              d.Get("snmpset").(string),
		Snmptraplogging:      d.Get("snmptraplogging").(string),
		Snmptraplogginglevel: d.Get("snmptraplogginglevel").(string),
	}

	err := client.UpdateUnnamedResource(service.Snmpoption.Type(), &snmpoption)
	if err != nil {
		return err
	}

	d.SetId(snmpoptionName)

	err = readSnmpoptionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpoption but we can't read it ??")
		return nil
	}
	return nil
}

func readSnmpoptionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpoptionFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpoption state ")
	data, err := client.FindResource(service.Snmpoption.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpoption state")
		d.SetId("")
		return nil
	}
	d.Set("partitionnameintrap", data["partitionnameintrap"])
	d.Set("snmpset", data["snmpset"])
	d.Set("snmptraplogging", data["snmptraplogging"])
	d.Set("snmptraplogginglevel", data["snmptraplogginglevel"])

	return nil

}

func updateSnmpoptionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpoptionFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpoption := snmp.Snmpoption{}
	
	hasChange := false
	if d.HasChange("partitionnameintrap") {
		log.Printf("[DEBUG]  citrixadc-provider: Partitionnameintrap has changed for snmpoption, starting update")
		snmpoption.Partitionnameintrap = d.Get("partitionnameintrap").(string)
		hasChange = true
	}
	if d.HasChange("snmpset") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmpset has changed for snmpoption, starting update")
		snmpoption.Snmpset = d.Get("snmpset").(string)
		hasChange = true
	}
	if d.HasChange("snmptraplogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmptraplogging has changed for snmpoption, starting update")
		snmpoption.Snmptraplogging = d.Get("snmptraplogging").(string)
		hasChange = true
	}
	if d.HasChange("snmptraplogginglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmptraplogginglevel has changed for snmpoption, starting update")
		snmpoption.Snmptraplogginglevel = d.Get("snmptraplogginglevel").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpoption.Type(), &snmpoption)
		if err != nil {
			return fmt.Errorf("Error updating snmpoption")
		}
	}
	return readSnmpoptionFunc(d, meta)
}

func deleteSnmpoptionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpoptionFunc")
	// snmpoption do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
