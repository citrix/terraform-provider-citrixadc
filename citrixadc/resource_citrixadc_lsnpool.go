package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcLsnpool() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnpoolFunc,
		Read:          readLsnpoolFunc,
		Update:        updateLsnpoolFunc,
		Delete:        deleteLsnpoolFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"poolname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nattype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"maxportrealloctmq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portblockallocation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"portrealloctimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnpoolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Get("poolname").(string)
	
	
	lsnpool := make(map[string]interface{})
	if v, ok := d.GetOkExists("portrealloctimeout"); ok {
		lsnpool["portrealloctimeout"] = v.(int)
	}
	if v, ok := d.GetOk("portblockallocation"); ok {
		lsnpool["portblockallocation"] = v.(string)
	}
	if v, ok := d.GetOkExists("poolname"); ok {
		lsnpool["poolname"] = v.(string)
	} 
	if v, ok := d.GetOk("nattype"); ok {
		lsnpool["nattype"] = v.(string)
	}
	if v, ok := d.GetOkExists("maxportrealloctmq"); ok {
		val, _ := strconv.Atoi(v.(string))
		lsnpool["maxportrealloctmq"] = val
	}

	_, err := client.AddResource("lsnpool", lsnpoolName, &lsnpool)
	if err != nil {
		return err
	}

	d.SetId(lsnpoolName)

	err = readLsnpoolFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnpool but we can't read it ?? %s", lsnpoolName)
		return nil
	}
	return nil
}

func readLsnpoolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnpool state %s", lsnpoolName)
	data, err := client.FindResource("lsnpool", lsnpoolName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnpool state %s", lsnpoolName)
		d.SetId("")
		return nil
	}
	d.Set("maxportrealloctmq", data["maxportrealloctmq"])
	d.Set("nattype", data["nattype"])
	d.Set("poolname", data["poolname"])
	d.Set("portblockallocation", data["portblockallocation"])
	d.Set("portrealloctimeout", data["portrealloctimeout"])

	return nil

}

func updateLsnpoolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Get("poolname").(string)

	lsnpool := lsn.Lsnpool{
		Poolname: d.Get("poolname").(string),
	}
	hasChange := false
	if d.HasChange("maxportrealloctmq") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxportrealloctmq has changed for lsnpool %s, starting update", lsnpoolName)
		val, _ := strconv.Atoi(d.Get("maxportrealloctmq").(string))
		lsnpool.Maxportrealloctmq = val
		hasChange = true
	}
	if d.HasChange("portrealloctimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Portrealloctimeout has changed for lsnpool %s, starting update", lsnpoolName)
		lsnpool.Portrealloctimeout = d.Get("portrealloctimeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnpool", &lsnpool)
		if err != nil {
			return fmt.Errorf("Error updating lsnpool %s", lsnpoolName)
		}
	}
	return readLsnpoolFunc(d, meta)
}

func deleteLsnpoolFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnpoolFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnpoolName := d.Id()
	err := client.DeleteResource("lsnpool", lsnpoolName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
