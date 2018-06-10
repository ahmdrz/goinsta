package goinsta

import (
	"encoding/json"
	"image"
	// Required for getImageDimensionFromReader in jpg and png format
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strconv"
	"unsafe"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func toString(i interface{}) string {
	switch s := i.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatInt(int64(s), 10)
	case uint64:
		return strconv.FormatInt(int64(s), 10)
	case uint32:
		return strconv.FormatInt(int64(s), 10)
	case uint16:
		return strconv.FormatInt(int64(s), 10)
	case uint8:
		return strconv.FormatInt(int64(s), 10)
	case []byte:
		return b2s(s)
	case error:
		return s.Error()
	}
	return ""
}

func prepareRecipients(cc interface{}) (bb string, err error) {
	var b []byte
	ids := make([][]int64, 0)
	switch c := cc.(type) {
	case *Conversation:
		for i := range c.Users {
			ids = append(ids, []int64{c.Users[i].ID})
		}
	case *Item:
		ids = append(ids, []int64{c.User.ID})
	case int64:
		ids = append(ids, []int64{c})
	}
	b, err = json.Marshal(ids)
	bb = b2s(b)
	return
}

// getImageDimensionFromReader return image dimension , types is .jpg and .png
func getImageDimensionFromReader(rdr io.Reader) (int, int, error) {
	image, _, err := image.DecodeConfig(rdr)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}
