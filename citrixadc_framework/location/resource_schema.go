package location

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LocationResourceModel describes the resource data model.
type LocationResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Ipfrom            types.String `tfsdk:"ipfrom"`
	Ipto              types.String `tfsdk:"ipto"`
	Latitude          types.Int64  `tfsdk:"latitude"`
	Longitude         types.Int64  `tfsdk:"longitude"`
	Preferredlocation types.String `tfsdk:"preferredlocation"`
}

func (r *LocationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the location resource.",
			},
			"ipfrom": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "First IP address in the range, in dotted decimal notation.",
			},
			"ipto": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Last IP address in the range, in dotted decimal notation.",
			},
			"latitude": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Numerical value, in degrees, specifying the latitude of the geographical location of the IP address-range.\nNote: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.",
			},
			"longitude": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Numerical value, in degrees, specifying the longitude of the geographical location of the IP address-range.\nNote: Longitude and latitude parameters are used for selecting a service with the static proximity GSLB method. If they are not specified, selection is based on the qualifiers specified for the location.",
			},
			"preferredlocation": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "String of qualifiers, in dotted notation, describing the geographical location of the IP address range. Each qualifier is more specific than the one that precedes it, as in continent.country.region.city.isp.organization. For example, \"NA.US.CA.San Jose.ATT.citrix\".\nNote: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.",
			},
		},
	}
}

func locationGetThePayloadFromtheConfig(ctx context.Context, data *LocationResourceModel) basic.Location {
	tflog.Debug(ctx, "In locationGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	location := basic.Location{}
	if !data.Ipfrom.IsNull() {
		location.Ipfrom = data.Ipfrom.ValueString()
	}
	if !data.Ipto.IsNull() {
		location.Ipto = data.Ipto.ValueString()
	}
	if !data.Latitude.IsNull() {
		location.Latitude = utils.IntPtr(int(data.Latitude.ValueInt64()))
	}
	if !data.Longitude.IsNull() {
		location.Longitude = utils.IntPtr(int(data.Longitude.ValueInt64()))
	}
	if !data.Preferredlocation.IsNull() {
		location.Preferredlocation = data.Preferredlocation.ValueString()
	}

	return location
}

func locationSetAttrFromGet(ctx context.Context, data *LocationResourceModel, getResponseData map[string]interface{}) *LocationResourceModel {
	tflog.Debug(ctx, "In locationSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipfrom"]; ok && val != nil {
		data.Ipfrom = types.StringValue(val.(string))
	} else {
		data.Ipfrom = types.StringNull()
	}
	if val, ok := getResponseData["ipto"]; ok && val != nil {
		data.Ipto = types.StringValue(val.(string))
	} else {
		data.Ipto = types.StringNull()
	}
	if val, ok := getResponseData["latitude"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Latitude = types.Int64Value(intVal)
		}
	} else {
		data.Latitude = types.Int64Null()
	}
	if val, ok := getResponseData["longitude"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Longitude = types.Int64Value(intVal)
		}
	} else {
		data.Longitude = types.Int64Null()
	}
	if val, ok := getResponseData["preferredlocation"]; ok && val != nil {
		data.Preferredlocation = types.StringValue(val.(string))
	} else {
		data.Preferredlocation = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Ipfrom.ValueString())

	return data
}
