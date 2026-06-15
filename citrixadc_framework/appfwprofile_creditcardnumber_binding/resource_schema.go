package appfwprofile_creditcardnumber_binding

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

// AppfwprofileCreditcardnumberBindingResourceModel describes the resource data model.
type AppfwprofileCreditcardnumberBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Alertonly           types.String `tfsdk:"alertonly"`
	Comment             types.String `tfsdk:"comment"`
	Creditcardnumber    types.String `tfsdk:"creditcardnumber"`
	Creditcardnumberurl types.String `tfsdk:"creditcardnumberurl"`
	Isautodeployed      types.String `tfsdk:"isautodeployed"`
	Name                types.String `tfsdk:"name"`
	Resourceid          types.String `tfsdk:"resourceid"`
	Ruletype            types.String `tfsdk:"ruletype"`
	State               types.String `tfsdk:"state"`
}

func (r *AppfwprofileCreditcardnumberBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_creditcardnumber_binding resource.",
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"creditcardnumber": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The object expression that is to be excluded from safe commerce check",
			},
			"creditcardnumberurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The url for which the list of credit card numbers are needed to be bypassed from inspection",
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
			"ruletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_creditcardnumber_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileCreditcardnumberBindingResourceModel) appfw.Appfwprofilecreditcardnumberbinding {
	tflog.Debug(ctx, "In appfwprofile_creditcardnumber_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_creditcardnumber_binding := appfw.Appfwprofilecreditcardnumberbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Comment = data.Comment.ValueString()
	}
	if !data.Creditcardnumber.IsNull() && !data.Creditcardnumber.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Creditcardnumber = data.Creditcardnumber.ValueString()
	}
	if !data.Creditcardnumberurl.IsNull() && !data.Creditcardnumberurl.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Creditcardnumberurl = data.Creditcardnumberurl.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_creditcardnumber_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_creditcardnumber_binding.State = data.State.ValueString()
	}

	return appfwprofile_creditcardnumber_binding
}

func appfwprofile_creditcardnumber_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileCreditcardnumberBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCreditcardnumberBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_creditcardnumber_bindingSetAttrFromGet Function")

	// NOTE: 'alertonly' and 'isautodeployed' are intentionally NOT read back from the
	// GET response here. The NITRO server returns server-side values for these (e.g.
	// "OFF"/"NOTAUTODEPLOYED") that may differ from what the user configured, which would
	// trigger an "inconsistent result after apply" error. The SDK v2 resource preserved
	// the configured values the same way (its d.Set calls for these were commented out).
	// The datasource setter (xxxSetAttrFromGetForDatasource) copies them faithfully.

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["creditcardnumber"]; ok && val != nil {
		data.Creditcardnumber = types.StringValue(val.(string))
	} else {
		data.Creditcardnumber = types.StringNull()
	}
	if val, ok := getResponseData["creditcardnumberurl"]; ok && val != nil {
		data.Creditcardnumberurl = types.StringValue(val.(string))
	} else {
		data.Creditcardnumberurl = types.StringNull()
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
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("creditcardnumber:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Creditcardnumber.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("creditcardnumberurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Creditcardnumberurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// appfwprofile_creditcardnumber_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response (including the server-side 'alertonly' and 'isautodeployed'
// values that the resource setter intentionally preserves) and sets the datasource ID.
func appfwprofile_creditcardnumber_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileCreditcardnumberBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileCreditcardnumberBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_creditcardnumber_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["creditcardnumber"]; ok && val != nil {
		data.Creditcardnumber = types.StringValue(val.(string))
	} else {
		data.Creditcardnumber = types.StringNull()
	}
	if val, ok := getResponseData["creditcardnumberurl"]; ok && val != nil {
		data.Creditcardnumberurl = types.StringValue(val.(string))
	} else {
		data.Creditcardnumberurl = types.StringNull()
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
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("creditcardnumber:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Creditcardnumber.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("creditcardnumberurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Creditcardnumberurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
