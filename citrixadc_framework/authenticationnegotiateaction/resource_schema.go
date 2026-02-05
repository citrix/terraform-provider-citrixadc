package authenticationnegotiateaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationnegotiateactionResourceModel describes the resource data model.
type AuthenticationnegotiateactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Domain                     types.String `tfsdk:"domain"`
	Domainuser                 types.String `tfsdk:"domainuser"`
	Domainuserpasswd           types.String `tfsdk:"domainuserpasswd"`
	Keytab                     types.String `tfsdk:"keytab"`
	Name                       types.String `tfsdk:"name"`
	Ntlmpath                   types.String `tfsdk:"ntlmpath"`
	Ou                         types.String `tfsdk:"ou"`
}

func (r *AuthenticationnegotiateactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationnegotiateaction resource.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the service principal that represnts Citrix ADC.",
			},
			"domainuser": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User name of the account that is mapped with Citrix ADC principal. This can be given along with domain and password when keytab file is not available. If username is given along with keytab file, then that keytab file will be searched for this user's credentials.",
			},
			"domainuserpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password of the account that is mapped to the Citrix ADC principal.",
			},
			"keytab": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the keytab file that is used to decrypt kerberos tickets presented to Citrix ADC. If keytab is not available, domain/username/password can be specified in the negotiate action configuration",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the AD KDC server profile (negotiate action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AD KDC server profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"ntlmpath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the site that is enabled for NTLM authentication, including FQDN of the server. This is used when clients fallback to NTLM.",
			},
			"ou": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Active Directory organizational units (OU) attribute.",
			},
		},
	}
}

func authenticationnegotiateactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationnegotiateactionResourceModel) authentication.Authenticationnegotiateaction {
	tflog.Debug(ctx, "In authenticationnegotiateactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationnegotiateaction := authentication.Authenticationnegotiateaction{}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationnegotiateaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Domain.IsNull() {
		authenticationnegotiateaction.Domain = data.Domain.ValueString()
	}
	if !data.Domainuser.IsNull() {
		authenticationnegotiateaction.Domainuser = data.Domainuser.ValueString()
	}
	if !data.Domainuserpasswd.IsNull() {
		authenticationnegotiateaction.Domainuserpasswd = data.Domainuserpasswd.ValueString()
	}
	if !data.Keytab.IsNull() {
		authenticationnegotiateaction.Keytab = data.Keytab.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationnegotiateaction.Name = data.Name.ValueString()
	}
	if !data.Ntlmpath.IsNull() {
		authenticationnegotiateaction.Ntlmpath = data.Ntlmpath.ValueString()
	}
	if !data.Ou.IsNull() {
		authenticationnegotiateaction.Ou = data.Ou.ValueString()
	}

	return authenticationnegotiateaction
}

func authenticationnegotiateactionSetAttrFromGet(ctx context.Context, data *AuthenticationnegotiateactionResourceModel, getResponseData map[string]interface{}) *AuthenticationnegotiateactionResourceModel {
	tflog.Debug(ctx, "In authenticationnegotiateactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["domainuser"]; ok && val != nil {
		data.Domainuser = types.StringValue(val.(string))
	} else {
		data.Domainuser = types.StringNull()
	}
	if val, ok := getResponseData["domainuserpasswd"]; ok && val != nil {
		data.Domainuserpasswd = types.StringValue(val.(string))
	} else {
		data.Domainuserpasswd = types.StringNull()
	}
	if val, ok := getResponseData["keytab"]; ok && val != nil {
		data.Keytab = types.StringValue(val.(string))
	} else {
		data.Keytab = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ntlmpath"]; ok && val != nil {
		data.Ntlmpath = types.StringValue(val.(string))
	} else {
		data.Ntlmpath = types.StringNull()
	}
	if val, ok := getResponseData["ou"]; ok && val != nil {
		data.Ou = types.StringValue(val.(string))
	} else {
		data.Ou = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
