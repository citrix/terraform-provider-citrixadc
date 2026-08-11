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
			// SDK v2: Optional. NITRO struct field "Params" (json:"Params"). Updateable in-place.
			"params": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Default added so removing it from config produces a plan diff,
				// letting Update fire the NITRO unset (spec has no documented
				// default; unset reverts it to empty).
				Default:     stringdefault.StaticString(""),
				Description: "Any comments associated with the protocol.",
			},
			// SDK v2: Optional+Computed, updateable.
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Default added so removing it from config produces a plan diff,
				// letting Update fire the NITRO unset (spec has no documented
				// default; unset reverts it to empty).
				Default:     stringdefault.StaticString(""),
				Description: "Any comments that you might want to associate with the virtual server.",
			},
			// SDK v2: Required, updateable (no ForceNew).
			"defaultlb": schema.StringAttribute{
				Required:    true,
				Description: "Name of the default Load Balancing virtual server used for load balancing of services. The protocol type of default Load Balancing virtual server should be a user type.",
			},
			// SDK v2: Required, updateable (no ForceNew).
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			"port": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number for the virtual server.",
			},
			// SDK v2: Optional+Computed, NO ForceNew, NO Default. Updateable via enable/disable action.
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the user vserver.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
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

// uservserverGetThePayloadFromthePlan builds the full create payload (AddResource).
// Mirrors SDK v2 createUservserverFunc which includes state in the add payload.
func uservserverGetThePayloadFromthePlan(ctx context.Context, data *UservserverResourceModel) user.Uservserver {
	tflog.Debug(ctx, "In uservserverGetThePayloadFromthePlan Function")

	uservserver := user.Uservserver{}
	if !data.Params.IsNull() && !data.Params.IsUnknown() {
		uservserver.Params = data.Params.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		uservserver.Comment = data.Comment.ValueString()
	}
	if !data.Defaultlb.IsNull() && !data.Defaultlb.IsUnknown() {
		uservserver.Defaultlb = data.Defaultlb.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		uservserver.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		uservserver.Name = data.Name.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		uservserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		uservserver.State = data.State.ValueString()
	}
	if !data.Userprotocol.IsNull() && !data.Userprotocol.IsUnknown() {
		uservserver.Userprotocol = data.Userprotocol.ValueString()
	}

	return uservserver
}

// uservserverGetTheUpdatablePayloadFromThePlan builds the in-place update payload
// (UpdateResource). Only the updateable attributes from SDK v2 updateUservserverFunc
// are included: comment, defaultlb, ipaddress, params (plus name as the identity key).
// state is handled separately via the enable/disable action, and
// userprotocol/port are ForceNew (they never reach Update).
func uservserverGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *UservserverResourceModel) user.Uservserver {
	tflog.Debug(ctx, "In uservserverGetTheUpdatablePayloadFromThePlan Function")

	uservserver := user.Uservserver{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		uservserver.Name = data.Name.ValueString()
	}
	if !data.Params.IsNull() && !data.Params.IsUnknown() {
		uservserver.Params = data.Params.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		uservserver.Comment = data.Comment.ValueString()
	}
	if !data.Defaultlb.IsNull() && !data.Defaultlb.IsUnknown() {
		uservserver.Defaultlb = data.Defaultlb.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		uservserver.Ipaddress = data.Ipaddress.ValueString()
	}

	return uservserver
}

// uservserverSetAttrFromGet is the RESOURCE state setter. It preserves a known
// configured value when NITRO omits the field from GET (omit-on-default guard),
// only nulling attributes that are still unknown.
func uservserverSetAttrFromGet(ctx context.Context, data *UservserverResourceModel, getResponseData map[string]interface{}) *UservserverResourceModel {
	tflog.Debug(ctx, "In uservserverSetAttrFromGet Function")

	// "Params" is the NITRO json key (capitalised); fall back to "params" defensively.
	if val, ok := getResponseData["Params"]; ok && val != nil {
		data.Params = types.StringValue(val.(string))
	} else if val, ok := getResponseData["params"]; ok && val != nil {
		data.Params = types.StringValue(val.(string))
	} else if data.Params.IsUnknown() {
		data.Params = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["defaultlb"]; ok && val != nil {
		data.Defaultlb = types.StringValue(val.(string))
	} else if data.Defaultlb.IsUnknown() {
		data.Defaultlb = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else if data.Ipaddress.IsUnknown() {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["userprotocol"]; ok && val != nil {
		data.Userprotocol = types.StringValue(val.(string))
	} else if data.Userprotocol.IsUnknown() {
		data.Userprotocol = types.StringNull()
	}

	// Set ID for the resource - single unique attribute (name). Matches SDK v2 d.SetId(name).
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// uservserverSetAttrFromGetForDatasource is the DATASOURCE state setter. It copies
// every attribute from the GET response (nulling omitted fields) and sets the ID.
func uservserverSetAttrFromGetForDatasource(ctx context.Context, data *UservserverResourceModel, getResponseData map[string]interface{}) *UservserverResourceModel {
	tflog.Debug(ctx, "In uservserverSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["Params"]; ok && val != nil {
		data.Params = types.StringValue(val.(string))
	} else if val, ok := getResponseData["params"]; ok && val != nil {
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
		} else {
			data.Port = types.Int64Null()
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

	// Set ID for the datasource - single unique attribute (name).
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
