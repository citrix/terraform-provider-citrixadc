package gslbvserver_domain_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbvserverDomainBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbvserverDomainBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbvserverDomainBindingResource)(nil)

func NewGslbvserverDomainBindingResource() resource.Resource {
	return &GslbvserverDomainBindingResource{}
}

// GslbvserverDomainBindingResource defines the resource implementation.
type GslbvserverDomainBindingResource struct {
	client *service.NitroClient
}

func (r *GslbvserverDomainBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbvserverDomainBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver_domain_binding"
}

func (r *GslbvserverDomainBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbvserverDomainBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbvserverDomainBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbvserver_domain_binding resource")
	gslbvserver_domain_binding := gslbvserver_domain_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Gslbvserver_domain_binding.Type(), &gslbvserver_domain_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbvserver_domain_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created gslbvserver_domain_binding resource")

	// Set ID for the resource before reading state.
	// Legacy SDK v2 ID order is name,domainname (see resource_id_mapping.json). Use the new
	// key:urlEncode(value) format; ParseIdString decodes both this and the legacy form.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("domainname:%s", utils.UrlEncode(data.Domainname.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readGslbvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverDomainBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbvserverDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbvserver_domain_binding resource")

	r.readGslbvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverDomainBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbvserverDomainBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating gslbvserver_domain_binding resource")

	// Check if there are any changes in the in-place-updateable attributes.
	// backupip / cookiedomain / backupipflag / cookiedomainflag / domainname / order are
	// all RequiresReplace, so they never reach Update.
	hasChange := false
	if !data.Cookietimeout.Equal(state.Cookietimeout) {
		tflog.Debug(ctx, fmt.Sprintf("cookietimeout has changed for gslbvserver_domain_binding"))
		hasChange = true
	}
	if !data.Sitedomainttl.Equal(state.Sitedomainttl) {
		tflog.Debug(ctx, fmt.Sprintf("sitedomainttl has changed for gslbvserver_domain_binding"))
		hasChange = true
	}
	if !data.Ttl.Equal(state.Ttl) {
		tflog.Debug(ctx, fmt.Sprintf("ttl has changed for gslbvserver_domain_binding"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		gslbvserver_domain_binding := gslbvserver_domain_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Gslbvserver_domain_binding.Type(), &gslbvserver_domain_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbvserver_domain_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated gslbvserver_domain_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for gslbvserver_domain_binding resource, skipping update")
	}

	// Read the updated state back
	r.readGslbvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverDomainBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbvserverDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbvserver_domain_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Legacy ID order: name (parent) + domainname (delete arg). ParseIdString handles both
	// the new key:value form and the legacy positional name,domainname form.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "domainname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	domainname_value, ok := idMap["domainname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Attribute 'domainname' not found in ID")
		return
	}

	// URL-encode the domainname value so slashy/special characters are passed safely as a
	// NITRO delete arg. (task hint b)
	args := []string{fmt.Sprintf("domainname:%s", utils.UrlEncode(domainname_value))}

	err = r.client.DeleteResourceWithArgs(service.Gslbvserver_domain_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbvserver_domain_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbvserver_domain_binding binding")
}

// Helper function to read gslbvserver_domain_binding data from API
func (r *GslbvserverDomainBindingResource) readGslbvserverDomainBindingFromApi(ctx context.Context, data *GslbvserverDomainBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "domainname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Gslbvserver_domain_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver_domain_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "gslbvserver_domain_binding returned empty array.")
		return
	}

	domainname_Name, ok := idMap["domainname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'domainname' not found in ID string")
		return
	}

	// Iterate through results to find the one matching domainname (the per-record unique key
	// under the parent name; backupipflag/cookiedomainflag are not real keys).
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["domainname"].(string); ok && val == domainname_Name {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("gslbvserver_domain_binding not found with the provided ID attributes"))
		return
	}

	gslbvserver_domain_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
