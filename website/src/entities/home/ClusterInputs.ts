export class ClusterInputs {
  totalClusters?: number
  lane?: string
  distanceMethod?: string
  attributes?: Set<string>
  constructor(totalClusters?: number, lane?: string, attributes?: Set<string>, distanceMethod?: string) {
    this.totalClusters = totalClusters
    this.lane = lane
    this.attributes = attributes
    this.distanceMethod = distanceMethod
  }

  static initial() {
    return new ClusterInputs()
  }

  static isTotalClustersOk(value: string) {
    const isPositiveInteger = /^\+?[1-9][\d]*$/.test(value)
    if (!isPositiveInteger) {
      return 'Has to be a positive integer'
    }
    return undefined
  }

  static isOk(self: ClusterInputs) {
    return self.attributes != null
      && self.lane != null
      && self.totalClusters != null
      && self.distanceMethod != null
  }
}