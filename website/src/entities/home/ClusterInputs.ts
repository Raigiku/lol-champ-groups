export class ClusterInputs {
  totalClusters?: number
  lane?: string
  attributes?: Set<string>
  constructor(totalClusters?: number, lane?: string, attributes?: Set<string>) {
    this.totalClusters = totalClusters
    this.lane = lane
    this.attributes = attributes
  }

  static initial() {
    return new ClusterInputs(undefined, undefined, undefined)
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
  }
}