package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcIpv6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIpv6Func,
		Read:          readIpv6Func,
		Update:        updateIpv6Func,
		Delete:        deleteIpv6Func,
		Schema: map[string]*schema.Schema{
			"dodad": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndbasereachtime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ndretransmissiontime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ralearning": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routerredirection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"usipnatprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIpv6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpv6Func")
	client := meta.(*NetScalerNitroClient).client
	ipv6Name := resource.PrefixedUniqueId("tf-ipv6-")
	ipv6 := make(map[string]interface{})

	if v, ok := d.GetOk("dodad"); ok {
		ipv6["dodad"] = v.(string)
	}
	if v, ok := d.GetOk("natprefix"); ok {
		ipv6["natprefix"] = v.(string)
	}
	if v, ok := d.GetOk("ndbasereachtime"); ok {
		ipv6["ndbasereachtime"] = v.(int)
	}
	if v, ok := d.GetOk("ndretransmissiontime"); ok {
		ipv6["ndretransmissiontime"] = v.(int)
	}
	if v, ok := d.GetOk("ralearning"); ok {
		ipv6["ralearning"] = v.(string)

	}
	if v, ok := d.GetOk("routerredirection"); ok {
		ipv6["routerredirection"] = v.(string)
	}
	if v, ok := d.GetOk("td"); ok {
		ipv6["td"] = v.(int)
	}
	if v, ok := d.GetOk("usipnatprefix"); ok {
		ipv6["usipnatprefix"] = v.(string)
	}

	err := client.UpdateUnnamedResource(service.Ipv6.Type(), &ipv6)
	if err != nil {
		return err
	}

	d.SetId(ipv6Name)

	err = readIpv6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ipv6 but we can't read it ??")
		return nil
	}
	return nil
}

func readIpv6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpv6Func")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ipv6 state")
	data, err := client.FindResource(service.Ipv6.Type(), strconv.Itoa(d.Get("td").(int)))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipv6 state")
		d.SetId("")
		return nil
	}
	d.Set("td", data["td"])
	d.Set("dodad", data["dodad"])
	d.Set("natprefix", data["natprefix"])
	d.Set("ndbasereachtime", data["ndbasereachtime"])
	d.Set("ndretransmissiontime", data["ndretransmissiontime"])
	d.Set("ralearning", data["ralearning"])
	d.Set("routerredirection", data["routerredirection"])
	d.Set("td", data["td"])
	d.Set("usipnatprefix", data["usipnatprefix"])

	return nil

}

func updateIpv6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpv6Func")
	client := meta.(*NetScalerNitroClient).client

	ipv6 := make(map[string]interface{})

	hasChange := false
	if d.HasChange("dodad") {
		log.Printf("[DEBUG]  citrixadc-provider: Dodad has changed for ipv6, starting update")
		ipv6["dodad"] = d.Get("dodad").(string)
		hasChange = true
	}
	if d.HasChange("natprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Natprefix has changed for ipv6, starting update")
		ipv6["natprefix"] = d.Get("natprefix").(string)
		hasChange = true
	}
	if d.HasChange("ndbasereachtime") {
		log.Printf("[DEBUG]  citrixadc-provider: Ndbasereachtime has changed for ipv6, starting update")
		ipv6["ndbasereachtime"] = d.Get("ndbasereachtime").(int)
		hasChange = true
	}
	if d.HasChange("ndretransmissiontime") {
		log.Printf("[DEBUG]  citrixadc-provider: Ndretransmissiontime has changed for ipv6, starting update")
		ipv6["ndretransmissiontime"] = d.Get("ndretransmissiontime").(int)
		hasChange = true
	}
	if d.HasChange("ralearning") {
		log.Printf("[DEBUG]  citrixadc-provider: Ralearning has changed for ipv6, starting update")
		ipv6["ralearning"] = d.Get("ralearning").(string)
		hasChange = true
	}
	if d.HasChange("routerredirection") {
		log.Printf("[DEBUG]  citrixadc-provider: Routerredirection has changed for ipv6, starting update")
		ipv6["routerredirection"] = d.Get("routerredirection").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for ipv6, starting update")
		ipv6["td"] = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("usipnatprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Usipnatprefix has changed for ipv6, starting update")
		ipv6["usipnatprefix"] = d.Get("usipnatprefix").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ipv6.Type(), &ipv6)
		if err != nil {
			return fmt.Errorf("Error updating ipv6")
		}
	}
	return readIpv6Func(d, meta)
}

func deleteIpv6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpv6Func")
	// ipv6 does not support DELETE operation
	d.SetId("")

	return nil
}
