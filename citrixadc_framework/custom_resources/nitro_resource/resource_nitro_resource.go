package nitro_resource

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"gopkg.in/yaml.v2"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NitroResourceResource{}
var _ resource.ResourceWithConfigure = (*NitroResourceResource)(nil)
var _ resource.ResourceWithImportState = (*NitroResourceResource)(nil)

func NewNitroResourceResource() resource.Resource {
	return &NitroResourceResource{}
}

// NitroResourceResource is the Plugin Framework port of the legacy SDKv2
// citrixadc_nitro_resource. It is a GENERIC PASSTHROUGH resource: it performs
// dynamic NITRO CRUD over an arbitrary object type that is described entirely by
// configuration.
//
// The behaviour is driven by a YAML "workflows" file (workflows_file) and a
// selected workflow name (workflow). The selected workflow's "lifecycle" key
// chooses which CRUD flavour to run:
//   - "object" / "non_updateable_object" -> named object CRUD
//     (AddResource / FindResourceArrayWithParams / UpdateResource / DeleteResource)
//   - "binding"                          -> binding CRUD
//     (UpdateResource to add / FindResourceArrayWithParams to read /
//     DeleteResourceWithArgs to remove)
//   - "object_by_args"                   -> unnamed/args-keyed object CRUD
//     (AddResource / FindResourceArrayWithParams by args /
//     UpdateUnnamedResource / DeleteResourceWithArgsMap)
//
// This is a faithful backward-compatible migration of the SDKv2 resource: the
// resource type name (citrixadc_nitro_resource), the four user-facing attributes
// (workflows_file, workflow, attributes, non_updateable_attributes), their types
// and optionality, the id scheme (per-lifecycle, see below) and the exact NITRO
// call sequence are all preserved.
type NitroResourceResource struct {
	client *service.NitroClient
}

// NitroResourceResourceModel describes the resource data model. Every schema
// attribute has a matching tfsdk field.
//
// attributes and non_updateable_attributes are string->string maps, exactly as
// in the SDKv2 schema (TypeMap with a TypeString element). The legacy resource
// stored every NITRO value back as a string (fmt.Sprintf("%v", ...)); that is
// preserved here.
type NitroResourceResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	WorkflowsFile           types.String `tfsdk:"workflows_file"`
	Workflow                types.String `tfsdk:"workflow"`
	Attributes              types.Map    `tfsdk:"attributes"`
	NonUpdateableAttributes types.Map    `tfsdk:"non_updateable_attributes"`
}

// nitroResourceMapRequiresReplace is an inline planmodifier.Map that mirrors the
// legacy ForceNew on non_updateable_attributes. The mapplanmodifier helper
// package is not vendored, so the RequiresReplace-on-change semantics are
// implemented directly: replace when the resource is being updated (prior state
// is non-null) and the planned map differs from the state map.
type nitroResourceMapRequiresReplace struct{}

func (m nitroResourceMapRequiresReplace) Description(_ context.Context) string {
	return "If the value of this attribute changes, Terraform will destroy and recreate the resource."
}

func (m nitroResourceMapRequiresReplace) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m nitroResourceMapRequiresReplace) PlanModifyMap(_ context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
	// No replace on create (no prior state) or destroy (no plan).
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	if req.PlanValue.Equal(req.StateValue) {
		return
	}
	resp.RequiresReplace = true
}

func (r *NitroResourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nitro_resource"
}

func (r *NitroResourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NitroResourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nitro_resource resource. Its format depends on the workflow lifecycle (single primary id for objects, `primaryId,secondaryId` for bindings, `key:value,...` for object_by_args).",
			},
			"workflows_file": schema.StringAttribute{
				Required:    true,
				Description: "Path to the YAML file that describes the available NITRO workflows.",
			},
			"workflow": schema.StringAttribute{
				Required:    true,
				Description: "The key of the workflow (inside workflows_file) that drives this resource's CRUD behaviour.",
			},
			"attributes": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The updateable NITRO object attributes, as a map of string values.",
			},
			"non_updateable_attributes": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Map{
					nitroResourceMapRequiresReplace{},
				},
				Description: "The non-updateable NITRO object attributes, as a map of string values. Changing any of these forces resource recreation.",
			},
		},
	}
}

