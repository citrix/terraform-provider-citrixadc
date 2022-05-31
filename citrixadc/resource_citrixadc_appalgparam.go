package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcAppalgparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppalgparamFunc,
		Read:          readAppalgparamFunc,
		Update:        updateAppalgparamFunc,
		Delete:        deleteAppalgparamFunc,
		Schema: map[string]*schema.Schema{
			"pptpgreidletimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAppalgparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppalgparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var appalgparamName string
	// there is no primary key in appalgparam resource. Hence generate one for terraform state maintenance
	appalgparamName = resource.PrefixedUniqueId("tf-appalgparam-")
	appalgparam := network.Appalgparam{
		Pptpgreidletimeout: d.Get("pptpgreidletimeout").(int),
	}

	err := client.UpdateUnnamedResource(service.Appalgparam.Type(), &appalgparam)
	if err != nil {
		return err
	}

	d.SetId(appalgparamName)

	err = readAppalgparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appalgparam but we can't read it ??")
		return nil
	}
	return nil
}

func readAppalgparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppalgparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading appalgparam state")
	data, err := client.FindResource(service.Appalgparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appalgparam state")
		d.SetId("")
		return nil
	}
	val,_ := strconv.Atoi(data["pptpgreidletimeout"].(string))
	d.Set("pptpgreidletimeout", val)

	return nil

}

func updateAppalgparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppalgparamFunc")
	client := meta.(*NetScalerNitroClient).client

	appalgparam := network.Appalgparam{}
	hasChange := false
	if d.HasChange("pptpgreidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Pptpgreidletimeout has changed for appalgparam, starting update")
		appalgparam.Pptpgreidletimeout = d.Get("pptpgreidletimeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appalgparam.Type(), &appalgparam)
		if err != nil {
			return fmt.Errorf("Error updating appalgparam")
		}
	}
	return readAppalgparamFunc(d, meta)
}

func deleteAppalgparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppalgparamFunc")


	d.SetId("")

	return nil
}
