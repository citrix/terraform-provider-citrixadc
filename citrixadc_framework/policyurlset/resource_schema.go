package policyurlset

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicyurlsetResourceModel describes the resource data model.
type PolicyurlsetResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Canaryurl           types.String `tfsdk:"canaryurl"`
	Comment             types.String `tfsdk:"comment"`
	Delimiter           types.String `tfsdk:"delimiter"`
	Imported            types.Bool   `tfsdk:"imported"`
	Interval            types.Int64  `tfsdk:"interval"`
	Matchedid           types.Int64  `tfsdk:"matchedid"`
	Name                types.String `tfsdk:"name"`
	Overwrite           types.Bool   `tfsdk:"overwrite"`
	Privateset          types.Bool   `tfsdk:"privateset"`
	Rowseparator        types.String `tfsdk:"rowseparator"`
	Subdomainexactmatch types.Bool   `tfsdk:"subdomainexactmatch"`
	Url                 types.String `tfsdk:"url"`
	UrlWo               types.String `tfsdk:"url_wo"`
	UrlWoVersion        types.Int64  `tfsdk:"url_wo_version"`
}

func (r *PolicyurlsetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policyurlset resource.",
			},
			"canaryurl": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Add this URL to this urlset. Used for testing when contents of urlset is kept confidential.",
			},
			"comment": schema.StringAttribute{
				// Pattern 15: comment belongs only to the secondary plain "add"
				// verb, NOT to the Import action used as Create. Kept Optional for
				// completeness but excluded from the Import payload builder.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this url set.",
			},
			"delimiter": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "CSV file record delimiter.",
			},
			"imported": schema.BoolAttribute{
				// Pattern 15: GET-only filter argument (x-unique-attr in metadata).
				// Excluded from the Import payload, ID composition, delete args and
				// Read match loop. Kept Optional (no Computed) for datasource filter.
				Optional:    true,
				Description: "when set, display shows all imported urlsets.",
			},
			"interval": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The interval, in seconds, rounded down to the nearest 15 minutes, at which the update of urlset occurs.",
			},
			"matchedid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "An ID that would be sent to AppFlow to indicate which URLSet was the last one that matched the requested URL.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Unique name of the url set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrites the existing file.",
			},
			"privateset": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Prevent this urlset from being exported.",
			},
			"rowseparator": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "CSV file row separator.",
			},
			"subdomainexactmatch": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Force exact subdomain matching, ex. given an entry 'google.com' in the urlset, a request to 'news.google.com' won't match, if subdomainExactMatch is set.",
			},
			"url": schema.StringAttribute{
				// x-secret-attr + mandatory (Pattern 17). Part of the url / url_wo /
				// url_wo_version triple; required-ness is enforced at plan time by
				// ValidateConfig (at-least-one-of url / url_wo).
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path and file name) from where the CSV (comma separated file) file will be imported or exported. Each record/line will one entry within the urlset. The first field contains the URL pattern, subsequent fields contains the metadata, if available. HTTP, HTTPS and FTP protocols are supported. NOTE: The operation fails if the destination HTTPS server requires client certificate authentication for access.",
			},
			"url_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path and file name) from where the CSV (comma separated file) file will be imported or exported. Each record/line will one entry within the urlset. The first field contains the URL pattern, subsequent fields contains the metadata, if available. HTTP, HTTPS and FTP protocols are supported. NOTE: The operation fails if the destination HTTPS server requires client certificate authentication for access.",
			},
			"url_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a url_wo update.",
			},
		},
	}
}

