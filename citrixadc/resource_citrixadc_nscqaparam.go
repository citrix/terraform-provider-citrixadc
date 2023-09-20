package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNscqaparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNscqaparamFunc,
		Read:          readNscqaparamFunc,
		Update:        updateNscqaparamFunc,
		Delete:        deleteNscqaparamFunc,
		Schema: map[string]*schema.Schema{
			"net1label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net2label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net3label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"harqretxdelay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lr1coeflist": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lr1probthresh": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"lr2coeflist": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lr2probthresh": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"minrttnet1": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minrttnet2": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minrttnet3": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"net1cclscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net1csqscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net1logcoef": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net2cclscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net2csqscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net2logcoef": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net3cclscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net3csqscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net3logcoef": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNscqaparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNscqaparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nscqaparamName := resource.PrefixedUniqueId("tf-nscqaparam-")

	nscqaparam := ns.Nscqaparam{
		Harqretxdelay: d.Get("harqretxdelay").(int),
		Lr1coeflist:   d.Get("lr1coeflist").(string),
		Lr1probthresh: d.Get("lr1probthresh").(float64),
		Lr2coeflist:   d.Get("lr2coeflist").(string),
		Lr2probthresh: d.Get("lr2probthresh").(float64),
		Minrttnet1:    d.Get("minrttnet1").(int),
		Minrttnet2:    d.Get("minrttnet2").(int),
		Minrttnet3:    d.Get("minrttnet3").(int),
		Net1cclscale:  d.Get("net1cclscale").(string),
		Net1csqscale:  d.Get("net1csqscale").(string),
		Net1label:     d.Get("net1label").(string),
		Net1logcoef:   d.Get("net1logcoef").(string),
		Net2cclscale:  d.Get("net2cclscale").(string),
		Net2csqscale:  d.Get("net2csqscale").(string),
		Net2label:     d.Get("net2label").(string),
		Net2logcoef:   d.Get("net2logcoef").(string),
		Net3cclscale:  d.Get("net3cclscale").(string),
		Net3csqscale:  d.Get("net3csqscale").(string),
		Net3label:     d.Get("net3label").(string),
		Net3logcoef:   d.Get("net3logcoef").(string),
	}

	err := client.UpdateUnnamedResource("nscqaparam", &nscqaparam)
	if err != nil {
		return err
	}

	d.SetId(nscqaparamName)

	err = readNscqaparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nscqaparam but we can't read it ?? ")
		return nil
	}
	return nil
}

func readNscqaparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNscqaparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nscqaparam state")
	data, err := client.FindResource("nscqaparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nscqaparam state")
		d.SetId("")
		return nil
	}
	d.Set("harqretxdelay", data["harqretxdelay"])
	d.Set("lr1coeflist", data["lr1coeflist"])
	d.Set("lr1probthresh", data["lr1probthresh"])
	d.Set("lr2coeflist", data["lr2coeflist"])
	d.Set("lr2probthresh", data["lr2probthresh"])
	d.Set("minrttnet1", data["minrttnet1"])
	d.Set("minrttnet2", data["minrttnet2"])
	d.Set("minrttnet3", data["minrttnet3"])
	d.Set("net1cclscale", data["net1cclscale"])
	d.Set("net1csqscale", data["net1csqscale"])
	d.Set("net1label", data["net1label"])
	d.Set("net1logcoef", data["net1logcoef"])
	d.Set("net2cclscale", data["net2cclscale"])
	d.Set("net2csqscale", data["net2csqscale"])
	d.Set("net2label", data["net2label"])
	d.Set("net2logcoef", data["net2logcoef"])
	d.Set("net3cclscale", data["net3cclscale"])
	d.Set("net3csqscale", data["net3csqscale"])
	d.Set("net3label", data["net3label"])
	d.Set("net3logcoef", data["net3logcoef"])

	return nil

}

func updateNscqaparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNscqaparamFunc")
	client := meta.(*NetScalerNitroClient).client

	nscqaparam := ns.Nscqaparam{}
	hasChange := false
	if d.HasChange("harqretxdelay") {
		log.Printf("[DEBUG]  citrixadc-provider: Harqretxdelay has changed for nscqaparam, starting update")
		nscqaparam.Harqretxdelay = d.Get("harqretxdelay").(int)
		hasChange = true
	}
	if d.HasChange("lr1coeflist") {
		log.Printf("[DEBUG]  citrixadc-provider: Lr1coeflist has changed for nscqaparam, starting update")
		nscqaparam.Lr1coeflist = d.Get("lr1coeflist").(string)
		hasChange = true
	}
	if d.HasChange("lr1probthresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Lr1probthresh has changed for nscqaparam, starting update")
		nscqaparam.Lr1probthresh = d.Get("lr1probthresh").(float64)
		hasChange = true
	}
	if d.HasChange("lr2coeflist") {
		log.Printf("[DEBUG]  citrixadc-provider: Lr2coeflist has changed for nscqaparam, starting update")
		nscqaparam.Lr2coeflist = d.Get("lr2coeflist").(string)
		hasChange = true
	}
	if d.HasChange("lr2probthresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Lr2probthresh has changed for nscqaparam, starting update")
		nscqaparam.Lr2probthresh = d.Get("lr2probthresh").(float64)
		hasChange = true
	}
	if d.HasChange("minrttnet1") {
		log.Printf("[DEBUG]  citrixadc-provider: Minrttnet1 has changed for nscqaparam, starting update")
		nscqaparam.Minrttnet1 = d.Get("minrttnet1").(int)
		hasChange = true
	}
	if d.HasChange("minrttnet2") {
		log.Printf("[DEBUG]  citrixadc-provider: Minrttnet2 has changed for nscqaparam, starting update")
		nscqaparam.Minrttnet2 = d.Get("minrttnet2").(int)
		hasChange = true
	}
	if d.HasChange("minrttnet3") {
		log.Printf("[DEBUG]  citrixadc-provider: Minrttnet3 has changed for nscqaparam, starting update")
		nscqaparam.Minrttnet3 = d.Get("minrttnet3").(int)
		hasChange = true
	}
	if d.HasChange("net1cclscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Net1cclscale has changed for nscqaparam, starting update")
		nscqaparam.Net1cclscale = d.Get("net1cclscale").(string)
		hasChange = true
	}
	if d.HasChange("net1csqscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Net1csqscale has changed for nscqaparam, starting update")
		nscqaparam.Net1csqscale = d.Get("net1csqscale").(string)
		hasChange = true
	}
	if d.HasChange("net1label") {
		log.Printf("[DEBUG]  citrixadc-provider: Net1label has changed for nscqaparam, starting update")
		nscqaparam.Net1label = d.Get("net1label").(string)
		hasChange = true
	}
	if d.HasChange("net1logcoef") {
		log.Printf("[DEBUG]  citrixadc-provider: Net1logcoef has changed for nscqaparam, starting update")
		nscqaparam.Net1logcoef = d.Get("net1logcoef").(string)
		hasChange = true
	}
	if d.HasChange("net2cclscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Net2cclscale has changed for nscqaparam, starting update")
		nscqaparam.Net2cclscale = d.Get("net2cclscale").(string)
		hasChange = true
	}
	if d.HasChange("net2csqscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Net2csqscale has changed for nscqaparam, starting update")
		nscqaparam.Net2csqscale = d.Get("net2csqscale").(string)
		hasChange = true
	}
	if d.HasChange("net2label") {
		log.Printf("[DEBUG]  citrixadc-provider: Net2label has changed for nscqaparam, starting update")
		nscqaparam.Net2label = d.Get("net2label").(string)
		hasChange = true
	}
	if d.HasChange("net2logcoef") {
		log.Printf("[DEBUG]  citrixadc-provider: Net2logcoef has changed for nscqaparam, starting update")
		nscqaparam.Net2logcoef = d.Get("net2logcoef").(string)
		hasChange = true
	}
	if d.HasChange("net3cclscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Net3cclscale has changed for nscqaparam, starting update")
		nscqaparam.Net3cclscale = d.Get("net3cclscale").(string)
		hasChange = true
	}
	if d.HasChange("net3csqscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Net3csqscale has changed for nscqaparam, starting update")
		nscqaparam.Net3csqscale = d.Get("net3csqscale").(string)
		hasChange = true
	}
	if d.HasChange("net3label") {
		log.Printf("[DEBUG]  citrixadc-provider: Net3label has changed for nscqaparam, starting update")
		nscqaparam.Net3label = d.Get("net3label").(string)
		hasChange = true
	}
	if d.HasChange("net3logcoef") {
		log.Printf("[DEBUG]  citrixadc-provider: Net3logcoef has changed for nscqaparam, starting update")
		nscqaparam.Net3logcoef = d.Get("net3logcoef").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("nscqaparam", &nscqaparam)
		if err != nil {
			return fmt.Errorf("Error updating nscqaparam")
		}
	}
	return readNscqaparamFunc(d, meta)
}

func deleteNscqaparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNscqaparamFunc")
	// nscqparam does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
