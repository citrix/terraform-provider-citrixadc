package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcNat64param() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNat64paramFunc,
		Read:          readNat64paramFunc,
		Update:        updateNat64paramFunc,
		Delete:        deleteNat64paramFunc,
		Schema: map[string]*schema.Schema{
			"nat64fragheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64ignoretos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64v6mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nat64zerochecksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNat64paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNat64paramFunc")
	client := meta.(*NetScalerNitroClient).client
	var nat64paramName string
	// there is no primary key in nat64param resource. Hence generate one for terraform state maintenance
	nat64paramName = resource.PrefixedUniqueId("tf-nat64param-")

	nat64param := network.Nat64param{
		Nat64fragheader:   d.Get("nat64fragheader").(string),
		Nat64ignoretos:    d.Get("nat64ignoretos").(string),
		Nat64v6mtu:        d.Get("nat64v6mtu").(int),
		Nat64zerochecksum: d.Get("nat64zerochecksum").(string),
		Td:                d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource("nat64param", &nat64param)
	if err != nil {
		return err
	}

	d.SetId(nat64paramName)

	err = readNat64paramFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nat64param but we can't read it ?? %s", nat64paramName)
		return nil
	}
	return nil
}

func readNat64paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNat64paramFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nat64param state")
	data, err := client.FindResource("nat64param","")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nat64param state")
		d.SetId("")
		return nil
	}
	d.Set("nat64fragheader", data["nat64fragheader"])
	d.Set("nat64ignoretos", data["nat64ignoretos"])
	val,_ := strconv.Atoi(data["nat64v6mtu"].(string))
	d.Set("nat64v6mtu", val)
	d.Set("nat64zerochecksum", data["nat64zerochecksum"])
	val,_ = strconv.Atoi(data["td"].(string))
	d.Set("td", val )

	return nil

}

func updateNat64paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNat64paramFunc")
	client := meta.(*NetScalerNitroClient).client

	nat64param := network.Nat64param{}
	hasChange := false
	if d.HasChange("nat64fragheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64fragheader has changed for nat64param, starting update")
		nat64param.Nat64fragheader = d.Get("nat64fragheader").(string)
		hasChange = true
	}
	if d.HasChange("nat64ignoretos") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64ignoretos has changed for nat64param, starting update")
		nat64param.Nat64ignoretos = d.Get("nat64ignoretos").(string)
		hasChange = true
	}
	if d.HasChange("nat64v6mtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64v6mtu has changed for nat64param, starting update")
		nat64param.Nat64v6mtu = d.Get("nat64v6mtu").(int)
		hasChange = true
	}
	if d.HasChange("nat64zerochecksum") {
		log.Printf("[DEBUG]  citrixadc-provider: Nat64zerochecksum has changed for nat64param, starting update")
		nat64param.Nat64zerochecksum = d.Get("nat64zerochecksum").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nat64param, starting update")
		nat64param.Td = d.Get("td").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("nat64param", &nat64param)
		if err != nil {
			return fmt.Errorf("Error updating nat64param")
		}
	}
	return readNat64paramFunc(d, meta)
}

func deleteNat64paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNat64paramFunc")

	d.SetId("")

	return nil
}
