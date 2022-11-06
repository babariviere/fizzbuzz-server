package fizzbuzz

import "errors"

type GetFizzbuzzReq struct {
    Int1 int `form:"int1"`
    Int2 int `form:"int2"`
    Limit int `form:"limit"`
    Str1 string `form:"str1"`
    Str2 string `form:"str2"`
}

func (req GetFizzbuzzReq) Validate() error {
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

// OrDefault fills default parameters if every values are null (except limit)
func (req *GetFizzbuzzReq) OrDefault() {
    if req.Int1 != 0 || req.Int2 != 0 || req.Str1 != "" || req.Str2 != "" {
        return
    }
    req.Int1 = 3
    req.Int2 = 5
    req.Str1 = "fizz"
    req.Str2 = "buzz"
}
