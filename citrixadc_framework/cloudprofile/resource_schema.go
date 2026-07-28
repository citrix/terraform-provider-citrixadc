package cloudprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CloudprofileResourceModel describes the resource data model.
type CloudprofileResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Azurepollperiod          types.Int64  `tfsdk:"azurepollperiod"`
	Azuretagname             types.String `tfsdk:"azuretagname"`
	Azuretagvalue            types.String `tfsdk:"azuretagvalue"`
	Boundservicegroupsvctype types.String `tfsdk:"boundservicegroupsvctype"`
	Delay                    types.Int64  `tfsdk:"delay"`
	Graceful                 types.String `tfsdk:"graceful"`
	Ipaddress                types.String `tfsdk:"ipaddress"`
	Name                     types.String `tfsdk:"name"`
	Port                     types.Int64  `tfsdk:"port"`
	Servicegroupname         types.String `tfsdk:"servicegroupname"`
	Servicetype              types.String `tfsdk:"servicetype"`
	Type                     types.String `tfsdk:"type"`
	Vservername              types.String `tfsdk:"vservername"`
	Vsvrbindsvcport          types.Int64  `tfsdk:"vsvrbindsvcport"`
}

func (r *CloudprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudprofile resource.",
			},
			"azurepollperiod": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Azure polling period (in seconds)",
			},
			"azuretagname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Azure tag name",
			},
			"azuretagvalue": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Azure tag value",
			},
			"boundservicegroupsvctype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The type of bound service",
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time, in seconds, after which all the services configured on the server are disabled.",
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Cloud profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"port": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number for the virtual server.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "servicegroups bind to this server",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used by the service (also called the service type).",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of cloud profile that you want to create, Vserver or based on Azure Tags",
			},
			"vservername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"vsvrbindsvcport": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The port number to be used for the bound service.",
			},
		},
	}
}

func cloudprofileGetThePayloadFromthePlan(ctx context.Context, data *CloudprofileResourceModel) cloud.Cloudprofile {
	tflog.Debug(ctx, "In cloudprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudprofile := cloud.Cloudprofile{}
	if !data.Azurepollperiod.IsNull() && !data.Azurepollperiod.IsUnknown() {
		cloudprofile.Azurepollperiod = utils.IntPtr(int(data.Azurepollperiod.ValueInt64()))
	}
	if !data.Azuretagname.IsNull() && !data.Azuretagname.IsUnknown() {
		cloudprofile.Azuretagname = data.Azuretagname.ValueString()
	}
	if !data.Azuretagvalue.IsNull() && !data.Azuretagvalue.IsUnknown() {
		cloudprofile.Azuretagvalue = data.Azuretagvalue.ValueString()
	}
	if !data.Boundservicegroupsvctype.IsNull() && !data.Boundservicegroupsvctype.IsUnknown() {
		cloudprofile.Boundservicegroupsvctype = data.Boundservicegroupsvctype.ValueString()
	}
	if !data.Delay.IsNull() && !data.Delay.IsUnknown() {
		cloudprofile.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Graceful.IsNull() && !data.Graceful.IsUnknown() {
		cloudprofile.Graceful = data.Graceful.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		cloudprofile.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		cloudprofile.Name = data.Name.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		cloudprofile.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		cloudprofile.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		cloudprofile.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		cloudprofile.Type = data.Type.ValueString()
	}
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		cloudprofile.Vservername = data.Vservername.ValueString()
	}
	if !data.Vsvrbindsvcport.IsNull() && !data.Vsvrbindsvcport.IsUnknown() {
		cloudprofile.Vsvrbindsvcport = utils.IntPtr(int(data.Vsvrbindsvcport.ValueInt64()))
	}

	return cloudprofile
}

func cloudprofileSetAttrFromGet(ctx context.Context, data *CloudprofileResourceModel, getResponseData map[string]interface{}) *CloudprofileResourceModel {
	tflog.Debug(ctx, "In cloudprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["azurepollperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Azurepollperiod = types.Int64Value(intVal)
		}
	} else {
		data.Azurepollperiod = types.Int64Null()
	}
	if val, ok := getResponseData["azuretagname"]; ok && val != nil {
		data.Azuretagname = types.StringValue(val.(string))
	} else {
		data.Azuretagname = types.StringNull()
	}
	if val, ok := getResponseData["azuretagvalue"]; ok && val != nil {
		data.Azuretagvalue = types.StringValue(val.(string))
	} else {
		data.Azuretagvalue = types.StringNull()
	}
	if val, ok := getResponseData["boundservicegroupsvctype"]; ok && val != nil {
		data.Boundservicegroupsvctype = types.StringValue(val.(string))
	} else {
		data.Boundservicegroupsvctype = types.StringNull()
	}
	if val, ok := getResponseData["delay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delay = types.Int64Value(intVal)
		}
	} else {
		data.Delay = types.Int64Null()
	}
	if val, ok := getResponseData["graceful"]; ok && val != nil {
		data.Graceful = types.StringValue(val.(string))
	} else {
		data.Graceful = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}
	if val, ok := getResponseData["vsvrbindsvcport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vsvrbindsvcport = types.Int64Value(intVal)
		}
	} else {
		data.Vsvrbindsvcport = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
