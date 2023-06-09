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
import MoreVertBtn from '../MoreVertBtn';
import clusterHttp from '../../http/cluster'
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import { BookmarkAddedOutlined, DeleteOutlineOutlined } from '@mui/icons-material';
import { useState,useEffect,useContext } from 'react';
import AddCluster from './AddCluster';
import { setCurrentCluster,setCurrentClusterName,currentCluster } from '../../store/cluster';
import { AppContext } from '../../App';
import listHttp from '../../http/list';

function chooseClusterHandler(row,setCurrentChoose,setCluster,setNamespace,setNamespaces) {
  return {
    title: 'choose',
    icon: <BookmarkAddedOutlined />,
    handle: () => {
        //获取namespace
       
        setCurrentCluster(row.id);
        return new Promise((resolve,reject) => {
          listHttp.listNamespace().then(res => {
              setNamespaces(res);
              setCurrentClusterName(row.name);
              setCurrentChoose(row.id);
              setNamespace(res[0]);
              setCluster(row.id);
              resolve({
                msg: 'change success',
                type: 'success'
              });
          }).catch(err => {
            reject(err);
          })
        })
       
    }
  }
}

function deleteClusterHandler(row,setRows) {
  return {
    title: 'Delete',
    icon: <DeleteOutlineOutlined />,
    handle: () => {
        return new Promise((resolve,reject) => {
            clusterHttp.deleteCluster(row.id).then(res => {
                getClusterList(setRows);
                resolve(res);
            }).catch(err => {
                reject(err);
            })
        })
    }
  }
}



function createHandler(row,setRows,setCurrentChoose,setCluster,setNamespace,setNamespaces) {
  return [
    chooseClusterHandler(row,setCurrentChoose,setCluster,setNamespace,setNamespaces),
    deleteClusterHandler(row,setRows),
  ]
}

const getClusterList = async (setRows) => {
  const res = await clusterHttp.listClusters();
  let items = res.map((item)=>{
    return {
      name:item.name,
      id: item.id,
      version: item.version,
      status: item.status
    }
  })
  setRows(items);
}

export default function Clusters() {
    const {setCluster,setNamespace,setNamespaces} = useContext(AppContext)
    const [rows, setRows] = useState([])
    const [openAdd, setOpenAdd] = useState(false)
    const [currentChoose, setCurrentChoose] = useState(null)

    useEffect(() => {
      getClusterList(setRows);
      setCurrentChoose(currentCluster());
    }, [currentChoose]);

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
        <TableContainer component={Paper} sx={{ maxHeight: 440 }}>
      <Table stickyHeader  sx={{ minWidth: 650 }} size="small">
        <TableHead>
          <TableRow >
            <TableCell sx={{ fontWeight: 'bold' }}>Name</TableCell>
            <TableCell sx={{ fontWeight: 'bold' }}>Version</TableCell>
            <TableCell sx={{ fontWeight: 'bold' }}>Status</TableCell>
            <TableCell></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow
              key={row.name}
              sx={{ '&:last-child td, &:last-child th': { border: 0 },
                backgroundColor: row.id === currentChoose ? '#0072E5' : 'white',
                
              }}
            >
              <TableCell 
                sx={{ color: row.id === currentChoose ? 'white' : 'black' }}
                component="th" scope="row">
                {row.name}
              </TableCell>
              <TableCell sx={{ color: row.id === currentChoose ? 'white' : 'black' }}>{row.version}</TableCell>
              <TableCell sx={{ color: row.id === currentChoose ? 'white' : 'black' }}>{row.status}</TableCell>
              <TableCell sx={{ color: row.id === currentChoose ? 'white' : 'black' }} align='right' width={1}>
                  <MoreVertBtn items={createHandler(row,setRows,setCurrentChoose,setCluster,setNamespace,setNamespaces)} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>

    <AddCluster open={openAdd} setOpen={setOpenAdd}/>
    </>
    )

}