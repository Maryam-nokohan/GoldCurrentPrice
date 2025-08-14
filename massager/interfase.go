package massager

import "github.com/Maryam-nokohan/GoldCurrentPrice/model"


type Massenger interface{
	Send(to string, message *model.Message)error
}