import {
    TableContainer,
    Table,
    TableHead,
    TableRow,
    TableCell,
    TableBody,
    Paper,
    IconButton,
} from '@mui/material';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';

import { useState } from 'react';
import AddCluster from './AddCluster';

export default function Clusters() {

    const [rows, setRows] = useState([])
    const [openAdd, setOpenAdd] = useState(false)
    return (
        <>
        <div style={{
            
            display: 'flex',
            justifyContent: 'space-between',
            padding: '10px 10px',
        }}>
            <h4 style={{margin:'0'}}>Clusters</h4>
            <h4 style={{margin:'0'}}>items:{rows.length}</h4>
            <IconButton onClick={()=>{
                setOpenAdd(true)
            }}>
                <AddCircleOutlineIcon />
            </IconButton>
        </div>
        <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} size="small">
        <TableHead>
          <TableRow sx={{
            backgroundColor:"rgba(100,100,100,0.1)"
          }}>
            <TableCell>Name</TableCell>
            <TableCell>Version</TableCell>
            <TableCell>Status</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow
              key={row.name}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {row.name}
              </TableCell>
              <TableCell>{row.version}</TableCell>
              <TableCell>{row.status}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>

    <AddCluster open={openAdd} setOpen={setOpenAdd}/>
    </>
    )

}