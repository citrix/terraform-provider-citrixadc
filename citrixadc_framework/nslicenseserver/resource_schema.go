package nslicenseserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslicenseserverResourceModel describes the resource data model.
type NslicenseserverResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Deviceprofilename types.String `tfsdk:"deviceprofilename"`
	Forceupdateip     types.Bool   `tfsdk:"forceupdateip"`
	Licensemode       types.String `tfsdk:"licensemode"`
	Licenseserverip   types.String `tfsdk:"licenseserverip"`
	Nodeid            types.Int64  `tfsdk:"nodeid"`
	Password          types.String `tfsdk:"password"`
	Port              types.Int64  `tfsdk:"port"`
	Servername        types.String `tfsdk:"servername"`
	Username          types.String `tfsdk:"username"`
}

func (r *NslicenseserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslicenseserver resource.",
			},
			// SDK v2: Optional+Computed+ForceNew
			"deviceprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Device profile is created on ADM and contains the user name and password of the instance(s). ADM will use this info to add the NS for registration",
			},
			// SDK v2: Optional+Computed+ForceNew
			"forceupdateip": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.",
			},
			// SDK v2: Optional+Computed (updateable)
			"licensemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This paramter indicates type of license customer interested while configuring add/set licenseserver",
			},
			// SDK v2: Optional+Computed+ForceNew
			"licenseserverip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "IP address of the License server.",
			},
			// SDK v2: Optional+Computed+ForceNew
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			// SDK v2: Optional+Computed+ForceNew (secret)
			"password": schema.StringAttribute{
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Password to use when authenticating with ADM Agent for LAS licensing.",
			},
			// SDK v2: Optional+Computed (updateable)
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "License server port.",
			},
			// SDK v2: Required+ForceNew
			"servername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Fully qualified domain name of the License server.",
			},
			// SDK v2: Optional+Computed+ForceNew
			"username": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
		},
	}
}

// nslicenseserverGetThePayloadFromthePlan builds the NITRO payload from the plan.
// Attributes that are unknown (Optional+Computed and not configured) are omitted so
// the ADC is not sent spurious zero values.
func nslicenseserverGetThePayloadFromthePlan(ctx context.Context, data *NslicenseserverResourceModel) ns.Nslicenseserver {
	tflog.Debug(ctx, "In nslicenseserverGetThePayloadFromthePlan Function")

	nslicenseserver := ns.Nslicenseserver{}
	if !data.Deviceprofilename.IsNull() && !data.Deviceprofilename.IsUnknown() {
		nslicenseserver.Deviceprofilename = data.Deviceprofilename.ValueString()
	}
	if !data.Forceupdateip.IsNull() && !data.Forceupdateip.IsUnknown() {
		nslicenseserver.Forceupdateip = data.Forceupdateip.ValueBool()
	}
	if !data.Licensemode.IsNull() && !data.Licensemode.IsUnknown() {
		nslicenseserver.Licensemode = data.Licensemode.ValueString()
	}
	if !data.Licenseserverip.IsNull() && !data.Licenseserverip.IsUnknown() {
		nslicenseserver.Licenseserverip = data.Licenseserverip.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		nslicenseserver.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		nslicenseserver.Password = data.Password.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		nslicenseserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		nslicenseserver.Servername = data.Servername.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		nslicenseserver.Username = data.Username.ValueString()
	}

	return nslicenseserver
}

// nslicenseserverSetAttrFromGet maps the GET response onto the RESOURCE state.
// Else-branches only null a value when it IsUnknown() so a configured value that
// NITRO omits from GET (secrets, 0/false defaults) is never clobbered
// (omit-on-default backward-compat guard).
func nslicenseserverSetAttrFromGet(ctx context.Context, data *NslicenseserverResourceModel, getResponseData map[string]interface{}) *NslicenseserverResourceModel {
	tflog.Debug(ctx, "In nslicenseserverSetAttrFromGet Function")

	if val, ok := getResponseData["deviceprofilename"]; ok && val != nil {
		data.Deviceprofilename = types.StringValue(val.(string))
	} else if data.Deviceprofilename.IsUnknown() {
		data.Deviceprofilename = types.StringNull()
	}
	if val, ok := getResponseData["forceupdateip"]; ok && val != nil {
		data.Forceupdateip = types.BoolValue(val.(bool))
	} else if data.Forceupdateip.IsUnknown() {
		data.Forceupdateip = types.BoolNull()
	}
	if val, ok := getResponseData["licensemode"]; ok && val != nil {
		data.Licensemode = types.StringValue(val.(string))
	} else if data.Licensemode.IsUnknown() {
		data.Licensemode = types.StringNull()
	}
	if val, ok := getResponseData["licenseserverip"]; ok && val != nil {
		data.Licenseserverip = types.StringValue(val.(string))
	} else if data.Licenseserverip.IsUnknown() {
		data.Licenseserverip = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		data.Nodeid = types.Int64Null()
	}
	// password is a secret - NITRO never returns it. Only resolve an unknown
	// (Create with no config) to null; never clobber a configured value.
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else if data.Password.IsUnknown() {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else if data.Username.IsUnknown() {
		data.Username = types.StringNull()
	}

	// ID matches SDK v2: d.SetId(servername)
	data.Id = types.StringValue(data.Servername.ValueString())

	return data
}

// nslicenseserverSetAttrFromGetForDatasource maps the GET response onto the
// DATASOURCE model, copying every attribute from the response (no state to
// preserve) and setting the ID.
func nslicenseserverSetAttrFromGetForDatasource(ctx context.Context, data *NslicenseserverResourceModel, getResponseData map[string]interface{}) *NslicenseserverResourceModel {
	tflog.Debug(ctx, "In nslicenseserverSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["deviceprofilename"]; ok && val != nil {
		data.Deviceprofilename = types.StringValue(val.(string))
	} else {
		data.Deviceprofilename = types.StringNull()
	}
	if val, ok := getResponseData["forceupdateip"]; ok && val != nil {
		data.Forceupdateip = types.BoolValue(val.(bool))
	} else {
		data.Forceupdateip = types.BoolNull()
	}
	if val, ok := getResponseData["licensemode"]; ok && val != nil {
		data.Licensemode = types.StringValue(val.(string))
	} else {
		data.Licensemode = types.StringNull()
	}
	if val, ok := getResponseData["licenseserverip"]; ok && val != nil {
		data.Licenseserverip = types.StringValue(val.(string))
	} else {
		data.Licenseserverip = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	data.Id = types.StringValue(data.Servername.ValueString())

	return data
}
