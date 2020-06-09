import React, { useReducer } from 'react'
import { Typography, Grid, Button } from '@material-ui/core'
import TextFieldWrapper from '../common/TextFieldWrapper'
import { homeReducer } from './HomeReducer'
import { HomeState } from './HomeState'
import { ClusterInputs } from '../../entities/home/ClusterInputs'
import SelectWrapper from '../common/SelectWrapper'
import { laneSelectItems, attributeCheckboxGroupItems, distanceMethodSelectItems } from './Constants'
import CheckboxGroupWrapper from '../common/CheckboxGroupWrapper'
import ClustersTable from './ClustersTable'

const HomePage = () => {
  const [state, dispatch] = useReducer(homeReducer, HomeState.initial())
  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Typography variant='h1'>LoL Champion Groups</Typography>
      </Grid>

      <Grid item xs={4} container spacing={2}>
        <Grid item xs={12}>
          <Button
            color='primary'
            variant='contained'
            onClick={HomeState.handleGenerateClusters(state.clusterInputs, dispatch)}
          >
            Generate Clusters
            </Button>
        </Grid>

        <Grid item xs={12}>
          <TextFieldWrapper
            label='Total Clusters'
            type='text'
            name='totalClusters'
            helperText='Total number of clusters to generate'
            required={true}
            updateState={(prop, value) => HomeState.setClusterInputsProp(prop, dispatch, value == null ? value : parseInt(value))}
            isValueOk={ClusterInputs.isTotalClustersOk}
          />
        </Grid>

        <Grid item xs={12}>
          <SelectWrapper
            label='Distance Method'
            name='distanceMethod'
            required={true}
            updateState={(prop, value) => HomeState.setClusterInputsProp(prop, dispatch, value)}
            helperText='Which method to use for calculating ditance between points'
            items={distanceMethodSelectItems}
          />
        </Grid>

        <Grid item xs={12}>
          <SelectWrapper
            label='Lane'
            name='lane'
            required={true}
            updateState={(prop, value) => HomeState.setClusterInputsProp(prop, dispatch, value)}
            helperText='Main lane/role of champions to analyze'
            items={laneSelectItems}
          />
        </Grid>

        <Grid item xs={12}>
          <CheckboxGroupWrapper
            label='Attributes'
            name='attributes'
            required={true}
            updateState={(prop, value) => HomeState.setClusterInputsProp(prop, dispatch, value)}
            helperText='Attributes to take into consideration when clustering champions'
            items={attributeCheckboxGroupItems}
          />
        </Grid>
      </Grid>

      <Grid item xs={8}>
        {state.clusters.length > 0 &&
          <ClustersTable
            clusters={state.clusters}
            clusterAttributes={state.clusterAttributes}
          />
        }
      </Grid>
    </Grid>
  )
}

export default HomePage
