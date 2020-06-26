// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// ContractSigner Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type ContractSigner struct {

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Party which is a signator to this Contract.
  Party *Reference `json:"party"`

  // Legally binding Contract DSIG signature contents in Base64.
  Signature []*Signature `json:"signature"`

  // Role of this Contract signer, e.g. notary, grantee.
  Type *Coding `json:"type"`
}

func (strct *ContractSigner) MarshalJSON() ([]byte, error) {
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
    // "Party" field is required
    if strct.Party == nil {
        return nil, errors.New("party is a required field")
    }
    // Marshal the "party" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"party\": ")
	if tmp, err := json.Marshal(strct.Party); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Signature" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "signature" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"signature\": ")
	if tmp, err := json.Marshal(strct.Signature); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Type" field is required
    if strct.Type == nil {
        return nil, errors.New("type is a required field")
    }
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ContractSigner) UnmarshalJSON(b []byte) error {
    partyReceived := false
    signatureReceived := false
    typeReceived := false
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
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "party":
            if err := json.Unmarshal([]byte(v), &strct.Party); err != nil {
                return err
             }
            partyReceived = true
        case "signature":
            if err := json.Unmarshal([]byte(v), &strct.Signature); err != nil {
                return err
             }
            signatureReceived = true
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            typeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if party (a required property) was received
    if !partyReceived {
        return errors.New("\"party\" is required but was not present")
    }
    // check if signature (a required property) was received
    if !signatureReceived {
        return errors.New("\"signature\" is required but was not present")
    }
    // check if type (a required property) was received
    if !typeReceived {
        return errors.New("\"type\" is required but was not present")
    }
    return nil
}