package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIptunnelparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIptunnelparamFunc,
		Read:          readIptunnelparamFunc,
		Update:        updateIptunnelparamFunc,
		Delete:        deleteIptunnelparamFunc,
		Schema: map[string]*schema.Schema{
			"dropfrag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropfragcputhreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enablestrictrx": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enablestricttx": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srciproundrobin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useclientsourceip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIptunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIptunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var iptunnelparamName string
	// there is no primary key in iptunnelparam resource. Hence generate one for terraform state maintenance
	iptunnelparamName = resource.PrefixedUniqueId("tf-iptunnelparam-")
	iptunnelparam := network.Iptunnelparam{
		Dropfrag:             d.Get("dropfrag").(string),
		Dropfragcputhreshold: d.Get("dropfragcputhreshold").(int),
		Enablestrictrx:       d.Get("enablestrictrx").(string),
		Enablestricttx:       d.Get("enablestricttx").(string),
		Mac:                  d.Get("mac").(string),
		Srcip:                d.Get("srcip").(string),
		Srciproundrobin:      d.Get("srciproundrobin").(string),
		Useclientsourceip:    d.Get("useclientsourceip").(string),
	}

	err := client.UpdateUnnamedResource(service.Iptunnelparam.Type(), &iptunnelparam)
	if err != nil {
		return err
	}

	d.SetId(iptunnelparamName)

	err = readIptunnelparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this iptunnelparam but we can't read it ?? %s", iptunnelparamName)
		return nil
	}
	return nil
}

func readIptunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIptunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading iptunnelparam state")
	data, err := client.FindResource(service.Iptunnelparam.Type(),"")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing iptunnelparam state")
		d.SetId("")
		return nil
	}
	d.Set("dropfrag", data["dropfrag"])
	d.Set("dropfragcputhreshold", data["dropfragcputhreshold"])
	d.Set("enablestrictrx", data["enablestrictrx"])
	d.Set("enablestricttx", data["enablestricttx"])
	d.Set("mac", data["mac"])
	d.Set("srcip", data["srcip"])
	d.Set("srciproundrobin", data["srciproundrobin"])
	d.Set("useclientsourceip", data["useclientsourceip"])

	return nil

}

func updateIptunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIptunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client

	iptunnelparam := network.Iptunnelparam{}
	hasChange := false
	if d.HasChange("dropfrag") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropfrag has changed for iptunnelparam, starting update")
		iptunnelparam.Dropfrag = d.Get("dropfrag").(string)
		hasChange = true
	}
	if d.HasChange("dropfragcputhreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropfragcputhreshold has changed for iptunnelparam, starting update")
		iptunnelparam.Dropfragcputhreshold = d.Get("dropfragcputhreshold").(int)
		hasChange = true
	}
	if d.HasChange("enablestrictrx") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablestrictrx has changed for iptunnelparam, starting update")
		iptunnelparam.Enablestrictrx = d.Get("enablestrictrx").(string)
		hasChange = true
	}
	if d.HasChange("enablestricttx") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablestricttx has changed for iptunnelparam, starting update")
		iptunnelparam.Enablestricttx = d.Get("enablestricttx").(string)
		hasChange = true
	}
	if d.HasChange("mac") {
		log.Printf("[DEBUG]  citrixadc-provider: Mac has changed for iptunnelparam, starting update")
		iptunnelparam.Mac = d.Get("mac").(string)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for iptunnelparam, starting update")
		iptunnelparam.Srcip = d.Get("srcip").(string)
		hasChange = true
	}
	if d.HasChange("srciproundrobin") {
		log.Printf("[DEBUG]  citrixadc-provider: Srciproundrobin has changed for iptunnelparam, starting update")
		iptunnelparam.Srciproundrobin = d.Get("srciproundrobin").(string)
		hasChange = true
	}
	if d.HasChange("useclientsourceip") {
		log.Printf("[DEBUG]  citrixadc-provider: Useclientsourceip has changed for iptunnelparam, starting update")
		iptunnelparam.Useclientsourceip = d.Get("useclientsourceip").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Iptunnelparam.Type(), &iptunnelparam)
		if err != nil {
			return fmt.Errorf("Error updating iptunnelparam")
		}
	}
	return readIptunnelparamFunc(d, meta)
}

func deleteIptunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIptunnelparamFunc")


	d.SetId("")

	return nil
}
