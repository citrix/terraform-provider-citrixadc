package appfwprofile_jsondosurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileJsondosurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alertonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send SNMP alert?",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"jsondosurl": schema.StringAttribute{
				Required:    true,
				Description: "The URL on which we need to enforce the specified JSON denial-of-service (JSONDoS) attack protections.\nAn JSON DoS configuration consists of the following items:\n* URL. PCRE-format regular expression for the URL.\n* Maximum-document-length-check toggle.  ON to enable this check, OFF to disable it.\n* Maximum document length. Positive integer representing the maximum length of the JSON document.\n* Maximum-container-depth-check toggle. ON to enable, OFF to disable.\n * Maximum container depth. Positive integer representing the maximum container depth of the JSON document.\n* Maximum-object-key-count-check toggle. ON to enable, OFF to disable.\n* Maximum object key count. Positive integer representing the maximum allowed number of keys in any of the  JSON object.\n* Maximum-object-key-length-check toggle. ON to enable, OFF to disable.\n* Maximum object key length. Positive integer representing the maximum allowed length of key in any of the  JSON object.\n* Maximum-array-value-count-check toggle. ON to enable, OFF to disable.\n* Maximum array value count. Positive integer representing the maximum allowed number of values in any of the JSON array.\n* Maximum-string-length-check toggle. ON to enable, OFF to disable.\n* Maximum string length. Positive integer representing the maximum length of string in JSON.",
			},
			"jsonmaxarraylength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum array length in the any of JSON object. This check protects against arrays having large lengths.",
			},
			"jsonmaxarraylengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max array value count check is ON or OFF.",
			},
			"jsonmaxcontainerdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum allowed nesting depth  of JSON document. JSON allows one to nest the containers (object and array) in any order to any depth. This check protects against documents that have excessive depth of hierarchy.",
			},
			"jsonmaxcontainerdepthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max depth check is ON or OFF.",
			},
			"jsonmaxdocumentlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum document length of JSON document, in bytes.",
			},
			"jsonmaxdocumentlengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max document length check is ON or OFF.",
			},
			"jsonmaxobjectkeycount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum key count in the any of JSON object. This check protects against objects that have large number of keys.",
			},
			"jsonmaxobjectkeycountcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max object key count check is ON or OFF.",
			},
			"jsonmaxobjectkeylength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum key length in the any of JSON object. This check protects against objects that have large keys.",
			},
			"jsonmaxobjectkeylengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max object key length check is ON or OFF.",
			},
			"jsonmaxstringlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum string length in the JSON. This check protects against strings that have large length.",
			},
			"jsonmaxstringlengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if JSON Max string value count check is ON or OFF.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}
