package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/rewrite"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRewriteaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewriteactionFunc,
		Read:          readRewriteactionFunc,
		Update:        updateRewriteactionFunc,
		Delete:        deleteRewriteactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refinesearch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"search": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stringbuilderexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createRewriteactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewriteactionFunc")
	client := meta.(*NetScalerNitroClient).client
	var rewriteactionName string
	if v, ok := d.GetOk("name"); ok {
		rewriteactionName = v.(string)
	} else {
		rewriteactionName = resource.PrefixedUniqueId("tf-rewriteaction-")
		d.Set("name", rewriteactionName)
	}
	rewriteaction := rewrite.Rewriteaction{
		Comment:           d.Get("comment").(string),
		Name:              d.Get("name").(string),
		Refinesearch:      d.Get("refinesearch").(string),
		Search:            d.Get("search").(string),
		Stringbuilderexpr: d.Get("stringbuilderexpr").(string),
		Target:            d.Get("target").(string),
		Type:              d.Get("type").(string),
	}

	_, err := client.AddResource(service.Rewriteaction.Type(), rewriteactionName, &rewriteaction)
	if err != nil {
		return err
	}

	d.SetId(rewriteactionName)

	err = readRewriteactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rewriteaction but we can't read it ?? %s", rewriteactionName)
		return nil
	}
	return nil
}

func readRewriteactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewriteactionFunc")
	client := meta.(*NetScalerNitroClient).client
	rewriteactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rewriteaction state %s", rewriteactionName)
	data, err := client.FindResource(service.Rewriteaction.Type(), rewriteactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewriteaction state %s", rewriteactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("refinesearch", data["refinesearch"])
	d.Set("search", data["search"])
	d.Set("stringbuilderexpr", data["stringbuilderexpr"])
	d.Set("target", data["target"])
	d.Set("type", data["type"])

	return nil

}

func updateRewriteactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewriteactionFunc")
	client := meta.(*NetScalerNitroClient).client
	rewriteactionName := d.Get("name").(string)

	rewriteaction := rewrite.Rewriteaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("refinesearch") {
		log.Printf("[DEBUG]  citrixadc-provider: Refinesearch has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Refinesearch = d.Get("refinesearch").(string)
		hasChange = true
	}
	if d.HasChange("search") {
		log.Printf("[DEBUG]  citrixadc-provider: Search has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Search = d.Get("search").(string)
		hasChange = true
	}
	if d.HasChange("stringbuilderexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Stringbuilderexpr has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Stringbuilderexpr = d.Get("stringbuilderexpr").(string)
		hasChange = true
	}
	if d.HasChange("target") {
		log.Printf("[DEBUG]  citrixadc-provider: Target has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Target = d.Get("target").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Rewriteaction.Type(), rewriteactionName, &rewriteaction)
		if err != nil {
			return fmt.Errorf("Error updating rewriteaction %s.\n%s", rewriteactionName, err)
		}
	}
	return readRewriteactionFunc(d, meta)
}

func deleteRewriteactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteactionFunc")
	client := meta.(*NetScalerNitroClient).client
	rewriteactionName := d.Id()
	err := client.DeleteResource(service.Rewriteaction.Type(), rewriteactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
