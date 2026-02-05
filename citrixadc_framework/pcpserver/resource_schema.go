package pcpserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/pcp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PcpserverResourceModel describes the resource data model.
type PcpserverResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Ipaddress  types.String `tfsdk:"ipaddress"`
	Name       types.String `tfsdk:"name"`
	Pcpprofile types.String `tfsdk:"pcpprofile"`
	Port       types.Int64  `tfsdk:"port"`
}

func (r *PcpserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the pcpserver resource.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the PCP server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the PCP server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my pcpServer\" or my pcpServer).",
			},
			"pcpprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "pcp profile name",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5351),
				Description: "Port number for the PCP server.",
			},
		},
	}
}

func pcpserverGetThePayloadFromtheConfig(ctx context.Context, data *PcpserverResourceModel) pcp.Pcpserver {
	tflog.Debug(ctx, "In pcpserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	pcpserver := pcp.Pcpserver{}
	if !data.Ipaddress.IsNull() {
		pcpserver.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		pcpserver.Name = data.Name.ValueString()
	}
	if !data.Pcpprofile.IsNull() {
		pcpserver.Pcpprofile = data.Pcpprofile.ValueString()
	}
	if !data.Port.IsNull() {
		pcpserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}

	return pcpserver
}

func pcpserverSetAttrFromGet(ctx context.Context, data *PcpserverResourceModel, getResponseData map[string]interface{}) *PcpserverResourceModel {
	tflog.Debug(ctx, "In pcpserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["pcpprofile"]; ok && val != nil {
		data.Pcpprofile = types.StringValue(val.(string))
	} else {
		data.Pcpprofile = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
