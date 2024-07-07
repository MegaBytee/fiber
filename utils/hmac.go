package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//Generating a SHA256 HMAC Hash
/*Symmetric Signatures / Message Authentication - HMAC-SHA512/256

When two parties share a secret key, they can use message authentication to
make sure that a piece of data hasn't been altered. You can think of it as a
"symmetric signature" - it proves both that the data is unchanged and that
someone who knows the shared secret key generated it. Anyone who does not know
the secret key can neither validate the data nor make valid alterations.

This comes up most often in the context of web stuff, such as:

1. Authenticating requests to your API. The most widely known example is
   probably the Amazon AWS API, which requires you to sign requests with
   HMAC-SHA256. In this type of use, the "secret key" is a token that the API
   provider issues to authorized API users.

2. Validating authenticated tokens (cookies, JWTs, etc) that are issued by a
   service but are stored by a user. In this case, the service wants to ensure
   that a user doesn't modify the data contained in the token.

As with encryption, you should always use a 256-bit random key to
authenticate messages. */

// sign data with secrect key
func Hmac256Signature(secret, data string) string {

	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(data))
	return hex.EncodeToString(hmac256.Sum(nil))
}
