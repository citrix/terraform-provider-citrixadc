package endpointinfo

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// EndpointinfoResourceModel describes the resource data model.
type EndpointinfoResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Endpointkind       types.String `tfsdk:"endpointkind"`
	Endpointlabelsjson types.String `tfsdk:"endpointlabelsjson"`
	Endpointmetadata   types.String `tfsdk:"endpointmetadata"`
	Endpointname       types.String `tfsdk:"endpointname"`
}

func (r *EndpointinfoResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the endpointinfo resource.",
			},
			"endpointkind": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("IP"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Endpoint kind. Currently, IP endpoints are supported",
			},
			"endpointlabelsjson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String representing labels in json form. Maximum length 16K",
			},
			"endpointmetadata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String of qualifiers, in dotted notation, structured metadata for an endpoint. Each qualifier is more specific than the one that precedes it, as in cluster.namespace.service. For example: cluster.default.frontend. \nNote: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.",
			},
			"endpointname": schema.StringAttribute{
				Required:    true,
				Description: "Name of endpoint, depends on kind. For IP Endpoint - IP address.",
			},
		},
	}
}

func endpointinfoGetThePayloadFromthePlan(ctx context.Context, data *EndpointinfoResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In endpointinfoGetThePayloadFromthePlan Function")

	// No vendored NITRO struct exists for endpointinfo (Pattern 3: missing
	// vendored struct -> use a map payload). AddResource / UpdateResource
	// accept any JSON-marshalable interface.
	endpointinfo := make(map[string]interface{})
	if !data.Endpointkind.IsNull() && !data.Endpointkind.IsUnknown() {
		endpointinfo["endpointkind"] = data.Endpointkind.ValueString()
	}
	if !data.Endpointlabelsjson.IsNull() && !data.Endpointlabelsjson.IsUnknown() {
		endpointinfo["endpointlabelsjson"] = data.Endpointlabelsjson.ValueString()
	}
	if !data.Endpointmetadata.IsNull() && !data.Endpointmetadata.IsUnknown() {
		endpointinfo["endpointmetadata"] = data.Endpointmetadata.ValueString()
	}
	if !data.Endpointname.IsNull() && !data.Endpointname.IsUnknown() {
		endpointinfo["endpointname"] = data.Endpointname.ValueString()
	}

	return endpointinfo
}

func endpointinfoSetAttrFromGet(ctx context.Context, data *EndpointinfoResourceModel, getResponseData map[string]interface{}) *EndpointinfoResourceModel {
	tflog.Debug(ctx, "In endpointinfoSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["endpointkind"]; ok && val != nil {
		data.Endpointkind = types.StringValue(val.(string))
	} else {
		data.Endpointkind = types.StringNull()
	}
	if val, ok := getResponseData["endpointlabelsjson"]; ok && val != nil {
		data.Endpointlabelsjson = types.StringValue(val.(string))
	} else {
		data.Endpointlabelsjson = types.StringNull()
	}
	if val, ok := getResponseData["endpointmetadata"]; ok && val != nil {
		data.Endpointmetadata = types.StringValue(val.(string))
	} else {
		data.Endpointmetadata = types.StringNull()
	}
	if val, ok := getResponseData["endpointname"]; ok && val != nil {
		data.Endpointname = types.StringValue(val.(string))
	} else {
		data.Endpointname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("endpointkind:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Endpointkind.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("endpointname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Endpointname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
