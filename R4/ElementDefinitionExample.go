// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// ElementDefinitionExample Captures constraints on each element within the resource, profile, or extension.
type ElementDefinitionExample struct {

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Describes the purpose of this example amoung the set of examples.
  Label string `json:"label,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueAddress *Address `json:"valueAddress,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueAge *Age `json:"valueAge,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueAttachment *Attachment `json:"valueAttachment,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueBase64Binary string `json:"valueBase64Binary,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueBoolean bool `json:"valueBoolean,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueCanonical string `json:"valueCanonical,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueCode string `json:"valueCode,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueCoding *Coding `json:"valueCoding,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueContributor *Contributor `json:"valueContributor,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueCount *Count `json:"valueCount,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueDate string `json:"valueDate,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueDateTime string `json:"valueDateTime,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueDecimal float64 `json:"valueDecimal,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueDistance *Distance `json:"valueDistance,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueDosage *Dosage `json:"valueDosage,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueDuration *Duration `json:"valueDuration,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueExpression *Expression `json:"valueExpression,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueHumanName *HumanName `json:"valueHumanName,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueId string `json:"valueId,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueInstant string `json:"valueInstant,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueInteger float64 `json:"valueInteger,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueMarkdown string `json:"valueMarkdown,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueMoney *Money `json:"valueMoney,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueOid string `json:"valueOid,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValuePeriod *Period `json:"valuePeriod,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValuePositiveInt float64 `json:"valuePositiveInt,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueQuantity *Quantity `json:"valueQuantity,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueRange *Range `json:"valueRange,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueRatio *Ratio `json:"valueRatio,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueReference *Reference `json:"valueReference,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueSampledData *SampledData `json:"valueSampledData,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueSignature *Signature `json:"valueSignature,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueString string `json:"valueString,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueTime string `json:"valueTime,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueTiming *Timing `json:"valueTiming,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueUnsignedInt float64 `json:"valueUnsignedInt,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueUri string `json:"valueUri,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueUrl string `json:"valueUrl,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`

  // The actual value for the element, which must be one of the types allowed for this element.
  ValueUuid string `json:"valueUuid,omitempty"`
}

func (strct *ElementDefinitionExample) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "label" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"label\": ")
	if tmp, err := json.Marshal(strct.Label); err != nil {
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
    // Marshal the "valueAddress" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueAddress\": ")
	if tmp, err := json.Marshal(strct.ValueAddress); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueAge" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueAge\": ")
	if tmp, err := json.Marshal(strct.ValueAge); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueAnnotation" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueAnnotation\": ")
	if tmp, err := json.Marshal(strct.ValueAnnotation); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueAttachment" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueAttachment\": ")
	if tmp, err := json.Marshal(strct.ValueAttachment); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueBase64Binary" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueBase64Binary\": ")
	if tmp, err := json.Marshal(strct.ValueBase64Binary); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueBoolean" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueBoolean\": ")
	if tmp, err := json.Marshal(strct.ValueBoolean); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueCanonical" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueCanonical\": ")
	if tmp, err := json.Marshal(strct.ValueCanonical); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueCode" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueCode\": ")
	if tmp, err := json.Marshal(strct.ValueCode); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueCodeableConcept" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueCodeableConcept\": ")
	if tmp, err := json.Marshal(strct.ValueCodeableConcept); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueCoding" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueCoding\": ")
	if tmp, err := json.Marshal(strct.ValueCoding); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueContactDetail" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueContactDetail\": ")
	if tmp, err := json.Marshal(strct.ValueContactDetail); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueContactPoint" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueContactPoint\": ")
	if tmp, err := json.Marshal(strct.ValueContactPoint); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueContributor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueContributor\": ")
	if tmp, err := json.Marshal(strct.ValueContributor); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueCount" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueCount\": ")
	if tmp, err := json.Marshal(strct.ValueCount); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDataRequirement" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDataRequirement\": ")
	if tmp, err := json.Marshal(strct.ValueDataRequirement); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDate" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDate\": ")
	if tmp, err := json.Marshal(strct.ValueDate); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDateTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDateTime\": ")
	if tmp, err := json.Marshal(strct.ValueDateTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDecimal" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDecimal\": ")
	if tmp, err := json.Marshal(strct.ValueDecimal); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDistance" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDistance\": ")
	if tmp, err := json.Marshal(strct.ValueDistance); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDosage" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDosage\": ")
	if tmp, err := json.Marshal(strct.ValueDosage); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueDuration" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueDuration\": ")
	if tmp, err := json.Marshal(strct.ValueDuration); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueExpression" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueExpression\": ")
	if tmp, err := json.Marshal(strct.ValueExpression); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueHumanName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueHumanName\": ")
	if tmp, err := json.Marshal(strct.ValueHumanName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueId\": ")
	if tmp, err := json.Marshal(strct.ValueId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueIdentifier" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueIdentifier\": ")
	if tmp, err := json.Marshal(strct.ValueIdentifier); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueInstant" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueInstant\": ")
	if tmp, err := json.Marshal(strct.ValueInstant); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueInteger" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueInteger\": ")
	if tmp, err := json.Marshal(strct.ValueInteger); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueMarkdown" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueMarkdown\": ")
	if tmp, err := json.Marshal(strct.ValueMarkdown); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueMoney" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueMoney\": ")
	if tmp, err := json.Marshal(strct.ValueMoney); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueOid" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueOid\": ")
	if tmp, err := json.Marshal(strct.ValueOid); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueParameterDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueParameterDefinition\": ")
	if tmp, err := json.Marshal(strct.ValueParameterDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valuePeriod" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valuePeriod\": ")
	if tmp, err := json.Marshal(strct.ValuePeriod); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valuePositiveInt" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valuePositiveInt\": ")
	if tmp, err := json.Marshal(strct.ValuePositiveInt); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueQuantity" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueQuantity\": ")
	if tmp, err := json.Marshal(strct.ValueQuantity); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueRange" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueRange\": ")
	if tmp, err := json.Marshal(strct.ValueRange); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueRatio" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueRatio\": ")
	if tmp, err := json.Marshal(strct.ValueRatio); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueReference" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueReference\": ")
	if tmp, err := json.Marshal(strct.ValueReference); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueRelatedArtifact" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueRelatedArtifact\": ")
	if tmp, err := json.Marshal(strct.ValueRelatedArtifact); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueSampledData" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueSampledData\": ")
	if tmp, err := json.Marshal(strct.ValueSampledData); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueSignature" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueSignature\": ")
	if tmp, err := json.Marshal(strct.ValueSignature); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueString" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueString\": ")
	if tmp, err := json.Marshal(strct.ValueString); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueTime" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueTime\": ")
	if tmp, err := json.Marshal(strct.ValueTime); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueTiming" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueTiming\": ")
	if tmp, err := json.Marshal(strct.ValueTiming); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueTriggerDefinition" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueTriggerDefinition\": ")
	if tmp, err := json.Marshal(strct.ValueTriggerDefinition); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueUnsignedInt" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueUnsignedInt\": ")
	if tmp, err := json.Marshal(strct.ValueUnsignedInt); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueUri" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueUri\": ")
	if tmp, err := json.Marshal(strct.ValueUri); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueUrl" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueUrl\": ")
	if tmp, err := json.Marshal(strct.ValueUrl); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueUsageContext" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueUsageContext\": ")
	if tmp, err := json.Marshal(strct.ValueUsageContext); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "valueUuid" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"valueUuid\": ")
	if tmp, err := json.Marshal(strct.ValueUuid); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ElementDefinitionExample) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "label":
            if err := json.Unmarshal([]byte(v), &strct.Label); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "valueAddress":
            if err := json.Unmarshal([]byte(v), &strct.ValueAddress); err != nil {
                return err
             }
        case "valueAge":
            if err := json.Unmarshal([]byte(v), &strct.ValueAge); err != nil {
                return err
             }
        case "valueAnnotation":
            if err := json.Unmarshal([]byte(v), &strct.ValueAnnotation); err != nil {
                return err
             }
        case "valueAttachment":
            if err := json.Unmarshal([]byte(v), &strct.ValueAttachment); err != nil {
                return err
             }
        case "valueBase64Binary":
            if err := json.Unmarshal([]byte(v), &strct.ValueBase64Binary); err != nil {
                return err
             }
        case "valueBoolean":
            if err := json.Unmarshal([]byte(v), &strct.ValueBoolean); err != nil {
                return err
             }
        case "valueCanonical":
            if err := json.Unmarshal([]byte(v), &strct.ValueCanonical); err != nil {
                return err
             }
        case "valueCode":
            if err := json.Unmarshal([]byte(v), &strct.ValueCode); err != nil {
                return err
             }
        case "valueCodeableConcept":
            if err := json.Unmarshal([]byte(v), &strct.ValueCodeableConcept); err != nil {
                return err
             }
        case "valueCoding":
            if err := json.Unmarshal([]byte(v), &strct.ValueCoding); err != nil {
                return err
             }
        case "valueContactDetail":
            if err := json.Unmarshal([]byte(v), &strct.ValueContactDetail); err != nil {
                return err
             }
        case "valueContactPoint":
            if err := json.Unmarshal([]byte(v), &strct.ValueContactPoint); err != nil {
                return err
             }
        case "valueContributor":
            if err := json.Unmarshal([]byte(v), &strct.ValueContributor); err != nil {
                return err
             }
        case "valueCount":
            if err := json.Unmarshal([]byte(v), &strct.ValueCount); err != nil {
                return err
             }
        case "valueDataRequirement":
            if err := json.Unmarshal([]byte(v), &strct.ValueDataRequirement); err != nil {
                return err
             }
        case "valueDate":
            if err := json.Unmarshal([]byte(v), &strct.ValueDate); err != nil {
                return err
             }
        case "valueDateTime":
            if err := json.Unmarshal([]byte(v), &strct.ValueDateTime); err != nil {
                return err
             }
        case "valueDecimal":
            if err := json.Unmarshal([]byte(v), &strct.ValueDecimal); err != nil {
                return err
             }
        case "valueDistance":
            if err := json.Unmarshal([]byte(v), &strct.ValueDistance); err != nil {
                return err
             }
        case "valueDosage":
            if err := json.Unmarshal([]byte(v), &strct.ValueDosage); err != nil {
                return err
             }
        case "valueDuration":
            if err := json.Unmarshal([]byte(v), &strct.ValueDuration); err != nil {
                return err
             }
        case "valueExpression":
            if err := json.Unmarshal([]byte(v), &strct.ValueExpression); err != nil {
                return err
             }
        case "valueHumanName":
            if err := json.Unmarshal([]byte(v), &strct.ValueHumanName); err != nil {
                return err
             }
        case "valueId":
            if err := json.Unmarshal([]byte(v), &strct.ValueId); err != nil {
                return err
             }
        case "valueIdentifier":
            if err := json.Unmarshal([]byte(v), &strct.ValueIdentifier); err != nil {
                return err
             }
        case "valueInstant":
            if err := json.Unmarshal([]byte(v), &strct.ValueInstant); err != nil {
                return err
             }
        case "valueInteger":
            if err := json.Unmarshal([]byte(v), &strct.ValueInteger); err != nil {
                return err
             }
        case "valueMarkdown":
            if err := json.Unmarshal([]byte(v), &strct.ValueMarkdown); err != nil {
                return err
             }
        case "valueMoney":
            if err := json.Unmarshal([]byte(v), &strct.ValueMoney); err != nil {
                return err
             }
        case "valueOid":
            if err := json.Unmarshal([]byte(v), &strct.ValueOid); err != nil {
                return err
             }
        case "valueParameterDefinition":
            if err := json.Unmarshal([]byte(v), &strct.ValueParameterDefinition); err != nil {
                return err
             }
        case "valuePeriod":
            if err := json.Unmarshal([]byte(v), &strct.ValuePeriod); err != nil {
                return err
             }
        case "valuePositiveInt":
            if err := json.Unmarshal([]byte(v), &strct.ValuePositiveInt); err != nil {
                return err
             }
        case "valueQuantity":
            if err := json.Unmarshal([]byte(v), &strct.ValueQuantity); err != nil {
                return err
             }
        case "valueRange":
            if err := json.Unmarshal([]byte(v), &strct.ValueRange); err != nil {
                return err
             }
        case "valueRatio":
            if err := json.Unmarshal([]byte(v), &strct.ValueRatio); err != nil {
                return err
             }
        case "valueReference":
            if err := json.Unmarshal([]byte(v), &strct.ValueReference); err != nil {
                return err
             }
        case "valueRelatedArtifact":
            if err := json.Unmarshal([]byte(v), &strct.ValueRelatedArtifact); err != nil {
                return err
             }
        case "valueSampledData":
            if err := json.Unmarshal([]byte(v), &strct.ValueSampledData); err != nil {
                return err
             }
        case "valueSignature":
            if err := json.Unmarshal([]byte(v), &strct.ValueSignature); err != nil {
                return err
             }
        case "valueString":
            if err := json.Unmarshal([]byte(v), &strct.ValueString); err != nil {
                return err
             }
        case "valueTime":
            if err := json.Unmarshal([]byte(v), &strct.ValueTime); err != nil {
                return err
             }
        case "valueTiming":
            if err := json.Unmarshal([]byte(v), &strct.ValueTiming); err != nil {
                return err
             }
        case "valueTriggerDefinition":
            if err := json.Unmarshal([]byte(v), &strct.ValueTriggerDefinition); err != nil {
                return err
             }
        case "valueUnsignedInt":
            if err := json.Unmarshal([]byte(v), &strct.ValueUnsignedInt); err != nil {
                return err
             }
        case "valueUri":
            if err := json.Unmarshal([]byte(v), &strct.ValueUri); err != nil {
                return err
             }
        case "valueUrl":
            if err := json.Unmarshal([]byte(v), &strct.ValueUrl); err != nil {
                return err
             }
        case "valueUsageContext":
            if err := json.Unmarshal([]byte(v), &strct.ValueUsageContext); err != nil {
                return err
             }
        case "valueUuid":
            if err := json.Unmarshal([]byte(v), &strct.ValueUuid); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}