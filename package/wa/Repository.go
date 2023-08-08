package wa

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"learn_native/api/contract"
	"learn_native/package/entites"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	SendWARepository(user *entites.Users, waContract contract.SendWAMessage, otp *entites.Otp) error
	CheckOtpRepository(otpCode contract.OtpCheck) (*entites.Otp, error)
}

type repository struct {
	WaURL      string
	Collection *mongo.Collection
}

func NewRepo(waUrl string, collection *mongo.Collection) Repository {
	return &repository{
		WaURL:      waUrl,
		Collection: collection,
	}
}

func (r *repository) SendWARepository(user *entites.Users, waContract contract.SendWAMessage, otp *entites.Otp) error {
	encode, _ := json.Marshal(waContract)
	bodyReader := bytes.NewBuffer(encode)

	url := r.WaURL + "v1/messages"
	request, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("D360-API-KEY", "OkhOfL_sandbox")
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return err
	}

	_, err = r.Collection.InsertOne(context.Background(), otp)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CheckOtpRepository(otpCode contract.OtpCheck) (*entites.Otp, error) {
	var otp entites.Otp
	err := r.Collection.FindOne(context.TODO(), bson.D{{Key: "otp_code", Value: otpCode.OtpCode}}).Decode(&otp)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Otp not found!")
		}
	}

	return &otp, nil
}
