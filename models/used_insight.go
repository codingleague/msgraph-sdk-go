package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsedInsight 
type UsedInsight struct {
    Entity
}
// NewUsedInsight instantiates a new usedInsight and sets the default values.
func NewUsedInsight()(*UsedInsight) {
    m := &UsedInsight{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUsedInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsedInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsedInsight(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsedInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUsageDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val.(UsageDetailsable))
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
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(Entityable))
        }
        return nil
    }
    res["resourceReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceReference(val.(ResourceReferenceable))
        }
        return nil
    }
    res["resourceVisualization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceVisualizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceVisualization(val.(ResourceVisualizationable))
        }
        return nil
    }
    return res
}
// GetLastUsed gets the lastUsed property value. Information about when the item was last viewed or modified by the user. Read only.
func (m *UsedInsight) GetLastUsed()(UsageDetailsable) {
    val, err := m.GetBackingStore().Get("lastUsed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UsageDetailsable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UsedInsight) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResource gets the resource property value. Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *UsedInsight) GetResource()(Entityable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Entityable)
    }
    return nil
}
// GetResourceReference gets the resourceReference property value. Reference properties of the used document, such as the url and type of the document. Read-only
func (m *UsedInsight) GetResourceReference()(ResourceReferenceable) {
    val, err := m.GetBackingStore().Get("resourceReference")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ResourceReferenceable)
    }
    return nil
}
// GetResourceVisualization gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *UsedInsight) GetResourceVisualization()(ResourceVisualizationable) {
    val, err := m.GetBackingStore().Get("resourceVisualization")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ResourceVisualizationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UsedInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("lastUsed", m.GetLastUsed())
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
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastUsed sets the lastUsed property value. Information about when the item was last viewed or modified by the user. Read only.
func (m *UsedInsight) SetLastUsed(value UsageDetailsable)() {
    err := m.GetBackingStore().Set("lastUsed", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UsedInsight) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResource sets the resource property value. Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *UsedInsight) SetResource(value Entityable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceReference sets the resourceReference property value. Reference properties of the used document, such as the url and type of the document. Read-only
func (m *UsedInsight) SetResourceReference(value ResourceReferenceable)() {
    err := m.GetBackingStore().Set("resourceReference", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceVisualization sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *UsedInsight) SetResourceVisualization(value ResourceVisualizationable)() {
    err := m.GetBackingStore().Set("resourceVisualization", value)
    if err != nil {
        panic(err)
    }
}
// UsedInsightable 
type UsedInsightable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastUsed()(UsageDetailsable)
    GetOdataType()(*string)
    GetResource()(Entityable)
    GetResourceReference()(ResourceReferenceable)
    GetResourceVisualization()(ResourceVisualizationable)
    SetLastUsed(value UsageDetailsable)()
    SetOdataType(value *string)()
    SetResource(value Entityable)()
    SetResourceReference(value ResourceReferenceable)()
    SetResourceVisualization(value ResourceVisualizationable)()
}
