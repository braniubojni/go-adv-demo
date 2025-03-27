package main

import (
	"context"
	"fmt"
)

func mainFake() {
	type key int
	const EmailKey key = 0
	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, EmailKey, "asd@zxc.ru")

	if userEmail, ok := ctxWithValue.Value(EmailKey).(string); ok {
		fmt.Println(userEmail)
	} else {
		fmt.Println("No value")
	}
}
