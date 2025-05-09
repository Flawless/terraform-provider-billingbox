package provider

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-billingbox/internal/client"
)

// Custom validator for matcho attribute.
type matchoValidator struct{}

func (v matchoValidator) Description(ctx context.Context) string {
	return "Matcho is required when engine is set to 'matcho'"
}

func (v matchoValidator) MarkdownDescription(ctx context.Context) string {
	return "Matcho is required when engine is set to 'matcho'"
}

func (v matchoValidator) ValidateDynamic(ctx context.Context, req validator.DynamicRequest, resp *validator.DynamicResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		// Get the engine value
		var engineVal types.String
		diags := req.Config.GetAttribute(ctx, path.Root("engine"), &engineVal)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			return
		}

		if engineVal.ValueString() == "matcho" {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				"The matcho attribute must be set when engine is 'matcho'",
			)
		}
	}
}

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AccessPolicyResource{}
var _ resource.ResourceWithImportState = &AccessPolicyResource{}
var _ validator.Dynamic = matchoValidator{}

func NewAccessPolicyResource() resource.Resource {
	return &AccessPolicyResource{}
}

// AccessPolicyResource defines the resource implementation.
type AccessPolicyResource struct {
	client *client.Client
}

// AccessPolicyResourceModel describes the resource data model.
type AccessPolicyResourceModel struct {
	ResourceType types.String  `json:"resourceType" tfsdk:"resource_type"`
	ID           types.String  `json:"id,omitempty" tfsdk:"id"`
	RoleName     types.String  `json:"roleName,omitempty" tfsdk:"role_name"`
	Engine       types.String  `json:"engine,omitempty" tfsdk:"engine"`
	Matcho       types.Dynamic `json:"matcho,omitempty" tfsdk:"matcho"`
	Meta         types.Object  `json:"meta,omitempty" tfsdk:"meta"`
	Description  types.String  `json:"description,omitempty" tfsdk:"description"`
}

func (r *AccessPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_policy"
}

func (r *AccessPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Access Policy resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Unique identifier for the Access Policy. If not set, a random ID will be generated.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Resource type of the Access Policy. Always set to 'AccessPolicy'.",
				Default:             stringdefault.StaticString("AccessPolicy"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"role_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the Role",
			},
			"engine": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Engine of the Access Policy. Must be one of: json-schema, allow, sql, complex, matcho, clj, matcho-rpc, allow-rpc, signed-rpc",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"json-schema",
						"allow",
						"sql",
						"complex",
						"matcho",
						"clj",
						"matcho-rpc",
						"allow-rpc",
						"signed-rpc",
					),
				},
			},
			"matcho": schema.DynamicAttribute{
				Optional:            true,
				MarkdownDescription: "Match object of the Access Policy. Can contain nested maps with string, number, or boolean values. Only used when engine is set to 'matcho'.",
				Validators: []validator.Dynamic{
					matchoValidator{},
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the Access Policy",
			},
			// shared metadata from schema_metadata.go
			"meta": MetaAttributes(),
		},
	}
}

