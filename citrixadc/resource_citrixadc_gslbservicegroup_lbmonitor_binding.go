package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcGslbservicegroup_lbmonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbservicegroup_lbmonitor_bindingFunc,
		Read:          readGslbservicegroup_lbmonitor_bindingFunc,
		Delete:        deleteGslbservicegroup_lbmonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"monitor_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"siteprefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createGslbservicegroup_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbservicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupname := d.Get("servicegroupname")
	monitor_name := d.Get("monitor_name")
	
	bindingId := fmt.Sprintf("%s,%s", servicegroupname, monitor_name)
	gslbservicegroup_lbmonitor_binding := gslb.Gslbservicegrouplbmonitorbinding{
		Hashid:           d.Get("hashid").(int),
		Monitorname:      d.Get("monitor_name").(string),
		Monstate:         d.Get("monstate").(string),
		Passive:          d.Get("passive").(bool),
		Port:             d.Get("port").(int),
		Publicip:         d.Get("publicip").(string),
		Publicport:       d.Get("publicport").(int),
		Servicegroupname: d.Get("servicegroupname").(string),
		Siteprefix:       d.Get("siteprefix").(string),
		State:            d.Get("state").(string),
		Weight:           d.Get("weight").(int),
	}

	err := client.UpdateUnnamedResource("gslbservicegroup_lbmonitor_binding", &gslbservicegroup_lbmonitor_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbservicegroup_lbmonitor_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservicegroup_lbmonitor_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbservicegroup_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbservicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicegroupname := idSlice[0]
	monitor_name := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservicegroup_lbmonitor_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbservicegroup_lbmonitor_binding",
		ResourceName:             servicegroupname,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservicegroup_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["monitor_name"].(string) == monitor_name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor_name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservicegroup_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("hashid", data["hashid"])
	d.Set("monitor_name", data["monitor_name"])
	d.Set("monstate", data["monstate"])
	d.Set("passive", data["passive"])
	d.Set("port", data["port"])
	d.Set("publicip", data["publicip"])
	d.Set("publicport", data["publicport"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("siteprefix", data["siteprefix"])
	d.Set("state", data["state"])
	d.Set("weight", data["weight"])

	return nil

}

func deleteGslbservicegroup_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbservicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	monitor_name := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("monitor_name:%s", monitor_name))

	err := client.DeleteResourceWithArgs("gslbservicegroup_lbmonitor_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
