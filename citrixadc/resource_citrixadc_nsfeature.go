package citrixadc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsfeature() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsfeatureFunc,
		ReadContext:   readNsfeatureFunc,
		UpdateContext: updateNsfeatureFunc,
		DeleteContext: deleteNsfeatureFunc,
		Schema: map[string]*schema.Schema{
			"wl": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"lb": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cs": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cr": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cmp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"pq": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ssl": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"gslb": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"hdosp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cf": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sslvpn": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"aaa": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ospf": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"bgp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rewrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipv6pt": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"appfw": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"responder": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"htmlinjection": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"push": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"appflow": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cloudbridge": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"isis": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ch": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"appqoe": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"contentaccelerator": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rise": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"feo": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"lsn": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rdpproxy": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"rep": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"urlfiltering": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"videooptimization": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"forwardproxy": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sslinterception": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"adaptivetcp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cqa": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ci": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"bot": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"apigateway": {
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
	"apigateway",
}

func createNsfeatureFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	if err := disableNsFeatureList(meta, disableList); err != nil {
		return diag.FromErr(err)
	}

	err := readNsfeatureFunc(ctx, d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsfeature but we can't read it ?? %s", nsfeatureId)
		return nil
	}
	return nil
}

func readNsfeatureFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func updateNsfeatureFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	if err := disableNsFeatureList(meta, disableList); err != nil {
		return diag.FromErr(err)
	}

	return readNsfeatureFunc(ctx, d, meta)
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

func deleteNsfeatureFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsfeatureFunc")

	d.SetId("")

	return nil
}
