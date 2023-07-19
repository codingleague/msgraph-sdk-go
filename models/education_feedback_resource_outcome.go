package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationFeedbackResourceOutcome 
type EducationFeedbackResourceOutcome struct {
    EducationOutcome
}
// NewEducationFeedbackResourceOutcome instantiates a new educationFeedbackResourceOutcome and sets the default values.
func NewEducationFeedbackResourceOutcome()(*EducationFeedbackResourceOutcome) {
    m := &EducationFeedbackResourceOutcome{
        EducationOutcome: *NewEducationOutcome(),
    }
    odataTypeValue := "#microsoft.graph.educationFeedbackResourceOutcome"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationFeedbackResourceOutcomeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationFeedbackResourceOutcomeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationFeedbackResourceOutcome(), nil
}
// GetFeedbackResource gets the feedbackResource property value. The actual feedback resource.
func (m *EducationFeedbackResourceOutcome) GetFeedbackResource()(EducationResourceable) {
    val, err := m.GetBackingStore().Get("feedbackResource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationResourceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationFeedbackResourceOutcome) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationOutcome.GetFieldDeserializers()
    res["feedbackResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedbackResource(val.(EducationResourceable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["resourceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationFeedbackResourceOutcomeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceStatus(val.(*EducationFeedbackResourceOutcomeStatus))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationFeedbackResourceOutcome) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceStatus gets the resourceStatus property value. The status of the feedback resource. The possible values are: notPublished, pendingPublish, published, failedPublish, unknownFutureValue.
func (m *EducationFeedbackResourceOutcome) GetResourceStatus()(*EducationFeedbackResourceOutcomeStatus) {
    val, err := m.GetBackingStore().Get("resourceStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EducationFeedbackResourceOutcomeStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationFeedbackResourceOutcome) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationOutcome.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("feedbackResource", m.GetFeedbackResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResourceStatus() != nil {
        cast := (*m.GetResourceStatus()).String()
        err = writer.WriteStringValue("resourceStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFeedbackResource sets the feedbackResource property value. The actual feedback resource.
func (m *EducationFeedbackResourceOutcome) SetFeedbackResource(value EducationResourceable)() {
    err := m.GetBackingStore().Set("feedbackResource", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationFeedbackResourceOutcome) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceStatus sets the resourceStatus property value. The status of the feedback resource. The possible values are: notPublished, pendingPublish, published, failedPublish, unknownFutureValue.
func (m *EducationFeedbackResourceOutcome) SetResourceStatus(value *EducationFeedbackResourceOutcomeStatus)() {
    err := m.GetBackingStore().Set("resourceStatus", value)
    if err != nil {
        panic(err)
    }
}
// EducationFeedbackResourceOutcomeable 
type EducationFeedbackResourceOutcomeable interface {
    EducationOutcomeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFeedbackResource()(EducationResourceable)
    GetOdataType()(*string)
    GetResourceStatus()(*EducationFeedbackResourceOutcomeStatus)
    SetFeedbackResource(value EducationResourceable)()
    SetOdataType(value *string)()
    SetResourceStatus(value *EducationFeedbackResourceOutcomeStatus)()
}
