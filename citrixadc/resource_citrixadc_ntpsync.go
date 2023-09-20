package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ntp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNtpsync() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNtpsyncFunc,
		Update:        updateNtpsyncFunc,
		Read:          readNtpsyncFunc,
		Delete:        deleteNtpsyncFunc,
		Schema: map[string]*schema.Schema{
			"state": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createNtpsyncFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNtpsyncFunc")
	ntpsyncName := resource.PrefixedUniqueId("tf-ntpsync-")
	client := meta.(*NetScalerNitroClient).client

	err := doNtpsyncChange(d, client)
	if err != nil {
		return err
	}

	d.SetId(ntpsyncName)

	err = readNtpsyncFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ntpsync but we can't read it ??")
		return nil
	}
	return nil
}

func readNtpsyncFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNtpsyncFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ntpsync state")
	data, err := client.FindResource(service.Ntpsync.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ntpsync state")
		d.SetId("")
		return nil
	}
	d.Set("state", data["state"].(string))

	return nil

}

func updateNtpsyncFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNtpSyncFunc")
	client := meta.(*NetScalerNitroClient).client

	hasChange := false
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: state has changed for ntpsync, starting update")
		hasChange = true
	}

	if hasChange {
		err := doNtpsyncChange(d, client)
		if err != nil {
			return err
		}
	}
	return readNtpsyncFunc(d, meta)
}
func doNtpsyncChange(d *schema.ResourceData, client *service.NitroClient) error {
	ntpsync := ntp.Ntpsync{}

	newstate := d.Get("state").(string)

	var err error
	// Enable action
	if newstate == "ENABLED" {
		err = client.ActOnResource(service.Ntpsync.Type(), &ntpsync, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err = client.ActOnResource(service.Ntpsync.Type(), &ntpsync, "disable")

	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}

	if err != nil {
		return err
	}
	return nil
}

func deleteNtpsyncFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNtpsyncFunc")
	// ntpsync does not support DELETE operation
	d.SetId("")

	return nil
}
