package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mitchellh/mapstructure"

	"log"
)

func resourceCitrixAdcRnatClear() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRnatClearFunc,
		Read:          readRnatClearFunc,
		Update:        updateRnatClearFunc,
		Delete:        deleteRnatClearFunc,
		Schema: map[string]*schema.Schema{
			"rnatsname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"rnat": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aclname": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"natip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"natip2": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"netmask": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirectport": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"td": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createRnatClearFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createRnatClearFunc")

	var rnatName string
	if v, ok := d.GetOk("rnatsname"); ok {
		rnatName = v.(string)
	} else {
		rnatName = resource.PrefixedUniqueId("tf-rnat-")
		d.Set("rnatsname", rnatName)
	}
	rnats := d.Get("rnat").(*schema.Set).List()
	for _, val := range rnats {
		rnat := val.(map[string]interface{})
		_ = createSingleRnat(rnat, meta)
	}

	d.SetId(rnatName)

	return nil
}

func readRnatClearFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readRnatClearFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading rnat state %s", rnatName)

	data, _ := client.FindAllResources(service.Rnat.Type())
	rnats := make([]map[string]interface{}, len(data))
	for i, a := range data {
		rnats[i] = a
	}
	d.Set("rnat", rnats)
	return nil
}

func updateRnatClearFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateRnatClearFunc")

	if d.HasChange("rnat") {
		orig, noo := d.GetChange("rnat")
		if orig == nil {
			orig = new(schema.Set)
		}
		if noo == nil {
			noo = new(schema.Set)
		}
		oset := orig.(*schema.Set)
		nset := noo.(*schema.Set)

		remove := oset.Difference(nset).List()
		add := nset.Difference(oset).List()
		log.Printf("[DEBUG]  netscaler-provider: need to remove %d rnat", len(remove))
		log.Printf("[DEBUG]  netscaler-provider: need to add %d rnat", len(add))

		for _, val := range remove {
			rnat := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to delete rnat %v", rnat)
			err := deleteSingleRnat(rnat, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error deleting rnat %v", rnat)
			}
		}

		for _, val := range add {
			rnat := val.(map[string]interface{})
			log.Printf("[DEBUG]  netscaler-provider: going to add rnat %s", rnat["rnatsname"].(string))
			err := createSingleRnat(rnat, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error adding rnat %s", rnat["rnatsname"].(string))
			}
		}
	}

	return readRnatClearFunc(d, meta)
}

func deleteRnatClearFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteRnatClearFunc")
	rnats := d.Get("rnat").(*schema.Set).List()

	log.Printf("[DEBUG]  netscaler-provider: deleteRnatClearFunc: found %d rnat rules to delete", len(rnats))
	for _, val := range rnats {
		rnat := val.(map[string]interface{})
		_ = deleteSingleRnat(rnat, meta)
	}
	d.SetId("")
	return nil
}

func createSingleRnat(rnat map[string]interface{}, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	rnat2 := network.Rnat{}
	mapstructure.Decode(rnat, &rnat2)

	err := client.UpdateUnnamedResource(service.Rnat.Type(), &rnat2)
	if err != nil {
		return err
	}
	return nil
}

func deleteSingleRnat(rnat map[string]interface{}, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteSingleRnat")

	rnat2 := network.Rnat{}
	mapstructure.Decode(rnat, &rnat2)
	client := meta.(*NetScalerNitroClient).client
	err := client.ActOnResource(service.Rnat.Type(), rnat2, "clear")
	if err != nil {
		return err
	}

	return nil
}
