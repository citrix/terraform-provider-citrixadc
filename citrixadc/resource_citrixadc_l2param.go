package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcL2param() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createL2paramFunc,
		Read:          readL2paramFunc,
		Update:        updateL2paramFunc,
		Delete:        deleteL2paramFunc,
		Schema: map[string]*schema.Schema{
			"bdggrpproxyarp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bdgsetting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bridgeagetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"garponvridintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"garpreply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macmodefwdmypkt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxbridgecollision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mbfinstlearning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbfpeermacupdate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxyarp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"returntoethernetsender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rstintfonhafo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skipproxyingbsdtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stopmacmoveupdate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usemymac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usenetprofilebsdtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createL2paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createL2paramFunc")
	client := meta.(*NetScalerNitroClient).client
	l2paramName := resource.PrefixedUniqueId("tf-l2param-")

	l2param := network.L2param{
		Bdggrpproxyarp:          d.Get("bdggrpproxyarp").(string),
		Bdgsetting:              d.Get("bdgsetting").(string),
		Bridgeagetimeout:        d.Get("bridgeagetimeout").(int),
		Garponvridintf:          d.Get("garponvridintf").(string),
		Garpreply:               d.Get("garpreply").(string),
		Macmodefwdmypkt:         d.Get("macmodefwdmypkt").(string),
		Maxbridgecollision:      d.Get("maxbridgecollision").(int),
		Mbfinstlearning:         d.Get("mbfinstlearning").(string),
		Mbfpeermacupdate:        d.Get("mbfpeermacupdate").(int),
		Proxyarp:                d.Get("proxyarp").(string),
		Returntoethernetsender:  d.Get("returntoethernetsender").(string),
		Rstintfonhafo:           d.Get("rstintfonhafo").(string),
		Skipproxyingbsdtraffic:  d.Get("skipproxyingbsdtraffic").(string),
		Stopmacmoveupdate:       d.Get("stopmacmoveupdate").(string),
		Usemymac:                d.Get("usemymac").(string),
		Usenetprofilebsdtraffic: d.Get("usenetprofilebsdtraffic").(string),
	}

	err := client.UpdateUnnamedResource(service.L2param.Type(), &l2param)
	if err != nil {
		return err
	}

	d.SetId(l2paramName)

	err = readL2paramFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this l2param but we can't read it ??")
		return nil
	}
	return nil
}

func readL2paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readL2paramFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading l2param state")
	data, err := client.FindResource(service.L2param.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing l2param state")
		d.SetId("")
		return nil
	}
	d.Set("bdggrpproxyarp", data["bdggrpproxyarp"])
	d.Set("bdgsetting", data["bdgsetting"])
	d.Set("bridgeagetimeout", data["bridgeagetimeout"])
	d.Set("garponvridintf", data["garponvridintf"])
	d.Set("garpreply", data["garpreply"])
	d.Set("macmodefwdmypkt", data["macmodefwdmypkt"])
	d.Set("maxbridgecollision", data["maxbridgecollision"])
	d.Set("mbfinstlearning", data["mbfinstlearning"])
	d.Set("mbfpeermacupdate", data["mbfpeermacupdate"])
	d.Set("proxyarp", data["proxyarp"])
	d.Set("returntoethernetsender", data["returntoethernetsender"])
	d.Set("rstintfonhafo", data["rstintfonhafo"])
	d.Set("skipproxyingbsdtraffic", data["skipproxyingbsdtraffic"])
	d.Set("stopmacmoveupdate", data["stopmacmoveupdate"])
	d.Set("usemymac", data["usemymac"])
	d.Set("usenetprofilebsdtraffic", data["usenetprofilebsdtraffic"])

	return nil

}

