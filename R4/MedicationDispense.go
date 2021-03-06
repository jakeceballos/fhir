// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// MedicationDispense Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispense struct {

  // Indicates the medication order that is being dispensed against.
  AuthorizingPrescription []*Reference `json:"authorizingPrescription,omitempty"`

  // Indicates the type of medication dispense (for example, where the medication is expected to be consumed or administered (i.e. inpatient or outpatient)).
  Category *CodeableConcept `json:"category,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The encounter or episode of care that establishes the context for this event.
  Context *Reference `json:"context,omitempty"`

  // The amount of medication expressed as a timing amount.
  DaysSupply *Quantity `json:"daysSupply,omitempty"`

  // Identification of the facility/location where the medication was shipped to, as part of the dispense event.
  Destination *Reference `json:"destination,omitempty"`

  // Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. drug-drug interaction, duplicate therapy, dosage alert etc.
  DetectedIssue []*Reference `json:"detectedIssue,omitempty"`

  // Indicates how the medication is to be used by the patient.
  DosageInstruction []*Dosage `json:"dosageInstruction,omitempty"`

  // A summary of the events of interest that have occurred, such as when the dispense was verified.
  EventHistory []*Reference `json:"eventHistory,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Identifiers associated with this Medication Dispense that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The principal physical location where the dispense was performed.
  Location *Reference `json:"location,omitempty"`

  // Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
  MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept,omitempty"`

  // Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
  MedicationReference *Reference `json:"medicationReference,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Extra information about the dispense that could not be conveyed in the other attributes.
  Note []*Annotation `json:"note,omitempty"`

  // The procedure that trigger the dispense.
  PartOf []*Reference `json:"partOf,omitempty"`

  // Indicates who or what performed the event.
  Performer []*MedicationDispensePerformer `json:"performer,omitempty"`

  // The amount of medication that has been dispensed. Includes unit of measure.
  Quantity *Quantity `json:"quantity,omitempty"`

  // Identifies the person who picked up the medication.  This will usually be a patient or their caregiver, but some cases exist where it can be a healthcare professional.
  Receiver []*Reference `json:"receiver,omitempty"`

  // This is a MedicationDispense resource
  ResourceType interface{} `json:"resourceType"`

  // A code specifying the state of the set of dispense events.
  Status string `json:"status,omitempty"`

  // Indicates the reason why a dispense was not performed.
  StatusReasonCodeableConcept *CodeableConcept `json:"statusReasonCodeableConcept,omitempty"`

  // Indicates the reason why a dispense was not performed.
  StatusReasonReference *Reference `json:"statusReasonReference,omitempty"`

  // A link to a resource representing the person or the group to whom the medication will be given.
  Subject *Reference `json:"subject,omitempty"`

  // Indicates whether or not substitution was made as part of the dispense.  In some cases, substitution will be expected but does not happen, in other cases substitution is not expected but does happen.  This block explains what substitution did or did not happen and why.  If nothing is specified, substitution was not done.
  Substitution *MedicationDispenseSubstitution `json:"substitution,omitempty"`

  // Additional information that supports the medication being dispensed.
  SupportingInformation []*Reference `json:"supportingInformation,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // Indicates the type of dispensing event that is performed. For example, Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
  Type *CodeableConcept `json:"type,omitempty"`

  // The time the dispensed product was provided to the patient or their representative.
  WhenHandedOver string `json:"whenHandedOver,omitempty"`

  // The time when the dispensed product was packaged and reviewed.
  WhenPrepared string `json:"whenPrepared,omitempty"`
}

