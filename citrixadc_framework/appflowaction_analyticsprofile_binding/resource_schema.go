package appflowaction_analyticsprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppflowactionAnalyticsprofileBindingResourceModel describes the resource data model.
type AppflowactionAnalyticsprofileBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Analyticsprofile types.String `tfsdk:"analyticsprofile"`
	Name             types.String `tfsdk:"name"`
}

func (r *AppflowactionAnalyticsprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appflowaction_analyticsprofile_binding resource.",
			},
			"analyticsprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Analytics profile to be bound to the appflow action",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
		},
	}
}

func appflowaction_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppflowactionAnalyticsprofileBindingResourceModel) appflow.Appflowactionanalyticsprofilebinding {
	tflog.Debug(ctx, "In appflowaction_analyticsprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appflowaction_analyticsprofile_binding := appflow.Appflowactionanalyticsprofilebinding{}
	if !data.Analyticsprofile.IsNull() {
		appflowaction_analyticsprofile_binding.Analyticsprofile = data.Analyticsprofile.ValueString()
	}
	if !data.Name.IsNull() {
		appflowaction_analyticsprofile_binding.Name = data.Name.ValueString()
	}

	return appflowaction_analyticsprofile_binding
}

func appflowaction_analyticsprofile_bindingSetAttrFromGet(ctx context.Context, data *AppflowactionAnalyticsprofileBindingResourceModel, getResponseData map[string]interface{}) *AppflowactionAnalyticsprofileBindingResourceModel {
	tflog.Debug(ctx, "In appflowaction_analyticsprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["analyticsprofile"]; ok && val != nil {
		data.Analyticsprofile = types.StringValue(val.(string))
	} else {
		data.Analyticsprofile = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("analyticsprofile:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Analyticsprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
