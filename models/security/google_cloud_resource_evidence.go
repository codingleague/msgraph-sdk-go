package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GoogleCloudResourceEvidence 
type GoogleCloudResourceEvidence struct {
    AlertEvidence
}
// NewGoogleCloudResourceEvidence instantiates a new googleCloudResourceEvidence and sets the default values.
func NewGoogleCloudResourceEvidence()(*GoogleCloudResourceEvidence) {
    m := &GoogleCloudResourceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateGoogleCloudResourceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGoogleCloudResourceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGoogleCloudResourceEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GoogleCloudResourceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["locationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGoogleCloudLocationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationType(val.(*GoogleCloudLocationType))
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
    res["projectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProjectId(val)
        }
        return nil
    }
    res["projectNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProjectNumber(val)
        }
        return nil
    }
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. The zone or region where the resource is located.
func (m *GoogleCloudResourceEvidence) GetLocation()(*string) {
    val, err := m.GetBackingStore().Get("location")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocationType gets the locationType property value. The type of location. Possible values are: unknown, regional, zonal, global, unknownFutureValue.
func (m *GoogleCloudResourceEvidence) GetLocationType()(*GoogleCloudLocationType) {
    val, err := m.GetBackingStore().Get("locationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GoogleCloudLocationType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GoogleCloudResourceEvidence) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProjectId gets the projectId property value. The Google project ID as defined by the user.
func (m *GoogleCloudResourceEvidence) GetProjectId()(*string) {
    val, err := m.GetBackingStore().Get("projectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProjectNumber gets the projectNumber property value. The project number assigned by Google.
func (m *GoogleCloudResourceEvidence) GetProjectNumber()(*int64) {
    val, err := m.GetBackingStore().Get("projectNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetResourceName gets the resourceName property value. The name of the resource.
func (m *GoogleCloudResourceEvidence) GetResourceName()(*string) {
    val, err := m.GetBackingStore().Get("resourceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceType gets the resourceType property value. The type of the resource.
func (m *GoogleCloudResourceEvidence) GetResourceType()(*string) {
    val, err := m.GetBackingStore().Get("resourceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GoogleCloudResourceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    if m.GetLocationType() != nil {
        cast := (*m.GetLocationType()).String()
        err = writer.WriteStringValue("locationType", &cast)
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
        err = writer.WriteStringValue("projectId", m.GetProjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("projectNumber", m.GetProjectNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceType", m.GetResourceType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLocation sets the location property value. The zone or region where the resource is located.
func (m *GoogleCloudResourceEvidence) SetLocation(value *string)() {
    err := m.GetBackingStore().Set("location", value)
    if err != nil {
        panic(err)
    }
}
// SetLocationType sets the locationType property value. The type of location. Possible values are: unknown, regional, zonal, global, unknownFutureValue.
func (m *GoogleCloudResourceEvidence) SetLocationType(value *GoogleCloudLocationType)() {
    err := m.GetBackingStore().Set("locationType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GoogleCloudResourceEvidence) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProjectId sets the projectId property value. The Google project ID as defined by the user.
func (m *GoogleCloudResourceEvidence) SetProjectId(value *string)() {
    err := m.GetBackingStore().Set("projectId", value)
    if err != nil {
        panic(err)
    }
}
// SetProjectNumber sets the projectNumber property value. The project number assigned by Google.
func (m *GoogleCloudResourceEvidence) SetProjectNumber(value *int64)() {
    err := m.GetBackingStore().Set("projectNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceName sets the resourceName property value. The name of the resource.
func (m *GoogleCloudResourceEvidence) SetResourceName(value *string)() {
    err := m.GetBackingStore().Set("resourceName", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceType sets the resourceType property value. The type of the resource.
func (m *GoogleCloudResourceEvidence) SetResourceType(value *string)() {
    err := m.GetBackingStore().Set("resourceType", value)
    if err != nil {
        panic(err)
    }
}
// GoogleCloudResourceEvidenceable 
type GoogleCloudResourceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocation()(*string)
    GetLocationType()(*GoogleCloudLocationType)
    GetOdataType()(*string)
    GetProjectId()(*string)
    GetProjectNumber()(*int64)
    GetResourceName()(*string)
    GetResourceType()(*string)
    SetLocation(value *string)()
    SetLocationType(value *GoogleCloudLocationType)()
    SetOdataType(value *string)()
    SetProjectId(value *string)()
    SetProjectNumber(value *int64)()
    SetResourceName(value *string)()
    SetResourceType(value *string)()
}
