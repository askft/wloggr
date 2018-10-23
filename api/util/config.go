package util

// TODO
//  Use JSON to specify configuration variables instead,
//  and then use the `init()` function to set everything up.
//  Also, the config doesn't need to be a struct. Just put
//  the variables as constants in a separate package `config`.

// Config is the configuration for building the API server.
var Config = struct {
	DefaultPort        string
	DBConnectionString string
	IsDevelopment      bool
	DomainName         string
	JwtSigningKey      string
}{
	// Port on which to start the server (unless explicitly specified).
	DefaultPort: "3000",

	// Used to establish a database connection.
	DBConnectionString: "root:password@(localhost:3306)/wloggr",

	// Used when setting up secure middleware (SSL etc).
	IsDevelopment: true,

	// The domain name on which this API server is hosted.
	DomainName: "wloggr.com",

	// Key used to sign newly issued tokens.
	JwtSigningKey: "VerySecretPasswordRightHere",
}
