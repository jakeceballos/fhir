// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "encoding/json"
    "fmt"
    "bytes"
)

// RelatedArtifact Related artifacts such as additional documentation, justification, or bibliographic references.
type RelatedArtifact struct {

  // A bibliographic citation for the related artifact. This text SHOULD be formatted according to an accepted citation format.
  Citation string `json:"citation,omitempty"`

  // A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
  Display string `json:"display,omitempty"`

  // The document being referenced, represented as an attachment. This is exclusive with the resource element.
  Document *Attachment `json:"document,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // A short label that can be used to reference the citation from elsewhere in the containing artifact, such as a footnote index.
  Label string `json:"label,omitempty"`

  // The related resource, such as a library, value set, profile, or other knowledge resource.
  Resource string `json:"resource,omitempty"`

  // The type of relationship to the related artifact.
  Type interface{} `json:"type,omitempty"`

  // A url for the artifact that can be followed to access the actual content.
  Url string `json:"url,omitempty"`
}

func (strct *RelatedArtifact) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "citation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"citation\": ")
	if tmp, err := json.Marshal(strct.Citation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "display" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"display\": ")
	if tmp, err := json.Marshal(strct.Display); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "document" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"document\": ")
	if tmp, err := json.Marshal(strct.Document); err != nil {
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
    // Marshal the "label" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"label\": ")
	if tmp, err := json.Marshal(strct.Label); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "resource" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"resource\": ")
	if tmp, err := json.Marshal(strct.Resource); err != nil {
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
    // Marshal the "url" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RelatedArtifact) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "citation":
            if err := json.Unmarshal([]byte(v), &strct.Citation); err != nil {
                return err
             }
        case "display":
            if err := json.Unmarshal([]byte(v), &strct.Display); err != nil {
                return err
             }
        case "document":
            if err := json.Unmarshal([]byte(v), &strct.Document); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "label":
            if err := json.Unmarshal([]byte(v), &strct.Label); err != nil {
                return err
             }
        case "resource":
            if err := json.Unmarshal([]byte(v), &strct.Resource); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        case "url":
            if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
