package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/policy"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcPolicydataset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicydatasetFunc,
		Read:          readPolicydatasetFunc,
		Delete:        deletePolicydatasetFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"indextype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicydatasetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicydatasetFunc")
	client := meta.(*NetScalerNitroClient).client
	var policydatasetName string
	if v, ok := d.GetOk("name"); ok {
		policydatasetName = v.(string)
	} else {
		policydatasetName = resource.PrefixedUniqueId("tf-policydataset-")
		d.Set("name", policydatasetName)
	}
	policydataset := policy.Policydataset{
		Comment:   d.Get("comment").(string),
		Indextype: d.Get("indextype").(string),
		Name:      d.Get("name").(string),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(netscaler.Policydataset.Type(), policydatasetName, &policydataset)
	if err != nil {
		return err
	}

	d.SetId(policydatasetName)

	err = readPolicydatasetFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policydataset but we can't read it ?? %s", policydatasetName)
		return nil
	}
	return nil
}

func readPolicydatasetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicydatasetFunc")
	client := meta.(*NetScalerNitroClient).client
	policydatasetName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policydataset state %s", policydatasetName)
	data, err := client.FindResource(netscaler.Policydataset.Type(), policydatasetName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policydataset state %s", policydatasetName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("indextype", data["indextype"])
	d.Set("name", data["name"])
	d.Set("type", data["type"])

	return nil

}

func deletePolicydatasetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicydatasetFunc")
	client := meta.(*NetScalerNitroClient).client
	policydatasetName := d.Id()
	err := client.DeleteResource(netscaler.Policydataset.Type(), policydatasetName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