func (r *NitroResourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// -----------------------------------------------------------------------------
// CRUD wrappers (dispatch on workflow["lifecycle"], mirroring the SDKv2 funcs)
// -----------------------------------------------------------------------------

func (r *NitroResourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In createNitroResourceFunc")

	var data NitroResourceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := nitroResourceReadWorkflow(&data)
	if err != nil {
		resp.Diagnostics.AddError("Workflow Error", err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("workflow read %v", workflow))

	switch workflow["lifecycle"] {
	case "object":
		err = r.createObjectFunc(ctx, &data, workflow, &resp.Diagnostics)
	case "non_updateable_object":
		err = r.createObjectFunc(ctx, &data, workflow, &resp.Diagnostics)
	case "binding":
		err = r.createBindingFunc(ctx, &data, workflow, &resp.Diagnostics)
	case "object_by_args":
		err = r.createObjectByArgsFunc(ctx, &data, workflow, &resp.Diagnostics)
	default:
		err = fmt.Errorf("Lifecycle \"%v\" does not have a create function", workflow["lifecycle"])
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NitroResourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "In readNitroResourceFunc")

	var data NitroResourceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := nitroResourceReadWorkflow(&data)
	if err != nil {
		resp.Diagnostics.AddError("Workflow Error", err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("workflow read %v", workflow))

	switch workflow["lifecycle"] {
	case "object":
		err = r.readObjectFunc(ctx, &data, workflow, &resp.Diagnostics)
	case "non_updateable_object":
		err = r.readObjectFunc(ctx, &data, workflow, &resp.Diagnostics)
	case "binding":
		err = r.readBindingFunc(ctx, &data, workflow, &resp.Diagnostics)
	case "object_by_args":
		err = r.readObjectByArgsFunc(ctx, &data, workflow, &resp.Diagnostics)
	default:
		err = fmt.Errorf("Lifecycle \"%v\" does not have a read function", workflow["lifecycle"])
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// SDKv2 signalled a vanished resource by clearing the id (d.SetId("")).
	if data.Id.ValueString() == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NitroResourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In updateNitroResourceFunc")

	var data, state NitroResourceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve the id from prior state; the id is not recomputed on update.
	data.Id = state.Id

	workflow, err := nitroResourceReadWorkflow(&data)
	if err != nil {
		resp.Diagnostics.AddError("Workflow Error", err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("workflow read %v", workflow))

	switch workflow["lifecycle"] {
	case "object":
		err = r.updateObjectFunc(ctx, &data, workflow, &resp.Diagnostics)
	case "object_by_args":
		err = r.updateObjectByArgsFunc(ctx, &data, workflow, &resp.Diagnostics)
	default:
		err = fmt.Errorf("Lifecycle \"%v\" does not have an update function", workflow["lifecycle"])
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NitroResourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In deleteNitroResourceFunc")

	var data NitroResourceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := nitroResourceReadWorkflow(&data)
	if err != nil {
		resp.Diagnostics.AddError("Workflow Error", err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("workflow read %v", workflow))

	switch workflow["lifecycle"] {
	case "object":
		err = r.deleteObjectFunc(ctx, &data, workflow)
	case "non_updateable_object":
		err = r.deleteObjectFunc(ctx, &data, workflow)
	case "binding":
		err = r.deleteBindingFunc(ctx, &data, workflow)
	case "object_by_args":
		err = r.deleteObjectByArgsFunc(ctx, &data, workflow)
	default:
		err = fmt.Errorf("Lifecycle \"%v\" does not have a delete function", workflow["lifecycle"])
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	// Framework removes the resource from state automatically once Delete returns
	// without error.
}

// -----------------------------------------------------------------------------
// object lifecycle
// -----------------------------------------------------------------------------

func (r *NitroResourceResource) createObjectFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In createObjectFunc")

	primaryIdAttribute := workflow["primary_id_attribute"]
	primaryId := nitroResourceGetConfiguredValue(ctx, data, primaryIdAttribute)
	if primaryId == nil {
		return fmt.Errorf("Configured object does not contain primary id attribute %v", primaryIdAttribute)
	}

	object := nitroResourceGetConfiguredMap(ctx, data)

	_, err := r.client.AddResource(workflow["endpoint"].(string), primaryId.(string), &object)
	if err != nil {
		return err
	}

	data.Id = types.StringValue(primaryId.(string))

	if err := r.readObjectFunc(ctx, data, workflow, diags); err != nil {
		return fmt.Errorf("Error when reading created object %s", err)
	}
	return nil
}

func (r *NitroResourceResource) readObjectFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In readObjectFunc")

	primaryId := data.Id.ValueString()

	findParams := service.FindParams{
		ResourceType:             workflow["endpoint"].(string),
		ResourceName:             primaryId,
		ResourceMissingErrorCode: workflow["resource_missing_errorcode"].(int),
	}

	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Error during FindResourceArrayWithParams %s", err.Error()))
		return err
	}

	if len(dataArr) == 0 {
		tflog.Warn(ctx, fmt.Sprintf("Clearing nitro resource state %s", primaryId))
		data.Id = types.StringValue("")
		return nil
	}

	if len(dataArr) > 1 {
		return fmt.Errorf("FindResourceArrayWithParams returned too many results")
	}

	nitroResourceSetConfiguredAttributes(ctx, data, dataArr[0], workflow, diags)
	return nil
}

func (r *NitroResourceResource) updateObjectFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In updateObjectFunc")

	primaryIdAttribute := workflow["primary_id_attribute"].(string)
	primaryId := nitroResourceGetConfiguredValue(ctx, data, primaryIdAttribute)
	if primaryId == nil {
		return fmt.Errorf("Configured object does not contain primary id attribute %v", primaryIdAttribute)
	}

	nitroObject := make(map[string]interface{})
	// The map copy works for simple (string) values; only "attributes" is updateable.
	for k, v := range nitroResourceMapToGo(ctx, data.Attributes) {
		nitroObject[k] = v
	}

	// Add primary id even if defined in non_updateable_attributes.
	if _, ok := nitroObject[primaryIdAttribute]; !ok {
		nitroObject[primaryIdAttribute] = primaryId
	}

	_, err := r.client.UpdateResource(workflow["endpoint"].(string), primaryId.(string), &nitroObject)
	if err != nil {
		return fmt.Errorf("Error updating object %s", err)
	}

	return r.readObjectFunc(ctx, data, workflow, diags)
}

func (r *NitroResourceResource) deleteObjectFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}) error {
	tflog.Debug(ctx, "In deleteObjectFunc")

	primaryId := data.Id.ValueString()
	err := r.client.DeleteResource(workflow["endpoint"].(string), primaryId)
	if err != nil {
		return err
	}
	data.Id = types.StringValue("")
	return nil
}

// -----------------------------------------------------------------------------
// binding lifecycle
// -----------------------------------------------------------------------------

func (r *NitroResourceResource) createBindingFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In createBindingFunc")

	primaryIdAttribute := workflow["primary_id_attribute"]
	primaryId := nitroResourceGetConfiguredValue(ctx, data, primaryIdAttribute)
	if primaryId == nil {
		return fmt.Errorf("Configured binding does not contain primary id attribute %v", primaryIdAttribute)
	}

	secondaryIdAttribute := workflow["secondary_id_attribute"]
	secondaryId := nitroResourceGetConfiguredValue(ctx, data, secondaryIdAttribute)
	if secondaryId == nil {
		return fmt.Errorf("Configured binding does not contain secondary id attribute %v", secondaryIdAttribute)
	}

	binding := nitroResourceGetConfiguredMap(ctx, data)

	_, err := r.client.UpdateResource(workflow["endpoint"].(string), primaryId.(string), &binding)
	if err != nil {
		return err
	}

	bindingId := fmt.Sprintf("%v,%v", primaryId, secondaryId)
	data.Id = types.StringValue(bindingId)

	if err := r.readBindingFunc(ctx, data, workflow, diags); err != nil {
		return fmt.Errorf("Error when reading created binding %s", err)
	}
	return nil
}

func (r *NitroResourceResource) readBindingFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In readBindingFunc")

	bindingId := data.Id.ValueString()
	idSlice := strings.SplitN(bindingId, ",", 2)
	primaryId := idSlice[0]
	secondaryId := idSlice[1]

	findParams := service.FindParams{
		ResourceType:             workflow["endpoint"].(string),
		ResourceName:             primaryId,
		ResourceMissingErrorCode: workflow["bound_resource_missing_errorcode"].(int),
	}

	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Error during FindResourceArrayWithParams %s", err.Error()))
		return err
	}

	if len(dataArr) == 0 {
		tflog.Warn(ctx, fmt.Sprintf("Clearing nitro resource state %s", primaryId))
		data.Id = types.StringValue("")
		return nil
	}

	secondaryIdAttribute := workflow["secondary_id_attribute"]
	foundIndex := -1
	for index, binding := range dataArr {
		if fmt.Sprintf("%v", binding[secondaryIdAttribute.(string)]) == secondaryId {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		tflog.Warn(ctx, fmt.Sprintf("Clearing binding state %s", bindingId))
		data.Id = types.StringValue("")
		return nil
	}

	nitroResourceSetConfiguredAttributes(ctx, data, dataArr[foundIndex], workflow, diags)
	return nil
}

