package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcClusternode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternodeFunc,
		Read:          readClusternodeFunc,
		Update:        updateClusternodeFunc,
		Delete:        deleteClusternodeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"nodeid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backplane": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clearnodegroupconfig": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nodegroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnelmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createClusternodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := strconv.Itoa(d.Get("nodeid").(int))

	clusternode := cluster.Clusternode{
		Backplane:  d.Get("backplane").(string),
		Delay:      d.Get("delay").(int),
		Ipaddress:  d.Get("ipaddress").(string),
		Nodegroup:  d.Get("nodegroup").(string),
		Nodeid:     d.Get("nodeid").(int),
		Priority:   d.Get("priority").(int),
		State:      d.Get("state").(string),
		Tunnelmode: d.Get("tunnelmode").(string),
	}

	_, err := client.AddResource(service.Clusternode.Type(), clusternodeId, &clusternode)
	if err != nil {
		return err
	}

	d.SetId(clusternodeId)

	err = readClusternodeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternode but we can't read it ?? %s", clusternodeId)
		return nil
	}
	return nil
}

func readClusternodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading clusternode state %s", clusternodeId)
	data, err := client.FindResource(service.Clusternode.Type(), clusternodeId)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing clusternode state %s", clusternodeId)
		d.SetId("")
		return nil
	}
	d.Set("nodeid", data["nodeid"])
	d.Set("backplane", data["backplane"])
	d.Set("delay", data["delay"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("nodegroup", data["nodegroup"])
	setToInt("priority", d, data["priority"])
	d.Set("state", data["state"])
	d.Set("tunnelmode", data["tunnelmode"])

	return nil

}

func updateClusternodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := strconv.Itoa(d.Get("nodeid").(int))

	clusternode := cluster.Clusternode{
		Nodeid: d.Get("nodeid").(int),
	}
	hasChange := false
	if d.HasChange("backplane") {
		log.Printf("[DEBUG]  citrixadc-provider: Backplane has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Backplane = d.Get("backplane").(string)
		hasChange = true
	}
	if d.HasChange("delay") {
		log.Printf("[DEBUG]  citrixadc-provider: Delay has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Delay = d.Get("delay").(int)
		hasChange = true
	}
	if d.HasChange("nodegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodegroup has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Nodegroup = d.Get("nodegroup").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Nodeid = d.Get("nodeid").(int)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for clusternode %s, starting update", clusternodeId)
		clusternode.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("tunnelmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Tunnelmode has changed for clusternode %s, starting update", clusternodeId)
		clusternode.Tunnelmode = d.Get("tunnelmode").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Clusternode.Type(), clusternodeId, &clusternode)
		if err != nil {
			return fmt.Errorf("Error updating clusternode %s", clusternodeId)
		}
	}
	return readClusternodeFunc(d, meta)
}

func deleteClusternodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodeFunc")
	client := meta.(*NetScalerNitroClient).client
	clusternodeId := d.Id()
	args := make([]string, 0)
	if v, ok := d.GetOk("clearnodegroupconfig"); ok {
		args = append(args, fmt.Sprintf("clearnodegroupconfig:%s", v.(string)))
	} else {
		args = append(args, fmt.Sprintf("clearnodegroupconfig:YES"))
	}
	err := client.DeleteResourceWithArgs(service.Clusternode.Type(), clusternodeId, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
