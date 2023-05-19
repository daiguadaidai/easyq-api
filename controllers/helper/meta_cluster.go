package helper

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/types"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/response"
	"strings"
)

func MetaClusterToClusterResponse(metaCluster *models.MetaCluster) *response.ClusterResponse {
	var resp response.ClusterResponse
	utils.CopyStruct(metaCluster, &resp)

	// 按逗号隔开, 拆分域名
	domainNames := make([]types.NullString, 0, 2)
	vipPorts := make([]types.NullString, 0, 2)
	vpcgwVipPorts := make([]types.NullString, 0, 2)
	if !metaCluster.DomainName.IsEmpty() {
		for _, domanName := range strings.Split(metaCluster.DomainName.String, ",") {
			domainNames = append(domainNames, types.NewNullString(strings.TrimSpace(domanName), true))
		}
	}
	if !metaCluster.VipPort.IsEmpty() {
		for _, vipPort := range strings.Split(metaCluster.VipPort.String, ",") {
			vipPorts = append(vipPorts, types.NewNullString(strings.TrimSpace(vipPort), true))
		}
	}
	if !metaCluster.VpcgwVipPort.IsEmpty() {
		for _, vpcgwVipPort := range strings.Split(metaCluster.VpcgwVipPort.String, ",") {
			vpcgwVipPorts = append(vpcgwVipPorts, types.NewNullString(strings.TrimSpace(vpcgwVipPort), true))
		}
	}
	resp.DomainNames = domainNames
	resp.VipPorts = vipPorts
	resp.VpcgwVipPorts = vpcgwVipPorts

	return &resp
}

func MetaClusterToClusterResponses(metaClusters []*models.MetaCluster) []*response.ClusterResponse {
	clusterResponses := make([]*response.ClusterResponse, 0, len(metaClusters))

	for _, metaCluster := range metaClusters {
		clusterResponse := MetaClusterToClusterResponse(metaCluster)
		clusterResponses = append(clusterResponses, clusterResponse)
	}

	return clusterResponses
}

func MetaClusterToNameResponses(metaClusters []*models.MetaCluster) []*response.ClusterNameResponse {
	clusterNameResponses := make([]*response.ClusterNameResponse, 0, len(metaClusters))

	for _, metaCluster := range metaClusters {
		var nameResp response.ClusterNameResponse
		utils.CopyStruct(metaCluster, &nameResp)
		clusterNameResponses = append(clusterNameResponses, &nameResp)
	}

	return clusterNameResponses
}

func MateClustersToIdMap(clusters []*models.MetaCluster) map[int64]*models.MetaCluster {
	clusterIdMap := make(map[int64]*models.MetaCluster)
	for _, cluster := range clusters {
		clusterIdMap[cluster.ID.Int64] = cluster
	}

	return clusterIdMap
}

// 获取master host和port
func GetMasterHostAndPortWithMetaCluster(cluster *models.MetaCluster) (string, int64, error) {
	var host string
	var port int64
	switch cluster.Category.Int64 {
	case models.ClusterCategoryTceSelfMySQL, // TCE自建MySQL
		models.ClusterCategoryPublicSelfMySQL, // 公有云自建MySQL
		models.ClusterCategoryFuZhouMySQL:     // 福州MySQL(自建)
		var ok bool
		host, port, ok = utils.AddrToHostPort(cluster.VipPort.String)
		if !ok {
			return "", 0, fmt.Errorf("通过自建MySQL VIP地址获取 host port 失败. VIP地址: %v", cluster.VpcgwVipPort.String)
		}
	case models.ClusterCategoryTceTDSQL: // TEC-TDSQL
		// 不支持 TCE tdsql 分布式
		if cluster.ShardType.String == models.ShardTypeTDSQL {
			// 不支持tdsql分布式
			return "", 0, fmt.Errorf("暂时不支持 分布式TDSQL")
		} else {
			// 通过proxy获取数据库
			var ok bool
			host, port, ok = utils.AddrToHostPort(cluster.VpcgwVipPort.String)
			if !ok {
				return "", 0, fmt.Errorf("通过单实例TDSQL proxy地址获取 host port 失败. proxy地址: %v", cluster.VpcgwVipPort.String)
			}
		}
	case models.ClusterCategoryPublicTDSQLC, // 公有云TDSQL-C
		models.ClusterCategoryPublicMySQL: // 公有云MySQL
		// 通过proxy获取数据库
		var ok bool
		host, port, ok = utils.AddrToHostPort(cluster.VpcgwVipPort.String)
		if !ok {
			return "", 0, fmt.Errorf("通过公有云MySQl proxy地址获取 host port 失败. proxy地址: %v", cluster.VpcgwVipPort.String)
		}
	case models.ClusterCategoryTIDB: // TIDB
		return "", 0, fmt.Errorf("暂时不支持TIDB审核")
	case models.ClusterCategoryUnknow: // 未知
		return "", 0, fmt.Errorf("遇到未知集群类别. 集群id: %v, 集群名: %v, 类别: %v", cluster.ID.Int64, cluster.Name.String, cluster.Category.Int64)
	default:
		return "", 0, fmt.Errorf("遇到无效集群类别. 集群id: %v, 集群名: %v, 类别: %v", cluster.ID.Int64, cluster.Name.String, cluster.Category.Int64)
	}

	return host, port, nil
}
