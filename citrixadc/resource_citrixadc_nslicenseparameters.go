package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcNslicenseparameters() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNslicenseparametersFunc,
		Read:          readNslicenseparametersFunc,
		Update:        updateNslicenseparametersFunc,
		Delete:        deleteNslicenseparametersFunc,
		Schema: map[string]*schema.Schema{
			"alert1gracetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"alert2gracetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNslicenseparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslicenseparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	var nslicenseparametersName string
	// there is no primary key in nslicenseparameters resource. Hence generate one for terraform state maintenance
	nslicenseparametersName = resource.PrefixedUniqueId("tf-nslicenseparameters-")
	nslicenseparameters := ns.Nslicenseparameters{
		Alert1gracetimeout: d.Get("alert1gracetimeout").(int),
		Alert2gracetimeout: d.Get("alert2gracetimeout").(int),
	}

	err := client.UpdateUnnamedResource("nslicenseparameters", &nslicenseparameters)
	if err != nil {
		return err
	}

	d.SetId(nslicenseparametersName)

	err = readNslicenseparametersFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nslicenseparameters but we can't read it ??")
		return nil
	}
	return nil
}

func readNslicenseparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslicenseparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nslicenseparameters state")
	data, err := client.FindResource("nslicenseparameters", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslicenseparameters state")
		d.SetId("")
		return nil
	}
	log.Println(data)
	val, _ := strconv.Atoi(data["alert1gracetimeout"].(string))
	d.Set("alert1gracetimeout", val)
	val, _ = strconv.Atoi(data["alert2gracetimeout"].(string))
	d.Set("alert2gracetimeout", val)

	return nil

}

func updateNslicenseparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNslicenseparametersFunc")
	client := meta.(*NetScalerNitroClient).client

	nslicenseparameters := ns.Nslicenseparameters{
		Alert1gracetimeout: d.Get("alert1gracetimeout").(int),
	}
	if d.HasChange("alert2gracetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Alert2gracetimeout has changed for nslicenseparameters, starting update")
		nslicenseparameters.Alert2gracetimeout = d.Get("alert2gracetimeout").(int)
	}

	err := client.UpdateUnnamedResource("nslicenseparameters", &nslicenseparameters)
	if err != nil {
		return fmt.Errorf("Error updating nslicenseparameters")
	}

	return readNslicenseparametersFunc(d, meta)
}

func deleteNslicenseparametersFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslicenseparametersFunc")

	d.SetId("")

	return nil
}
