// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "errors"
    "bytes"
    "encoding/json"
    "fmt"
)

// StructureMapRule A Map of relationships between 2 structures that can be used to transform data.
type StructureMapRule struct {

  // Which other rules to apply in the context of this rule.
  Dependent []*StructureMapDependent `json:"dependent,omitempty"`

  // Documentation for this instance of data.
  Documentation string `json:"documentation,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Name of the rule for internal references.
  Name string `json:"name,omitempty"`

  // Rules contained in this rule.
  Rule []*StructureMapRule `json:"rule,omitempty"`

  // Source inputs to the mapping.
  Source []*StructureMapSource `json:"source"`

  // Content to create because of this mapping rule.
  Target []*StructureMapTarget `json:"target,omitempty"`
}

func (strct *StructureMapRule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "dependent" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dependent\": ")
	if tmp, err := json.Marshal(strct.Dependent); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "documentation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"documentation\": ")
	if tmp, err := json.Marshal(strct.Documentation); err != nil {
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
    // Marshal the "rule" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"rule\": ")
	if tmp, err := json.Marshal(strct.Rule); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Source" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "source" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"source\": ")
	if tmp, err := json.Marshal(strct.Source); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "target" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"target\": ")
	if tmp, err := json.Marshal(strct.Target); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *StructureMapRule) UnmarshalJSON(b []byte) error {
    sourceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "dependent":
            if err := json.Unmarshal([]byte(v), &strct.Dependent); err != nil {
                return err
             }
        case "documentation":
            if err := json.Unmarshal([]byte(v), &strct.Documentation); err != nil {
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
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "rule":
            if err := json.Unmarshal([]byte(v), &strct.Rule); err != nil {
                return err
             }
        case "source":
            if err := json.Unmarshal([]byte(v), &strct.Source); err != nil {
                return err
             }
            sourceReceived = true
        case "target":
            if err := json.Unmarshal([]byte(v), &strct.Target); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if source (a required property) was received
    if !sourceReceived {
        return errors.New("\"source\" is required but was not present")
    }
    return nil
}
