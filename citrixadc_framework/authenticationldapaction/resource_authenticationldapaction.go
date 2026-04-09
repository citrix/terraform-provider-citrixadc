package authenticationldapaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthenticationldapactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationldapactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationldapactionResource)(nil)

func NewAuthenticationldapactionResource() resource.Resource {
	return &AuthenticationldapactionResource{}
}

// AuthenticationldapactionResource defines the resource implementation.
type AuthenticationldapactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationldapactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationldapactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationldapaction"
}

func (r *AuthenticationldapactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationldapactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationldapactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationldapaction resource")
	// Get payload from plan (regular attributes)
	authenticationldapaction := authenticationldapactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationldapactionGetThePayloadFromtheConfig(ctx, &config, &authenticationldapaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationldapaction.Type(), name_value, &authenticationldapaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationldapaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationldapaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationldapactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationldapactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationldapactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationldapaction resource")

	r.readAuthenticationldapactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationldapactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationldapactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationldapaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Alternateemailattr.Equal(state.Alternateemailattr) {
		tflog.Debug(ctx, fmt.Sprintf("alternateemailattr has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute1.Equal(state.Attribute1) {
		tflog.Debug(ctx, fmt.Sprintf("attribute1 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute10.Equal(state.Attribute10) {
		tflog.Debug(ctx, fmt.Sprintf("attribute10 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute11.Equal(state.Attribute11) {
		tflog.Debug(ctx, fmt.Sprintf("attribute11 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute12.Equal(state.Attribute12) {
		tflog.Debug(ctx, fmt.Sprintf("attribute12 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute13.Equal(state.Attribute13) {
		tflog.Debug(ctx, fmt.Sprintf("attribute13 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute14.Equal(state.Attribute14) {
		tflog.Debug(ctx, fmt.Sprintf("attribute14 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute15.Equal(state.Attribute15) {
		tflog.Debug(ctx, fmt.Sprintf("attribute15 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute16.Equal(state.Attribute16) {
		tflog.Debug(ctx, fmt.Sprintf("attribute16 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute2.Equal(state.Attribute2) {
		tflog.Debug(ctx, fmt.Sprintf("attribute2 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute3.Equal(state.Attribute3) {
		tflog.Debug(ctx, fmt.Sprintf("attribute3 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute4.Equal(state.Attribute4) {
		tflog.Debug(ctx, fmt.Sprintf("attribute4 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute5.Equal(state.Attribute5) {
		tflog.Debug(ctx, fmt.Sprintf("attribute5 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute6.Equal(state.Attribute6) {
		tflog.Debug(ctx, fmt.Sprintf("attribute6 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute7.Equal(state.Attribute7) {
		tflog.Debug(ctx, fmt.Sprintf("attribute7 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute8.Equal(state.Attribute8) {
		tflog.Debug(ctx, fmt.Sprintf("attribute8 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attribute9.Equal(state.Attribute9) {
		tflog.Debug(ctx, fmt.Sprintf("attribute9 has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Attributes.Equal(state.Attributes) {
		tflog.Debug(ctx, fmt.Sprintf("attributes has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Authentication.Equal(state.Authentication) {
		tflog.Debug(ctx, fmt.Sprintf("authentication has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Authtimeout.Equal(state.Authtimeout) {
		tflog.Debug(ctx, fmt.Sprintf("authtimeout has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Cloudattributes.Equal(state.Cloudattributes) {
		tflog.Debug(ctx, fmt.Sprintf("cloudattributes has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Email.Equal(state.Email) {
		tflog.Debug(ctx, fmt.Sprintf("email has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Followreferrals.Equal(state.Followreferrals) {
		tflog.Debug(ctx, fmt.Sprintf("followreferrals has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Groupattrname.Equal(state.Groupattrname) {
		tflog.Debug(ctx, fmt.Sprintf("groupattrname has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Groupnameidentifier.Equal(state.Groupnameidentifier) {
		tflog.Debug(ctx, fmt.Sprintf("groupnameidentifier has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Groupsearchattribute.Equal(state.Groupsearchattribute) {
		tflog.Debug(ctx, fmt.Sprintf("groupsearchattribute has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Groupsearchfilter.Equal(state.Groupsearchfilter) {
		tflog.Debug(ctx, fmt.Sprintf("groupsearchfilter has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Groupsearchsubattribute.Equal(state.Groupsearchsubattribute) {
		tflog.Debug(ctx, fmt.Sprintf("groupsearchsubattribute has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Kbattribute.Equal(state.Kbattribute) {
		tflog.Debug(ctx, fmt.Sprintf("kbattribute has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Ldapbase.Equal(state.Ldapbase) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbase has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Ldapbinddn.Equal(state.Ldapbinddn) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbinddn has changed for authenticationldapaction"))
		hasChange = true
	}
	// Check secret attribute ldapbinddnpassword or its version tracker
	if !data.Ldapbinddnpassword.Equal(state.Ldapbinddnpassword) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbinddnpassword has changed for authenticationldapaction"))
		hasChange = true
	} else if !data.LdapbinddnpasswordWoVersion.Equal(state.LdapbinddnpasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbinddnpassword_wo_version has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Ldaphostname.Equal(state.Ldaphostname) {
		tflog.Debug(ctx, fmt.Sprintf("ldaphostname has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Ldaploginname.Equal(state.Ldaploginname) {
		tflog.Debug(ctx, fmt.Sprintf("ldaploginname has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Maxldapreferrals.Equal(state.Maxldapreferrals) {
		tflog.Debug(ctx, fmt.Sprintf("maxldapreferrals has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Maxnestinglevel.Equal(state.Maxnestinglevel) {
		tflog.Debug(ctx, fmt.Sprintf("maxnestinglevel has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Mssrvrecordlocation.Equal(state.Mssrvrecordlocation) {
		tflog.Debug(ctx, fmt.Sprintf("mssrvrecordlocation has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Nestedgroupextraction.Equal(state.Nestedgroupextraction) {
		tflog.Debug(ctx, fmt.Sprintf("nestedgroupextraction has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Otpsecret.Equal(state.Otpsecret) {
		tflog.Debug(ctx, fmt.Sprintf("otpsecret has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Passwdchange.Equal(state.Passwdchange) {
		tflog.Debug(ctx, fmt.Sprintf("passwdchange has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Pushservice.Equal(state.Pushservice) {
		tflog.Debug(ctx, fmt.Sprintf("pushservice has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Referraldnslookup.Equal(state.Referraldnslookup) {
		tflog.Debug(ctx, fmt.Sprintf("referraldnslookup has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Requireuser.Equal(state.Requireuser) {
		tflog.Debug(ctx, fmt.Sprintf("requireuser has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Searchfilter.Equal(state.Searchfilter) {
		tflog.Debug(ctx, fmt.Sprintf("searchfilter has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Sectype.Equal(state.Sectype) {
		tflog.Debug(ctx, fmt.Sprintf("sectype has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, fmt.Sprintf("serverip has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Servername.Equal(state.Servername) {
		tflog.Debug(ctx, fmt.Sprintf("servername has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, fmt.Sprintf("serverport has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Sshpublickey.Equal(state.Sshpublickey) {
		tflog.Debug(ctx, fmt.Sprintf("sshpublickey has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Ssonameattribute.Equal(state.Ssonameattribute) {
		tflog.Debug(ctx, fmt.Sprintf("ssonameattribute has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Subattributename.Equal(state.Subattributename) {
		tflog.Debug(ctx, fmt.Sprintf("subattributename has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Svrtype.Equal(state.Svrtype) {
		tflog.Debug(ctx, fmt.Sprintf("svrtype has changed for authenticationldapaction"))
		hasChange = true
	}
	if !data.Validateservercert.Equal(state.Validateservercert) {
		tflog.Debug(ctx, fmt.Sprintf("validateservercert has changed for authenticationldapaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationldapaction := authenticationldapactionGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationldapactionGetThePayloadFromtheConfig(ctx, &config, &authenticationldapaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationldapaction.Type(), name_value, &authenticationldapaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationldapaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationldapaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationldapaction resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationldapactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationldapactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationldapactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationldapaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationldapaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationldapaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationldapaction resource")
}

// Helper function to read authenticationldapaction data from API
func (r *AuthenticationldapactionResource) readAuthenticationldapactionFromApi(ctx context.Context, data *AuthenticationldapactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationldapaction.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationldapaction, got error: %s", err))
		return
	}

	authenticationldapactionSetAttrFromGet(ctx, data, getResponseData)

}
