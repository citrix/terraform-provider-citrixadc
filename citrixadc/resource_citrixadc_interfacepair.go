package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"strconv"
	"log"
)

func resourceCitrixAdcInterfacepair() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createInterfacepairFunc,
		Read:          readInterfacepairFunc,
		Delete:        deleteInterfacepairFunc,
		Schema: map[string]*schema.Schema{
			"interface_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ifnum": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createInterfacepairFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createInterfacepairFunc")
	client := meta.(*NetScalerNitroClient).client
	interfacepairName := strconv.Itoa(d.Get("interface_id").(int))
	
	interfacepair := network.Interfacepair{
		Id:    d.Get("interface_id").(int),
		Ifnum: toStringList(d.Get("ifnum").([]interface{})),
	}

	_, err := client.AddResource(service.Interfacepair.Type(), interfacepairName, &interfacepair)
	if err != nil {
		return err
	}

	d.SetId(interfacepairName)

	err = readInterfacepairFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this interfacepair but we can't read it ?? %s", interfacepairName)
		return nil
	}
	return nil
}

func readInterfacepairFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readInterfacepairFunc")
	client := meta.(*NetScalerNitroClient).client
	interfacepairName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading interfacepair state %s", interfacepairName)
	data, err := client.FindResource(service.Interfacepair.Type(), interfacepairName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing interfacepair state %s", interfacepairName)
		d.SetId("")
		return nil
	}
	d.Set("id", data["id"])
	//d.Set("ifnum", data["ifnum"])

	return nil

}

func deleteInterfacepairFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteInterfacepairFunc")
	client := meta.(*NetScalerNitroClient).client
	interfacepairName := d.Id()
	err := client.DeleteResource(service.Interfacepair.Type(), interfacepairName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
