package nsaigwprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// unsetOnRemoveInt64Modifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-zero value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. It intentionally does nothing when the config
// still carries a value, on create (no prior state), or when the prior value is
// already zero (avoids churn).
type unsetOnRemoveInt64Modifier struct{}

func (m unsetOnRemoveInt64Modifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-zero value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveInt64Modifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveInt64Modifier) PlanModifyInt64(_ context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueInt64() != 0 {
		resp.PlanValue = types.Int64Unknown()
	}
}

// NsaigwprofileResourceModel describes the resource data model.
type NsaigwprofileResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Endpointtype          types.String `tfsdk:"endpointtype"`
	Profiletype           types.String `tfsdk:"profiletype"`
	Tokenquota            types.Int64  `tfsdk:"tokenquota"`
	Quotarefreshfrequency types.Int64  `tfsdk:"quotarefreshfrequency"`
	Authtoken             types.String `tfsdk:"authtoken"`
	AuthtokenWo           types.String `tfsdk:"authtoken_wo"`
	AuthtokenWoVersion    types.Int64  `tfsdk:"authtoken_wo_version"`
}

func (r *NsaigwprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsaigwprofile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AIGW Profile.",
			},
			// endpointtype is required by NITRO on create and is create-only: `set
			// nsaigwprofile -endpointType ...` is rejected (errorcode 278), so a
			// change forces resource replacement. UseStateForUnknown keeps the
			// computed value stable across refreshes (avoids GH #1436 spurious
			// replace). On the current firmware the only accepted value is azureopenai.
			"endpointtype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The type of AI GW endpoint type. Possible values = azureopenai",
			},
			// profiletype is the binding entity (frontend|backend) fixed at creation;
			// it is create-only, so a change forces resource replacement.
			"profiletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The binding entity for the aigw profile. Possible values = frontend, backend",
			},
			"tokenquota": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{},
				},
				Description: "Token capacity of the backend server.",
			},
			"quotarefreshfrequency": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{},
				},
				Description: "Quota refresh rate, in minutes.",
			},
			// authtoken is create-only: NITRO rejects `set nsaigwprofile -authtoken`
			// (errorcode 278). A change therefore forces resource replacement (the
			// new token is applied via a fresh add). Only valid on backend profiles.
			"authtoken": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Authentication token/API Key for the AI GW Endpoint.",
			},
			"authtoken_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Authentication token/API Key for the AI GW Endpoint.",
			},
			// Because authtoken is create-only, rotating the write-only secret (a
			// bump of authtoken_wo_version) cannot be applied in place either; the
			// version change forces resource replacement so the new authtoken_wo is
			// re-added on create.
			"authtoken_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Increment this version to signal an authtoken_wo update.",
			},
		},
	}
}

func nsaigwprofileGetThePayloadFromthePlan(ctx context.Context, data *NsaigwprofileResourceModel) ns.Nsaigwprofile {
	tflog.Debug(ctx, "In nsaigwprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsaigwprofile := ns.Nsaigwprofile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsaigwprofile.Name = data.Name.ValueString()
	}
	if !data.Endpointtype.IsNull() && !data.Endpointtype.IsUnknown() {
		nsaigwprofile.Endpointtype = data.Endpointtype.ValueString()
	}
	if !data.Profiletype.IsNull() && !data.Profiletype.IsUnknown() {
		nsaigwprofile.Profiletype = data.Profiletype.ValueString()
	}
	if !data.Tokenquota.IsNull() && !data.Tokenquota.IsUnknown() {
		nsaigwprofile.Tokenquota = utils.IntPtr(int(data.Tokenquota.ValueInt64()))
	}
	if !data.Quotarefreshfrequency.IsNull() && !data.Quotarefreshfrequency.IsUnknown() {
		nsaigwprofile.Quotarefreshfrequency = utils.IntPtr(int(data.Quotarefreshfrequency.ValueInt64()))
	}
	if !data.Authtoken.IsNull() && !data.Authtoken.IsUnknown() {
		nsaigwprofile.Authtoken = data.Authtoken.ValueString()
	}
	// Skip write-only attribute: authtoken_wo
	// Skip version tracker attribute: authtoken_wo_version

	return nsaigwprofile
}

func nsaigwprofileGetThePayloadFromtheConfig(ctx context.Context, data *NsaigwprofileResourceModel, payload *ns.Nsaigwprofile) {
	tflog.Debug(ctx, "In nsaigwprofileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: authtoken_wo -> authtoken
	if !data.AuthtokenWo.IsNull() {
		authtokenWo := data.AuthtokenWo.ValueString()
		if authtokenWo != "" {
			payload.Authtoken = authtokenWo
		}
	}
}

// nsaigwprofileGetTheUpdatablePayloadFromThePlan builds the NITRO UPDATE payload.
// Only tokenquota and quotarefreshfrequency are mutable via `set nsaigwprofile`;
// endpointtype, profiletype and authtoken are create-only (NITRO rejects them on
// set with errorcode 278) and are RequiresReplace, so they are intentionally
// excluded here (sending them would fail every update).
func nsaigwprofileGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *NsaigwprofileResourceModel) ns.Nsaigwprofile {
	tflog.Debug(ctx, "In nsaigwprofileGetTheUpdatablePayloadFromThePlan Function")

	nsaigwprofile := ns.Nsaigwprofile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsaigwprofile.Name = data.Name.ValueString()
	}
	if !data.Tokenquota.IsNull() && !data.Tokenquota.IsUnknown() {
		nsaigwprofile.Tokenquota = utils.IntPtr(int(data.Tokenquota.ValueInt64()))
	}
	if !data.Quotarefreshfrequency.IsNull() && !data.Quotarefreshfrequency.IsUnknown() {
		nsaigwprofile.Quotarefreshfrequency = utils.IntPtr(int(data.Quotarefreshfrequency.ValueInt64()))
	}

	return nsaigwprofile
}

func nsaigwprofileSetAttrFromGet(ctx context.Context, data *NsaigwprofileResourceModel, getResponseData map[string]interface{}) *NsaigwprofileResourceModel {
	tflog.Debug(ctx, "In nsaigwprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["endpointtype"]; ok && val != nil {
		data.Endpointtype = types.StringValue(val.(string))
	} else {
		data.Endpointtype = types.StringNull()
	}
	if val, ok := getResponseData["profiletype"]; ok && val != nil {
		data.Profiletype = types.StringValue(val.(string))
	} else {
		data.Profiletype = types.StringNull()
	}
	if val, ok := getResponseData["tokenquota"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tokenquota = types.Int64Value(intVal)
		}
	} else {
		data.Tokenquota = types.Int64Null()
	}
	if val, ok := getResponseData["quotarefreshfrequency"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Quotarefreshfrequency = types.Int64Value(intVal)
		}
	} else {
		data.Quotarefreshfrequency = types.Int64Null()
	}
	// authtoken is a secret and is not read back into state - retain from config
	// authtoken_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// authtoken_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
