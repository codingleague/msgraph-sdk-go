package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterCreateOperation 
type PrinterCreateOperation struct {
    PrintOperation
}
// NewPrinterCreateOperation instantiates a new printerCreateOperation and sets the default values.
func NewPrinterCreateOperation()(*PrinterCreateOperation) {
    m := &PrinterCreateOperation{
        PrintOperation: *NewPrintOperation(),
    }
    odataTypeValue := "#microsoft.graph.printerCreateOperation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrinterCreateOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterCreateOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterCreateOperation(), nil
}
// GetCertificate gets the certificate property value. The signed certificate created during the registration process. Read-only.
func (m *PrinterCreateOperation) GetCertificate()(*string) {
    val, err := m.GetBackingStore().Get("certificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterCreateOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrintOperation.GetFieldDeserializers()
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
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
    res["printer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinter(val.(Printerable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterCreateOperation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrinter gets the printer property value. The created printer entity. Read-only.
func (m *PrinterCreateOperation) GetPrinter()(Printerable) {
    val, err := m.GetBackingStore().Get("printer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Printerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrinterCreateOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrintOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificate", m.GetCertificate())
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
        err = writer.WriteObjectValue("printer", m.GetPrinter())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificate sets the certificate property value. The signed certificate created during the registration process. Read-only.
func (m *PrinterCreateOperation) SetCertificate(value *string)() {
    err := m.GetBackingStore().Set("certificate", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterCreateOperation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrinter sets the printer property value. The created printer entity. Read-only.
func (m *PrinterCreateOperation) SetPrinter(value Printerable)() {
    err := m.GetBackingStore().Set("printer", value)
    if err != nil {
        panic(err)
    }
}
// PrinterCreateOperationable 
type PrinterCreateOperationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrintOperationable
    GetCertificate()(*string)
    GetOdataType()(*string)
    GetPrinter()(Printerable)
    SetCertificate(value *string)()
    SetOdataType(value *string)()
    SetPrinter(value Printerable)()
}
