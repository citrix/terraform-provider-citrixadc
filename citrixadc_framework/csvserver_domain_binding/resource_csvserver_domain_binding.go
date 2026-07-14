package csvserver_domain_binding

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
var _ resource.Resource = &CsvserverDomainBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverDomainBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverDomainBindingResource)(nil)

func NewCsvserverDomainBindingResource() resource.Resource {
	return &CsvserverDomainBindingResource{}
}

// CsvserverDomainBindingResource defines the resource implementation.
type CsvserverDomainBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverDomainBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverDomainBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_domain_binding"
}

func (r *CsvserverDomainBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverDomainBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverDomainBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_domain_binding resource")
	csvserver_domain_binding := csvserver_domain_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Csvserver_domain_binding.Type(), &csvserver_domain_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_domain_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created csvserver_domain_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("domainname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Domainname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readCsvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverDomainBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_domain_binding resource")

	r.readCsvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverDomainBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CsvserverDomainBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating csvserver_domain_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Backupip.Equal(state.Backupip) {
		tflog.Debug(ctx, fmt.Sprintf("backupip has changed for csvserver_domain_binding"))
		hasChange = true
	}
	if !data.Cookiedomain.Equal(state.Cookiedomain) {
		tflog.Debug(ctx, fmt.Sprintf("cookiedomain has changed for csvserver_domain_binding"))
		hasChange = true
	}
	if !data.Cookietimeout.Equal(state.Cookietimeout) {
		tflog.Debug(ctx, fmt.Sprintf("cookietimeout has changed for csvserver_domain_binding"))
		hasChange = true
	}
	if !data.Sitedomainttl.Equal(state.Sitedomainttl) {
		tflog.Debug(ctx, fmt.Sprintf("sitedomainttl has changed for csvserver_domain_binding"))
		hasChange = true
	}
	if !data.Ttl.Equal(state.Ttl) {
		tflog.Debug(ctx, fmt.Sprintf("ttl has changed for csvserver_domain_binding"))
		hasChange = true
	}

	if hasChange {
		// This binding exposes no NITRO "set"/update endpoint (only add/delete/get/count),
		// and the ADC rejects re-binding an already-bound domain via PUT with errorcode 1842
		// ("The domain is already bound to a GSLB vserver"). To change an updateable
		// attribute (ttl, backupip, cookiedomain, cookietimeout, sitedomainttl) we must
		// unbind the existing domain and rebind it with the new values. name + domainname
		// are RequiresReplace identity keys, so they are unchanged here (data.Id == state.Id).
		idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
		if err != nil {
			resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for update: %s", err))
			return
		}
		name_value, ok := idMap["name"]
		if !ok {
			resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
			return
		}
		argsMap := make(map[string]string)
		if val, ok := idMap["domainname"]; ok && val != "" {
			argsMap["domainname"] = val
		}
		// Unbind the existing domain binding.
		if err := r.client.DeleteResourceWithArgsMap(service.Csvserver_domain_binding.Type(), name_value, argsMap); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_domain_binding (unbind step), got error: %s", err))
			return
		}
		// Rebind the domain with the new values.
		csvserver_domain_binding := csvserver_domain_bindingGetThePayloadFromthePlan(ctx, &data)
		if err := r.client.UpdateUnnamedResource(service.Csvserver_domain_binding.Type(), &csvserver_domain_binding); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_domain_binding (rebind step), got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated csvserver_domain_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for csvserver_domain_binding resource, skipping update")
	}

	// Read the updated state back
	r.readCsvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverDomainBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_domain_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["domainname"]; ok && val != "" {
		argsMap["domainname"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Csvserver_domain_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csvserver_domain_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csvserver_domain_binding binding")
}

// Helper function to read csvserver_domain_binding data from API
func (r *CsvserverDomainBindingResource) readCsvserverDomainBindingFromApi(ctx context.Context, data *CsvserverDomainBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
		ResourceType:             service.Csvserver_domain_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_domain_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "csvserver_domain_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check domainname
		if idVal, ok := idMap["domainname"]; ok {
			if val, ok := v["domainname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["domainname"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("csvserver_domain_binding not found with the provided ID attributes"))
		return
	}

	csvserver_domain_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
