// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// CapabilityStatementRest A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server for a particular version of FHIR that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatementRest struct {

  // An absolute URI which is a reference to the definition of a compartment that the system supports. The reference is to a CompartmentDefinition resource by its canonical URL .
  Compartment []string `json:"compartment,omitempty"`

  // Information about the system's restful capabilities that apply across all applications, such as security.
  Documentation string `json:"documentation,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // A specification of restful operations supported by the system.
  Interaction []*CapabilityStatementInteraction1 `json:"interaction,omitempty"`

  // Identifies whether this portion of the statement is describing the ability to initiate or receive restful operations.
  Mode interface{} `json:"mode,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Definition of an operation or a named query together with its parameters and their meaning and type.
  Operation []*CapabilityStatementOperation `json:"operation,omitempty"`

  // A specification of the restful capabilities of the solution for a specific resource type.
  Resource []*CapabilityStatementResource `json:"resource,omitempty"`

  // Search parameters that are supported for searching all resources for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
  SearchParam []*CapabilityStatementSearchParam `json:"searchParam,omitempty"`

  // Information about security implementation from an interface perspective - what a client needs to know.
  Security *CapabilityStatementSecurity `json:"security,omitempty"`
}

func (strct *CapabilityStatementRest) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "compartment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"compartment\": ")
	if tmp, err := json.Marshal(strct.Compartment); err != nil {
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
    // Marshal the "interaction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"interaction\": ")
	if tmp, err := json.Marshal(strct.Interaction); err != nil {
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
    // Marshal the "operation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"operation\": ")
	if tmp, err := json.Marshal(strct.Operation); err != nil {
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
    // Marshal the "searchParam" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"searchParam\": ")
	if tmp, err := json.Marshal(strct.SearchParam); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "security" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"security\": ")
	if tmp, err := json.Marshal(strct.Security); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CapabilityStatementRest) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "compartment":
            if err := json.Unmarshal([]byte(v), &strct.Compartment); err != nil {
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
        case "interaction":
            if err := json.Unmarshal([]byte(v), &strct.Interaction); err != nil {
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
        case "operation":
            if err := json.Unmarshal([]byte(v), &strct.Operation); err != nil {
                return err
             }
        case "resource":
            if err := json.Unmarshal([]byte(v), &strct.Resource); err != nil {
                return err
             }
        case "searchParam":
            if err := json.Unmarshal([]byte(v), &strct.SearchParam); err != nil {
                return err
             }
        case "security":
            if err := json.Unmarshal([]byte(v), &strct.Security); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}