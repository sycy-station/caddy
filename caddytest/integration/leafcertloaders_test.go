package integration

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddytest"
)

func TestLeafCertLoaders(t *testing.T) {
	tester := caddytest.StartHarness(t)
	tester.LoadConfig(`
	{
		"admin": {
			"listen": "{$TESTING_ADMIN_API}"
		},
		"apps": {
			"http": {
				"http_port": 9080,
       			"https_port": 9443,
				"grace_period": 1,
				"servers": {
					"srv0": {
						"listen": [
							":9443"
						],
						"routes": [
							{
								"match": [
									{
										"host": [
											"localhost"
										]
									}
								],
								"terminal": true
							}
						],
						"tls_connection_policies": [
							{
								"client_authentication": {
									"verifiers": [
										{
											"verifier": "leaf",
											"leaf_certs_loaders": [
												{
													"loader": "file",
													"files": ["../leafcert.pem"]
												}, 
												{
													"loader": "folder", 
													"folders": ["../"]
												},
												{
													"loader": "storage"
												},
												{
													"loader": "pem"
												}
											]
										}
									]
								}
							}
						]
					}
				}
			}
		}
	}`, "json")
}
