package kafkacluster_servicegroup_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/kafka"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// KafkaclusterServicegroupBindingResourceModel describes the resource data model.
type KafkaclusterServicegroupBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
}

func (r *KafkaclusterServicegroupBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the kafkacluster_servicegroup_binding resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Kafka cluster",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the bound servicegroup.",
			},
		},
	}
}

func kafkacluster_servicegroup_bindingGetThePayloadFromthePlan(ctx context.Context, data *KafkaclusterServicegroupBindingResourceModel) kafka.Kafkaclusterservicegroupbinding {
	tflog.Debug(ctx, "In kafkacluster_servicegroup_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	kafkacluster_servicegroup_binding := kafka.Kafkaclusterservicegroupbinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		kafkacluster_servicegroup_binding.Name = data.Name.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		kafkacluster_servicegroup_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}

	return kafkacluster_servicegroup_binding
}

func kafkacluster_servicegroup_bindingSetAttrFromGet(ctx context.Context, data *KafkaclusterServicegroupBindingResourceModel, getResponseData map[string]interface{}) *KafkaclusterServicegroupBindingResourceModel {
	tflog.Debug(ctx, "In kafkacluster_servicegroup_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
