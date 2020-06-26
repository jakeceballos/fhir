// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// Media A photo, video, or audio recording acquired or used in healthcare. The actual content may be inline or provided by direct reference.
type Media struct {

  // A procedure that is fulfilled in whole or in part by the creation of this media.
  BasedOn []*Reference `json:"basedOn,omitempty"`

  // Indicates the site on the subject's body where the observation was made (i.e. the target site).
  BodySite *CodeableConcept `json:"bodySite,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The actual content of the media - inline or by direct reference to the media source file.
  Content *Attachment `json:"content"`

  // The date and time(s) at which the media was collected.
  CreatedDateTime string `json:"createdDateTime,omitempty"`

  // The date and time(s) at which the media was collected.
  CreatedPeriod *Period `json:"createdPeriod,omitempty"`

  // The device used to collect the media.
  Device *Reference `json:"device,omitempty"`

  // The name of the device / manufacturer of the device  that was used to make the recording.
  DeviceName string `json:"deviceName,omitempty"`

  // The duration of the recording in seconds - for audio and video.
  Duration float64 `json:"duration,omitempty"`

  // The encounter that establishes the context for this media.
  Encounter *Reference `json:"encounter,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The number of frames in a photo. This is used with a multi-page fax, or an imaging acquisition context that takes multiple slices in a single image, or an animated gif. If there is more than one frame, this SHALL have a value in order to alert interface software that a multi-frame capable rendering widget is required.
  Frames float64 `json:"frames,omitempty"`

  // Height of the image in pixels (photo/video).
  Height float64 `json:"height,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Identifiers associated with the image - these may include identifiers for the image itself, identifiers for the context of its collection (e.g. series ids) and context ids such as accession numbers or other workflow identifiers.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The date and time this version of the media was made available to providers, typically after having been reviewed.
  Issued string `json:"issued,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // Details of the type of the media - usually, how it was acquired (what type of device). If images sourced from a DICOM system, are wrapped in a Media resource, then this is the modality.
  Modality *CodeableConcept `json:"modality,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Comments made about the media by the performer, subject or other participants.
  Note []*Annotation `json:"note,omitempty"`

  // The person who administered the collection of the image.
  Operator *Reference `json:"operator,omitempty"`

  // A larger event of which this particular event is a component or step.
  PartOf []*Reference `json:"partOf,omitempty"`

  // Describes why the event occurred in coded or textual form.
  ReasonCode []*CodeableConcept `json:"reasonCode,omitempty"`

  // This is a Media resource
  ResourceType interface{} `json:"resourceType"`

  // The current state of the {{title}}.
  Status string `json:"status,omitempty"`

  // Who/What this Media is a record of.
  Subject *Reference `json:"subject,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // A code that classifies whether the media is an image, video or audio recording or some other media category.
  Type *CodeableConcept `json:"type,omitempty"`

  // The name of the imaging view e.g. Lateral or Antero-posterior (AP).
  View *CodeableConcept `json:"view,omitempty"`

  // Width of the image in pixels (photo/video).
  Width float64 `json:"width,omitempty"`
}

