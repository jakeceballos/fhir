// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// TestScriptVariable A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScriptVariable struct {

  // A default, hard-coded, or user-defined value for this variable.
  DefaultValue string `json:"defaultValue,omitempty"`

  // A free text natural language description of the variable and its purpose.
  Description string `json:"description,omitempty"`

  // The FHIRPath expression to evaluate against the fixture body. When variables are defined, only one of either expression, headerField or path must be specified.
  Expression string `json:"expression,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Will be used to grab the HTTP header field value from the headers that sourceId is pointing to.
  HeaderField string `json:"headerField,omitempty"`

  // Displayable text string with hint help information to the user when entering a default value.
  Hint string `json:"hint,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Descriptive name for this variable.
  Name string `json:"name,omitempty"`

  // XPath or JSONPath to evaluate against the fixture body.  When variables are defined, only one of either expression, headerField or path must be specified.
  Path string `json:"path,omitempty"`

  // Fixture to evaluate the XPath/JSONPath expression or the headerField  against within this variable.
  SourceId string `json:"sourceId,omitempty"`
}

func (strct *TestScriptVariable) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "defaultValue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"defaultValue\": ")
	if tmp, err := json.Marshal(strct.DefaultValue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "description" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "expression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"expression\": ")
	if tmp, err := json.Marshal(strct.Expression); err != nil {
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
    // Marshal the "headerField" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"headerField\": ")
	if tmp, err := json.Marshal(strct.HeaderField); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "hint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"hint\": ")
	if tmp, err := json.Marshal(strct.Hint); err != nil {
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
    // Marshal the "name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "path" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"path\": ")
	if tmp, err := json.Marshal(strct.Path); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "sourceId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sourceId\": ")
	if tmp, err := json.Marshal(strct.SourceId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TestScriptVariable) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "defaultValue":
            if err := json.Unmarshal([]byte(v), &strct.DefaultValue); err != nil {
                return err
             }
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "expression":
            if err := json.Unmarshal([]byte(v), &strct.Expression); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "headerField":
            if err := json.Unmarshal([]byte(v), &strct.HeaderField); err != nil {
                return err
             }
        case "hint":
            if err := json.Unmarshal([]byte(v), &strct.Hint); err != nil {
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
        case "name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "path":
            if err := json.Unmarshal([]byte(v), &strct.Path); err != nil {
                return err
             }
        case "sourceId":
            if err := json.Unmarshal([]byte(v), &strct.SourceId); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
