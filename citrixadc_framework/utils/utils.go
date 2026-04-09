package utils

import (
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"log"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NetScalerNitroClient struct {
	Username string
	Password string
	Endpoint string
	client   *service.NitroClient
	lock     sync.Mutex
}

func IsTargetAdcCluster(nsClient *service.NitroClient) bool {
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

func ToStringList(in []interface{}) []string {
	out := make([]string, 0, len(in))
	for _, val := range in {
		out = append(out, val.(string))
	}
	return out
}

func ToIntegerList(in []interface{}) []int {
	out := make([]int, len(in))
	for i := range in {
		out[i] = in[i].(int)
	}
	return out
}

func StringListToIntList(in []interface{}) []int {
	out := make([]int, 0, len(in))
	for _, val := range in {
		res, _ := strconv.Atoi(val.(string))
		out = append(out, res)
	}
	return out
}

// Check if the attribute is int, if not convert to int and set the value
func SetToInt(attributeName string, d *schema.ResourceData, value interface{}) {
	log.Printf("[DEBUG] netscaler-provider: In setToInt for attribute %s", attributeName)

	var v int
	var err error

	switch valueTyped := value.(type) {
	case int:
		v = valueTyped
	case float64:
		v = int(valueTyped)
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
func ToString(value interface{}) string {
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
func RebootNetScaler(d *schema.ResourceData, meta interface{}, warm bool) error {
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
func HashString(input string) int {
	h := fnv.New32a()
	h.Write([]byte(input))
	return int(h.Sum32() & 0x7fffffff) // Ensure positive int
}

// intPtr returns a pointer to the provided int value
// This is useful for optional fields in structs that require *int
func IntPtr(i int) *int {
	return &i
}

// boolPtr returns a pointer to the provided bool value
// This is useful for optional fields in structs that require *bool
func BoolPtr(b bool) *bool {
	return &b
}

// Helper function to convert interface{} to int64
func ConvertToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", value)
	}
}

// EncodeToBase64 encodes a string to base64
func EncodeToBase64(input interface{}) string {
	stringified := ToString(input)
	return base64.StdEncoding.EncodeToString([]byte(stringified))
}

// UrlEncode URL-encodes a value so it is safe to embed in a comma-separated ID string.
// Regular ADC resource names (alphanumeric, hyphens, underscores) pass through unchanged;
// only characters that would break parsing (e.g. commas, colons) are percent-encoded.
func UrlEncode(input interface{}) string {
	return url.QueryEscape(ToString(input))
}

// UrlDecode decodes a URL-encoded string back to its original value.
// Returns the decoded string and an error if the input is malformed.
func UrlDecode(encoded string) (string, error) {
	return url.QueryUnescape(encoded)
}

// isValidIdAttrName returns true if s is a valid NITRO attribute name
// (lowercase letters, digits, underscores only).
func isValidIdAttrName(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_') {
			return false
		}
	}
	return true
}

// isNewIdFormat returns true when every comma-separated segment looks like
// "attrname:value", which is the Framework multi-part ID format.
func isNewIdFormat(parts []string) bool {
	for _, part := range parts {
		colonIdx := strings.Index(part, ":")
		if colonIdx < 0 {
			return false
		}
		if !isValidIdAttrName(part[:colonIdx]) {
			return false
		}
	}
	return true
}

// ParseIdString parses a resource ID string into a map of attribute name → value,
// plus a set of optional attributes that were absent from a legacy-format ID.
//
// It handles two formats:
//   - New Framework format: "attr1:urlEncoded(val1),attr2:urlEncoded(val2),..."
//     (self-describing key:value pairs, values are URL-decoded)
//   - Legacy SDK v2 format: "val1,val2,..."
//     (positional values mapped to names via legacyAttrOrder)
//
// legacyAttrOrder lists attribute names in the SDK v2 d.SetId() order
// (sourced from resource_id_mapping.json, '?' suffixes already stripped).
//
// legacyOptionalAttrs lists attribute names that were marked optional ('?') in
// resource_id_mapping.json. When a legacy ID omits trailing optional values,
// those attrs are returned in the optionalAbsent map so callers can apply
// zero-value matching rather than skipping the comparison entirely.
func ParseIdString(idStr string, legacyAttrOrder []string, legacyOptionalAttrs []string) (map[string]string, map[string]bool, error) {
	if idStr == "" {
		return map[string]string{}, map[string]bool{}, nil
	}

	parts := strings.Split(idStr, ",")

	optionalSet := make(map[string]bool, len(legacyOptionalAttrs))
	for _, attr := range legacyOptionalAttrs {
		optionalSet[attr] = true
	}

	var result map[string]string

	if isNewIdFormat(parts) {
		result = make(map[string]string, len(parts))
		for _, part := range parts {
			colonIdx := strings.Index(part, ":")
			key := part[:colonIdx]
			val, err := url.QueryUnescape(part[colonIdx+1:])
			if err != nil {
				return nil, nil, fmt.Errorf("failed to URL-decode value for key %q: %w", key, err)
			}
			result[key] = val
		}
	} else {
		// Legacy positional format
		if len(legacyAttrOrder) == 0 {
			return nil, nil, fmt.Errorf("cannot parse legacy ID %q: no attribute order provided", idStr)
		}
		result = make(map[string]string, len(parts))
		for i, val := range parts {
			if i >= len(legacyAttrOrder) {
				break
			}
			result[legacyAttrOrder[i]] = val
		}
	}

	// Any optional attr absent from the parsed result (regardless of format)
	optionalAbsent := make(map[string]bool)
	for attr := range optionalSet {
		if _, found := result[attr]; !found {
			optionalAbsent[attr] = true
		}
	}

	return result, optionalAbsent, nil
}
