// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// CompositionSection A set of healthcare-related information that is assembled together into a single logical package that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. A Composition defines the structure and narrative content necessary for a document. However, a Composition alone does not constitute a document. Rather, the Composition must be the first entry in a Bundle where Bundle.type=document, and any other resources referenced from Composition must be included as subsequent entries in the Bundle (for example Patient, Practitioner, Encounter, etc.).
type CompositionSection struct {

  // Identifies who is responsible for the information in this section, not necessarily who typed it in.
  Author []*Reference `json:"author,omitempty"`

  // A code identifying the kind of content contained within the section. This must be consistent with the section title.
  Code *CodeableConcept `json:"code,omitempty"`

  // If the section is empty, why the list is empty. An empty section typically has some text explaining the empty reason.
  EmptyReason *CodeableConcept `json:"emptyReason,omitempty"`

  // A reference to the actual resource from which the narrative in the section is derived.
  Entry []*Reference `json:"entry,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The actual focus of the section when it is not the subject of the composition, but instead represents something or someone associated with the subject such as (for a patient subject) a spouse, parent, fetus, or donor. If not focus is specified, the focus is assumed to be focus of the parent section, or, for a section in the Composition itself, the subject of the composition. Sections with a focus SHALL only include resources where the logical subject (patient, subject, focus, etc.) matches the section focus, or the resources have no logical subject (few resources).
  Focus *Reference `json:"focus,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // How the entry list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted.
  Mode string `json:"mode,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Specifies the order applied to the items in the section entries.
  OrderedBy *CodeableConcept `json:"orderedBy,omitempty"`

  // A nested sub-section within this section.
  Section []*CompositionSection `json:"section,omitempty"`

  // A human-readable narrative that contains the attested content of the section, used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative.
  Text *Narrative `json:"text,omitempty"`

  // The label for this particular section.  This will be part of the rendered content for the document, and is often used to build a table of contents.
  Title string `json:"title,omitempty"`
}

func (strct *CompositionSection) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "author" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"author\": ")
	if tmp, err := json.Marshal(strct.Author); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "code" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"code\": ")
	if tmp, err := json.Marshal(strct.Code); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "emptyReason" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"emptyReason\": ")
	if tmp, err := json.Marshal(strct.EmptyReason); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "entry" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"entry\": ")
	if tmp, err := json.Marshal(strct.Entry); err != nil {
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
    // Marshal the "focus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"focus\": ")
	if tmp, err := json.Marshal(strct.Focus); err != nil {
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
    // Marshal the "mode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"mode\": ")
	if tmp, err := json.Marshal(strct.Mode); err != nil {
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
    // Marshal the "orderedBy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"orderedBy\": ")
	if tmp, err := json.Marshal(strct.OrderedBy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "section" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"section\": ")
	if tmp, err := json.Marshal(strct.Section); err != nil {
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
    // Marshal the "title" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"title\": ")
	if tmp, err := json.Marshal(strct.Title); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CompositionSection) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "author":
            if err := json.Unmarshal([]byte(v), &strct.Author); err != nil {
                return err
             }
        case "code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
                return err
             }
        case "emptyReason":
            if err := json.Unmarshal([]byte(v), &strct.EmptyReason); err != nil {
                return err
             }
        case "entry":
            if err := json.Unmarshal([]byte(v), &strct.Entry); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "focus":
            if err := json.Unmarshal([]byte(v), &strct.Focus); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "mode":
            if err := json.Unmarshal([]byte(v), &strct.Mode); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "orderedBy":
            if err := json.Unmarshal([]byte(v), &strct.OrderedBy); err != nil {
                return err
             }
        case "section":
            if err := json.Unmarshal([]byte(v), &strct.Section); err != nil {
                return err
             }
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        case "title":
            if err := json.Unmarshal([]byte(v), &strct.Title); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
