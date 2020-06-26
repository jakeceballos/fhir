// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "errors"
    "encoding/json"
    "fmt"
    "bytes"
)

// DiagnosticReport The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
type DiagnosticReport struct {

  // Details concerning a service requested.
  BasedOn []*Reference `json:"basedOn,omitempty"`

  // A code that classifies the clinical discipline, department or diagnostic service that created the report (e.g. cardiology, biochemistry, hematology, MRI). This is used for searching, sorting and display purposes.
  Category []*CodeableConcept `json:"category,omitempty"`

  // A code or name that describes this diagnostic report.
  Code *CodeableConcept `json:"code"`

  // Concise and clinically contextualized summary conclusion (interpretation/impression) of the diagnostic report.
  Conclusion string `json:"conclusion,omitempty"`

  // One or more codes that represent the summary conclusion (interpretation/impression) of the diagnostic report.
  ConclusionCode []*CodeableConcept `json:"conclusionCode,omitempty"`

  // These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
  Contained []interface{} `json:"contained,omitempty"`

  // The time or time-period the observed values are related to. When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
  EffectiveDateTime string `json:"effectiveDateTime,omitempty"`

  // The time or time-period the observed values are related to. When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
  EffectivePeriod *Period `json:"effectivePeriod,omitempty"`

  // The healthcare event  (e.g. a patient and healthcare provider interaction) which this DiagnosticReport is about.
  Encounter *Reference `json:"encounter,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
  Id string `json:"id,omitempty"`

  // Identifiers assigned to this report by the performer or other systems.
  Identifier []*Identifier `json:"identifier,omitempty"`

  // One or more links to full details of any imaging performed during the diagnostic investigation. Typically, this is imaging performed by DICOM enabled modalities, but this is not required. A fully enabled PACS viewer can use this information to provide views of the source images.
  ImagingStudy []*Reference `json:"imagingStudy,omitempty"`

  // A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
  ImplicitRules string `json:"implicitRules,omitempty"`

  // The date and time that this version of the report was made available to providers, typically after the report was reviewed and verified.
  Issued string `json:"issued,omitempty"`

  // The base language in which the resource is written.
  Language string `json:"language,omitempty"`

  // A list of key images associated with this report. The images are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
  Media []*DiagnosticReportMedia `json:"media,omitempty"`

  // The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
  Meta *Meta `json:"meta,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The diagnostic service that is responsible for issuing the report.
  Performer []*Reference `json:"performer,omitempty"`

  // Rich text representation of the entire result as issued by the diagnostic service. Multiple formats are allowed but they SHALL be semantically equivalent.
  PresentedForm []*Attachment `json:"presentedForm,omitempty"`

  // This is a DiagnosticReport resource
  ResourceType interface{} `json:"resourceType"`

  // [Observations](observation.html)  that are part of this diagnostic report.
  Result []*Reference `json:"result,omitempty"`

  // The practitioner or organization that is responsible for the report's conclusions and interpretations.
  ResultsInterpreter []*Reference `json:"resultsInterpreter,omitempty"`

  // Details about the specimens on which this diagnostic report is based.
  Specimen []*Reference `json:"specimen,omitempty"`

  // The status of the diagnostic report.
  Status interface{} `json:"status,omitempty"`

  // The subject of the report. Usually, but not always, this is a patient. However, diagnostic services also perform analyses on specimens collected from a variety of other sources.
  Subject *Reference `json:"subject,omitempty"`

  // A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
  Text *Narrative `json:"text,omitempty"`
}

func (strct *DiagnosticReport) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // "Code" field is required
    if strct.Code == nil {
        return nil, errors.New("code is a required field")
    }
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
    // Marshal the "conclusion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"conclusion\": ")
	if tmp, err := json.Marshal(strct.Conclusion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "conclusionCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"conclusionCode\": ")
	if tmp, err := json.Marshal(strct.ConclusionCode); err != nil {
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
    // Marshal the "effectiveDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"effectiveDateTime\": ")
	if tmp, err := json.Marshal(strct.EffectiveDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "effectivePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"effectivePeriod\": ")
	if tmp, err := json.Marshal(strct.EffectivePeriod); err != nil {
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
    // Marshal the "imagingStudy" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"imagingStudy\": ")
	if tmp, err := json.Marshal(strct.ImagingStudy); err != nil {
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
    // Marshal the "issued" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"issued\": ")
	if tmp, err := json.Marshal(strct.Issued); err != nil {
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
    // Marshal the "media" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"media\": ")
	if tmp, err := json.Marshal(strct.Media); err != nil {
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
    // Marshal the "presentedForm" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"presentedForm\": ")
	if tmp, err := json.Marshal(strct.PresentedForm); err != nil {
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
    // Marshal the "result" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"result\": ")
	if tmp, err := json.Marshal(strct.Result); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "resultsInterpreter" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"resultsInterpreter\": ")
	if tmp, err := json.Marshal(strct.ResultsInterpreter); err != nil {
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

func (strct *DiagnosticReport) UnmarshalJSON(b []byte) error {
    codeReceived := false
    resourceTypeReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "basedOn":
            if err := json.Unmarshal([]byte(v), &strct.BasedOn); err != nil {
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
            codeReceived = true
        case "conclusion":
            if err := json.Unmarshal([]byte(v), &strct.Conclusion); err != nil {
                return err
             }
        case "conclusionCode":
            if err := json.Unmarshal([]byte(v), &strct.ConclusionCode); err != nil {
                return err
             }
        case "contained":
            if err := json.Unmarshal([]byte(v), &strct.Contained); err != nil {
                return err
             }
        case "effectiveDateTime":
            if err := json.Unmarshal([]byte(v), &strct.EffectiveDateTime); err != nil {
                return err
             }
        case "effectivePeriod":
            if err := json.Unmarshal([]byte(v), &strct.EffectivePeriod); err != nil {
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
        case "imagingStudy":
            if err := json.Unmarshal([]byte(v), &strct.ImagingStudy); err != nil {
                return err
             }
        case "implicitRules":
            if err := json.Unmarshal([]byte(v), &strct.ImplicitRules); err != nil {
                return err
             }
        case "issued":
            if err := json.Unmarshal([]byte(v), &strct.Issued); err != nil {
                return err
             }
        case "language":
            if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
                return err
             }
        case "media":
            if err := json.Unmarshal([]byte(v), &strct.Media); err != nil {
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
        case "performer":
            if err := json.Unmarshal([]byte(v), &strct.Performer); err != nil {
                return err
             }
        case "presentedForm":
            if err := json.Unmarshal([]byte(v), &strct.PresentedForm); err != nil {
                return err
             }
        case "resourceType":
            if err := json.Unmarshal([]byte(v), &strct.ResourceType); err != nil {
                return err
             }
            resourceTypeReceived = true
        case "result":
            if err := json.Unmarshal([]byte(v), &strct.Result); err != nil {
                return err
             }
        case "resultsInterpreter":
            if err := json.Unmarshal([]byte(v), &strct.ResultsInterpreter); err != nil {
                return err
             }
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
        case "text":
            if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if code (a required property) was received
    if !codeReceived {
        return errors.New("\"code\" is required but was not present")
    }
    // check if resourceType (a required property) was received
    if !resourceTypeReceived {
        return errors.New("\"resourceType\" is required but was not present")
    }
    return nil
}
