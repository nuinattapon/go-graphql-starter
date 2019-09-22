package resolver

import (
	"log"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	gcontext "github.com/nuinattapon/go-graphql-starter/context"
	"github.com/nuinattapon/go-graphql-starter/schema"
	"github.com/nuinattapon/go-graphql-starter/service"
	"golang.org/x/net/context"
)

var (
	rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &Resolver{})
	ctx        context.Context
)

func init() {
	config := gcontext.LoadConfig("../")
	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	log := service.NewLogger(config)
	roleService := service.NewRoleService(db, log)
	userService := service.NewUserService(db, roleService, log)
	ctx = context.WithValue(context.Background(), "userService", userService)
}

func TestBasic(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
				{
					user(email:"test@1.com") {
						id
						email
						password
					}
				}
			`,
			ExpectedResult: `
				{
					"user": {
					  "id": "1",
					  "email": "test@1.com",
					  "password": "********"
					}
				}
			`,
		},
	})
}
