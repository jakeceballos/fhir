// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "fmt"
    "bytes"
    "errors"
    "encoding/json"
)

// ServiceRequest A record of a request for service such as diagnostic investigations, treatments, or operations to be performed.
type ServiceRequest struct {

  // If a CodeableConcept is present, it indicates the pre-condition for performing the service.  For example "pain", "on flare-up", etc.
  AsNeededBoolean bool `json:"asNeededBoolean,omitempty"`

  // If a CodeableConcept is present, it indicates the pre-condition for performing the service.  For example "pain", "on flare-up", etc.
  AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

  // When the request transitioned to being actionable.
  AuthoredOn string `json:"authoredOn,omitempty"`

  // Plan/proposal/order fulfilled by this request.
  BasedOn []*Reference `json:"basedOn,omitempty"`

  // Anatomic location where the procedure should be performed. This is the target site.
  BodySite []*CodeableConcept `json:"bodySite,omitempty"`

  // A code that classifies the service for searching, sorting and display purposes (e.g. "Surgical Procedure").
  Category []*CodeableConcept `json:"category,omitempty"`

  // A code that identifies a particular service (i.e., procedure, diagnostic investigation, or panel of investigations) that have been requested.
  Code *CodeableConcept `json:"code,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // Set this to true if the record is saying that the service/procedure should NOT be performed.
  DoNotPerform bool `json:"doNotPerform,omitempty"`

  // An encounter that provides additional information about the healthcare context in which this request is made.
  Encounter *Reference `json:"encounter,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Identifiers assigned to this order instance by the orderer and/or the receiver and/or order fulfiller.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
  InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

  // The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
  InstantiatesUri []string `json:"instantiatesUri,omitempty"`

  // Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be needed for delivering the requested service.
  Insurance []*Reference `json:"insurance,omitempty"`

  // Whether the request is a proposal, plan, an original order or a reflex order.
  Intent string `json:"intent,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The preferred location(s) where the procedure should actually happen in coded or free text form. E.g. at home or nursing day care center.
  LocationCode []*CodeableConcept `json:"locationCode,omitempty"`

  // A reference to the the preferred location(s) where the procedure should actually happen. E.g. at home or nursing day care center.
  LocationReference []*Reference `json:"locationReference,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Any other notes and comments made about the service request. For example, internal billing notes.
  Note []*Annotation `json:"note,omitempty"`

  // The date/time at which the requested service should occur.
  OccurrenceDateTime string `json:"occurrenceDateTime,omitempty"`

  // The date/time at which the requested service should occur.
  OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`

  // The date/time at which the requested service should occur.
  OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`

  // Additional details and instructions about the how the services are to be delivered.   For example, and order for a urinary catheter may have an order detail for an external or indwelling catheter, or an order for a bandage may require additional instructions specifying how the bandage should be applied.
  OrderDetail []*CodeableConcept `json:"orderDetail,omitempty"`

  // Instructions in terms that are understood by the patient or consumer.
  PatientInstruction string `json:"patientInstruction,omitempty"`

  // The desired performer for doing the requested service.  For example, the surgeon, dermatopathologist, endoscopist, etc.
  Performer []*Reference `json:"performer,omitempty"`

  // Desired type of performer for doing the requested service.
  PerformerType *CodeableConcept `json:"performerType,omitempty"`

  // Indicates how quickly the ServiceRequest should be addressed with respect to other requests.
  Priority string `json:"priority,omitempty"`

  // An amount of service being requested which can be a quantity ( for example $1,500 home modification), a ratio ( for example, 20 half day visits per month), or a range (2.0 to 1.8 Gy per fraction).
  QuantityQuantity *Quantity `json:"quantityQuantity,omitempty"`

  // An amount of service being requested which can be a quantity ( for example $1,500 home modification), a ratio ( for example, 20 half day visits per month), or a range (2.0 to 1.8 Gy per fraction).
  QuantityRange *Range `json:"quantityRange,omitempty"`

  // An amount of service being requested which can be a quantity ( for example $1,500 home modification), a ratio ( for example, 20 half day visits per month), or a range (2.0 to 1.8 Gy per fraction).
  QuantityRatio *Ratio `json:"quantityRatio,omitempty"`

  // An explanation or justification for why this service is being requested in coded or textual form.   This is often for billing purposes.  May relate to the resources referred to in `supportingInfo`.
  ReasonCode []*CodeableConcept `json:"reasonCode,omitempty"`

  // Indicates another resource that provides a justification for why this service is being requested.   May relate to the resources referred to in `supportingInfo`.
  ReasonReference []*Reference `json:"reasonReference,omitempty"`

  // Key events in the history of the request.
  RelevantHistory []*Reference `json:"relevantHistory,omitempty"`

  // The request takes the place of the referenced completed or terminated request(s).
  Replaces []*Reference `json:"replaces,omitempty"`

  // The individual who initiated the request and has responsibility for its activation.
  Requester *Reference `json:"requester,omitempty"`

  // A shared identifier common to all service requests that were authorized more or less simultaneously by a single author, representing the composite or group identifier.
  Requisition *Identifier `json:"requisition,omitempty"`

  // This is a ServiceRequest resource
  ResourceType interface{} `json:"resourceType"`

  // One or more specimens that the laboratory procedure will use.
  Specimen []*Reference `json:"specimen,omitempty"`

  // The status of the order.
  Status string `json:"status,omitempty"`

  // On whom or what the service is to be performed. This is usually a human patient, but can also be requested on animals, groups of humans or animals, devices such as dialysis machines, or even locations (typically for environmental scans).
  Subject *Reference `json:"subject"`

  // Additional clinical information about the patient or specimen that may influence the services or their interpretations.     This information includes diagnosis, clinical findings and other observations.  In laboratory ordering these are typically referred to as "ask at order entry questions (AOEs)".  This includes observations explicitly requested by the producer (filler) to provide context or supporting information needed to complete the order. For example,  reporting the amount of inspired oxygen for blood gas measurements.
  SupportingInfo []*Reference `json:"supportingInfo,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *ServiceRequest) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "asNeededBoolean" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"asNeededBoolean\": ")
	if tmp, err := json.Marshal(strct.AsNeededBoolean); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "asNeededCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"asNeededCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.AsNeededCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "authoredOn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"authoredOn\": ")
	if tmp, err := json.Marshal(strct.AuthoredOn); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "basedOn" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"basedOn\": ")
	if tmp, err := json.Marshal(strct.BasedOn); err != nil {
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
    // Marshal the "code" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"code\": ")
	if tmp, err := json.Marshal(strct.Code); err != nil {
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
    // Marshal the "doNotPerform" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"doNotPerform\": ")
	if tmp, err := json.Marshal(strct.DoNotPerform); err != nil {
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
    // Marshal the "instantiatesCanonical" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"instantiatesCanonical\": ")
	if tmp, err := json.Marshal(strct.InstantiatesCanonical); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "instantiatesUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"instantiatesUri\": ")
	if tmp, err := json.Marshal(strct.InstantiatesUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "intent" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"intent\": ")
	if tmp, err := json.Marshal(strct.Intent); err != nil {
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
    // Marshal the "locationCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"locationCode\": ")
	if tmp, err := json.Marshal(strct.LocationCode); err != nil {
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
    // Marshal the "occurrenceDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"occurrenceDateTime\": ")
	if tmp, err := json.Marshal(strct.OccurrenceDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "occurrencePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"occurrencePeriod\": ")
	if tmp, err := json.Marshal(strct.OccurrencePeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "occurrenceTiming" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"occurrenceTiming\": ")
	if tmp, err := json.Marshal(strct.OccurrenceTiming); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "orderDetail" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"orderDetail\": ")
	if tmp, err := json.Marshal(strct.OrderDetail); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "patientInstruction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"patientInstruction\": ")
	if tmp, err := json.Marshal(strct.PatientInstruction); err != nil {
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
    // Marshal the "performerType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"performerType\": ")
	if tmp, err := json.Marshal(strct.PerformerType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "quantityQuantity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"quantityQuantity\": ")
	if tmp, err := json.Marshal(strct.QuantityQuantity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "quantityRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"quantityRange\": ")
	if tmp, err := json.Marshal(strct.QuantityRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "quantityRatio" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"quantityRatio\": ")
	if tmp, err := json.Marshal(strct.QuantityRatio); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "reasonCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reasonCode\": ")
	if tmp, err := json.Marshal(strct.ReasonCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "reasonReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"reasonReference\": ")
	if tmp, err := json.Marshal(strct.ReasonReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "relevantHistory" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"relevantHistory\": ")
	if tmp, err := json.Marshal(strct.RelevantHistory); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "replaces" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"replaces\": ")
	if tmp, err := json.Marshal(strct.Replaces); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "requester" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"requester\": ")
	if tmp, err := json.Marshal(strct.Requester); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "requisition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"requisition\": ")
	if tmp, err := json.Marshal(strct.Requisition); err != nil {
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
    // Marshal the "specimen" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"specimen\": ")
	if tmp, err := json.Marshal(strct.Specimen); err != nil {
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
    // "Subject" field is required
    if strct.Subject == nil {
        return nil, errors.New("subject is a required field")
    }
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ServiceRequest) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    subjectReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "asNeededBoolean":
            if err := json.Unmarshal([]byte(v), &strct.AsNeededBoolean); err != nil {
                return err
             }
        case "asNeededCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.AsNeededCodeableConcept); err != nil {
                return err
             }
        case "authoredOn":
            if err := json.Unmarshal([]byte(v), &strct.AuthoredOn); err != nil {
                return err
             }
        case "basedOn":
            if err := json.Unmarshal([]byte(v), &strct.BasedOn); err != nil {
                return err
             }
        case "bodySite":
            if err := json.Unmarshal([]byte(v), &strct.BodySite); err != nil {
                return err
             }
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
        case "code":
            if err := json.Unmarshal([]byte(v), &strct.Code); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "doNotPerform":
            if err := json.Unmarshal([]byte(v), &strct.DoNotPerform); err != nil {
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
        case "instantiatesCanonical":
            if err := json.Unmarshal([]byte(v), &strct.InstantiatesCanonical); err != nil {
                return err
             }
        case "instantiatesUri":
            if err := json.Unmarshal([]byte(v), &strct.InstantiatesUri); err != nil {
                return err
             }
        case "insurance":
            if err := json.Unmarshal([]byte(v), &strct.Insurance); err != nil {
                return err
             }
        case "intent":
            if err := json.Unmarshal([]byte(v), &strct.Intent); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "locationCode":
            if err := json.Unmarshal([]byte(v), &strct.LocationCode); err != nil {
                return err
             }
        case "locationReference":
            if err := json.Unmarshal([]byte(v), &strct.LocationReference); err != nil {
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
        case "occurrenceDateTime":
            if err := json.Unmarshal([]byte(v), &strct.OccurrenceDateTime); err != nil {
                return err
             }
        case "occurrencePeriod":
            if err := json.Unmarshal([]byte(v), &strct.OccurrencePeriod); err != nil {
                return err
             }
        case "occurrenceTiming":
            if err := json.Unmarshal([]byte(v), &strct.OccurrenceTiming); err != nil {
                return err
             }
        case "orderDetail":
            if err := json.Unmarshal([]byte(v), &strct.OrderDetail); err != nil {
                return err
             }
        case "patientInstruction":
            if err := json.Unmarshal([]byte(v), &strct.PatientInstruction); err != nil {
                return err
             }
        case "performer":
            if err := json.Unmarshal([]byte(v), &strct.Performer); err != nil {
                return err
             }
        case "performerType":
            if err := json.Unmarshal([]byte(v), &strct.PerformerType); err != nil {
                return err
             }
        case "priority":
            if err := json.Unmarshal([]byte(v), &strct.Priority); err != nil {
                return err
             }
        case "quantityQuantity":
            if err := json.Unmarshal([]byte(v), &strct.QuantityQuantity); err != nil {
                return err
             }
        case "quantityRange":
            if err := json.Unmarshal([]byte(v), &strct.QuantityRange); err != nil {
                return err
             }
        case "quantityRatio":
            if err := json.Unmarshal([]byte(v), &strct.QuantityRatio); err != nil {
                return err
             }
        case "reasonCode":
            if err := json.Unmarshal([]byte(v), &strct.ReasonCode); err != nil {
                return err
             }
        case "reasonReference":
            if err := json.Unmarshal([]byte(v), &strct.ReasonReference); err != nil {
                return err
             }
        case "relevantHistory":
            if err := json.Unmarshal([]byte(v), &strct.RelevantHistory); err != nil {
                return err
             }
        case "replaces":
            if err := json.Unmarshal([]byte(v), &strct.Replaces); err != nil {
                return err
             }
        case "requester":
            if err := json.Unmarshal([]byte(v), &strct.Requester); err != nil {
                return err
             }
        case "requisition":
            if err := json.Unmarshal([]byte(v), &strct.Requisition); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "specimen":
            if err := json.Unmarshal([]byte(v), &strct.Specimen); err != nil {
                return err
             }
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "subject":
            if err := json.Unmarshal([]byte(v), &strct.Subject); err != nil {
                return err
             }
            subjectReceived = true
        case "supportingInfo":
            if err := json.Unmarshal([]byte(v), &strct.SupportingInfo); err != nil {
                return err
             }
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
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
    // check if subject (a required property) was received
    if !subjectReceived {
        return errors.New("\"subject\" is required but was not present")
    }
    return nil
}
