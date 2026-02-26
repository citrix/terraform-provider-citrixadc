package csvserver_analyticsprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CsvserverAnalyticsprofileBindingResourceModel describes the resource data model.
type CsvserverAnalyticsprofileBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Analyticsprofile types.String `tfsdk:"analyticsprofile"`
	Name             types.String `tfsdk:"name"`
}

func (r *CsvserverAnalyticsprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver_analyticsprofile_binding resource.",
			},
			"analyticsprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the analytics profile bound to the LB vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
		},
	}
}

func csvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *CsvserverAnalyticsprofileBindingResourceModel) cs.Csvserveranalyticsprofilebinding {
	tflog.Debug(ctx, "In csvserver_analyticsprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	csvserver_analyticsprofile_binding := cs.Csvserveranalyticsprofilebinding{}
	if !data.Analyticsprofile.IsNull() {
		csvserver_analyticsprofile_binding.Analyticsprofile = data.Analyticsprofile.ValueString()
	}
	if !data.Name.IsNull() {
		csvserver_analyticsprofile_binding.Name = data.Name.ValueString()
	}

	return csvserver_analyticsprofile_binding
}

func csvserver_analyticsprofile_bindingSetAttrFromGet(ctx context.Context, data *CsvserverAnalyticsprofileBindingResourceModel, getResponseData map[string]interface{}) *CsvserverAnalyticsprofileBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_analyticsprofile_bindingSetAttrFromGet Function")

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
