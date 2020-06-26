// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// HealthcareService The details of a healthcare service available at a location.
type HealthcareService struct {

  // This flag is used to mark the record to not be used. This is not used when a center is closed for maintenance, or for holidays, the notAvailable period is to be used for this.
  Active bool `json:"active,omitempty"`

  // Indicates whether or not a prospective consumer will require an appointment for a particular service at a site to be provided by the Organization. Indicates if an appointment is required for access to this service.
  AppointmentRequired bool `json:"appointmentRequired,omitempty"`

  // A description of site availability exceptions, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as details in the available Times and not available Times.
  AvailabilityExceptions string `json:"availabilityExceptions,omitempty"`

  // A collection of times that the Service Site is available.
  AvailableTime []*HealthcareServiceAvailableTime `json:"availableTime,omitempty"`

  // Identifies the broad category of service being performed or delivered.
  Category []*CodeableConcept `json:"category,omitempty"`

  // Collection of characteristics (attributes).
  Characteristic []*CodeableConcept `json:"characteristic,omitempty"`

  // Any additional description of the service and/or any specific issues not covered by the other attributes, which can be displayed as further detail under the serviceName.
  Comment string `json:"comment,omitempty"`

  // Some services are specifically made available in multiple languages, this property permits a directory to declare the languages this is offered in. Typically this is only provided where a service operates in communities with mixed languages used.
  Communication []*CodeableConcept `json:"communication,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The location(s) that this service is available to (not where the service is provided).
  CoverageArea []*Reference `json:"coverageArea,omitempty"`

  // Does this service have specific eligibility requirements that need to be met in order to use the service?
  Eligibility []*HealthcareServiceEligibility `json:"eligibility,omitempty"`

  // Technical endpoints providing access to services operated for the specific healthcare services defined at this resource.
  Endpoint []*Reference `json:"endpoint,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Extra details about the service that can't be placed in the other fields.
  ExtraDetails string `json:"extraDetails,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // External identifiers for this item.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The location(s) where this healthcare service may be provided.
  Location []*Reference `json:"location,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Further description of the service as it would be presented to a consumer while searching.
  Name string `json:"name,omitempty"`

  // The HealthcareService is not available during this period of time due to the provided reason.
  NotAvailable []*HealthcareServiceNotAvailable `json:"notAvailable,omitempty"`

  // If there is a photo/symbol associated with this HealthcareService, it may be included here to facilitate quick identification of the service in a list.
  Photo *Attachment `json:"photo,omitempty"`

  // Programs that this service is applicable to.
  Program []*CodeableConcept `json:"program,omitempty"`

  // The organization that provides this healthcare service.
  ProvidedBy *Reference `json:"providedBy,omitempty"`

  // Ways that the service accepts referrals, if this is not provided then it is implied that no referral is required.
  ReferralMethod []*CodeableConcept `json:"referralMethod,omitempty"`

  // This is a HealthcareService resource
  ResourceType interface{} `json:"resourceType"`

  // The code(s) that detail the conditions under which the healthcare service is available/offered.
  ServiceProvisionCode []*CodeableConcept `json:"serviceProvisionCode,omitempty"`

  // Collection of specialties handled by the service site. This is more of a medical term.
  Specialty []*CodeableConcept `json:"specialty,omitempty"`

  // List of contacts related to this specific healthcare service.
  Telecom []*ContactPoint `json:"telecom,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // The specific type of service that may be delivered or performed.
  Type []*CodeableConcept `json:"type,omitempty"`
}

func (strct *HealthcareService) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "active" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"active\": ")
	if tmp, err := json.Marshal(strct.Active); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "appointmentRequired" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"appointmentRequired\": ")
	if tmp, err := json.Marshal(strct.AppointmentRequired); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "availabilityExceptions" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"availabilityExceptions\": ")
	if tmp, err := json.Marshal(strct.AvailabilityExceptions); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "availableTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"availableTime\": ")
	if tmp, err := json.Marshal(strct.AvailableTime); err != nil {
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
    // Marshal the "characteristic" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"characteristic\": ")
	if tmp, err := json.Marshal(strct.Characteristic); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "comment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"comment\": ")
	if tmp, err := json.Marshal(strct.Comment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "communication" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"communication\": ")
	if tmp, err := json.Marshal(strct.Communication); err != nil {
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
    // Marshal the "coverageArea" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"coverageArea\": ")
	if tmp, err := json.Marshal(strct.CoverageArea); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "eligibility" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"eligibility\": ")
	if tmp, err := json.Marshal(strct.Eligibility); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "endpoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"endpoint\": ")
	if tmp, err := json.Marshal(strct.Endpoint); err != nil {
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
    // Marshal the "extraDetails" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"extraDetails\": ")
	if tmp, err := json.Marshal(strct.ExtraDetails); err != nil {
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
    // Marshal the "name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "notAvailable" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"notAvailable\": ")
	if tmp, err := json.Marshal(strct.NotAvailable); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "photo" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"photo\": ")
	if tmp, err := json.Marshal(strct.Photo); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "program" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"program\": ")
	if tmp, err := json.Marshal(strct.Program); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "providedBy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"providedBy\": ")
	if tmp, err := json.Marshal(strct.ProvidedBy); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "referralMethod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"referralMethod\": ")
	if tmp, err := json.Marshal(strct.ReferralMethod); err != nil {
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
    // Marshal the "serviceProvisionCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"serviceProvisionCode\": ")
	if tmp, err := json.Marshal(strct.ServiceProvisionCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "specialty" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"specialty\": ")
	if tmp, err := json.Marshal(strct.Specialty); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "telecom" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"telecom\": ")
	if tmp, err := json.Marshal(strct.Telecom); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *HealthcareService) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "active":
            if err := json.Unmarshal([]byte(v), &strct.Active); err != nil {
                return err
             }
        case "appointmentRequired":
            if err := json.Unmarshal([]byte(v), &strct.AppointmentRequired); err != nil {
                return err
             }
        case "availabilityExceptions":
            if err := json.Unmarshal([]byte(v), &strct.AvailabilityExceptions); err != nil {
                return err
             }
        case "availableTime":
            if err := json.Unmarshal([]byte(v), &strct.AvailableTime); err != nil {
                return err
             }
        case "category":
            if err := json.Unmarshal([]byte(v), &strct.Category); err != nil {
                return err
             }
        case "characteristic":
            if err := json.Unmarshal([]byte(v), &strct.Characteristic); err != nil {
                return err
             }
        case "comment":
            if err := json.Unmarshal([]byte(v), &strct.Comment); err != nil {
                return err
             }
        case "communication":
            if err := json.Unmarshal([]byte(v), &strct.Communication); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "coverageArea":
            if err := json.Unmarshal([]byte(v), &strct.CoverageArea); err != nil {
                return err
             }
        case "eligibility":
            if err := json.Unmarshal([]byte(v), &strct.Eligibility); err != nil {
                return err
             }
        case "endpoint":
            if err := json.Unmarshal([]byte(v), &strct.Endpoint); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "extraDetails":
            if err := json.Unmarshal([]byte(v), &strct.ExtraDetails); err != nil {
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
        case "meta":
            if err := json.Unmarshal([]byte(v), &strct.Meta); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
        case "notAvailable":
            if err := json.Unmarshal([]byte(v), &strct.NotAvailable); err != nil {
                return err
             }
        case "photo":
            if err := json.Unmarshal([]byte(v), &strct.Photo); err != nil {
                return err
             }
        case "program":
            if err := json.Unmarshal([]byte(v), &strct.Program); err != nil {
                return err
             }
        case "providedBy":
            if err := json.Unmarshal([]byte(v), &strct.ProvidedBy); err != nil {
                return err
             }
        case "referralMethod":
            if err := json.Unmarshal([]byte(v), &strct.ReferralMethod); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "serviceProvisionCode":
            if err := json.Unmarshal([]byte(v), &strct.ServiceProvisionCode); err != nil {
                return err
             }
        case "specialty":
            if err := json.Unmarshal([]byte(v), &strct.Specialty); err != nil {
                return err
             }
        case "telecom":
            if err := json.Unmarshal([]byte(v), &strct.Telecom); err != nil {
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