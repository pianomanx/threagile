/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package types

import (
	"encoding/json"
	"errors"
	"strings"
)

type RiskExploitationLikelihood int

const (
	Unlikely RiskExploitationLikelihood = iota
	Likely
	VeryLikely
	Frequent
)

func RiskExploitationLikelihoodValues() []TypeEnum {
	return []TypeEnum{
		Unlikely,
		Likely,
		VeryLikely,
		Frequent,
	}
}

var RiskExploitationLikelihoodTypeDescription = [...]TypeDescription{
	{"unlikely", "Unlikely"},
	{"likely", "Likely"},
	{"very-likely", "Very-Likely"},
	{"frequent", "Frequent"},
}

func ParseRiskExploitationLikelihood(value string) (riskExploitationLikelihood RiskExploitationLikelihood, err error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Likely, nil
	}
	for _, candidate := range RiskExploitationLikelihoodValues() {
		if candidate.String() == value {
			return candidate.(RiskExploitationLikelihood), err
		}
	}
	return riskExploitationLikelihood, errors.New("Unable to parse into type: " + value)
}

func (what RiskExploitationLikelihood) String() string {
	// NOTE: maintain list also in schema.json for validation in IDEs
	return RiskExploitationLikelihoodTypeDescription[what].Name
}

func (what RiskExploitationLikelihood) Explain() string {
	return RiskExploitationLikelihoodTypeDescription[what].Description
}

func (what RiskExploitationLikelihood) Title() string {
	return [...]string{"Unlikely", "Likely", "Very Likely", "Frequent"}[what]
}

func (what RiskExploitationLikelihood) Weight() int {
	return [...]int{1, 2, 3, 4}[what]
}

func (what RiskExploitationLikelihood) MarshalJSON() ([]byte, error) {
	return json.Marshal(what.String())
}
