package types

type Risk struct {
	CategoryId                      string                     `yaml:"category,omitempty" json:"category,omitempty"`       // used for better JSON marshalling, is assigned in risk evaluation phase automatically
	RiskStatus                      RiskStatus                 `yaml:"risk_status,omitempty" json:"risk_status,omitempty"` // used for better JSON marshalling, is assigned in risk evaluation phase automatically
	Severity                        RiskSeverity               `yaml:"severity,omitempty" json:"severity,omitempty"`
	ExploitationLikelihood          RiskExploitationLikelihood `yaml:"exploitation_likelihood,omitempty" json:"exploitation_likelihood,omitempty"`
	ExploitationImpact              RiskExploitationImpact     `yaml:"exploitation_impact,omitempty" json:"exploitation_impact,omitempty"`
	Title                           string                     `yaml:"title,omitempty" json:"title,omitempty"`
	SyntheticId                     string                     `yaml:"synthetic_id,omitempty" json:"synthetic_id,omitempty"`
	MostRelevantDataAssetId         string                     `yaml:"most_relevant_data_asset,omitempty" json:"most_relevant_data_asset,omitempty"`
	MostRelevantTechnicalAssetId    string                     `yaml:"most_relevant_technical_asset,omitempty" json:"most_relevant_technical_asset,omitempty"`
	MostRelevantTrustBoundaryId     string                     `yaml:"most_relevant_trust_boundary,omitempty" json:"most_relevant_trust_boundary,omitempty"`
	MostRelevantSharedRuntimeId     string                     `yaml:"most_relevant_shared_runtime,omitempty" json:"most_relevant_shared_runtime,omitempty"`
	MostRelevantCommunicationLinkId string                     `yaml:"most_relevant_communication_link,omitempty" json:"most_relevant_communication_link,omitempty"`
	DataBreachProbability           DataBreachProbability      `yaml:"data_breach_probability,omitempty" json:"data_breach_probability,omitempty"`
	DataBreachTechnicalAssetIDs     []string                   `yaml:"data_breach_technical_assets,omitempty" json:"data_breach_technical_assets,omitempty"`
	// TODO: refactor all "ID" here to "ID"?
}

func (what Risk) GetRiskTracking(model *Model) *RiskTracking { // TODO: Unify function naming regarding Get etc.
	if riskTracking, ok := model.RiskTracking[what.SyntheticId]; ok {
		return riskTracking
	}
	return nil
}

func (what Risk) GetRiskTrackingWithDefault(model *Model) RiskTracking { // TODO: Unify function naming regarding Get etc.
	if riskTracking, ok := model.RiskTracking[what.SyntheticId]; ok {
		return *riskTracking
	}
	return RiskTracking{}
}

func (what Risk) IsRiskTracked(model *Model) bool {
	if _, ok := model.RiskTracking[what.SyntheticId]; ok {
		return true
	}
	return false
}
