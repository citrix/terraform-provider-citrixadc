package nslicenseserver

import (
	"context"
	"fmt"
	"strings"

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
			"deviceprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Device profile is created on ADM and contains the user name and password of the instance(s). ADM will use this info to add the NS for registration",
			},
			"forceupdateip": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.",
			},
			"licensemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This paramter indicates type of license customer interested while configuring add/set licenseserver",
			},
			"licenseserverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the License server.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"password": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password to use when authenticating with ADM Agent for LAS licensing.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "License server port.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the License server.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
		},
	}
}

func nslicenseserverGetThePayloadFromtheConfig(ctx context.Context, data *NslicenseserverResourceModel) ns.Nslicenseserver {
	tflog.Debug(ctx, "In nslicenseserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nslicenseserver := ns.Nslicenseserver{}
	if !data.Deviceprofilename.IsNull() {
		nslicenseserver.Deviceprofilename = data.Deviceprofilename.ValueString()
	}
	if !data.Forceupdateip.IsNull() {
		nslicenseserver.Forceupdateip = data.Forceupdateip.ValueBool()
	}
	if !data.Licensemode.IsNull() {
		nslicenseserver.Licensemode = data.Licensemode.ValueString()
	}
	if !data.Licenseserverip.IsNull() {
		nslicenseserver.Licenseserverip = data.Licenseserverip.ValueString()
	}
	if !data.Nodeid.IsNull() {
		nslicenseserver.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Password.IsNull() {
		nslicenseserver.Password = data.Password.ValueString()
	}
	if !data.Port.IsNull() {
		nslicenseserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Servername.IsNull() {
		nslicenseserver.Servername = data.Servername.ValueString()
	}
	if !data.Username.IsNull() {
		nslicenseserver.Username = data.Username.ValueString()
	}

	return nslicenseserver
}

func nslicenseserverSetAttrFromGet(ctx context.Context, data *NslicenseserverResourceModel, getResponseData map[string]interface{}) *NslicenseserverResourceModel {
	tflog.Debug(ctx, "In nslicenseserverSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource based on which identifiers are present
	var idParts []string
	if !data.Licenseserverip.IsNull() {
		idParts = append(idParts, fmt.Sprintf("licenseserverip:%s", data.Licenseserverip.ValueString()))
	}
	if !data.Servername.IsNull() {
		idParts = append(idParts, fmt.Sprintf("servername:%s", data.Servername.ValueString()))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
