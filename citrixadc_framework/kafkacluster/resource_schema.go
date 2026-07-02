package kafkacluster

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/kafka"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

)

// KafkaclusterResourceModel describes the resource data model.
type KafkaclusterResourceModel struct {
	Id types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *KafkaclusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the kafkacluster resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Kafka cluster",
			},
		},
	}
}

func kafkaclusterGetThePayloadFromthePlan(ctx context.Context, data *KafkaclusterResourceModel) kafka.Kafkacluster {
	tflog.Debug(ctx, "In kafkaclusterGetThePayloadFromthePlan Function")

	// Create API request body from the model
	kafkacluster := kafka.Kafkacluster{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		kafkacluster.Name = data.Name.ValueString()
	}

	return kafkacluster
}

func kafkaclusterSetAttrFromGet(ctx context.Context, data *KafkaclusterResourceModel, getResponseData map[string]interface{}) *KafkaclusterResourceModel {
	tflog.Debug(ctx, "In kafkaclusterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}