import { DataPoint } from "./DataPoint";

export class Cluster {
  centroid: DataPoint
  dataPoints: DataPoint[]
  constructor(centroid: DataPoint, dataPoints: DataPoint[]) {
    this.centroid = centroid
    this.dataPoints = dataPoints
  }

  static initialize(centroid: DataPoint, dataPoints: DataPoint[]) {
    return new Cluster(centroid, dataPoints)
  }
}