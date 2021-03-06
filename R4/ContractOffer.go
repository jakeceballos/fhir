// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// ContractOffer Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type ContractOffer struct {

  // Response to offer text.
  Answer []*ContractAnswer `json:"answer,omitempty"`

  // Type of choice made by accepting party with respect to an offer made by an offeror/ grantee.
  Decision *CodeableConcept `json:"decision,omitempty"`

  // How the decision about a Contract was conveyed.
  DecisionMode []*CodeableConcept `json:"decisionMode,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Unique identifier for this particular Contract Provision.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // The id of the clause or question text of the offer in the referenced questionnaire/response.
  LinkId []string `json:"linkId,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Offer Recipient.
  Party []*ContractParty `json:"party,omitempty"`

  // Security labels that protects the offer.
  SecurityLabelNumber []float64 `json:"securityLabelNumber,omitempty"`

  // Human readable form of this Contract Offer.
  Text string `json:"text,omitempty"`

  // The owner of an asset has the residual control rights over the asset: the right to decide all usages of the asset in any way not inconsistent with a prior contract, custom, or law (Hart, 1995, p. 30).
  Topic *Reference `json:"topic,omitempty"`

  // Type of Contract Provision such as specific requirements, purposes for actions, obligations, prohibitions, e.g. life time maximum benefit.
  Type *CodeableConcept `json:"type,omitempty"`
}

func (strct *ContractOffer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "answer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"answer\": ")
	if tmp, err := json.Marshal(strct.Answer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "decision" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"decision\": ")
	if tmp, err := json.Marshal(strct.Decision); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "decisionMode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"decisionMode\": ")
	if tmp, err := json.Marshal(strct.DecisionMode); err != nil {
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
    // Marshal the "linkId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"linkId\": ")
	if tmp, err := json.Marshal(strct.LinkId); err != nil {
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
    // Marshal the "securityLabelNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"securityLabelNumber\": ")
	if tmp, err := json.Marshal(strct.SecurityLabelNumber); err != nil {
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
    // Marshal the "topic" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"topic\": ")
	if tmp, err := json.Marshal(strct.Topic); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ContractOffer) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "answer":
            if err := json.Unmarshal([]byte(v), &strct.Answer); err != nil {
                return err
             }
        case "decision":
            if err := json.Unmarshal([]byte(v), &strct.Decision); err != nil {
                return err
             }
        case "decisionMode":
            if err := json.Unmarshal([]byte(v), &strct.DecisionMode); err != nil {
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
        case "identifier":
            if err := json.Unmarshal([]byte(v), &strct.Identifier); err != nil {
                return err
             }
        case "linkId":
            if err := json.Unmarshal([]byte(v), &strct.LinkId); err != nil {
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
        case "securityLabelNumber":
            if err := json.Unmarshal([]byte(v), &strct.SecurityLabelNumber); err != nil {
                return err
             }
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        case "topic":
            if err := json.Unmarshal([]byte(v), &strct.Topic); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
