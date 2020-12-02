package schema

// This may not be the best way to maintain this, but it's quick

// _document represents the schema for the project. It is written to be processed
// by the Go templating engine to provide required parameters to be injected into
// the schema. In a production system the schema should be embedded using a tool
// like pakcr from the Buffalo project (https://github.com/gobuffalo/packr).
var _document = `
type City {
	id: ID!
	name: String! @search(by: [exact])
	lat: Float!
	lng: Float!
}
`
