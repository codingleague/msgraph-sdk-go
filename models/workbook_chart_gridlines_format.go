package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartGridlinesFormat 
type WorkbookChartGridlinesFormat struct {
    Entity
}
// NewWorkbookChartGridlinesFormat instantiates a new workbookChartGridlinesFormat and sets the default values.
func NewWorkbookChartGridlinesFormat()(*WorkbookChartGridlinesFormat) {
    m := &WorkbookChartGridlinesFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartGridlinesFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartGridlinesFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartGridlinesFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartGridlinesFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartLineFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val.(WorkbookChartLineFormatable))
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
    return res
}
// GetLine gets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartGridlinesFormat) GetLine()(WorkbookChartLineFormatable) {
    val, err := m.GetBackingStore().Get("line")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartLineFormatable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkbookChartGridlinesFormat) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookChartGridlinesFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("line", m.GetLine())
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
    return nil
}
// SetLine sets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartGridlinesFormat) SetLine(value WorkbookChartLineFormatable)() {
    err := m.GetBackingStore().Set("line", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookChartGridlinesFormat) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookChartGridlinesFormatable 
type WorkbookChartGridlinesFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLine()(WorkbookChartLineFormatable)
    GetOdataType()(*string)
    SetLine(value WorkbookChartLineFormatable)()
    SetOdataType(value *string)()
}
