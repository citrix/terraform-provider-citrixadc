package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNd6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNd6Func,
		Read:          readNd6Func,
		Delete:        deleteNd6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"mac": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"neighbor": {
				Type:     schema.TypeString,
				Required: true,
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
			"td": {
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
			"vtep": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vxlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNd6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNd6Func")
	client := meta.(*NetScalerNitroClient).client
	nd6Id := d.Get("neighbor").(string)
	nd6 := make(map[string]interface{})
	nd6["neighbor"] = d.Get("neighbor").(string)
	nd6["mac"] = d.Get("mac").(string)
	if v, ok := d.GetOk("ifnum"); ok {
		nd6["ifnum"] = v.(string)
	}
	if v, ok := d.GetOk("td"); ok {
		nd6["td"] = v.(int)
	}
	if v, ok := d.GetOk("vlan"); ok {
		nd6["vlan"] = v.(int)
	}
	if v, ok := d.GetOk("vtep"); ok {
		nd6["vtep"] = v.(string)
	}
	if v, ok := d.GetOk("vxlan"); ok {
		nd6["vxlan"] = v.(int)
	}
	// nd6 := network.Nd6{
	// 	Ifnum:    d.Get("ifnum").(string),
	// 	Mac:      d.Get("mac").(string),
	// 	Neighbor: d.Get("neighbor").(string),
	// 	Nodeid:   d.Get("nodeid").(int),
	// 	Td:       d.Get("td").(int),
	// 	Vlan:     d.Get("vlan").(int),
	// 	Vtep:     d.Get("vtep").(string),
	// 	Vxlan:    d.Get("vxlan").(int),
	// }

	_, err := client.AddResource(service.Nd6.Type(), nd6Id, &nd6)
	if err != nil {
		return err
	}

	d.SetId(nd6Id)

	err = readNd6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nd6 but we can't read it ?? %s", nd6Id)
		return nil
	}
	return nil
}

func readNd6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNd6Func")
	client := meta.(*NetScalerNitroClient).client
	nd6Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nd6 state %s", nd6Name)
	dataArr, err := client.FindAllResources(service.Nd6.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nd6 state %s", nd6Name)
		d.SetId("")
		return nil
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["neighbor"] == nd6Name {
			foundIndex = i
		}
	}
	data := dataArr[foundIndex]
	d.Set("ifnum", data["ifnum"])
	d.Set("mac", data["mac"])
	d.Set("neighbor", data["neighbor"])
	d.Set("nodeid", data["nodeid"])
	d.Set("td", data["td"])
	d.Set("vlan", data["vlan"])
	d.Set("vtep", data["vtep"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func deleteNd6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNd6Func")
	client := meta.(*NetScalerNitroClient).client
	nd6Name := d.Id()
	args := make([]string, 0)
	if v, ok := d.GetOk("vlan"); ok {
		vlan := v.(int)
		args = append(args, fmt.Sprintf("vlan:%v", vlan))
	}
	if v, ok := d.GetOk("vxlan"); ok {
		vxlan := v.(int)
		args = append(args, fmt.Sprintf("vlan:%v", vxlan))
	}
	if v, ok := d.GetOk("td"); ok {
		td := v.(int)
		args = append(args, fmt.Sprintf("vlan:%v", td))
	}
	err := client.DeleteResourceWithArgs(service.Nd6.Type(), nd6Name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
