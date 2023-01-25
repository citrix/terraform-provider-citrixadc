package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcPtp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPtpFunc,
		Read:          readPtpFunc,
		Update:        updatePtpFunc,
		Delete:        deletePtpFunc,
		Schema: map[string]*schema.Schema{
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createPtpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPtpFunc")
	client := meta.(*NetScalerNitroClient).client
	var ptpName string
	// there is no primary key in ptp resource. Hence generate one for terraform state maintenance
	ptpName = resource.PrefixedUniqueId("tf-ptp-")
	ptp := network.Ptp{
		State: d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Ptp.Type(), &ptp)
	if err != nil {
		return err
	}

	d.SetId(ptpName)

	err = readPtpFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ptp but we can't read it ?? %s", ptpName)
		return nil
	}
	return nil
}

func readPtpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPtpFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ptp state")
	data, err := client.FindResource(service.Ptp.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ptp state")
		d.SetId("")
		return nil
	}
	d.Set("state", data["state"])

	return nil

}

func updatePtpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePtpFunc")
	client := meta.(*NetScalerNitroClient).client

	ptp := network.Ptp{}
	hasChange := false
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for ptp, starting update")
		ptp.State = d.Get("state").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ptp.Type(), &ptp)
		if err != nil {
			return fmt.Errorf("Error updating ptp ")
		}
	}
	return readPtpFunc(d, meta)
}

func deletePtpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePtpFunc")

	d.SetId("")

	return nil
}
