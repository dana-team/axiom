package e2e_tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/dana-team/axiom/backend/internal/controllers"
	"github.com/dana-team/axiom/backend/internal/types"
	"github.com/dana-team/axiom/backend/internal/utils"
	"github.com/dana-team/axiom/backend/internal/utils/testutils"
	v1alpha "github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ = Describe("ClusterController", func() {
	var (
		router      *gin.Engine
		mongoClient *utils.MongoClient
		collection  *mongo.Collection
		recorder    *httptest.ResponseRecorder
		ctx         context.Context
	)

	BeforeEach(func() {
		ctx = context.TODO()
		// Set Gin to test mode
		gin.SetMode(gin.TestMode)

		// Create a new Gin engine
		router = gin.New()

		// Create test MongoDB client
		client, err := testutils.ConnectTestMongo()
		Expect(err).NotTo(HaveOccurred())

		mongoClient = &utils.MongoClient{
			Client:   client,
			Database: client.Database(testutils.TestDbName),
		}

		collection = mongoClient.GetCollection(testutils.TestDbCollection)

		// Create a new recorder
		recorder = httptest.NewRecorder()

		// Initialize controller and routes
		clusterController := controllers.NewClusterController(mongoClient)
		router.GET("/v1/clusters", clusterController.GetClusters)
	})

	AfterEach(func() {
		// Clean up the test database
		Expect(collection.Drop(ctx)).To(Succeed())
		Expect(mongoClient.Disconnect()).To(Succeed())
	})

	Context("GetClusters", func() {
		BeforeEach(func() {
			// Insert test data
			_, err := collection.InsertMany(ctx, testutils.TestClusters)
			Expect(err).NotTo(HaveOccurred())
		})

		When("no query parameters are provided", func() {
			It("should return all clusters", func() {
				req := httptest.NewRequest(http.MethodGet, "/v1/clusters", nil)
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))

				var response types.APIResponse
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				Expect(err).NotTo(HaveOccurred())
				Expect(response.Success).To(BeTrue())
				Expect(response.Count).To(Equal(int64(2)))

				clusters, err := decodeData[v1alpha.ClusterInfoStatus](response.Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(clusters).To(HaveLen(2))
			})
		})

		When("filtering by cluster ID", func() {
			It("should return filtered clusters", func() {
				req := httptest.NewRequest(http.MethodGet, "/v1/clusters?clusterId=cluster-1", nil)
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))

				var response types.APIResponse
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				Expect(err).NotTo(HaveOccurred())
				Expect(response.Success).To(BeTrue())
				Expect(response.Count).To(Equal(int64(1)))

				clusters, err := decodeData[v1alpha.ClusterInfoStatus](response.Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(clusters).To(HaveLen(1))
				Expect(clusters[0].ClusterID).To(Equal(testutils.TestCluster1ID))
			})
		})

		When("filtering by Kubernetes version", func() {
			It("should return clusters with matching version", func() {
				req := httptest.NewRequest(http.MethodGet, "/v1/clusters?kubeVersion=1.24", nil)
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))

				var response types.APIResponse
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				Expect(err).NotTo(HaveOccurred())
				Expect(response.Success).To(BeTrue())
				Expect(response.Count).To(Equal(int64(1)))

				clusters, err := decodeData[v1alpha.ClusterInfoStatus](response.Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(clusters).To(HaveLen(1))
				Expect(clusters[0].KubernetesVersion).To(Equal(testutils.TestCluster1Version))
			})
		})

		When("limiting results", func() {
			It("should respect the limit parameter", func() {
				req := httptest.NewRequest(http.MethodGet, "/v1/clusters?limit=1", nil)
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))

				var response types.APIResponse
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				Expect(err).NotTo(HaveOccurred())
				Expect(response.Success).To(BeTrue())

				clusters, err := decodeData[v1alpha.ClusterInfoStatus](response.Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(clusters).To(HaveLen(1))
			})
		})
	})
})

func decodeData[T any](data interface{}) ([]T, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var result []T
	err = json.Unmarshal(dataBytes, &result)
	return result, err
}
