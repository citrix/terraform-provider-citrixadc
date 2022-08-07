package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcClusterinstance() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusterinstanceFunc,
		Read:          readClusterinstanceFunc,
		Update:        updateClusterinstanceFunc,
		Delete:        deleteClusterinstanceFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"backplanebasedview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deadinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hellointerval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"inc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodegroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preemption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"processlocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quorumtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retainconnectionsoncluster": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncstatusstrictmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createClusterinstanceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Get("clid").(int)

	clusterinstance := cluster.Clusterinstance{
		Backplanebasedview:         d.Get("backplanebasedview").(string),
		Clid:                       d.Get("clid").(int),
		Deadinterval:               d.Get("deadinterval").(int),
		Hellointerval:              d.Get("hellointerval").(int),
		Inc:                        d.Get("inc").(string),
		Nodegroup:                  d.Get("nodegroup").(string),
		Preemption:                 d.Get("preemption").(string),
		Processlocal:               d.Get("processlocal").(string),
		Quorumtype:                 d.Get("quorumtype").(string),
		Retainconnectionsoncluster: d.Get("retainconnectionsoncluster").(string),
		Syncstatusstrictmode:       d.Get("syncstatusstrictmode").(string),
	}


	_, err := client.AddResource(service.Clusterinstance.Type(), strconv.Itoa(clusterinstanceName), &clusterinstance)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(clusterinstanceName))

	err = readClusterinstanceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusterinstance but we can't read it ?? %s", strconv.Itoa(clusterinstanceName))
		return nil
	}
	return nil
}

func readClusterinstanceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading clusterinstance state %s", clusterinstanceName)
	data, err := client.FindResource(service.Clusterinstance.Type(), clusterinstanceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing clusterinstance state %s", clusterinstanceName)
		d.SetId("")
		return nil
	}
	d.Set("backplanebasedview", data["backplanebasedview"])
	d.Set("clid", data["clid"])
	d.Set("deadinterval", data["deadinterval"])
	d.Set("hellointerval", data["hellointerval"])
	d.Set("inc", data["inc"])
	d.Set("nodegroup", data["nodegroup"])
	d.Set("preemption", data["preemption"])
	d.Set("processlocal", data["processlocal"])
	d.Set("quorumtype", data["quorumtype"])
	d.Set("retainconnectionsoncluster", data["retainconnectionsoncluster"])
	d.Set("syncstatusstrictmode", data["syncstatusstrictmode"])

	return nil

}

func updateClusterinstanceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Get("clid").(int)

	clusterinstance := cluster.Clusterinstance{
		Clid: d.Get("clid").(int),
	}
	hasChange := false
	if d.HasChange("backplanebasedview") {
		log.Printf("[DEBUG]  citrixadc-provider: Backplanebasedview has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Backplanebasedview = d.Get("backplanebasedview").(string)
		hasChange = true
	}
	if d.HasChange("deadinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Deadinterval has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Deadinterval = d.Get("deadinterval").(int)
		hasChange = true
	}
	if d.HasChange("hellointerval") {
		log.Printf("[DEBUG]  citrixadc-provider: Hellointerval has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Hellointerval = d.Get("hellointerval").(int)
		hasChange = true
	}
	if d.HasChange("inc") {
		log.Printf("[DEBUG]  citrixadc-provider: Inc has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Inc = d.Get("inc").(string)
		hasChange = true
	}
	if d.HasChange("nodegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodegroup has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Nodegroup = d.Get("nodegroup").(string)
		hasChange = true
	}
	if d.HasChange("preemption") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemption has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Preemption = d.Get("preemption").(string)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  citrixadc-provider: Processlocal has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("quorumtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Quorumtype has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Quorumtype = d.Get("quorumtype").(string)
		hasChange = true
	}
	if d.HasChange("retainconnectionsoncluster") {
		log.Printf("[DEBUG]  citrixadc-provider: Retainconnectionsoncluster has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Retainconnectionsoncluster = d.Get("retainconnectionsoncluster").(string)
		hasChange = true
	}
	if d.HasChange("syncstatusstrictmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Syncstatusstrictmode has changed for clusterinstance %s, starting update", strconv.Itoa(clusterinstanceName))
		clusterinstance.Syncstatusstrictmode = d.Get("syncstatusstrictmode").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Clusterinstance.Type(), &clusterinstance)
		if err != nil {
			return fmt.Errorf("Error updating clusterinstance %s", strconv.Itoa(clusterinstanceName))
		}
	}
	return readClusterinstanceFunc(d, meta)
}

func deleteClusterinstanceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusterinstanceFunc")
	client := meta.(*NetScalerNitroClient).client
	clusterinstanceName := d.Id()
	err := client.DeleteResource(service.Clusterinstance.Type(), clusterinstanceName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
