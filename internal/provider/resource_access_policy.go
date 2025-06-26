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

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AccessPolicyResource{}
var _ resource.ResourceWithImportState = &AccessPolicyResource{}

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
	SQL          types.Object  `json:"sql,omitempty" tfsdk:"sql"`
	Schema       types.Dynamic `json:"schema,omitempty" tfsdk:"schema"`
	And          types.Dynamic `json:"and,omitempty" tfsdk:"and"`
	Or           types.Dynamic `json:"or,omitempty" tfsdk:"or"`
	RPC          types.Dynamic `json:"rpc,omitempty" tfsdk:"rpc"`
	Link         types.List    `json:"link,omitempty" tfsdk:"link"`
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
			},
			"sql": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "SQL configuration for the Access Policy. Only used when engine is set to 'sql'.",
				Attributes: map[string]schema.Attribute{
					"query": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "SQL query to execute. Should return a single row with one column containing a boolean result.",
					},
				},
			},
			"schema": schema.DynamicAttribute{
				Optional:            true,
				MarkdownDescription: "JSON Schema object for validation. Only used when engine is set to 'json-schema'.",
			},
			"and": schema.DynamicAttribute{
				Optional:            true,
				MarkdownDescription: "Array of engine rules that must all be satisfied (AND logic). Only used when engine is set to 'complex'. Cannot be used together with 'or'.",
			},
			"or": schema.DynamicAttribute{
				Optional:            true,
				MarkdownDescription: "Array of engine rules where at least one must be satisfied (OR logic). Only used when engine is set to 'complex'. Cannot be used together with 'and'.",
			},
			"rpc": schema.DynamicAttribute{
				Optional:            true,
				MarkdownDescription: "RPC configuration object. Only used when engine is set to 'matcho-rpc' or 'allow-rpc'.",
			},
			"link": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "List of links to Users, Clients, or Operations. If empty, the policy is considered global.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"resource_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Type of resource to link to. Must be one of: User, Client, Operation",
							Validators: []validator.String{
								stringvalidator.OneOf("User", "Client", "Operation"),
							},
						},
						"id": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "ID of the resource to link to",
						},
					},
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

	if !data.SQL.IsNull() && !data.SQL.IsUnknown() {
		sqlAttrs := data.SQL.Attributes()
		if queryVal, ok := sqlAttrs["query"]; ok {
			if queryStr, ok := queryVal.(types.String); ok {
				accessPolicy.SQL = &client.SQLConfig{
					Query: queryStr.ValueString(),
				}
			}
		}
	}

	if !data.Schema.IsNull() && !data.Schema.IsUnknown() {
		if objValue, ok := data.Schema.UnderlyingValue().(types.Object); ok {
			accessPolicy.Schema = convertObjectToMap(objValue)
		}
	}

	if !data.And.IsNull() && !data.And.IsUnknown() {
		if listValue, ok := data.And.UnderlyingValue().(types.List); ok {
			andElements := make([]interface{}, 0)
			for _, elem := range listValue.Elements() {
				if objValue, ok := elem.(types.Dynamic); ok {
					if underlyingObj, ok := objValue.UnderlyingValue().(types.Object); ok {
						andElements = append(andElements, convertObjectToMap(underlyingObj))
					}
				}
			}
			accessPolicy.And = andElements
		}
	}

	if !data.Or.IsNull() && !data.Or.IsUnknown() {
		if listValue, ok := data.Or.UnderlyingValue().(types.List); ok {
			orElements := make([]interface{}, 0)
			for _, elem := range listValue.Elements() {
				if objValue, ok := elem.(types.Dynamic); ok {
					if underlyingObj, ok := objValue.UnderlyingValue().(types.Object); ok {
						orElements = append(orElements, convertObjectToMap(underlyingObj))
					}
				}
			}
			accessPolicy.Or = orElements
		}
	}

	if !data.RPC.IsNull() && !data.RPC.IsUnknown() {
		if objValue, ok := data.RPC.UnderlyingValue().(types.Object); ok {
			accessPolicy.RPC = convertObjectToMap(objValue)
		}
	}

	if !data.Link.IsNull() && !data.Link.IsUnknown() {
		linkElements := make([]client.AccessPolicyLink, 0)
		for _, elem := range data.Link.Elements() {
			if objValue, ok := elem.(types.Object); ok {
				linkAttrs := objValue.Attributes()
				var link client.AccessPolicyLink
				if resourceTypeVal, ok := linkAttrs["resource_type"]; ok {
					if resourceTypeStr, ok := resourceTypeVal.(types.String); ok {
						link.ResourceType = resourceTypeStr.ValueString()
					}
				}
				if idVal, ok := linkAttrs["id"]; ok {
					if idStr, ok := idVal.(types.String); ok {
						link.ID = idStr.ValueString()
					}
				}
				linkElements = append(linkElements, link)
			}
		}
		accessPolicy.Link = linkElements
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
		// If the resource doesn't exist, mark it for recreation
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
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

	if sql, ok := policy["sql"]; ok {
		if sqlMap, ok := sql.(map[string]interface{}); ok {
			sqlValues := map[string]attr.Value{}
			if query, ok := sqlMap["query"].(string); ok {
				sqlValues["query"] = types.StringValue(query)
			}
			sqlTypes := map[string]attr.Type{
				"query": types.StringType,
			}
			sqlObj, diags := types.ObjectValue(sqlTypes, sqlValues)
			if !diags.HasError() {
				data.SQL = sqlObj
			}
		}
	}

	if schema, ok := policy["schema"]; ok {
		if schemaMap, ok := schema.(map[string]interface{}); ok {
			objValue := convertMapToObject(schemaMap)
			data.Schema = types.DynamicValue(objValue)
		}
	}

	if and, ok := policy["and"]; ok {
		if andArray, ok := and.([]interface{}); ok {
			andElements := make([]attr.Value, 0, len(andArray))
			for _, item := range andArray {
				if itemMap, ok := item.(map[string]interface{}); ok {
					objValue := convertMapToObject(itemMap)
					andElements = append(andElements, types.DynamicValue(objValue))
				}
			}
			andList, diags := types.ListValue(types.DynamicType, andElements)
			if !diags.HasError() {
				data.And = types.DynamicValue(andList)
			}
		}
	}

	if or, ok := policy["or"]; ok {
		if orArray, ok := or.([]interface{}); ok {
			orElements := make([]attr.Value, 0, len(orArray))
			for _, item := range orArray {
				if itemMap, ok := item.(map[string]interface{}); ok {
					objValue := convertMapToObject(itemMap)
					orElements = append(orElements, types.DynamicValue(objValue))
				}
			}
			orList, diags := types.ListValue(types.DynamicType, orElements)
			if !diags.HasError() {
				data.Or = types.DynamicValue(orList)
			}
		}
	}

	if rpc, ok := policy["rpc"]; ok {
		if rpcMap, ok := rpc.(map[string]interface{}); ok {
			objValue := convertMapToObject(rpcMap)
			data.RPC = types.DynamicValue(objValue)
		}
	}

	if link, ok := policy["link"]; ok {
		if linkArray, ok := link.([]interface{}); ok {
			linkElements := make([]attr.Value, 0, len(linkArray))
			for _, item := range linkArray {
				if itemMap, ok := item.(map[string]interface{}); ok {
					linkValues := map[string]attr.Value{}
					if resourceType, ok := itemMap["resourceType"].(string); ok {
						linkValues["resource_type"] = types.StringValue(resourceType)
					}
					if id, ok := itemMap["id"].(string); ok {
						linkValues["id"] = types.StringValue(id)
					}
					linkTypes := map[string]attr.Type{
						"resource_type": types.StringType,
						"id":            types.StringType,
					}
					linkObj, diags := types.ObjectValue(linkTypes, linkValues)
					if !diags.HasError() {
						linkElements = append(linkElements, linkObj)
					}
				}
			}
			linkElementType := types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"resource_type": types.StringType,
					"id":            types.StringType,
				},
			}
			linkList, diags := types.ListValue(linkElementType, linkElements)
			if !diags.HasError() {
				data.Link = linkList
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

	if !data.SQL.IsNull() && !data.SQL.IsUnknown() {
		sqlAttrs := data.SQL.Attributes()
		if queryVal, ok := sqlAttrs["query"]; ok {
			if queryStr, ok := queryVal.(types.String); ok {
				accessPolicy.SQL = &client.SQLConfig{
					Query: queryStr.ValueString(),
				}
			}
		}
	}

	if !data.Schema.IsNull() && !data.Schema.IsUnknown() {
		if objValue, ok := data.Schema.UnderlyingValue().(types.Object); ok {
			accessPolicy.Schema = convertObjectToMap(objValue)
		}
	}

	if !data.And.IsNull() && !data.And.IsUnknown() {
		if listValue, ok := data.And.UnderlyingValue().(types.List); ok {
			andElements := make([]interface{}, 0)
			for _, elem := range listValue.Elements() {
				if objValue, ok := elem.(types.Dynamic); ok {
					if underlyingObj, ok := objValue.UnderlyingValue().(types.Object); ok {
						andElements = append(andElements, convertObjectToMap(underlyingObj))
					}
				}
			}
			accessPolicy.And = andElements
		}
	}

	if !data.Or.IsNull() && !data.Or.IsUnknown() {
		if listValue, ok := data.Or.UnderlyingValue().(types.List); ok {
			orElements := make([]interface{}, 0)
			for _, elem := range listValue.Elements() {
				if objValue, ok := elem.(types.Dynamic); ok {
					if underlyingObj, ok := objValue.UnderlyingValue().(types.Object); ok {
						orElements = append(orElements, convertObjectToMap(underlyingObj))
					}
				}
			}
			accessPolicy.Or = orElements
		}
	}

	if !data.RPC.IsNull() && !data.RPC.IsUnknown() {
		if objValue, ok := data.RPC.UnderlyingValue().(types.Object); ok {
			accessPolicy.RPC = convertObjectToMap(objValue)
		}
	}

	if !data.Link.IsNull() && !data.Link.IsUnknown() {
		linkElements := make([]client.AccessPolicyLink, 0)
		for _, elem := range data.Link.Elements() {
			if objValue, ok := elem.(types.Object); ok {
				linkAttrs := objValue.Attributes()
				var link client.AccessPolicyLink
				if resourceTypeVal, ok := linkAttrs["resource_type"]; ok {
					if resourceTypeStr, ok := resourceTypeVal.(types.String); ok {
						link.ResourceType = resourceTypeStr.ValueString()
					}
				}
				if idVal, ok := linkAttrs["id"]; ok {
					if idStr, ok := idVal.(types.String); ok {
						link.ID = idStr.ValueString()
					}
				}
				linkElements = append(linkElements, link)
			}
		}
		accessPolicy.Link = linkElements
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

// Helper function to convert a map[string]interface{} to a types.Object.
func convertMapToObject(m map[string]interface{}) types.Object {
	attrValues := make(map[string]attr.Value)
	attrTypes := make(map[string]attr.Type)

	for k, v := range m {
		if v == nil {
			continue
		}
		switch val := v.(type) {
		case string:
			attrValues[k] = types.StringValue(val)
			attrTypes[k] = types.StringType
		case float64:
			attrValues[k] = types.NumberValue(big.NewFloat(val))
			attrTypes[k] = types.NumberType
		case bool:
			attrValues[k] = types.BoolValue(val)
			attrTypes[k] = types.BoolType
		case map[string]interface{}:
			nestedObj := convertMapToObject(val)
			attrValues[k] = nestedObj
			attrTypes[k] = nestedObj.Type(context.Background())
		case []interface{}:
			// Handle arrays
			listElements := make([]attr.Value, 0, len(val))
			var elementType attr.Type = types.StringType // default
			for _, item := range val {
				switch itemVal := item.(type) {
				case string:
					listElements = append(listElements, types.StringValue(itemVal))
					elementType = types.StringType
				case float64:
					listElements = append(listElements, types.NumberValue(big.NewFloat(itemVal)))
					elementType = types.NumberType
				case bool:
					listElements = append(listElements, types.BoolValue(itemVal))
					elementType = types.BoolType
				case map[string]interface{}:
					nestedObj := convertMapToObject(itemVal)
					listElements = append(listElements, nestedObj)
					elementType = nestedObj.Type(context.Background())
				}
			}
			if len(listElements) > 0 {
				listValue, _ := types.ListValue(elementType, listElements)
				attrValues[k] = listValue
				attrTypes[k] = types.ListType{ElemType: elementType}
			}
		}
	}

	objValue, _ := types.ObjectValue(attrTypes, attrValues)
	return objValue
}
