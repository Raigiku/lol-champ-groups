import { ClusterInputs } from "../../entities/home/ClusterInputs"
import { HomeAction, HomeActionTypes } from "./HomeReducer"
import { LanesService } from "../../services/LanesService"
import { Cluster } from "../../entities/home/Cluster"

export class HomeState {
  clusterInputs: ClusterInputs
  clusters: Cluster[]
  clusterAttributes: string[]
  constructor(clusterInputs: ClusterInputs, clusters: Cluster[], clusterAttributes: string[]) {
    this.clusterInputs = clusterInputs
    this.clusters = clusters
    this.clusterAttributes = clusterAttributes
  }

  static initial() {
    return new HomeState(ClusterInputs.initial(), [], [])
  }

  static setClusterInputsProp<T>(prop: string, dispatch: React.Dispatch<HomeAction>, value?: T) {
    dispatch({
      type: HomeActionTypes.SetClusterInputsProp,
      payload: { prop, value }
    })
  }

  static handleGenerateClusters(clusterInputs: ClusterInputs, dispatch: React.Dispatch<HomeAction>) {
    return async () => {
      if (ClusterInputs.isOk(clusterInputs)) {
        const response = await LanesService.getClustersByLane(
          clusterInputs.lane!,
          clusterInputs.totalClusters!,
          clusterInputs.attributes!,
          clusterInputs.distanceMethod!
        )
        if (response.ok) {
          dispatch({ type: HomeActionTypes.SetProp, payload: { prop: 'clusters', value: response.data } })
          dispatch({ type: HomeActionTypes.SetProp, payload: { prop: 'clusterAttributes', value: Array.from(clusterInputs.attributes!) }})
        } else {
          alert(response.data)
        }
      }
    }
  }
}