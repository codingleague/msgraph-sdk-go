package policies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody 
type FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ids property
    ids []string
    // The types property
    types []string
}
// NewFeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody instantiates a new FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody()(*FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) {
    m := &FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    res["types"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTypes(res)
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetTypes gets the types property value. The types property
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) GetTypes()([]string) {
    return m.types
}
// Serialize serializes information the current object
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    if m.GetTypes() != nil {
        err := writer.WriteCollectionOfStringValues("types", m.GetTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetTypes sets the types property value. The types property
func (m *FeatureRolloutPoliciesItemAppliesToMicrosoftGraphGetByIdsGetByIdsPostRequestBody) SetTypes(value []string)() {
    m.types = value
}
