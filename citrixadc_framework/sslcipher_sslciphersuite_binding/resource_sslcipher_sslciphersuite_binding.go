package sslcipher_sslciphersuite_binding

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
var _ resource.Resource = &SslcipherSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslcipherSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslcipherSslciphersuiteBindingResource)(nil)

func NewSslcipherSslciphersuiteBindingResource() resource.Resource {
	return &SslcipherSslciphersuiteBindingResource{}
}

// SslcipherSslciphersuiteBindingResource defines the resource implementation.
type SslcipherSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslcipherSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcipherSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcipher_sslciphersuite_binding"
}

func (r *SslcipherSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcipherSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcipher_sslciphersuite_binding resource")
	sslcipher_sslciphersuite_binding := sslcipher_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - SDK v2 used AddResource (POST); match for backward compatibility
	_, err := r.client.AddResource(service.Sslcipher_sslciphersuite_binding.Type(), "", &sslcipher_sslciphersuite_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcipher_sslciphersuite_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphergroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphergroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcipher_sslciphersuite_binding resource")

	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslcipher_sslciphersuite_binding; all attributes are
	// RequiresReplace (SDK v2 marked every field ForceNew and had no UpdateContext).
	// Terraform recreates on any change, so this branch is never reached with a diff.
	tflog.Debug(ctx, "Update is a no-op for sslcipher_sslciphersuite_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcipher_sslciphersuite_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"ciphergroupname", "ciphername"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	ciphergroupname_value, ok := idMap["ciphergroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'ciphergroupname' not found in ID")
		return
	}

	// Delete args are appended raw to the URL by the NITRO client, so URL-encode the
	// value here to handle ciphernames containing slashes/special characters.
	args := make([]string, 0)
	if val, ok := idMap["ciphername"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ciphername:%s", url.QueryEscape(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcipher_sslciphersuite_binding binding")
}

// Helper function to read sslcipher_sslciphersuite_binding data from API
func (r *SslcipherSslciphersuiteBindingResource) readSslcipherSslciphersuiteBindingFromApi(ctx context.Context, data *SslcipherSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"ciphergroupname", "ciphername"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	ciphergroupname_Name, ok := idMap["ciphergroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ciphergroupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslcipher_sslciphersuite_binding.Type(),
		ResourceName:             ciphergroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslcipher_sslciphersuite_binding returned empty array.")
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
		diags.AddError("Client Error", fmt.Sprintf("sslcipher_sslciphersuite_binding not found with the provided ID attributes"))
		return
	}

	sslcipher_sslciphersuite_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
