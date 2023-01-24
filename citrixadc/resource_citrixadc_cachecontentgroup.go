package citrixadc

import (

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCachecontentgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCachecontentgroupFunc,
		Read:          readCachecontentgroupFunc,
		Update:        updateCachecontentgroupFunc,
		Delete:        deleteCachecontentgroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"absexpiry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"absexpirygmt": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"alwaysevalpolicies": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachecontrol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expireatlastbyte": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flashcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heurexpiryparam": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hitparams": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hitselector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignoreparamvaluecase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignorereloadreq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignorereqcachinghdrs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertetag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertvia": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalparams": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"invalrestrictedtohost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalselector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lazydnsresolve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"matchcookies": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minhits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"persistha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pinned": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polleverytime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefetch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefetchmaxpending": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"prefetchperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"prefetchperiodmillisec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quickabortsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"relexpiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"relexpirymillisec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"removecookies": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selectorvalue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tosecondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weaknegrelexpiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weakposrelexpiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCachecontentgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCachecontentgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	cachecontentgroupName := d.Get("name").(string)

	cachecontentgroup := make(map[string]interface{})
	if v, ok := d.GetOk("prefetch"); ok {
		cachecontentgroup["prefetch"] = v.(string)
	}
	if v, ok := d.GetOk("polleverytime"); ok {
		cachecontentgroup["polleverytime"] = v.(string)
	} 
	if v, ok := d.GetOk("pinned"); ok {
		cachecontentgroup["pinned"] = v.(string)
	}
	if v, ok := d.GetOk("persistha"); ok {
		cachecontentgroup["persistha"] = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		cachecontentgroup["name"] = v.(string)
	}
	if v, ok := d.GetOkExists("minressize"); ok {
		cachecontentgroup["minressize"] = v.(int)
	} 
	if v, ok := d.GetOkExists("minhits"); ok {
		cachecontentgroup["minhits"] = v.(int)
	}
	if v, ok := d.GetOk("memlimit"); ok {
		cachecontentgroup["memlimit"] = v.(int)
	}	
	if v, ok := d.GetOkExists("maxressize"); ok {
		cachecontentgroup["maxressize"] = v.(int)
	}
	if v, ok := d.GetOk("matchcookies"); ok {
		cachecontentgroup["matchcookies"] = v.(string)
	} 
	if v, ok := d.GetOk("lazydnsresolve"); ok {
		cachecontentgroup["lazydnsresolve"] = v.(string)
	}
	if v, ok := d.GetOk("invalselector"); ok {
		cachecontentgroup["invalselector"] = v.(string)
	}	
	if v, ok := d.GetOk("invalrestrictedtohost"); ok {
		cachecontentgroup["invalrestrictedtohost"] = v.(string)
	}
	if v, ok := d.GetOk("invalparams"); ok {
		cachecontentgroup["invalparams"] = toStringList(v.([]interface{}))
	} 
	if v, ok := d.GetOk("insertvia"); ok {
		cachecontentgroup["insertvia"] = v.(string)
	}
	if v, ok := d.GetOk("insertetag"); ok {
		cachecontentgroup["insertetag"] = v.(string)
	}	
	if v, ok := d.GetOk("insertage"); ok {
		cachecontentgroup["insertage"] = v.(string)
	}
	if v, ok := d.GetOk("ignoreparamvaluecase"); ok {
		cachecontentgroup["ignoreparamvaluecase"] = v.(string)
	} 
	if v, ok := d.GetOk("host"); ok {
		cachecontentgroup["host"] = v.(string)
	}
	if v, ok := d.GetOk("hitselector"); ok {
		cachecontentgroup["hitselector"] = v.(string)
	}	
	if v, ok := d.GetOk("hitparams"); ok {
		cachecontentgroup["hitparams"] = toStringList(v.([]interface{}))
	}
	if v, ok := d.GetOkExists("heurexpiryparam"); ok {
		cachecontentgroup["heurexpiryparam"] = v.(int)
	} 
	if v, ok := d.GetOk("flashcache"); ok {
		cachecontentgroup["flashcache"] = v.(string)
	}
	if v, ok := d.GetOk("expireatlastbyte"); ok {
		cachecontentgroup["expireatlastbyte"] = v.(string)
	}	
	if v, ok := d.GetOk("cachecontrol"); ok {
		cachecontentgroup["cachecontrol"] = v.(int)
	}
	if v, ok := d.GetOk("alwaysevalpolicies"); ok {
		cachecontentgroup["alwaysevalpolicies"] = v.(string)
	} 
	if v, ok := d.GetOk("absexpirygmt"); ok {
		cachecontentgroup["absexpirygmt"] = toStringList(v.([]interface{}))
	}
	if v, ok := d.GetOk("absexpiry"); ok {
		cachecontentgroup["absexpiry"] = toStringList(v.([]interface{}))
	}
	if v, ok := d.GetOkExists("prefetchmaxpending"); ok {
		cachecontentgroup["prefetchmaxpending"] = v.(int)
	}	
	if v, ok := d.GetOkExists("prefetchperiod"); ok {
		cachecontentgroup["prefetchperiod"] = v.(int)
	}	
	if v, ok := d.GetOkExists("prefetchperiodmillisec"); ok {
		cachecontentgroup["prefetchperiodmillisec"] = v.(int)
	}	
	if v, ok := d.GetOk("query"); ok {
		cachecontentgroup["query"] = v.(string)
	}	
	if v, ok := d.GetOkExists("quickabortsize"); ok {
		cachecontentgroup["quickabortsize"] = v.(int)
	}	
	if v, ok := d.GetOkExists("relexpiry"); ok {
		cachecontentgroup["relexpiry"] = v.(int)
	}
	if v, ok := d.GetOkExists("relexpirymillisec"); ok {
		cachecontentgroup["relexpirymillisec"] = v.(int)
	}
	if v, ok := d.GetOk("removecookies"); ok {
		cachecontentgroup["removecookies"] = v.(string)
	}
	if v, ok := d.GetOk("selectorvalue"); ok {
		cachecontentgroup["selectorvalue"] = v.(string)
	}
	if v, ok := d.GetOk("tosecondary"); ok {
		cachecontentgroup["tosecondary"] = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		cachecontentgroup["type"] = v.(string)
	}
	if v, ok := d.GetOkExists("weaknegrelexpiry"); ok {
		cachecontentgroup["weaknegrelexpiry"] = v.(int)
	}
	if v, ok := d.GetOk("weakposrelexpiry"); ok {
		cachecontentgroup["weakposrelexpiry"] = v.(int)
	}



	// cachecontentgroup := cache.Cachecontentgroup{
	// 	Absexpiry:              toStringList(d.Get("absexpiry").([]interface{})),
	// 	Absexpirygmt:           toStringList(d.Get("absexpirygmt").([]interface{})),
	// 	Alwaysevalpolicies:     d.Get("alwaysevalpolicies").(string),
	// 	Cachecontrol:           d.Get("cachecontrol").(string),
	// 	Expireatlastbyte:       d.Get("expireatlastbyte").(string),
	// 	Flashcache:             d.Get("flashcache").(string),
	// 	Heurexpiryparam:        d.Get("heurexpiryparam").(int),
	// 	Hitparams:              toStringList(d.Get("hitparams").([]interface{})),
	// 	Hitselector:            d.Get("hitselector").(string),
	// 	Host:                   d.Get("host").(string),
	// 	Ignoreparamvaluecase:   d.Get("ignoreparamvaluecase").(string),
	// 	Ignorereloadreq:        d.Get("ignorereloadreq").(string),
	// 	Ignorereqcachinghdrs:   d.Get("ignorereqcachinghdrs").(string),
	// 	Insertage:              d.Get("insertage").(string),
	// 	Insertetag:             d.Get("insertetag").(string),
	// 	Insertvia:              d.Get("insertvia").(string),
	// 	Invalparams:            toStringList(d.Get("invalparams").([]interface{})),
	// 	Invalrestrictedtohost:  d.Get("invalrestrictedtohost").(string),
	// 	Invalselector:          d.Get("invalselector").(string),
	// 	Lazydnsresolve:         d.Get("lazydnsresolve").(string),
	// 	Matchcookies:           d.Get("matchcookies").(string),
	// 	Maxressize:             d.Get("maxressize").(int),
	// 	Memlimit:               d.Get("memlimit").(int),
	// 	Minhits:                d.Get("minhits").(int),
	// 	Minressize:             d.Get("minressize").(int),
	// 	Name:                   d.Get("name").(string),
	// 	Persistha:              d.Get("persistha").(string),
	// 	Pinned:                 d.Get("pinned").(string),
	// 	Polleverytime:          d.Get("polleverytime").(string),
	// 	Prefetch:               d.Get("prefetch").(string),
	// 	Prefetchmaxpending:     d.Get("prefetchmaxpending").(int),
	// 	Prefetchperiod:         d.Get("prefetchperiod").(int),
	// 	Prefetchperiodmillisec: d.Get("prefetchperiodmillisec").(int),
	// 	Query:                  d.Get("query").(string),
	// 	Quickabortsize:         d.Get("quickabortsize").(int),
	// 	Relexpiry:              d.Get("relexpiry").(int),
	// 	Relexpirymillisec:      d.Get("relexpirymillisec").(int),
	// 	Removecookies:          d.Get("removecookies").(string),
	// 	Selectorvalue:          d.Get("selectorvalue").(string),
	// 	Tosecondary:            d.Get("tosecondary").(string),
	// 	Type:                   d.Get("type").(string),
	// 	Weaknegrelexpiry:       d.Get("weaknegrelexpiry").(int),
	// 	Weakposrelexpiry:       d.Get("weakposrelexpiry").(int),
	// }

	_, err := client.AddResource(service.Cachecontentgroup.Type(), cachecontentgroupName, &cachecontentgroup)
	if err != nil {
		return err
	}

	d.SetId(cachecontentgroupName)

	err = readCachecontentgroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cachecontentgroup but we can't read it ?? %s", cachecontentgroupName)
		return nil
	}
	return nil
}

func readCachecontentgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCachecontentgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	cachecontentgroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cachecontentgroup state %s", cachecontentgroupName)
	data, err := client.FindResource(service.Cachecontentgroup.Type(), cachecontentgroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cachecontentgroup state %s", cachecontentgroupName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("absexpiry", data["absexpiry"])
	d.Set("absexpirygmt", data["absexpirygmt"])
	d.Set("alwaysevalpolicies", data["alwaysevalpolicies"])
	d.Set("cachecontrol", data["cachecontrol"])
	d.Set("expireatlastbyte", data["expireatlastbyte"])
	d.Set("flashcache", data["flashcache"])
	d.Set("heurexpiryparam", data["heurexpiryparam"])
	d.Set("hitparams", data["hitparams"])
	d.Set("hitselector", data["hitselector"])
	d.Set("host", data["host"])
	d.Set("ignoreparamvaluecase", data["ignoreparamvaluecase"])
	d.Set("ignorereloadreq", data["ignorereloadreq"])
	d.Set("ignorereqcachinghdrs", data["ignorereqcachinghdrs"])
	d.Set("insertage", data["insertage"])
	d.Set("insertetag", data["insertetag"])
	d.Set("insertvia", data["insertvia"])
	d.Set("invalparams", data["invalparams"])
	d.Set("invalrestrictedtohost", data["invalrestrictedtohost"])
	d.Set("invalselector", data["invalselector"])
	d.Set("lazydnsresolve", data["lazydnsresolve"])
	d.Set("matchcookies", data["matchcookies"])
	d.Set("maxressize", data["maxressize"])
	d.Set("memlimit", data["memlimit"])
	d.Set("minhits", data["minhits"])
	d.Set("minressize", data["minressize"])
	d.Set("persistha", data["persistha"])
	d.Set("pinned", data["pinned"])
	d.Set("polleverytime", data["polleverytime"])
	d.Set("prefetch", data["prefetch"])
	d.Set("prefetchmaxpending", data["prefetchmaxpending"])
	d.Set("prefetchperiod", data["prefetchperiod"])
	d.Set("prefetchperiodmillisec", data["prefetchperiodmillisec"])
	d.Set("query", data["query"])
	d.Set("quickabortsize", data["quickabortsize"])
	d.Set("relexpiry", data["relexpiry"])
	d.Set("relexpirymillisec", data["relexpirymillisec"])
	d.Set("removecookies", data["removecookies"])
	d.Set("selectorvalue", data["selectorvalue"])
	d.Set("tosecondary", data["tosecondary"])
	d.Set("type", data["type"])
	d.Set("weaknegrelexpiry", data["weaknegrelexpiry"])
	d.Set("weakposrelexpiry", data["weakposrelexpiry"])

	return nil

}

func updateCachecontentgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCachecontentgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	cachecontentgroupName := d.Get("name").(string)

	cachecontentgroup := make(map[string]interface{})
	cachecontentgroup["name"] =  d.Get("name").(string)
	
	hasChange := false
	if d.HasChange("absexpiry") {
		log.Printf("[DEBUG]  citrixadc-provider: Absexpiry has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["absexpiry"] = toStringList(d.Get("absexpiry").([]interface{}))
		hasChange = true
	}
	if d.HasChange("absexpirygmt") {
		log.Printf("[DEBUG]  citrixadc-provider: Absexpirygmt has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["absexpirygmt"] = toStringList(d.Get("absexpirygmt").([]interface{}))
		hasChange = true
	}
	if d.HasChange("alwaysevalpolicies") {
		log.Printf("[DEBUG]  citrixadc-provider: Alwaysevalpolicies has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["alwaysevalpolicies"] = d.Get("alwaysevalpolicies").(string)
		hasChange = true
	}
	if d.HasChange("cachecontrol") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachecontrol has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["cachecontrol"] = d.Get("cachecontrol").(string)
		hasChange = true
	}
	if d.HasChange("expireatlastbyte") {
		log.Printf("[DEBUG]  citrixadc-provider: Expireatlastbyte has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["expireatlastbyte"] = d.Get("expireatlastbyte").(string)
		hasChange = true
	}
	if d.HasChange("flashcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Flashcache has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["flashcache"] = d.Get("flashcache").(string)
		hasChange = true
	}
	if d.HasChange("heurexpiryparam") {
		log.Printf("[DEBUG]  citrixadc-provider: Heurexpiryparam has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["heurexpiryparam"] = d.Get("heurexpiryparam").(int)
		hasChange = true
	}
	if d.HasChange("hitparams") {
		log.Printf("[DEBUG]  citrixadc-provider: Hitparams has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["hitparams"] = toStringList(d.Get("hitparams").([]interface{}))
		hasChange = true
	}
	if d.HasChange("hitselector") {
		log.Printf("[DEBUG]  citrixadc-provider: Hitselector has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["hitselector"] = d.Get("hitselector").(string)
		hasChange = true
	}
	if d.HasChange("host") {
		log.Printf("[DEBUG]  citrixadc-provider: Host has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["host"] = d.Get("host").(string)
		hasChange = true
	}
	if d.HasChange("ignoreparamvaluecase") {
		log.Printf("[DEBUG]  citrixadc-provider: Ignoreparamvaluecase has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["ignoreparamvaluecase"] = d.Get("ignoreparamvaluecase").(string)
		hasChange = true
	}
	if d.HasChange("ignorereloadreq") {
		log.Printf("[DEBUG]  citrixadc-provider: Ignorereloadreq has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["ignorereloadreq"] = d.Get("ignorereloadreq").(string)
		hasChange = true
	}
	if d.HasChange("ignorereqcachinghdrs") {
		log.Printf("[DEBUG]  citrixadc-provider: Ignorereqcachinghdrs has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["ignorereqcachinghdrs"] = d.Get("ignorereqcachinghdrs").(string)
		hasChange = true
	}
	if d.HasChange("insertage") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertage has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["insertage"] = d.Get("insertage").(string)
		hasChange = true
	}
	if d.HasChange("insertetag") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertetag has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["insertetag"] = d.Get("insertetag").(string)
		hasChange = true
	}
	if d.HasChange("insertvia") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertvia has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["insertvia"] = d.Get("insertvia").(string)
		hasChange = true
	}
	if d.HasChange("invalparams") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalparams has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["invalparams"] = toStringList(d.Get("invalparams").([]interface{}))
		hasChange = true
	}
	if d.HasChange("invalrestrictedtohost") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalrestrictedtohost has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["invalrestrictedtohost"] = d.Get("invalrestrictedtohost").(string)
		hasChange = true
	}
	if d.HasChange("invalselector") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalselector has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["invalselector"] = d.Get("invalselector").(string)
		hasChange = true
	}
	if d.HasChange("lazydnsresolve") {
		log.Printf("[DEBUG]  citrixadc-provider: Lazydnsresolve has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["lazydnsresolve"] = d.Get("lazydnsresolve").(string)
		hasChange = true
	}
	if d.HasChange("matchcookies") {
		log.Printf("[DEBUG]  citrixadc-provider: Matchcookies has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["matchcookies"] = d.Get("matchcookies").(string)
		hasChange = true
	}
	if d.HasChange("maxressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxressize has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["maxressize"] = d.Get("maxressize").(int)
		hasChange = true
	}
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["memlimit"] = d.Get("memlimit").(int)
		hasChange = true
	}
	if d.HasChange("minhits") {
		log.Printf("[DEBUG]  citrixadc-provider: Minhits has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["minhits"] = d.Get("minhits").(int)
		hasChange = true
	}
	if d.HasChange("minressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Minressize has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["minressize"] = d.Get("minressize").(int)
		hasChange = true
	}
	if d.HasChange("persistha") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistha has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["persistha"] = d.Get("persistha").(string)
		hasChange = true
	}
	if d.HasChange("pinned") {
		log.Printf("[DEBUG]  citrixadc-provider: Pinned has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["pinned"] = d.Get("pinned").(string)
		hasChange = true
	}
	if d.HasChange("polleverytime") {
		log.Printf("[DEBUG]  citrixadc-provider: Polleverytime has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["polleverytime"] = d.Get("polleverytime").(string)
		hasChange = true
	}
	if d.HasChange("prefetch") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefetch has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["prefetch"] = d.Get("prefetch").(string)
		hasChange = true
	}
	if d.HasChange("prefetchmaxpending") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefetchmaxpending has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["prefetchmaxpending"] = d.Get("prefetchmaxpending").(int)
		hasChange = true
	}
	if d.HasChange("prefetchperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefetchperiod has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["prefetchperiod"] = d.Get("prefetchperiod").(int)
		hasChange = true
	}
	if d.HasChange("prefetchperiodmillisec") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefetchperiodmillisec has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["prefetchperiodmillisec"] = d.Get("prefetchperiodmillisec").(int)
		hasChange = true
	}
	if d.HasChange("query") {
		log.Printf("[DEBUG]  citrixadc-provider: Query has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["query"] = d.Get("query").(string)
		hasChange = true
	}
	if d.HasChange("quickabortsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Quickabortsize has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["quickabortsize"] = d.Get("quickabortsize").(int)
		hasChange = true
	}
	if d.HasChange("relexpiry") {
		log.Printf("[DEBUG]  citrixadc-provider: Relexpiry has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["relexpiry"] = d.Get("relexpiry").(int)
		hasChange = true
	}
	if d.HasChange("relexpirymillisec") {
		log.Printf("[DEBUG]  citrixadc-provider: Relexpirymillisec has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["relexpirymillisec"] = d.Get("relexpirymillisec").(int)
		hasChange = true
	}
	if d.HasChange("removecookies") {
		log.Printf("[DEBUG]  citrixadc-provider: Removecookies has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["removecookies"] = d.Get("removecookies").(string)
		hasChange = true
	}
	if d.HasChange("selectorvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Selectorvalue has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["selectorvalue"] = d.Get("selectorvalue").(string)
		hasChange = true
	}
	if d.HasChange("tosecondary") {
		log.Printf("[DEBUG]  citrixadc-provider: Tosecondary has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["tosecondary"] = d.Get("tosecondary").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["type"] = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("weaknegrelexpiry") {
		log.Printf("[DEBUG]  citrixadc-provider: Weaknegrelexpiry has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["weaknegrelexpiry"] = d.Get("weaknegrelexpiry").(int)
		hasChange = true
	}
	if d.HasChange("weakposrelexpiry") {
		log.Printf("[DEBUG]  citrixadc-provider: Weakposrelexpiry has changed for cachecontentgroup %s, starting update", cachecontentgroupName)
		cachecontentgroup["weakposrelexpiry"] = d.Get("weakposrelexpiry").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Cachecontentgroup.Type(), &cachecontentgroup)
		if err != nil {
			return fmt.Errorf("Error updating cachecontentgroup %s", cachecontentgroupName)
		}
	}
	return readCachecontentgroupFunc(d, meta)
}

func deleteCachecontentgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCachecontentgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	cachecontentgroupName := d.Id()
	err := client.DeleteResource(service.Cachecontentgroup.Type(), cachecontentgroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
