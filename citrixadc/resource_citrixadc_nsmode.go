package citrixadc

import (
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcNsmode() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsmodeFunc,
		Read:          readNsmodeFunc,
		Delete:        deleteNsmodeFunc,
		Schema: map[string]*schema.Schema{
			"fr": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"l2": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"usip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cka": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpb": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mbf": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"edge": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"usnip": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"l3": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pmtud": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mediaclassification": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sradv": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dradv": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"iradv": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sradv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dradv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bridgebpdus": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ulfd": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

var modesList = [...]string{
	"fr",
	"l2",
	"usip",
	"cka",
	"tcpb",
	"mbf",
	"edge",
	"usnip",
	"l3",
	"pmtud",
	"mediaclassification",
	"sradv",
	"dradv",
	"iradv",
	"sradv6",
	"dradv6",
	"bridgebpdus",
	"ulfd",
}

func createNsmodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsmodeFunc")
	nsmodeName := resource.PrefixedUniqueId("tf-nsmode-")

	err := syncNsmode(d, meta)
	if err != nil {
		return err
	}

	d.SetId(nsmodeName)

	err = readNsmodeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsmode but we can't read it ?? %s", nsmodeName)
		return nil
	}
	return nil
}

func readNsmodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsmodeFunc")
	client := meta.(*NetScalerNitroClient).client
	nsmodeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsmode state %s", nsmodeName)
	findParams := service.FindParams{
		ResourceType: "nsmode",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Error during read %s", err)
		log.Printf("[WARN] citrixadc-provider: Clearing nsmode state %s", nsmodeName)
		d.SetId("")
		return nil
	}
	if len(dataArr) != 1 {
		return fmt.Errorf("Unexpected fetched nsmode result %v", dataArr)
	}
	data := dataArr[0]

	for _, mode := range modesList {
		if val, ok := data[mode]; ok {
			if val.(bool) {
				d.Set(mode, true)
			} else {
				d.Set(mode, false)
			}
		}
	}

	return nil

}

func deleteNsmodeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsmodeFunc")

	d.SetId("")

	return nil
}

func syncNsmode(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In syncNsmode")
	enableList := make([]string, 0, len(modesList))
	disableList := make([]string, 0, len(modesList))

	for _, modeName := range modesList {
		if val, ok := d.GetOkExists(modeName); ok {
			if val.(bool) {
				enableList = append(enableList, modeName)
			} else {
				disableList = append(disableList, modeName)
			}
		}
	}

	if len(enableList) > 0 {
		if err := enableNsModeList(meta, enableList); err != nil {
			return err
		}
	}

	if len(disableList) > 0 {
		if err := disableNsModeList(meta, disableList); err != nil {
			return err
		}
	}

	return nil
}

type nsfeature struct {
	Mode []string `json:"mode"`
}

func enableNsModeList(meta interface{}, enableList []string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In enableNsModeList")
	client := meta.(*NetScalerNitroClient).client
	features := nsfeature{
		Mode: enableList,
	}
	err := client.ActOnResource("nsmode", &features, "enable")
	if err != nil {
		return err
	}
	return nil
}

func disableNsModeList(meta interface{}, disableList []string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In disableNsModeList")
	client := meta.(*NetScalerNitroClient).client
	features := nsfeature{
		Mode: disableList,
	}
	err := client.ActOnResource("nsmode", &features, "disable")
	if err != nil {
		return err
	}
	return nil
}
