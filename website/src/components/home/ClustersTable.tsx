import React from 'react'
import { TableContainer, Paper, Table, TableHead, TableRow, TableBody, TableCell, withStyles, Theme, createStyles } from '@material-ui/core'
import { Cluster } from '../../entities/home/Cluster'
import { v4 as uuidv4 } from 'uuid'

type Props = {
  clusters: Cluster[]
  clusterAttributes: string[]
}

const StyledTableRow = withStyles((theme: Theme) =>
  createStyles({
    root: {
      '&:nth-of-type(odd)': {
        backgroundColor: theme.palette.action.hover,
      },
    },
  }),
)(TableRow);

const ClustersTable = ({
  clusterAttributes,
  clusters
}: Props) => {

  return (
    <>
      {clusters.map((cluster, i) => (
        <React.Fragment key={uuidv4()}>
          <TableContainer component={Paper}>
            <Table>
              <TableHead>
                <StyledTableRow>
                  <TableCell>Cluster {i + 1}</TableCell>
                  <TableCell></TableCell>
                </StyledTableRow>
                <TableRow>
                  <TableCell>name</TableCell>
                  {clusterAttributes.map(attr => (
                    <TableCell key={uuidv4()}>{attr}</TableCell>
                  ))}
                </TableRow>
              </TableHead>

              <TableBody>
                <StyledTableRow>
                  <TableCell>Centroid</TableCell >
                  {cluster.centroid.positions.map(position => (
                    <TableCell key={uuidv4()}>{position.toFixed(2)}</TableCell>
                  ))}
                </StyledTableRow>

                {cluster.dataPoints.map(dataPoint => (
                  <TableRow key={uuidv4()}>
                    <TableCell>{dataPoint.championName}</TableCell>
                    {dataPoint.positions.map(position => (
                      <TableCell key={uuidv4()}>{position.toFixed(2)}</TableCell>
                    ))}
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          <br />
          <br />
        </React.Fragment>
      ))}
    </>
  )
}

export default ClustersTable
