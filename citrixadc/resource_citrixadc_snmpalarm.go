package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSnmpalarm() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSnmpalarmFunc,
		Read:          readSnmpalarmFunc,
		Update:        updateSnmpalarmFunc,
		Delete:        deleteSnmpalarmFunc,
		Schema: map[string]*schema.Schema{
			"logging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"normalvalue": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"thresholdvalue": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trapname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSnmpalarmFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSnmpalarmFunc")
	client := meta.(*NetScalerNitroClient).client
	snmpalarmName := resource.PrefixedUniqueId("tf-snmpalarm-")

	// As time attriute can take zero value, we need to check if it is set or not
	// if we use the snmpalarm struct directly, it will set the time to 0 (provided we remove omitempty in struct in adc-nitro-go repo(https://github.com/netscaler/adc-nitro-go/blob/main/resource/config/snmp/snmpalarm.go)), even if not set by customers
	// So, removing the snmpalarm struct and adding the attributes directly
	snmpalarm := make(map[string]interface{})
	if v, ok := d.GetOk("logging"); ok {
		snmpalarm["logging"] = v.(string)
	}
	if v, ok := d.GetOk("normalvalue"); ok {
		snmpalarm["normalvalue"] = v.(int)
	}
	if v, ok := d.GetOk("severity"); ok {
		snmpalarm["severity"] = v.(string)
	}
	if v, ok := d.GetOk("state"); ok {
		snmpalarm["state"] = v.(string)
	}
	if v, ok := d.GetOk("thresholdvalue"); ok {
		snmpalarm["thresholdvalue"] = v.(int)
	}
	if v, ok := d.GetOk("time"); ok {
		snmpalarm["time"] = v.(int)
	}
	if v, ok := d.GetOk("trapname"); ok {
		snmpalarm["trapname"] = v.(string)
	}

	err := client.UpdateUnnamedResource(service.Snmpalarm.Type(), &snmpalarm)
	if err != nil {
		return err
	}

	d.SetId(snmpalarmName)

	err = readSnmpalarmFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this snmpalarm but we can't read it ??")
		return nil
	}
	return nil
}

func readSnmpalarmFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSnmpalarmFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading snmpalarm state")
	data, err := client.FindResource(service.Snmpalarm.Type(), d.Get("trapname").(string))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing snmpalarm state")
		d.SetId("")
		return nil
	}
	d.Set("trapname", data["trapname"])
	d.Set("logging", data["logging"])
	// d.Set("normalvalue", data["normalvalue"]) TODO: Not received from NetScaler
	d.Set("severity", data["severity"])
	d.Set("state", data["state"])
	d.Set("thresholdvalue", data["thresholdvalue"])
	d.Set("time", data["time"])
	d.Set("trapname", data["trapname"])

	return nil

}

func updateSnmpalarmFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSnmpalarmFunc")
	client := meta.(*NetScalerNitroClient).client

	snmpalarm := make(map[string]interface{})

	snmpalarm["trapname"] = d.Get("trapname").(string)

	hasChange := false
	stateChange := false
	if d.HasChange("logging") {
		log.Printf("[DEBUG]  citrixadc-provider: Logging has changed for snmpalarm, starting update")
		snmpalarm["logging"] = d.Get("logging").(string)
		hasChange = true
	}
	if d.HasChange("normalvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Normalvalue has changed for snmpalarm, starting update")
		snmpalarm["normalvalue"] = d.Get("normalvalue").(int)
		snmpalarm["thresholdvalue"] = d.Get("thresholdvalue").(int)
		hasChange = true
	}
	if d.HasChange("severity") {
		log.Printf("[DEBUG]  citrixadc-provider: Severity has changed for snmpalarm, starting update")
		snmpalarm["severity"] = d.Get("severity").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for snmpalarm, starting update")
		snmpalarm["state"] = d.Get("state").(string)
		stateChange = true
	}
	if d.HasChange("thresholdvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Thresholdvalue has changed for snmpalarm, starting update")
		snmpalarm["thresholdvalue"] = d.Get("thresholdvalue").(int)
		hasChange = true
	}
	if d.HasChange("time") {
		log.Printf("[DEBUG]  citrixadc-provider: Time has changed for snmpalarm, starting update")
		snmpalarm["time"] = d.Get("time").(int)
		hasChange = true
	}
	if stateChange {
		err := doSnmpalarmStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling snmpalarm %s", d.Get("trapname").(string))
		}
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Snmpalarm.Type(), &snmpalarm)
		if err != nil {
			return fmt.Errorf("Error updating snmpalarm")
		}
	}
	return readSnmpalarmFunc(d, meta)
}

func deleteSnmpalarmFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSnmpalarmFunc")
	// snmpalarm do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}

func doSnmpalarmStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doSnmpalarmStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes

	snmpalarm := make(map[string]interface{})

	snmpalarm["trapname"] = d.Get("trapname").(string)

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Snmpalarm.Type(), snmpalarm, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Snmpalarm.Type(), snmpalarm, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
