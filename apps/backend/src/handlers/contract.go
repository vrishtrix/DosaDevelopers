package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"dosadevelopers.devsoc/backend/src/config"
	"dosadevelopers.devsoc/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var accessToken string

func AuthenticateWithVenly() {
	client := &http.Client{}
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", config.Client_ID)
	data.Set("client_secret", config.Client_Secret)

	// Encode the data and create a new request
	encodedData := strings.NewReader(data.Encode())
	request, err := http.NewRequest("POST", config.Login_URL, encodedData)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	if token, exists := result["access_token"]; exists {
		accessToken = token.(string)
		log.Println("Access token:", accessToken)
	} else {
		log.Fatal("Access token not found in response")
	}
}

func ContractDeployment(c *fiber.Ctx) error {
	var contract models.Contract
	if err := c.BodyParser(&contract); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	existingName := bson.M{"name": contract.Name}
	counter, err := contractcollection.CountDocuments(context.Background(), existingName)
	if err != nil {
		log.Fatal(err)
	}
	if counter > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name already exists",
		})
	}

	newContract := bson.M{
		"name":        contract.Name,
		"description": contract.Description,
		"image":       contract.Image,
		"externalUrl": contract.ExternalUrl,
		"chain":       "MATIC",
	}

	// Here you make the request to the external API.
	err = deployContract(newContract)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to deploy contract",
		})
	}

	_, err = contractcollection.InsertOne(context.Background(), newContract)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{
		"message": "Contract deployment initiated successfully",
	})
}

func deployContract(contractData bson.M) error {
	url := config.Contract_URL
	jsonData, err := json.Marshal(contractData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken) // Replace <YourAuthTokenHere> with your actual auth token

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var responseObj models.ContractResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	_, err = contractrespcollection.InsertOne(context.Background(), bson.M{
		"name":            responseObj.Result.Name,
		"description":     responseObj.Result.Description,
		"id":              responseObj.Result.ID,
		"secretType":      responseObj.Result.SecretType,
		"symbol":          responseObj.Result.Symbol,
		"externalUrl":     responseObj.Result.ExternalUrl,
		"image":           responseObj.Result.Image,
		"transactionHash": responseObj.Result.TransactionHash,
		"status":          responseObj.Result.Status,
		"storageType":     responseObj.Result.Storage.Type,
		"storageLocation": responseObj.Result.Storage.Location,
		"contractUri":     responseObj.Result.ContractUri,
		"externalLink":    responseObj.Result.ExternalLink,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response from API:", string(body))
	if resp.StatusCode != http.StatusOK {
		// Handle non-successful response
		return fmt.Errorf("API responded with status code: %d", resp.StatusCode)
	}

	return nil
}
