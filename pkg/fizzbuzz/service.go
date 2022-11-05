package fizzbuzz

import (
	"errors"
	"strconv"
	"strings"
)

type FizzbuzzRequest struct {
    Int1 int
    Int2 int
    Limit int
    Str1 string
    Str2 string
}

func (req FizzbuzzRequest) Validate() error {
    if req.Int1 <= 0 {
        return errors.New("int1 must be superior to 0")
    }
    if req.Int2 <= 0 {
        return errors.New("int2 must be superior to 0")
    }
    if req.Str1 == "" {
        return errors.New("str1 must be a non empty string")
    }
    if req.Str2 == "" {
        return errors.New("str2 must be a non empty string")
    }
    return nil
}


type Service interface {
    GetFizzbuzz(req FizzbuzzRequest) (string, error)
}

func NewFizzbuzz() Service {
    return service{}
}

type service struct {}

func (s service) GetFizzbuzz(req FizzbuzzRequest) (string, error) {
    var sb strings.Builder

    if err := req.Validate(); err != nil {
        return "", err
    }

    for i := 1; i <= req.Limit; i++ {
        if i >= 2 {
            sb.WriteByte(',')
        }

        switch {
        case i % req.Int1 == 0 && i % req.Int2 == 0:
            sb.WriteString(req.Str1)
            sb.WriteString(req.Str2)
        case i % req.Int1 == 0:
            sb.WriteString(req.Str1)
        case i % req.Int2 == 0:
            sb.WriteString(req.Str2)
        default:
            sb.WriteString(strconv.Itoa(i))
        }
    }

    return sb.String(), nil
}