func updateL2paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateL2paramFunc")
	client := meta.(*NetScalerNitroClient).client

	l2param := network.L2param{}
	hasChange := false
	if d.HasChange("bdggrpproxyarp") {
		log.Printf("[DEBUG]  citrixadc-provider: Bdggrpproxyarp has changed for l2param, starting update")
		l2param.Bdggrpproxyarp = d.Get("bdggrpproxyarp").(string)
		hasChange = true
	}
	if d.HasChange("bdgsetting") {
		log.Printf("[DEBUG]  citrixadc-provider: Bdgsetting has changed for l2param, starting update")
		l2param.Bdgsetting = d.Get("bdgsetting").(string)
		hasChange = true
	}
	if d.HasChange("bridgeagetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Bridgeagetimeout has changed for l2param, starting update")
		l2param.Bridgeagetimeout = d.Get("bridgeagetimeout").(int)
		hasChange = true
	}
	if d.HasChange("garponvridintf") {
		log.Printf("[DEBUG]  citrixadc-provider: Garponvridintf has changed for l2param, starting update")
		l2param.Garponvridintf = d.Get("garponvridintf").(string)
		hasChange = true
	}
	if d.HasChange("garpreply") {
		log.Printf("[DEBUG]  citrixadc-provider: Garpreply has changed for l2param, starting update")
		l2param.Garpreply = d.Get("garpreply").(string)
		hasChange = true
	}
	if d.HasChange("macmodefwdmypkt") {
		log.Printf("[DEBUG]  citrixadc-provider: Macmodefwdmypkt has changed for l2param, starting update")
		l2param.Macmodefwdmypkt = d.Get("macmodefwdmypkt").(string)
		hasChange = true
	}
	if d.HasChange("maxbridgecollision") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbridgecollision has changed for l2param, starting update")
		l2param.Maxbridgecollision = d.Get("maxbridgecollision").(int)
		hasChange = true
	}
	if d.HasChange("mbfinstlearning") {
		log.Printf("[DEBUG]  citrixadc-provider: Mbfinstlearning has changed for l2param, starting update")
		l2param.Mbfinstlearning = d.Get("mbfinstlearning").(string)
		hasChange = true
	}
	if d.HasChange("mbfpeermacupdate") {
		log.Printf("[DEBUG]  citrixadc-provider: Mbfpeermacupdate has changed for l2param, starting update")
		l2param.Mbfpeermacupdate = d.Get("mbfpeermacupdate").(int)
		hasChange = true
	}
	if d.HasChange("proxyarp") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyarp has changed for l2param, starting update")
		l2param.Proxyarp = d.Get("proxyarp").(string)
		hasChange = true
	}
	if d.HasChange("returntoethernetsender") {
		log.Printf("[DEBUG]  citrixadc-provider: Returntoethernetsender has changed for l2param, starting update")
		l2param.Returntoethernetsender = d.Get("returntoethernetsender").(string)
		hasChange = true
	}
	if d.HasChange("rstintfonhafo") {
		log.Printf("[DEBUG]  citrixadc-provider: Rstintfonhafo has changed for l2param, starting update")
		l2param.Rstintfonhafo = d.Get("rstintfonhafo").(string)
		hasChange = true
	}
	if d.HasChange("skipproxyingbsdtraffic") {
		log.Printf("[DEBUG]  citrixadc-provider: Skipproxyingbsdtraffic has changed for l2param, starting update")
		l2param.Skipproxyingbsdtraffic = d.Get("skipproxyingbsdtraffic").(string)
		hasChange = true
	}
	if d.HasChange("stopmacmoveupdate") {
		log.Printf("[DEBUG]  citrixadc-provider: Stopmacmoveupdate has changed for l2param, starting update")
		l2param.Stopmacmoveupdate = d.Get("stopmacmoveupdate").(string)
		hasChange = true
	}
	if d.HasChange("usemymac") {
		log.Printf("[DEBUG]  citrixadc-provider: Usemymac has changed for l2param, starting update")
		l2param.Usemymac = d.Get("usemymac").(string)
		hasChange = true
	}
	if d.HasChange("usenetprofilebsdtraffic") {
		log.Printf("[DEBUG]  citrixadc-provider: Usenetprofilebsdtraffic has changed for l2param, starting update")
		l2param.Usenetprofilebsdtraffic = d.Get("usenetprofilebsdtraffic").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.L2param.Type(), &l2param)
		if err != nil {
			return fmt.Errorf("Error updating l2param")
		}
	}
	return readL2paramFunc(d, meta)
}

func deleteL2paramFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteL2paramFunc")
	// l2param does not suppor DELETE operation
	d.SetId("")

	return nil
}
