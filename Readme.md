# co

[![GoDoc](https://godoc.org/github.com/nowk/co?status.svg)](http://godoc.org/github.com/nowk/co)

Signature utility around `crypto/hmac`

## Examples

Given a struct

    type S3 struct {
      Filename string
      Mime     string
      Expires  int64
    }

    func (s S3) Message() ([]byte, error) {
      return []byte(strings.Join([]string{
        "GET",
        "",
        s.Mime,
        strconv.FormatInt(s.Expires, 10),
        s.Path(),
      }, "\n")), nil
    }

    func (s S3) Path() string {
      return fmt.Sprintf("bucket/%s", s.Filename)
    }

---

Pass a struct that implements `co.Messenger` and provide your `sha` & `key`

    sig, err := co.Sign(S3{}, sha1.New, "asecret")
    if err != nil {
      // handle err
    }

---

Signature can be `base64'd`

    b, err := sig.Base64()
    if err != nil {
      // handle err
    }



## License

MIT