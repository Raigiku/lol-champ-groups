import { HomeState } from "./HomeState"

export enum HomeActionTypes {
  SetClusterInputsProp,
  SetProp
}

export class HomeAction {
  type: HomeActionTypes
  payload?: any
  constructor(type: HomeActionTypes, payload?: any) {
    this.type = type
    this.payload = payload
  }
}

export const homeReducer = (state: HomeState, action: HomeAction) => {
  switch (action.type) {
    case HomeActionTypes.SetClusterInputsProp: {
      const { prop, value } = action.payload
      return {
        ...state,
        clusterInputs: {
          ...state.clusterInputs,
          [prop]: value
        }
      } as HomeState
    }

     case HomeActionTypes.SetProp: {
       const { prop, value } = action.payload
       return {
        ...state,
        [prop]: value
       } as HomeState
     }

    default:
      throw new Error('No case for undefined action type')
  }
}