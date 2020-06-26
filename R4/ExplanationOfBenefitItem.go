// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// ExplanationOfBenefitItem This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefitItem struct {

  // If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
  Adjudication []*ExplanationOfBenefitAdjudication `json:"adjudication,omitempty"`

  // Physical service site on the patient (limb, tooth, etc.).
  BodySite *CodeableConcept `json:"bodySite,omitempty"`

  // Care team members related to this service or product.
  CareTeamSequence []float64 `json:"careTeamSequence,omitempty"`

  // Code to identify the general type of benefits under which products and services are provided.
  Category *CodeableConcept `json:"category,omitempty"`

  // Second-tier of goods and services.
  Detail []*ExplanationOfBenefitDetail `json:"detail,omitempty"`

  // Diagnoses applicable for this service or product.
  DiagnosisSequence []float64 `json:"diagnosisSequence,omitempty"`

  // A billed item may include goods or services provided in multiple encounters.
  Encounter []*Reference `json:"encounter,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
  Factor float64 `json:"factor,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Exceptions, special conditions and supporting information applicable for this service or product.
  InformationSequence []float64 `json:"informationSequence,omitempty"`

  // Where the product or service was provided.
  LocationAddress *Address `json:"locationAddress,omitempty"`

  // Where the product or service was provided.
  LocationCodeableConcept *CodeableConcept `json:"locationCodeableConcept,omitempty"`

  // Where the product or service was provided.
  LocationReference *Reference `json:"locationReference,omitempty"`

  // Item typification or modifiers codes to convey additional context for the product or service.
  Modifier []*CodeableConcept `json:"modifier,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The quantity times the unit price for an additional service or product or charge.
  Net *Money `json:"net,omitempty"`

  // The numbers associated with notes below which apply to the adjudication of this item.
  NoteNumber []float64 `json:"noteNumber,omitempty"`

  // Procedures applicable for this service or product.
  ProcedureSequence []float64 `json:"procedureSequence,omitempty"`

  // When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
  ProductOrService *CodeableConcept `json:"productOrService"`

  // Identifies the program under which this may be recovered.
  ProgramCode []*CodeableConcept `json:"programCode,omitempty"`

  // The number of repetitions of a service or product.
  Quantity *Quantity `json:"quantity,omitempty"`

  // The type of revenue or cost center providing the product and/or service.
  Revenue *CodeableConcept `json:"revenue,omitempty"`

  // A number to uniquely identify item entries.
  Sequence float64 `json:"sequence,omitempty"`

  // The date or dates when the service or product was supplied, performed or completed.
  ServicedDate string `json:"servicedDate,omitempty"`

  // The date or dates when the service or product was supplied, performed or completed.
  ServicedPeriod *Period `json:"servicedPeriod,omitempty"`

  // A region or surface of the bodySite, e.g. limb region or tooth surface(s).
  SubSite []*CodeableConcept `json:"subSite,omitempty"`

  // Unique Device Identifiers associated with this line item.
  Udi []*Reference `json:"udi,omitempty"`

  // If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
  UnitPrice *Money `json:"unitPrice,omitempty"`
}

func (strct *ExplanationOfBenefitItem) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "adjudication" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"adjudication\": ")
	if tmp, err := json.Marshal(strct.Adjudication); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "bodySite" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"bodySite\": ")
	if tmp, err := json.Marshal(strct.BodySite); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "careTeamSequence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"careTeamSequence\": ")
	if tmp, err := json.Marshal(strct.CareTeamSequence); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "category" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"category\": ")
	if tmp, err := json.Marshal(strct.Category); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "detail" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"detail\": ")
	if tmp, err := json.Marshal(strct.Detail); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "diagnosisSequence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"diagnosisSequence\": ")
	if tmp, err := json.Marshal(strct.DiagnosisSequence); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "encounter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"encounter\": ")
	if tmp, err := json.Marshal(strct.Encounter); err != nil {
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
    // Marshal the "factor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"factor\": ")
	if tmp, err := json.Marshal(strct.Factor); err != nil {
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
    // Marshal the "informationSequence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"informationSequence\": ")
	if tmp, err := json.Marshal(strct.InformationSequence); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "locationAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"locationAddress\": ")
	if tmp, err := json.Marshal(strct.LocationAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "locationCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"locationCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.LocationCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "locationReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"locationReference\": ")
	if tmp, err := json.Marshal(strct.LocationReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "modifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"modifier\": ")
	if tmp, err := json.Marshal(strct.Modifier); err != nil {
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
    // Marshal the "net" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"net\": ")
	if tmp, err := json.Marshal(strct.Net); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "noteNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"noteNumber\": ")
	if tmp, err := json.Marshal(strct.NoteNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "procedureSequence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"procedureSequence\": ")
	if tmp, err := json.Marshal(strct.ProcedureSequence); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ProductOrService" field is required
    if strct.ProductOrService == nil {
        return nil, errors.New("productOrService is a required field")
    }
    // Marshal the "productOrService" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"productOrService\": ")
	if tmp, err := json.Marshal(strct.ProductOrService); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "programCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"programCode\": ")
	if tmp, err := json.Marshal(strct.ProgramCode); err != nil {
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
    // Marshal the "revenue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"revenue\": ")
	if tmp, err := json.Marshal(strct.Revenue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "sequence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"sequence\": ")
	if tmp, err := json.Marshal(strct.Sequence); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "servicedDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"servicedDate\": ")
	if tmp, err := json.Marshal(strct.ServicedDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "servicedPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"servicedPeriod\": ")
	if tmp, err := json.Marshal(strct.ServicedPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "subSite" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subSite\": ")
	if tmp, err := json.Marshal(strct.SubSite); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "udi" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"udi\": ")
	if tmp, err := json.Marshal(strct.Udi); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "unitPrice" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"unitPrice\": ")
	if tmp, err := json.Marshal(strct.UnitPrice); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ExplanationOfBenefitItem) UnmarshalJSON(b []byte) error {
    productOrServiceReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "adjudication":
            if err := json.Unmarshal([]byte(v), &strct.Adjudication); err != nil {
                return err
             }
        case "bodySite":
            if err := json.Unmarshal([]byte(v), &strct.BodySite); err != nil {
                return err
             }
        case "careTeamSequence":
            if err := json.Unmarshal([]byte(v), &strct.CareTeamSequence); err != nil {
                return err
             }
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
        case "detail":
            if err := json.Unmarshal([]byte(v), &strct.Detail); err != nil {
                return err
             }
        case "diagnosisSequence":
            if err := json.Unmarshal([]byte(v), &strct.DiagnosisSequence); err != nil {
                return err
             }
        case "encounter":
            if err := json.Unmarshal([]byte(v), &strct.Encounter); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "factor":
            if err := json.Unmarshal([]byte(v), &strct.Factor); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "informationSequence":
            if err := json.Unmarshal([]byte(v), &strct.InformationSequence); err != nil {
                return err
             }
        case "locationAddress":
            if err := json.Unmarshal([]byte(v), &strct.LocationAddress); err != nil {
                return err
             }
        case "locationCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.LocationCodeableConcept); err != nil {
                return err
             }
        case "locationReference":
            if err := json.Unmarshal([]byte(v), &strct.LocationReference); err != nil {
                return err
             }
        case "modifier":
            if err := json.Unmarshal([]byte(v), &strct.Modifier); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "net":
            if err := json.Unmarshal([]byte(v), &strct.Net); err != nil {
                return err
             }
        case "noteNumber":
            if err := json.Unmarshal([]byte(v), &strct.NoteNumber); err != nil {
                return err
             }
        case "procedureSequence":
            if err := json.Unmarshal([]byte(v), &strct.ProcedureSequence); err != nil {
                return err
             }
        case "productOrService":
            if err := json.Unmarshal([]byte(v), &strct.ProductOrService); err != nil {
                return err
             }
            productOrServiceReceived = true
        case "programCode":
            if err := json.Unmarshal([]byte(v), &strct.ProgramCode); err != nil {
                return err
             }
        case "quantity":
            if err := json.Unmarshal([]byte(v), &strct.Quantity); err != nil {
                return err
             }
        case "revenue":
            if err := json.Unmarshal([]byte(v), &strct.Revenue); err != nil {
                return err
             }
        case "sequence":
            if err := json.Unmarshal([]byte(v), &strct.Sequence); err != nil {
                return err
             }
        case "servicedDate":
            if err := json.Unmarshal([]byte(v), &strct.ServicedDate); err != nil {
                return err
             }
        case "servicedPeriod":
            if err := json.Unmarshal([]byte(v), &strct.ServicedPeriod); err != nil {
                return err
             }
        case "subSite":
            if err := json.Unmarshal([]byte(v), &strct.SubSite); err != nil {
                return err
             }
        case "udi":
            if err := json.Unmarshal([]byte(v), &strct.Udi); err != nil {
                return err
             }
        case "unitPrice":
            if err := json.Unmarshal([]byte(v), &strct.UnitPrice); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if productOrService (a required property) was received
    if !productOrServiceReceived {
        return errors.New("\"productOrService\" is required but was not present")
    }
    return nil
}
