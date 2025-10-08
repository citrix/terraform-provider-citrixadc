package citrixadc

import (
	"hash/fnv"
	"log"
	"strconv"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func isTargetAdcCluster(nsClient *service.NitroClient) bool {
	log.Printf("[DEBUG]  citrixadc-provider-test: In isTargetAdcCluster")

	datalist, err := nsClient.FindAllResources(service.Clusterinstance.Type())
	if err != nil {
		//lintignore:R009
		panic(err)
	}

	if len(datalist) == 0 {
		return false
	} else {
		return true
	}
}

func toStringList(in []interface{}) []string {
	out := make([]string, 0, len(in))
	for _, val := range in {
		out = append(out, val.(string))
	}
	return out
}

func toIntegerList(in []interface{}) []int {
	out := make([]int, len(in))
	for i := range in {
		out[i] = in[i].(int)
	}
	return out
}

func stringListToIntList(in []interface{}) []int {
	out := make([]int, 0, len(in))
	for _, val := range in {
		res, _ := strconv.Atoi(val.(string))
		out = append(out, res)
	}
	return out
}

// Check if the attribute is int, if not convert to int and set the value
func setToInt(attributeName string, d *schema.ResourceData, value interface{}) {
	log.Printf("[DEBUG] netscaler-provider: In setToInt for attribute %s", attributeName)

	var v int
	var err error

	switch valueTyped := value.(type) {
	case int:
		v = valueTyped
	case string:
		v, err = strconv.Atoi(valueTyped)
	case nil:
		return
	default:
		log.Printf("[DEBUG] got unexpected type %T for int", value)
		return
	}

	if err != nil {
		return
	}

	d.Set(attributeName, v)
}

// Convert various types (int, float64, etc.) to string
func toString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case bool:
		return strconv.FormatBool(v)
	case nil:
		return ""
	default:
		return ""
	}
}

// TODO: Better implementation can be done to check whether NetScaler is UP after reboot, by polling the NetScaler
func rebootNetScaler(d *schema.ResourceData, meta interface{}, warm bool) error {
	log.Printf("[DEBUG] netscaler-provider: In rebootAdc")

	client := meta.(*NetScalerNitroClient).client
	reboot := ns.Reboot{
		Warm: warm,
	}
	if err := client.ActOnResource("reboot", &reboot, ""); err != nil {
		return err
	}
	// wait for NetScaler to Reboot. If warm reboot then wait for 120s, else wait for 240s
	sleepTimeout := 120
	if !warm {
		sleepTimeout = 240
	}
	time.Sleep(time.Second * time.Duration(sleepTimeout))

	return nil
}

// hashString returns a hash of the input string using FNV-1a
// This replaces the hashcode.String function from SDK v1
func hashString(input string) int {
	h := fnv.New32a()
	h.Write([]byte(input))
	return int(h.Sum32() & 0x7fffffff) // Ensure positive int
}
