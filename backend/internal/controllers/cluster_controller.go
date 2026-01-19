package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dana-team/axiom/backend/internal/types"
	"github.com/dana-team/axiom/backend/internal/utils"
	v1alpha "github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClusterController handles HTTP requests related to cluster operations
// using MongoDB as the data store
type ClusterController struct {
	mongoClient *utils.MongoClient
}

// NewClusterController creates a new instance of ClusterController with the provided
// MongoDB client for handling cluster-related operations
func NewClusterController(mongoClient *utils.MongoClient) *ClusterController {
	return &ClusterController{
		mongoClient: mongoClient,
	}
}

// GetClusters handles HTTP GET requests to retrieve cluster information
// with support for filtering, sorting, and pagination
// @Summary      Get Clusters
// @Description  Retrieve a list of clusters with optional filtering and pagination
// @Tags         clusters
// @Accept       json
// @Produce      json
// @Param        clusterId    query     string  false  "Cluster ID"
// @Param        kubeVersion  query     string  false  "Kubernetes Version"
// @Param        name         query     string  false  "Cluster Name"
// @Param        limit        query     int     false  "Limit"
// @Param        sortBy       query     string  false  "Sort By"
// @Param        sortOrder    query     string  false  "Sort Order"
// @Success      200          {object}  types.APIResponse
// @Failure      400          {object}  types.APIResponse
// @Failure      500          {object}  types.APIResponse
// @Router       /clusters [get]
func (cc *ClusterController) GetClusters(c *gin.Context) {
	params := types.QueryParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Success: false,
			Message: "Invalid query parameters" + err.Error(),
		})
		return
	}

	collection := cc.mongoClient.GetCollection(utils.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if params.ClusterID != "" {
		filter["clusterID"] = params.ClusterID
	}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex":   params.Name,
			"$options": "i",
		}
	}

	if params.KubeVersion != "" {
		filter["kubernetesVersion"] = bson.M{
			"$regex":   fmt.Sprintf("^v?%s(\\.|$)", params.KubeVersion),
			"$options": "i",
		}
	}

	sortOrder := 1
	if params.SortOrder == "desc" {
		sortOrder = -1
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(params.Limit))
	findOptions.SetSort(bson.D{{Key: params.SortBy, Value: sortOrder}})

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Success: false,
			Message: "Failed to query clusters:" + err.Error(),
		})
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	clusters := []v1alpha.ClusterInfoStatus{}
	if err := cursor.All(ctx, &clusters); err != nil {
		c.JSON(http.StatusInternalServerError, types.APIResponse{
			Success: false,
			Message: "Failed to decode clusters:" + err.Error(),
		})
		return
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		count = 0
	}

	c.JSON(http.StatusOK, types.APIResponse{
		Success: true,
		Data:    clusters,
		Count:   count,
	})
}