func (r *NitroResourceResource) deleteBindingFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}) error {
	tflog.Debug(ctx, "In deleteBindingFunc")

	bindingId := data.Id.ValueString()
	idSlice := strings.SplitN(bindingId, ",", 2)
	primaryId := idSlice[0]
	secondaryId := idSlice[1]

	secondaryIdAttribute := workflow["secondary_id_attribute"]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("%v:%s", secondaryIdAttribute, secondaryId))

	err := r.client.DeleteResourceWithArgs(workflow["endpoint"].(string), primaryId, args)
	if err != nil {
		return err
	}
	data.Id = types.StringValue("")
	return nil
}

// -----------------------------------------------------------------------------
// object_by_args lifecycle
// -----------------------------------------------------------------------------

func (r *NitroResourceResource) createObjectByArgsFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In createObjectByArgsFunc")

	deleteIdAttributes := workflow["delete_id_attributes"].([]interface{})
	idSlice := make([]string, 0)

	for _, deleteIdAttribute := range deleteIdAttributes {
		attributeValue := nitroResourceGetConfiguredValue(ctx, data, deleteIdAttribute)
		if attributeValue != nil {
			idItem := fmt.Sprintf("%s:%s", deleteIdAttribute, attributeValue)
			idSlice = append(idSlice, idItem)
		}
	}

	if len(idSlice) == 0 {
		return fmt.Errorf("Configured object does not contain any id attribute")
	}

	idString := strings.Join(idSlice, ",")

	object := nitroResourceGetConfiguredMap(ctx, data)

	_, err := r.client.AddResource(workflow["endpoint"].(string), "", &object)
	if err != nil {
		return err
	}

	data.Id = types.StringValue(idString)

	if err := r.readObjectByArgsFunc(ctx, data, workflow, diags); err != nil {
		return fmt.Errorf("Error when reading created object %s", err)
	}
	return nil
}

