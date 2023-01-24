package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNshmackey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNshmackeyFunc,
		Read:          readNshmackeyFunc,
		Update:        updateNshmackeyFunc,
		Delete:        deleteNshmackeyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"digest": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyvalue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNshmackeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNshmackeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nshmackeyName := d.Get("name").(string)
	nshmackey := ns.Nshmackey{
		Comment:  d.Get("comment").(string),
		Digest:   d.Get("digest").(string),
		Keyvalue: d.Get("keyvalue").(string),
		Name:     d.Get("name").(string),
	}

	_, err := client.AddResource("nshmackey", nshmackeyName, &nshmackey)
	if err != nil {
		return err
	}

	d.SetId(nshmackeyName)

	err = readNshmackeyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nshmackey but we can't read it ?? %s", nshmackeyName)
		return nil
	}
	return nil
}

func readNshmackeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNshmackeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nshmackeyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nshmackey state %s", nshmackeyName)
	data, err := client.FindResource("nshmackey", nshmackeyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nshmackey state %s", nshmackeyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("digest", data["digest"])
	// d.Set("keyvalue", data["keyvalue"])

	return nil

}

func updateNshmackeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNshmackeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nshmackeyName := d.Get("name").(string)

	nshmackey := ns.Nshmackey{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for nshmackey %s, starting update", nshmackeyName)
		nshmackey.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("digest") {
		log.Printf("[DEBUG]  citrixadc-provider: Digest has changed for nshmackey %s, starting update", nshmackeyName)
		nshmackey.Digest = d.Get("digest").(string)
		hasChange = true
	}
	if d.HasChange("keyvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Keyvalue has changed for nshmackey %s, starting update", nshmackeyName)
		nshmackey.Keyvalue = d.Get("keyvalue").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("nshmackey", nshmackeyName, &nshmackey)
		if err != nil {
			return fmt.Errorf("Error updating nshmackey %s", nshmackeyName)
		}
	}
	return readNshmackeyFunc(d, meta)
}

func deleteNshmackeyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNshmackeyFunc")
	client := meta.(*NetScalerNitroClient).client
	nshmackeyName := d.Id()
	err := client.DeleteResource("nshmackey", nshmackeyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
