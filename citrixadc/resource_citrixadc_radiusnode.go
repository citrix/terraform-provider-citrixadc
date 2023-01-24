package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcRadiusnode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRadiusnodeFunc,
		Read:          readRadiusnodeFunc,
		Update:        updateRadiusnodeFunc,
		Delete:        deleteRadiusnodeFunc,
		Schema: map[string]*schema.Schema{
			"nodeprefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"radkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createRadiusnodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Get("nodeprefix").(string)
	radiusnode := basic.Radiusnode{
		Nodeprefix: d.Get("nodeprefix").(string),
		Radkey:     d.Get("radkey").(string),
	}

	_, err := client.AddResource("radiusnode", radiusnodeName, &radiusnode)
	if err != nil {
		return err
	}

	d.SetId(radiusnodeName)

	err = readRadiusnodeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this radiusnode but we can't read it ?? %s", radiusnodeName)
		return nil
	}
	return nil
}

func readRadiusnodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading radiusnode state %s", radiusnodeName)
	radiusnodeescaped := url.PathEscape(url.QueryEscape(radiusnodeName))
	data, err := client.FindResource("radiusnode", radiusnodeescaped)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing radiusnode state %s", radiusnodeName)
		d.SetId("")
		return nil
	}
	d.Set("nodeprefix", data["nodeprefix"])
	// d.Set("radkey", data["radkey"])

	return nil

}

func updateRadiusnodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Get("nodeprefix").(string)
	radiusnodeescaped := url.PathEscape(url.QueryEscape(radiusnodeName))


	radiusnode := basic.Radiusnode{
		Nodeprefix: d.Get("nodeprefix").(string),
	}
	hasChange := false
	if d.HasChange("radkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Radkey has changed for radiusnode %s, starting update", radiusnodeName)
		radiusnode.Radkey = d.Get("radkey").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("radiusnode", radiusnodeescaped, &radiusnode)
		if err != nil {
			return fmt.Errorf("Error updating radiusnode %s", radiusnodeName)
		}
	}
	return readRadiusnodeFunc(d, meta)
}

func deleteRadiusnodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRadiusnodeFunc")
	client := meta.(*NetScalerNitroClient).client
	radiusnodeName := d.Id()
	err := client.DeleteResource("radiusnode", url.QueryEscape(url.QueryEscape(radiusnodeName)))
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
