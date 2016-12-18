package cmd

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestStackShow(t *testing.T) {
	cases := []struct {
		req  AppStackShowRequest
		res  string
		want *AppStackShowResponse
	}{
		{
			req: AppStackShowRequest{
				ID: 1,
			},
			res: `{
				"app_stack":{
					"id":1,
					"name":"pretty",
					"inserted_at":1234567890,
					"updated_at":1481537486
				}
			}`,
			want: &AppStackShowResponse{
				AppStack: AppStack{
					ID:         1,
					Name:       "pretty",
					InsertedAt: IntToUnixtime(1234567890),
					UpdatedAt:  IntToUnixtime(1481537486),
				},
			},
		},
		{
			req: AppStackShowRequest{
				ID: 2,
			},
			res: `{"app_stack":{"updated_at":1481953207,"name":"compact_unordered","inserted_at":1234567890,"id":2}}`,
			want: &AppStackShowResponse{
				AppStack: AppStack{
					ID:         2,
					Name:       "compact_unordered",
					InsertedAt: IntToUnixtime(1234567890),
					UpdatedAt:  IntToUnixtime(1481953207),
				},
			},
		},
	}

	mux, mockServerURL := newMockServer()
	client := newTestClient(mockServerURL)

	for _, tc := range cases {
		hundlePath := fmt.Sprintf("/api/app_stacks/%d", tc.req.ID)
		mux.HandleFunc(hundlePath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, tc.res)
		})

		got, err := client.StackShow(context.Background(), tc.req)

		if err != nil {
			t.Fatalf("StackShow was failed: req = %+v, got = %+v, err = %+v", tc.req, got, err)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got=%+v, want=%+v", got, tc.want)
		}
	}
}
