package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strconv"
)

func resourceCitrixAdcBridgetable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBridgetableFunc,
		Read:          readBridgetableFunc,
		Update:        updateBridgetableFunc,
		Delete:        deleteBridgetableFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vtep": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vxlan": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bridgeage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"devicevlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ifnum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vni": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBridgetableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBridgetableFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgetableName := fmt.Sprintf("%s,%s,%s", d.Get("mac").(string), strconv.Itoa(d.Get("vxlan").(int)), d.Get("vtep").(string))
	bridgetable := network.Bridgetable{
		Devicevlan: d.Get("devicevlan").(int),
		Mac:        d.Get("mac").(string),
		Vni:        d.Get("vni").(int),
		Vtep:       d.Get("vtep").(string),
		Vxlan:      d.Get("vxlan").(int),
		Ifnum:      d.Get("ifnum").(string),
		Nodeid:     d.Get("nodeid").(int),
		Vlan:       d.Get("vlan").(int),
	}
	
	_, err := client.AddResource(service.Bridgetable.Type(), bridgetableName, &bridgetable)
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("bridgeage"); ok {
		bridgetable2 := network.Bridgetable{
			Bridgeage:  d.Get("bridgeage").(int),
		}
		err1 := client.UpdateUnnamedResource(service.Bridgetable.Type(),&bridgetable2)
		if err1 != nil {
			return err1
		}
	}

	d.SetId(bridgetableName)

	err = readBridgetableFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this bridgetable but we can't read it ?? %s", bridgetableName)
		return nil
	}
	return nil
}

func readBridgetableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBridgetableFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgetableName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading bridgetable state %s", bridgetableName)
	findParams := service.FindParams{
		ResourceType: service.Bridgetable.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing bridgetable state %s", bridgetableName)
		d.SetId("")
		return nil
	}
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: Bridge table does not exist. Clearing state.")
		d.SetId("")
		return nil
	}
	foundIndex := -1
	for i, bridgetable := range dataArray {
		match := true
		if bridgetable["mac"] != d.Get("mac").(string) {
			match = false
		}
		if bridgetable["vxlan"] != strconv.Itoa(d.Get("vxlan").(int)) {
			match = false
		}
		if bridgetable["vtep"] != d.Get("vtep").(string) {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams bridgetable not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing bridgetable state %s", bridgetableName)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	//d.Set("bridgeage", data["bridgeage"])
	d.Set("devicevlan", data["devicevlan"])
	d.Set("ifnum", data["ifnum"])
	d.Set("mac", data["mac"])
	d.Set("nodeid", data["nodeid"])
	d.Set("vlan", data["vlan"])
	d.Set("vni", data["vni"])
	d.Set("vtep", data["vtep"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func updateBridgetableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateBridgetableFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgetableName := d.Id()

	bridgetable := network.Bridgetable{}
	hasChange := false
	if d.HasChange("bridgeage") {
		log.Printf("[DEBUG]  citrixadc-provider: Bridgeage has changed for bridgetable %s, starting update", bridgetableName)
		bridgetable.Bridgeage = d.Get("bridgeage").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Bridgetable.Type(), &bridgetable)
		if err != nil {
			return fmt.Errorf("Error updating bridgetable %s", bridgetableName)
		}
	}
	return readBridgetableFunc(d, meta)
}

func deleteBridgetableFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBridgetableFunc")
	client := meta.(*NetScalerNitroClient).client
	argsMap := make(map[string]string)

	argsMap["mac"] = url.QueryEscape(d.Get("mac").(string))
	argsMap["vtep"] = url.QueryEscape(d.Get("vtep").(string))
	argsMap["vxlan"] = strconv.Itoa(d.Get("vxlan").(int))
	argsMap["devicevlan"] = strconv.Itoa(d.Get("devicevlan").(int))
	err := client.DeleteResourceWithArgsMap(service.Bridgetable.Type(), "",argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
