package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/feo"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcFeoaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFeoactionFunc,
		Read:          readFeoactionFunc,
		Update:        updateFeoactionFunc,
		Delete:        deleteFeoactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cachemaxage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"clientsidemeasurements": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"convertimporttolink": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"csscombine": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cssimginline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cssinline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cssminify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cssmovetohead": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"dnsshards": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"domainsharding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htmlminify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"imggiftopng": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"imginline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"imglazyload": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"imgshrinktoattrib": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"imgtojpegxr": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"imgtowebp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jpgoptimize": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jsinline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jsminify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jsmovetoend": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"pageextendcache": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createFeoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFeoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	feoactionName := d.Get("name").(string)
	
	feoaction := make(map[string]interface{})

	if v, ok := d.GetOkExists("cachemaxage"); ok {
		feoaction["cachemaxage"] = v.(int)
	}
	if v, ok := d.GetOk("clientsidemeasurements"); ok {
		feoaction["clientsidemeasurements"] = v.(bool)
	}
	if v, ok := d.GetOk("convertimporttolink"); ok {
		feoaction["convertimporttolink"] = v.(bool)
	}
	if v, ok := d.GetOk("cssminify"); ok {
		feoaction["cssminify"] = v.(bool)
	}
	if v, ok := d.GetOk("cssmovetohead"); ok {
		feoaction["cssmovetohead"] = v.(bool)
	}
	if v, ok := d.GetOk("cssinline"); ok {
		feoaction["cssinline"] = v.(bool)
	}
	if v, ok := d.GetOk("cssimginline"); ok {
		feoaction["cssimginline"] = v.(bool)
	}
	if v, ok := d.GetOk("csscombine"); ok {
		feoaction["csscombine"] = v.(bool)
	}
	if v, ok := d.GetOk("dnsshards"); ok {
		feoaction["dnsshards"] = toIntegerList(v.([]interface{}))
	}
	if v, ok := d.GetOk("imgtowebp"); ok {
		feoaction["imgtowebp"] = v.(bool)
	}
	if v, ok := d.GetOk("imgtojpegxr"); ok {
		feoaction["imgtojpegxr"] = v.(bool)
	}
	if v, ok := d.GetOk("imgshrinktoattrib"); ok {
		feoaction["imgshrinktoattrib"] = v.(bool)
	}
	if v, ok := d.GetOk("imglazyload"); ok {
		feoaction["imglazyload"] = v.(bool)
	}
	if v, ok := d.GetOk("imggiftopng"); ok {
		feoaction["imggiftopng"] = v.(bool)
	}
	if v, ok := d.GetOk("htmlminify"); ok {
		feoaction["htmlminify"] = v.(bool)
	}
	if v, ok := d.GetOk("domainsharding"); ok {
		feoaction["domainsharding"] = v.(string)
	}
	if v, ok := d.GetOk("dnsshards"); ok {
		feoaction["dnsshards"] = toIntegerList(v.([]interface{}))
	}
	if v, ok := d.GetOk("pageextendcache"); ok {
		feoaction["pageextendcache"] = v.(bool)
	}
	if v, ok := d.GetOk("name"); ok {
		feoaction["name"] = v.(string)
	}
	if v, ok := d.GetOk("jsmovetoend"); ok {
		feoaction["jsmovetoend"] = v.(bool)
	}
	if v, ok := d.GetOk("jsminify"); ok {
		feoaction["jsminify"] = v.(bool)
	}
	if v, ok := d.GetOk("jsinline"); ok {
		feoaction["jsinline"] = v.(bool)
	}
	if v, ok := d.GetOk("jpgoptimize"); ok {
		feoaction["jpgoptimize"] = v.(string)
	}
	_, err := client.AddResource("feoaction", feoactionName, &feoaction)
	if err != nil {
		return err
	}

	d.SetId(feoactionName)

	err = readFeoactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this feoaction but we can't read it ?? %s", feoactionName)
		return nil
	}
	return nil
}

func readFeoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readFeoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	feoactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading feoaction state %s", feoactionName)
	data, err := client.FindResource("feoaction", feoactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing feoaction state %s", feoactionName)
		d.SetId("")
		return nil
	}
	d.Set("cachemaxage", data["cachemaxage"])
	d.Set("clientsidemeasurements", data["clientsidemeasurements"])
	d.Set("convertimporttolink", data["convertimporttolink"])
	d.Set("csscombine", data["csscombine"])
	d.Set("cssimginline", data["cssimginline"])
	d.Set("cssinline", data["cssinline"])
	d.Set("cssminify", data["cssminify"])
	d.Set("cssmovetohead", data["cssmovetohead"])
	d.Set("dnsshards", data["dnsshards"])
	d.Set("domainsharding", data["domainsharding"])
	d.Set("htmlminify", data["htmlminify"])
	d.Set("imggiftopng", data["imggiftopng"])
	d.Set("imginline", data["imginline"])
	d.Set("imglazyload", data["imglazyload"])
	d.Set("imgshrinktoattrib", data["imgshrinktoattrib"])
	d.Set("imgtojpegxr", data["imgtojpegxr"])
	d.Set("imgtowebp", data["imgtowebp"])
	d.Set("jpgoptimize", data["jpgoptimize"])
	d.Set("jsinline", data["jsinline"])
	d.Set("jsminify", data["jsminify"])
	d.Set("jsmovetoend", data["jsmovetoend"])
	d.Set("pageextendcache", data["pageextendcache"])

	return nil

}

func updateFeoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFeoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	feoactionName := d.Get("name").(string)

	feoaction := feo.Feoaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("cachemaxage") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachemaxage has changed for feoaction %s, starting update", feoactionName)
		feoaction.Cachemaxage = d.Get("cachemaxage").(int)
		hasChange = true
	}
	if d.HasChange("clientsidemeasurements") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsidemeasurements has changed for feoaction %s, starting update", feoactionName)
		feoaction.Clientsidemeasurements = d.Get("clientsidemeasurements").(bool)
		hasChange = true
	}
	if d.HasChange("convertimporttolink") {
		log.Printf("[DEBUG]  citrixadc-provider: Convertimporttolink has changed for feoaction %s, starting update", feoactionName)
		feoaction.Convertimporttolink = d.Get("convertimporttolink").(bool)
		hasChange = true
	}
	if d.HasChange("csscombine") {
		log.Printf("[DEBUG]  citrixadc-provider: Csscombine has changed for feoaction %s, starting update", feoactionName)
		feoaction.Csscombine = d.Get("csscombine").(bool)
		hasChange = true
	}
	if d.HasChange("cssimginline") {
		log.Printf("[DEBUG]  citrixadc-provider: Cssimginline has changed for feoaction %s, starting update", feoactionName)
		feoaction.Cssimginline = d.Get("cssimginline").(bool)
		hasChange = true
	}
	if d.HasChange("cssinline") {
		log.Printf("[DEBUG]  citrixadc-provider: Cssinline has changed for feoaction %s, starting update", feoactionName)
		feoaction.Cssinline = d.Get("cssinline").(bool)
		hasChange = true
	}
	if d.HasChange("cssminify") {
		log.Printf("[DEBUG]  citrixadc-provider: Cssminify has changed for feoaction %s, starting update", feoactionName)
		feoaction.Cssminify = d.Get("cssminify").(bool)
		hasChange = true
	}
	if d.HasChange("cssmovetohead") {
		log.Printf("[DEBUG]  citrixadc-provider: Cssmovetohead has changed for feoaction %s, starting update", feoactionName)
		feoaction.Cssmovetohead = d.Get("cssmovetohead").(bool)
		hasChange = true
	}
	if d.HasChange("dnsshards") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsshards has changed for feoaction %s, starting update", feoactionName)
		feoaction.Dnsshards = toStringList(d.Get("dnsshards").([]interface{}))
		hasChange = true
	}
	if d.HasChange("domainsharding") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainsharding has changed for feoaction %s, starting update", feoactionName)
		feoaction.Domainsharding = d.Get("domainsharding").(string)
		hasChange = true
	}
	if d.HasChange("htmlminify") {
		log.Printf("[DEBUG]  citrixadc-provider: Htmlminify has changed for feoaction %s, starting update", feoactionName)
		feoaction.Htmlminify = d.Get("htmlminify").(bool)
		hasChange = true
	}
	if d.HasChange("imggiftopng") {
		log.Printf("[DEBUG]  citrixadc-provider: Imggiftopng has changed for feoaction %s, starting update", feoactionName)
		feoaction.Imggiftopng = d.Get("imggiftopng").(bool)
		hasChange = true
	}
	if d.HasChange("imginline") {
		log.Printf("[DEBUG]  citrixadc-provider: Imginline has changed for feoaction %s, starting update", feoactionName)
		feoaction.Imginline = d.Get("imginline").(bool)
		hasChange = true
	}
	if d.HasChange("imglazyload") {
		log.Printf("[DEBUG]  citrixadc-provider: Imglazyload has changed for feoaction %s, starting update", feoactionName)
		feoaction.Imglazyload = d.Get("imglazyload").(bool)
		hasChange = true
	}
	if d.HasChange("imgshrinktoattrib") {
		log.Printf("[DEBUG]  citrixadc-provider: Imgshrinktoattrib has changed for feoaction %s, starting update", feoactionName)
		feoaction.Imgshrinktoattrib = d.Get("imgshrinktoattrib").(bool)
		hasChange = true
	}
	if d.HasChange("imgtojpegxr") {
		log.Printf("[DEBUG]  citrixadc-provider: Imgtojpegxr has changed for feoaction %s, starting update", feoactionName)
		feoaction.Imgtojpegxr = d.Get("imgtojpegxr").(bool)
		hasChange = true
	}
	if d.HasChange("imgtowebp") {
		log.Printf("[DEBUG]  citrixadc-provider: Imgtowebp has changed for feoaction %s, starting update", feoactionName)
		feoaction.Imgtowebp = d.Get("imgtowebp").(bool)
		hasChange = true
	}
	if d.HasChange("jpgoptimize") {
		log.Printf("[DEBUG]  citrixadc-provider: Jpgoptimize has changed for feoaction %s, starting update", feoactionName)
		feoaction.Jpgoptimize = d.Get("jpgoptimize").(bool)
		hasChange = true
	}
	if d.HasChange("jsinline") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsinline has changed for feoaction %s, starting update", feoactionName)
		feoaction.Jsinline = d.Get("jsinline").(bool)
		hasChange = true
	}
	if d.HasChange("jsminify") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsminify has changed for feoaction %s, starting update", feoactionName)
		feoaction.Jsminify = d.Get("jsminify").(bool)
		hasChange = true
	}
	if d.HasChange("jsmovetoend") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsmovetoend has changed for feoaction %s, starting update", feoactionName)
		feoaction.Jsmovetoend = d.Get("jsmovetoend").(bool)
		hasChange = true
	}
	if d.HasChange("pageextendcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Pageextendcache has changed for feoaction %s, starting update", feoactionName)
		feoaction.Pageextendcache = d.Get("pageextendcache").(bool)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("feoaction", feoactionName, &feoaction)
		if err != nil {
			return fmt.Errorf("Error updating feoaction %s", feoactionName)
		}
	}
	return readFeoactionFunc(d, meta)
}

func deleteFeoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFeoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	feoactionName := d.Id()
	err := client.DeleteResource("feoaction", feoactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
