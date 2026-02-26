package csvserver_gslbvserver_binding

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

// CsvserverGslbvserverBindingResourceModel describes the resource data model.
type CsvserverGslbvserverBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Vserver types.String `tfsdk:"vserver"`
}

func (r *CsvserverGslbvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver_gslbvserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the default gslb or vpn vserver bound to CS vserver of type GSLB/VPN. For Example: bind cs vserver cs1 -vserver gslb1 or bind cs vserver cs1 -vserver vpn1",
			},
		},
	}
}

func csvserver_gslbvserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *CsvserverGslbvserverBindingResourceModel) cs.Csvservergslbvserverbinding {
	tflog.Debug(ctx, "In csvserver_gslbvserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	csvserver_gslbvserver_binding := cs.Csvservergslbvserverbinding{}
	if !data.Name.IsNull() {
		csvserver_gslbvserver_binding.Name = data.Name.ValueString()
	}
	if !data.Vserver.IsNull() {
		csvserver_gslbvserver_binding.Vserver = data.Vserver.ValueString()
	}

	return csvserver_gslbvserver_binding
}

func csvserver_gslbvserver_bindingSetAttrFromGet(ctx context.Context, data *CsvserverGslbvserverBindingResourceModel, getResponseData map[string]interface{}) *CsvserverGslbvserverBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_gslbvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vserver:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
