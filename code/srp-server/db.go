package main

// AuthDatabaseValue stores SRP credentials retrieved from the authentication database.
type AuthDatabaseValue struct {
	salt     string
	verifier string
}

// authDatabase stores the username password mapping. This would normally be stored in a database, such as MySQL.
var authDatabase = map[string]AuthDatabaseValue{
	// password: "testPassword"
	"test@example.com": {
		salt:     "l5tM9mLMqAd4q57Z89aBopHUyAfFTFaOjt7AWcKo/5I=",
		verifier: "78ion/89QmSELqrTPpe/57c7uojZOUGgKs7RRXlaJSvzABBmGovgqOPHo4PZbEGU12eyTeTIyTw8xEiEpDiB1/wa+oi+rX6ivdZjCc51s2hf6xBldk8dhDCValSGh8XsGmzQwOXpMkjzIeSHuoPjbR6Yriq01h45pyga0AyggmnlGTo3L1keqGJYtOJJhsf1xdule/a/y5krhqOp8X3beFOmCZx4MLPwyhL0sVhLgqDMIbSF/DrVC7WMvfXV2KTo6tCdfjlsnW1E2iNwdVf3NsxyuIQpEPGXGJylIh22vR2MZS02kFoGyzGl0t54p8YS7LzSW5eksDgFEQOMkb16VvLf/3rMbnN5AYTtIE52DmxUG77WMyo4StX5RMIkF/aZWCkG7SFd+5Lvo0agVRlkE3HcVLPNlX8XIv5E5EMnKWkGKvIKz0Qk7suzcbEVv3n8HdAOzcw0M5hkK9FgYXO9Ktu95mLh7peaq3S4nKXcPgtKAwaECQLPlbvoYXBU8O83SpBG5x5It4PlclHqrU/giQgids4qiJXByG8idsFdpZY5UmLZKGSGs6vRyk6rKTyF5sSOcHq28C8SqhJAJ9EEEbIvk+rZIWdTNjXWCKL6JRbHxrJRubpH/Otx8OsgOf702c72VIf0TJfibRK+iX5XEMBcGMgtt8Cn1sym7Yv0mDM=",
	},
}
