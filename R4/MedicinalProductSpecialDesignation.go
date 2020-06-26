// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// MedicinalProductSpecialDesignation Detailed definition of a medicinal product, typically for uses other than direct patient care (e.g. regulatory use).
type MedicinalProductSpecialDesignation struct {

  // Date when the designation was granted.
  Date string `json:"date,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Identifier for the designation, or procedure number.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // Condition for which the medicinal use applies.
  IndicationCodeableConcept *CodeableConcept `json:"indicationCodeableConcept,omitempty"`

  // Condition for which the medicinal use applies.
  IndicationReference *Reference `json:"indicationReference,omitempty"`

  // The intended use of the product, e.g. prevention, treatment.
  IntendedUse *CodeableConcept `json:"intendedUse,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Animal species for which this applies.
  Species *CodeableConcept `json:"species,omitempty"`

  // For example granted, pending, expired or withdrawn.
  Status *CodeableConcept `json:"status,omitempty"`

  // The type of special designation, e.g. orphan drug, minor use.
  Type *CodeableConcept `json:"type,omitempty"`
}

func (strct *MedicinalProductSpecialDesignation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "indicationCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"indicationCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.IndicationCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "indicationReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"indicationReference\": ")
	if tmp, err := json.Marshal(strct.IndicationReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "intendedUse" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"intendedUse\": ")
	if tmp, err := json.Marshal(strct.IntendedUse); err != nil {
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
    // Marshal the "species" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"species\": ")
	if tmp, err := json.Marshal(strct.Species); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "status" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
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

func (strct *MedicinalProductSpecialDesignation) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
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
        case "indicationCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.IndicationCodeableConcept); err != nil {
                return err
             }
        case "indicationReference":
            if err := json.Unmarshal([]byte(v), &strct.IndicationReference); err != nil {
                return err
             }
        case "intendedUse":
            if err := json.Unmarshal([]byte(v), &strct.IntendedUse); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "species":
            if err := json.Unmarshal([]byte(v), &strct.Species); err != nil {
                return err
             }
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
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
