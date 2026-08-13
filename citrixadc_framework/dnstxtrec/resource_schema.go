package dnstxtrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnstxtrecResourceModel describes the resource data model.
type DnstxtrecResourceModel struct {
	Id        types.String `tfsdk:"id"`
	String    types.List   `tfsdk:"string"`
	Domain    types.String `tfsdk:"domain"`
	Ecssubnet types.String `tfsdk:"ecssubnet"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Recordid  types.Int64  `tfsdk:"recordid"`
	Ttl       types.Int64  `tfsdk:"ttl"`
}

func (r *DnstxtrecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnstxtrec resource.",
			},
			"string": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Information to store in the TXT resource record. Enclose the string in single or double quotation marks. A TXT resource record can contain up to six strings, each of which can contain up to 255 characters. If you want to add a string of more than 255 characters, evaluate whether splitting it into two or more smaller strings, subject to the six-string limit, works for you.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the domain for the TXT record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: preserve computed value on upgrade; replace only when configured.
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Subnet for which the cached TXT record need to be removed.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436: preserve computed value on upgrade; replace only when configured.
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"recordid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436: preserve computed value on upgrade; replace only when configured.
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Unique, internally generated record ID. View the details of the TXT record to obtain its record ID. Mutually exclusive with the string parameter.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436: preserve computed value on upgrade; replace only when configured.
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}

func dnstxtrecGetThePayloadFromthePlan(ctx context.Context, data *DnstxtrecResourceModel, diags *diag.Diagnostics) dns.Dnstxtrec {
	tflog.Debug(ctx, "In dnstxtrecGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// dnstxtrec has only an "add" NITRO verb (no update); the create payload mirrors
	// the SDK v2 resource, which sent only domain, String and ttl. recordid is
	// server-generated (mutually exclusive with String), and ecssubnet/nodeid are
	// read/delete-time arguments, so none of them are part of the add payload.
	dnstxtrec := dns.Dnstxtrec{}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		dnstxtrec.Domain = data.Domain.ValueString()
	}
	if !data.String.IsNull() && !data.String.IsUnknown() {
		stringList := make([]string, 0, len(data.String.Elements()))
		diags.Append(data.String.ElementsAs(ctx, &stringList, false)...)
		dnstxtrec.String = stringList
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		dnstxtrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return dnstxtrec
}

// dnstxtrecSetAttrFromGet populates the resource state from a NITRO GET response.
// The configured "string" (RequiresReplace) is intentionally not overwritten from
// the GET response so the resource preserves the user's configured value and avoids
// "inconsistent result after apply" churn; the datasource setter handles it instead.
func dnstxtrecSetAttrFromGet(ctx context.Context, data *DnstxtrecResourceModel, getResponseData map[string]interface{}) *DnstxtrecResourceModel {
	tflog.Debug(ctx, "In dnstxtrecSetAttrFromGet Function")

	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else if data.Ecssubnet.IsUnknown() {
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["recordid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Recordid = types.Int64Value(intVal)
		}
	} else if data.Recordid.IsUnknown() {
		data.Recordid = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else if data.Ttl.IsUnknown() {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the resource - single unique attribute (domain), plain value.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Domain.ValueString()))

	return data
}

// dnstxtrecSetAttrFromGetForDatasource populates all readable attributes (including
// "string") from a NITRO GET response for the datasource.
func dnstxtrecSetAttrFromGetForDatasource(ctx context.Context, data *DnstxtrecResourceModel, getResponseData map[string]interface{}) *DnstxtrecResourceModel {
	tflog.Debug(ctx, "In dnstxtrecSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["String"]; ok && val != nil {
		if listVal, listOk := val.([]interface{}); listOk {
			strs := make([]string, 0, len(listVal))
			for _, item := range listVal {
				strs = append(strs, fmt.Sprintf("%v", item))
			}
			listValue, d := types.ListValueFrom(ctx, types.StringType, strs)
			if !d.HasError() {
				data.String = listValue
			}
		}
	} else {
		data.String = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else {
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["recordid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Recordid = types.Int64Value(intVal)
		}
	} else {
		data.Recordid = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the datasource - single unique attribute (domain), plain value.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Domain.ValueString()))

	return data
}
