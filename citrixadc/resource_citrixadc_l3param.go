package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcL3param() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createL3paramFunc,
		Read:          readL3paramFunc,
		Update:        updateL3paramFunc,
		Delete:        deleteL3paramFunc,
		Schema: map[string]*schema.Schema{
			"acllogtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowclasseipv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropdfflag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropipfragments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamicrouting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"externalloopback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forwardicmpfragments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpgenratethreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"implicitaclallow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6dynamicrouting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"miproundrobin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overridernat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcnat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tnlpmtuwoconn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usipserverstraypkt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createL3paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createL3paramFunc")
	client := meta.(*NetScalerNitroClient).client
	l3paramName := resource.PrefixedUniqueId("tf-l3param-")

	l3param := network.L3param{
		Acllogtime:           d.Get("acllogtime").(int),
		Allowclasseipv4:      d.Get("allowclasseipv4").(string),
		Dropdfflag:           d.Get("dropdfflag").(string),
		Dropipfragments:      d.Get("dropipfragments").(string),
		Dynamicrouting:       d.Get("dynamicrouting").(string),
		Externalloopback:     d.Get("externalloopback").(string),
		Forwardicmpfragments: d.Get("forwardicmpfragments").(string),
		Icmpgenratethreshold: d.Get("icmpgenratethreshold").(int),
		Implicitaclallow:     d.Get("implicitaclallow").(string),
		Ipv6dynamicrouting:   d.Get("ipv6dynamicrouting").(string),
		Miproundrobin:        d.Get("miproundrobin").(string),
		Overridernat:         d.Get("overridernat").(string),
		Srcnat:               d.Get("srcnat").(string),
		Tnlpmtuwoconn:        d.Get("tnlpmtuwoconn").(string),
		Usipserverstraypkt:   d.Get("usipserverstraypkt").(string),
	}

	err := client.UpdateUnnamedResource(service.L3param.Type(), &l3param)
	if err != nil {
		return err
	}

	d.SetId(l3paramName)

	err = readL3paramFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this l3param but we can't read it ??")
		return nil
	}
	return nil
}

func readL3paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readL3paramFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading l3param state")
	data, err := client.FindResource(service.L3param.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing l3param state")
		d.SetId("")
		return nil
	}
	d.Set("acllogtime", data["acllogtime"])
	d.Set("allowclasseipv4", data["allowclasseipv4"])
	d.Set("dropdfflag", data["dropdfflag"])
	d.Set("dropipfragments", data["dropipfragments"])
	d.Set("dynamicrouting", data["dynamicrouting"])
	d.Set("externalloopback", data["externalloopback"])
	d.Set("forwardicmpfragments", data["forwardicmpfragments"])
	d.Set("icmpgenratethreshold", data["icmpgenratethreshold"])
	d.Set("implicitaclallow", data["implicitaclallow"])
	d.Set("ipv6dynamicrouting", data["ipv6dynamicrouting"])
	d.Set("miproundrobin", data["miproundrobin"])
	d.Set("overridernat", data["overridernat"])
	d.Set("srcnat", data["srcnat"])
	d.Set("tnlpmtuwoconn", data["tnlpmtuwoconn"])
	d.Set("usipserverstraypkt", data["usipserverstraypkt"])

	return nil

}

func updateL3paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateL3paramFunc")
	client := meta.(*NetScalerNitroClient).client

	l3param := network.L3param{}
	hasChange := false
	if d.HasChange("acllogtime") {
		log.Printf("[DEBUG]  citrixadc-provider: Acllogtime has changed for l3param, starting update")
		l3param.Acllogtime = d.Get("acllogtime").(int)
		hasChange = true
	}
	if d.HasChange("allowclasseipv4") {
		log.Printf("[DEBUG]  citrixadc-provider: Allowclasseipv4 has changed for l3param, starting update")
		l3param.Allowclasseipv4 = d.Get("allowclasseipv4").(string)
		hasChange = true
	}
	if d.HasChange("dropdfflag") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropdfflag has changed for l3param, starting update")
		l3param.Dropdfflag = d.Get("dropdfflag").(string)
		hasChange = true
	}
	if d.HasChange("dropipfragments") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropipfragments has changed for l3param, starting update")
		l3param.Dropipfragments = d.Get("dropipfragments").(string)
		hasChange = true
	}
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for l3param, starting update")
		l3param.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("externalloopback") {
		log.Printf("[DEBUG]  citrixadc-provider: Externalloopback has changed for l3param, starting update")
		l3param.Externalloopback = d.Get("externalloopback").(string)
		hasChange = true
	}
	if d.HasChange("forwardicmpfragments") {
		log.Printf("[DEBUG]  citrixadc-provider: Forwardicmpfragments has changed for l3param, starting update")
		l3param.Forwardicmpfragments = d.Get("forwardicmpfragments").(string)
		hasChange = true
	}
	if d.HasChange("icmpgenratethreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmpgenratethreshold has changed for l3param, starting update")
		l3param.Icmpgenratethreshold = d.Get("icmpgenratethreshold").(int)
		hasChange = true
	}
	if d.HasChange("implicitaclallow") {
		log.Printf("[DEBUG]  citrixadc-provider: Implicitaclallow has changed for l3param, starting update")
		l3param.Implicitaclallow = d.Get("implicitaclallow").(string)
		hasChange = true
	}
	if d.HasChange("ipv6dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6dynamicrouting has changed for l3param, starting update")
		l3param.Ipv6dynamicrouting = d.Get("ipv6dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("miproundrobin") {
		log.Printf("[DEBUG]  citrixadc-provider: Miproundrobin has changed for l3param, starting update")
		l3param.Miproundrobin = d.Get("miproundrobin").(string)
		hasChange = true
	}
	if d.HasChange("overridernat") {
		log.Printf("[DEBUG]  citrixadc-provider: Overridernat has changed for l3param, starting update")
		l3param.Overridernat = d.Get("overridernat").(string)
		hasChange = true
	}
	if d.HasChange("srcnat") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcnat has changed for l3param, starting update")
		l3param.Srcnat = d.Get("srcnat").(string)
		hasChange = true
	}
	if d.HasChange("tnlpmtuwoconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Tnlpmtuwoconn has changed for l3param, starting update")
		l3param.Tnlpmtuwoconn = d.Get("tnlpmtuwoconn").(string)
		hasChange = true
	}
	if d.HasChange("usipserverstraypkt") {
		log.Printf("[DEBUG]  citrixadc-provider: Usipserverstraypkt has changed for l3param, starting update")
		l3param.Usipserverstraypkt = d.Get("usipserverstraypkt").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.L3param.Type(), &l3param)
		if err != nil {
			return fmt.Errorf("Error updating l3param")
		}
	}
	return readL3paramFunc(d, meta)
}

func deleteL3paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteL3paramFunc")
	// l3param does not support delete operation
	d.SetId("")

	return nil
}
