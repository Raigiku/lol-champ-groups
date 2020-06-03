export class DataPoint {
  championName: string
  positions: number[]
  constructor(championName: string, positions: number[]) {
    this.championName = championName
    this.positions = positions
  }
}
