package search

import (
	"testing"
)

func TestUserSearch(t *testing.T) {
	//	tests := []struct {
	//		name      string
	//		validator *Validator
	//	}{
	//		{
	//			name: "Should return no error for a valid lambda descriptor resource",
	//			validator: NewValidator(NewResource(ResourceTypeDescriptor, "valid resource", strings.NewReader(`description: "Valid lambda descriptor"
	//attributesOverride:
	//  memory: 128
	//  prefetchMicrosEnvVars: true
	//`), string(types.ComponentTypeLambda))),
	//		},
	//	}
	//
	//	for _, tt := range tests {
	//		t.Run(tt.name, func(t *testing.T) {
	//			validationResult := tt.validator.ValidateResource()
	//			if validationResult.Err != nil {
	//				t.Errorf("Expected no error, got %v", validationResult.Err)
	//			}
	//
	//			if validationResult.Result.Valid() == false {
	//				t.Errorf("Expected valid result, got %v", validationResult.Result)
	//			}
	//		})
	//	}
}

func TestTicketSearch(t *testing.T) {
	//tests := []struct {
	//	name      string
	//	validator *Validator
	//}{
	//	{
	//		name: "Should return no error for a valid workflow resource",
	//		validator: NewValidator(
	//			NewResource(
	//				ResourceTypeWorkflow,
	//				"valid resource",
	//				strings.NewReader(`{}`),
	//				string(types.ComponentTypeStepFunctions),
	//			),
	//		),
	//	},
	//}
	//
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		validationResult := tt.validator.ValidateResource()
	//		if validationResult.Err != nil {
	//			t.Errorf("Expected no error, got %v", validationResult.Err)
	//		}
	//	})
	//}
}

func TestInvalidateResourceDescriptor(t *testing.T) {
	//	tests := []struct {
	//		name      string
	//		validator *Validator
	//	}{
	//		{
	//			name: "Should return error for an invalid lambda descriptor resource",
	//			validator: NewValidator(
	//				NewResource(ResourceTypeDescriptor, "missing attributesOverride", strings.NewReader(`description: "invalid lambda descriptor"
	//`), string(types.ComponentTypeLambda))),
	//		},
	//	}
	//
	//	for _, tt := range tests {
	//		t.Run(tt.name, func(t *testing.T) {
	//			validationResult := tt.validator.ValidateResource()
	//
	//			if validationResult.Err != nil {
	//				t.Errorf("Expected no error, got %v", validationResult.Err)
	//			}
	//
	//			if validationResult.Result.Valid() == true {
	//				t.Errorf("Expected invalid result, got %v: %v", validationResult.Result, validationResult.Result.Valid())
	//			}
	//		})
	//	}
}

func TestOrganizationSearch(t *testing.T) {
	//tests := []struct {
	//	name      string
	//	validator *Validator
	//}{
	//	{
	//		name:      "Should return error for an invalid workflow resource",
	//		validator: NewValidator(NewResource(ResourceTypeDescriptor, "missing attributesOverride", strings.NewReader(`{`), string(types.ComponentTypeLambda))),
	//	},
	//}
	//
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		validationResult := tt.validator.ValidateResource()
	//
	//		if validationResult.Err == nil {
	//			t.Errorf("Expected error, got %v", validationResult.Err)
	//		}
	//	})
	//}
}
