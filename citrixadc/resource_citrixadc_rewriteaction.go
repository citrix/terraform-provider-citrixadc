package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/rewrite"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

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
		CustomizeDiff: customizeRewriteactionDiff,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bypasssafetycheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refinesearch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stringbuilderexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
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
		Bypasssafetycheck: d.Get("bypasssafetycheck").(string),
		Comment:           d.Get("comment").(string),
		Name:              d.Get("name").(string),
		Pattern:           d.Get("pattern").(string),
		Refinesearch:      d.Get("refinesearch").(string),
		Search:            d.Get("search").(string),
		Stringbuilderexpr: d.Get("stringbuilderexpr").(string),
		Target:            d.Get("target").(string),
		Type:              d.Get("type").(string),
	}

	_, err := client.AddResource(netscaler.Rewriteaction.Type(), rewriteactionName, &rewriteaction)
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
	data, err := client.FindResource(netscaler.Rewriteaction.Type(), rewriteactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewriteaction state %s", rewriteactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("bypasssafetycheck", data["bypasssafetycheck"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("pattern", data["pattern"])
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
	if d.HasChange("bypasssafetycheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Bypasssafetycheck has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Bypasssafetycheck = d.Get("bypasssafetycheck").(string)
		hasChange = true
	}
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
	if d.HasChange("pattern") {
		log.Printf("[DEBUG]  citrixadc-provider: Pattern has changed for rewriteaction %s, starting update", rewriteactionName)
		rewriteaction.Pattern = d.Get("pattern").(string)
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
		_, err := client.UpdateResource(netscaler.Rewriteaction.Type(), rewriteactionName, &rewriteaction)
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
	err := client.DeleteResource(netscaler.Rewriteaction.Type(), rewriteactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func customizeRewriteactionDiff(diff *schema.ResourceDiff, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In customizeRewriteactionDiff")
	o := diff.GetChangedKeysPrefix("")

	// Check if target and bypasssafetycheck is in changed keys
	targetDefined := false
	bypasssafetycheckDefined := false

	for _, element := range o {

		if element == "target" {
			targetDefined = true
		}

		if element == "bypasssafetycheck" {
			bypasssafetycheckDefined = true
		}
	}

	// Clear bypasssafetycheck when present without target
	if bypasssafetycheckDefined && !targetDefined {
		diff.Clear("bypasssafetycheck")
	}

	return nil
}
