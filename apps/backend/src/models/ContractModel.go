package models

type Contract struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Image       string `json:"image" bson:"image"`
	ExternalUrl string `json:"externalUrl" bson:"externalUrl"`
}

type ContractResponse struct {
    Success bool `json:"success"`
    Result  struct {
        Name            string `json:"name"`
        Description     string `json:"description"`
        ID              string `json:"id"`
        SecretType      string `json:"secretType"`
        Symbol          string `json:"symbol"`
        ExternalUrl     string `json:"externalUrl"`
        Image           string `json:"image"`
        TransactionHash string `json:"transactionHash"`
        Status          string `json:"status"`
        Storage         struct {
            Type     string `json:"type"`
            Location string `json:"location"`
        } `json:"storage"`
        ContractUri   string `json:"contractUri"`
        ExternalLink  string `json:"external_link"`
    } `json:"result"`
}
