// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "encoding/json"
    "fmt"
    "bytes"
)

// ImplementationGuidePage A set of rules of how a particular interoperability or standards problem is solved - typically through the use of FHIR resources. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuidePage struct {

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // A code that indicates how the page is generated.
  Generation interface{} `json:"generation,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The source address for the page.
  NameReference *Reference `json:"nameReference,omitempty"`

  // The source address for the page.
  NameUrl string `json:"nameUrl,omitempty"`

  // Nested Pages/Sections under this page.
  Page []*ImplementationGuidePage `json:"page,omitempty"`

  // A short title used to represent this page in navigational structures such as table of contents, bread crumbs, etc.
  Title string `json:"title,omitempty"`
}

func (strct *ImplementationGuidePage) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "generation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"generation\": ")
	if tmp, err := json.Marshal(strct.Generation); err != nil {
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
    // Marshal the "nameReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"nameReference\": ")
	if tmp, err := json.Marshal(strct.NameReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "nameUrl" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"nameUrl\": ")
	if tmp, err := json.Marshal(strct.NameUrl); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "page" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"page\": ")
	if tmp, err := json.Marshal(strct.Page); err != nil {
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

func (strct *ImplementationGuidePage) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "generation":
            if err := json.Unmarshal([]byte(v), &strct.Generation); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "nameReference":
            if err := json.Unmarshal([]byte(v), &strct.NameReference); err != nil {
                return err
             }
        case "nameUrl":
            if err := json.Unmarshal([]byte(v), &strct.NameUrl); err != nil {
                return err
             }
        case "page":
            if err := json.Unmarshal([]byte(v), &strct.Page); err != nil {
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
