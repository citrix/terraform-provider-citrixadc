package appfwprofile_restvalidation_binding

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

// AppfwprofileRestvalidationBindingResourceModel describes the resource data model.
type AppfwprofileRestvalidationBindingResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Alertonly            types.String `tfsdk:"alertonly"`
	Comment              types.String `tfsdk:"comment"`
	Isautodeployed       types.String `tfsdk:"isautodeployed"`
	Name                 types.String `tfsdk:"name"`
	Resourceid           types.String `tfsdk:"resourceid"`
	RestValidationAction types.String `tfsdk:"rest_validation_action"`
	Restvalidation       types.String `tfsdk:"restvalidation"`
	State                types.String `tfsdk:"state"`
}

func (r *AppfwprofileRestvalidationBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_restvalidation_binding resource.",
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
			"rest_validation_action": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Action to be taken for traffic matching the configured relaxation rule.",
			},
			"restvalidation": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Exempt REST endpoints or any other URLs matching the given pattern from the API schema validation check. Example: GET:/v1/bookstore/viewbooks",
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

func appfwprofile_restvalidation_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileRestvalidationBindingResourceModel) appfw.Appfwprofilerestvalidationbinding {
	tflog.Debug(ctx, "In appfwprofile_restvalidation_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// NOTE (Pattern 15/16 sanctioned exclusions): alertonly (GET-only response
	// field), isautodeployed (read-only/derived AUTODEPLOYED|NOTAUTODEPLOYED) and
	// resourceid (server-assigned read-only id) are NOT accepted by the CLI for
	// `bind appfw profile ... -restValidation ...`. They are kept Computed in the
	// schema (populated from GET) but excluded from the write payload. ruletype is a
	// cross-branch (SQLInjection) attribute and is not part of this resource's model
	// at all, so there is nothing to exclude for it here.
	appfwprofile_restvalidation_binding := appfw.Appfwprofilerestvalidationbinding{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_restvalidation_binding.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_restvalidation_binding.Name = data.Name.ValueString()
	}
	if !data.RestValidationAction.IsNull() && !data.RestValidationAction.IsUnknown() {
		appfwprofile_restvalidation_binding.Restvalidationaction = data.RestValidationAction.ValueString()
	}
	if !data.Restvalidation.IsNull() && !data.Restvalidation.IsUnknown() {
		appfwprofile_restvalidation_binding.Restvalidation = data.Restvalidation.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_restvalidation_binding.State = data.State.ValueString()
	}

	return appfwprofile_restvalidation_binding
}

// appfwprofile_restvalidation_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-supplied plan/state values for the write-able / identity
// attributes (name, restvalidation, rest_validation_action, comment, state) so
// that a GET response (which may normalize or default these) does not produce an
// "inconsistent result after apply" diff. It only copies the server-managed
// read-only attributes (resourceid, alertonly, isautodeployed) from the response.
// The ID is composed once in Create and never recomputed here (Pattern 6).
func appfwprofile_restvalidation_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileRestvalidationBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileRestvalidationBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_restvalidation_bindingSetAttrFromGet Function")

	// Server-managed read-only attributes - safe to copy from the GET response.
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

	// comment and state are echoed verbatim by the GET row (the appliance does not
	// normalize them), so populate them here to make `terraform import` round-trip.
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

	// name, restvalidation and rest_validation_action are identity / ID-component
	// attributes; they are backfilled from the parsed composite ID in
	// readAppfwprofileRestvalidationBindingFromApi (authoritative on import).

	return data
}

// appfwprofile_restvalidation_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every attribute from the GET response and composes the ID.
func appfwprofile_restvalidation_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileRestvalidationBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileRestvalidationBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_restvalidation_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["rest_validation_action"]; ok && val != nil {
		data.RestValidationAction = types.StringValue(val.(string))
	} else {
		data.RestValidationAction = types.StringNull()
	}
	if val, ok := getResponseData["restvalidation"]; ok && val != nil {
		data.Restvalidation = types.StringValue(val.(string))
	} else {
		data.Restvalidation = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Compose the ID (datasource has no Create).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("rest_validation_action:%s", utils.UrlEncode(fmt.Sprintf("%v", data.RestValidationAction.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("restvalidation:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Restvalidation.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
