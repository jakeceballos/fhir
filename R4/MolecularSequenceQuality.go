// Code generated by schema-generate. DO NOT EDIT.

package R4

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// MolecularSequenceQuality Raw data describing a biological sequence.
type MolecularSequenceQuality struct {

  // End position of the sequence. If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
  End float64 `json:"end,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
  Extension []*Extension `json:"extension,omitempty"`

  // Harmonic mean of Recall and Precision, computed as: 2 * precision * recall / (precision + recall).
  FScore float64 `json:"fScore,omitempty"`

  // The number of false positives where the non-REF alleles in the Truth and Query Call Sets match (i.e. cases where the truth is 1/1 and the query is 0/1 or similar).
  GtFP float64 `json:"gtFP,omitempty"`

  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
  Id string `json:"id,omitempty"`

  // Which method is used to get sequence quality.
  Method *CodeableConcept `json:"method,omitempty"`

  // May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
  // 
  // Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
  ModifierExtension []*Extension `json:"modifierExtension,omitempty"`

  // QUERY.TP / (QUERY.TP + QUERY.FP).
  Precision float64 `json:"precision,omitempty"`

  // False positives, i.e. the number of sites in the Query Call Set for which there is no path through the Truth Call Set that is consistent with this site. Sites with correct variant but incorrect genotype are counted here.
  QueryFP float64 `json:"queryFP,omitempty"`

  // True positives, from the perspective of the query data, i.e. the number of sites in the Query Call Set for which there are paths through the Truth Call Set that are consistent with all of the alleles at this site, and for which there is an accurate genotype call for the event.
  QueryTP float64 `json:"queryTP,omitempty"`

  // TRUTH.TP / (TRUTH.TP + TRUTH.FN).
  Recall float64 `json:"recall,omitempty"`

  // Receiver Operator Characteristic (ROC) Curve  to give sensitivity/specificity tradeoff.
  Roc *MolecularSequenceRoc `json:"roc,omitempty"`

  // The score of an experimentally derived feature such as a p-value ([SO:0001685](http://www.sequenceontology.org/browser/current_svn/term/SO:0001685)).
  Score *Quantity `json:"score,omitempty"`

  // Gold standard sequence used for comparing against.
  StandardSequence *CodeableConcept `json:"standardSequence,omitempty"`

  // Start position of the sequence. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
  Start float64 `json:"start,omitempty"`

  // False negatives, i.e. the number of sites in the Truth Call Set for which there is no path through the Query Call Set that is consistent with all of the alleles at this site, or sites for which there is an inaccurate genotype call for the event. Sites with correct variant but incorrect genotype are counted here.
  TruthFN float64 `json:"truthFN,omitempty"`

  // True positives, from the perspective of the truth data, i.e. the number of sites in the Truth Call Set for which there are paths through the Query Call Set that are consistent with all of the alleles at this site, and for which there is an accurate genotype call for the event.
  TruthTP float64 `json:"truthTP,omitempty"`

  // INDEL / SNP / Undefined variant.
  Type interface{} `json:"type,omitempty"`
}

func (strct *MolecularSequenceQuality) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
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
    // Marshal the "fScore" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"fScore\": ")
	if tmp, err := json.Marshal(strct.FScore); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "gtFP" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"gtFP\": ")
	if tmp, err := json.Marshal(strct.GtFP); err != nil {
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
    // Marshal the "method" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"method\": ")
	if tmp, err := json.Marshal(strct.Method); err != nil {
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
    // Marshal the "precision" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"precision\": ")
	if tmp, err := json.Marshal(strct.Precision); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "queryFP" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"queryFP\": ")
	if tmp, err := json.Marshal(strct.QueryFP); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "queryTP" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"queryTP\": ")
	if tmp, err := json.Marshal(strct.QueryTP); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "recall" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"recall\": ")
	if tmp, err := json.Marshal(strct.Recall); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "roc" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"roc\": ")
	if tmp, err := json.Marshal(strct.Roc); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "score" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"score\": ")
	if tmp, err := json.Marshal(strct.Score); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "standardSequence" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"standardSequence\": ")
	if tmp, err := json.Marshal(strct.StandardSequence); err != nil {
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
    // Marshal the "truthFN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"truthFN\": ")
	if tmp, err := json.Marshal(strct.TruthFN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "truthTP" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"truthTP\": ")
	if tmp, err := json.Marshal(strct.TruthTP); err != nil {
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

func (strct *MolecularSequenceQuality) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "end":
            if err := json.Unmarshal([]byte(v), &strct.End); err != nil {
                return err
             }
        case "extension":
            if err := json.Unmarshal([]byte(v), &strct.Extension); err != nil {
                return err
             }
        case "fScore":
            if err := json.Unmarshal([]byte(v), &strct.FScore); err != nil {
                return err
             }
        case "gtFP":
            if err := json.Unmarshal([]byte(v), &strct.GtFP); err != nil {
                return err
             }
        case "id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
        case "method":
            if err := json.Unmarshal([]byte(v), &strct.Method); err != nil {
                return err
             }
        case "modifierExtension":
            if err := json.Unmarshal([]byte(v), &strct.ModifierExtension); err != nil {
                return err
             }
        case "precision":
            if err := json.Unmarshal([]byte(v), &strct.Precision); err != nil {
                return err
             }
        case "queryFP":
            if err := json.Unmarshal([]byte(v), &strct.QueryFP); err != nil {
                return err
             }
        case "queryTP":
            if err := json.Unmarshal([]byte(v), &strct.QueryTP); err != nil {
                return err
             }
        case "recall":
            if err := json.Unmarshal([]byte(v), &strct.Recall); err != nil {
                return err
             }
        case "roc":
            if err := json.Unmarshal([]byte(v), &strct.Roc); err != nil {
                return err
             }
        case "score":
            if err := json.Unmarshal([]byte(v), &strct.Score); err != nil {
                return err
             }
        case "standardSequence":
            if err := json.Unmarshal([]byte(v), &strct.StandardSequence); err != nil {
                return err
             }
        case "start":
            if err := json.Unmarshal([]byte(v), &strct.Start); err != nil {
                return err
             }
        case "truthFN":
            if err := json.Unmarshal([]byte(v), &strct.TruthFN); err != nil {
                return err
             }
        case "truthTP":
            if err := json.Unmarshal([]byte(v), &strct.TruthTP); err != nil {
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
    return nil
}
