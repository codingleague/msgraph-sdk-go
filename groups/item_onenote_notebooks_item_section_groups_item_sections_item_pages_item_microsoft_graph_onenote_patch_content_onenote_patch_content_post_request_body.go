package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody 
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commands property
    commands []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable
}
// NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody instantiates a new ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody and sets the default values.
func NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody()(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody) {
    m := &ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommands gets the commands property value. The commands property
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody) GetCommands()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable) {
    return m.commands
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commands"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePatchContentCommandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable)
            }
            m.SetCommands(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCommands() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("commands", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommands sets the commands property value. The commands property
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBody) SetCommands(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable)() {
    m.commands = value
}
