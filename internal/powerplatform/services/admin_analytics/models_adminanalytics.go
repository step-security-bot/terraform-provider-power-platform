package powerplatform

type AppInsightConnectionDto struct {
	Source       string
	Environments []EnvironmentDto
	Sink         SinkDto
	Scenarios    []string
	Status       string
	PackageName  string
}

type AppInsightConnectionDtoArray struct {
	Value []AppInsightConnectionDto `json:"value"`
}

type EnvironmentDto struct {
	EnvironmentId  string
	OrganizationId string
}

type SinkDto struct {
	Key          string
	Type         string
	Id           string
	ResourceName string
}
