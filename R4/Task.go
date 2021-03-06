// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "encoding/json"
    "fmt"
    "errors"
    "bytes"
)

// Task A task to be performed.
type Task struct {

  // The date and time this task was created.
  AuthoredOn string `json:"authoredOn,omitempty"`

  // BasedOn refers to a higher-level authorization that triggered the creation of the task.  It references a "request" resource such as a ServiceRequest, MedicationRequest, ServiceRequest, CarePlan, etc. which is distinct from the "request" resource the task is seeking to fulfill.  This latter resource is referenced by FocusOn.  For example, based on a ServiceRequest (= BasedOn), a task is created to fulfill a procedureRequest ( = FocusOn ) to collect a specimen from a patient.
  BasedOn []*Reference `json:"basedOn,omitempty"`

  // Contains business-specific nuances of the business state.
  BusinessStatus *CodeableConcept `json:"businessStatus,omitempty"`

  // A name or code (or both) briefly describing what the task involves.
  Code *CodeableConcept `json:"code,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // A free-text description of what is to be performed.
  Description string `json:"description,omitempty"`

  // The healthcare event  (e.g. a patient and healthcare provider interaction) during which this task was created.
  Encounter *Reference `json:"encounter,omitempty"`

  // Identifies the time action was first taken against the task (start) and/or the time final action was taken against the task prior to marking it as completed (end).
  ExecutionPeriod *Period `json:"executionPeriod,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The request being actioned or the resource being manipulated by this task.
  Focus *Reference `json:"focus,omitempty"`

  // The entity who benefits from the performance of the service specified in the task (e.g., the patient).
  For *Reference `json:"for,omitempty"`

  // An identifier that links together multiple tasks and other requests that were created in the same context.
  GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // The business identifier for this task.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // Additional information that may be needed in the execution of the task.
  Input []*TaskInput `json:"input,omitempty"`

  // The URL pointing to a *FHIR*-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
  InstantiatesCanonical string `json:"instantiatesCanonical,omitempty"`

  // The URL pointing to an *externally* maintained  protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
  InstantiatesUri string `json:"instantiatesUri,omitempty"`

  // Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Task.
  Insurance []*Reference `json:"insurance,omitempty"`

  // Indicates the "level" of actionability associated with the Task, i.e. i+R[9]Cs this a proposed task, a planned task, an actionable task, etc.
  Intent interface{} `json:"intent,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The date and time of last modification to this task.
  LastModified string `json:"lastModified,omitempty"`

  // Principal physical location where the this task is performed.
  Location *Reference `json:"location,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // Free-text information captured about the task as it progresses.
  Note []*Annotation `json:"note,omitempty"`

  // Outputs produced by the Task.
  Output []*TaskOutput `json:"output,omitempty"`

  // Individual organization or Device currently responsible for task execution.
  Owner *Reference `json:"owner,omitempty"`

  // Task that this particular task is part of.
  PartOf []*Reference `json:"partOf,omitempty"`

  // The kind of participant that should perform the task.
  PerformerType []*CodeableConcept `json:"performerType,omitempty"`

  // Indicates how quickly the Task should be addressed with respect to other requests.
  Priority string `json:"priority,omitempty"`

  // A description or code indicating why this task needs to be performed.
  ReasonCode *CodeableConcept `json:"reasonCode,omitempty"`

  // A resource reference indicating why this task needs to be performed.
  ReasonReference *Reference `json:"reasonReference,omitempty"`

  // Links to Provenance records for past versions of this Task that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the task.
  RelevantHistory []*Reference `json:"relevantHistory,omitempty"`

  // The creator of the task.
  Requester *Reference `json:"requester,omitempty"`

  // This is a Task resource
  ResourceType interface{} `json:"resourceType"`

  // If the Task.focus is a request resource and the task is seeking fulfillment (i.e. is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
  Restriction *TaskRestriction `json:"restriction,omitempty"`

  // The current status of the task.
  Status interface{} `json:"status,omitempty"`

  // An explanation as to why this task is held, failed, was refused, etc.
  StatusReason *CodeableConcept `json:"statusReason,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *Task) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "businessStatus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"businessStatus\": ")
	if tmp, err := json.Marshal(strct.BusinessStatus); err != nil {
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
    // Marshal the "executionPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"executionPeriod\": ")
	if tmp, err := json.Marshal(strct.ExecutionPeriod); err != nil {
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
    // Marshal the "focus" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"focus\": ")
	if tmp, err := json.Marshal(strct.Focus); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "for" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"for\": ")
	if tmp, err := json.Marshal(strct.For); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "groupIdentifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"groupIdentifier\": ")
	if tmp, err := json.Marshal(strct.GroupIdentifier); err != nil {
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
    // Marshal the "input" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"input\": ")
	if tmp, err := json.Marshal(strct.Input); err != nil {
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
    // Marshal the "lastModified" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"lastModified\": ")
	if tmp, err := json.Marshal(strct.LastModified); err != nil {
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
    // Marshal the "output" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"output\": ")
	if tmp, err := json.Marshal(strct.Output); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "owner" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"owner\": ")
	if tmp, err := json.Marshal(strct.Owner); err != nil {
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
    // Marshal the "restriction" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"restriction\": ")
	if tmp, err := json.Marshal(strct.Restriction); err != nil {
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
    // Marshal the "statusReason" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"statusReason\": ")
	if tmp, err := json.Marshal(strct.StatusReason); err != nil {
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

func (strct *Task) UnmarshalJSON(b []byte) error {
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "authoredOn":
            if err := json.Unmarshal([]byte(v), &strct.AuthoredOn); err != nil {
                return err
             }
        case "basedOn":
            if err := json.Unmarshal([]byte(v), &strct.BasedOn); err != nil {
                return err
             }
        case "businessStatus":
            if err := json.Unmarshal([]byte(v), &strct.BusinessStatus); err != nil {
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
        case "description":
            if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
                return err
             }
        case "encounter":
            if err := json.Unmarshal([]byte(v), &strct.Encounter); err != nil {
                return err
             }
        case "executionPeriod":
            if err := json.Unmarshal([]byte(v), &strct.ExecutionPeriod); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "focus":
            if err := json.Unmarshal([]byte(v), &strct.Focus); err != nil {
                return err
             }
        case "for":
            if err := json.Unmarshal([]byte(v), &strct.For); err != nil {
                return err
             }
        case "groupIdentifier":
            if err := json.Unmarshal([]byte(v), &strct.GroupIdentifier); err != nil {
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
        case "input":
            if err := json.Unmarshal([]byte(v), &strct.Input); err != nil {
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
        case "lastModified":
            if err := json.Unmarshal([]byte(v), &strct.LastModified); err != nil {
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
        case "note":
            if err := json.Unmarshal([]byte(v), &strct.Note); err != nil {
                return err
             }
        case "output":
            if err := json.Unmarshal([]byte(v), &strct.Output); err != nil {
                return err
             }
        case "owner":
            if err := json.Unmarshal([]byte(v), &strct.Owner); err != nil {
                return err
             }
        case "partOf":
            if err := json.Unmarshal([]byte(v), &strct.PartOf); err != nil {
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
        case "requester":
            if err := json.Unmarshal([]byte(v), &strct.Requester); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "restriction":
            if err := json.Unmarshal([]byte(v), &strct.Restriction); err != nil {
                return err
             }
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
        case "statusReason":
            if err := json.Unmarshal([]byte(v), &strct.StatusReason); err != nil {
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
    return nil
}
