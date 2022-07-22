package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcIp6tunnelparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIp6tunnelparamFunc,
		Read:          readIp6tunnelparamFunc,
		Update:        updateIp6tunnelparamFunc,
		Delete:        deleteIp6tunnelparamFunc,
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
			"useclientsourceipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIp6tunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIp6tunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var ip6tunnelparamName string
	// there is no primary key in ip6tunnelparam resource. Hence generate one for terraform state maintenance
	ip6tunnelparamName = resource.PrefixedUniqueId("tf-ip6tunnelparam-")
	ip6tunnelparam := network.Ip6tunnelparam{
		Dropfrag:             d.Get("dropfrag").(string),
		Dropfragcputhreshold: d.Get("dropfragcputhreshold").(int),
		Srcip:                d.Get("srcip").(string),
		Srciproundrobin:      d.Get("srciproundrobin").(string),
		Useclientsourceipv6:  d.Get("useclientsourceipv6").(string),
	}

	err := client.UpdateUnnamedResource(service.Ip6tunnelparam.Type(), &ip6tunnelparam)
	if err != nil {
		return err
	}

	d.SetId(ip6tunnelparamName)

	err = readIp6tunnelparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ip6tunnelparam but we can't read it ??")
		return nil
	}
	return nil
}

func readIp6tunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIp6tunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ip6tunnelparam state")
	data, err := client.FindResource(service.Ip6tunnelparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ip6tunnelparam state")
		d.SetId("")
		return nil
	}
	d.Set("dropfrag", data["dropfrag"])
	val,_ := strconv.Atoi(data["dropfragcputhreshold"].(string))
	d.Set("dropfragcputhreshold", val)
	d.Set("srcip", data["srcip"])
	d.Set("srciproundrobin", data["srciproundrobin"])
	d.Set("useclientsourceipv6", data["useclientsourceipv6"])

	return nil

}

func updateIp6tunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIp6tunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client

	ip6tunnelparam := network.Ip6tunnelparam{}
	hasChange := false
	if d.HasChange("dropfrag") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropfrag has changed for ip6tunnelparam, starting update")
		ip6tunnelparam.Dropfrag = d.Get("dropfrag").(string)
		hasChange = true
	}
	if d.HasChange("dropfragcputhreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropfragcputhreshold has changed for ip6tunnelparam, starting update")
		ip6tunnelparam.Dropfragcputhreshold = d.Get("dropfragcputhreshold").(int)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for ip6tunnelparam, starting update")
		ip6tunnelparam.Srcip = d.Get("srcip").(string)
		hasChange = true
	}
	if d.HasChange("srciproundrobin") {
		log.Printf("[DEBUG]  citrixadc-provider: Srciproundrobin has changed for ip6tunnelparam, starting update")
		ip6tunnelparam.Srciproundrobin = d.Get("srciproundrobin").(string)
		hasChange = true
	}
	if d.HasChange("useclientsourceipv6") {
		log.Printf("[DEBUG]  citrixadc-provider: Useclientsourceipv6 has changed for ip6tunnelparam, starting update")
		ip6tunnelparam.Useclientsourceipv6 = d.Get("useclientsourceipv6").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ip6tunnelparam.Type(), &ip6tunnelparam)
		if err != nil {
			return fmt.Errorf("Error updating ip6tunnelparam")
		}
	}
	return readIp6tunnelparamFunc(d, meta)
}

func deleteIp6tunnelparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIp6tunnelparamFunc")
	client := meta.(*NetScalerNitroClient).client
	type ip6tunnelparamRemove struct {
		Srcip 					bool `json:"srcip,omitempty"`
		Dropfrag  				bool `json:"dropfrag,omitempty"`
		Dropfragcputhreshold    bool `json:"dropfragcputhreshold,omitempty"`
		Srciproundrobin      	bool `json:"srciproundrobin,omitempty"`
		Useclientsourceipv6     bool `json:"useclientsourceipv6,omitempty"`
	}
	ip6tunnelparam := ip6tunnelparamRemove{}
	log.Printf("ip6tunnelparamdelete struct %v", ip6tunnelparam)

	if _, ok := d.GetOk("srcip"); ok {
		ip6tunnelparam.Srcip = true
	}

	if _, ok := d.GetOk("dropfrag"); ok {
		ip6tunnelparam.Dropfrag = true
	}

	if _, ok := d.GetOk("dropfragcputhreshold"); ok {
		ip6tunnelparam.Dropfragcputhreshold = true
	}
	if _, ok := d.GetOk("srciproundrobin"); ok {
		ip6tunnelparam.Srciproundrobin = true
	}
	if _, ok := d.GetOk("useclientsourceipv6"); ok {
		ip6tunnelparam.Useclientsourceipv6 = true
	}

	err := client.ActOnResource(service.Ip6tunnelparam.Type(), &ip6tunnelparam, "unset")
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
