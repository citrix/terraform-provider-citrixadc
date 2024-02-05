package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

// We do not use the go-nitro struct because we need the
// Ownernode to be string so that it correctly behaves for 0 values
// The problem is that if it is int there is no way to avoid sending ownernode param to VPX
// which will fail if the NSIP is that of a standalone VPX.
type Nsvpxparam struct {
	Cpuyield        string `json:"cpuyield,omitempty"`
	Masterclockcpu1 string `json:"masterclockcpu1,omitempty"`
	Ownernode       string `json:"ownernode,omitempty"`
}

func resourceCitrixAdcNsvpxparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsvpxparamFunc,
		Read:          readNsvpxparamFunc,
		Delete:        deleteNsvpxparamFunc,
		Schema: map[string]*schema.Schema{
			"cpuyield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"masterclockcpu1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownernode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNsvpxparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsvpxparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvpxparamName := resource.PrefixedUniqueId("tf-nsvpxparam-")

	nsvpxparam := Nsvpxparam{
		Cpuyield:        d.Get("cpuyield").(string),
		Masterclockcpu1: d.Get("masterclockcpu1").(string),
		Ownernode:       d.Get("ownernode").(string),
	}

	err := client.UpdateUnnamedResource("nsvpxparam", &nsvpxparam)
	if err != nil {
		return err
	}

	d.SetId(nsvpxparamName)

	err = readNsvpxparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsvpxparam but we can't read it ?? %s", nsvpxparamName)
		return nil
	}
	return nil
}

func readNsvpxparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsvpxparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nsvpxparamName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsvpxparam state %s", nsvpxparamName)
	findParams := service.FindParams{
		ResourceType: "nsvpxparam",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}

	ownernode, ownernodeOk := d.GetOk("ownernode")

	foundIndex := -1
	if ownernodeOk {
		for index, value := range dataArr {
			if ownernode == value["ownernode"] {
				foundIndex = index
			}
		}
	} else {
		// In standalone VPX there is only one entry for nsvpxparam
		foundIndex = 0
	}

	if foundIndex == -1 {
		// Clear state for resource
		d.SetId("")
		return nil
	}

	data := dataArr[foundIndex]

	d.Set("cpuyield", data["cpuyield"])
	// d.Set("masterclockcpu1", data["masterclockcpu1"])
	d.Set("ownernode", data["ownernode"])

	return nil

}

func deleteNsvpxparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsvpxparamFunc")
	// Just delete the reference
	// Actual configuration cannot be deleted
	d.SetId("")

	return nil
}
