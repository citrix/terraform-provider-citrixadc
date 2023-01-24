package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcArp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createArpFunc,
		Read:          readArpFunc,
		Delete:        deleteArpFunc,
		Schema: map[string]*schema.Schema{
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew:true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"all": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
			"ownernode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": &schema.Schema{
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
			"vtep": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vxlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createArpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createArpFunc")
	client := meta.(*NetScalerNitroClient).client
	arpName := d.Get("ipaddress").(string)
	arp := make(map[string]interface{})
	arp["ipaddress"] = d.Get("ipaddress").(string)
	arp["mac"] = d.Get("mac").(string)
	if v, ok := d.GetOk("ifnum"); ok {
		arp["ifnum"] = v.(string)
	}
	if v, ok := d.GetOk("td"); ok {
		arp["td"] = v.(int)
	}
	if v, ok := d.GetOk("vlan"); ok {
		arp["vlan"] = v.(int)
	}
	if v, ok := d.GetOk("vtep"); ok {
		arp["vtep"] = v.(string)
	}
	if v, ok := d.GetOk("vxlan"); ok {
		arp["vxlan"] = v.(int)
	}
	if v, ok := d.GetOk("ownernode"); ok {
		arp["ownernode"] = v.(int)
	}

	_, err := client.AddResource(service.Arp.Type(), arpName, &arp)
	if err != nil {
		return err
	}

	d.SetId(arpName)

	err = readArpFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this arp but we can't read it ?? %s", arpName)
		return nil
	}
	return nil
}

func readArpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readArpFunc")
	client := meta.(*NetScalerNitroClient).client
	arpName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading arp state %s", arpName)
	dataArr, err := client.FindAllResources(service.Arp.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing arp state %s", arpName)
		d.SetId("")
		return nil
	}

	if len(dataArr) == 0 {
		log.Printf("[WARN] citrixadc-provider: arp does not exist. Clearing state.")
		d.SetId("")
		return nil
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == arpName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceAllresources arp not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing arp state %s", arpName)
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]
	d.Set("all", data["all"])
	d.Set("ifnum", data["ifnum"])
	d.Set("ipaddress", data["ipaddress"])
	//d.Set("mac", data["mac"])
	d.Set("nodeid", data["nodeid"])
	d.Set("ownernode", data["ownernode"])
	d.Set("td", data["td"])
	d.Set("vlan", data["vlan"])
	d.Set("vtep", data["vtep"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func deleteArpFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteArpFunc")
	client := meta.(*NetScalerNitroClient).client
	arpName := d.Id()
	args := make([]string, 0)
	if v, ok := d.GetOk("td"); ok {
		td := v.(int)
		args = append(args, fmt.Sprintf("td:%v", td))
	}
	if v, ok := d.GetOk("all"); ok {
		all := v.(bool)
		args = append(args, fmt.Sprintf("all:%s", strconv.FormatBool(all)))
	}
	if v, ok := d.GetOk("ownernode"); ok {
		ownernode := v.(int)
		args = append(args, fmt.Sprintf("ownernode:%v", ownernode))
	}
	var err error
	if len(args) == 0 {
		err = client.DeleteResource(service.Arp.Type(), arpName)
	} else  if len(args) > 0 {
		err = client.DeleteResourceWithArgs(service.Arp.Type(), arpName,args)
	}	
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
