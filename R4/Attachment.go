// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// Attachment For referring to data content defined in other formats.
type Attachment struct {

  // Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
  ContentType string `json:"contentType,omitempty"`

  // The date that the attachment was first created.
  Creation string `json:"creation,omitempty"`

  // The actual data of the attachment - a sequence of bytes, base64 encoded.
  Data string `json:"data,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The calculated hash of the data using SHA-1. Represented using base64.
  Hash string `json:"hash,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // The human language of the content. The value can be any valid value according to BCP 47.
  Language string `json:"language,omitempty"`

  // The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
  Size float64 `json:"size,omitempty"`

  // A label or set of text to display in place of the data.
  Title string `json:"title,omitempty"`

  // A location where the data can be accessed.
  Url string `json:"url,omitempty"`
}

func (strct *Attachment) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "contentType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contentType\": ")
	if tmp, err := json.Marshal(strct.ContentType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "creation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"creation\": ")
	if tmp, err := json.Marshal(strct.Creation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "data" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"data\": ")
	if tmp, err := json.Marshal(strct.Data); err != nil {
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
    // Marshal the "hash" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"hash\": ")
	if tmp, err := json.Marshal(strct.Hash); err != nil {
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
    // Marshal the "size" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"size\": ")
	if tmp, err := json.Marshal(strct.Size); err != nil {
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

func (strct *Attachment) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "contentType":
            if err := json.Unmarshal([]byte(v), &strct.ContentType); err != nil {
                return err
             }
        case "creation":
            if err := json.Unmarshal([]byte(v), &strct.Creation); err != nil {
                return err
             }
        case "data":
            if err := json.Unmarshal([]byte(v), &strct.Data); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "hash":
            if err := json.Unmarshal([]byte(v), &strct.Hash); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "size":
            if err := json.Unmarshal([]byte(v), &strct.Size); err != nil {
                return err
             }
        case "title":
            if err := json.Unmarshal([]byte(v), &strct.Title); err != nil {
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
