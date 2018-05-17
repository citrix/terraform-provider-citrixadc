package netscaler

import (
	"github.com/chiradeep/go-nitro/config/network"
	"github.com/mitchellh/mapstructure"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceNetScalerRnats() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRnatsFunc,
		Read:          readRnatsFunc,
		Update:        updateRnatsFunc,
		Delete:        deleteRnatsFunc,
		Schema: map[string]*schema.Schema{
			"rnatsname": &schema.Schema{
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
						"aclname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"natip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"natip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"netmask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirectport": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"td": &schema.Schema{
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

func createRnatsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createRnatsFunc")

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

func readRnatsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readRnatsFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading rnat state %s", rnatName)

	data, _ := client.FindAllResources(netscaler.Rnat.Type())
	rnats := make([]map[string]interface{}, len(data))
	for i, a := range data {
		rnats[i] = a
	}
	d.Set("rnat", rnats)
	return nil
}

func updateRnatsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateRnatsFunc")

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
			log.Printf("[DEBUG]  netscaler-provider: going to add rnat %s", rnat["rnatname"].(string))
			err := createSingleRnat(rnat, meta)
			if err != nil {
				log.Printf("[DEBUG]  netscaler-provider: error adding rnat %s", rnat["rnatname"].(string))
			}
		}
	}

	return readRnatsFunc(d, meta)
}

func deleteRnatsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteRnatsFunc")
	rnats := d.Get("rnat").(*schema.Set).List()

	log.Printf("[DEBUG]  netscaler-provider: deleteRnatsFunc: found %d rnat rules to delete", len(rnats))
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

	err := client.UpdateUnnamedResource(netscaler.Rnat.Type(), &rnat2)
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
	err := client.ActOnResource(netscaler.Rnat.Type(), rnat2, "clear")
	if err != nil {
		return err
	}

	return nil
}
