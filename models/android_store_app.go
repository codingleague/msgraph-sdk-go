package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidStoreApp contains properties and inherited properties for Android store apps.
type AndroidStoreApp struct {
    MobileApp
    // The OdataType property
    OdataType *string
}
// NewAndroidStoreApp instantiates a new androidStoreApp and sets the default values.
func NewAndroidStoreApp()(*AndroidStoreApp) {
    m := &AndroidStoreApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.androidStoreApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidStoreApp(), nil
}
// GetAppStoreUrl gets the appStoreUrl property value. The Android app store URL.
func (m *AndroidStoreApp) GetAppStoreUrl()(*string) {
    val, err := m.GetBackingStore().Get("appStoreUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(AndroidMinimumOperatingSystemable))
        }
        return nil
    }
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
        }
        return nil
    }
    return res
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *AndroidStoreApp) GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable) {
    val, err := m.GetBackingStore().Get("minimumSupportedOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidMinimumOperatingSystemable)
    }
    return nil
}
// GetPackageId gets the packageId property value. The package identifier.
func (m *AndroidStoreApp) GetPackageId()(*string) {
    val, err := m.GetBackingStore().Get("packageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidStoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageId", m.GetPackageId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppStoreUrl sets the appStoreUrl property value. The Android app store URL.
func (m *AndroidStoreApp) SetAppStoreUrl(value *string)() {
    err := m.GetBackingStore().Set("appStoreUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *AndroidStoreApp) SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)() {
    err := m.GetBackingStore().Set("minimumSupportedOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *AndroidStoreApp) SetPackageId(value *string)() {
    err := m.GetBackingStore().Set("packageId", value)
    if err != nil {
        panic(err)
    }
}
// AndroidStoreAppable 
type AndroidStoreAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppStoreUrl()(*string)
    GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable)
    GetPackageId()(*string)
    SetAppStoreUrl(value *string)()
    SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)()
    SetPackageId(value *string)()
}
