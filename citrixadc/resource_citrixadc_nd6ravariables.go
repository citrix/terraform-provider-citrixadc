package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcNd6ravariables() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNd6ravariablesFunc,
		Read:          readNd6ravariablesFunc,
		Update:        updateNd6ravariablesFunc,
		Delete:        deleteNd6ravariablesFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ceaserouteradv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"currhoplimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultlifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"linkmtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"managedaddrconfig": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxrtadvinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minrtadvinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"onlyunicastrtadvresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"otheraddrconfig": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reachabletime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retranstime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sendrouteradv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srclinklayeraddroption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNd6ravariablesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNd6ravariablesFunc")
	client := meta.(*NetScalerNitroClient).client
	nd6ravariablesName := strconv.Itoa(d.Get("vlan").(int))
	nd6ravariables := network.Nd6ravariables{
		Ceaserouteradv:           d.Get("ceaserouteradv").(string),
		Currhoplimit:             d.Get("currhoplimit").(int),
		Defaultlifetime:          d.Get("defaultlifetime").(int),
		Linkmtu:                  d.Get("linkmtu").(int),
		Managedaddrconfig:        d.Get("managedaddrconfig").(string),
		Maxrtadvinterval:         d.Get("maxrtadvinterval").(int),
		Minrtadvinterval:         d.Get("minrtadvinterval").(int),
		Onlyunicastrtadvresponse: d.Get("onlyunicastrtadvresponse").(string),
		Otheraddrconfig:          d.Get("otheraddrconfig").(string),
		Reachabletime:            d.Get("reachabletime").(int),
		Retranstime:              d.Get("retranstime").(int),
		Sendrouteradv:            d.Get("sendrouteradv").(string),
		Srclinklayeraddroption:   d.Get("srclinklayeraddroption").(string),
		Vlan:                     d.Get("vlan").(int),
	}

	err := client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
	if err != nil {
		return err
	}

	d.SetId(nd6ravariablesName)

	err = readNd6ravariablesFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nd6ravariables but we can't read it ?? %s", nd6ravariablesName)
		return nil
	}
	return nil
}

func readNd6ravariablesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNd6ravariablesFunc")
	client := meta.(*NetScalerNitroClient).client
	nd6ravariablesName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nd6ravariables state %s", nd6ravariablesName)
	data, err := client.FindResource(service.Nd6ravariables.Type(), nd6ravariablesName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nd6ravariables state %s", nd6ravariablesName)
		d.SetId("")
		return nil
	}
	vlan_int,_ := strconv.Atoi(data["vlan"].(string))
	d.Set("vlan", vlan_int)
	d.Set("ceaserouteradv", data["ceaserouteradv"])
	d.Set("currhoplimit", data["currhoplimit"])
	d.Set("defaultlifetime", data["defaultlifetime"])
	d.Set("linkmtu", data["linkmtu"])
	d.Set("managedaddrconfig", data["managedaddrconfig"])
	d.Set("maxrtadvinterval", data["maxrtadvinterval"])
	d.Set("minrtadvinterval", data["minrtadvinterval"])
	d.Set("onlyunicastrtadvresponse", data["onlyunicastrtadvresponse"])
	d.Set("otheraddrconfig", data["otheraddrconfig"])
	d.Set("reachabletime", data["reachabletime"])
	d.Set("retranstime", data["retranstime"])
	d.Set("sendrouteradv", data["sendrouteradv"])
	d.Set("srclinklayeraddroption", data["srclinklayeraddroption"])
	d.Set("vlan", data["vlan"])

	return nil

}

func updateNd6ravariablesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNd6ravariablesFunc")
	client := meta.(*NetScalerNitroClient).client
	nd6ravariablesName := strconv.Itoa(d.Get("vlan").(int))

	nd6ravariables := network.Nd6ravariables{
		Vlan: d.Get("vlan").(int),
	}
	hasChange := false
	if d.HasChange("ceaserouteradv") {
		log.Printf("[DEBUG]  citrixadc-provider: Ceaserouteradv has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Ceaserouteradv = d.Get("ceaserouteradv").(string)
		hasChange = true
	}
	if d.HasChange("currhoplimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Currhoplimit has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Currhoplimit = d.Get("currhoplimit").(int)
		hasChange = true
	}
	if d.HasChange("defaultlifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultlifetime has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Defaultlifetime = d.Get("defaultlifetime").(int)
		hasChange = true
	}
	if d.HasChange("linkmtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Linkmtu has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Linkmtu = d.Get("linkmtu").(int)
		hasChange = true
	}
	if d.HasChange("managedaddrconfig") {
		log.Printf("[DEBUG]  citrixadc-provider: Managedaddrconfig has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Managedaddrconfig = d.Get("managedaddrconfig").(string)
		hasChange = true
	}
	if d.HasChange("maxrtadvinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxrtadvinterval has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Maxrtadvinterval = d.Get("maxrtadvinterval").(int)
		hasChange = true
	}
	if d.HasChange("minrtadvinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Minrtadvinterval has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Minrtadvinterval = d.Get("minrtadvinterval").(int)
		hasChange = true
	}
	if d.HasChange("onlyunicastrtadvresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Onlyunicastrtadvresponse has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Onlyunicastrtadvresponse = d.Get("onlyunicastrtadvresponse").(string)
		hasChange = true
	}
	if d.HasChange("otheraddrconfig") {
		log.Printf("[DEBUG]  citrixadc-provider: Otheraddrconfig has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Otheraddrconfig = d.Get("otheraddrconfig").(string)
		hasChange = true
	}
	if d.HasChange("reachabletime") {
		log.Printf("[DEBUG]  citrixadc-provider: Reachabletime has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Reachabletime = d.Get("reachabletime").(int)
		hasChange = true
	}
	if d.HasChange("retranstime") {
		log.Printf("[DEBUG]  citrixadc-provider: Retranstime has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Retranstime = d.Get("retranstime").(int)
		hasChange = true
	}
	if d.HasChange("sendrouteradv") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendrouteradv has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Sendrouteradv = d.Get("sendrouteradv").(string)
		hasChange = true
	}
	if d.HasChange("srclinklayeraddroption") {
		log.Printf("[DEBUG]  citrixadc-provider: Srclinklayeraddroption has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Srclinklayeraddroption = d.Get("srclinklayeraddroption").(string)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nd6ravariables %s, starting update", nd6ravariablesName)
		nd6ravariables.Vlan = d.Get("vlan").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
		if err != nil {
			return fmt.Errorf("Error updating nd6ravariables %s", nd6ravariablesName)
		}
	}
	return readNd6ravariablesFunc(d, meta)
}

func deleteNd6ravariablesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNd6ravariablesFunc")
	// nd6ravariables does not support DELETE operation
	d.SetId("")

	return nil
}
