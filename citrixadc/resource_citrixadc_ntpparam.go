package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNtpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNtpparamFunc,
		Read:          readNtpparamFunc,
		Update:        updateNtpparamFunc,
		Delete:        deleteNtpparamFunc,
		Schema: map[string]*schema.Schema{
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autokeylogsec": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"revokelogsec": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trustedkey": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func createNtpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNtpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	ntpparamName := resource.PrefixedUniqueId("tf-ntpparam-")

	ntpparam := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok {
		ntpparam["authentication"] = v.(string)
	}
	if v, ok := d.GetOk("autokeylogsec"); ok {
		ntpparam["autokeylogsec"] = v.(int)
	}
	if v, ok := d.GetOk("revokelogsec"); ok {
		ntpparam["revokelogsec"] = v.(int)
	}
	if v, ok := d.GetOk("trustedkey"); ok {
		ntpparam["trustedkey"] = toIntegerList(v.([]interface{}))
	}

	err := client.UpdateUnnamedResource(service.Ntpparam.Type(), &ntpparam)
	if err != nil {
		return err
	}

	d.SetId(ntpparamName)

	err = readNtpparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ntpparam but we can't read it ??")
		return nil
	}
	return nil
}

func readNtpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNtpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ntpparam state")
	data, err := client.FindResource(service.Ntpparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ntpparam state")
		d.SetId("")
		return nil
	}
	d.Set("authentication", data["authentication"])
	d.Set("autokeylogsec", data["autokeylogsec"])
	d.Set("revokelogsec", data["revokelogsec"])
	d.Set("trustedkey", data["trustedkey"])

	return nil

}

func updateNtpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNtpparamFunc")
	client := meta.(*NetScalerNitroClient).client

	ntpparam := make(map[string]interface{})
	hasChange := false
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for ntpparam, starting update")
		ntpparam["authentication"] = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("autokeylogsec") {
		log.Printf("[DEBUG]  citrixadc-provider: Autokeylogsec has changed for ntpparam, starting update")
		ntpparam["autokeylogsec"] = d.Get("autokeylogsec").(int)
		hasChange = true
	}
	if d.HasChange("revokelogsec") {
		log.Printf("[DEBUG]  citrixadc-provider: Revokelogsec has changed for ntpparam, starting update")
		ntpparam["revokelogsec"] = d.Get("revokelogsec").(int)
		hasChange = true
	}
	if d.HasChange("trustedkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Trustedkey has changed for ntpparam, starting update")
		ntpparam["trustedkey"] = toIntegerList(d.Get("trustedkey").([]interface{}))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ntpparam.Type(), &ntpparam)
		if err != nil {
			return fmt.Errorf("Error updating ntpparam")
		}
	}
	return readNtpparamFunc(d, meta)
}

func deleteNtpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNtpparamFunc")
	// ntpparam does not support DELETE operation
	d.SetId("")

	return nil
}
