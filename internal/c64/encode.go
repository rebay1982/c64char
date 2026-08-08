package c64

func Encode(buff []uint8, w, h int) (string, error) {
	
	err := validateSize(buff, w, h)
	if err != nil {
		return "", err
	}




	return "", nil
}

