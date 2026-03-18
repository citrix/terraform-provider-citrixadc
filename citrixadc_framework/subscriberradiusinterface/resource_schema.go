package subscriberradiusinterface

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SubscriberradiusinterfaceResourceModel describes the resource data model.
type SubscriberradiusinterfaceResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Listeningservice     types.String `tfsdk:"listeningservice"`
	Radiusinterimasstart types.String `tfsdk:"radiusinterimasstart"`
}

func (r *SubscriberradiusinterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the subscriberradiusinterface resource.",
			},
			"listeningservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of RADIUS LISTENING service that will process RADIUS accounting requests.",
			},
			"radiusinterimasstart": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Treat radius interim message as start radius messages.",
			},
		},
	}
}

func subscriberradiusinterfaceGetThePayloadFromtheConfig(ctx context.Context, data *SubscriberradiusinterfaceResourceModel) subscriber.Subscriberradiusinterface {
	tflog.Debug(ctx, "In subscriberradiusinterfaceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	subscriberradiusinterface := subscriber.Subscriberradiusinterface{}
	if !data.Listeningservice.IsNull() {
		subscriberradiusinterface.Listeningservice = data.Listeningservice.ValueString()
	}
	if !data.Radiusinterimasstart.IsNull() {
		subscriberradiusinterface.Radiusinterimasstart = data.Radiusinterimasstart.ValueString()
	}

	return subscriberradiusinterface
}

func subscriberradiusinterfaceSetAttrFromGet(ctx context.Context, data *SubscriberradiusinterfaceResourceModel, getResponseData map[string]interface{}) *SubscriberradiusinterfaceResourceModel {
	tflog.Debug(ctx, "In subscriberradiusinterfaceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["listeningservice"]; ok && val != nil {
		data.Listeningservice = types.StringValue(val.(string))
	} else {
		data.Listeningservice = types.StringNull()
	}
	if val, ok := getResponseData["radiusinterimasstart"]; ok && val != nil {
		data.Radiusinterimasstart = types.StringValue(val.(string))
	} else {
		data.Radiusinterimasstart = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("subscriberradiusinterface-config")

	return data
}
