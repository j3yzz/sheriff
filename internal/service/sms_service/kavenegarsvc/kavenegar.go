package kavenegarsvc

import (
	"fmt"
	"github.com/kavenegar/kavenegar-go"
)

type KavenegarSvc struct {
	Service  *kavenegar.Kavenegar
	Template string
}

func New(cfg Config) *KavenegarSvc {
	return &KavenegarSvc{
		Service:  kavenegar.New(cfg.ApiKey),
		Template: cfg.Template,
	}
}

func (svc *KavenegarSvc) Send(receptors []string, message string) {
	if res, err := svc.Service.Message.Send("", receptors, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
		}
	}
}

func SendOTP(receptor string, token string) {
	//api := kavenegar.New(apiKey)
	//params := &kavenegar.VerifyLookupParam{}
	//
	//if res, err := api.Verify.Lookup(receptor, template, token, params); err != nil {
	//	switch err := err.(type) {
	//	case *kavenegar.APIError:
	//		fmt.Println(err)
	//		// do something
	//	case *kavenegar.HTTPError:
	//		fmt.Println(err)
	//		// do something
	//	default:
	//		fmt.Println(err)
	//		// do something
	//	}
	//} else {
	//	// do something
	//	fmt.Println(res.MessageID, res.Status)
	//}

}
