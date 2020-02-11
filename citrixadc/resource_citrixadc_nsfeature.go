package citrixadc

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"strings"
)

func resourceCitrixAdcNsfeature() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsfeatureFunc,
		Read:          readNsfeatureFunc,
		Update:        updateNsfeatureFunc,
		Delete:        deleteNsfeatureFunc,
		Schema: map[string]*schema.Schema{
			"wl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"lb": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cr": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cmp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"pq": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"gslb": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"hdosp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ic": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"aaa": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ospf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"bgp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rewrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipv6pt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"appfw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"responder": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"htmlinjection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"push": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"appflow": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cloudbridge": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"isis": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ch": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"appqoe": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"contentaccelerator": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rise": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"feo": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"lsn": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rdpproxy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rep": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"urlfiltering": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"videooptimization": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"forwardproxy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sslinterception": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"adaptivetcp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cqa": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ci": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"bot": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

var featureList = [...]string{
	"wl",
	"sp",
	"lb",
	"cs",
	"cr",
	"sc",
	"cmp",
	"pq",
	"ssl",
	"gslb",
	"hdosp",
	"cf",
	"ic",
	"sslvpn",
	"aaa",
	"ospf",
	"rip",
	"bgp",
	"rewrite",
	"ipv6pt",
	"appfw",
	"responder",
	"htmlinjection",
	"push",
	"appflow",
	"cloudbridge",
	"isis",
	"ch",
	"appqoe",
	"contentaccelerator",
	"rise",
	"feo",
	"lsn",
	"rdpproxy",
	"rep",
	"urlfiltering",
	"videooptimization",
	"forwardproxy",
	"sslinterception",
	"adaptivetcp",
	"cqa",
	"ci",
	"bot",
}

func createNsfeatureFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsfeatureFunc")
	client := meta.(*NetScalerNitroClient).client
	_ = client
	var nsfeatureId string
	nsfeatureId = resource.PrefixedUniqueId("tf-nsfeature-")

	d.SetId(nsfeatureId)
	enableList := make([]string, 0, len(featureList))
	disableList := make([]string, 0, len(featureList))

	for _, featureName := range featureList {
		if val, ok := d.GetOkExists(featureName); ok {
			if val.(bool) {
				enableList = append(enableList, featureName)
			} else {
				disableList = append(disableList, featureName)
			}
		}
	}

	if err := enableNsFeatureList(meta, enableList); err != nil {
		return err
	}

	if err := disableNsFeatureList(meta, disableList); err != nil {
		return err
	}

	err := readNsfeatureFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsfeature but we can't read it ?? %s", nsfeatureId)
		return nil
	}
	return nil
}

func readNsfeatureFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsfeatureFunc")
	client := meta.(*NetScalerNitroClient).client
	nsfeatureId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsfeature state %s", nsfeatureId)
	data, err := client.ListEnabledFeatures()
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsfeature state %s", nsfeatureId)
		d.SetId("")
		return nil
	}
	featuresRead := make([]string, len(data))
	for i, val := range data {
		featuresRead[i] = strings.ToLower(val)
	}
	log.Printf("features enabled %v\n", featuresRead)

	for _, featureName := range featureList {

		found := false
		for _, featureRead := range featuresRead {
			if featureRead == featureName {
				found = true
				break
			}
		}
		if found {
			d.Set(featureName, true)
		} else {
			d.Set(featureName, false)
		}
	}

	return nil

}

func updateNsfeatureFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsfeatureFunc")
	//client := meta.(*NetScalerNitroClient).client

	enableList := make([]string, 0, len(featureList))
	disableList := make([]string, 0, len(featureList))

	for _, featureName := range featureList {
		if d.HasChange(featureName) {
			val := d.Get(featureName)
			log.Printf("[DEBUG]  citrixadc-provider: Feature %v has value %v", featureName, val.(bool))
			if val.(bool) {
				enableList = append(enableList, featureName)
			} else {
				disableList = append(disableList, featureName)
			}
		}
	}

	if err := enableNsFeatureList(meta, enableList); err != nil {
		return err
	}

	if err := disableNsFeatureList(meta, disableList); err != nil {
		return err
	}

	return readNsfeatureFunc(d, meta)
}

func enableNsFeatureList(meta interface{}, featureList []string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In enableNsFeatureList")
	client := meta.(*NetScalerNitroClient).client
	if len(featureList) == 0 {
		log.Printf("")
		return nil
	}
	log.Printf("[DEBUG]  citrixadc-provider: Enabling features %v", featureList)
	if err := client.EnableFeatures(featureList); err != nil {
		return err
	}
	return nil
}

func disableNsFeatureList(meta interface{}, featureList []string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In disableNsFeatureList")
	client := meta.(*NetScalerNitroClient).client
	if len(featureList) == 0 {
		return nil
	}
	log.Printf("[DEBUG]  citrixadc-provider: Disabling features %v", featureList)
	if err := client.DisableFeatures(featureList); err != nil {
		return err
	}
	return nil
}

func deleteNsfeatureFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsfeatureFunc")

	d.SetId("")

	return nil
}
