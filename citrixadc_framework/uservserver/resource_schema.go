package uservserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/user"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// UservserverResourceModel describes the resource data model.
type UservserverResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Params       types.String `tfsdk:"params"`
	Comment      types.String `tfsdk:"comment"`
	Defaultlb    types.String `tfsdk:"defaultlb"`
	Ipaddress    types.String `tfsdk:"ipaddress"`
	Name         types.String `tfsdk:"name"`
	Port         types.Int64  `tfsdk:"port"`
	State        types.String `tfsdk:"state"`
	Userprotocol types.String `tfsdk:"userprotocol"`
}

func (r *UservserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the uservserver resource.",
			},
			"params": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the protocol.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the virtual server.",
			},
			"defaultlb": schema.StringAttribute{
				Required:    true,
				Description: "Name of the default Load Balancing virtual server used for load balancing of services. The protocol type of default Load Balancing virtual server should be a user type.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"port": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number for the virtual server.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Initial state of the user vserver.",
			},
			"userprotocol": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User protocol uesd by the service.",
			},
		},
	}
}

func uservserverGetThePayloadFromtheConfig(ctx context.Context, data *UservserverResourceModel) user.Uservserver {
	tflog.Debug(ctx, "In uservserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	uservserver := user.Uservserver{}
	if !data.Params.IsNull() {
		uservserver.Params = data.Params.ValueString()
	}
	if !data.Comment.IsNull() {
		uservserver.Comment = data.Comment.ValueString()
	}
	if !data.Defaultlb.IsNull() {
		uservserver.Defaultlb = data.Defaultlb.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		uservserver.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		uservserver.Name = data.Name.ValueString()
	}
	if !data.Port.IsNull() {
		uservserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.State.IsNull() {
		uservserver.State = data.State.ValueString()
	}
	if !data.Userprotocol.IsNull() {
		uservserver.Userprotocol = data.Userprotocol.ValueString()
	}

	return uservserver
}

func uservserverSetAttrFromGet(ctx context.Context, data *UservserverResourceModel, getResponseData map[string]interface{}) *UservserverResourceModel {
	tflog.Debug(ctx, "In uservserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Params"]; ok && val != nil {
		data.Params = types.StringValue(val.(string))
	} else {
		data.Params = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["defaultlb"]; ok && val != nil {
		data.Defaultlb = types.StringValue(val.(string))
	} else {
		data.Defaultlb = types.StringNull()
	}
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
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["userprotocol"]; ok && val != nil {
		data.Userprotocol = types.StringValue(val.(string))
	} else {
		data.Userprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
