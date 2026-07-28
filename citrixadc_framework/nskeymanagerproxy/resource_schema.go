package nskeymanagerproxy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NskeymanagerproxyResourceModel describes the resource data model.
type NskeymanagerproxyResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Port       types.Int64  `tfsdk:"port"`
	Serverip   types.String `tfsdk:"serverip"`
	Servername types.String `tfsdk:"servername"`
}

func (r *NskeymanagerproxyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nskeymanagerproxy resource.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Key Manager proxy server port.",
			},
			"serverip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the Key Manager proxy server.",
			},
			"servername": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Fully qualified domain name of the Key Manager proxy server.",
			},
		},
	}
}

func nskeymanagerproxyGetThePayloadFromthePlan(ctx context.Context, data *NskeymanagerproxyResourceModel) ns.Nskeymanagerproxy {
	tflog.Debug(ctx, "In nskeymanagerproxyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nskeymanagerproxy := ns.Nskeymanagerproxy{}
	// nodeid is a cluster GET-context filter, not an add/delete payload field (Pattern 15) - excluded.
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		nskeymanagerproxy.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() {
		nskeymanagerproxy.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		nskeymanagerproxy.Servername = data.Servername.ValueString()
	}

	return nskeymanagerproxy
}

func nskeymanagerproxySetAttrFromGet(ctx context.Context, data *NskeymanagerproxyResourceModel, getResponseData map[string]interface{}) *NskeymanagerproxyResourceModel {
	tflog.Debug(ctx, "In nskeymanagerproxySetAttrFromGet Function")

	// Convert API response to model
	// nodeid is a cluster GET-context filter (Pattern 15), not a returned property -
	// preserve the existing plan/state value instead of nulling it.
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil {
		data.Serverip = types.StringValue(val.(string))
	} else {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}

	// ID is set once in Create (Pattern 6); do not recompute it here.

	return data
}

// nskeymanagerproxySetAttrFromGetForDatasource faithfully copies the GET response
// and sets the ID, since the datasource has no Create step to set it.
func nskeymanagerproxySetAttrFromGetForDatasource(ctx context.Context, data *NskeymanagerproxyResourceModel, getResponseData map[string]interface{}) *NskeymanagerproxyResourceModel {
	tflog.Debug(ctx, "In nskeymanagerproxySetAttrFromGetForDatasource Function")

	nskeymanagerproxySetAttrFromGet(ctx, data, getResponseData)

	// ID is serverip when set, otherwise servername (both x-unique-attr).
	if !data.Serverip.IsNull() && data.Serverip.ValueString() != "" {
		data.Id = types.StringValue(data.Serverip.ValueString())
	} else {
		data.Id = types.StringValue(data.Servername.ValueString())
	}

	return data
}