func (strct *Media) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "basedOn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"basedOn\": ")
	if tmp, err := json.Marshal(strct.BasedOn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "bodySite" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"bodySite\": ")
	if tmp, err := json.Marshal(strct.BodySite); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "contained" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contained\": ")
	if tmp, err := json.Marshal(strct.Contained); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Content" field is required
    if strct.Content == nil {
        return nil, errors.New("content is a required field")
    }
    // Marshal the "content" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"content\": ")
	if tmp, err := json.Marshal(strct.Content); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "createdDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"createdDateTime\": ")
	if tmp, err := json.Marshal(strct.CreatedDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "createdPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"createdPeriod\": ")
	if tmp, err := json.Marshal(strct.CreatedPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "device" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"device\": ")
	if tmp, err := json.Marshal(strct.Device); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "deviceName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"deviceName\": ")
	if tmp, err := json.Marshal(strct.DeviceName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "duration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"duration\": ")
	if tmp, err := json.Marshal(strct.Duration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "encounter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"encounter\": ")
	if tmp, err := json.Marshal(strct.Encounter); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "extension" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"extension\": ")
	if tmp, err := json.Marshal(strct.Extension); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "frames" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"frames\": ")
	if tmp, err := json.Marshal(strct.Frames); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "height" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"height\": ")
	if tmp, err := json.Marshal(strct.Height); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "id" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "identifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"identifier\": ")
	if tmp, err := json.Marshal(strct.Identifier); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "implicitRules" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"implicitRules\": ")
	if tmp, err := json.Marshal(strct.ImplicitRules); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "issued" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"issued\": ")
	if tmp, err := json.Marshal(strct.Issued); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "language" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"language\": ")
	if tmp, err := json.Marshal(strct.Language); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "meta" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"meta\": ")
	if tmp, err := json.Marshal(strct.Meta); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "modality" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"modality\": ")
	if tmp, err := json.Marshal(strct.Modality); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "modifierExtension" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"modifierExtension\": ")
	if tmp, err := json.Marshal(strct.ModifierExtension); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "note" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"note\": ")
	if tmp, err := json.Marshal(strct.Note); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "operator" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"operator\": ")
	if tmp, err := json.Marshal(strct.Operator); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "partOf" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"partOf\": ")
	if tmp, err := json.Marshal(strct.PartOf); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "reasonCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reasonCode\": ")
	if tmp, err := json.Marshal(strct.ReasonCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ResourceType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "resourceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"resourceType\": ")
	if tmp, err := json.Marshal(strct.ResourceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "status" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "subject" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subject\": ")
	if tmp, err := json.Marshal(strct.Subject); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "text" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"text\": ")
	if tmp, err := json.Marshal(strct.Text); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "type" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "view" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"view\": ")
	if tmp, err := json.Marshal(strct.View); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "width" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"width\": ")
	if tmp, err := json.Marshal(strct.Width); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Media) UnmarshalJSON(b []byte) error {
    contentReceived := false
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "basedOn":
            if err := json.Unmarshal([]byte(v), &strct.BasedOn); err != nil {
                return err
             }
        case "bodySite":
            if err := json.Unmarshal([]byte(v), &strct.BodySite); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "content":
            if err := json.Unmarshal([]byte(v), &strct.Content); err != nil {
                return err
             }
            contentReceived = true
        case "createdDateTime":
            if err := json.Unmarshal([]byte(v), &strct.CreatedDateTime); err != nil {
                return err
             }
        case "createdPeriod":
            if err := json.Unmarshal([]byte(v), &strct.CreatedPeriod); err != nil {
                return err
             }
        case "device":
            if err := json.Unmarshal([]byte(v), &strct.Device); err != nil {
                return err
             }
        case "deviceName":
            if err := json.Unmarshal([]byte(v), &strct.DeviceName); err != nil {
                return err
             }
        case "duration":
            if err := json.Unmarshal([]byte(v), &strct.Duration); err != nil {
                return err
             }
        case "encounter":
            if err := json.Unmarshal([]byte(v), &strct.Encounter); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "frames":
            if err := json.Unmarshal([]byte(v), &strct.Frames); err != nil {
                return err
             }
        case "height":
            if err := json.Unmarshal([]byte(v), &strct.Height); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "identifier":
            if err := json.Unmarshal([]byte(v), &strct.Identifier); err != nil {
                return err
             }
        case "implicitRules":
            if err := json.Unmarshal([]byte(v), &strct.ImplicitRules); err != nil {
                return err
             }
        case "issued":
            if err := json.Unmarshal([]byte(v), &strct.Issued); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "meta":
            if err := json.Unmarshal([]byte(v), &strct.Meta); err != nil {
                return err
             }
        case "modality":
            if err := json.Unmarshal([]byte(v), &strct.Modality); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "note":
            if err := json.Unmarshal([]byte(v), &strct.Note); err != nil {
                return err
             }
        case "operator":
            if err := json.Unmarshal([]byte(v), &strct.Operator); err != nil {
                return err
             }
        case "partOf":
            if err := json.Unmarshal([]byte(v), &strct.PartOf); err != nil {
                return err
             }
        case "reasonCode":
            if err := json.Unmarshal([]byte(v), &strct.ReasonCode); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "subject":
            if err := json.Unmarshal([]byte(v), &strct.Subject); err != nil {
                return err
             }
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        case "view":
            if err := json.Unmarshal([]byte(v), &strct.View); err != nil {
                return err
             }
        case "width":
            if err := json.Unmarshal([]byte(v), &strct.Width); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if content (a required property) was received
    if !contentReceived {
        return errors.New("\"content\" is required but was not present")
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    return nil
}