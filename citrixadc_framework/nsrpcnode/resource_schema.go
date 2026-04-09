package nsrpcnode

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsrpcnodeResourceModel describes the resource data model.
type NsrpcnodeResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Ipaddress         types.String `tfsdk:"ipaddress"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Secure            types.String `tfsdk:"secure"`
	Srcip             types.String `tfsdk:"srcip"`
	Validatecert      types.String `tfsdk:"validatecert"`
}

func (r *NsrpcnodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsrpcnode resource.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the node. This has to be in the same subnet as the NSIP address.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password to be used in authentication with the peer system node.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password to be used in authentication with the peer system node.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"secure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the channel when talking to the node.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.",
			},
			"validatecert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "validate the server certificate for secure SSL connections",
			},
		},
	}
}

func nsrpcnodeGetThePayloadFromthePlan(ctx context.Context, data *NsrpcnodeResourceModel) ns.Nsrpcnode {
	tflog.Debug(ctx, "In nsrpcnodeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsrpcnode := ns.Nsrpcnode{}
	if !data.Ipaddress.IsNull() {
		nsrpcnode.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Password.IsNull() {
		nsrpcnode.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Secure.IsNull() {
		nsrpcnode.Secure = data.Secure.ValueString()
	}
	if !data.Srcip.IsNull() {
		nsrpcnode.Srcip = data.Srcip.ValueString()
	}
	if !data.Validatecert.IsNull() {
		nsrpcnode.Validatecert = data.Validatecert.ValueString()
	}

	return nsrpcnode
}

func nsrpcnodeGetThePayloadFromtheConfig(ctx context.Context, data *NsrpcnodeResourceModel, payload *ns.Nsrpcnode) {
	tflog.Debug(ctx, "In nsrpcnodeGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func nsrpcnodeSetAttrFromGet(ctx context.Context, data *NsrpcnodeResourceModel, getResponseData map[string]interface{}) *NsrpcnodeResourceModel {
	tflog.Debug(ctx, "In nsrpcnodeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["secure"]; ok && val != nil {
		data.Secure = types.StringValue(val.(string))
	} else {
		data.Secure = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["validatecert"]; ok && val != nil {
		data.Validatecert = types.StringValue(val.(string))
	} else {
		data.Validatecert = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Ipaddress.ValueString()))

	return data
}
