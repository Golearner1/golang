package secure

import "golang.org/x/crypto/bcrypt"

func Lockpassword(PASSWORD string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(PASSWORD), 14)
	return string(bytes), err
}
func Checkpassword(passwrod,lockedpassword string)bool{
err:=bcrypt.CompareHashAndPassword([]byte(lockedpassword),[]byte(passwrod))
return err==nil
}