func policyurlsetGetThePayloadFromthePlan(ctx context.Context, data *PolicyurlsetResourceModel) policy.Policyurlset {
	tflog.Debug(ctx, "In policyurlsetGetThePayloadFromthePlan Function")

	// Create API request body from the model
	policyurlset := policy.Policyurlset{}
	if !data.Canaryurl.IsNull() && !data.Canaryurl.IsUnknown() {
		policyurlset.Canaryurl = data.Canaryurl.ValueString()
	}
	// Skip comment: belongs only to the plain "add" verb, not the Import action.
	// Skip imported: GET-only filter argument (Pattern 15), not an Import payload field.
	if !data.Delimiter.IsNull() && !data.Delimiter.IsUnknown() {
		policyurlset.Delimiter = data.Delimiter.ValueString()
	}
	if !data.Interval.IsNull() && !data.Interval.IsUnknown() {
		policyurlset.Interval = utils.IntPtr(int(data.Interval.ValueInt64()))
	}
	if !data.Matchedid.IsNull() && !data.Matchedid.IsUnknown() {
		policyurlset.Matchedid = utils.IntPtr(int(data.Matchedid.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policyurlset.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		policyurlset.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Privateset.IsNull() && !data.Privateset.IsUnknown() {
		policyurlset.Privateset = data.Privateset.ValueBool()
	}
	if !data.Rowseparator.IsNull() && !data.Rowseparator.IsUnknown() {
		policyurlset.Rowseparator = data.Rowseparator.ValueString()
	}
	if !data.Subdomainexactmatch.IsNull() && !data.Subdomainexactmatch.IsUnknown() {
		policyurlset.Subdomainexactmatch = data.Subdomainexactmatch.ValueBool()
	}
	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		policyurlset.Url = data.Url.ValueString()
	}
	// Skip write-only attribute: url_wo
	// Skip version tracker attribute: url_wo_version

	return policyurlset
}

func policyurlsetGetThePayloadFromtheConfig(ctx context.Context, data *PolicyurlsetResourceModel, payload *policy.Policyurlset) {
	tflog.Debug(ctx, "In policyurlsetGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: url_wo -> url
	if !data.UrlWo.IsNull() {
		urlWo := data.UrlWo.ValueString()
		if urlWo != "" {
			payload.Url = urlWo
		}
	}
}

// policyurlsetSetAttrFromGet is used by the resource Read. It preserves the
// user-configured plan/state values for RequiresReplace import attributes that
// the import action does not faithfully echo back, while refreshing the fields
// NITRO reliably returns. It does NOT set data.Id (set once in Create).
func policyurlsetSetAttrFromGet(ctx context.Context, data *PolicyurlsetResourceModel, getResponseData map[string]interface{}) *PolicyurlsetResourceModel {
	tflog.Debug(ctx, "In policyurlsetSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["canaryurl"]; ok && val != nil {
		data.Canaryurl = types.StringValue(val.(string))
	}
	// comment is not part of the Import action; preserve state value.
	if val, ok := getResponseData["delimiter"]; ok && val != nil {
		data.Delimiter = types.StringValue(val.(string))
	}
	// imported is a GET-only filter (Pattern 15); preserve state value.
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["matchedid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Matchedid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["overwrite"]; ok && val != nil {
		data.Overwrite = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["privateset"]; ok && val != nil {
		data.Privateset = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["rowseparator"]; ok && val != nil {
		data.Rowseparator = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["subdomainexactmatch"]; ok && val != nil {
		data.Subdomainexactmatch = types.BoolValue(val.(bool))
	}
	// url is not returned by NITRO API (secret/ephemeral) - retain from config
	// url_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// url_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	return data
}

// policyurlsetSetAttrFromGetForDatasource faithfully copies every field returned
// by GET into the model (datasources have no prior plan/state to preserve) and
// sets the datasource ID to the resource name.
func policyurlsetSetAttrFromGetForDatasource(ctx context.Context, data *PolicyurlsetResourceModel, getResponseData map[string]interface{}) *PolicyurlsetResourceModel {
	tflog.Debug(ctx, "In policyurlsetSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["canaryurl"]; ok && val != nil {
		data.Canaryurl = types.StringValue(val.(string))
	} else {
		data.Canaryurl = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["delimiter"]; ok && val != nil {
		data.Delimiter = types.StringValue(val.(string))
	} else {
		data.Delimiter = types.StringNull()
	}
	if val, ok := getResponseData["imported"]; ok && val != nil {
		data.Imported = types.BoolValue(val.(bool))
	} else {
		data.Imported = types.BoolNull()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
		}
	} else {
		data.Interval = types.Int64Null()
	}
	if val, ok := getResponseData["matchedid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Matchedid = types.Int64Value(intVal)
		}
	} else {
		data.Matchedid = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["overwrite"]; ok && val != nil {
		data.Overwrite = types.BoolValue(val.(bool))
	} else {
		data.Overwrite = types.BoolNull()
	}
	if val, ok := getResponseData["privateset"]; ok && val != nil {
		data.Privateset = types.BoolValue(val.(bool))
	} else {
		data.Privateset = types.BoolNull()
	}
	if val, ok := getResponseData["rowseparator"]; ok && val != nil {
		data.Rowseparator = types.StringValue(val.(string))
	} else {
		data.Rowseparator = types.StringNull()
	}
	if val, ok := getResponseData["subdomainexactmatch"]; ok && val != nil {
		data.Subdomainexactmatch = types.BoolValue(val.(bool))
	} else {
		data.Subdomainexactmatch = types.BoolNull()
	}
	// url is a secret; never returned by NITRO.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
