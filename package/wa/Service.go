package wa

import (
	"learn_native/api/contract"
	"learn_native/package/entites"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	SendWA(user entites.Users, message string) error
	CheckOtpService(otpContract contract.OtpCheck) (*entites.Otp, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) SendWA(user entites.Users, message string) error {
	otpCode := 21233
	var otpUser entites.Otp
	otpUser.ID = primitive.NewObjectID()
	otpUser.OtpCode = otpCode
	otpUser.UserID = user.ID

	messages := ""
	if message != "" {
		messages = message
	} else {
		messages = "Halo " + user.Name + " \n" + "jangan berikan kode otp ini kesiapapun " + strconv.Itoa(otpCode)
	}

	waContract := &contract.SendWAMessage{
		ReceipientType: "individual",
		To:             strconv.Itoa(user.PhoneNumber),
		Type:           "text",
		Text: contract.Message{
			Body: messages,
		},
	}

	return s.repository.SendWARepository(&user, *waContract, &otpUser)
}

func (s *service) CheckOtpService(otpContract contract.OtpCheck) (*entites.Otp, error) {
	return s.repository.CheckOtpRepository(otpContract)
}
