package v2

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usememos/memos/store"
	apiv2pb "github.com/usememos/memos/proto/gen/api/v2"	
)

func TestAPIV2Service_SetMemoRelations(t *testing.T) {
	ctx := context.Background()
	mockStore := &MockStore{}
	service := &APIV2Service{
		Store: mockStore,
	}

	// Prepare test data
	request := &apiv2pb.SetMemoRelationsRequest{
		Id: "memo-id",
		Relations: []*apiv2pb.MemoRelation{
			{
				RelatedMemoId: "related-memo-id-1",
				Type:          apiv2pb.MemoRelation_REFERENCE,
			},
			{
				RelatedMemoId: "related-memo-id-2",
				Type:          apiv2pb.MemoRelation_REFERENCE,
			},
		},
	}

	// Mock the DeleteMemoRelation function
	mockStore.On("DeleteMemoRelation", ctx, &store.DeleteMemoRelation{
		MemoID: &request.Id,
		Type:   store.MemoRelationReference,
	}).Return(nil)

	// Mock the UpsertMemoRelation function
	mockStore.On("UpsertMemoRelation", ctx, &store.MemoRelation{
		MemoID:        request.Id,
		RelatedMemoID: "related-memo-id-1",
		Type:          store.MemoRelationReference,
	}).Return(nil)

	mockStore.On("UpsertMemoRelation", ctx, &store.MemoRelation{
		MemoID:        request.Id,
		RelatedMemoID: "related-memo-id-2",
		Type:          store.MemoRelationReference,
	}).Return(nil)

	// Call the function under test
	response, err := service.SetMemoRelations(ctx, request)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, response)

	// Verify the function calls
	mockStore.AssertCalled(t, "DeleteMemoRelation", ctx, &store.DeleteMemoRelation{
		MemoID: &request.Id,
		Type:   store.MemoRelationReference,
	})

	mockStore.AssertCalled(t, "UpsertMemoRelation", ctx, &store.MemoRelation{
		MemoID:        request.Id,
		RelatedMemoID: "related-memo-id-1",
		Type:          store.MemoRelationReference,
	})

	mockStore.AssertCalled(t, "UpsertMemoRelation", ctx, &store.MemoRelation{
		MemoID:        request.Id,
		RelatedMemoID: "related-memo-id-2",
		Type:          store.MemoRelationReference,
	})
}