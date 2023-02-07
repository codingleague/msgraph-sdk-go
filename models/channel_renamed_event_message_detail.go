package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelRenamedEventMessageDetail 
type ChannelRenamedEventMessageDetail struct {
    EventMessageDetail
    // The updated name of the channel.
    channelDisplayName *string
    // Unique identifier of the channel.
    channelId *string
    // Initiator of the event.
    initiator IdentitySetable
}
// NewChannelRenamedEventMessageDetail instantiates a new ChannelRenamedEventMessageDetail and sets the default values.
func NewChannelRenamedEventMessageDetail()(*ChannelRenamedEventMessageDetail) {
    m := &ChannelRenamedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.channelRenamedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateChannelRenamedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelRenamedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelRenamedEventMessageDetail(), nil
}
// GetChannelDisplayName gets the channelDisplayName property value. The updated name of the channel.
func (m *ChannelRenamedEventMessageDetail) GetChannelDisplayName()(*string) {
    return m.channelDisplayName
}
// GetChannelId gets the channelId property value. Unique identifier of the channel.
func (m *ChannelRenamedEventMessageDetail) GetChannelId()(*string) {
    return m.channelId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelRenamedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["channelDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelDisplayName(val)
        }
        return nil
    }
    res["channelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *ChannelRenamedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// Serialize serializes information the current object
func (m *ChannelRenamedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("channelDisplayName", m.GetChannelDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("channelId", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChannelDisplayName sets the channelDisplayName property value. The updated name of the channel.
func (m *ChannelRenamedEventMessageDetail) SetChannelDisplayName(value *string)() {
    m.channelDisplayName = value
}
// SetChannelId sets the channelId property value. Unique identifier of the channel.
func (m *ChannelRenamedEventMessageDetail) SetChannelId(value *string)() {
    m.channelId = value
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *ChannelRenamedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