func (strct *MedicationDispense) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "authorizingPrescription" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"authorizingPrescription\": ")
	if tmp, err := json.Marshal(strct.AuthorizingPrescription); err != nil {
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
    // Marshal the "context" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"context\": ")
	if tmp, err := json.Marshal(strct.Context); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "daysSupply" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"daysSupply\": ")
	if tmp, err := json.Marshal(strct.DaysSupply); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "destination" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"destination\": ")
	if tmp, err := json.Marshal(strct.Destination); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "detectedIssue" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"detectedIssue\": ")
	if tmp, err := json.Marshal(strct.DetectedIssue); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "dosageInstruction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"dosageInstruction\": ")
	if tmp, err := json.Marshal(strct.DosageInstruction); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "eventHistory" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"eventHistory\": ")
	if tmp, err := json.Marshal(strct.EventHistory); err != nil {
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
    // Marshal the "location" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "medicationCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"medicationCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.MedicationCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "medicationReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"medicationReference\": ")
	if tmp, err := json.Marshal(strct.MedicationReference); err != nil {
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
    // Marshal the "note" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"note\": ")
	if tmp, err := json.Marshal(strct.Note); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "partOf" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"partOf\": ")
	if tmp, err := json.Marshal(strct.PartOf); err != nil {
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
    // Marshal the "receiver" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"receiver\": ")
	if tmp, err := json.Marshal(strct.Receiver); err != nil {
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
    // Marshal the "statusReasonCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"statusReasonCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.StatusReasonCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "statusReasonReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"statusReasonReference\": ")
	if tmp, err := json.Marshal(strct.StatusReasonReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "subject" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"subject\": ")
	if tmp, err := json.Marshal(strct.Subject); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "substitution" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"substitution\": ")
	if tmp, err := json.Marshal(strct.Substitution); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "supportingInformation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"supportingInformation\": ")
	if tmp, err := json.Marshal(strct.SupportingInformation); err != nil {
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
    // Marshal the "whenHandedOver" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"whenHandedOver\": ")
	if tmp, err := json.Marshal(strct.WhenHandedOver); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "whenPrepared" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"whenPrepared\": ")
	if tmp, err := json.Marshal(strct.WhenPrepared); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MedicationDispense) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "authorizingPrescription":
            if err := json.Unmarshal([]byte(v), &strct.AuthorizingPrescription); err != nil {
                return err
             }
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "context":
            if err := json.Unmarshal([]byte(v), &strct.Context); err != nil {
                return err
             }
        case "daysSupply":
            if err := json.Unmarshal([]byte(v), &strct.DaysSupply); err != nil {
                return err
             }
        case "destination":
            if err := json.Unmarshal([]byte(v), &strct.Destination); err != nil {
                return err
             }
        case "detectedIssue":
            if err := json.Unmarshal([]byte(v), &strct.DetectedIssue); err != nil {
                return err
             }
        case "dosageInstruction":
            if err := json.Unmarshal([]byte(v), &strct.DosageInstruction); err != nil {
                return err
             }
        case "eventHistory":
            if err := json.Unmarshal([]byte(v), &strct.EventHistory); err != nil {
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
        case "implicitRules":
            if err := json.Unmarshal([]byte(v), &strct.ImplicitRules); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "location":
            if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
                return err
             }
        case "medicationCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.MedicationCodeableConcept); err != nil {
                return err
             }
        case "medicationReference":
            if err := json.Unmarshal([]byte(v), &strct.MedicationReference); err != nil {
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
        case "note":
            if err := json.Unmarshal([]byte(v), &strct.Note); err != nil {
                return err
             }
        case "partOf":
            if err := json.Unmarshal([]byte(v), &strct.PartOf); err != nil {
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
        case "receiver":
            if err := json.Unmarshal([]byte(v), &strct.Receiver); err != nil {
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
        case "statusReasonCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.StatusReasonCodeableConcept); err != nil {
                return err
             }
        case "statusReasonReference":
            if err := json.Unmarshal([]byte(v), &strct.StatusReasonReference); err != nil {
                return err
             }
        case "subject":
            if err := json.Unmarshal([]byte(v), &strct.Subject); err != nil {
                return err
             }
        case "substitution":
            if err := json.Unmarshal([]byte(v), &strct.Substitution); err != nil {
                return err
             }
        case "supportingInformation":
            if err := json.Unmarshal([]byte(v), &strct.SupportingInformation); err != nil {
                return err
             }
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        case "type":
            if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
                return err
             }
        case "whenHandedOver":
            if err := json.Unmarshal([]byte(v), &strct.WhenHandedOver); err != nil {
                return err
             }
        case "whenPrepared":
            if err := json.Unmarshal([]byte(v), &strct.WhenPrepared); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    return nil
}
