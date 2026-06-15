package appfwprofile_starturl_binding

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

// AppfwprofileStarturlBindingResourceModel describes the resource data model.
type AppfwprofileStarturlBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Alertonly      types.String `tfsdk:"alertonly"`
	Comment        types.String `tfsdk:"comment"`
	Isautodeployed types.String `tfsdk:"isautodeployed"`
	Name           types.String `tfsdk:"name"`
	Resourceid     types.String `tfsdk:"resourceid"`
	Ruletype       types.String `tfsdk:"ruletype"`
	Starturl       types.String `tfsdk:"starturl"`
	State          types.String `tfsdk:"state"`
}

func (r *AppfwprofileStarturlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_starturl_binding resource.",
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
				// GET does not echo back comment; keep Optional only so a
				// configured value is not clobbered to null on Read.
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
			"ruletype": schema.StringAttribute{
				// GET does not echo back ruletype; keep Optional only so a
				// configured value is not clobbered to null on Read.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding.",
			},
			"starturl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A regular expression that designates a URL on the Start URL list.",
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

func appfwprofile_starturl_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileStarturlBindingResourceModel) appfw.Appfwprofilestarturlbinding {
	tflog.Debug(ctx, "In appfwprofile_starturl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_starturl_binding := appfw.Appfwprofilestarturlbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_starturl_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_starturl_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_starturl_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_starturl_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_starturl_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_starturl_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.Starturl.IsNull() && !data.Starturl.IsUnknown() {
		appfwprofile_starturl_binding.Starturl = data.Starturl.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_starturl_binding.State = data.State.ValueString()
	}

	return appfwprofile_starturl_binding
}

// appfwprofile_starturl_bindingSetAttrFromGet is the resource-side setter. It
// preserves user-configured fields that the GET response does not echo back
// (comment, ruletype) to avoid spurious "inconsistent result" diffs (P7).
func appfwprofile_starturl_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileStarturlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileStarturlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_starturl_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	}
	// comment is not echoed back by GET; preserve the existing plan/state value.
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	}
	// ruletype is not echoed back by GET; preserve the existing plan/state value.
	if val, ok := getResponseData["starturl"]; ok && val != nil {
		data.Starturl = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	}

	return data
}

// appfwprofile_starturl_bindingSetAttrFromGetForDatasource faithfully copies the
// GET response into the model and sets the composite ID (datasource has no Create).
func appfwprofile_starturl_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileStarturlBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileStarturlBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_starturl_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["starturl"]; ok && val != nil {
		data.Starturl = types.StringValue(val.(string))
	} else {
		data.Starturl = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("starturl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Starturl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