func (r *AccessPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *AccessPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AccessPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accessPolicy := &client.AccessPolicy{
		Resource: client.Resource{
			ResourceType: "AccessPolicy",
		},
		RoleName:    data.RoleName.ValueString(),
		Engine:      data.Engine.ValueString(),
		Description: data.Description.ValueString(),
	}

	if !data.ID.IsNull() && !data.ID.IsUnknown() {
		accessPolicy.ID = data.ID.ValueString()
	}

	if !data.Matcho.IsNull() && !data.Matcho.IsUnknown() {
		if objValue, ok := data.Matcho.UnderlyingValue().(types.Object); ok {
			accessPolicy.Matcho = convertObjectToMap(objValue)
		}
	}

	// Create the access policy
	result, err := r.client.CreateResource("AccessPolicy", accessPolicy)
	if err != nil {
		resp.Diagnostics.AddError("Error creating access policy", err.Error())
		return
	}

	// Update the model with the response data
	if v, ok := getStringField(result, "id", &resp.Diagnostics, "CreateAccessPolicy"); ok {
		data.ID = types.StringValue(v)
	} else {
		return
	}
	if meta, ok := result["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{}
		if versionID, ok := meta["versionId"].(string); ok {
			metaValues["version_id"] = types.StringValue(versionID)
		}
		if createdAt, ok := meta["createdAt"].(string); ok {
			metaValues["created_at"] = types.StringValue(createdAt)
		}
		if lastUpdated, ok := meta["lastUpdated"].(string); ok {
			metaValues["last_updated"] = types.StringValue(lastUpdated)
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AccessPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := r.client.GetResource("AccessPolicy", data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading access policy", err.Error())
		return
	}

	// Update the model with the response data
	if v, ok := getStringField(policy, "id", &resp.Diagnostics, "ReadAccessPolicy"); ok {
		data.ID = types.StringValue(v)
	} else {
		return
	}
	if v, ok := getStringField(policy, "roleName", &resp.Diagnostics, "ReadAccessPolicy"); ok {
		data.RoleName = types.StringValue(v)
	}
	if v, ok := getStringField(policy, "engine", &resp.Diagnostics, "ReadAccessPolicy"); ok {
		data.Engine = types.StringValue(v)
	}
	if v, ok := getStringField(policy, "resourceType", &resp.Diagnostics, "ReadAccessPolicy"); ok {
		data.ResourceType = types.StringValue(v)
	}

	if matcho, ok := policy["matcho"]; ok {
		// Convert matcho to a map[string]interface{}
		if matchoMap, ok := matcho.(map[string]interface{}); ok {
			// Create a map of attr.Value with the expected structure
			matchoValues := make(map[string]attr.Value)
			for k, v := range matchoMap {
				if v == nil {
					continue
				}
				switch val := v.(type) {
				case string:
					matchoValues[k] = types.StringValue(val)
				case float64:
					matchoValues[k] = types.NumberValue(big.NewFloat(val))
				case bool:
					matchoValues[k] = types.BoolValue(val)
				case map[string]interface{}:
					// Handle nested maps
					nestedValues := make(map[string]attr.Value)
					for nk, nv := range val {
						if nv == nil {
							continue
						}
						switch nval := nv.(type) {
						case string:
							nestedValues[nk] = types.StringValue(nval)
						case float64:
							nestedValues[nk] = types.NumberValue(big.NewFloat(nval))
						case bool:
							nestedValues[nk] = types.BoolValue(nval)
						}
					}
					matchoValues[k] = types.MapValueMust(types.StringType, nestedValues)
				}
			}
			// Create object type map
			attrTypes := make(map[string]attr.Type)
			for k := range matchoValues {
				attrTypes[k] = types.StringType
			}
			// Create dynamic value from object
			objValue, diags := types.ObjectValue(attrTypes, matchoValues)
			if !diags.HasError() {
				data.Matcho = types.DynamicValue(objValue)
			}
		}
	}

	if meta, ok := policy["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{}
		if versionID, ok := meta["versionId"].(string); ok {
			metaValues["version_id"] = types.StringValue(versionID)
		}
		if createdAt, ok := meta["createdAt"].(string); ok {
			metaValues["created_at"] = types.StringValue(createdAt)
		}
		if lastUpdated, ok := meta["lastUpdated"].(string); ok {
			metaValues["last_updated"] = types.StringValue(lastUpdated)
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AccessPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accessPolicy := &client.AccessPolicy{
		RoleName:    data.RoleName.ValueString(),
		Engine:      data.Engine.ValueString(),
		Description: data.Description.ValueString(),
	}

	if obj, ok := data.Matcho.UnderlyingValue().(types.Object); ok {
		accessPolicy.Matcho = convertObjectToMap(obj)
	}

	// Update the access policy
	result, err := r.client.UpdateResource("AccessPolicy", data.ID.ValueString(), accessPolicy)
	if err != nil {
		resp.Diagnostics.AddError("Error updating access policy", err.Error())
		return
	}

	// Update the model with the response data
	if id, ok := result["id"].(string); ok {
		data.ID = types.StringValue(id)
	}
	if meta, ok := result["meta"].(map[string]interface{}); ok {
		metaValues := map[string]attr.Value{}
		if versionID, ok := meta["versionId"].(string); ok {
			metaValues["version_id"] = types.StringValue(versionID)
		}
		if createdAt, ok := meta["createdAt"].(string); ok {
			metaValues["created_at"] = types.StringValue(createdAt)
		}
		if lastUpdated, ok := meta["lastUpdated"].(string); ok {
			metaValues["last_updated"] = types.StringValue(lastUpdated)
		}
		metaTypes := map[string]attr.Type{
			"version_id":   types.StringType,
			"created_at":   types.StringType,
			"last_updated": types.StringType,
		}
		metaObj, diags := types.ObjectValue(metaTypes, metaValues)
		if !diags.HasError() {
			data.Meta = metaObj
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AccessPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteResource("AccessPolicy", data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting access policy", err.Error())
		return
	}
}

func (r *AccessPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Add this function before the Create method.
func convertObjectToMap(obj types.Object) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range obj.Attributes() {
		switch val := v.(type) {
		case types.String:
			result[k] = val.ValueString()
		case types.Number:
			if f, _ := val.ValueBigFloat().Float64(); true {
				result[k] = f
			}
		case types.Bool:
			result[k] = val.ValueBool()
		case types.Object:
			result[k] = convertObjectToMap(val)
		case types.List:
			// Handle lists if needed
			listValues := make([]interface{}, 0)
			for _, item := range val.Elements() {
				switch itemVal := item.(type) {
				case types.String:
					listValues = append(listValues, itemVal.ValueString())
				case types.Number:
					if f, _ := itemVal.ValueBigFloat().Float64(); true {
						listValues = append(listValues, f)
					}
				case types.Bool:
					listValues = append(listValues, itemVal.ValueBool())
				case types.Object:
					listValues = append(listValues, convertObjectToMap(itemVal))
				}
			}
			result[k] = listValues
		}
	}
	return result
}
