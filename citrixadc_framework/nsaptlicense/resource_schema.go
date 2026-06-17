package nsaptlicense

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NOTE: Allocating APT license counts via this resource is DISRUPTIVE and
// non-idempotent on the appliance. It performs the NITRO `change` action
// (POST ?action=update); re-running it re-allocates licenses.

// NsaptlicenseResourceModel describes the resource data model.
// The synthetic Terraform identifier is the NITRO License ID ("id"), which is
// the mandatory key for the change/update action. "serialno" is a GET-only
// filter key (see Operations section of the NITRO doc) and is therefore NOT
// sent in the change payload.
type NsaptlicenseResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Bindtype       types.String `tfsdk:"bindtype"`
	Countavailable types.String `tfsdk:"countavailable"`
	Licensedir     types.String `tfsdk:"licensedir"`
	Serialno       types.String `tfsdk:"serialno"`
	Sessionid      types.String `tfsdk:"sessionid"`
	Useproxy       types.String `tfsdk:"useproxy"`
}

func (r *NsaptlicenseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				// NITRO License ID. CLI-mandatory for the change action and used
				// as the Terraform resource identifier.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "License ID",
			},
			"bindtype": schema.StringAttribute{
				Required:    true,
				Description: "Bind type",
			},
			"countavailable": schema.StringAttribute{
				Required:    true,
				Description: "The user can allocate one or more licenses. Ensure the value is less than (for partial allocation) or equal to the total number of available licenses",
			},
			"licensedir": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "License Directory",
			},
			"serialno": schema.StringAttribute{
				// GET-only filter key (Pattern 15) - used to read state back, not
				// part of the change action payload.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Hardware Serial Number/License Activation Code(LAC)",
			},
			"sessionid": schema.StringAttribute{
				Required:    true,
				Description: "Session ID",
			},
			"useproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option.",
			},
		},
	}
}

func nsaptlicenseGetThePayloadFromthePlan(ctx context.Context, data *NsaptlicenseResourceModel) ns.Nsaptlicense {
	tflog.Debug(ctx, "In nsaptlicenseGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// serialno is a GET-only filter (Pattern 15) and is intentionally excluded
	// from the change/update action payload.
	nsaptlicense := ns.Nsaptlicense{}
	if !data.Bindtype.IsNull() && !data.Bindtype.IsUnknown() {
		nsaptlicense.Bindtype = data.Bindtype.ValueString()
	}
	if !data.Countavailable.IsNull() && !data.Countavailable.IsUnknown() {
		nsaptlicense.Countavailable = data.Countavailable.ValueString()
	}
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		nsaptlicense.Id = data.Id.ValueString()
	}
	if !data.Licensedir.IsNull() && !data.Licensedir.IsUnknown() {
		nsaptlicense.Licensedir = data.Licensedir.ValueString()
	}
	if !data.Sessionid.IsNull() && !data.Sessionid.IsUnknown() {
		nsaptlicense.Sessionid = data.Sessionid.ValueString()
	}
	if !data.Useproxy.IsNull() && !data.Useproxy.IsUnknown() {
		nsaptlicense.Useproxy = data.Useproxy.ValueString()
	}

	return nsaptlicense
}

// nsaptlicenseSetAttrFromGet populates the resource model from a GET response
// while preserving the existing Id (set once in Create). It only adopts values
// that the GET response echoes back.
func nsaptlicenseSetAttrFromGet(ctx context.Context, data *NsaptlicenseResourceModel, getResponseData map[string]interface{}) *NsaptlicenseResourceModel {
	tflog.Debug(ctx, "In nsaptlicenseSetAttrFromGet Function")

	if val, ok := getResponseData["bindtype"]; ok && val != nil {
		data.Bindtype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["countavailable"]; ok && val != nil {
		data.Countavailable = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Id = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["licensedir"]; ok && val != nil {
		data.Licensedir = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["serialno"]; ok && val != nil {
		data.Serialno = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["sessionid"]; ok && val != nil {
		data.Sessionid = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["useproxy"]; ok && val != nil {
		data.Useproxy = types.StringValue(val.(string))
	}

	return data
}

// nsaptlicenseSetAttrFromGetForDatasource faithfully copies the GET response
// into the model for the datasource (which has no prior state to preserve) and
// sets the Id from the License ID returned by NITRO.
func nsaptlicenseSetAttrFromGetForDatasource(ctx context.Context, data *NsaptlicenseResourceModel, getResponseData map[string]interface{}) *NsaptlicenseResourceModel {
	tflog.Debug(ctx, "In nsaptlicenseSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bindtype"]; ok && val != nil {
		data.Bindtype = types.StringValue(val.(string))
	} else {
		data.Bindtype = types.StringNull()
	}
	if val, ok := getResponseData["countavailable"]; ok && val != nil {
		data.Countavailable = types.StringValue(val.(string))
	} else {
		data.Countavailable = types.StringNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Id = types.StringValue(val.(string))
	} else {
		data.Id = types.StringNull()
	}
	if val, ok := getResponseData["licensedir"]; ok && val != nil {
		data.Licensedir = types.StringValue(val.(string))
	} else {
		data.Licensedir = types.StringNull()
	}
	if val, ok := getResponseData["serialno"]; ok && val != nil {
		data.Serialno = types.StringValue(val.(string))
	} else {
		data.Serialno = types.StringNull()
	}
	if val, ok := getResponseData["sessionid"]; ok && val != nil {
		data.Sessionid = types.StringValue(val.(string))
	} else {
		data.Sessionid = types.StringNull()
	}
	if val, ok := getResponseData["useproxy"]; ok && val != nil {
		data.Useproxy = types.StringValue(val.(string))
	} else {
		data.Useproxy = types.StringNull()
	}

	return data
}
