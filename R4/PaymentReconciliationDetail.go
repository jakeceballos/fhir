// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// PaymentReconciliationDetail This resource provides the details including amount of a payment and allocates the payment items being paid.
type PaymentReconciliationDetail struct {

  // The monetary amount allocated from the total payment to the payable.
  Amount *Money `json:"amount,omitempty"`

  // The date from the response resource containing a commitment to pay.
  Date string `json:"date,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Unique identifier for the current payment item for the referenced payable.
  Identifier *Identifier `json:"identifier,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The party which is receiving the payment.
  Payee *Reference `json:"payee,omitempty"`

  // Unique identifier for the prior payment item for the referenced payable.
  Predecessor *Identifier `json:"predecessor,omitempty"`

  // A resource, such as a Claim, the evaluation of which could lead to payment.
  Request *Reference `json:"request,omitempty"`

  // A resource, such as a ClaimResponse, which contains a commitment to payment.
  Response *Reference `json:"response,omitempty"`

  // A reference to the individual who is responsible for inquiries regarding the response and its payment.
  Responsible *Reference `json:"responsible,omitempty"`

  // The party which submitted the claim or financial transaction.
  Submitter *Reference `json:"submitter,omitempty"`

  // Code to indicate the nature of the payment.
  Type *CodeableConcept `json:"type"`
}

func (strct *PaymentReconciliationDetail) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "amount" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"amount\": ")
	if tmp, err := json.Marshal(strct.Amount); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "date" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"date\": ")
	if tmp, err := json.Marshal(strct.Date); err != nil {
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
    // Marshal the "payee" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"payee\": ")
	if tmp, err := json.Marshal(strct.Payee); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "predecessor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"predecessor\": ")
	if tmp, err := json.Marshal(strct.Predecessor); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "request" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"request\": ")
	if tmp, err := json.Marshal(strct.Request); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "response" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"response\": ")
	if tmp, err := json.Marshal(strct.Response); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "responsible" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"responsible\": ")
	if tmp, err := json.Marshal(strct.Responsible); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "submitter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"submitter\": ")
	if tmp, err := json.Marshal(strct.Submitter); err != nil {
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

func (strct *PaymentReconciliationDetail) UnmarshalJSON(b []byte) error {
    typeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "amount":
            if err := json.Unmarshal([]byte(v), &strct.Amount); err != nil {
                return err
             }
        case "date":
            if err := json.Unmarshal([]byte(v), &strct.Date); err != nil {
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
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "payee":
            if err := json.Unmarshal([]byte(v), &strct.Payee); err != nil {
                return err
             }
        case "predecessor":
            if err := json.Unmarshal([]byte(v), &strct.Predecessor); err != nil {
                return err
             }
        case "request":
            if err := json.Unmarshal([]byte(v), &strct.Request); err != nil {
                return err
             }
        case "response":
            if err := json.Unmarshal([]byte(v), &strct.Response); err != nil {
                return err
             }
        case "responsible":
            if err := json.Unmarshal([]byte(v), &strct.Responsible); err != nil {
                return err
             }
        case "submitter":
            if err := json.Unmarshal([]byte(v), &strct.Submitter); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            typeReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if type (a required property) was received
    if !typeReceived {
        return errors.New("\"type\" is required but was not present")
    }
    return nil
}
