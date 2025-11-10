package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNd6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNd6Func,
		ReadContext:   readNd6Func,
		DeleteContext: deleteNd6Func,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
				ForceNew: true,
				Computed: true,
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

func createNd6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	_, err := client.AddResource(service.Nd6.Type(), nd6Id, &nd6)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nd6Id)

	return readNd6Func(ctx, d, meta)
}

func readNd6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	setToInt("nodeid", d, data["nodeid"])
	setToInt("td", d, data["td"])
	setToInt("vlan", d, data["vlan"])
	d.Set("vtep", data["vtep"])
	setToInt("vxlan", d, data["vxlan"])

	return nil

}

func deleteNd6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNd6Func")
	client := meta.(*NetScalerNitroClient).client
	nd6Name := d.Id()
	args := make([]string, 0)
	if v, ok := d.GetOk("vlan"); ok {
		vlan := v.(int)
		args = append(args, fmt.Sprintf("vlan:%v", vlan))
	} else if v, ok := d.GetOk("vxlan"); ok {
		vxlan := v.(int)
		args = append(args, fmt.Sprintf("vxlan:%v", vxlan))
	} else {
		args = append(args, fmt.Sprintf("vlan:%v", 1))
	}
	if v, ok := d.GetOk("td"); ok {
		td := v.(int)
		args = append(args, fmt.Sprintf("td:%v", td))
	}
	err := client.DeleteResourceWithArgs(service.Nd6.Type(), nd6Name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