func (r *NitroResourceResource) readObjectByArgsFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In readObjectByArgsFunc")

	primaryId := data.Id.ValueString()
	idItems := strings.Split(primaryId, ",")
	argsMap := make(map[string]string)

	for _, idItem := range idItems {
		idSlice := strings.Split(idItem, ":")
		key := url.QueryEscape(idSlice[0])
		value := url.QueryEscape(idSlice[1])
		argsMap[key] = value
	}

	findParams := service.FindParams{
		ResourceType:             workflow["endpoint"].(string),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: workflow["resource_missing_errorcode"].(int),
	}

	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Error during FindResourceArrayWithParams %s", err.Error()))
		return err
	}

	if len(dataArr) == 0 {
		tflog.Warn(ctx, fmt.Sprintf("Clearing nitro resource state %s", primaryId))
		data.Id = types.StringValue("")
		return nil
	}

	if len(dataArr) > 1 {
		return fmt.Errorf("FindResourceArrayWithParams returned too many results")
	}

	nitroResourceSetConfiguredAttributes(ctx, data, dataArr[0], workflow, diags)
	return nil
}

func (r *NitroResourceResource) updateObjectByArgsFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In updateObjectByArgsFunc")

	nitroObject := make(map[string]interface{})
	// The map copy works for simple (string) values; only "attributes" is updateable.
	for k, v := range nitroResourceMapToGo(ctx, data.Attributes) {
		nitroObject[k] = v
	}

	for _, key := range workflow["delete_id_attributes"].([]interface{}) {
		value := nitroResourceGetConfiguredValue(ctx, data, key)
		if value == nil {
			continue
		}
		// Add primary ids even if defined in non_updateable_attributes.
		if _, ok := nitroObject[key.(string)]; !ok {
			nitroObject[key.(string)] = value
		}
	}

	err := r.client.UpdateUnnamedResource(workflow["endpoint"].(string), &nitroObject)
	if err != nil {
		return fmt.Errorf("Error updating object %s", err)
	}

	return r.readObjectByArgsFunc(ctx, data, workflow, diags)
}

func (r *NitroResourceResource) deleteObjectByArgsFunc(ctx context.Context, data *NitroResourceResourceModel, workflow map[interface{}]interface{}) error {
	tflog.Debug(ctx, "In deleteObjectByArgsFunc")

	primaryId := data.Id.ValueString()
	idItems := strings.Split(primaryId, ",")
	argsMap := make(map[string]string)

	for _, idItem := range idItems {
		idSlice := strings.Split(idItem, ":")
		key := url.QueryEscape(idSlice[0])
		value := url.QueryEscape(idSlice[1])
		argsMap[key] = value
	}

	err := r.client.DeleteResourceWithArgsMap(workflow["endpoint"].(string), "", argsMap)
	if err != nil {
		return err
	}
	data.Id = types.StringValue("")
	return nil
}

