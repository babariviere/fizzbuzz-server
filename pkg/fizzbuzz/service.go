package fizzbuzz

import (
	"strconv"
	"strings"
)


type Service interface {
    GetFizzbuzz(req GetFizzbuzzReq) (string, error)
}

func NewFizzbuzz() Service {
    return service{}
}

type service struct {}

func (s service) GetFizzbuzz(req GetFizzbuzzReq) (string, error) {
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
