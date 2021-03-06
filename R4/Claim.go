// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "errors"
    "encoding/json"
    "fmt"
    "bytes"
)

// Claim A provider issued list of professional services and products which have been provided, or are to be provided, to a patient which is sent to an insurer for reimbursement.
type Claim struct {

  // Details of an accident which resulted in injuries which required the products and services listed in the claim.
  Accident *ClaimAccident `json:"accident,omitempty"`

  // The period for which charges are being submitted.
  BillablePeriod *Period `json:"billablePeriod,omitempty"`

  // The members of the team who provided the products and services.
  CareTeam []*ClaimCareTeam `json:"careTeam,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The date this resource was created.
  Created string `json:"created,omitempty"`

  // Information about diagnoses relevant to the claim items.
  Diagnosis []*ClaimDiagnosis `json:"diagnosis,omitempty"`

  // Individual who created the claim, predetermination or preauthorization.
  Enterer *Reference `json:"enterer,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Facility where the services were provided.
  Facility *Reference `json:"facility,omitempty"`

  // A code to indicate whether and for whom funds are to be reserved for future claims.
  FundsReserve *CodeableConcept `json:"fundsReserve,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // A unique identifier assigned to this claim.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // Financial instruments for reimbursement for the health care products and services specified on the claim.
  Insurance []*ClaimInsurance `json:"insurance"`

  // The Insurer who is target of the request.
  Insurer *Reference `json:"insurer,omitempty"`

  // A claim line. Either a simple  product or service or a 'group' of details which can each be a simple items or groups of sub-details.
  Item []*ClaimItem `json:"item,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Original prescription which has been superseded by this prescription to support the dispensing of pharmacy services, medications or products.
  OriginalPrescription *Reference `json:"originalPrescription,omitempty"`

  // The party to whom the professional services and/or products have been supplied or are being considered and for whom actual or forecast reimbursement is sought.
  Patient *Reference `json:"patient"`

  // The party to be reimbursed for cost of the products and services according to the terms of the policy.
  Payee *ClaimPayee `json:"payee,omitempty"`

  // Prescription to support the dispensing of pharmacy, device or vision products.
  Prescription *Reference `json:"prescription,omitempty"`

  // The provider-required urgency of processing the request. Typical values include: stat, routine deferred.
  Priority *CodeableConcept `json:"priority"`

  // Procedures performed on the patient relevant to the billing items with the claim.
  Procedure []*ClaimProcedure `json:"procedure,omitempty"`

  // The provider which is responsible for the claim, predetermination or preauthorization.
  Provider *Reference `json:"provider"`

  // A reference to a referral resource.
  Referral *Reference `json:"referral,omitempty"`

  // Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
  Related []*ClaimRelated `json:"related,omitempty"`

  // This is a Claim resource
  ResourceType interface{} `json:"resourceType"`

  // The status of the resource instance.
  Status string `json:"status,omitempty"`

  // A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
  SubType *CodeableConcept `json:"subType,omitempty"`

  // Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
  SupportingInfo []*ClaimSupportingInfo `json:"supportingInfo,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // The total value of the all the items in the claim.
  Total *Money `json:"total,omitempty"`

  // The category of claim, e.g. oral, pharmacy, vision, institutional, professional.
  Type *CodeableConcept `json:"type"`

  // A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
  Use interface{} `json:"use,omitempty"`
}

func (strct *Claim) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "accident" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"accident\": ")
	if tmp, err := json.Marshal(strct.Accident); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "billablePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"billablePeriod\": ")
	if tmp, err := json.Marshal(strct.BillablePeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "careTeam" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"careTeam\": ")
	if tmp, err := json.Marshal(strct.CareTeam); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "contained" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"contained\": ")
	if tmp, err := json.Marshal(strct.Contained); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "created" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"created\": ")
	if tmp, err := json.Marshal(strct.Created); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "diagnosis" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"diagnosis\": ")
	if tmp, err := json.Marshal(strct.Diagnosis); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "enterer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"enterer\": ")
	if tmp, err := json.Marshal(strct.Enterer); err != nil {
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
    // Marshal the "facility" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"facility\": ")
	if tmp, err := json.Marshal(strct.Facility); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "fundsReserve" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"fundsReserve\": ")
	if tmp, err := json.Marshal(strct.FundsReserve); err != nil {
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
    // Marshal the "implicitRules" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"implicitRules\": ")
	if tmp, err := json.Marshal(strct.ImplicitRules); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Insurance" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "insurance" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"insurance\": ")
	if tmp, err := json.Marshal(strct.Insurance); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "insurer" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"insurer\": ")
	if tmp, err := json.Marshal(strct.Insurer); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "item" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"item\": ")
	if tmp, err := json.Marshal(strct.Item); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "language" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"language\": ")
	if tmp, err := json.Marshal(strct.Language); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "meta" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"meta\": ")
	if tmp, err := json.Marshal(strct.Meta); err != nil {
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
    // Marshal the "originalPrescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"originalPrescription\": ")
	if tmp, err := json.Marshal(strct.OriginalPrescription); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Patient" field is required
    if strct.Patient == nil {
        return nil, errors.New("patient is a required field")
    }
    // Marshal the "patient" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"patient\": ")
	if tmp, err := json.Marshal(strct.Patient); err != nil {
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
    // Marshal the "prescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"prescription\": ")
	if tmp, err := json.Marshal(strct.Prescription); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Priority" field is required
    if strct.Priority == nil {
        return nil, errors.New("priority is a required field")
    }
    // Marshal the "priority" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"priority\": ")
	if tmp, err := json.Marshal(strct.Priority); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "procedure" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"procedure\": ")
	if tmp, err := json.Marshal(strct.Procedure); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Provider" field is required
    if strct.Provider == nil {
        return nil, errors.New("provider is a required field")
    }
    // Marshal the "provider" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"provider\": ")
	if tmp, err := json.Marshal(strct.Provider); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "referral" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"referral\": ")
	if tmp, err := json.Marshal(strct.Referral); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "related" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"related\": ")
	if tmp, err := json.Marshal(strct.Related); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ResourceType" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "resourceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"resourceType\": ")
	if tmp, err := json.Marshal(strct.ResourceType); err != nil {
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
    // Marshal the "subType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subType\": ")
	if tmp, err := json.Marshal(strct.SubType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "supportingInfo" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"supportingInfo\": ")
	if tmp, err := json.Marshal(strct.SupportingInfo); err != nil {
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
    // Marshal the "total" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"total\": ")
	if tmp, err := json.Marshal(strct.Total); err != nil {
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
    // Marshal the "use" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"use\": ")
	if tmp, err := json.Marshal(strct.Use); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Claim) UnmarshalJSON(b []byte) error {
    insuranceReceived := false
    patientReceived := false
    priorityReceived := false
    providerReceived := false
    resourceTypeReceived := false
    typeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "accident":
            if err := json.Unmarshal([]byte(v), &strct.Accident); err != nil {
                return err
             }
        case "billablePeriod":
            if err := json.Unmarshal([]byte(v), &strct.BillablePeriod); err != nil {
                return err
             }
        case "careTeam":
            if err := json.Unmarshal([]byte(v), &strct.CareTeam); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "created":
            if err := json.Unmarshal([]byte(v), &strct.Created); err != nil {
                return err
             }
        case "diagnosis":
            if err := json.Unmarshal([]byte(v), &strct.Diagnosis); err != nil {
                return err
             }
        case "enterer":
            if err := json.Unmarshal([]byte(v), &strct.Enterer); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "facility":
            if err := json.Unmarshal([]byte(v), &strct.Facility); err != nil {
                return err
             }
        case "fundsReserve":
            if err := json.Unmarshal([]byte(v), &strct.FundsReserve); err != nil {
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
        case "implicitRules":
            if err := json.Unmarshal([]byte(v), &strct.ImplicitRules); err != nil {
                return err
             }
        case "insurance":
            if err := json.Unmarshal([]byte(v), &strct.Insurance); err != nil {
                return err
             }
            insuranceReceived = true
        case "insurer":
            if err := json.Unmarshal([]byte(v), &strct.Insurer); err != nil {
                return err
             }
        case "item":
            if err := json.Unmarshal([]byte(v), &strct.Item); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "meta":
            if err := json.Unmarshal([]byte(v), &strct.Meta); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "originalPrescription":
            if err := json.Unmarshal([]byte(v), &strct.OriginalPrescription); err != nil {
                return err
             }
        case "patient":
            if err := json.Unmarshal([]byte(v), &strct.Patient); err != nil {
                return err
             }
            patientReceived = true
        case "payee":
            if err := json.Unmarshal([]byte(v), &strct.Payee); err != nil {
                return err
             }
        case "prescription":
            if err := json.Unmarshal([]byte(v), &strct.Prescription); err != nil {
                return err
             }
        case "priority":
            if err := json.Unmarshal([]byte(v), &strct.Priority); err != nil {
                return err
             }
            priorityReceived = true
        case "procedure":
            if err := json.Unmarshal([]byte(v), &strct.Procedure); err != nil {
                return err
             }
        case "provider":
            if err := json.Unmarshal([]byte(v), &strct.Provider); err != nil {
                return err
             }
            providerReceived = true
        case "referral":
            if err := json.Unmarshal([]byte(v), &strct.Referral); err != nil {
                return err
             }
        case "related":
            if err := json.Unmarshal([]byte(v), &strct.Related); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "subType":
            if err := json.Unmarshal([]byte(v), &strct.SubType); err != nil {
                return err
             }
        case "supportingInfo":
            if err := json.Unmarshal([]byte(v), &strct.SupportingInfo); err != nil {
                return err
             }
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        case "total":
            if err := json.Unmarshal([]byte(v), &strct.Total); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
            typeReceived = true
        case "use":
            if err := json.Unmarshal([]byte(v), &strct.Use); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if insurance (a required property) was received
    if !insuranceReceived {
        return errors.New("\"insurance\" is required but was not present")
    }
    // check if patient (a required property) was received
    if !patientReceived {
        return errors.New("\"patient\" is required but was not present")
    }
    // check if priority (a required property) was received
    if !priorityReceived {
        return errors.New("\"priority\" is required but was not present")
    }
    // check if provider (a required property) was received
    if !providerReceived {
        return errors.New("\"provider\" is required but was not present")
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    // check if type (a required property) was received
    if !typeReceived {
        return errors.New("\"type\" is required but was not present")
    }
    return nil
}
