package cipher

import "io"

func (r *Rot128Reader) DecodeToString() string {
	decodedValue := ""

	for {
		arr := make([]byte, 256)
		// n is number of bytes read
		n, err := r.Read(arr)

		if err == io.EOF {
			break
		}
		decodedValue += string(arr[:n])
	}
	return decodedValue
}

func (r *Rot128Reader) DecodeToBytes() []byte {
	decodedString := r.DecodeToString()

	return []byte(decodedString)
}