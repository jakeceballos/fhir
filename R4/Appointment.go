// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// Appointment A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment struct {

  // The style of appointment or patient that has been booked in the slot (not service type).
  AppointmentType *CodeableConcept `json:"appointmentType,omitempty"`

  // The service request this appointment is allocated to assess (e.g. incoming referral or procedure request).
  BasedOn []*Reference `json:"basedOn,omitempty"`

  // The coded reason for the appointment being cancelled. This is often used in reporting/billing/futher processing to determine if further actions are required, or specific fees apply.
  CancelationReason *CodeableConcept `json:"cancelationReason,omitempty"`

  // Additional comments about the appointment.
  Comment string `json:"comment,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The date that this appointment was initially created. This could be different to the meta.lastModified value on the initial entry, as this could have been before the resource was created on the FHIR server, and should remain unchanged over the lifespan of the appointment.
  Created string `json:"created,omitempty"`

  // The brief description of the appointment as would be shown on a subject line in a meeting request, or appointment list. Detailed or expanded information should be put in the comment field.
  Description string `json:"description,omitempty"`

  // Date/Time that the appointment is to conclude.
  End string `json:"end,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // This records identifiers associated with this appointment concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // Number of minutes that the appointment is to take. This can be less than the duration between the start and end times.  For example, where the actual time of appointment is only an estimate or if a 30 minute appointment is being requested, but any time would work.  Also, if there is, for example, a planned 15 minute break in the middle of a long appointment, the duration may be 15 minutes less than the difference between the start and end.
  MinutesDuration float64 `json:"minutesDuration,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // List of participants involved in the appointment.
  Participant []*AppointmentParticipant `json:"participant"`

  // While Appointment.comment contains information for internal use, Appointment.patientInstructions is used to capture patient facing information about the Appointment (e.g. please bring your referral or fast from 8pm night before).
  PatientInstruction string `json:"patientInstruction,omitempty"`

  // The priority of the appointment. Can be used to make informed decisions if needing to re-prioritize appointments. (The iCal Standard specifies 0 as undefined, 1 as highest, 9 as lowest priority).
  Priority float64 `json:"priority,omitempty"`

  // The coded reason that this appointment is being scheduled. This is more clinical than administrative.
  ReasonCode []*CodeableConcept `json:"reasonCode,omitempty"`

  // Reason the appointment has been scheduled to take place, as specified using information from another resource. When the patient arrives and the encounter begins it may be used as the admission diagnosis. The indication will typically be a Condition (with other resources referenced in the evidence.detail), or a Procedure.
  ReasonReference []*Reference `json:"reasonReference,omitempty"`

  // A set of date ranges (potentially including times) that the appointment is preferred to be scheduled within.
  // 
  // The duration (usually in minutes) could also be provided to indicate the length of the appointment to fill and populate the start/end times for the actual allocated time. However, in other situations the duration may be calculated by the scheduling system.
  RequestedPeriod []*Period `json:"requestedPeriod,omitempty"`

  // This is a Appointment resource
  ResourceType interface{} `json:"resourceType"`

  // A broad categorization of the service that is to be performed during this appointment.
  ServiceCategory []*CodeableConcept `json:"serviceCategory,omitempty"`

  // The specific service that is to be performed during this appointment.
  ServiceType []*CodeableConcept `json:"serviceType,omitempty"`

  // The slots from the participants' schedules that will be filled by the appointment.
  Slot []*Reference `json:"slot,omitempty"`

  // The specialty of a practitioner that would be required to perform the service requested in this appointment.
  Specialty []*CodeableConcept `json:"specialty,omitempty"`

  // Date/Time that the appointment is to take place.
  Start string `json:"start,omitempty"`

  // The overall status of the Appointment. Each of the participants has their own participation status which indicates their involvement in the process, however this status indicates the shared status.
  Status interface{} `json:"status,omitempty"`

  // Additional information to support the appointment provided when making the appointment.
  SupportingInformation []*Reference `json:"supportingInformation,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *Appointment) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "appointmentType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"appointmentType\": ")
	if tmp, err := json.Marshal(strct.AppointmentType); err != nil {
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
    // Marshal the "cancelationReason" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"cancelationReason\": ")
	if tmp, err := json.Marshal(strct.CancelationReason); err != nil {
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
    // Marshal the "description" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "end" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"end\": ")
	if tmp, err := json.Marshal(strct.End); err != nil {
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
    // Marshal the "minutesDuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"minutesDuration\": ")
	if tmp, err := json.Marshal(strct.MinutesDuration); err != nil {
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
    // "Participant" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "participant" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"participant\": ")
	if tmp, err := json.Marshal(strct.Participant); err != nil {
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
    // Marshal the "requestedPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"requestedPeriod\": ")
	if tmp, err := json.Marshal(strct.RequestedPeriod); err != nil {
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
    // Marshal the "serviceCategory" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"serviceCategory\": ")
	if tmp, err := json.Marshal(strct.ServiceCategory); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "serviceType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"serviceType\": ")
	if tmp, err := json.Marshal(strct.ServiceType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "slot" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"slot\": ")
	if tmp, err := json.Marshal(strct.Slot); err != nil {
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
    // Marshal the "start" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"start\": ")
	if tmp, err := json.Marshal(strct.Start); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Appointment) UnmarshalJSON(b []byte) error {
    participantReceived := false
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "appointmentType":
            if err := json.Unmarshal([]byte(v), &strct.AppointmentType); err != nil {
                return err
             }
        case "basedOn":
            if err := json.Unmarshal([]byte(v), &strct.BasedOn); err != nil {
                return err
             }
        case "cancelationReason":
            if err := json.Unmarshal([]byte(v), &strct.CancelationReason); err != nil {
                return err
             }
        case "comment":
            if err := json.Unmarshal([]byte(v), &strct.Comment); err != nil {
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
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "end":
            if err := json.Unmarshal([]byte(v), &strct.End); err != nil {
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
        case "meta":
            if err := json.Unmarshal([]byte(v), &strct.Meta); err != nil {
                return err
             }
        case "minutesDuration":
            if err := json.Unmarshal([]byte(v), &strct.MinutesDuration); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "participant":
            if err := json.Unmarshal([]byte(v), &strct.Participant); err != nil {
                return err
             }
            participantReceived = true
        case "patientInstruction":
            if err := json.Unmarshal([]byte(v), &strct.PatientInstruction); err != nil {
                return err
             }
        case "priority":
            if err := json.Unmarshal([]byte(v), &strct.Priority); err != nil {
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
        case "requestedPeriod":
            if err := json.Unmarshal([]byte(v), &strct.RequestedPeriod); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "serviceCategory":
            if err := json.Unmarshal([]byte(v), &strct.ServiceCategory); err != nil {
                return err
             }
        case "serviceType":
            if err := json.Unmarshal([]byte(v), &strct.ServiceType); err != nil {
                return err
             }
        case "slot":
            if err := json.Unmarshal([]byte(v), &strct.Slot); err != nil {
                return err
             }
        case "specialty":
            if err := json.Unmarshal([]byte(v), &strct.Specialty); err != nil {
                return err
             }
        case "start":
            if err := json.Unmarshal([]byte(v), &strct.Start); err != nil {
                return err
             }
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
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
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if participant (a required property) was received
    if !participantReceived {
        return errors.New("\"participant\" is required but was not present")
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    return nil
}
