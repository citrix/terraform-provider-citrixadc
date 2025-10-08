package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcBridgetable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBridgetableFunc,
		ReadContext:   readBridgetableFunc,
		UpdateContext: updateBridgetableFunc,
		DeleteContext: deleteBridgetableFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"mac": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vtep": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vxlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bridgeage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"devicevlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ifnum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vni": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBridgetableFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	if _, ok := d.GetOk("bridgeage"); ok {
		bridgetable2 := network.Bridgetable{
			Bridgeage: d.Get("bridgeage").(int),
		}
		err1 := client.UpdateUnnamedResource(service.Bridgetable.Type(), &bridgetable2)
		if err1 != nil {
			return diag.FromErr(err1)
		}
	}

	d.SetId(bridgetableName)

	return readBridgetableFunc(ctx, d, meta)
}

func readBridgetableFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	//setToInt("bridgeage", d, data["bridgeage"])
	setToInt("devicevlan", d, data["devicevlan"])
	d.Set("ifnum", data["ifnum"])
	d.Set("mac", data["mac"])
	setToInt("nodeid", d, data["nodeid"])
	setToInt("vlan", d, data["vlan"])
	setToInt("vni", d, data["vni"])
	d.Set("vtep", data["vtep"])
	setToInt("vxlan", d, data["vxlan"])

	return nil

}

func updateBridgetableFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("Error updating bridgetable %s", bridgetableName)
		}
	}
	return readBridgetableFunc(ctx, d, meta)
}

func deleteBridgetableFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBridgetableFunc")
	client := meta.(*NetScalerNitroClient).client
	argsMap := make(map[string]string)

	argsMap["mac"] = url.QueryEscape(d.Get("mac").(string))
	argsMap["vtep"] = url.QueryEscape(d.Get("vtep").(string))
	argsMap["vxlan"] = strconv.Itoa(d.Get("vxlan").(int))
	argsMap["devicevlan"] = strconv.Itoa(d.Get("devicevlan").(int))
	err := client.DeleteResourceWithArgsMap(service.Bridgetable.Type(), "", argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
