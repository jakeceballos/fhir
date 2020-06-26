// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "bytes"
    "encoding/json"
)

// NutritionOrderEnteralFormula A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrderEnteralFormula struct {

  // The product or brand name of the type of modular component to be added to the formula.
  AdditiveProductName string `json:"additiveProductName,omitempty"`

  // Indicates the type of modular component such as protein, carbohydrate, fat or fiber to be provided in addition to or mixed with the base formula.
  AdditiveType *CodeableConcept `json:"additiveType,omitempty"`

  // Formula administration instructions as structured data.  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
  Administration []*NutritionOrderAdministration `json:"administration,omitempty"`

  // Free text formula administration, feeding instructions or additional instructions or information.
  AdministrationInstruction string `json:"administrationInstruction,omitempty"`

  // The product or brand name of the enteral or infant formula product such as "ACME Adult Standard Formula".
  BaseFormulaProductName string `json:"baseFormulaProductName,omitempty"`

  // The type of enteral or infant formula such as an adult standard formula with fiber or a soy-based infant formula.
  BaseFormulaType *CodeableConcept `json:"baseFormulaType,omitempty"`

  // The amount of energy (calories) that the formula should provide per specified volume, typically per mL or fluid oz.  For example, an infant may require a formula that provides 24 calories per fluid ounce or an adult may require an enteral formula that provides 1.5 calorie/mL.
  CaloricDensity *Quantity `json:"caloricDensity,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // The maximum total quantity of formula that may be administered to a subject over the period of time, e.g. 1440 mL over 24 hours.
  MaxVolumeToDeliver *Quantity `json:"maxVolumeToDeliver,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The route or physiological path of administration into the patient's gastrointestinal  tract for purposes of providing the formula feeding, e.g. nasogastric tube.
  RouteofAdministration *CodeableConcept `json:"routeofAdministration,omitempty"`
}

func (strct *NutritionOrderEnteralFormula) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "additiveProductName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"additiveProductName\": ")
	if tmp, err := json.Marshal(strct.AdditiveProductName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "additiveType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"additiveType\": ")
	if tmp, err := json.Marshal(strct.AdditiveType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "administration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"administration\": ")
	if tmp, err := json.Marshal(strct.Administration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "administrationInstruction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"administrationInstruction\": ")
	if tmp, err := json.Marshal(strct.AdministrationInstruction); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "baseFormulaProductName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"baseFormulaProductName\": ")
	if tmp, err := json.Marshal(strct.BaseFormulaProductName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "baseFormulaType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"baseFormulaType\": ")
	if tmp, err := json.Marshal(strct.BaseFormulaType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "caloricDensity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"caloricDensity\": ")
	if tmp, err := json.Marshal(strct.CaloricDensity); err != nil {
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
    // Marshal the "maxVolumeToDeliver" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"maxVolumeToDeliver\": ")
	if tmp, err := json.Marshal(strct.MaxVolumeToDeliver); err != nil {
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
    // Marshal the "routeofAdministration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"routeofAdministration\": ")
	if tmp, err := json.Marshal(strct.RouteofAdministration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *NutritionOrderEnteralFormula) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "additiveProductName":
            if err := json.Unmarshal([]byte(v), &strct.AdditiveProductName); err != nil {
                return err
             }
        case "additiveType":
            if err := json.Unmarshal([]byte(v), &strct.AdditiveType); err != nil {
                return err
             }
        case "administration":
            if err := json.Unmarshal([]byte(v), &strct.Administration); err != nil {
                return err
             }
        case "administrationInstruction":
            if err := json.Unmarshal([]byte(v), &strct.AdministrationInstruction); err != nil {
                return err
             }
        case "baseFormulaProductName":
            if err := json.Unmarshal([]byte(v), &strct.BaseFormulaProductName); err != nil {
                return err
             }
        case "baseFormulaType":
            if err := json.Unmarshal([]byte(v), &strct.BaseFormulaType); err != nil {
                return err
             }
        case "caloricDensity":
            if err := json.Unmarshal([]byte(v), &strct.CaloricDensity); err != nil {
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
        case "maxVolumeToDeliver":
            if err := json.Unmarshal([]byte(v), &strct.MaxVolumeToDeliver); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "routeofAdministration":
            if err := json.Unmarshal([]byte(v), &strct.RouteofAdministration); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}