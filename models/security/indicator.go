package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Indicator 
type Indicator struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
}
// NewIndicator instantiates a new indicator and sets the default values.
func NewIndicator()(*Indicator) {
    m := &Indicator{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateIndicatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIndicatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.articleIndicator":
                        return NewArticleIndicator(), nil
                    case "#microsoft.graph.security.intelligenceProfileIndicator":
                        return NewIntelligenceProfileIndicator(), nil
                }
            }
        }
    }
    return NewIndicator(), nil
}
// GetArtifact gets the artifact property value. The artifact property
func (m *Indicator) GetArtifact()(Artifactable) {
    val, err := m.GetBackingStore().Get("artifact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Artifactable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Indicator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["artifact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateArtifactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifact(val.(Artifactable))
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIndicatorSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*IndicatorSource))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Indicator) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSource gets the source property value. The source property
func (m *Indicator) GetSource()(*IndicatorSource) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IndicatorSource)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Indicator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("artifact", m.GetArtifact())
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
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArtifact sets the artifact property value. The artifact property
func (m *Indicator) SetArtifact(value Artifactable)() {
    err := m.GetBackingStore().Set("artifact", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Indicator) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. The source property
func (m *Indicator) SetSource(value *IndicatorSource)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
// Indicatorable 
type Indicatorable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArtifact()(Artifactable)
    GetOdataType()(*string)
    GetSource()(*IndicatorSource)
    SetArtifact(value Artifactable)()
    SetOdataType(value *string)()
    SetSource(value *IndicatorSource)()
}
