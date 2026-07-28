package appfwprofile_grpcvalidation_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileGrpcvalidationBindingResourceModel describes the resource data model.
type AppfwprofileGrpcvalidationBindingResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Alertonly                 types.String `tfsdk:"alertonly"`
	Comment                   types.String `tfsdk:"comment"`
	GrpcRelaxValidationAction types.String `tfsdk:"grpc_relax_validation_action"`
	Grpcvalidation            types.String `tfsdk:"grpcvalidation"`
	Isautodeployed            types.String `tfsdk:"isautodeployed"`
	Name                      types.String `tfsdk:"name"`
	Resourceid                types.String `tfsdk:"resourceid"`
	State                     types.String `tfsdk:"state"`
}

func (r *AppfwprofileGrpcvalidationBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_grpcvalidation_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"grpc_relax_validation_action": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Action to be taken for traffic matching the configured relaxation rule.",
			},
			"grpcvalidation": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Exempt any gRPC method matching the given pattern from the API schema validation check. Example: bookstore.api.doc.AddBook",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_grpcvalidation_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileGrpcvalidationBindingResourceModel) appfw.Appfwprofilegrpcvalidationbinding {
	tflog.Debug(ctx, "In appfwprofile_grpcvalidation_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// NOTE (Pattern 15 sanctioned exclusion): alertonly and resourceid are
	// read-only / server-assigned GET-response attributes for the grpcValidation
	// bind branch. The CLI does not accept them for
	// `bind appfw profile ... -grpcValidation ...`, so they are kept Computed in the
	// schema (populated from GET) but NOT sent in the write payload. (ruletype is a
	// cross-branch attribute and is not part of this resource's model at all, so
	// there is nothing to exclude for it here.)
	appfwprofile_grpcvalidation_binding := appfw.Appfwprofilegrpcvalidationbinding{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_grpcvalidation_binding.Comment = data.Comment.ValueString()
	}
	if !data.GrpcRelaxValidationAction.IsNull() && !data.GrpcRelaxValidationAction.IsUnknown() {
		appfwprofile_grpcvalidation_binding.Grpcrelaxvalidationaction = data.GrpcRelaxValidationAction.ValueString()
	}
	if !data.Grpcvalidation.IsNull() && !data.Grpcvalidation.IsUnknown() {
		appfwprofile_grpcvalidation_binding.Grpcvalidation = data.Grpcvalidation.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_grpcvalidation_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_grpcvalidation_binding.Name = data.Name.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_grpcvalidation_binding.State = data.State.ValueString()
	}

	return appfwprofile_grpcvalidation_binding
}

// appfwprofile_grpcvalidation_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-supplied plan/state values for the write-able / identity
// attributes (name, grpcvalidation, grpc_relax_validation_action, comment, state,
// isautodeployed) so that a GET response (which may normalize or default these)
// does not produce an "inconsistent result after apply" diff. It only copies the
// server-managed read-only attributes (resourceid, alertonly) from the response.
// The ID is composed once in Create and never recomputed here (Pattern 6).
func appfwprofile_grpcvalidation_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileGrpcvalidationBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileGrpcvalidationBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_grpcvalidation_bindingSetAttrFromGet Function")

	// Server-managed read-only Computed attributes - copied from the GET response
	// so the Computed values become known after apply (Pattern 7 ECHOED branch).
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}

	// comment and state are echoed verbatim by the GET row, so read them back to
	// make `terraform import` round-trip (category (b)). The appliance does not
	// normalize these values, so populating them from the response does not
	// introduce an "inconsistent result after apply" diff in the basic test.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// name, grpcvalidation and grpc_relax_validation_action are identity /
	// composite-ID components; they are backfilled from the parsed ID in
	// readAppfwprofileGrpcvalidationBindingFromApi (so import round-trips) and are
	// not overwritten from the GET response here.

	return data
}

// appfwprofile_grpcvalidation_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every attribute from the GET response and composes the ID.
func appfwprofile_grpcvalidation_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileGrpcvalidationBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileGrpcvalidationBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_grpcvalidation_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["grpc_relax_validation_action"]; ok && val != nil {
		data.GrpcRelaxValidationAction = types.StringValue(val.(string))
	} else {
		data.GrpcRelaxValidationAction = types.StringNull()
	}
	if val, ok := getResponseData["grpcvalidation"]; ok && val != nil {
		data.Grpcvalidation = types.StringValue(val.(string))
	} else {
		data.Grpcvalidation = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Compose the ID (datasource has no Create).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("grpc_relax_validation_action:%s", utils.UrlEncode(fmt.Sprintf("%v", data.GrpcRelaxValidationAction.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("grpcvalidation:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Grpcvalidation.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
