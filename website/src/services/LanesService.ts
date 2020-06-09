import { Cluster } from "../entities/home/Cluster"

export class GetClustersByLaneResponse {
  ok: boolean
  data: Cluster[] | any
  constructor(ok: boolean, data: Cluster[] | any) {
    this.ok = ok
    this.data = data
  }
}

export class LanesService {
  static async getClustersByLane(lane: string, totalClusters: number, attributes: Set<string>, distanceMethod: string) {
    const queryParams = new URLSearchParams()
    queryParams.append('totalClusters', totalClusters.toString())
    queryParams.append('distanceMethod', distanceMethod)
    queryParams.append('attributes', Array.from(attributes).toString())
    const url = `http://localhost:8000/api/lanes/${lane}?${queryParams.toString()}`
    const response = await fetch(url)
    const data = await response.json()
    return {
      data,
      ok: response.ok
    } as GetClustersByLaneResponse
  }
}