// -----------------------------------------------------------------------------
// helpers (ports of the SDKv2 free functions)
// -----------------------------------------------------------------------------

// nitroResourceReadWorkflow loads the workflows_file YAML and returns the
// selected workflow map, mirroring the SDKv2 readWorkflow.
func nitroResourceReadWorkflow(data *NitroResourceResourceModel) (map[interface{}]interface{}, error) {
	yamlFileName := data.WorkflowsFile.ValueString()
	fileData, err := os.ReadFile(yamlFileName)
	if err != nil {
		return nil, err
	}

	var parsed interface{}
	if err := yaml.Unmarshal(fileData, &parsed); err != nil {
		return nil, err
	}

	topLevel, ok := parsed.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("Top level workflow key not found in workflows yaml file")
	}

	workflowsDict, ok := topLevel["workflow"]
	if !ok {
		return nil, fmt.Errorf("Top level workflow key not found in workflows yaml file")
	}

	specificWorkflow, ok := workflowsDict.(map[interface{}]interface{})[data.Workflow.ValueString()]
	if !ok {
		return nil, fmt.Errorf("Key %v not found in workflows map", data.Workflow.ValueString())
	}

	return specificWorkflow.(map[interface{}]interface{}), nil
}

// nitroResourceMapToGo converts a types.Map of strings to a Go map[string]string.
// A null/unknown map yields an empty (non-nil) map, matching the SDKv2 "not set"
// behaviour.
func nitroResourceMapToGo(ctx context.Context, m types.Map) map[string]string {
	out := make(map[string]string)
	if m.IsNull() || m.IsUnknown() {
		return out
	}
	m.ElementsAs(ctx, &out, false)
	return out
}

// nitroResourceGetConfiguredValue returns the configured value for key, looking
// first in attributes then in non_updateable_attributes (SDKv2 getConfiguredValue).
func nitroResourceGetConfiguredValue(ctx context.Context, data *NitroResourceResourceModel, key interface{}) interface{} {
	keyStr, ok := key.(string)
	if !ok {
		return nil
	}

	if v, ok := nitroResourceMapToGo(ctx, data.Attributes)[keyStr]; ok {
		return v
	}
	if v, ok := nitroResourceMapToGo(ctx, data.NonUpdateableAttributes)[keyStr]; ok {
		return v
	}
	return nil
}

// nitroResourceGetConfiguredMap merges attributes and non_updateable_attributes
// into a single map[string]interface{} for the NITRO payload (SDKv2
// getConfiguredMap).
func nitroResourceGetConfiguredMap(ctx context.Context, data *NitroResourceResourceModel) map[string]interface{} {
	retVal := make(map[string]interface{})
	for k, v := range nitroResourceMapToGo(ctx, data.Attributes) {
		retVal[k] = v
	}
	for k, v := range nitroResourceMapToGo(ctx, data.NonUpdateableAttributes) {
		retVal[k] = v
	}
	return retVal
}

// nitroResourceSetConfiguredAttributes rebuilds the attributes /
// non_updateable_attributes maps from a NITRO GET result. Only the keys that were
// already configured are kept, and every value is stored as a string
// (fmt.Sprintf("%v", ...)), exactly as in SDKv2 setConfiguredAttributes.
func nitroResourceSetConfiguredAttributes(ctx context.Context, data *NitroResourceResourceModel, nitroData map[string]interface{}, workflow map[interface{}]interface{}, diags *diag.Diagnostics) {
	configuredAttrs := nitroResourceMapToGo(ctx, data.Attributes)
	configuredNonUpd := nitroResourceMapToGo(ctx, data.NonUpdateableAttributes)

	attributesMap := make(map[string]string)
	nonUpdateableAttributesMap := make(map[string]string)

	for dataKey, dataValueRaw := range nitroData {
		dataValue := fmt.Sprintf("%v", dataValueRaw)
		if _, ok := configuredAttrs[dataKey]; ok {
			attributesMap[dataKey] = dataValue
		}
		if _, ok := configuredNonUpd[dataKey]; ok {
			nonUpdateableAttributesMap[dataKey] = dataValue
		}
	}

	attrVal, d1 := types.MapValueFrom(ctx, types.StringType, attributesMap)
	nonUpdVal, d2 := types.MapValueFrom(ctx, types.StringType, nonUpdateableAttributesMap)
	diags.Append(d1...)
	diags.Append(d2...)

	data.Attributes = attrVal
	data.NonUpdateableAttributes = nonUpdVal
}
