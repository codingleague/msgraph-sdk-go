package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionStageSetting 
type CustomExtensionStageSetting struct {
    Entity
}
// NewCustomExtensionStageSetting instantiates a new customExtensionStageSetting and sets the default values.
func NewCustomExtensionStageSetting()(*CustomExtensionStageSetting) {
    m := &CustomExtensionStageSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomExtensionStageSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionStageSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionStageSetting(), nil
}
// GetCustomExtension gets the customExtension property value. The customExtension property
func (m *CustomExtensionStageSetting) GetCustomExtension()(CustomCalloutExtensionable) {
    val, err := m.GetBackingStore().Get("customExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomCalloutExtensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionStageSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomCalloutExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(CustomCalloutExtensionable))
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
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*AccessPackageCustomExtensionStage))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomExtensionStageSetting) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStage gets the stage property value. The stage property
func (m *CustomExtensionStageSetting) GetStage()(*AccessPackageCustomExtensionStage) {
    val, err := m.GetBackingStore().Get("stage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AccessPackageCustomExtensionStage)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomExtensionStageSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
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
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err = writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtension sets the customExtension property value. The customExtension property
func (m *CustomExtensionStageSetting) SetCustomExtension(value CustomCalloutExtensionable)() {
    err := m.GetBackingStore().Set("customExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionStageSetting) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStage sets the stage property value. The stage property
func (m *CustomExtensionStageSetting) SetStage(value *AccessPackageCustomExtensionStage)() {
    err := m.GetBackingStore().Set("stage", value)
    if err != nil {
        panic(err)
    }
}
// CustomExtensionStageSettingable 
type CustomExtensionStageSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomExtension()(CustomCalloutExtensionable)
    GetOdataType()(*string)
    GetStage()(*AccessPackageCustomExtensionStage)
    SetCustomExtension(value CustomCalloutExtensionable)()
    SetOdataType(value *string)()
    SetStage(value *AccessPackageCustomExtensionStage)()
}
