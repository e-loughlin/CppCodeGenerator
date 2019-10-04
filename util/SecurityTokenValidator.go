package util

/// securityToken ... Ensures that the server will not accept unauthorized POST requests,
/// and the executable will not run without it as an argument.
const securityToken = "EC4307A23E260072141C85CF4B0B1911726EEF65205877D0D03EF933AAA4D453"

func IsValidSecurityToken(token string) bool {
	return (token == securityToken)
}
