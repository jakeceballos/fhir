// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// MedicationRequestDispenseRequest An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequestDispenseRequest struct {

  // The minimum period of time that must occur between dispenses of the medication.
  DispenseInterval *Duration `json:"dispenseInterval,omitempty"`

  // Identifies the period time over which the supplied product is expected to be used, or the length of time the dispense is expected to last.
  ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Indicates the quantity or duration for the first dispense of the medication.
  InitialFill *MedicationRequestInitialFill `json:"initialFill,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // An integer indicating the number of times, in addition to the original dispense, (aka refills or repeats) that the patient can receive the prescribed medication. Usage Notes: This integer does not include the original order dispense. This means that if an order indicates dispense 30 tablets plus "3 repeats", then the order can be dispensed a total of 4 times and the patient can receive a total of 120 tablets.  A prescriber may explicitly say that zero refills are permitted after the initial dispense.
  NumberOfRepeatsAllowed float64 `json:"numberOfRepeatsAllowed,omitempty"`

  // Indicates the intended dispensing Organization specified by the prescriber.
  Performer *Reference `json:"performer,omitempty"`

  // The amount that is to be dispensed for one fill.
  Quantity *Quantity `json:"quantity,omitempty"`

  // This indicates the validity period of a prescription (stale dating the Prescription).
  ValidityPeriod *Period `json:"validityPeriod,omitempty"`
}

func (strct *MedicationRequestDispenseRequest) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "dispenseInterval" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dispenseInterval\": ")
	if tmp, err := json.Marshal(strct.DispenseInterval); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "expectedSupplyDuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"expectedSupplyDuration\": ")
	if tmp, err := json.Marshal(strct.ExpectedSupplyDuration); err != nil {
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
    // Marshal the "initialFill" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"initialFill\": ")
	if tmp, err := json.Marshal(strct.InitialFill); err != nil {
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
    // Marshal the "numberOfRepeatsAllowed" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"numberOfRepeatsAllowed\": ")
	if tmp, err := json.Marshal(strct.NumberOfRepeatsAllowed); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "performer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"performer\": ")
	if tmp, err := json.Marshal(strct.Performer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "quantity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"quantity\": ")
	if tmp, err := json.Marshal(strct.Quantity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "validityPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"validityPeriod\": ")
	if tmp, err := json.Marshal(strct.ValidityPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MedicationRequestDispenseRequest) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "dispenseInterval":
            if err := json.Unmarshal([]byte(v), &strct.DispenseInterval); err != nil {
                return err
             }
        case "expectedSupplyDuration":
            if err := json.Unmarshal([]byte(v), &strct.ExpectedSupplyDuration); err != nil {
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
        case "initialFill":
            if err := json.Unmarshal([]byte(v), &strct.InitialFill); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "numberOfRepeatsAllowed":
            if err := json.Unmarshal([]byte(v), &strct.NumberOfRepeatsAllowed); err != nil {
                return err
             }
        case "performer":
            if err := json.Unmarshal([]byte(v), &strct.Performer); err != nil {
                return err
             }
        case "quantity":
            if err := json.Unmarshal([]byte(v), &strct.Quantity); err != nil {
                return err
             }
        case "validityPeriod":
            if err := json.Unmarshal([]byte(v), &strct.ValidityPeriod); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
