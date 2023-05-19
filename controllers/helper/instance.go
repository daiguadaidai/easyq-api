package helper

import (
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/models/view_models"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/response"
)

// 实例集群信息转化成api返回信息
// 如果只有slave没有master, 则会不显示
func InstanceClustersToMasterSlavesResps(instances []*view_models.InstanceCluster) []*response.InstanceClusterResponse {
	if len(instances) == 0 {
		return nil
	}

	// 将实例信息转化成 response 结构体
	instanceResps := InstanceClustersToResps(instances)
	// 获取从实例Map, key: host:port, value: slaves
	masterHostrPortSlavesMap := GetMasterHostrPortSlavesMap(instanceResps)

	// 获取所有master
	masters := GetMasters(instances)

	// 为master添加slave
	masterResps := make([]*response.InstanceClusterResponse, 0, len(masters))
	for _, master := range masters {
		var masterResp response.InstanceClusterResponse
		utils.CopyStruct(master, &masterResp)

		setedSlaveMap := make(map[string]struct{}) // 用户保存已经添加过的slave hostport
		setRespSlave(&masterResp, masterHostrPortSlavesMap, setedSlaveMap)

		masterResps = append(masterResps, &masterResp)
	}

	return masterResps
}

// 获取slave对应的 master map 实例集群map, key: host:port, value: slaves
// key是空字符串, 是主实例
func GetMasterHostrPortSlavesMap(instances []*response.InstanceClusterResponse) map[string][]*response.InstanceClusterResponse {
	masterHostrPortSlavesMap := make(map[string][]*response.InstanceClusterResponse)

	for _, instance := range instances {
		if instance.Role.String == models.InstanceRoleMaster {
			continue
		}
		// 如果 hostPort 等于空字符串, 代表是主实例
		var hostPort string
		if !instance.MasterHost.IsEmpty() { // 该实例是slave
			hostPort = utils.AddrI64(instance.MasterHost.String, instance.MasterPort.Int64)
		}

		tmpMap, ok := masterHostrPortSlavesMap[hostPort]
		if !ok {
			tmpMap = make([]*response.InstanceClusterResponse, 0, 5)
			masterHostrPortSlavesMap[hostPort] = tmpMap
		}

		masterHostrPortSlavesMap[hostPort] = append(masterHostrPortSlavesMap[hostPort], instance)
	}

	return masterHostrPortSlavesMap
}

// 获取所有master
func GetMasters(instances []*view_models.InstanceCluster) []*view_models.InstanceCluster {
	masters := make([]*view_models.InstanceCluster, 0, 4)
	for _, instance := range instances {
		if instance.Role.String == models.InstanceRoleMaster {
			masters = append(masters, instance)
		}
	}

	return masters
}

// 用户保存已经添加过的slave hostport
// 设置slave
func setRespSlave(
	master *response.InstanceClusterResponse,
	masterHostrPortSlavesMap map[string][]*response.InstanceClusterResponse,
	setedSlaveMap map[string]struct{},
) {
	// master 没有对应的ip
	if master.MachineHost.IsEmpty() && master.VpcgwRip.IsEmpty() {
		return
	}

	machineMaster := utils.AddrI64(master.MachineHost.String, master.Port.Int64)
	// 机器ip找slave
	var slaves []*response.InstanceClusterResponse
	var ok bool
	if !master.MachineHost.IsEmpty() {
		if slaves, ok = masterHostrPortSlavesMap[machineMaster]; ok {
			for _, slave := range slaves {
				if slave.MachineHost.IsEmpty() { // 如果该slave没有机器ip就不需要再查找他的slave了
					continue
				}
				slaveAddr := utils.AddrI64(slave.MachineHost.String, slave.Port.Int64)
				// 该ip已经设置过则就不再寻找他的slave了
				if _, ok := setedSlaveMap[slaveAddr]; ok {
					continue
				}

				// 标记该slave已经处理过
				setedSlaveMap[slaveAddr] = struct{}{}

				// 查找该slave的 slaves
				setRespSlave(slave, masterHostrPortSlavesMap, setedSlaveMap)
			}
		}
	}
	if len(slaves) > 0 {
		master.Slaves = slaves
		return
	}

	ripMaster := utils.AddrI64(master.VpcgwRip.String, master.Port.Int64)
	// 在 vpc rip port找不到slave 使用 rip找
	if !master.VpcgwRip.IsEmpty() {
		if slaves, ok = masterHostrPortSlavesMap[ripMaster]; ok {
			for _, slave := range slaves {
				if slave.VpcgwRip.IsEmpty() { // 如果该slave没有 vpc rip就不需要再查找他的slave了
					continue
				}
				slaveAddr := utils.AddrI64(slave.VpcgwRip.String, slave.Port.Int64)
				// 该ip已经设置过则就不再寻找他的slave了
				if _, ok := setedSlaveMap[slaveAddr]; ok {
					continue
				}

				// 标记该slave已经处理过
				setedSlaveMap[slaveAddr] = struct{}{}

				// 查找该slave的 slaves
				setRespSlave(slave, masterHostrPortSlavesMap, setedSlaveMap)
			}
		}
	}
	if len(slaves) > 0 {
		master.Slaves = slaves
	}
}

func InstanceClustersToResps(instances []*view_models.InstanceCluster) []*response.InstanceClusterResponse {
	resps := make([]*response.InstanceClusterResponse, 0, len(instances))
	for _, instance := range instances {
		var resp response.InstanceClusterResponse
		utils.CopyStruct(instance, &resp)
		resps = append(resps, &resp)
	}

	return resps
}
