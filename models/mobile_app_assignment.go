package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppAssignment a class containing the properties used for Group Assignment of a Mobile App.
type MobileAppAssignment struct {
    Entity
}
// NewMobileAppAssignment instantiates a new mobileAppAssignment and sets the default values.
func NewMobileAppAssignment()(*MobileAppAssignment) {
    m := &MobileAppAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*InstallIntent))
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
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppAssignmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(MobileAppAssignmentSettingsable))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetIntent gets the intent property value. Possible values for the install intent chosen by the admin.
func (m *MobileAppAssignment) GetIntent()(*InstallIntent) {
    val, err := m.GetBackingStore().Get("intent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*InstallIntent)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MobileAppAssignment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettings gets the settings property value. The settings for target assignment defined by the admin.
func (m *MobileAppAssignment) GetSettings()(MobileAppAssignmentSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MobileAppAssignmentSettingsable)
    }
    return nil
}
// GetTarget gets the target property value. The target group assignment defined by the admin.
func (m *MobileAppAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceAndAppManagementAssignmentTargetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
        err = writer.WriteStringValue("intent", &cast)
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
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIntent sets the intent property value. Possible values for the install intent chosen by the admin.
func (m *MobileAppAssignment) SetIntent(value *InstallIntent)() {
    err := m.GetBackingStore().Set("intent", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MobileAppAssignment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings for target assignment defined by the admin.
func (m *MobileAppAssignment) SetSettings(value MobileAppAssignmentSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The target group assignment defined by the admin.
func (m *MobileAppAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppAssignmentable 
type MobileAppAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntent()(*InstallIntent)
    GetOdataType()(*string)
    GetSettings()(MobileAppAssignmentSettingsable)
    GetTarget()(DeviceAndAppManagementAssignmentTargetable)
    SetIntent(value *InstallIntent)()
    SetOdataType(value *string)()
    SetSettings(value MobileAppAssignmentSettingsable)()
    SetTarget(value DeviceAndAppManagementAssignmentTargetable)()
}
