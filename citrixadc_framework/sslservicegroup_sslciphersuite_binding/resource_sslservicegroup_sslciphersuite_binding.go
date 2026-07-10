package sslservicegroup_sslciphersuite_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &SslservicegroupSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupSslciphersuiteBindingResource)(nil)

func NewSslservicegroupSslciphersuiteBindingResource() resource.Resource {
	return &SslservicegroupSslciphersuiteBindingResource{}
}

// SslservicegroupSslciphersuiteBindingResource defines the resource implementation.
type SslservicegroupSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslciphersuite_binding"
}

func (r *SslservicegroupSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup_sslciphersuite_binding resource")
	sslservicegroup_sslciphersuite_binding := sslservicegroup_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslciphersuite_binding.Type(), &sslservicegroup_sslciphersuite_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservicegroup_sslciphersuite_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslservicegroupSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup_sslciphersuite_binding resource")

	r.readSslservicegroupSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservicegroup_sslciphersuite_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservicegroup_sslciphersuite_binding := sslservicegroup_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslciphersuite_binding.Type(), &sslservicegroup_sslciphersuite_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup_sslciphersuite_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservicegroup_sslciphersuite_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservicegroup_sslciphersuite_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslservicegroupSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup_sslciphersuite_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "ciphername"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicegroupname_value, ok := idMap["servicegroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ciphername"]; ok && val != "" {
		// URL-encode the arg value: ciphername can contain slashes/special chars
		// and the NITRO delete helper does not encode arg values.
		argsMap["ciphername"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslservicegroup_sslciphersuite_binding.Type(), servicegroupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservicegroup_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservicegroup_sslciphersuite_binding binding")
}

// Helper function to read sslservicegroup_sslciphersuite_binding data from API
func (r *SslservicegroupSslciphersuiteBindingResource) readSslservicegroupSslciphersuiteBindingFromApi(ctx context.Context, data *SslservicegroupSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "ciphername"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicegroupname_Name, ok := idMap["servicegroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicegroupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslservicegroup_sslciphersuite_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslciphersuite_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslservicegroup_sslciphersuite_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ciphername
		if idVal, ok := idMap["ciphername"]; ok {
			if val, ok := v["ciphername"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ciphername"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("sslservicegroup_sslciphersuite_binding not found with the provided ID attributes"))
		return
	}

	sslservicegroup_sslciphersuite_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
