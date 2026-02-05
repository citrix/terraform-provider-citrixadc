package ssllogprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SsllogprofileResourceModel describes the resource data model.
type SsllogprofileResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Ssllogclauth         types.String `tfsdk:"ssllogclauth"`
	Ssllogclauthfailures types.String `tfsdk:"ssllogclauthfailures"`
	Sslloghs             types.String `tfsdk:"sslloghs"`
	Sslloghsfailures     types.String `tfsdk:"sslloghsfailures"`
}

func (r *SsllogprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssllogprofile resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the ssllogprofile.",
			},
			"ssllogclauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "log all SSL ClAuth events.",
			},
			"ssllogclauthfailures": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "log all SSL ClAuth error events.",
			},
			"sslloghs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "log all SSL HS events.",
			},
			"sslloghsfailures": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "log all SSL HS error events.",
			},
		},
	}
}

func ssllogprofileGetThePayloadFromtheConfig(ctx context.Context, data *SsllogprofileResourceModel) ssl.Ssllogprofile {
	tflog.Debug(ctx, "In ssllogprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ssllogprofile := ssl.Ssllogprofile{}
	if !data.Name.IsNull() {
		ssllogprofile.Name = data.Name.ValueString()
	}
	if !data.Ssllogclauth.IsNull() {
		ssllogprofile.Ssllogclauth = data.Ssllogclauth.ValueString()
	}
	if !data.Ssllogclauthfailures.IsNull() {
		ssllogprofile.Ssllogclauthfailures = data.Ssllogclauthfailures.ValueString()
	}
	if !data.Sslloghs.IsNull() {
		ssllogprofile.Sslloghs = data.Sslloghs.ValueString()
	}
	if !data.Sslloghsfailures.IsNull() {
		ssllogprofile.Sslloghsfailures = data.Sslloghsfailures.ValueString()
	}

	return ssllogprofile
}

func ssllogprofileSetAttrFromGet(ctx context.Context, data *SsllogprofileResourceModel, getResponseData map[string]interface{}) *SsllogprofileResourceModel {
	tflog.Debug(ctx, "In ssllogprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ssllogclauth"]; ok && val != nil {
		data.Ssllogclauth = types.StringValue(val.(string))
	} else {
		data.Ssllogclauth = types.StringNull()
	}
	if val, ok := getResponseData["ssllogclauthfailures"]; ok && val != nil {
		data.Ssllogclauthfailures = types.StringValue(val.(string))
	} else {
		data.Ssllogclauthfailures = types.StringNull()
	}
	if val, ok := getResponseData["sslloghs"]; ok && val != nil {
		data.Sslloghs = types.StringValue(val.(string))
	} else {
		data.Sslloghs = types.StringNull()
	}
	if val, ok := getResponseData["sslloghsfailures"]; ok && val != nil {
		data.Sslloghsfailures = types.StringValue(val.(string))
	} else {
		data.Sslloghsfailures = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
