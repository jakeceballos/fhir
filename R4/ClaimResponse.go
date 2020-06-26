// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "errors"
    "encoding/json"
    "fmt"
)

// ClaimResponse This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse struct {

  // The first-tier service adjudications for payor added product or service lines.
  AddItem []*ClaimResponseAddItem `json:"addItem,omitempty"`

  // The adjudication results which are presented at the header level rather than at the line-item or add-item levels.
  Adjudication []*ClaimResponseAdjudication `json:"adjudication,omitempty"`

  // Request for additional supporting or authorizing information.
  CommunicationRequest []*Reference `json:"communicationRequest,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The date this resource was created.
  Created string `json:"created,omitempty"`

  // A human readable description of the status of the adjudication.
  Disposition string `json:"disposition,omitempty"`

  // Errors encountered during the processing of the adjudication.
  Error []*ClaimResponseError `json:"error,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The actual form, by reference or inclusion, for printing the content or an EOB.
  Form *Attachment `json:"form,omitempty"`

  // A code for the form to be used for printing the content.
  FormCode *CodeableConcept `json:"formCode,omitempty"`

  // A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom.
  FundsReserve *CodeableConcept `json:"fundsReserve,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // A unique identifier assigned to this claim response.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // Financial instruments for reimbursement for the health care products and services specified on the claim.
  Insurance []*ClaimResponseInsurance `json:"insurance,omitempty"`

  // The party responsible for authorization, adjudication and reimbursement.
  Insurer *Reference `json:"insurer"`

  // A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
  Item []*ClaimResponseItem `json:"item,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The outcome of the claim, predetermination, or preauthorization processing.
  Outcome string `json:"outcome,omitempty"`

  // The party to whom the professional services and/or products have been supplied or are being considered and for whom actual for facast reimbursement is sought.
  Patient *Reference `json:"patient"`

  // Type of Party to be reimbursed: subscriber, provider, other.
  PayeeType *CodeableConcept `json:"payeeType,omitempty"`

  // Payment details for the adjudication of the claim.
  Payment *ClaimResponsePayment `json:"payment,omitempty"`

  // The time frame during which this authorization is effective.
  PreAuthPeriod *Period `json:"preAuthPeriod,omitempty"`

  // Reference from the Insurer which is used in later communications which refers to this adjudication.
  PreAuthRef string `json:"preAuthRef,omitempty"`

  // A note that describes or explains adjudication results in a human readable form.
  ProcessNote []*ClaimResponseProcessNote `json:"processNote,omitempty"`

  // Original request resource reference.
  Request *Reference `json:"request,omitempty"`

  // The provider which is responsible for the claim, predetermination or preauthorization.
  Requestor *Reference `json:"requestor,omitempty"`

  // This is a ClaimResponse resource
  ResourceType interface{} `json:"resourceType"`

  // The status of the resource instance.
  Status string `json:"status,omitempty"`

  // A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
  SubType *CodeableConcept `json:"subType,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`

  // Categorized monetary totals for the adjudication.
  Total []*ClaimResponseTotal `json:"total,omitempty"`

  // A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
  Type *CodeableConcept `json:"type"`

  // A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
  Use string `json:"use,omitempty"`
}

func (strct *ClaimResponse) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "addItem" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"addItem\": ")
	if tmp, err := json.Marshal(strct.AddItem); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
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
    // Marshal the "communicationRequest" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"communicationRequest\": ")
	if tmp, err := json.Marshal(strct.CommunicationRequest); err != nil {
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
    // Marshal the "disposition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"disposition\": ")
	if tmp, err := json.Marshal(strct.Disposition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "error" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"error\": ")
	if tmp, err := json.Marshal(strct.Error); err != nil {
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
    // Marshal the "form" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"form\": ")
	if tmp, err := json.Marshal(strct.Form); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "formCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"formCode\": ")
	if tmp, err := json.Marshal(strct.FormCode); err != nil {
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
    // "Insurer" field is required
    if strct.Insurer == nil {
        return nil, errors.New("insurer is a required field")
    }
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
    // Marshal the "outcome" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"outcome\": ")
	if tmp, err := json.Marshal(strct.Outcome); err != nil {
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
    // Marshal the "payeeType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"payeeType\": ")
	if tmp, err := json.Marshal(strct.PayeeType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "payment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"payment\": ")
	if tmp, err := json.Marshal(strct.Payment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "preAuthPeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"preAuthPeriod\": ")
	if tmp, err := json.Marshal(strct.PreAuthPeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "preAuthRef" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"preAuthRef\": ")
	if tmp, err := json.Marshal(strct.PreAuthRef); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "processNote" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"processNote\": ")
	if tmp, err := json.Marshal(strct.ProcessNote); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "request" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"request\": ")
	if tmp, err := json.Marshal(strct.Request); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "requestor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"requestor\": ")
	if tmp, err := json.Marshal(strct.Requestor); err != nil {
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

func (strct *ClaimResponse) UnmarshalJSON(b []byte) error {
    insurerReceived := false
    patientReceived := false
    resourceTypeReceived := false
    typeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "addItem":
            if err := json.Unmarshal([]byte(v), &strct.AddItem); err != nil {
                return err
             }
        case "adjudication":
            if err := json.Unmarshal([]byte(v), &strct.Adjudication); err != nil {
                return err
             }
        case "communicationRequest":
            if err := json.Unmarshal([]byte(v), &strct.CommunicationRequest); err != nil {
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
        case "disposition":
            if err := json.Unmarshal([]byte(v), &strct.Disposition); err != nil {
                return err
             }
        case "error":
            if err := json.Unmarshal([]byte(v), &strct.Error); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "form":
            if err := json.Unmarshal([]byte(v), &strct.Form); err != nil {
                return err
             }
        case "formCode":
            if err := json.Unmarshal([]byte(v), &strct.FormCode); err != nil {
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
        case "insurer":
            if err := json.Unmarshal([]byte(v), &strct.Insurer); err != nil {
                return err
             }
            insurerReceived = true
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
        case "outcome":
            if err := json.Unmarshal([]byte(v), &strct.Outcome); err != nil {
                return err
             }
        case "patient":
            if err := json.Unmarshal([]byte(v), &strct.Patient); err != nil {
                return err
             }
            patientReceived = true
        case "payeeType":
            if err := json.Unmarshal([]byte(v), &strct.PayeeType); err != nil {
                return err
             }
        case "payment":
            if err := json.Unmarshal([]byte(v), &strct.Payment); err != nil {
                return err
             }
        case "preAuthPeriod":
            if err := json.Unmarshal([]byte(v), &strct.PreAuthPeriod); err != nil {
                return err
             }
        case "preAuthRef":
            if err := json.Unmarshal([]byte(v), &strct.PreAuthRef); err != nil {
                return err
             }
        case "processNote":
            if err := json.Unmarshal([]byte(v), &strct.ProcessNote); err != nil {
                return err
             }
        case "request":
            if err := json.Unmarshal([]byte(v), &strct.Request); err != nil {
                return err
             }
        case "requestor":
            if err := json.Unmarshal([]byte(v), &strct.Requestor); err != nil {
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
    // check if insurer (a required property) was received
    if !insurerReceived {
        return errors.New("\"insurer\" is required but was not present")
    }
    // check if patient (a required property) was received
    if !patientReceived {
        return errors.New("\"patient\" is required but was not present")
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
