package jwt

import (
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("parin")
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(input string, k int) string {
	t1 := time.Now().Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["sub"] = input
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["exp"] = t1 + 500
	claims["iat"] = t1
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return ""
	}
	s1 := strings.Split(tokenString, ".")
	s3 := strings.Split(s1[1], "")
	var s4 []string
	// k := 9
	for i := 0; i <= len(s3); i++ {
		j := (len(s3) / k)
		if i%k == 0 && i != 0 {
			s4uum := ``
			for l := 0; l < k; l++ {
				// s4uum = s4uum + s3[i-7] + s3[i-6] + s3[i-5] + s3[i-4] + s3[i-2] + s3[i-1]
				// s4uum = s4uum + s3[i-1-(k-l-1)]
				s4uum = s4uum + s3[i-1-(l)]
			}
			s4 = append(s4, s4uum)
		} else if i >= j*k {
			s4 = append(s4, s3[i-1])
		}
	}
	s5 := ``
	for i := 0; i < len(s4); i++ {
		s5 = s5 + s4[i]
	}
	output_encode := s5 + `.` + s1[2]
	return output_encode
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string, k_d int) string {
	head := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`
	input_decode := tokenStr
	s1_d := strings.Split(input_decode, ".")
	s3_d := strings.Split(s1_d[0], "")
	var s4_d []string
	// k_d := 9
	for i := 0; i <= len(s3_d); i++ {
		j := (len(s3_d) / k_d)
		if i%k_d == 0 && i != 0 {
			s4_duum := ``
			for l := 0; l < k_d; l++ {
				// s4_duum = s4uum + s3_d[i-7] + s3_d[i-6] + s3_d[i-5] + s3_d[i-4] + s3_d[i-2] + s3_d[i-1]
				// s4_duum = s4uum + s3_d[i-1-(k_d-l-1)]
				s4_duum = s4_duum + s3_d[i-1-l]
			}
			s4_d = append(s4_d, s4_duum)
		} else if i >= j*k_d {
			s4_d = append(s4_d, s3_d[i-1])
			// fmt.Println("-->", i)
		}
	}

	s5_d := ``
	for i := 0; i < len(s4_d); i++ {
		s5_d = s5_d + s4_d[i]
	}

	output_decode := head + `.` + s5_d + `.` + s1_d[1]

	token, err := jwt.Parse(output_decode, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		// return ""
	}
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		output := claims["sub"].(string)
		return output
	} else {
		return ""
	}
}
